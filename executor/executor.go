package executor

import (
	"context"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcmm "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	eabi "github.com/pancakeswap/pancake-statas/abi"
	"github.com/pancakeswap/pancake-statas/common"
	"github.com/pancakeswap/pancake-statas/util"
)

type Executor interface {
	GetBlockAndTxEvents(height int64) (*common.BlockAndEventLogs, error)
	GetPairList() []ethcmm.Address
}

type DecimalQuerier interface {
	GetDecimals(addr ethcmm.Address) (uint8, uint8, error)
}

type ChainExecutor struct {
	mux         sync.Mutex
	PairList    []ethcmm.Address
	SwapPairABI abi.ABI
	Client      *ethclient.Client
	Factory     string

	infoQuery DecimalQuerier
}

func NewExecutor(provider, factory string) *ChainExecutor {
	swapPairAbi, err := abi.JSON(strings.NewReader(eabi.SwappairABI))
	if err != nil {
		panic("marshal abi error")
	}
	client, err := ethclient.Dial(provider)
	if err != nil {
		panic("new eth client error")
	}
	factoryIns, err := eabi.NewFactory(ethcmm.HexToAddress(factory), client)
	if err != nil {
		panic(err)
	}
	pairLength, err := factoryIns.AllPairsLength(nil)
	if err != nil {
		panic(err)
	}
	pairList := make([]ethcmm.Address, 0, pairLength.Int64())
	for i := int64(0); i < pairLength.Int64(); i++ {
		pair, err := factoryIns.AllPairs(nil, big.NewInt(i))
		if err != nil {
			panic(err)
		}
		pairList = append(pairList, pair)
	}
	return &ChainExecutor{
		SwapPairABI: swapPairAbi,
		Client:      client,
		PairList:    pairList,
		Factory:     factory,
	}
}

func (e *ChainExecutor) SetInfoQuery(infoQuery DecimalQuerier) {
	e.infoQuery = infoQuery
}

func (e *ChainExecutor) Start() {
	go func() {
		for {
			time.Sleep(1000 * time.Second)
			factoryIns, err := eabi.NewFactory(ethcmm.HexToAddress(e.Factory), e.Client)
			if err != nil {
				panic(err)
			}
			pairLength, err := factoryIns.AllPairsLength(nil)
			if err != nil {
				panic(err)
			}
			pairList := make([]ethcmm.Address, 0, pairLength.Int64())
			for i := int64(0); i < pairLength.Int64(); i++ {
				pair, err := factoryIns.AllPairs(nil, big.NewInt(i))
				if err != nil {
					panic(err)
				}
				pairList = append(pairList, pair)
			}
			e.mux.Lock()
			e.PairList = pairList
			e.mux.Unlock()
		}
	}()
}

func (e *ChainExecutor) GetPairList() []ethcmm.Address {
	e.mux.Lock()
	defer e.mux.Unlock()
	return e.PairList
}

func (e *ChainExecutor) GetBlockAndTxEvents(height int64) (*common.BlockAndEventLogs, error) {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	header, err := e.Client.HeaderByNumber(ctxWithTimeout, big.NewInt(height))
	if err != nil {
		return nil, err
	}

	packageLogs, err := e.GetLogs(header)
	if err != nil {
		return nil, err
	}

	return &common.BlockAndEventLogs{
		Height:          height,
		BlockHash:       header.Hash().String(),
		ParentBlockHash: header.ParentHash.String(),
		BlockTime:       int64(header.Time),
		Events:          packageLogs,
	}, nil
}

func (e *ChainExecutor) GetLogs(header *types.Header) ([]interface{}, error) {
	topics := [][]ethcmm.Hash{{SwapEventHash}}

	blockHash := header.Hash()

	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	logs, err := e.Client.FilterLogs(ctxWithTimeout, ethereum.FilterQuery{
		BlockHash: &blockHash,
		Topics:    topics,
		Addresses: e.GetPairList(),
	})
	if err != nil {
		return nil, err
	}
	eventModels := make([]interface{}, 0)
	for _, log := range logs {
		event, err := ParseSwapEvent(&e.SwapPairABI, &log)
		if err != nil {
			util.Logger.Errorf("parse event log error, er=%s", err.Error())
			continue
		}

		if event == nil {
			continue
		}

		util.Logger.Infof("get log: %d, %s, %s", log.BlockNumber, log.Topics[0].String(), log.TxHash.String())
		d0, d1, err := e.infoQuery.GetDecimals(event.Contract)
		if err != nil {
			util.Logger.Error("Decimal can not found log: %d, %s, %s", log.BlockNumber, log.Topics[0].String(), log.TxHash.String())
			continue
		}
		eventModel := event.ToTxLog(&log, d0, d1)
		eventModel.BlockTime = int64(header.Time)
		eventModels = append(eventModels, eventModel)
	}
	return eventModels, nil
}

