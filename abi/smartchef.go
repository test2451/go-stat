// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SmartchefABI is the input ABI used to generate the binding from.
const SmartchefABI = "[{\"inputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"_syrup\",\"type\":\"address\"},{\"internalType\":\"contractIBEP20\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rewardPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bonusEndBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EmergencyWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bonusEndBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"emergencyRewardWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_to\",\"type\":\"uint256\"}],\"name\":\"getMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"massUpdatePools\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"pendingReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolInfo\",\"outputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"lpToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accCakePerShare\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardToken\",\"outputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stopReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"syrup\",\"outputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"updatePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDebt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Smartchef is an auto generated Go binding around an Ethereum contract.
type Smartchef struct {
	SmartchefCaller     // Read-only binding to the contract
	SmartchefTransactor // Write-only binding to the contract
	SmartchefFilterer   // Log filterer for contract events
}

// SmartchefCaller is an auto generated read-only Go binding around an Ethereum contract.
type SmartchefCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartchefTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SmartchefTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartchefFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SmartchefFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartchefSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SmartchefSession struct {
	Contract     *Smartchef        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SmartchefCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SmartchefCallerSession struct {
	Contract *SmartchefCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SmartchefTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SmartchefTransactorSession struct {
	Contract     *SmartchefTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SmartchefRaw is an auto generated low-level Go binding around an Ethereum contract.
type SmartchefRaw struct {
	Contract *Smartchef // Generic contract binding to access the raw methods on
}

// SmartchefCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SmartchefCallerRaw struct {
	Contract *SmartchefCaller // Generic read-only contract binding to access the raw methods on
}

// SmartchefTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SmartchefTransactorRaw struct {
	Contract *SmartchefTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSmartchef creates a new instance of Smartchef, bound to a specific deployed contract.
func NewSmartchef(address common.Address, backend bind.ContractBackend) (*Smartchef, error) {
	contract, err := bindSmartchef(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Smartchef{SmartchefCaller: SmartchefCaller{contract: contract}, SmartchefTransactor: SmartchefTransactor{contract: contract}, SmartchefFilterer: SmartchefFilterer{contract: contract}}, nil
}

// NewSmartchefCaller creates a new read-only instance of Smartchef, bound to a specific deployed contract.
func NewSmartchefCaller(address common.Address, caller bind.ContractCaller) (*SmartchefCaller, error) {
	contract, err := bindSmartchef(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SmartchefCaller{contract: contract}, nil
}

// NewSmartchefTransactor creates a new write-only instance of Smartchef, bound to a specific deployed contract.
func NewSmartchefTransactor(address common.Address, transactor bind.ContractTransactor) (*SmartchefTransactor, error) {
	contract, err := bindSmartchef(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SmartchefTransactor{contract: contract}, nil
}

// NewSmartchefFilterer creates a new log filterer instance of Smartchef, bound to a specific deployed contract.
func NewSmartchefFilterer(address common.Address, filterer bind.ContractFilterer) (*SmartchefFilterer, error) {
	contract, err := bindSmartchef(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SmartchefFilterer{contract: contract}, nil
}

// bindSmartchef binds a generic wrapper to an already deployed contract.
func bindSmartchef(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SmartchefABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Smartchef *SmartchefRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Smartchef.Contract.SmartchefCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Smartchef *SmartchefRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Smartchef.Contract.SmartchefTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Smartchef *SmartchefRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Smartchef.Contract.SmartchefTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Smartchef *SmartchefCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Smartchef.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Smartchef *SmartchefTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Smartchef.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Smartchef *SmartchefTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Smartchef.Contract.contract.Transact(opts, method, params...)
}

// BonusEndBlock is a free data retrieval call binding the contract method 0x1aed6553.
//
// Solidity: function bonusEndBlock() constant returns(uint256)
func (_Smartchef *SmartchefCaller) BonusEndBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Smartchef.contract.Call(opts, out, "bonusEndBlock")
	return *ret0, err
}

// BonusEndBlock is a free data retrieval call binding the contract method 0x1aed6553.
//
// Solidity: function bonusEndBlock() constant returns(uint256)
func (_Smartchef *SmartchefSession) BonusEndBlock() (*big.Int, error) {
	return _Smartchef.Contract.BonusEndBlock(&_Smartchef.CallOpts)
}

// BonusEndBlock is a free data retrieval call binding the contract method 0x1aed6553.
//
// Solidity: function bonusEndBlock() constant returns(uint256)
func (_Smartchef *SmartchefCallerSession) BonusEndBlock() (*big.Int, error) {
	return _Smartchef.Contract.BonusEndBlock(&_Smartchef.CallOpts)
}

// GetMultiplier is a free data retrieval call binding the contract method 0x8dbb1e3a.
//
// Solidity: function getMultiplier(uint256 _from, uint256 _to) constant returns(uint256)
func (_Smartchef *SmartchefCaller) GetMultiplier(opts *bind.CallOpts, _from *big.Int, _to *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Smartchef.contract.Call(opts, out, "getMultiplier", _from, _to)
	return *ret0, err
}

// GetMultiplier is a free data retrieval call binding the contract method 0x8dbb1e3a.
//
// Solidity: function getMultiplier(uint256 _from, uint256 _to) constant returns(uint256)
func (_Smartchef *SmartchefSession) GetMultiplier(_from *big.Int, _to *big.Int) (*big.Int, error) {
	return _Smartchef.Contract.GetMultiplier(&_Smartchef.CallOpts, _from, _to)
}

// GetMultiplier is a free data retrieval call binding the contract method 0x8dbb1e3a.
//
// Solidity: function getMultiplier(uint256 _from, uint256 _to) constant returns(uint256)
func (_Smartchef *SmartchefCallerSession) GetMultiplier(_from *big.Int, _to *big.Int) (*big.Int, error) {
	return _Smartchef.Contract.GetMultiplier(&_Smartchef.CallOpts, _from, _to)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Smartchef *SmartchefCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Smartchef.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Smartchef *SmartchefSession) Owner() (common.Address, error) {
	return _Smartchef.Contract.Owner(&_Smartchef.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Smartchef *SmartchefCallerSession) Owner() (common.Address, error) {
	return _Smartchef.Contract.Owner(&_Smartchef.CallOpts)
}

// PendingReward is a free data retrieval call binding the contract method 0xf40f0f52.
//
// Solidity: function pendingReward(address _user) constant returns(uint256)
func (_Smartchef *SmartchefCaller) PendingReward(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Smartchef.contract.Call(opts, out, "pendingReward", _user)
	return *ret0, err
}

// PendingReward is a free data retrieval call binding the contract method 0xf40f0f52.
//
// Solidity: function pendingReward(address _user) constant returns(uint256)
func (_Smartchef *SmartchefSession) PendingReward(_user common.Address) (*big.Int, error) {
	return _Smartchef.Contract.PendingReward(&_Smartchef.CallOpts, _user)
}

// PendingReward is a free data retrieval call binding the contract method 0xf40f0f52.
//
// Solidity: function pendingReward(address _user) constant returns(uint256)
func (_Smartchef *SmartchefCallerSession) PendingReward(_user common.Address) (*big.Int, error) {
	return _Smartchef.Contract.PendingReward(&_Smartchef.CallOpts, _user)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) constant returns(address lpToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accCakePerShare)
func (_Smartchef *SmartchefCaller) PoolInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	LpToken         common.Address
	AllocPoint      *big.Int
	LastRewardBlock *big.Int
	AccCakePerShare *big.Int
}, error) {
	ret := new(struct {
		LpToken         common.Address
		AllocPoint      *big.Int
		LastRewardBlock *big.Int
		AccCakePerShare *big.Int
	})
	out := ret
	err := _Smartchef.contract.Call(opts, out, "poolInfo", arg0)
	return *ret, err
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) constant returns(address lpToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accCakePerShare)
func (_Smartchef *SmartchefSession) PoolInfo(arg0 *big.Int) (struct {
	LpToken         common.Address
	AllocPoint      *big.Int
	LastRewardBlock *big.Int
	AccCakePerShare *big.Int
}, error) {
	return _Smartchef.Contract.PoolInfo(&_Smartchef.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) constant returns(address lpToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accCakePerShare)
func (_Smartchef *SmartchefCallerSession) PoolInfo(arg0 *big.Int) (struct {
	LpToken         common.Address
	AllocPoint      *big.Int
	LastRewardBlock *big.Int
	AccCakePerShare *big.Int
}, error) {
	return _Smartchef.Contract.PoolInfo(&_Smartchef.CallOpts, arg0)
}

// RewardPerBlock is a free data retrieval call binding the contract method 0x8ae39cac.
//
// Solidity: function rewardPerBlock() constant returns(uint256)
func (_Smartchef *SmartchefCaller) RewardPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Smartchef.contract.Call(opts, out, "rewardPerBlock")
	return *ret0, err
}

// RewardPerBlock is a free data retrieval call binding the contract method 0x8ae39cac.
//
// Solidity: function rewardPerBlock() constant returns(uint256)
func (_Smartchef *SmartchefSession) RewardPerBlock() (*big.Int, error) {
	return _Smartchef.Contract.RewardPerBlock(&_Smartchef.CallOpts)
}

// RewardPerBlock is a free data retrieval call binding the contract method 0x8ae39cac.
//
// Solidity: function rewardPerBlock() constant returns(uint256)
func (_Smartchef *SmartchefCallerSession) RewardPerBlock() (*big.Int, error) {
	return _Smartchef.Contract.RewardPerBlock(&_Smartchef.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() constant returns(address)
func (_Smartchef *SmartchefCaller) RewardToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Smartchef.contract.Call(opts, out, "rewardToken")
	return *ret0, err
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() constant returns(address)
func (_Smartchef *SmartchefSession) RewardToken() (common.Address, error) {
	return _Smartchef.Contract.RewardToken(&_Smartchef.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() constant returns(address)
func (_Smartchef *SmartchefCallerSession) RewardToken() (common.Address, error) {
	return _Smartchef.Contract.RewardToken(&_Smartchef.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() constant returns(uint256)
func (_Smartchef *SmartchefCaller) StartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Smartchef.contract.Call(opts, out, "startBlock")
	return *ret0, err
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() constant returns(uint256)
func (_Smartchef *SmartchefSession) StartBlock() (*big.Int, error) {
	return _Smartchef.Contract.StartBlock(&_Smartchef.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() constant returns(uint256)
func (_Smartchef *SmartchefCallerSession) StartBlock() (*big.Int, error) {
	return _Smartchef.Contract.StartBlock(&_Smartchef.CallOpts)
}

// Syrup is a free data retrieval call binding the contract method 0x86a952c4.
//
// Solidity: function syrup() constant returns(address)
func (_Smartchef *SmartchefCaller) Syrup(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Smartchef.contract.Call(opts, out, "syrup")
	return *ret0, err
}

// Syrup is a free data retrieval call binding the contract method 0x86a952c4.
//
// Solidity: function syrup() constant returns(address)
func (_Smartchef *SmartchefSession) Syrup() (common.Address, error) {
	return _Smartchef.Contract.Syrup(&_Smartchef.CallOpts)
}

// Syrup is a free data retrieval call binding the contract method 0x86a952c4.
//
// Solidity: function syrup() constant returns(address)
func (_Smartchef *SmartchefCallerSession) Syrup() (common.Address, error) {
	return _Smartchef.Contract.Syrup(&_Smartchef.CallOpts)
}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) constant returns(uint256 amount, uint256 rewardDebt)
func (_Smartchef *SmartchefCaller) UserInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	ret := new(struct {
		Amount     *big.Int
		RewardDebt *big.Int
	})
	out := ret
	err := _Smartchef.contract.Call(opts, out, "userInfo", arg0)
	return *ret, err
}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) constant returns(uint256 amount, uint256 rewardDebt)
func (_Smartchef *SmartchefSession) UserInfo(arg0 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _Smartchef.Contract.UserInfo(&_Smartchef.CallOpts, arg0)
}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) constant returns(uint256 amount, uint256 rewardDebt)
func (_Smartchef *SmartchefCallerSession) UserInfo(arg0 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _Smartchef.Contract.UserInfo(&_Smartchef.CallOpts, arg0)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_Smartchef *SmartchefTransactor) Deposit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Smartchef.contract.Transact(opts, "deposit", _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_Smartchef *SmartchefSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _Smartchef.Contract.Deposit(&_Smartchef.TransactOpts, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_Smartchef *SmartchefTransactorSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _Smartchef.Contract.Deposit(&_Smartchef.TransactOpts, _amount)
}

// EmergencyRewardWithdraw is a paid mutator transaction binding the contract method 0x3279beab.
//
// Solidity: function emergencyRewardWithdraw(uint256 _amount) returns()
func (_Smartchef *SmartchefTransactor) EmergencyRewardWithdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Smartchef.contract.Transact(opts, "emergencyRewardWithdraw", _amount)
}

// EmergencyRewardWithdraw is a paid mutator transaction binding the contract method 0x3279beab.
//
// Solidity: function emergencyRewardWithdraw(uint256 _amount) returns()
func (_Smartchef *SmartchefSession) EmergencyRewardWithdraw(_amount *big.Int) (*types.Transaction, error) {
	return _Smartchef.Contract.EmergencyRewardWithdraw(&_Smartchef.TransactOpts, _amount)
}

// EmergencyRewardWithdraw is a paid mutator transaction binding the contract method 0x3279beab.
//
// Solidity: function emergencyRewardWithdraw(uint256 _amount) returns()
func (_Smartchef *SmartchefTransactorSession) EmergencyRewardWithdraw(_amount *big.Int) (*types.Transaction, error) {
	return _Smartchef.Contract.EmergencyRewardWithdraw(&_Smartchef.TransactOpts, _amount)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_Smartchef *SmartchefTransactor) EmergencyWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Smartchef.contract.Transact(opts, "emergencyWithdraw")
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_Smartchef *SmartchefSession) EmergencyWithdraw() (*types.Transaction, error) {
	return _Smartchef.Contract.EmergencyWithdraw(&_Smartchef.TransactOpts)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_Smartchef *SmartchefTransactorSession) EmergencyWithdraw() (*types.Transaction, error) {
	return _Smartchef.Contract.EmergencyWithdraw(&_Smartchef.TransactOpts)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_Smartchef *SmartchefTransactor) MassUpdatePools(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Smartchef.contract.Transact(opts, "massUpdatePools")
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_Smartchef *SmartchefSession) MassUpdatePools() (*types.Transaction, error) {
	return _Smartchef.Contract.MassUpdatePools(&_Smartchef.TransactOpts)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_Smartchef *SmartchefTransactorSession) MassUpdatePools() (*types.Transaction, error) {
	return _Smartchef.Contract.MassUpdatePools(&_Smartchef.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Smartchef *SmartchefTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Smartchef.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Smartchef *SmartchefSession) RenounceOwnership() (*types.Transaction, error) {
	return _Smartchef.Contract.RenounceOwnership(&_Smartchef.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Smartchef *SmartchefTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Smartchef.Contract.RenounceOwnership(&_Smartchef.TransactOpts)
}

// StopReward is a paid mutator transaction binding the contract method 0x80dc0672.
//
// Solidity: function stopReward() returns()
func (_Smartchef *SmartchefTransactor) StopReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Smartchef.contract.Transact(opts, "stopReward")
}

// StopReward is a paid mutator transaction binding the contract method 0x80dc0672.
//
// Solidity: function stopReward() returns()
func (_Smartchef *SmartchefSession) StopReward() (*types.Transaction, error) {
	return _Smartchef.Contract.StopReward(&_Smartchef.TransactOpts)
}

// StopReward is a paid mutator transaction binding the contract method 0x80dc0672.
//
// Solidity: function stopReward() returns()
func (_Smartchef *SmartchefTransactorSession) StopReward() (*types.Transaction, error) {
	return _Smartchef.Contract.StopReward(&_Smartchef.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Smartchef *SmartchefTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Smartchef.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Smartchef *SmartchefSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Smartchef.Contract.TransferOwnership(&_Smartchef.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Smartchef *SmartchefTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Smartchef.Contract.TransferOwnership(&_Smartchef.TransactOpts, newOwner)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_Smartchef *SmartchefTransactor) UpdatePool(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _Smartchef.contract.Transact(opts, "updatePool", _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_Smartchef *SmartchefSession) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _Smartchef.Contract.UpdatePool(&_Smartchef.TransactOpts, _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_Smartchef *SmartchefTransactorSession) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _Smartchef.Contract.UpdatePool(&_Smartchef.TransactOpts, _pid)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_Smartchef *SmartchefTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Smartchef.contract.Transact(opts, "withdraw", _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_Smartchef *SmartchefSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _Smartchef.Contract.Withdraw(&_Smartchef.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_Smartchef *SmartchefTransactorSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _Smartchef.Contract.Withdraw(&_Smartchef.TransactOpts, _amount)
}

// SmartchefDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Smartchef contract.
type SmartchefDepositIterator struct {
	Event *SmartchefDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SmartchefDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartchefDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SmartchefDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SmartchefDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartchefDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartchefDeposit represents a Deposit event raised by the Smartchef contract.
type SmartchefDeposit struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed user, uint256 amount)
func (_Smartchef *SmartchefFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address) (*SmartchefDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Smartchef.contract.FilterLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return &SmartchefDepositIterator{contract: _Smartchef.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed user, uint256 amount)
func (_Smartchef *SmartchefFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *SmartchefDeposit, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Smartchef.contract.WatchLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartchefDeposit)
				if err := _Smartchef.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed user, uint256 amount)
func (_Smartchef *SmartchefFilterer) ParseDeposit(log types.Log) (*SmartchefDeposit, error) {
	event := new(SmartchefDeposit)
	if err := _Smartchef.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SmartchefEmergencyWithdrawIterator is returned from FilterEmergencyWithdraw and is used to iterate over the raw logs and unpacked data for EmergencyWithdraw events raised by the Smartchef contract.
type SmartchefEmergencyWithdrawIterator struct {
	Event *SmartchefEmergencyWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SmartchefEmergencyWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartchefEmergencyWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SmartchefEmergencyWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SmartchefEmergencyWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartchefEmergencyWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartchefEmergencyWithdraw represents a EmergencyWithdraw event raised by the Smartchef contract.
type SmartchefEmergencyWithdraw struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyWithdraw is a free log retrieval operation binding the contract event 0x5fafa99d0643513820be26656b45130b01e1c03062e1266bf36f88cbd3bd9695.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 amount)
func (_Smartchef *SmartchefFilterer) FilterEmergencyWithdraw(opts *bind.FilterOpts, user []common.Address) (*SmartchefEmergencyWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Smartchef.contract.FilterLogs(opts, "EmergencyWithdraw", userRule)
	if err != nil {
		return nil, err
	}
	return &SmartchefEmergencyWithdrawIterator{contract: _Smartchef.contract, event: "EmergencyWithdraw", logs: logs, sub: sub}, nil
}

// WatchEmergencyWithdraw is a free log subscription operation binding the contract event 0x5fafa99d0643513820be26656b45130b01e1c03062e1266bf36f88cbd3bd9695.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 amount)
func (_Smartchef *SmartchefFilterer) WatchEmergencyWithdraw(opts *bind.WatchOpts, sink chan<- *SmartchefEmergencyWithdraw, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Smartchef.contract.WatchLogs(opts, "EmergencyWithdraw", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartchefEmergencyWithdraw)
				if err := _Smartchef.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEmergencyWithdraw is a log parse operation binding the contract event 0x5fafa99d0643513820be26656b45130b01e1c03062e1266bf36f88cbd3bd9695.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 amount)
func (_Smartchef *SmartchefFilterer) ParseEmergencyWithdraw(log types.Log) (*SmartchefEmergencyWithdraw, error) {
	event := new(SmartchefEmergencyWithdraw)
	if err := _Smartchef.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SmartchefOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Smartchef contract.
type SmartchefOwnershipTransferredIterator struct {
	Event *SmartchefOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SmartchefOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartchefOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SmartchefOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SmartchefOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartchefOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartchefOwnershipTransferred represents a OwnershipTransferred event raised by the Smartchef contract.
type SmartchefOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Smartchef *SmartchefFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SmartchefOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Smartchef.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SmartchefOwnershipTransferredIterator{contract: _Smartchef.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Smartchef *SmartchefFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SmartchefOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Smartchef.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartchefOwnershipTransferred)
				if err := _Smartchef.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Smartchef *SmartchefFilterer) ParseOwnershipTransferred(log types.Log) (*SmartchefOwnershipTransferred, error) {
	event := new(SmartchefOwnershipTransferred)
	if err := _Smartchef.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SmartchefWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Smartchef contract.
type SmartchefWithdrawIterator struct {
	Event *SmartchefWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SmartchefWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartchefWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SmartchefWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SmartchefWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartchefWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartchefWithdraw represents a Withdraw event raised by the Smartchef contract.
type SmartchefWithdraw struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 amount)
func (_Smartchef *SmartchefFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address) (*SmartchefWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Smartchef.contract.FilterLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return &SmartchefWithdrawIterator{contract: _Smartchef.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 amount)
func (_Smartchef *SmartchefFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *SmartchefWithdraw, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Smartchef.contract.WatchLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartchefWithdraw)
				if err := _Smartchef.contract.UnpackLog(event, "Withdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 amount)
func (_Smartchef *SmartchefFilterer) ParseWithdraw(log types.Log) (*SmartchefWithdraw, error) {
	event := new(SmartchefWithdraw)
	if err := _Smartchef.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}
