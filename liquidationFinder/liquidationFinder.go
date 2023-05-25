// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package liquidationFinder

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

// LiquidationFinderLiquidation is an auto generated low-level Go binding around an user-defined struct.
type LiquidationFinderLiquidation struct {
	Account    common.Address
	OvenId     *big.Int
	Collateral *big.Int
	Amount     *big.Int
	Ratio      *big.Int
}

// LiquidationFinderMetaData contains all meta data concerning the LiquidationFinder contract.
var LiquidationFinderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fassetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRatio\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"skip\",\"type\":\"int256\"}],\"name\":\"getLiquidation\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ovenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ratio\",\"type\":\"uint256\"}],\"internalType\":\"structLiquidationFinder.Liquidation\",\"name\":\"liq\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// LiquidationFinderABI is the input ABI used to generate the binding from.
// Deprecated: Use LiquidationFinderMetaData.ABI instead.
var LiquidationFinderABI = LiquidationFinderMetaData.ABI

// LiquidationFinder is an auto generated Go binding around an Ethereum contract.
type LiquidationFinder struct {
	LiquidationFinderCaller     // Read-only binding to the contract
	LiquidationFinderTransactor // Write-only binding to the contract
	LiquidationFinderFilterer   // Log filterer for contract events
}

// LiquidationFinderCaller is an auto generated read-only Go binding around an Ethereum contract.
type LiquidationFinderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidationFinderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LiquidationFinderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidationFinderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LiquidationFinderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidationFinderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LiquidationFinderSession struct {
	Contract     *LiquidationFinder // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LiquidationFinderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LiquidationFinderCallerSession struct {
	Contract *LiquidationFinderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// LiquidationFinderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LiquidationFinderTransactorSession struct {
	Contract     *LiquidationFinderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// LiquidationFinderRaw is an auto generated low-level Go binding around an Ethereum contract.
type LiquidationFinderRaw struct {
	Contract *LiquidationFinder // Generic contract binding to access the raw methods on
}

// LiquidationFinderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LiquidationFinderCallerRaw struct {
	Contract *LiquidationFinderCaller // Generic read-only contract binding to access the raw methods on
}

// LiquidationFinderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LiquidationFinderTransactorRaw struct {
	Contract *LiquidationFinderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLiquidationFinder creates a new instance of LiquidationFinder, bound to a specific deployed contract.
func NewLiquidationFinder(address common.Address, backend bind.ContractBackend) (*LiquidationFinder, error) {
	contract, err := bindLiquidationFinder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LiquidationFinder{LiquidationFinderCaller: LiquidationFinderCaller{contract: contract}, LiquidationFinderTransactor: LiquidationFinderTransactor{contract: contract}, LiquidationFinderFilterer: LiquidationFinderFilterer{contract: contract}}, nil
}

// NewLiquidationFinderCaller creates a new read-only instance of LiquidationFinder, bound to a specific deployed contract.
func NewLiquidationFinderCaller(address common.Address, caller bind.ContractCaller) (*LiquidationFinderCaller, error) {
	contract, err := bindLiquidationFinder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LiquidationFinderCaller{contract: contract}, nil
}

// NewLiquidationFinderTransactor creates a new write-only instance of LiquidationFinder, bound to a specific deployed contract.
func NewLiquidationFinderTransactor(address common.Address, transactor bind.ContractTransactor) (*LiquidationFinderTransactor, error) {
	contract, err := bindLiquidationFinder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LiquidationFinderTransactor{contract: contract}, nil
}

// NewLiquidationFinderFilterer creates a new log filterer instance of LiquidationFinder, bound to a specific deployed contract.
func NewLiquidationFinderFilterer(address common.Address, filterer bind.ContractFilterer) (*LiquidationFinderFilterer, error) {
	contract, err := bindLiquidationFinder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LiquidationFinderFilterer{contract: contract}, nil
}

// bindLiquidationFinder binds a generic wrapper to an already deployed contract.
func bindLiquidationFinder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LiquidationFinderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LiquidationFinder *LiquidationFinderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LiquidationFinder.Contract.LiquidationFinderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LiquidationFinder *LiquidationFinderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidationFinder.Contract.LiquidationFinderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LiquidationFinder *LiquidationFinderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LiquidationFinder.Contract.LiquidationFinderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LiquidationFinder *LiquidationFinderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LiquidationFinder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LiquidationFinder *LiquidationFinderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidationFinder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LiquidationFinder *LiquidationFinderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LiquidationFinder.Contract.contract.Transact(opts, method, params...)
}

// GetLiquidation is a free data retrieval call binding the contract method 0xcaf0f0f9.
//
// Solidity: function getLiquidation(uint256 fassetId, uint256 maxRatio, int256 skip) view returns((address,uint256,uint256,uint256,uint256) liq)
func (_LiquidationFinder *LiquidationFinderCaller) GetLiquidation(opts *bind.CallOpts, fassetId *big.Int, maxRatio *big.Int, skip *big.Int) (LiquidationFinderLiquidation, error) {
	var out []interface{}
	err := _LiquidationFinder.contract.Call(opts, &out, "getLiquidation", fassetId, maxRatio, skip)

	if err != nil {
		return *new(LiquidationFinderLiquidation), err
	}

	out0 := *abi.ConvertType(out[0], new(LiquidationFinderLiquidation)).(*LiquidationFinderLiquidation)

	return out0, err

}

// GetLiquidation is a free data retrieval call binding the contract method 0xcaf0f0f9.
//
// Solidity: function getLiquidation(uint256 fassetId, uint256 maxRatio, int256 skip) view returns((address,uint256,uint256,uint256,uint256) liq)
func (_LiquidationFinder *LiquidationFinderSession) GetLiquidation(fassetId *big.Int, maxRatio *big.Int, skip *big.Int) (LiquidationFinderLiquidation, error) {
	return _LiquidationFinder.Contract.GetLiquidation(&_LiquidationFinder.CallOpts, fassetId, maxRatio, skip)
}

// GetLiquidation is a free data retrieval call binding the contract method 0xcaf0f0f9.
//
// Solidity: function getLiquidation(uint256 fassetId, uint256 maxRatio, int256 skip) view returns((address,uint256,uint256,uint256,uint256) liq)
func (_LiquidationFinder *LiquidationFinderCallerSession) GetLiquidation(fassetId *big.Int, maxRatio *big.Int, skip *big.Int) (LiquidationFinderLiquidation, error) {
	return _LiquidationFinder.Contract.GetLiquidation(&_LiquidationFinder.CallOpts, fassetId, maxRatio, skip)
}
