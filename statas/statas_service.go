package statas

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/pancakeswap/pancake-statas/executor"
	"math/big"
	"sync"
	"time"

	"github.com/jinzhu/gorm"

	ethcmm "github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pancakeswap/pancake-statas/abi"
	"github.com/pancakeswap/pancake-statas/common"
	"github.com/pancakeswap/pancake-statas/model"
	"github.com/pancakeswap/pancake-statas/util"
)

const (
	QulifiedVolume = 100
)

var STABLE_TOKENS = map[string]bool{"BUSD": true}
var BASE_TOKENS = []string{"WBNB", "BUSD"}

type SwapPairInfo struct {
	SwapPairContract string  `json:"swap_pair_contract"`
	BaseSymbol       string  `json:"base_symbol"`
	QuoteSymbol      string  `json:"quote_symbol"`
	LastPrice        float64 `json:"last_price"`
	BaseVolume24h    float64 `json:"base_volume_24_h"`
	QuoteVolume24h   float64 `json:"quote_volume_24_h"`

	decimal0 uint8
	decimal1 uint8
	reserve0 float64
	reserve1 float64
}

type SyrupTVL struct {
	Name string  `json:"name"`
	Tvl  float64 `json:"tvl"`
}

type StatasSvc struct {
	mux       sync.Mutex
	statasDB  *gorm.DB
	config    *util.Config
	bscClient *ethclient.Client
	executor  executor.Executor

	updateAt        time.Time
	tokenPrice      map[string]float64
	totalVolume     float64
	totalLockVolume float64

	TVL        float64
	SyrupPools []SyrupTVL

	poolList        []ethcmm.Address
	swapPairList    []ethcmm.Address
	CertPairList    []ethcmm.Address
	swapPairInfoMap map[ethcmm.Address]*SwapPairInfo
	swapPairInfos   []SwapPairInfo
}

func NewStatasSvc(statasDB *gorm.DB, config *util.Config, executor executor.Executor) *StatasSvc {
	bscClient, err := ethclient.Dial(config.ChainConfig.BSCProvider)
	if err != nil {
		panic("new eth client error")
	}
	pairList := make([]ethcmm.Address, 0, len(config.ChainConfig.CertificatedPairs))
	for _, addr := range config.ChainConfig.CertificatedPairs {
		pairList = append(pairList, ethcmm.HexToAddress(addr))
	}
	poolList := make([]ethcmm.Address, 0, len(config.ChainConfig.SynupPools))
	for _, addr := range config.ChainConfig.SynupPools {
		poolList = append(poolList, ethcmm.HexToAddress(addr))
	}
	return &StatasSvc{
		statasDB:     statasDB,
		bscClient:    bscClient,
		CertPairList: pairList,
		poolList:     poolList,
		config:       config,
		executor:     executor,
	}
}

func (r *StatasSvc) refreshSwapPairs() []ethcmm.Address {
	totalSwapPairList := r.executor.GetPairList()
	certiMap := make(map[string]ethcmm.Address, 0)
	swapPairList := make([]ethcmm.Address, 0)
	for _, swapPairAddr := range r.CertPairList {
		pairInstance, err := abi.NewSwappair(swapPairAddr, r.bscClient)
		if err != nil {
			continue
		}
		token0, err := pairInstance.Token0(nil)
		if err != nil {
			continue
		}
		token1, err := pairInstance.Token1(nil)
		if err != nil {
			continue
		}
		token0Instance, err := abi.NewBep20(token0, r.bscClient)
		if err != nil {
			continue
		}
		token1Instance, err := abi.NewBep20(token1, r.bscClient)
		if err != nil {
			continue
		}
		symbol0, err := token0Instance.Symbol(nil)
		if err != nil {
			continue
		}

		symbol1, err := token1Instance.Symbol(nil)
		if err != nil {
			continue
		}
		certiMap[symbol0] = token0
		certiMap[symbol1] = token1
	}
	for _, swapPairAddr := range totalSwapPairList {
		pairInstance, err := abi.NewSwappair(swapPairAddr, r.bscClient)
		if err != nil {
			continue
		}
		token0, err := pairInstance.Token0(nil)
		if err != nil {
			continue
		}
		token1, err := pairInstance.Token1(nil)
		if err != nil {
			continue
		}
		token0Instance, err := abi.NewBep20(token0, r.bscClient)
		if err != nil {
			continue
		}
		token1Instance, err := abi.NewBep20(token1, r.bscClient)
		if err != nil {
			continue
		}
		symbol0, err := token0Instance.Symbol(nil)
		if err != nil {
			continue
		}

		symbol1, err := token1Instance.Symbol(nil)
		if err != nil {
			continue
		}
		if addr, exist := certiMap[symbol0]; exist && addr != token0 {
			continue
		}
		if addr, exist := certiMap[symbol1]; exist && addr != token1 {
			continue
		}
		swapPairList = append(swapPairList, swapPairAddr)
	}
	return swapPairList
}

