// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth

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

// ConsumeContentABI is the input ABI used to generate the binding from.
const ConsumeContentABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withDraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"consume\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"platform\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ratio\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"publisher\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"pPublisher\",\"type\":\"address\"},{\"name\":\"pPlatform\",\"type\":\"address\"},{\"name\":\"pPrice\",\"type\":\"uint32\"},{\"name\":\"pRatio\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"EventConsume\",\"type\":\"event\"}]"

// ConsumeContentBin is the compiled bytecode used for deploying new contracts.
const ConsumeContentBin = `0x6060604052341561000f57600080fd5b60405160808061042d8339810160405280805191906020018051919060200180519190602001805160008054600160a060020a03978816600160a060020a0319918216179091556001805460ff90931678010000000000000000000000000000000000000000000000000260c060020a60ff021963ffffffff909716740100000000000000000000000000000000000000000260a060020a63ffffffff021998909916939092169290921795909516959095179290921692909217909255505061034f806100de6000396000f3006060604052600436106100825763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306661abd81146100875780630fdb1c10146100ac5780631dedc6f7146100c15780634bde38c8146100c957806371ca337d146100f85780638c72c54e14610121578063a035b1fe14610134575b600080fd5b341561009257600080fd5b61009a610160565b60405190815260200160405180910390f35b34156100b757600080fd5b6100bf610166565b005b6100bf6101d7565b34156100d457600080fd5b6100dc6102bc565b604051600160a060020a03909116815260200160405180910390f35b341561010357600080fd5b61010b6102cb565b60405160ff909116815260200160405180910390f35b341561012c57600080fd5b6100dc6102f0565b341561013f57600080fd5b6101476102ff565b60405163ffffffff909116815260200160405180910390f35b60025481565b600160a060020a03331660009081526003602052604081205481901161018b57600080fd5b50600160a060020a033316600081815260036020526040808220805492905590919082156108fc0290839051600060405180830381858888f1935050505015156101d457600080fd5b50565b60015460009074010000000000000000000000000000000000000000900463ffffffff16341461020657600080fd5b600280546001908101909155546064907801000000000000000000000000000000000000000000000000900460ff16340260008054600160a060020a039081168252600360205260408083208054959094049485019093556001541681528190208054348490030190556002549192507fe7092102cafaeb2e518d7d55510eec03f722fc75b987ac4d6b7fb8ea57d63f6791339151600160a060020a03909216825260208201526040908101905180910390a150565b600154600160a060020a031681565b6001547801000000000000000000000000000000000000000000000000900460ff1681565b600054600160a060020a031681565b60015474010000000000000000000000000000000000000000900463ffffffff16815600a165627a7a72305820320b341b56ea34622013a51e474b65e2bb011111290d4934f077dceb3a9b313a0029`

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

// ConsumeContentEventConsumeIterator is returned from FilterEventConsume and is used to iterate over the raw logs and unpacked data for EventConsume events raised by the ConsumeContent contract.
type ConsumeContentEventConsumeIterator struct {
	Event *ConsumeContentEventConsume // Event containing the contract specifics and raw log

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
func (it *ConsumeContentEventConsumeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConsumeContentEventConsume)
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
		it.Event = new(ConsumeContentEventConsume)
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
func (it *ConsumeContentEventConsumeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConsumeContentEventConsumeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConsumeContentEventConsume represents a EventConsume event raised by the ConsumeContent contract.
type ConsumeContentEventConsume struct {
	User  common.Address
	Count *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEventConsume is a free log retrieval operation binding the contract event 0xe7092102cafaeb2e518d7d55510eec03f722fc75b987ac4d6b7fb8ea57d63f67.
//
// Solidity: e EventConsume(user address, count uint256)
func (_ConsumeContent *ConsumeContentFilterer) FilterEventConsume(opts *bind.FilterOpts) (*ConsumeContentEventConsumeIterator, error) {

	logs, sub, err := _ConsumeContent.contract.FilterLogs(opts, "EventConsume")
	if err != nil {
		return nil, err
	}
	return &ConsumeContentEventConsumeIterator{contract: _ConsumeContent.contract, event: "EventConsume", logs: logs, sub: sub}, nil
}

// WatchEventConsume is a free log subscription operation binding the contract event 0xe7092102cafaeb2e518d7d55510eec03f722fc75b987ac4d6b7fb8ea57d63f67.
//
// Solidity: e EventConsume(user address, count uint256)
func (_ConsumeContent *ConsumeContentFilterer) WatchEventConsume(opts *bind.WatchOpts, sink chan<- *ConsumeContentEventConsume) (event.Subscription, error) {

	logs, sub, err := _ConsumeContent.contract.WatchLogs(opts, "EventConsume")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConsumeContentEventConsume)
				if err := _ConsumeContent.contract.UnpackLog(event, "EventConsume", log); err != nil {
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
