package executor

import (
	"github.com/ethereum/go-ethereum/common/math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/pancakeswap/pancake-statas/model"
)

var (
	SwapEventName = "Swap"
	SwapEventHash = common.HexToHash("0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822")

	defaultDecimal = new(big.Float).SetInt64(1e18)
)

type SwapEvent struct {
	Contract   common.Address
	Sender     common.Address
	To         common.Address
	Amount0In  *big.Int
	Amount1In  *big.Int
	Amount0Out *big.Int
	Amount1Out *big.Int
}

func (ev *SwapEvent) ToTxLog(log *types.Log, decimal0, decimal1 uint8) *model.TxEventLog {
	var amount0, amount1 float64
	d0 := new(big.Float).SetInt(math.Exp(big.NewInt(10), big.NewInt(int64(decimal0))))
	d1 := new(big.Float).SetInt(math.Exp(big.NewInt(10), big.NewInt(int64(decimal1))))


	if ev.Amount0In.Cmp(ev.Amount0Out) > 0 {
		amount0, _ = new(big.Float).Quo(new(big.Float).SetInt(new(big.Int).Sub(ev.Amount0In, ev.Amount0Out)), d0).Float64()
	} else {
		amount0, _ = new(big.Float).Quo(new(big.Float).SetInt(new(big.Int).Sub(ev.Amount0Out, ev.Amount0In)), d0).Float64()
	}
	if ev.Amount1In.Cmp(ev.Amount1Out) > 0 {
		amount1, _ = new(big.Float).Quo(new(big.Float).SetInt(new(big.Int).Sub(ev.Amount1In, ev.Amount1Out)), d1).Float64()
	} else {
		amount1, _ = new(big.Float).Quo(new(big.Float).SetInt(new(big.Int).Sub(ev.Amount1Out, ev.Amount1In)), d1).Float64()
	}
	pack := &model.TxEventLog{
		ContractAddress: ev.Contract.String(),
		Amount0:         amount0,
		Amount1:         amount1,
		BlockHash:       log.BlockHash.Hex(),
		TxHash:          log.TxHash.String(),
		Height:          int64(log.BlockNumber),
	}
	return pack
}

func ParseSwapEvent(abi *abi.ABI, log *types.Log) (*SwapEvent, error) {
	var ev SwapEvent

	err := abi.Unpack(&ev, SwapEventName, log.Data)
	if err != nil {
		return nil, err
	}

	ev.Sender = common.BytesToAddress(log.Topics[1].Bytes())
	ev.To = common.BytesToAddress(log.Topics[2].Bytes())
	ev.Contract = log.Address

	return &ev, nil
}