func (r *StatasSvc) Start() {
	r.swapPairList = r.refreshSwapPairs()
	r.refreshSwapPairInfos()
	go r.refreshLoop()
	go r.refreshSwapPairsRoute()
}

func (r *StatasSvc) refreshSwapPairsRoute() {
	for {
		swapList := r.refreshSwapPairs()
		r.mux.Lock()
		r.swapPairList = swapList
		r.mux.Unlock()
		time.Sleep(1000 * time.Second)
	}
}

func (r *StatasSvc) refreshLoop() {
	tick := time.Tick(common.RefreshInterval)
	for {
		select {
		case <-tick:
			r.refreshSwapPairInfos()
		}
	}
}

func (r *StatasSvc) GetDecimals(addr ethcmm.Address) (uint8, uint8, error) {
	r.mux.Lock()
	defer r.mux.Unlock()
	info, exist := r.swapPairInfoMap[addr]
	if !exist {
		return 0, 0, fmt.Errorf("not exist")
	}
	return info.decimal0, info.decimal1, nil
}

func (r *StatasSvc) GetSwapPairInfos() ([]SwapPairInfo, float64, float64, time.Time) {
	r.mux.Lock()
	defer r.mux.Unlock()
	return r.swapPairInfos, r.totalVolume, r.totalLockVolume, r.updateAt
}

func (r *StatasSvc) GetPrice() (map[string]float64, time.Time) {
	r.mux.Lock()
	defer r.mux.Unlock()
	return r.tokenPrice, r.updateAt
}

func (r *StatasSvc) GetSynup() ([]SyrupTVL, float64, time.Time) {
	r.mux.Lock()
	defer r.mux.Unlock()
	return r.SyrupPools, r.TVL, r.updateAt
}

