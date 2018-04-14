// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethereum

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ConsumeContentABI is the input ABI used to generate the binding from.
const ConsumeContentABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withDraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"consume\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"platform\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ratio\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"publisher\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"pPublisher\",\"type\":\"address\"},{\"name\":\"pPlatform\",\"type\":\"address\"},{\"name\":\"pPrice\",\"type\":\"uint32\"},{\"name\":\"pRatio\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// ConsumeContentBin is the compiled bytecode used for deploying new contracts.
const ConsumeContentBin = `0x6060604052341561000f57600080fd5b6040516080806104338339810160405280805191906020018051919060200180519190602001805160008054600160a060020a03978816600160a060020a0319918216179091556001805460ff90931678010000000000000000000000000000000000000000000000000260c060020a60ff021963ffffffff909716740100000000000000000000000000000000000000000260a060020a63ffffffff0219989099169390921692909217959095169590951792909216929092179092555050610355806100de6000396000f3006060604052600436106100825763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306661abd81146100875780630fdb1c10146100ac5780631dedc6f7146100c15780634bde38c8146100c957806371ca337d146101055780638c72c54e1461012e578063a035b1fe14610141575b600080fd5b341561009257600080fd5b61009a61016d565b60405190815260200160405180910390f35b34156100b757600080fd5b6100bf610173565b005b6100bf6101fe565b34156100d457600080fd5b6100dc6102a8565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b341561011057600080fd5b6101186102c4565b60405160ff909116815260200160405180910390f35b341561013957600080fd5b6100dc6102e9565b341561014c57600080fd5b610154610305565b60405163ffffffff909116815260200160405180910390f35b60025481565b73ffffffffffffffffffffffffffffffffffffffff33166000908152600360205260408120548190116101a557600080fd5b5073ffffffffffffffffffffffffffffffffffffffff3316600081815260036020526040808220805492905590919082156108fc0290839051600060405180830381858888f1935050505015156101fb57600080fd5b50565b60015460009074010000000000000000000000000000000000000000900463ffffffff16341461022d57600080fd5b600280546001908101909155546064907801000000000000000000000000000000000000000000000000900460ff1634026000805473ffffffffffffffffffffffffffffffffffffffff9081168252600360205260408083208054959094049485019093556001541681522080543492909203909101905550565b60015473ffffffffffffffffffffffffffffffffffffffff1681565b6001547801000000000000000000000000000000000000000000000000900460ff1681565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b60015474010000000000000000000000000000000000000000900463ffffffff16815600a165627a7a72305820a7dc09e862d3c4dcca2b37099b128529eb6eb6e498860210bc444a1505ca46e00029`

// DeployConsumeContent deploys a new Ethereum contract, binding an instance of ConsumeContent to it.
func DeployConsumeContent(auth *bind.TransactOpts, backend bind.ContractBackend, pPublisher common.Address, pPlatform common.Address, pPrice uint32, pRatio uint8) (common.Address, *types.Transaction, *ConsumeContent, error) {
	parsed, err := abi.JSON(strings.NewReader(ConsumeContentABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ConsumeContentBin), backend, pPublisher, pPlatform, pPrice, pRatio)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ConsumeContent{ConsumeContentCaller: ConsumeContentCaller{contract: contract}, ConsumeContentTransactor: ConsumeContentTransactor{contract: contract}, ConsumeContentFilterer: ConsumeContentFilterer{contract: contract}}, nil
}

// ConsumeContent is an auto generated Go binding around an Ethereum contract.
type ConsumeContent struct {
	ConsumeContentCaller     // Read-only binding to the contract
	ConsumeContentTransactor // Write-only binding to the contract
	ConsumeContentFilterer   // Log filterer for contract events
}

// ConsumeContentCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConsumeContentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConsumeContentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConsumeContentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConsumeContentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConsumeContentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConsumeContentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConsumeContentSession struct {
	Contract     *ConsumeContent   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConsumeContentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConsumeContentCallerSession struct {
	Contract *ConsumeContentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ConsumeContentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConsumeContentTransactorSession struct {
	Contract     *ConsumeContentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ConsumeContentRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConsumeContentRaw struct {
	Contract *ConsumeContent // Generic contract binding to access the raw methods on
}

// ConsumeContentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConsumeContentCallerRaw struct {
	Contract *ConsumeContentCaller // Generic read-only contract binding to access the raw methods on
}

// ConsumeContentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConsumeContentTransactorRaw struct {
	Contract *ConsumeContentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConsumeContent creates a new instance of ConsumeContent, bound to a specific deployed contract.
func NewConsumeContent(address common.Address, backend bind.ContractBackend) (*ConsumeContent, error) {
	contract, err := bindConsumeContent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ConsumeContent{ConsumeContentCaller: ConsumeContentCaller{contract: contract}, ConsumeContentTransactor: ConsumeContentTransactor{contract: contract}, ConsumeContentFilterer: ConsumeContentFilterer{contract: contract}}, nil
}

// NewConsumeContentCaller creates a new read-only instance of ConsumeContent, bound to a specific deployed contract.
func NewConsumeContentCaller(address common.Address, caller bind.ContractCaller) (*ConsumeContentCaller, error) {
	contract, err := bindConsumeContent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConsumeContentCaller{contract: contract}, nil
}

// NewConsumeContentTransactor creates a new write-only instance of ConsumeContent, bound to a specific deployed contract.
func NewConsumeContentTransactor(address common.Address, transactor bind.ContractTransactor) (*ConsumeContentTransactor, error) {
	contract, err := bindConsumeContent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConsumeContentTransactor{contract: contract}, nil
}

// NewConsumeContentFilterer creates a new log filterer instance of ConsumeContent, bound to a specific deployed contract.
func NewConsumeContentFilterer(address common.Address, filterer bind.ContractFilterer) (*ConsumeContentFilterer, error) {
	contract, err := bindConsumeContent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConsumeContentFilterer{contract: contract}, nil
}

// bindConsumeContent binds a generic wrapper to an already deployed contract.
func bindConsumeContent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConsumeContentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConsumeContent *ConsumeContentRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ConsumeContent.Contract.ConsumeContentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConsumeContent *ConsumeContentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConsumeContent.Contract.ConsumeContentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConsumeContent *ConsumeContentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConsumeContent.Contract.ConsumeContentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConsumeContent *ConsumeContentCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ConsumeContent.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConsumeContent *ConsumeContentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConsumeContent.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConsumeContent *ConsumeContentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConsumeContent.Contract.contract.Transact(opts, method, params...)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() constant returns(uint256)
func (_ConsumeContent *ConsumeContentCaller) Count(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ConsumeContent.contract.Call(opts, out, "count")
	return *ret0, err
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() constant returns(uint256)
func (_ConsumeContent *ConsumeContentSession) Count() (*big.Int, error) {
	return _ConsumeContent.Contract.Count(&_ConsumeContent.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() constant returns(uint256)
func (_ConsumeContent *ConsumeContentCallerSession) Count() (*big.Int, error) {
	return _ConsumeContent.Contract.Count(&_ConsumeContent.CallOpts)
}

// Platform is a free data retrieval call binding the contract method 0x4bde38c8.
//
// Solidity: function platform() constant returns(address)
func (_ConsumeContent *ConsumeContentCaller) Platform(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ConsumeContent.contract.Call(opts, out, "platform")
	return *ret0, err
}

// Platform is a free data retrieval call binding the contract method 0x4bde38c8.
//
// Solidity: function platform() constant returns(address)
func (_ConsumeContent *ConsumeContentSession) Platform() (common.Address, error) {
	return _ConsumeContent.Contract.Platform(&_ConsumeContent.CallOpts)
}

// Platform is a free data retrieval call binding the contract method 0x4bde38c8.
//
// Solidity: function platform() constant returns(address)
func (_ConsumeContent *ConsumeContentCallerSession) Platform() (common.Address, error) {
	return _ConsumeContent.Contract.Platform(&_ConsumeContent.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint32)
func (_ConsumeContent *ConsumeContentCaller) Price(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _ConsumeContent.contract.Call(opts, out, "price")
	return *ret0, err
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint32)
func (_ConsumeContent *ConsumeContentSession) Price() (uint32, error) {
	return _ConsumeContent.Contract.Price(&_ConsumeContent.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint32)
func (_ConsumeContent *ConsumeContentCallerSession) Price() (uint32, error) {
	return _ConsumeContent.Contract.Price(&_ConsumeContent.CallOpts)
}

// Publisher is a free data retrieval call binding the contract method 0x8c72c54e.
//
// Solidity: function publisher() constant returns(address)
func (_ConsumeContent *ConsumeContentCaller) Publisher(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ConsumeContent.contract.Call(opts, out, "publisher")
	return *ret0, err
}

// Publisher is a free data retrieval call binding the contract method 0x8c72c54e.
//
// Solidity: function publisher() constant returns(address)
func (_ConsumeContent *ConsumeContentSession) Publisher() (common.Address, error) {
	return _ConsumeContent.Contract.Publisher(&_ConsumeContent.CallOpts)
}

// Publisher is a free data retrieval call binding the contract method 0x8c72c54e.
//
// Solidity: function publisher() constant returns(address)
func (_ConsumeContent *ConsumeContentCallerSession) Publisher() (common.Address, error) {
	return _ConsumeContent.Contract.Publisher(&_ConsumeContent.CallOpts)
}

// Ratio is a free data retrieval call binding the contract method 0x71ca337d.
//
// Solidity: function ratio() constant returns(uint8)
func (_ConsumeContent *ConsumeContentCaller) Ratio(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ConsumeContent.contract.Call(opts, out, "ratio")
	return *ret0, err
}

// Ratio is a free data retrieval call binding the contract method 0x71ca337d.
//
// Solidity: function ratio() constant returns(uint8)
func (_ConsumeContent *ConsumeContentSession) Ratio() (uint8, error) {
	return _ConsumeContent.Contract.Ratio(&_ConsumeContent.CallOpts)
}

// Ratio is a free data retrieval call binding the contract method 0x71ca337d.
//
// Solidity: function ratio() constant returns(uint8)
func (_ConsumeContent *ConsumeContentCallerSession) Ratio() (uint8, error) {
	return _ConsumeContent.Contract.Ratio(&_ConsumeContent.CallOpts)
}

// Consume is a paid mutator transaction binding the contract method 0x1dedc6f7.
//
// Solidity: function consume() returns()
func (_ConsumeContent *ConsumeContentTransactor) Consume(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConsumeContent.contract.Transact(opts, "consume")
}

// Consume is a paid mutator transaction binding the contract method 0x1dedc6f7.
//
// Solidity: function consume() returns()
func (_ConsumeContent *ConsumeContentSession) Consume() (*types.Transaction, error) {
	return _ConsumeContent.Contract.Consume(&_ConsumeContent.TransactOpts)
}

// Consume is a paid mutator transaction binding the contract method 0x1dedc6f7.
//
// Solidity: function consume() returns()
func (_ConsumeContent *ConsumeContentTransactorSession) Consume() (*types.Transaction, error) {
	return _ConsumeContent.Contract.Consume(&_ConsumeContent.TransactOpts)
}

// WithDraw is a paid mutator transaction binding the contract method 0x0fdb1c10.
//
// Solidity: function withDraw() returns()
func (_ConsumeContent *ConsumeContentTransactor) WithDraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConsumeContent.contract.Transact(opts, "withDraw")
}

// WithDraw is a paid mutator transaction binding the contract method 0x0fdb1c10.
//
// Solidity: function withDraw() returns()
func (_ConsumeContent *ConsumeContentSession) WithDraw() (*types.Transaction, error) {
	return _ConsumeContent.Contract.WithDraw(&_ConsumeContent.TransactOpts)
}

// WithDraw is a paid mutator transaction binding the contract method 0x0fdb1c10.
//
// Solidity: function withDraw() returns()
func (_ConsumeContent *ConsumeContentTransactorSession) WithDraw() (*types.Transaction, error) {
	return _ConsumeContent.Contract.WithDraw(&_ConsumeContent.TransactOpts)
}
