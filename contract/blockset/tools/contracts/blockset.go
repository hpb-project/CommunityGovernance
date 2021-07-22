// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ContractsABI is the input ABI used to generate the binding from.
const ContractsABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"},{\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"addProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"deleteAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"},{\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"resetProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newThreshold\",\"type\":\"uint256\"}],\"name\":\"setThreshold\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"voteProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AddAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"DelAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"keywords\",\"type\":\"string\"}],\"name\":\"PassedThreshold\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAdmins\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"getProposolThreshold\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getThreshold\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"getValue\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() view returns(address[])
func (_Contracts *ContractsCaller) GetAdmins(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getAdmins")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() view returns(address[])
func (_Contracts *ContractsSession) GetAdmins() ([]common.Address, error) {
	return _Contracts.Contract.GetAdmins(&_Contracts.CallOpts)
}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() view returns(address[])
func (_Contracts *ContractsCallerSession) GetAdmins() ([]common.Address, error) {
	return _Contracts.Contract.GetAdmins(&_Contracts.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Contracts *ContractsCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Contracts *ContractsSession) GetOwner() (common.Address, error) {
	return _Contracts.Contract.GetOwner(&_Contracts.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Contracts *ContractsCallerSession) GetOwner() (common.Address, error) {
	return _Contracts.Contract.GetOwner(&_Contracts.CallOpts)
}

// GetProposolThreshold is a free data retrieval call binding the contract method 0x35a708ec.
//
// Solidity: function getProposolThreshold(string key) view returns(uint256)
func (_Contracts *ContractsCaller) GetProposolThreshold(opts *bind.CallOpts, key string) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getProposolThreshold", key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProposolThreshold is a free data retrieval call binding the contract method 0x35a708ec.
//
// Solidity: function getProposolThreshold(string key) view returns(uint256)
func (_Contracts *ContractsSession) GetProposolThreshold(key string) (*big.Int, error) {
	return _Contracts.Contract.GetProposolThreshold(&_Contracts.CallOpts, key)
}

// GetProposolThreshold is a free data retrieval call binding the contract method 0x35a708ec.
//
// Solidity: function getProposolThreshold(string key) view returns(uint256)
func (_Contracts *ContractsCallerSession) GetProposolThreshold(key string) (*big.Int, error) {
	return _Contracts.Contract.GetProposolThreshold(&_Contracts.CallOpts, key)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (_Contracts *ContractsCaller) GetThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (_Contracts *ContractsSession) GetThreshold() (*big.Int, error) {
	return _Contracts.Contract.GetThreshold(&_Contracts.CallOpts)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (_Contracts *ContractsCallerSession) GetThreshold() (*big.Int, error) {
	return _Contracts.Contract.GetThreshold(&_Contracts.CallOpts)
}

// GetValue is a free data retrieval call binding the contract method 0x960384a0.
//
// Solidity: function getValue(string key) view returns(uint256)
func (_Contracts *ContractsCaller) GetValue(opts *bind.CallOpts, key string) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getValue", key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValue is a free data retrieval call binding the contract method 0x960384a0.
//
// Solidity: function getValue(string key) view returns(uint256)
func (_Contracts *ContractsSession) GetValue(key string) (*big.Int, error) {
	return _Contracts.Contract.GetValue(&_Contracts.CallOpts, key)
}

// GetValue is a free data retrieval call binding the contract method 0x960384a0.
//
// Solidity: function getValue(string key) view returns(uint256)
func (_Contracts *ContractsCallerSession) GetValue(key string) (*big.Int, error) {
	return _Contracts.Contract.GetValue(&_Contracts.CallOpts, key)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address admin) returns()
func (_Contracts *ContractsTransactor) AddAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "addAdmin", admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address admin) returns()
func (_Contracts *ContractsSession) AddAdmin(admin common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.AddAdmin(&_Contracts.TransactOpts, admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address admin) returns()
func (_Contracts *ContractsTransactorSession) AddAdmin(admin common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.AddAdmin(&_Contracts.TransactOpts, admin)
}

// AddProposal is a paid mutator transaction binding the contract method 0xc208abdc.
//
// Solidity: function addProposal(string key, uint256 number) returns()
func (_Contracts *ContractsTransactor) AddProposal(opts *bind.TransactOpts, key string, number *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "addProposal", key, number)
}

// AddProposal is a paid mutator transaction binding the contract method 0xc208abdc.
//
// Solidity: function addProposal(string key, uint256 number) returns()
func (_Contracts *ContractsSession) AddProposal(key string, number *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AddProposal(&_Contracts.TransactOpts, key, number)
}

// AddProposal is a paid mutator transaction binding the contract method 0xc208abdc.
//
// Solidity: function addProposal(string key, uint256 number) returns()
func (_Contracts *ContractsTransactorSession) AddProposal(key string, number *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AddProposal(&_Contracts.TransactOpts, key, number)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_Contracts *ContractsTransactor) ChangeOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "changeOwner", newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_Contracts *ContractsSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.ChangeOwner(&_Contracts.TransactOpts, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_Contracts *ContractsTransactorSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.ChangeOwner(&_Contracts.TransactOpts, newOwner)
}

// DeleteAdmin is a paid mutator transaction binding the contract method 0x27e1f7df.
//
// Solidity: function deleteAdmin(address admin) returns()
func (_Contracts *ContractsTransactor) DeleteAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "deleteAdmin", admin)
}

// DeleteAdmin is a paid mutator transaction binding the contract method 0x27e1f7df.
//
// Solidity: function deleteAdmin(address admin) returns()
func (_Contracts *ContractsSession) DeleteAdmin(admin common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.DeleteAdmin(&_Contracts.TransactOpts, admin)
}

// DeleteAdmin is a paid mutator transaction binding the contract method 0x27e1f7df.
//
// Solidity: function deleteAdmin(address admin) returns()
func (_Contracts *ContractsTransactorSession) DeleteAdmin(admin common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.DeleteAdmin(&_Contracts.TransactOpts, admin)
}

// ResetProposal is a paid mutator transaction binding the contract method 0x9390ea46.
//
// Solidity: function resetProposal(string key, uint256 number) returns()
func (_Contracts *ContractsTransactor) ResetProposal(opts *bind.TransactOpts, key string, number *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "resetProposal", key, number)
}

// ResetProposal is a paid mutator transaction binding the contract method 0x9390ea46.
//
// Solidity: function resetProposal(string key, uint256 number) returns()
func (_Contracts *ContractsSession) ResetProposal(key string, number *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.ResetProposal(&_Contracts.TransactOpts, key, number)
}

// ResetProposal is a paid mutator transaction binding the contract method 0x9390ea46.
//
// Solidity: function resetProposal(string key, uint256 number) returns()
func (_Contracts *ContractsTransactorSession) ResetProposal(key string, number *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.ResetProposal(&_Contracts.TransactOpts, key, number)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x960bfe04.
//
// Solidity: function setThreshold(uint256 newThreshold) returns()
func (_Contracts *ContractsTransactor) SetThreshold(opts *bind.TransactOpts, newThreshold *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setThreshold", newThreshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x960bfe04.
//
// Solidity: function setThreshold(uint256 newThreshold) returns()
func (_Contracts *ContractsSession) SetThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetThreshold(&_Contracts.TransactOpts, newThreshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x960bfe04.
//
// Solidity: function setThreshold(uint256 newThreshold) returns()
func (_Contracts *ContractsTransactorSession) SetThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetThreshold(&_Contracts.TransactOpts, newThreshold)
}

// VoteProposal is a paid mutator transaction binding the contract method 0x55fd06f1.
//
// Solidity: function voteProposal(string key) returns()
func (_Contracts *ContractsTransactor) VoteProposal(opts *bind.TransactOpts, key string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "voteProposal", key)
}

// VoteProposal is a paid mutator transaction binding the contract method 0x55fd06f1.
//
// Solidity: function voteProposal(string key) returns()
func (_Contracts *ContractsSession) VoteProposal(key string) (*types.Transaction, error) {
	return _Contracts.Contract.VoteProposal(&_Contracts.TransactOpts, key)
}

// VoteProposal is a paid mutator transaction binding the contract method 0x55fd06f1.
//
// Solidity: function voteProposal(string key) returns()
func (_Contracts *ContractsTransactorSession) VoteProposal(key string) (*types.Transaction, error) {
	return _Contracts.Contract.VoteProposal(&_Contracts.TransactOpts, key)
}

// ContractsAddAdminIterator is returned from FilterAddAdmin and is used to iterate over the raw logs and unpacked data for AddAdmin events raised by the Contracts contract.
type ContractsAddAdminIterator struct {
	Event *ContractsAddAdmin // Event containing the contract specifics and raw log

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
func (it *ContractsAddAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsAddAdmin)
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
		it.Event = new(ContractsAddAdmin)
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
func (it *ContractsAddAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsAddAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsAddAdmin represents a AddAdmin event raised by the Contracts contract.
type ContractsAddAdmin struct {
	User     common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAddAdmin is a free log retrieval operation binding the contract event 0x99ad91c72e4fcc0a37d0481b982aec5c0c78491d30d39e65338bec15fa9cc4b7.
//
// Solidity: event AddAdmin(address indexed user, address indexed newAdmin)
func (_Contracts *ContractsFilterer) FilterAddAdmin(opts *bind.FilterOpts, user []common.Address, newAdmin []common.Address) (*ContractsAddAdminIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "AddAdmin", userRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return &ContractsAddAdminIterator{contract: _Contracts.contract, event: "AddAdmin", logs: logs, sub: sub}, nil
}

// WatchAddAdmin is a free log subscription operation binding the contract event 0x99ad91c72e4fcc0a37d0481b982aec5c0c78491d30d39e65338bec15fa9cc4b7.
//
// Solidity: event AddAdmin(address indexed user, address indexed newAdmin)
func (_Contracts *ContractsFilterer) WatchAddAdmin(opts *bind.WatchOpts, sink chan<- *ContractsAddAdmin, user []common.Address, newAdmin []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "AddAdmin", userRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsAddAdmin)
				if err := _Contracts.contract.UnpackLog(event, "AddAdmin", log); err != nil {
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

// ParseAddAdmin is a log parse operation binding the contract event 0x99ad91c72e4fcc0a37d0481b982aec5c0c78491d30d39e65338bec15fa9cc4b7.
//
// Solidity: event AddAdmin(address indexed user, address indexed newAdmin)
func (_Contracts *ContractsFilterer) ParseAddAdmin(log types.Log) (*ContractsAddAdmin, error) {
	event := new(ContractsAddAdmin)
	if err := _Contracts.contract.UnpackLog(event, "AddAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsDelAdminIterator is returned from FilterDelAdmin and is used to iterate over the raw logs and unpacked data for DelAdmin events raised by the Contracts contract.
type ContractsDelAdminIterator struct {
	Event *ContractsDelAdmin // Event containing the contract specifics and raw log

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
func (it *ContractsDelAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsDelAdmin)
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
		it.Event = new(ContractsDelAdmin)
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
func (it *ContractsDelAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsDelAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsDelAdmin represents a DelAdmin event raised by the Contracts contract.
type ContractsDelAdmin struct {
	User  common.Address
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDelAdmin is a free log retrieval operation binding the contract event 0xff210541965015a61784a98072bef6701ce8c6087580100ed42c393224329051.
//
// Solidity: event DelAdmin(address indexed user, address indexed admin)
func (_Contracts *ContractsFilterer) FilterDelAdmin(opts *bind.FilterOpts, user []common.Address, admin []common.Address) (*ContractsDelAdminIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "DelAdmin", userRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &ContractsDelAdminIterator{contract: _Contracts.contract, event: "DelAdmin", logs: logs, sub: sub}, nil
}

// WatchDelAdmin is a free log subscription operation binding the contract event 0xff210541965015a61784a98072bef6701ce8c6087580100ed42c393224329051.
//
// Solidity: event DelAdmin(address indexed user, address indexed admin)
func (_Contracts *ContractsFilterer) WatchDelAdmin(opts *bind.WatchOpts, sink chan<- *ContractsDelAdmin, user []common.Address, admin []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "DelAdmin", userRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsDelAdmin)
				if err := _Contracts.contract.UnpackLog(event, "DelAdmin", log); err != nil {
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

// ParseDelAdmin is a log parse operation binding the contract event 0xff210541965015a61784a98072bef6701ce8c6087580100ed42c393224329051.
//
// Solidity: event DelAdmin(address indexed user, address indexed admin)
func (_Contracts *ContractsFilterer) ParseDelAdmin(log types.Log) (*ContractsDelAdmin, error) {
	event := new(ContractsDelAdmin)
	if err := _Contracts.contract.UnpackLog(event, "DelAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsOwnerSetIterator is returned from FilterOwnerSet and is used to iterate over the raw logs and unpacked data for OwnerSet events raised by the Contracts contract.
type ContractsOwnerSetIterator struct {
	Event *ContractsOwnerSet // Event containing the contract specifics and raw log

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
func (it *ContractsOwnerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsOwnerSet)
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
		it.Event = new(ContractsOwnerSet)
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
func (it *ContractsOwnerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsOwnerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsOwnerSet represents a OwnerSet event raised by the Contracts contract.
type ContractsOwnerSet struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerSet is a free log retrieval operation binding the contract event 0x342827c97908e5e2f71151c08502a66d44b6f758e3ac2f1de95f02eb95f0a735.
//
// Solidity: event OwnerSet(address indexed oldOwner, address indexed newOwner)
func (_Contracts *ContractsFilterer) FilterOwnerSet(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*ContractsOwnerSetIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "OwnerSet", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractsOwnerSetIterator{contract: _Contracts.contract, event: "OwnerSet", logs: logs, sub: sub}, nil
}

// WatchOwnerSet is a free log subscription operation binding the contract event 0x342827c97908e5e2f71151c08502a66d44b6f758e3ac2f1de95f02eb95f0a735.
//
// Solidity: event OwnerSet(address indexed oldOwner, address indexed newOwner)
func (_Contracts *ContractsFilterer) WatchOwnerSet(opts *bind.WatchOpts, sink chan<- *ContractsOwnerSet, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "OwnerSet", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsOwnerSet)
				if err := _Contracts.contract.UnpackLog(event, "OwnerSet", log); err != nil {
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

// ParseOwnerSet is a log parse operation binding the contract event 0x342827c97908e5e2f71151c08502a66d44b6f758e3ac2f1de95f02eb95f0a735.
//
// Solidity: event OwnerSet(address indexed oldOwner, address indexed newOwner)
func (_Contracts *ContractsFilterer) ParseOwnerSet(log types.Log) (*ContractsOwnerSet, error) {
	event := new(ContractsOwnerSet)
	if err := _Contracts.contract.UnpackLog(event, "OwnerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsPassedThresholdIterator is returned from FilterPassedThreshold and is used to iterate over the raw logs and unpacked data for PassedThreshold events raised by the Contracts contract.
type ContractsPassedThresholdIterator struct {
	Event *ContractsPassedThreshold // Event containing the contract specifics and raw log

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
func (it *ContractsPassedThresholdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsPassedThreshold)
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
		it.Event = new(ContractsPassedThreshold)
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
func (it *ContractsPassedThresholdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsPassedThresholdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsPassedThreshold represents a PassedThreshold event raised by the Contracts contract.
type ContractsPassedThreshold struct {
	Keywords common.Hash
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPassedThreshold is a free log retrieval operation binding the contract event 0xcd5e25901a3200193fd4e05510166dd567108536fd9d3b5309934183539298e1.
//
// Solidity: event PassedThreshold(string indexed keywords)
func (_Contracts *ContractsFilterer) FilterPassedThreshold(opts *bind.FilterOpts, keywords []string) (*ContractsPassedThresholdIterator, error) {

	var keywordsRule []interface{}
	for _, keywordsItem := range keywords {
		keywordsRule = append(keywordsRule, keywordsItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "PassedThreshold", keywordsRule)
	if err != nil {
		return nil, err
	}
	return &ContractsPassedThresholdIterator{contract: _Contracts.contract, event: "PassedThreshold", logs: logs, sub: sub}, nil
}

// WatchPassedThreshold is a free log subscription operation binding the contract event 0xcd5e25901a3200193fd4e05510166dd567108536fd9d3b5309934183539298e1.
//
// Solidity: event PassedThreshold(string indexed keywords)
func (_Contracts *ContractsFilterer) WatchPassedThreshold(opts *bind.WatchOpts, sink chan<- *ContractsPassedThreshold, keywords []string) (event.Subscription, error) {

	var keywordsRule []interface{}
	for _, keywordsItem := range keywords {
		keywordsRule = append(keywordsRule, keywordsItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "PassedThreshold", keywordsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsPassedThreshold)
				if err := _Contracts.contract.UnpackLog(event, "PassedThreshold", log); err != nil {
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

// ParsePassedThreshold is a log parse operation binding the contract event 0xcd5e25901a3200193fd4e05510166dd567108536fd9d3b5309934183539298e1.
//
// Solidity: event PassedThreshold(string indexed keywords)
func (_Contracts *ContractsFilterer) ParsePassedThreshold(log types.Log) (*ContractsPassedThreshold, error) {
	event := new(ContractsPassedThreshold)
	if err := _Contracts.contract.UnpackLog(event, "PassedThreshold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
