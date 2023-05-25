// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package flashloan

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

// FlashloanMetaData contains all meta data concerning the Flashloan contract.
var FlashloanMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPoolAddressesProvider\",\"name\":\"provider\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ADDRESSES_PROVIDER\",\"outputs\":[{\"internalType\":\"contractIPoolAddressesProvider\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POOL\",\"outputs\":[{\"internalType\":\"contractIPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"loanToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"loanAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"aaveFlashloan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_params\",\"type\":\"bytes\"}],\"name\":\"executeOperation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// FlashloanABI is the input ABI used to generate the binding from.
// Deprecated: Use FlashloanMetaData.ABI instead.
var FlashloanABI = FlashloanMetaData.ABI

// Flashloan is an auto generated Go binding around an Ethereum contract.
type Flashloan struct {
	FlashloanCaller     // Read-only binding to the contract
	FlashloanTransactor // Write-only binding to the contract
	FlashloanFilterer   // Log filterer for contract events
}

// FlashloanCaller is an auto generated read-only Go binding around an Ethereum contract.
type FlashloanCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlashloanTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FlashloanTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlashloanFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FlashloanFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlashloanSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FlashloanSession struct {
	Contract     *Flashloan        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FlashloanCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FlashloanCallerSession struct {
	Contract *FlashloanCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// FlashloanTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FlashloanTransactorSession struct {
	Contract     *FlashloanTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// FlashloanRaw is an auto generated low-level Go binding around an Ethereum contract.
type FlashloanRaw struct {
	Contract *Flashloan // Generic contract binding to access the raw methods on
}

// FlashloanCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FlashloanCallerRaw struct {
	Contract *FlashloanCaller // Generic read-only contract binding to access the raw methods on
}

// FlashloanTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FlashloanTransactorRaw struct {
	Contract *FlashloanTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFlashloan creates a new instance of Flashloan, bound to a specific deployed contract.
func NewFlashloan(address common.Address, backend bind.ContractBackend) (*Flashloan, error) {
	contract, err := bindFlashloan(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Flashloan{FlashloanCaller: FlashloanCaller{contract: contract}, FlashloanTransactor: FlashloanTransactor{contract: contract}, FlashloanFilterer: FlashloanFilterer{contract: contract}}, nil
}

// NewFlashloanCaller creates a new read-only instance of Flashloan, bound to a specific deployed contract.
func NewFlashloanCaller(address common.Address, caller bind.ContractCaller) (*FlashloanCaller, error) {
	contract, err := bindFlashloan(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FlashloanCaller{contract: contract}, nil
}

// NewFlashloanTransactor creates a new write-only instance of Flashloan, bound to a specific deployed contract.
func NewFlashloanTransactor(address common.Address, transactor bind.ContractTransactor) (*FlashloanTransactor, error) {
	contract, err := bindFlashloan(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FlashloanTransactor{contract: contract}, nil
}

// NewFlashloanFilterer creates a new log filterer instance of Flashloan, bound to a specific deployed contract.
func NewFlashloanFilterer(address common.Address, filterer bind.ContractFilterer) (*FlashloanFilterer, error) {
	contract, err := bindFlashloan(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FlashloanFilterer{contract: contract}, nil
}

// bindFlashloan binds a generic wrapper to an already deployed contract.
func bindFlashloan(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FlashloanMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Flashloan *FlashloanRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Flashloan.Contract.FlashloanCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Flashloan *FlashloanRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Flashloan.Contract.FlashloanTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Flashloan *FlashloanRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Flashloan.Contract.FlashloanTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Flashloan *FlashloanCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Flashloan.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Flashloan *FlashloanTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Flashloan.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Flashloan *FlashloanTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Flashloan.Contract.contract.Transact(opts, method, params...)
}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_Flashloan *FlashloanCaller) ADDRESSESPROVIDER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Flashloan.contract.Call(opts, &out, "ADDRESSES_PROVIDER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_Flashloan *FlashloanSession) ADDRESSESPROVIDER() (common.Address, error) {
	return _Flashloan.Contract.ADDRESSESPROVIDER(&_Flashloan.CallOpts)
}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_Flashloan *FlashloanCallerSession) ADDRESSESPROVIDER() (common.Address, error) {
	return _Flashloan.Contract.ADDRESSESPROVIDER(&_Flashloan.CallOpts)
}

// POOL is a free data retrieval call binding the contract method 0x7535d246.
//
// Solidity: function POOL() view returns(address)
func (_Flashloan *FlashloanCaller) POOL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Flashloan.contract.Call(opts, &out, "POOL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// POOL is a free data retrieval call binding the contract method 0x7535d246.
//
// Solidity: function POOL() view returns(address)
func (_Flashloan *FlashloanSession) POOL() (common.Address, error) {
	return _Flashloan.Contract.POOL(&_Flashloan.CallOpts)
}

// POOL is a free data retrieval call binding the contract method 0x7535d246.
//
// Solidity: function POOL() view returns(address)
func (_Flashloan *FlashloanCallerSession) POOL() (common.Address, error) {
	return _Flashloan.Contract.POOL(&_Flashloan.CallOpts)
}

// GetTime is a free data retrieval call binding the contract method 0x557ed1ba.
//
// Solidity: function getTime() view returns(uint256)
func (_Flashloan *FlashloanCaller) GetTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Flashloan.contract.Call(opts, &out, "getTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTime is a free data retrieval call binding the contract method 0x557ed1ba.
//
// Solidity: function getTime() view returns(uint256)
func (_Flashloan *FlashloanSession) GetTime() (*big.Int, error) {
	return _Flashloan.Contract.GetTime(&_Flashloan.CallOpts)
}

// GetTime is a free data retrieval call binding the contract method 0x557ed1ba.
//
// Solidity: function getTime() view returns(uint256)
func (_Flashloan *FlashloanCallerSession) GetTime() (*big.Int, error) {
	return _Flashloan.Contract.GetTime(&_Flashloan.CallOpts)
}

// AaveFlashloan is a paid mutator transaction binding the contract method 0xe45b9aa2.
//
// Solidity: function aaveFlashloan(address loanToken, uint256 loanAmount, bytes params) returns()
func (_Flashloan *FlashloanTransactor) AaveFlashloan(opts *bind.TransactOpts, loanToken common.Address, loanAmount *big.Int, params []byte) (*types.Transaction, error) {
	return _Flashloan.contract.Transact(opts, "aaveFlashloan", loanToken, loanAmount, params)
}

// AaveFlashloan is a paid mutator transaction binding the contract method 0xe45b9aa2.
//
// Solidity: function aaveFlashloan(address loanToken, uint256 loanAmount, bytes params) returns()
func (_Flashloan *FlashloanSession) AaveFlashloan(loanToken common.Address, loanAmount *big.Int, params []byte) (*types.Transaction, error) {
	return _Flashloan.Contract.AaveFlashloan(&_Flashloan.TransactOpts, loanToken, loanAmount, params)
}

// AaveFlashloan is a paid mutator transaction binding the contract method 0xe45b9aa2.
//
// Solidity: function aaveFlashloan(address loanToken, uint256 loanAmount, bytes params) returns()
func (_Flashloan *FlashloanTransactorSession) AaveFlashloan(loanToken common.Address, loanAmount *big.Int, params []byte) (*types.Transaction, error) {
	return _Flashloan.Contract.AaveFlashloan(&_Flashloan.TransactOpts, loanToken, loanAmount, params)
}

// ExecuteOperation is a paid mutator transaction binding the contract method 0x1b11d0ff.
//
// Solidity: function executeOperation(address asset, uint256 amount, uint256 premium, address , bytes _params) returns(bool)
func (_Flashloan *FlashloanTransactor) ExecuteOperation(opts *bind.TransactOpts, asset common.Address, amount *big.Int, premium *big.Int, arg3 common.Address, _params []byte) (*types.Transaction, error) {
	return _Flashloan.contract.Transact(opts, "executeOperation", asset, amount, premium, arg3, _params)
}

// ExecuteOperation is a paid mutator transaction binding the contract method 0x1b11d0ff.
//
// Solidity: function executeOperation(address asset, uint256 amount, uint256 premium, address , bytes _params) returns(bool)
func (_Flashloan *FlashloanSession) ExecuteOperation(asset common.Address, amount *big.Int, premium *big.Int, arg3 common.Address, _params []byte) (*types.Transaction, error) {
	return _Flashloan.Contract.ExecuteOperation(&_Flashloan.TransactOpts, asset, amount, premium, arg3, _params)
}

// ExecuteOperation is a paid mutator transaction binding the contract method 0x1b11d0ff.
//
// Solidity: function executeOperation(address asset, uint256 amount, uint256 premium, address , bytes _params) returns(bool)
func (_Flashloan *FlashloanTransactorSession) ExecuteOperation(asset common.Address, amount *big.Int, premium *big.Int, arg3 common.Address, _params []byte) (*types.Transaction, error) {
	return _Flashloan.Contract.ExecuteOperation(&_Flashloan.TransactOpts, asset, amount, premium, arg3, _params)
}
