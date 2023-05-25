// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package forge

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ForgeMetaData contains all meta data concerning the Forge contract.
var ForgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIForgeAsset\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_depositRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_liquidationThreshold\",\"type\":\"uint256\"}],\"name\":\"addAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"close\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_currency\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_currencyDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_minCollateral\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Close\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"edit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Edit\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"liquidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"Liquidate\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"open\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"token\",\"type\":\"uint256\"}],\"name\":\"Open\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"pauseAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resume\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"resumeAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"setAssetOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_collector\",\"type\":\"address\"}],\"name\":\"setCollector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_feeLiquidatorPercent\",\"type\":\"uint256\"}],\"name\":\"setFeeLiquidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minCollateral\",\"type\":\"uint256\"}],\"name\":\"setMinCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pauser\",\"type\":\"address\"}],\"name\":\"setPauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"assets\",\"outputs\":[{\"internalType\":\"contractIForgeAsset\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"oracleDecimals\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"depositRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationThreshold\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cRatioDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currency\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currencyDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeLiquidatorPercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeLiquidatorPercentBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getOvenRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ratio\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getOvensLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"numOvens\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUsersLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"numUsers\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isUser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ovens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"token\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"users\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ForgeABI is the input ABI used to generate the binding from.
// Deprecated: Use ForgeMetaData.ABI instead.
var ForgeABI = ForgeMetaData.ABI

// Forge is an auto generated Go binding around an Ethereum contract.
type Forge struct {
	ForgeCaller     // Read-only binding to the contract
	ForgeTransactor // Write-only binding to the contract
	ForgeFilterer   // Log filterer for contract events
}

// ForgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ForgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ForgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ForgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ForgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ForgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ForgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ForgeSession struct {
	Contract     *Forge            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ForgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ForgeCallerSession struct {
	Contract *ForgeCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ForgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ForgeTransactorSession struct {
	Contract     *ForgeTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ForgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ForgeRaw struct {
	Contract *Forge // Generic contract binding to access the raw methods on
}

// ForgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ForgeCallerRaw struct {
	Contract *ForgeCaller // Generic read-only contract binding to access the raw methods on
}

// ForgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ForgeTransactorRaw struct {
	Contract *ForgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewForge creates a new instance of Forge, bound to a specific deployed contract.
func NewForge(address common.Address, backend bind.ContractBackend) (*Forge, error) {
	contract, err := bindForge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Forge{ForgeCaller: ForgeCaller{contract: contract}, ForgeTransactor: ForgeTransactor{contract: contract}, ForgeFilterer: ForgeFilterer{contract: contract}}, nil
}

// NewForgeCaller creates a new read-only instance of Forge, bound to a specific deployed contract.
func NewForgeCaller(address common.Address, caller bind.ContractCaller) (*ForgeCaller, error) {
	contract, err := bindForge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ForgeCaller{contract: contract}, nil
}

// NewForgeTransactor creates a new write-only instance of Forge, bound to a specific deployed contract.
func NewForgeTransactor(address common.Address, transactor bind.ContractTransactor) (*ForgeTransactor, error) {
	contract, err := bindForge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ForgeTransactor{contract: contract}, nil
}

// NewForgeFilterer creates a new log filterer instance of Forge, bound to a specific deployed contract.
func NewForgeFilterer(address common.Address, filterer bind.ContractFilterer) (*ForgeFilterer, error) {
	contract, err := bindForge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ForgeFilterer{contract: contract}, nil
}

// bindForge binds a generic wrapper to an already deployed contract.
func bindForge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ForgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Forge *ForgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Forge.Contract.ForgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Forge *ForgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Forge.Contract.ForgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Forge *ForgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Forge.Contract.ForgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Forge *ForgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Forge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Forge *ForgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Forge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Forge *ForgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Forge.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Forge *ForgeCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Forge.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Forge *ForgeSession) Admin() (common.Address, error) {
	return _Forge.Contract.Admin(&_Forge.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Forge *ForgeCallerSession) Admin() (common.Address, error) {
	return _Forge.Contract.Admin(&_Forge.CallOpts)
}

// AllPaused is a free data retrieval call binding the contract method 0x697a80b7.
//
// Solidity: function allPaused() view returns(bool)
func (_Forge *ForgeCaller) AllPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Forge.contract.Call(opts, &out, "allPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllPaused is a free data retrieval call binding the contract method 0x697a80b7.
//
// Solidity: function allPaused() view returns(bool)
func (_Forge *ForgeSession) AllPaused() (bool, error) {
	return _Forge.Contract.AllPaused(&_Forge.CallOpts)
}

// AllPaused is a free data retrieval call binding the contract method 0x697a80b7.
//
// Solidity: function allPaused() view returns(bool)
func (_Forge *ForgeCallerSession) AllPaused() (bool, error) {
	return _Forge.Contract.AllPaused(&_Forge.CallOpts)
}

// Assets is a free data retrieval call binding the contract method 0xcf35bdd0.
//
// Solidity: function assets(uint256 ) view returns(address token, uint8 decimals, address oracle, uint8 oracleDecimals, bool paused, uint256 depositRatio, uint256 liquidationThreshold)
func (_Forge *ForgeCaller) Assets(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Token                common.Address
	Decimals             uint8
	Oracle               common.Address
	OracleDecimals       uint8
	Paused               bool
	DepositRatio         *big.Int
	LiquidationThreshold *big.Int
}, error) {
	var out []interface{}
	err := _Forge.contract.Call(opts, &out, "assets", arg0)

	outstruct := new(struct {
		Token                common.Address
		Decimals             uint8
		Oracle               common.Address
		OracleDecimals       uint8
		Paused               bool
		DepositRatio         *big.Int
		LiquidationThreshold *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Token = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Decimals = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.Oracle = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.OracleDecimals = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.Paused = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.DepositRatio = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.LiquidationThreshold = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Assets is a free data retrieval call binding the contract method 0xcf35bdd0.
//
// Solidity: function assets(uint256 ) view returns(address token, uint8 decimals, address oracle, uint8 oracleDecimals, bool paused, uint256 depositRatio, uint256 liquidationThreshold)
func (_Forge *ForgeSession) Assets(arg0 *big.Int) (struct {
	Token                common.Address
	Decimals             uint8
	Oracle               common.Address
	OracleDecimals       uint8
	Paused               bool
	DepositRatio         *big.Int
	LiquidationThreshold *big.Int
}, error) {
	return _Forge.Contract.Assets(&_Forge.CallOpts, arg0)
}

// Assets is a free data retrieval call binding the contract method 0xcf35bdd0.
//
// Solidity: function assets(uint256 ) view returns(address token, uint8 decimals, address oracle, uint8 oracleDecimals, bool paused, uint256 depositRatio, uint256 liquidationThreshold)
func (_Forge *ForgeCallerSession) Assets(arg0 *big.Int) (struct {
	Token                common.Address
	Decimals             uint8
	Oracle               common.Address
	OracleDecimals       uint8
	Paused               bool
	DepositRatio         *big.Int
	LiquidationThreshold *big.Int
}, error) {
	return _Forge.Contract.Assets(&_Forge.CallOpts, arg0)
}

// CRatioDecimals is a free data retrieval call binding the contract method 0x9e231c9e.
//
// Solidity: function cRatioDecimals() view returns(uint8)
func (_Forge *ForgeCaller) CRatioDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Forge.contract.Call(opts, &out, "cRatioDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CRatioDecimals is a free data retrieval call binding the contract method 0x9e231c9e.
//
// Solidity: function cRatioDecimals() view returns(uint8)
func (_Forge *ForgeSession) CRatioDecimals() (uint8, error) {
	return _Forge.Contract.CRatioDecimals(&_Forge.CallOpts)
}

// CRatioDecimals is a free data retrieval call binding the contract method 0x9e231c9e.
//
// Solidity: function cRatioDecimals() view returns(uint8)
func (_Forge *ForgeCallerSession) CRatioDecimals() (uint8, error) {
	return _Forge.Contract.CRatioDecimals(&_Forge.CallOpts)
}

// Collector is a free data retrieval call binding the contract method 0x913e77ad.
//
// Solidity: function collector() view returns(address)
func (_Forge *ForgeCaller) Collector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Forge.contract.Call(opts, &out, "collector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Collector is a free data retrieval call binding the contract method 0x913e77ad.
//
// Solidity: function collector() view returns(address)
func (_Forge *ForgeSession) Collector() (common.Address, error) {
	return _Forge.Contract.Collector(&_Forge.CallOpts)
}

// Collector is a free data retrieval call binding the contract method 0x913e77ad.
//
// Solidity: function collector() view returns(address)
func (_Forge *ForgeCallerSession) Collector() (common.Address, error) {
	return _Forge.Contract.Collector(&_Forge.CallOpts)
}

// Currency is a free data retrieval call binding the contract method 0xe5a6b10f.
//
// Solidity: function currency() view returns(address)
func (_Forge *ForgeCaller) Currency(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Forge.contract.Call(opts, &out, "currency")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Currency is a free data retrieval call binding the contract method 0xe5a6b10f.
//
// Solidity: function currency() view returns(address)
func (_Forge *ForgeSession) Currency() (common.Address, error) {
	return _Forge.Contract.Currency(&_Forge.CallOpts)
}

// Currency is a free data retrieval call binding the contract method 0xe5a6b10f.
//
// Solidity: function currency() view returns(address)
func (_Forge *ForgeCallerSession) Currency() (common.Address, error) {
	return _Forge.Contract.Currency(&_Forge.CallOpts)
}

// CurrencyDecimals is a free data retrieval call binding the contract method 0x7624388e.
//
// Solidity: function currencyDecimals() view returns(uint8)
func (_Forge *ForgeCaller) CurrencyDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Forge.contract.Call(opts, &out, "currencyDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CurrencyDecimals is a free data retrieval call binding the contract method 0x7624388e.
//
// Solidity: function currencyDecimals() view returns(uint8)
func (_Forge *ForgeSession) CurrencyDecimals() (uint8, error) {
	return _Forge.Contract.CurrencyDecimals(&_Forge.CallOpts)
}

// CurrencyDecimals is a free data retrieval call binding the contract method 0x7624388e.
//
// Solidity: function currencyDecimals() view returns(uint8)
func (_Forge *ForgeCallerSession) CurrencyDecimals() (uint8, error) {
	return _Forge.Contract.CurrencyDecimals(&_Forge.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Forge *ForgeCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Forge.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Forge *ForgeSession) Fee() (*big.Int, error) {
	return _Forge.Contract.Fee(&_Forge.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Forge *ForgeCallerSession) Fee() (*big.Int, error) {
	return _Forge.Contract.Fee(&_Forge.CallOpts)
}

// FeeBase is a free data retrieval call binding the contract method 0x95e911a8.
//
// Solidity: function feeBase() view returns(uint256)
func (_Forge *ForgeCaller) FeeBase(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Forge.contract.Call(opts, &out, "feeBase")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeBase is a free data retrieval call binding the contract method 0x95e911a8.
//
// Solidity: function feeBase() view returns(uint256)
func (_Forge *ForgeSession) FeeBase() (*big.Int, error) {
	return _Forge.Contract.FeeBase(&_Forge.CallOpts)
}

// FeeBase is a free data retrieval call binding the contract method 0x95e911a8.
//
// Solidity: function feeBase() view returns(uint256)
func (_Forge *ForgeCallerSession) FeeBase() (*big.Int, error) {
	return _Forge.Contract.FeeBase(&_Forge.CallOpts)
}

// FeeLiquidatorPercent is a free data retrieval call binding the contract method 0x29c929ff.
//
// Solidity: function feeLiquidatorPercent() view returns(uint256)
func (_Forge *ForgeCaller) FeeLiquidatorPercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Forge.contract.Call(opts, &out, "feeLiquidatorPercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeLiquidatorPercent is a free data retrieval call binding the contract method 0x29c929ff.
//
// Solidity: function feeLiquidatorPercent() view returns(uint256)
func (_Forge *ForgeSession) FeeLiquidatorPercent() (*big.Int, error) {
	return _Forge.Contract.FeeLiquidatorPercent(&_Forge.CallOpts)
}

// FeeLiquidatorPercent is a free data retrieval call binding the contract method 0x29c929ff.
//
// Solidity: function feeLiquidatorPercent() view returns(uint256)
func (_Forge *ForgeCallerSession) FeeLiquidatorPercent() (*big.Int, error) {
	return _Forge.Contract.FeeLiquidatorPercent(&_Forge.CallOpts)
}

// FeeLiquidatorPercentBase is a free data retrieval call binding the contract method 0x4a0d05f6.
//
// Solidity: function feeLiquidatorPercentBase() view returns(uint256)
func (_Forge *ForgeCaller) FeeLiquidatorPercentBase(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Forge.contract.Call(opts, &out, "feeLiquidatorPercentBase")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeLiquidatorPercentBase is a free data retrieval call binding the contract method 0x4a0d05f6.
//
// Solidity: function feeLiquidatorPercentBase() view returns(uint256)
func (_Forge *ForgeSession) FeeLiquidatorPercentBase() (*big.Int, error) {
	return _Forge.Contract.FeeLiquidatorPercentBase(&_Forge.CallOpts)
}

// FeeLiquidatorPercentBase is a free data retrieval call binding the contract method 0x4a0d05f6.
//
// Solidity: function feeLiquidatorPercentBase() view returns(uint256)
func (_Forge *ForgeCallerSession) FeeLiquidatorPercentBase() (*big.Int, error) {
	return _Forge.Contract.FeeLiquidatorPercentBase(&_Forge.CallOpts)
}

// GetOvenRatio is a free data retrieval call binding the contract method 0x909b7a77.
//
// Solidity: function getOvenRatio(address account, uint256 id) view returns(uint256 ratio)
func (_Forge *ForgeCaller) GetOvenRatio(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Forge.contract.Call(opts, &out, "getOvenRatio", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOvenRatio is a free data retrieval call binding the contract method 0x909b7a77.
//
// Solidity: function getOvenRatio(address account, uint256 id) view returns(uint256 ratio)
func (_Forge *ForgeSession) GetOvenRatio(account common.Address, id *big.Int) (*big.Int, error) {
	return _Forge.Contract.GetOvenRatio(&_Forge.CallOpts, account, id)
}

// GetOvenRatio is a free data retrieval call binding the contract method 0x909b7a77.
//
// Solidity: function getOvenRatio(address account, uint256 id) view returns(uint256 ratio)
func (_Forge *ForgeCallerSession) GetOvenRatio(account common.Address, id *big.Int) (*big.Int, error) {
	return _Forge.Contract.GetOvenRatio(&_Forge.CallOpts, account, id)
}

// GetOvensLength is a free data retrieval call binding the contract method 0x7fba91d1.
//
// Solidity: function getOvensLength(address account) view returns(uint256 numOvens)
func (_Forge *ForgeCaller) GetOvensLength(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Forge.contract.Call(opts, &out, "getOvensLength", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOvensLength is a free data retrieval call binding the contract method 0x7fba91d1.
//
// Solidity: function getOvensLength(address account) view returns(uint256 numOvens)
func (_Forge *ForgeSession) GetOvensLength(account common.Address) (*big.Int, error) {
	return _Forge.Contract.GetOvensLength(&_Forge.CallOpts, account)
}

// GetOvensLength is a free data retrieval call binding the contract method 0x7fba91d1.
//
// Solidity: function getOvensLength(address account) view returns(uint256 numOvens)
func (_Forge *ForgeCallerSession) GetOvensLength(account common.Address) (*big.Int, error) {
	return _Forge.Contract.GetOvensLength(&_Forge.CallOpts, account)
}

// GetUsersLength is a free data retrieval call binding the contract method 0x595b1a3e.
//
// Solidity: function getUsersLength() view returns(uint256 numUsers)
func (_Forge *ForgeCaller) GetUsersLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Forge.contract.Call(opts, &out, "getUsersLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUsersLength is a free data retrieval call binding the contract method 0x595b1a3e.
//
// Solidity: function getUsersLength() view returns(uint256 numUsers)
func (_Forge *ForgeSession) GetUsersLength() (*big.Int, error) {
	return _Forge.Contract.GetUsersLength(&_Forge.CallOpts)
}

// GetUsersLength is a free data retrieval call binding the contract method 0x595b1a3e.
//
// Solidity: function getUsersLength() view returns(uint256 numUsers)
func (_Forge *ForgeCallerSession) GetUsersLength() (*big.Int, error) {
	return _Forge.Contract.GetUsersLength(&_Forge.CallOpts)
}

// IsUser is a free data retrieval call binding the contract method 0x4209fff1.
//
// Solidity: function isUser(address ) view returns(bool)
func (_Forge *ForgeCaller) IsUser(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Forge.contract.Call(opts, &out, "isUser", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUser is a free data retrieval call binding the contract method 0x4209fff1.
//
// Solidity: function isUser(address ) view returns(bool)
func (_Forge *ForgeSession) IsUser(arg0 common.Address) (bool, error) {
	return _Forge.Contract.IsUser(&_Forge.CallOpts, arg0)
}

// IsUser is a free data retrieval call binding the contract method 0x4209fff1.
//
// Solidity: function isUser(address ) view returns(bool)
func (_Forge *ForgeCallerSession) IsUser(arg0 common.Address) (bool, error) {
	return _Forge.Contract.IsUser(&_Forge.CallOpts, arg0)
}

// MinCollateral is a free data retrieval call binding the contract method 0xba2de9bc.
//
// Solidity: function minCollateral() view returns(uint256)
func (_Forge *ForgeCaller) MinCollateral(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Forge.contract.Call(opts, &out, "minCollateral")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinCollateral is a free data retrieval call binding the contract method 0xba2de9bc.
//
// Solidity: function minCollateral() view returns(uint256)
func (_Forge *ForgeSession) MinCollateral() (*big.Int, error) {
	return _Forge.Contract.MinCollateral(&_Forge.CallOpts)
}

// MinCollateral is a free data retrieval call binding the contract method 0xba2de9bc.
//
// Solidity: function minCollateral() view returns(uint256)
func (_Forge *ForgeCallerSession) MinCollateral() (*big.Int, error) {
	return _Forge.Contract.MinCollateral(&_Forge.CallOpts)
}

// Ovens is a free data retrieval call binding the contract method 0x4f6ada8d.
//
// Solidity: function ovens(address , uint256 ) view returns(address account, uint256 collateral, uint256 amount, uint256 token)
func (_Forge *ForgeCaller) Ovens(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Account    common.Address
	Collateral *big.Int
	Amount     *big.Int
	Token      *big.Int
}, error) {
	var out []interface{}
	err := _Forge.contract.Call(opts, &out, "ovens", arg0, arg1)

	outstruct := new(struct {
		Account    common.Address
		Collateral *big.Int
		Amount     *big.Int
		Token      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Account = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Collateral = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Token = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Ovens is a free data retrieval call binding the contract method 0x4f6ada8d.
//
// Solidity: function ovens(address , uint256 ) view returns(address account, uint256 collateral, uint256 amount, uint256 token)
func (_Forge *ForgeSession) Ovens(arg0 common.Address, arg1 *big.Int) (struct {
	Account    common.Address
	Collateral *big.Int
	Amount     *big.Int
	Token      *big.Int
}, error) {
	return _Forge.Contract.Ovens(&_Forge.CallOpts, arg0, arg1)
}

// Ovens is a free data retrieval call binding the contract method 0x4f6ada8d.
//
// Solidity: function ovens(address , uint256 ) view returns(address account, uint256 collateral, uint256 amount, uint256 token)
func (_Forge *ForgeCallerSession) Ovens(arg0 common.Address, arg1 *big.Int) (struct {
	Account    common.Address
	Collateral *big.Int
	Amount     *big.Int
	Token      *big.Int
}, error) {
	return _Forge.Contract.Ovens(&_Forge.CallOpts, arg0, arg1)
}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_Forge *ForgeCaller) Pauser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Forge.contract.Call(opts, &out, "pauser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_Forge *ForgeSession) Pauser() (common.Address, error) {
	return _Forge.Contract.Pauser(&_Forge.CallOpts)
}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_Forge *ForgeCallerSession) Pauser() (common.Address, error) {
	return _Forge.Contract.Pauser(&_Forge.CallOpts)
}

// Users is a free data retrieval call binding the contract method 0x365b98b2.
//
// Solidity: function users(uint256 ) view returns(address)
func (_Forge *ForgeCaller) Users(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Forge.contract.Call(opts, &out, "users", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Users is a free data retrieval call binding the contract method 0x365b98b2.
//
// Solidity: function users(uint256 ) view returns(address)
func (_Forge *ForgeSession) Users(arg0 *big.Int) (common.Address, error) {
	return _Forge.Contract.Users(&_Forge.CallOpts, arg0)
}

// Users is a free data retrieval call binding the contract method 0x365b98b2.
//
// Solidity: function users(uint256 ) view returns(address)
func (_Forge *ForgeCallerSession) Users(arg0 *big.Int) (common.Address, error) {
	return _Forge.Contract.Users(&_Forge.CallOpts, arg0)
}

// AddAsset is a paid mutator transaction binding the contract method 0x206ec0ed.
//
// Solidity: function addAsset(address _token, address _oracle, bool _paused, uint256 _depositRatio, uint256 _liquidationThreshold) returns()
func (_Forge *ForgeTransactor) AddAsset(opts *bind.TransactOpts, _token common.Address, _oracle common.Address, _paused bool, _depositRatio *big.Int, _liquidationThreshold *big.Int) (*types.Transaction, error) {
	return _Forge.contract.Transact(opts, "addAsset", _token, _oracle, _paused, _depositRatio, _liquidationThreshold)
}

// AddAsset is a paid mutator transaction binding the contract method 0x206ec0ed.
//
// Solidity: function addAsset(address _token, address _oracle, bool _paused, uint256 _depositRatio, uint256 _liquidationThreshold) returns()
func (_Forge *ForgeSession) AddAsset(_token common.Address, _oracle common.Address, _paused bool, _depositRatio *big.Int, _liquidationThreshold *big.Int) (*types.Transaction, error) {
	return _Forge.Contract.AddAsset(&_Forge.TransactOpts, _token, _oracle, _paused, _depositRatio, _liquidationThreshold)
}

// AddAsset is a paid mutator transaction binding the contract method 0x206ec0ed.
//
// Solidity: function addAsset(address _token, address _oracle, bool _paused, uint256 _depositRatio, uint256 _liquidationThreshold) returns()
func (_Forge *ForgeTransactorSession) AddAsset(_token common.Address, _oracle common.Address, _paused bool, _depositRatio *big.Int, _liquidationThreshold *big.Int) (*types.Transaction, error) {
	return _Forge.Contract.AddAsset(&_Forge.TransactOpts, _token, _oracle, _paused, _depositRatio, _liquidationThreshold)
}

// Close is a paid mutator transaction binding the contract method 0x0aebeb4e.
//
// Solidity: function close(uint256 id) returns()
func (_Forge *ForgeTransactor) Close(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _Forge.contract.Transact(opts, "close", id)
}

// Close is a paid mutator transaction binding the contract method 0x0aebeb4e.
//
// Solidity: function close(uint256 id) returns()
func (_Forge *ForgeSession) Close(id *big.Int) (*types.Transaction, error) {
	return _Forge.Contract.Close(&_Forge.TransactOpts, id)
}

// Close is a paid mutator transaction binding the contract method 0x0aebeb4e.
//
// Solidity: function close(uint256 id) returns()
func (_Forge *ForgeTransactorSession) Close(id *big.Int) (*types.Transaction, error) {
	return _Forge.Contract.Close(&_Forge.TransactOpts, id)
}

// Edit is a paid mutator transaction binding the contract method 0x771612ab.
//
// Solidity: function edit(uint256 id, uint256 collateral, uint256 amount) returns()
func (_Forge *ForgeTransactor) Edit(opts *bind.TransactOpts, id *big.Int, collateral *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Forge.contract.Transact(opts, "edit", id, collateral, amount)
}

// Edit is a paid mutator transaction binding the contract method 0x771612ab.
//
// Solidity: function edit(uint256 id, uint256 collateral, uint256 amount) returns()
func (_Forge *ForgeSession) Edit(id *big.Int, collateral *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Forge.Contract.Edit(&_Forge.TransactOpts, id, collateral, amount)
}

// Edit is a paid mutator transaction binding the contract method 0x771612ab.
//
// Solidity: function edit(uint256 id, uint256 collateral, uint256 amount) returns()
func (_Forge *ForgeTransactorSession) Edit(id *big.Int, collateral *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Forge.Contract.Edit(&_Forge.TransactOpts, id, collateral, amount)
}

// Liquidate is a paid mutator transaction binding the contract method 0x0710285c.
//
// Solidity: function liquidate(address account, uint256 id, uint256 amount) returns()
func (_Forge *ForgeTransactor) Liquidate(opts *bind.TransactOpts, account common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Forge.contract.Transact(opts, "liquidate", account, id, amount)
}

// Liquidate is a paid mutator transaction binding the contract method 0x0710285c.
//
// Solidity: function liquidate(address account, uint256 id, uint256 amount) returns()
func (_Forge *ForgeSession) Liquidate(account common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Forge.Contract.Liquidate(&_Forge.TransactOpts, account, id, amount)
}

// Liquidate is a paid mutator transaction binding the contract method 0x0710285c.
//
// Solidity: function liquidate(address account, uint256 id, uint256 amount) returns()
func (_Forge *ForgeTransactorSession) Liquidate(account common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Forge.Contract.Liquidate(&_Forge.TransactOpts, account, id, amount)
}

// Open is a paid mutator transaction binding the contract method 0x3d57b2a9.
//
// Solidity: function open(uint256 collateral, uint256 amount, uint256 assetId) returns()
func (_Forge *ForgeTransactor) Open(opts *bind.TransactOpts, collateral *big.Int, amount *big.Int, assetId *big.Int) (*types.Transaction, error) {
	return _Forge.contract.Transact(opts, "open", collateral, amount, assetId)
}

// Open is a paid mutator transaction binding the contract method 0x3d57b2a9.
//
// Solidity: function open(uint256 collateral, uint256 amount, uint256 assetId) returns()
func (_Forge *ForgeSession) Open(collateral *big.Int, amount *big.Int, assetId *big.Int) (*types.Transaction, error) {
	return _Forge.Contract.Open(&_Forge.TransactOpts, collateral, amount, assetId)
}

// Open is a paid mutator transaction binding the contract method 0x3d57b2a9.
//
// Solidity: function open(uint256 collateral, uint256 amount, uint256 assetId) returns()
func (_Forge *ForgeTransactorSession) Open(collateral *big.Int, amount *big.Int, assetId *big.Int) (*types.Transaction, error) {
	return _Forge.Contract.Open(&_Forge.TransactOpts, collateral, amount, assetId)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Forge *ForgeTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Forge.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Forge *ForgeSession) Pause() (*types.Transaction, error) {
	return _Forge.Contract.Pause(&_Forge.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Forge *ForgeTransactorSession) Pause() (*types.Transaction, error) {
	return _Forge.Contract.Pause(&_Forge.TransactOpts)
}

// PauseAsset is a paid mutator transaction binding the contract method 0x6cdbf1f9.
//
// Solidity: function pauseAsset(uint256 id) returns()
func (_Forge *ForgeTransactor) PauseAsset(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _Forge.contract.Transact(opts, "pauseAsset", id)
}

// PauseAsset is a paid mutator transaction binding the contract method 0x6cdbf1f9.
//
// Solidity: function pauseAsset(uint256 id) returns()
func (_Forge *ForgeSession) PauseAsset(id *big.Int) (*types.Transaction, error) {
	return _Forge.Contract.PauseAsset(&_Forge.TransactOpts, id)
}

// PauseAsset is a paid mutator transaction binding the contract method 0x6cdbf1f9.
//
// Solidity: function pauseAsset(uint256 id) returns()
func (_Forge *ForgeTransactorSession) PauseAsset(id *big.Int) (*types.Transaction, error) {
	return _Forge.Contract.PauseAsset(&_Forge.TransactOpts, id)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Forge *ForgeTransactor) Resume(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Forge.contract.Transact(opts, "resume")
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Forge *ForgeSession) Resume() (*types.Transaction, error) {
	return _Forge.Contract.Resume(&_Forge.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Forge *ForgeTransactorSession) Resume() (*types.Transaction, error) {
	return _Forge.Contract.Resume(&_Forge.TransactOpts)
}

// ResumeAsset is a paid mutator transaction binding the contract method 0xfdea14e6.
//
// Solidity: function resumeAsset(uint256 id) returns()
func (_Forge *ForgeTransactor) ResumeAsset(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _Forge.contract.Transact(opts, "resumeAsset", id)
}

// ResumeAsset is a paid mutator transaction binding the contract method 0xfdea14e6.
//
// Solidity: function resumeAsset(uint256 id) returns()
func (_Forge *ForgeSession) ResumeAsset(id *big.Int) (*types.Transaction, error) {
	return _Forge.Contract.ResumeAsset(&_Forge.TransactOpts, id)
}

// ResumeAsset is a paid mutator transaction binding the contract method 0xfdea14e6.
//
// Solidity: function resumeAsset(uint256 id) returns()
func (_Forge *ForgeTransactorSession) ResumeAsset(id *big.Int) (*types.Transaction, error) {
	return _Forge.Contract.ResumeAsset(&_Forge.TransactOpts, id)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_Forge *ForgeTransactor) SetAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _Forge.contract.Transact(opts, "setAdmin", _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_Forge *ForgeSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _Forge.Contract.SetAdmin(&_Forge.TransactOpts, _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_Forge *ForgeTransactorSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _Forge.Contract.SetAdmin(&_Forge.TransactOpts, _admin)
}

// SetAssetOracle is a paid mutator transaction binding the contract method 0xae92e247.
//
// Solidity: function setAssetOracle(address _oracle, uint256 id) returns()
func (_Forge *ForgeTransactor) SetAssetOracle(opts *bind.TransactOpts, _oracle common.Address, id *big.Int) (*types.Transaction, error) {
	return _Forge.contract.Transact(opts, "setAssetOracle", _oracle, id)
}

// SetAssetOracle is a paid mutator transaction binding the contract method 0xae92e247.
//
// Solidity: function setAssetOracle(address _oracle, uint256 id) returns()
func (_Forge *ForgeSession) SetAssetOracle(_oracle common.Address, id *big.Int) (*types.Transaction, error) {
	return _Forge.Contract.SetAssetOracle(&_Forge.TransactOpts, _oracle, id)
}

// SetAssetOracle is a paid mutator transaction binding the contract method 0xae92e247.
//
// Solidity: function setAssetOracle(address _oracle, uint256 id) returns()
func (_Forge *ForgeTransactorSession) SetAssetOracle(_oracle common.Address, id *big.Int) (*types.Transaction, error) {
	return _Forge.Contract.SetAssetOracle(&_Forge.TransactOpts, _oracle, id)
}

// SetCollector is a paid mutator transaction binding the contract method 0xfb5b82d0.
//
// Solidity: function setCollector(address _collector) returns()
func (_Forge *ForgeTransactor) SetCollector(opts *bind.TransactOpts, _collector common.Address) (*types.Transaction, error) {
	return _Forge.contract.Transact(opts, "setCollector", _collector)
}

// SetCollector is a paid mutator transaction binding the contract method 0xfb5b82d0.
//
// Solidity: function setCollector(address _collector) returns()
func (_Forge *ForgeSession) SetCollector(_collector common.Address) (*types.Transaction, error) {
	return _Forge.Contract.SetCollector(&_Forge.TransactOpts, _collector)
}

// SetCollector is a paid mutator transaction binding the contract method 0xfb5b82d0.
//
// Solidity: function setCollector(address _collector) returns()
func (_Forge *ForgeTransactorSession) SetCollector(_collector common.Address) (*types.Transaction, error) {
	return _Forge.Contract.SetCollector(&_Forge.TransactOpts, _collector)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _fee) returns()
func (_Forge *ForgeTransactor) SetFee(opts *bind.TransactOpts, _fee *big.Int) (*types.Transaction, error) {
	return _Forge.contract.Transact(opts, "setFee", _fee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _fee) returns()
func (_Forge *ForgeSession) SetFee(_fee *big.Int) (*types.Transaction, error) {
	return _Forge.Contract.SetFee(&_Forge.TransactOpts, _fee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _fee) returns()
func (_Forge *ForgeTransactorSession) SetFee(_fee *big.Int) (*types.Transaction, error) {
	return _Forge.Contract.SetFee(&_Forge.TransactOpts, _fee)
}

// SetFeeLiquidation is a paid mutator transaction binding the contract method 0x115c2ec7.
//
// Solidity: function setFeeLiquidation(uint256 _feeLiquidatorPercent) returns()
func (_Forge *ForgeTransactor) SetFeeLiquidation(opts *bind.TransactOpts, _feeLiquidatorPercent *big.Int) (*types.Transaction, error) {
	return _Forge.contract.Transact(opts, "setFeeLiquidation", _feeLiquidatorPercent)
}

// SetFeeLiquidation is a paid mutator transaction binding the contract method 0x115c2ec7.
//
// Solidity: function setFeeLiquidation(uint256 _feeLiquidatorPercent) returns()
func (_Forge *ForgeSession) SetFeeLiquidation(_feeLiquidatorPercent *big.Int) (*types.Transaction, error) {
	return _Forge.Contract.SetFeeLiquidation(&_Forge.TransactOpts, _feeLiquidatorPercent)
}

// SetFeeLiquidation is a paid mutator transaction binding the contract method 0x115c2ec7.
//
// Solidity: function setFeeLiquidation(uint256 _feeLiquidatorPercent) returns()
func (_Forge *ForgeTransactorSession) SetFeeLiquidation(_feeLiquidatorPercent *big.Int) (*types.Transaction, error) {
	return _Forge.Contract.SetFeeLiquidation(&_Forge.TransactOpts, _feeLiquidatorPercent)
}

// SetMinCollateral is a paid mutator transaction binding the contract method 0x846321a4.
//
// Solidity: function setMinCollateral(uint256 _minCollateral) returns()
func (_Forge *ForgeTransactor) SetMinCollateral(opts *bind.TransactOpts, _minCollateral *big.Int) (*types.Transaction, error) {
	return _Forge.contract.Transact(opts, "setMinCollateral", _minCollateral)
}

// SetMinCollateral is a paid mutator transaction binding the contract method 0x846321a4.
//
// Solidity: function setMinCollateral(uint256 _minCollateral) returns()
func (_Forge *ForgeSession) SetMinCollateral(_minCollateral *big.Int) (*types.Transaction, error) {
	return _Forge.Contract.SetMinCollateral(&_Forge.TransactOpts, _minCollateral)
}

// SetMinCollateral is a paid mutator transaction binding the contract method 0x846321a4.
//
// Solidity: function setMinCollateral(uint256 _minCollateral) returns()
func (_Forge *ForgeTransactorSession) SetMinCollateral(_minCollateral *big.Int) (*types.Transaction, error) {
	return _Forge.Contract.SetMinCollateral(&_Forge.TransactOpts, _minCollateral)
}

// SetPauser is a paid mutator transaction binding the contract method 0x2d88af4a.
//
// Solidity: function setPauser(address _pauser) returns()
func (_Forge *ForgeTransactor) SetPauser(opts *bind.TransactOpts, _pauser common.Address) (*types.Transaction, error) {
	return _Forge.contract.Transact(opts, "setPauser", _pauser)
}

// SetPauser is a paid mutator transaction binding the contract method 0x2d88af4a.
//
// Solidity: function setPauser(address _pauser) returns()
func (_Forge *ForgeSession) SetPauser(_pauser common.Address) (*types.Transaction, error) {
	return _Forge.Contract.SetPauser(&_Forge.TransactOpts, _pauser)
}

// SetPauser is a paid mutator transaction binding the contract method 0x2d88af4a.
//
// Solidity: function setPauser(address _pauser) returns()
func (_Forge *ForgeTransactorSession) SetPauser(_pauser common.Address) (*types.Transaction, error) {
	return _Forge.Contract.SetPauser(&_Forge.TransactOpts, _pauser)
}

// ForgeCloseIterator is returned from FilterClose and is used to iterate over the raw logs and unpacked data for Close events raised by the Forge contract.
type ForgeCloseIterator struct {
	Event *ForgeClose // Event containing the contract specifics and raw log

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
func (it *ForgeCloseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ForgeClose)
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
		it.Event = new(ForgeClose)
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
func (it *ForgeCloseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ForgeCloseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ForgeClose represents a Close event raised by the Forge contract.
type ForgeClose struct {
	Account common.Address
	Id      *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClose is a free log retrieval operation binding the contract event 0x684222b0069d4a2e5e0d986611cc5182d543904c4e4264bf770d4e51faefc822.
//
// Solidity: event Close(address indexed account, uint256 indexed id)
func (_Forge *ForgeFilterer) FilterClose(opts *bind.FilterOpts, account []common.Address, id []*big.Int) (*ForgeCloseIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Forge.contract.FilterLogs(opts, "Close", accountRule, idRule)
	if err != nil {
		return nil, err
	}
	return &ForgeCloseIterator{contract: _Forge.contract, event: "Close", logs: logs, sub: sub}, nil
}

// WatchClose is a free log subscription operation binding the contract event 0x684222b0069d4a2e5e0d986611cc5182d543904c4e4264bf770d4e51faefc822.
//
// Solidity: event Close(address indexed account, uint256 indexed id)
func (_Forge *ForgeFilterer) WatchClose(opts *bind.WatchOpts, sink chan<- *ForgeClose, account []common.Address, id []*big.Int) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Forge.contract.WatchLogs(opts, "Close", accountRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ForgeClose)
				if err := _Forge.contract.UnpackLog(event, "Close", log); err != nil {
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

// ParseClose is a log parse operation binding the contract event 0x684222b0069d4a2e5e0d986611cc5182d543904c4e4264bf770d4e51faefc822.
//
// Solidity: event Close(address indexed account, uint256 indexed id)
func (_Forge *ForgeFilterer) ParseClose(log types.Log) (*ForgeClose, error) {
	event := new(ForgeClose)
	if err := _Forge.contract.UnpackLog(event, "Close", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ForgeEditIterator is returned from FilterEdit and is used to iterate over the raw logs and unpacked data for Edit events raised by the Forge contract.
type ForgeEditIterator struct {
	Event *ForgeEdit // Event containing the contract specifics and raw log

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
func (it *ForgeEditIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ForgeEdit)
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
		it.Event = new(ForgeEdit)
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
func (it *ForgeEditIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ForgeEditIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ForgeEdit represents a Edit event raised by the Forge contract.
type ForgeEdit struct {
	Account    common.Address
	Id         *big.Int
	Collateral *big.Int
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEdit is a free log retrieval operation binding the contract event 0xb126f3e2aa336dc5cbfec51e086edcde100a0d591c02286dcbc45b63cf941f0a.
//
// Solidity: event Edit(address indexed account, uint256 indexed id, uint256 collateral, uint256 amount)
func (_Forge *ForgeFilterer) FilterEdit(opts *bind.FilterOpts, account []common.Address, id []*big.Int) (*ForgeEditIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Forge.contract.FilterLogs(opts, "Edit", accountRule, idRule)
	if err != nil {
		return nil, err
	}
	return &ForgeEditIterator{contract: _Forge.contract, event: "Edit", logs: logs, sub: sub}, nil
}

// WatchEdit is a free log subscription operation binding the contract event 0xb126f3e2aa336dc5cbfec51e086edcde100a0d591c02286dcbc45b63cf941f0a.
//
// Solidity: event Edit(address indexed account, uint256 indexed id, uint256 collateral, uint256 amount)
func (_Forge *ForgeFilterer) WatchEdit(opts *bind.WatchOpts, sink chan<- *ForgeEdit, account []common.Address, id []*big.Int) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Forge.contract.WatchLogs(opts, "Edit", accountRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ForgeEdit)
				if err := _Forge.contract.UnpackLog(event, "Edit", log); err != nil {
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

// ParseEdit is a log parse operation binding the contract event 0xb126f3e2aa336dc5cbfec51e086edcde100a0d591c02286dcbc45b63cf941f0a.
//
// Solidity: event Edit(address indexed account, uint256 indexed id, uint256 collateral, uint256 amount)
func (_Forge *ForgeFilterer) ParseEdit(log types.Log) (*ForgeEdit, error) {
	event := new(ForgeEdit)
	if err := _Forge.contract.UnpackLog(event, "Edit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ForgeLiquidateIterator is returned from FilterLiquidate and is used to iterate over the raw logs and unpacked data for Liquidate events raised by the Forge contract.
type ForgeLiquidateIterator struct {
	Event *ForgeLiquidate // Event containing the contract specifics and raw log

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
func (it *ForgeLiquidateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ForgeLiquidate)
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
		it.Event = new(ForgeLiquidate)
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
func (it *ForgeLiquidateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ForgeLiquidateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ForgeLiquidate represents a Liquidate event raised by the Forge contract.
type ForgeLiquidate struct {
	Account    common.Address
	Liquidator common.Address
	Id         *big.Int
	Collateral *big.Int
	Amount     *big.Int
	Fee        *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLiquidate is a free log retrieval operation binding the contract event 0xdcfb2ea46125b5c04f5311ff0585da9998aa5151b11409482cd84c8ec08e8215.
//
// Solidity: event Liquidate(address indexed account, address indexed liquidator, uint256 indexed id, uint256 collateral, uint256 amount, uint256 fee)
func (_Forge *ForgeFilterer) FilterLiquidate(opts *bind.FilterOpts, account []common.Address, liquidator []common.Address, id []*big.Int) (*ForgeLiquidateIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Forge.contract.FilterLogs(opts, "Liquidate", accountRule, liquidatorRule, idRule)
	if err != nil {
		return nil, err
	}
	return &ForgeLiquidateIterator{contract: _Forge.contract, event: "Liquidate", logs: logs, sub: sub}, nil
}

// WatchLiquidate is a free log subscription operation binding the contract event 0xdcfb2ea46125b5c04f5311ff0585da9998aa5151b11409482cd84c8ec08e8215.
//
// Solidity: event Liquidate(address indexed account, address indexed liquidator, uint256 indexed id, uint256 collateral, uint256 amount, uint256 fee)
func (_Forge *ForgeFilterer) WatchLiquidate(opts *bind.WatchOpts, sink chan<- *ForgeLiquidate, account []common.Address, liquidator []common.Address, id []*big.Int) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Forge.contract.WatchLogs(opts, "Liquidate", accountRule, liquidatorRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ForgeLiquidate)
				if err := _Forge.contract.UnpackLog(event, "Liquidate", log); err != nil {
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

// ParseLiquidate is a log parse operation binding the contract event 0xdcfb2ea46125b5c04f5311ff0585da9998aa5151b11409482cd84c8ec08e8215.
//
// Solidity: event Liquidate(address indexed account, address indexed liquidator, uint256 indexed id, uint256 collateral, uint256 amount, uint256 fee)
func (_Forge *ForgeFilterer) ParseLiquidate(log types.Log) (*ForgeLiquidate, error) {
	event := new(ForgeLiquidate)
	if err := _Forge.contract.UnpackLog(event, "Liquidate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ForgeOpenIterator is returned from FilterOpen and is used to iterate over the raw logs and unpacked data for Open events raised by the Forge contract.
type ForgeOpenIterator struct {
	Event *ForgeOpen // Event containing the contract specifics and raw log

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
func (it *ForgeOpenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ForgeOpen)
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
		it.Event = new(ForgeOpen)
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
func (it *ForgeOpenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ForgeOpenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ForgeOpen represents a Open event raised by the Forge contract.
type ForgeOpen struct {
	Account    common.Address
	Id         *big.Int
	Amount     *big.Int
	Collateral *big.Int
	Token      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOpen is a free log retrieval operation binding the contract event 0x86b71d373e8d15f7c924340f2b1c55729b133d62d8fd4a7bbefd7fe0fbe38db8.
//
// Solidity: event Open(address indexed account, uint256 indexed id, uint256 amount, uint256 collateral, uint256 token)
func (_Forge *ForgeFilterer) FilterOpen(opts *bind.FilterOpts, account []common.Address, id []*big.Int) (*ForgeOpenIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Forge.contract.FilterLogs(opts, "Open", accountRule, idRule)
	if err != nil {
		return nil, err
	}
	return &ForgeOpenIterator{contract: _Forge.contract, event: "Open", logs: logs, sub: sub}, nil
}

// WatchOpen is a free log subscription operation binding the contract event 0x86b71d373e8d15f7c924340f2b1c55729b133d62d8fd4a7bbefd7fe0fbe38db8.
//
// Solidity: event Open(address indexed account, uint256 indexed id, uint256 amount, uint256 collateral, uint256 token)
func (_Forge *ForgeFilterer) WatchOpen(opts *bind.WatchOpts, sink chan<- *ForgeOpen, account []common.Address, id []*big.Int) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Forge.contract.WatchLogs(opts, "Open", accountRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ForgeOpen)
				if err := _Forge.contract.UnpackLog(event, "Open", log); err != nil {
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

// ParseOpen is a log parse operation binding the contract event 0x86b71d373e8d15f7c924340f2b1c55729b133d62d8fd4a7bbefd7fe0fbe38db8.
//
// Solidity: event Open(address indexed account, uint256 indexed id, uint256 amount, uint256 collateral, uint256 token)
func (_Forge *ForgeFilterer) ParseOpen(log types.Log) (*ForgeOpen, error) {
	event := new(ForgeOpen)
	if err := _Forge.contract.UnpackLog(event, "Open", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