func (r *StatasSvc) refreshSwapPairInfos() {
	swapPairInfoMap := make(map[ethcmm.Address]*SwapPairInfo, 0)
	swapPairInfos := make([]SwapPairInfo, 0)
	symbols := make(map[string]bool, 0)
	tokenPrice := make(map[string]float64, 0)
	tokePriceMetrics := make(map[string]map[string]*PriceVolume, 0)
	for _, swapContract := range r.swapPairList {
		swapInfo, err := r.refreshSwapPairInfo(swapContract)
		if err!=nil{
			continue
		}
		if swapInfo.reserve0*swapInfo.reserve1 < 100 {
			continue
		}
		if err != nil {
			util.Logger.Errorf("refreshSwapPairInfo failed, err=%v, will retry refresh later", err)
			return
		}
		swapPairInfoMap[swapContract] = swapInfo
		if tokePriceMetrics[swapInfo.BaseSymbol] == nil {
			tokePriceMetrics[swapInfo.BaseSymbol] = make(map[string]*PriceVolume, 0)
		}
		if tokePriceMetrics[swapInfo.QuoteSymbol] == nil {
			tokePriceMetrics[swapInfo.QuoteSymbol] = make(map[string]*PriceVolume, 0)
		}
		tokePriceMetrics[swapInfo.BaseSymbol][swapInfo.QuoteSymbol] = &PriceVolume{Price: swapInfo.LastPrice}
		if swapInfo.LastPrice != 0 {
			tokePriceMetrics[swapInfo.QuoteSymbol][swapInfo.BaseSymbol] = &PriceVolume{Price: 1 / swapInfo.LastPrice}
		}
		symbols[swapInfo.BaseSymbol] = true
		symbols[swapInfo.QuoteSymbol] = true
	}

	for symbol := range symbols {
		if STABLE_TOKENS[symbol] {
			tokenPrice[symbol] = 1
		}
	}

	totalStatas, err := model.GetLast24HourTotalAccount(r.statasDB)
	if err != nil {
		util.Logger.Errorf("refreshSwapPairInfo failed, err=%v, will retry refresh later", err)
		return
	}
	for _, stata := range totalStatas {
		smartAddr := ethcmm.HexToAddress(stata.ContractAddress)
		swapInfo := swapPairInfoMap[smartAddr]
		if swapInfo == nil {
			continue
		}
		swapInfo.BaseVolume24h = stata.TotalAmount0
		swapInfo.QuoteVolume24h = stata.TotalAmount1
		if _, exist := tokePriceMetrics[swapInfo.BaseSymbol]; exist {
			tokePriceMetrics[swapInfo.BaseSymbol][swapInfo.QuoteSymbol].Volume = swapInfo.BaseVolume24h
		}
		if _, exist := tokePriceMetrics[swapInfo.QuoteSymbol]; exist {
			tokePriceMetrics[swapInfo.QuoteSymbol][swapInfo.BaseSymbol].Volume = swapInfo.QuoteVolume24h
		}
	}
	// to decrease the impact of low liquidity
	tokenPrice["WBNB"] = tokePriceMetrics["WBNB"]["BUSD"].Price

	for _, s := range BASE_TOKENS {
		p := tokenPrice[s]
		for os, op := range tokePriceMetrics[s] {
			if _, exist := tokenPrice[os]; !exist && op.Price != 0 && op.Volume*p > QulifiedVolume {
				tokenPrice[os] = p / op.Price
			}
		}
	}

	// for cake,
	s := "Cake"
	p := tokenPrice[s]
	for os, op := range tokePriceMetrics[s] {
		if _, exist := tokenPrice[os]; !exist && op.Price != 0 && op.Volume*p > QulifiedVolume {
			tokenPrice[os] = p / op.Price
		}
	}

	var totalVolume, totalLock float64
	for _, swapInfo := range swapPairInfoMap {
		var swapPairVolume, swapLock float64
		if price, exist := tokenPrice[swapInfo.BaseSymbol]; exist {
			swapPairVolume = swapPairVolume + swapInfo.BaseVolume24h*price
			swapLock = swapLock + swapInfo.reserve0*price
		}
		if price, exist := tokenPrice[swapInfo.QuoteSymbol]; exist {
			swapPairVolume = swapPairVolume + swapInfo.QuoteVolume24h*price
			swapLock = swapLock + swapInfo.reserve1*price
		}
		if swapPairVolume >= QulifiedVolume {
			totalVolume = totalVolume + swapPairVolume
			totalLock = totalLock + swapLock
			swapPairInfos = append(swapPairInfos, *swapInfo)
		}
	}

	cakeIns, err := abi.NewBep20(ethcmm.HexToAddress("0x0E09FaBB73Bd3Ade0a17ECC321fD13a19e81cE82"), r.bscClient)
	if err != nil {
		util.Logger.Errorf("failed to init cake Ins", err)
		return
	}
	syrupPools := make([]SyrupTVL, 0)
	totalSynupTvl := float64(0)
	for idx, addr := range r.poolList {
		poolIns, err := abi.NewSmartchef(addr, r.bscClient)
		if err != nil {
			util.Logger.Errorf("failed to init poolIns Ins %v, %s", err, addr.String())
			continue
		}
		var name string
		if idx == 0 {
			name = "Cake"
		} else {
			rewardToken, err := poolIns.RewardToken(nil)
			if err != nil {
				util.Logger.Errorf("failed to init rewardToken Ins %v, %s", err, addr.String())
				continue
			}
			rewardTokenIns, err := abi.NewBep20(rewardToken, r.bscClient)
			if err != nil {
				util.Logger.Errorf("failed to init rewardTokenIns Ins %v, %s", err, addr.String())
				continue
			}
			name, err = rewardTokenIns.Name(nil)
			if err != nil {
				continue
			}
		}
		balance, err := cakeIns.BalanceOf(nil, addr)
		if err != nil {
			util.Logger.Errorf("failed to get cake balance Ins %v, %s", err, addr.String())
			continue
		}
		cakePrice := tokenPrice["Cake"]
		tvl := float64(new(big.Int).Div(balance, big.NewInt(1e18)).Int64()) * cakePrice
		syrupPools = append(syrupPools, SyrupTVL{
			Name: name,
			Tvl:  tvl,
		})
		totalSynupTvl += tvl
	}

	r.mux.Lock()
	r.TVL = totalSynupTvl
	r.SyrupPools = syrupPools
	r.tokenPrice = tokenPrice
	r.swapPairInfoMap = swapPairInfoMap
	r.swapPairInfos = swapPairInfos
	r.totalVolume = totalVolume
	r.totalLockVolume = totalLock
	r.updateAt = time.Now()
	r.mux.Unlock()
}

