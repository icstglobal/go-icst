// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethereum

import (
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// ConsumeSkillABI is the input ABI used to generate the binding from.
const ConsumeSkillABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"Complete\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hash\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withDraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"abort\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"platform\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ratio\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"publisher\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"consumer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"start\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"pHash\",\"type\":\"string\"},{\"name\":\"pPublisher\",\"type\":\"address\"},{\"name\":\"pPlatform\",\"type\":\"address\"},{\"name\":\"pConsumer\",\"type\":\"address\"},{\"name\":\"pPrice\",\"type\":\"uint32\"},{\"name\":\"pRatio\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"actor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"from\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"uint8\"}],\"name\":\"StateChange\",\"type\":\"event\"}]"

// ConsumeSkillBin is the compiled bytecode used for deploying new contracts.
const ConsumeSkillBin = `0x6060604052341561000f57600080fd5b60405161092e38038061092e8339810160405280805182019190602001805191906020018051919060200180519190602001805191906020018051915060039050868051610061929160200190610114565b5060008054600160a060020a0319908116600160a060020a03888116919091178355600180548316888316178155600280549093169187169190911760a060020a63ffffffff0219167401000000000000000000000000000000000000000063ffffffff8716021760c060020a60ff021916780100000000000000000000000000000000000000000000000060ff8616021790915560058054909160ff19909116908302179055505050505050506101af565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061015557805160ff1916838001178555610182565b82800160010185558215610182579182015b82811115610182578251825591602001919060010190610167565b5061018e929150610192565b5090565b6101ac91905b8082111561018e5760008155600101610198565b90565b610770806101be6000396000f3006060604052600436106100ae5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166301b7dcb481146100b057806309bd5a60146100c35780630fdb1c101461014d57806335a063b4146101605780634bde38c81461017357806371ca337d146101a25780638c72c54e146101cb578063a035b1fe146101de578063b4fd72961461020a578063be9a65551461021d578063c19d93fb14610225575b005b34156100bb57600080fd5b6100ae61025c565b34156100ce57600080fd5b6100d66103b9565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156101125780820151838201526020016100fa565b50505050905090810190601f16801561013f5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561015857600080fd5b6100ae610457565b341561016b57600080fd5b6100ae610501565b341561017e57600080fd5b6101866105b7565b604051600160a060020a03909116815260200160405180910390f35b34156101ad57600080fd5b6101b56105c6565b60405160ff909116815260200160405180910390f35b34156101d657600080fd5b6101866105eb565b34156101e957600080fd5b6101f16105fa565b60405163ffffffff909116815260200160405180910390f35b341561021557600080fd5b61018661061e565b6100ae61062d565b341561023057600080fd5b61023861073b565b6040518082600381111561024857fe5b60ff16815260200191505060405180910390f35b60408051908101604052600054600160a060020a0390811682526002541660208201528051600160a060020a031633600160a060020a031614806102b557506020810151600160a060020a031633600160a060020a0316145b15156102c057600080fd5b600160055460ff1660038111156102d357fe5b146102dd576103b6565b600160a060020a03338116600090815260066020526040808220805460ff1916600117905581549092168152205460ff1680156103345750600254600160a060020a031660009081526006602052604090205460ff165b15610347576005805460ff191660021790555b7fedab99fa21b022682599c9a93397843b51e5a93cec30fc4940f231578a32bd0e3360016002604051600160a060020a03841681526020810183600381111561038c57fe5b60ff1681526020018260038111156103a057fe5b60ff168152602001935050505060405180910390a15b50565b60038054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561044f5780601f106104245761010080835404028352916020019161044f565b820191906000526020600020905b81548152906001019060200180831161043257829003601f168201915b505050505081565b6000600260055460ff16600381111561046c57fe5b1415801561048b5750600360055460ff16600381111561048857fe5b14155b15610495576103b6565b600160a060020a033316600090815260046020526040812054116104b857600080fd5b50600160a060020a033316600081815260046020526040808220805492905590919082156108fc0290839051600060405180830381858888f1935050505015156103b657600080fd5b60005433600160a060020a0390811691161461051c57600080fd5b600060055460ff16600381111561052f57fe5b14610539576105b5565b6005805460ff191660021790557fedab99fa21b022682599c9a93397843b51e5a93cec30fc4940f231578a32bd0e3360006003604051600160a060020a03841681526020810183600381111561058b57fe5b60ff16815260200182600381111561059f57fe5b60ff168152602001935050505060405180910390a15b565b600154600160a060020a031681565b6002547801000000000000000000000000000000000000000000000000900460ff1681565b600054600160a060020a031681565b60025474010000000000000000000000000000000000000000900463ffffffff1681565b600254600160a060020a031681565b60025460009033600160a060020a0390811691161461064b57600080fd5b600060055460ff16600381111561065e57fe5b14610668576103b6565b60025474010000000000000000000000000000000000000000900463ffffffff16341461069457600080fd5b6002546064907801000000000000000000000000000000000000000000000000900460ff16340260008054600160a060020a039081168252600460205260408083208054959094049485019093556001805490911682528282208054348690030190559293507fedab99fa21b022682599c9a93397843b51e5a93cec30fc4940f231578a32bd0e92339251600160a060020a03841681526020810183600381111561038c57fe5b60055460ff16815600a165627a7a72305820e5479a8ade0e893064c646b902d0cc81e7a370dc6fb12688ed7803d72c8774000029`

// DeployConsumeSkill deploys a new Ethereum contract, binding an instance of ConsumeSkill to it.
func DeployConsumeSkill(auth *bind.TransactOpts, backend bind.ContractBackend, pHash string, pPublisher common.Address, pPlatform common.Address, pConsumer common.Address, pPrice uint32, pRatio uint8) (common.Address, *types.Transaction, *ConsumeSkill, error) {
	parsed, err := abi.JSON(strings.NewReader(ConsumeSkillABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ConsumeSkillBin), backend, pHash, pPublisher, pPlatform, pConsumer, pPrice, pRatio)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ConsumeSkill{ConsumeSkillCaller: ConsumeSkillCaller{contract: contract}, ConsumeSkillTransactor: ConsumeSkillTransactor{contract: contract}, ConsumeSkillFilterer: ConsumeSkillFilterer{contract: contract}}, nil
}

// ConsumeSkill is an auto generated Go binding around an Ethereum contract.
type ConsumeSkill struct {
	ConsumeSkillCaller     // Read-only binding to the contract
	ConsumeSkillTransactor // Write-only binding to the contract
	ConsumeSkillFilterer   // Log filterer for contract events
}

// ConsumeSkillCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConsumeSkillCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConsumeSkillTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConsumeSkillTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConsumeSkillFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConsumeSkillFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConsumeSkillSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConsumeSkillSession struct {
	Contract     *ConsumeSkill     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConsumeSkillCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConsumeSkillCallerSession struct {
	Contract *ConsumeSkillCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ConsumeSkillTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConsumeSkillTransactorSession struct {
	Contract     *ConsumeSkillTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ConsumeSkillRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConsumeSkillRaw struct {
	Contract *ConsumeSkill // Generic contract binding to access the raw methods on
}

// ConsumeSkillCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConsumeSkillCallerRaw struct {
	Contract *ConsumeSkillCaller // Generic read-only contract binding to access the raw methods on
}

// ConsumeSkillTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConsumeSkillTransactorRaw struct {
	Contract *ConsumeSkillTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConsumeSkill creates a new instance of ConsumeSkill, bound to a specific deployed contract.
func NewConsumeSkill(address common.Address, backend bind.ContractBackend) (*ConsumeSkill, error) {
	contract, err := bindConsumeSkill(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ConsumeSkill{ConsumeSkillCaller: ConsumeSkillCaller{contract: contract}, ConsumeSkillTransactor: ConsumeSkillTransactor{contract: contract}, ConsumeSkillFilterer: ConsumeSkillFilterer{contract: contract}}, nil
}

// NewConsumeSkillCaller creates a new read-only instance of ConsumeSkill, bound to a specific deployed contract.
func NewConsumeSkillCaller(address common.Address, caller bind.ContractCaller) (*ConsumeSkillCaller, error) {
	contract, err := bindConsumeSkill(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConsumeSkillCaller{contract: contract}, nil
}

// NewConsumeSkillTransactor creates a new write-only instance of ConsumeSkill, bound to a specific deployed contract.
func NewConsumeSkillTransactor(address common.Address, transactor bind.ContractTransactor) (*ConsumeSkillTransactor, error) {
	contract, err := bindConsumeSkill(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConsumeSkillTransactor{contract: contract}, nil
}

// NewConsumeSkillFilterer creates a new log filterer instance of ConsumeSkill, bound to a specific deployed contract.
func NewConsumeSkillFilterer(address common.Address, filterer bind.ContractFilterer) (*ConsumeSkillFilterer, error) {
	contract, err := bindConsumeSkill(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConsumeSkillFilterer{contract: contract}, nil
}

// bindConsumeSkill binds a generic wrapper to an already deployed contract.
func bindConsumeSkill(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConsumeSkillABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConsumeSkill *ConsumeSkillRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ConsumeSkill.Contract.ConsumeSkillCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConsumeSkill *ConsumeSkillRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConsumeSkill.Contract.ConsumeSkillTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConsumeSkill *ConsumeSkillRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConsumeSkill.Contract.ConsumeSkillTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConsumeSkill *ConsumeSkillCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ConsumeSkill.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConsumeSkill *ConsumeSkillTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConsumeSkill.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConsumeSkill *ConsumeSkillTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConsumeSkill.Contract.contract.Transact(opts, method, params...)
}

// Consumer is a free data retrieval call binding the contract method 0xb4fd7296.
//
// Solidity: function consumer() constant returns(address)
func (_ConsumeSkill *ConsumeSkillCaller) Consumer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ConsumeSkill.contract.Call(opts, out, "consumer")
	return *ret0, err
}

// Consumer is a free data retrieval call binding the contract method 0xb4fd7296.
//
// Solidity: function consumer() constant returns(address)
func (_ConsumeSkill *ConsumeSkillSession) Consumer() (common.Address, error) {
	return _ConsumeSkill.Contract.Consumer(&_ConsumeSkill.CallOpts)
}

// Consumer is a free data retrieval call binding the contract method 0xb4fd7296.
//
// Solidity: function consumer() constant returns(address)
func (_ConsumeSkill *ConsumeSkillCallerSession) Consumer() (common.Address, error) {
	return _ConsumeSkill.Contract.Consumer(&_ConsumeSkill.CallOpts)
}

// Hash is a free data retrieval call binding the contract method 0x09bd5a60.
//
// Solidity: function hash() constant returns(string)
func (_ConsumeSkill *ConsumeSkillCaller) Hash(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ConsumeSkill.contract.Call(opts, out, "hash")
	return *ret0, err
}

// Hash is a free data retrieval call binding the contract method 0x09bd5a60.
//
// Solidity: function hash() constant returns(string)
func (_ConsumeSkill *ConsumeSkillSession) Hash() (string, error) {
	return _ConsumeSkill.Contract.Hash(&_ConsumeSkill.CallOpts)
}

// Hash is a free data retrieval call binding the contract method 0x09bd5a60.
//
// Solidity: function hash() constant returns(string)
func (_ConsumeSkill *ConsumeSkillCallerSession) Hash() (string, error) {
	return _ConsumeSkill.Contract.Hash(&_ConsumeSkill.CallOpts)
}

// Platform is a free data retrieval call binding the contract method 0x4bde38c8.
//
// Solidity: function platform() constant returns(address)
func (_ConsumeSkill *ConsumeSkillCaller) Platform(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ConsumeSkill.contract.Call(opts, out, "platform")
	return *ret0, err
}

// Platform is a free data retrieval call binding the contract method 0x4bde38c8.
//
// Solidity: function platform() constant returns(address)
func (_ConsumeSkill *ConsumeSkillSession) Platform() (common.Address, error) {
	return _ConsumeSkill.Contract.Platform(&_ConsumeSkill.CallOpts)
}

// Platform is a free data retrieval call binding the contract method 0x4bde38c8.
//
// Solidity: function platform() constant returns(address)
func (_ConsumeSkill *ConsumeSkillCallerSession) Platform() (common.Address, error) {
	return _ConsumeSkill.Contract.Platform(&_ConsumeSkill.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint32)
func (_ConsumeSkill *ConsumeSkillCaller) Price(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _ConsumeSkill.contract.Call(opts, out, "price")
	return *ret0, err
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint32)
func (_ConsumeSkill *ConsumeSkillSession) Price() (uint32, error) {
	return _ConsumeSkill.Contract.Price(&_ConsumeSkill.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint32)
func (_ConsumeSkill *ConsumeSkillCallerSession) Price() (uint32, error) {
	return _ConsumeSkill.Contract.Price(&_ConsumeSkill.CallOpts)
}

// Publisher is a free data retrieval call binding the contract method 0x8c72c54e.
//
// Solidity: function publisher() constant returns(address)
func (_ConsumeSkill *ConsumeSkillCaller) Publisher(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ConsumeSkill.contract.Call(opts, out, "publisher")
	return *ret0, err
}

// Publisher is a free data retrieval call binding the contract method 0x8c72c54e.
//
// Solidity: function publisher() constant returns(address)
func (_ConsumeSkill *ConsumeSkillSession) Publisher() (common.Address, error) {
	return _ConsumeSkill.Contract.Publisher(&_ConsumeSkill.CallOpts)
}

// Publisher is a free data retrieval call binding the contract method 0x8c72c54e.
//
// Solidity: function publisher() constant returns(address)
func (_ConsumeSkill *ConsumeSkillCallerSession) Publisher() (common.Address, error) {
	return _ConsumeSkill.Contract.Publisher(&_ConsumeSkill.CallOpts)
}

// Ratio is a free data retrieval call binding the contract method 0x71ca337d.
//
// Solidity: function ratio() constant returns(uint8)
func (_ConsumeSkill *ConsumeSkillCaller) Ratio(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ConsumeSkill.contract.Call(opts, out, "ratio")
	return *ret0, err
}

// Ratio is a free data retrieval call binding the contract method 0x71ca337d.
//
// Solidity: function ratio() constant returns(uint8)
func (_ConsumeSkill *ConsumeSkillSession) Ratio() (uint8, error) {
	return _ConsumeSkill.Contract.Ratio(&_ConsumeSkill.CallOpts)
}

// Ratio is a free data retrieval call binding the contract method 0x71ca337d.
//
// Solidity: function ratio() constant returns(uint8)
func (_ConsumeSkill *ConsumeSkillCallerSession) Ratio() (uint8, error) {
	return _ConsumeSkill.Contract.Ratio(&_ConsumeSkill.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() constant returns(uint8)
func (_ConsumeSkill *ConsumeSkillCaller) State(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ConsumeSkill.contract.Call(opts, out, "state")
	return *ret0, err
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() constant returns(uint8)
func (_ConsumeSkill *ConsumeSkillSession) State() (uint8, error) {
	return _ConsumeSkill.Contract.State(&_ConsumeSkill.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() constant returns(uint8)
func (_ConsumeSkill *ConsumeSkillCallerSession) State() (uint8, error) {
	return _ConsumeSkill.Contract.State(&_ConsumeSkill.CallOpts)
}

// Complete is a paid mutator transaction binding the contract method 0x01b7dcb4.
//
// Solidity: function Complete() returns()
func (_ConsumeSkill *ConsumeSkillTransactor) Complete(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConsumeSkill.contract.Transact(opts, "Complete")
}

// Complete is a paid mutator transaction binding the contract method 0x01b7dcb4.
//
// Solidity: function Complete() returns()
func (_ConsumeSkill *ConsumeSkillSession) Complete() (*types.Transaction, error) {
	return _ConsumeSkill.Contract.Complete(&_ConsumeSkill.TransactOpts)
}

// Complete is a paid mutator transaction binding the contract method 0x01b7dcb4.
//
// Solidity: function Complete() returns()
func (_ConsumeSkill *ConsumeSkillTransactorSession) Complete() (*types.Transaction, error) {
	return _ConsumeSkill.Contract.Complete(&_ConsumeSkill.TransactOpts)
}

// Abort is a paid mutator transaction binding the contract method 0x35a063b4.
//
// Solidity: function abort() returns()
func (_ConsumeSkill *ConsumeSkillTransactor) Abort(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConsumeSkill.contract.Transact(opts, "abort")
}

// Abort is a paid mutator transaction binding the contract method 0x35a063b4.
//
// Solidity: function abort() returns()
func (_ConsumeSkill *ConsumeSkillSession) Abort() (*types.Transaction, error) {
	return _ConsumeSkill.Contract.Abort(&_ConsumeSkill.TransactOpts)
}

// Abort is a paid mutator transaction binding the contract method 0x35a063b4.
//
// Solidity: function abort() returns()
func (_ConsumeSkill *ConsumeSkillTransactorSession) Abort() (*types.Transaction, error) {
	return _ConsumeSkill.Contract.Abort(&_ConsumeSkill.TransactOpts)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_ConsumeSkill *ConsumeSkillTransactor) Start(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConsumeSkill.contract.Transact(opts, "start")
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_ConsumeSkill *ConsumeSkillSession) Start() (*types.Transaction, error) {
	return _ConsumeSkill.Contract.Start(&_ConsumeSkill.TransactOpts)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_ConsumeSkill *ConsumeSkillTransactorSession) Start() (*types.Transaction, error) {
	return _ConsumeSkill.Contract.Start(&_ConsumeSkill.TransactOpts)
}

// WithDraw is a paid mutator transaction binding the contract method 0x0fdb1c10.
//
// Solidity: function withDraw() returns()
func (_ConsumeSkill *ConsumeSkillTransactor) WithDraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConsumeSkill.contract.Transact(opts, "withDraw")
}

// WithDraw is a paid mutator transaction binding the contract method 0x0fdb1c10.
//
// Solidity: function withDraw() returns()
func (_ConsumeSkill *ConsumeSkillSession) WithDraw() (*types.Transaction, error) {
	return _ConsumeSkill.Contract.WithDraw(&_ConsumeSkill.TransactOpts)
}

// WithDraw is a paid mutator transaction binding the contract method 0x0fdb1c10.
//
// Solidity: function withDraw() returns()
func (_ConsumeSkill *ConsumeSkillTransactorSession) WithDraw() (*types.Transaction, error) {
	return _ConsumeSkill.Contract.WithDraw(&_ConsumeSkill.TransactOpts)
}

// ConsumeSkillStateChangeIterator is returned from FilterStateChange and is used to iterate over the raw logs and unpacked data for StateChange events raised by the ConsumeSkill contract.
type ConsumeSkillStateChangeIterator struct {
	Event *ConsumeSkillStateChange // Event containing the contract specifics and raw log

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
func (it *ConsumeSkillStateChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConsumeSkillStateChange)
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
		it.Event = new(ConsumeSkillStateChange)
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
func (it *ConsumeSkillStateChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConsumeSkillStateChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConsumeSkillStateChange represents a StateChange event raised by the ConsumeSkill contract.
type ConsumeSkillStateChange struct {
	Actor common.Address
	From  uint8
	To    uint8
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStateChange is a free log retrieval operation binding the contract event 0xedab99fa21b022682599c9a93397843b51e5a93cec30fc4940f231578a32bd0e.
//
// Solidity: event StateChange(actor address, from uint8, to uint8)
func (_ConsumeSkill *ConsumeSkillFilterer) FilterStateChange(opts *bind.FilterOpts) (*ConsumeSkillStateChangeIterator, error) {

	logs, sub, err := _ConsumeSkill.contract.FilterLogs(opts, "StateChange")
	if err != nil {
		return nil, err
	}
	return &ConsumeSkillStateChangeIterator{contract: _ConsumeSkill.contract, event: "StateChange", logs: logs, sub: sub}, nil
}

// WatchStateChange is a free log subscription operation binding the contract event 0xedab99fa21b022682599c9a93397843b51e5a93cec30fc4940f231578a32bd0e.
//
// Solidity: event StateChange(actor address, from uint8, to uint8)
func (_ConsumeSkill *ConsumeSkillFilterer) WatchStateChange(opts *bind.WatchOpts, sink chan<- *ConsumeSkillStateChange) (event.Subscription, error) {

	logs, sub, err := _ConsumeSkill.contract.WatchLogs(opts, "StateChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConsumeSkillStateChange)
				if err := _ConsumeSkill.contract.UnpackLog(event, "StateChange", log); err != nil {
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