func (r *StatasSvc) refreshSwapPairInfo(swapPairAddr ethcmm.Address) (*SwapPairInfo, error) {
	pairInstance, err := abi.NewSwappair(swapPairAddr, r.bscClient)
	if err != nil {
		return nil, err
	}
	reserve, err := pairInstance.GetReserves(nil)
	if err != nil {
		return nil, err
	}

	token0, err := pairInstance.Token0(nil)
	if err != nil {
		return nil, err
	}
	token1, err := pairInstance.Token1(nil)
	if err != nil {
		return nil, err
	}
	token0Instance, err := abi.NewBep20(token0, r.bscClient)
	if err != nil {
		return nil, err
	}
	token1Instance, err := abi.NewBep20(token1, r.bscClient)
	if err != nil {
		return nil, err
	}
	symbol0, err := token0Instance.Symbol(nil)
	if err != nil {
		return nil, err
	}

	symbol1, err := token1Instance.Symbol(nil)
	if err != nil {
		return nil, err
	}

	decimal0, err := token0Instance.Decimals(nil)
	if err != nil {
		return nil, err
	}
	decimal1, err := token1Instance.Decimals(nil)

	if err != nil {
		return nil, err
	}

	reserve0, _ := new(big.Float).Quo(new(big.Float).SetInt(reserve.Reserve0), new(big.Float).SetInt(math.Exp(big.NewInt(10), big.NewInt(int64(decimal0))))).Float64()
	reserve1, _ := new(big.Float).Quo(new(big.Float).SetInt(reserve.Reserve1), new(big.Float).SetInt(math.Exp(big.NewInt(10), big.NewInt(int64(decimal1))))).Float64()

	var price float64
	if reserve.Reserve0.Cmp(new(big.Int).SetInt64(0)) != 0 {
		price = reserve1 / reserve0
	}

	return &SwapPairInfo{
		SwapPairContract: swapPairAddr.String(),
		BaseSymbol:       symbol0,
		QuoteSymbol:      symbol1,
		LastPrice:        price,
		decimal0:         decimal0,
		decimal1:         decimal1,
		reserve0:         reserve0,
		reserve1:         reserve1,
	}, nil
}

type PriceVolume struct {
	Price  float64
	Volume float64
}
