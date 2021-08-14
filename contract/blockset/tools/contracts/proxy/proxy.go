// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package proxy

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

// ProxyABI is the input ABI used to generate the binding from.
const ProxyABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setcontract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"deleteAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAdmins\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"getValue\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"adminMap\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getcontract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"}]"

// ProxyBin is the compiled bytecode used for deploying new contracts.
var ProxyBin = "0x6080604052336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061007d6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff16610083640100000000026401000000009004565b50610267565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156100e057600080fd5b600073ffffffffffffffffffffffffffffffffffffffff16600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614151561017a57600080fd5b81600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060018290806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505060019050919050565b610ec0806102766000396000f3fe60806040526004361061009e576000357c01000000000000000000000000000000000000000000000000000000009004806307f8f29a146100a357806327e1f7df146100f457806331ae450b1461015d57806370480275146101c9578063893d20e8146102325780638da5cb5b14610289578063960384a0146102e0578063dbbc830b1461037a578063e89408881461040b578063f2fde38b14610462575b600080fd5b3480156100af57600080fd5b506100f2600480360360208110156100c657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506104cb565b005b34801561010057600080fd5b506101436004803603602081101561011757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506105eb565b604051808215151515815260200191505060405180910390f35b34801561016957600080fd5b5061017261094b565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156101b557808201518184015260208101905061019a565b505050509050019250505060405180910390f35b3480156101d557600080fd5b50610218600480360360208110156101ec57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506109d9565b604051808215151515815260200191505060405180910390f35b34801561023e57600080fd5b50610247610bbd565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561029557600080fd5b5061029e610be6565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156102ec57600080fd5b506103646004803603602081101561030357600080fd5b810190808035906020019064010000000081111561032057600080fd5b82018360208201111561033257600080fd5b8035906020019184600183028401116401000000008311171561035457600080fd5b9091929391929390505050610c0b565b6040518082815260200191505060405180910390f35b34801561038657600080fd5b506103c96004803603602081101561039d57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610d0b565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561041757600080fd5b50610420610d3e565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561046e57600080fd5b506104b16004803603602081101561048557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610d68565b604051808215151515815260200191505060405180910390f35b600073ffffffffffffffffffffffffffffffffffffffff16600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415151561056657600080fd5b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561064857600080fd5b600073ffffffffffffffffffffffffffffffffffffffff16600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141515156106e357600080fd5b6000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905560008090505b600180549050811015610941578273ffffffffffffffffffffffffffffffffffffffff166001828154811015156107fd57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614156109345760018080805490500381548110151561085957fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660018281548110151561089357fe5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001808080549050038154811015156108f057fe5b9060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905560018054809190600190036109329190610e43565b505b80806001019150506107ca565b5060019050919050565b606060018054806020026020016040519081016040528092919081815260200182805480156109cf57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610985575b5050505050905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610a3657600080fd5b600073ffffffffffffffffffffffffffffffffffffffff16600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141515610ad057600080fd5b81600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060018290806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505060019050919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663960384a084846040518363ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180806020018281038252848482818152602001925080828437600081840152601f19601f820116905080830192505050935050505060206040518083038186803b158015610cc857600080fd5b505afa158015610cdc573d6000803e3d6000fd5b505050506040513d6020811015610cf257600080fd5b8101908080519060200190929190505050905092915050565b60026020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610dc557600080fd5b816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610e0e826109d9565b50610e396000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff166105eb565b5060019050919050565b815481835581811115610e6a57818360005260206000209182019101610e699190610e6f565b5b505050565b610e9191905b80821115610e8d576000816000905550600101610e75565b5090565b9056fea165627a7a723058209c6cd8420485a971a89a2c3faed3586930fdcae8d802c295d2546c38c7c064480029"

// DeployProxy deploys a new Ethereum contract, binding an instance of Proxy to it.
func DeployProxy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Proxy, error) {
	parsed, err := abi.JSON(strings.NewReader(ProxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ProxyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Proxy{ProxyCaller: ProxyCaller{contract: contract}, ProxyTransactor: ProxyTransactor{contract: contract}, ProxyFilterer: ProxyFilterer{contract: contract}}, nil
}

// Proxy is an auto generated Go binding around an Ethereum contract.
type Proxy struct {
	ProxyCaller     // Read-only binding to the contract
	ProxyTransactor // Write-only binding to the contract
	ProxyFilterer   // Log filterer for contract events
}

// ProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProxySession struct {
	Contract     *Proxy            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProxyCallerSession struct {
	Contract *ProxyCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProxyTransactorSession struct {
	Contract     *ProxyTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProxyRaw struct {
	Contract *Proxy // Generic contract binding to access the raw methods on
}

// ProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProxyCallerRaw struct {
	Contract *ProxyCaller // Generic read-only contract binding to access the raw methods on
}

// ProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProxyTransactorRaw struct {
	Contract *ProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProxy creates a new instance of Proxy, bound to a specific deployed contract.
func NewProxy(address common.Address, backend bind.ContractBackend) (*Proxy, error) {
	contract, err := bindProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Proxy{ProxyCaller: ProxyCaller{contract: contract}, ProxyTransactor: ProxyTransactor{contract: contract}, ProxyFilterer: ProxyFilterer{contract: contract}}, nil
}

// NewProxyCaller creates a new read-only instance of Proxy, bound to a specific deployed contract.
func NewProxyCaller(address common.Address, caller bind.ContractCaller) (*ProxyCaller, error) {
	contract, err := bindProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyCaller{contract: contract}, nil
}

// NewProxyTransactor creates a new write-only instance of Proxy, bound to a specific deployed contract.
func NewProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*ProxyTransactor, error) {
	contract, err := bindProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyTransactor{contract: contract}, nil
}

// NewProxyFilterer creates a new log filterer instance of Proxy, bound to a specific deployed contract.
func NewProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*ProxyFilterer, error) {
	contract, err := bindProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProxyFilterer{contract: contract}, nil
}

// bindProxy binds a generic wrapper to an already deployed contract.
func bindProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proxy *ProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proxy.Contract.ProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proxy *ProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxy.Contract.ProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proxy *ProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proxy.Contract.ProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proxy *ProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proxy *ProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proxy *ProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proxy.Contract.contract.Transact(opts, method, params...)
}

// AdminMap is a free data retrieval call binding the contract method 0xdbbc830b.
//
// Solidity: function adminMap(address ) view returns(address)
func (_Proxy *ProxyCaller) AdminMap(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "adminMap", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AdminMap is a free data retrieval call binding the contract method 0xdbbc830b.
//
// Solidity: function adminMap(address ) view returns(address)
func (_Proxy *ProxySession) AdminMap(arg0 common.Address) (common.Address, error) {
	return _Proxy.Contract.AdminMap(&_Proxy.CallOpts, arg0)
}

// AdminMap is a free data retrieval call binding the contract method 0xdbbc830b.
//
// Solidity: function adminMap(address ) view returns(address)
func (_Proxy *ProxyCallerSession) AdminMap(arg0 common.Address) (common.Address, error) {
	return _Proxy.Contract.AdminMap(&_Proxy.CallOpts, arg0)
}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() view returns(address[])
func (_Proxy *ProxyCaller) GetAdmins(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "getAdmins")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() view returns(address[])
func (_Proxy *ProxySession) GetAdmins() ([]common.Address, error) {
	return _Proxy.Contract.GetAdmins(&_Proxy.CallOpts)
}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() view returns(address[])
func (_Proxy *ProxyCallerSession) GetAdmins() ([]common.Address, error) {
	return _Proxy.Contract.GetAdmins(&_Proxy.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Proxy *ProxyCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Proxy *ProxySession) GetOwner() (common.Address, error) {
	return _Proxy.Contract.GetOwner(&_Proxy.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Proxy *ProxyCallerSession) GetOwner() (common.Address, error) {
	return _Proxy.Contract.GetOwner(&_Proxy.CallOpts)
}

// GetValue is a free data retrieval call binding the contract method 0x960384a0.
//
// Solidity: function getValue(string key) view returns(uint256)
func (_Proxy *ProxyCaller) GetValue(opts *bind.CallOpts, key string) (*big.Int, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "getValue", key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValue is a free data retrieval call binding the contract method 0x960384a0.
//
// Solidity: function getValue(string key) view returns(uint256)
func (_Proxy *ProxySession) GetValue(key string) (*big.Int, error) {
	return _Proxy.Contract.GetValue(&_Proxy.CallOpts, key)
}

// GetValue is a free data retrieval call binding the contract method 0x960384a0.
//
// Solidity: function getValue(string key) view returns(uint256)
func (_Proxy *ProxyCallerSession) GetValue(key string) (*big.Int, error) {
	return _Proxy.Contract.GetValue(&_Proxy.CallOpts, key)
}

// Getcontract is a free data retrieval call binding the contract method 0xe8940888.
//
// Solidity: function getcontract() view returns(address)
func (_Proxy *ProxyCaller) Getcontract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "getcontract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Getcontract is a free data retrieval call binding the contract method 0xe8940888.
//
// Solidity: function getcontract() view returns(address)
func (_Proxy *ProxySession) Getcontract() (common.Address, error) {
	return _Proxy.Contract.Getcontract(&_Proxy.CallOpts)
}

// Getcontract is a free data retrieval call binding the contract method 0xe8940888.
//
// Solidity: function getcontract() view returns(address)
func (_Proxy *ProxyCallerSession) Getcontract() (common.Address, error) {
	return _Proxy.Contract.Getcontract(&_Proxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Proxy *ProxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Proxy *ProxySession) Owner() (common.Address, error) {
	return _Proxy.Contract.Owner(&_Proxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Proxy *ProxyCallerSession) Owner() (common.Address, error) {
	return _Proxy.Contract.Owner(&_Proxy.CallOpts)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address addr) returns(bool)
func (_Proxy *ProxyTransactor) AddAdmin(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "addAdmin", addr)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address addr) returns(bool)
func (_Proxy *ProxySession) AddAdmin(addr common.Address) (*types.Transaction, error) {
	return _Proxy.Contract.AddAdmin(&_Proxy.TransactOpts, addr)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address addr) returns(bool)
func (_Proxy *ProxyTransactorSession) AddAdmin(addr common.Address) (*types.Transaction, error) {
	return _Proxy.Contract.AddAdmin(&_Proxy.TransactOpts, addr)
}

// DeleteAdmin is a paid mutator transaction binding the contract method 0x27e1f7df.
//
// Solidity: function deleteAdmin(address addr) returns(bool)
func (_Proxy *ProxyTransactor) DeleteAdmin(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "deleteAdmin", addr)
}

// DeleteAdmin is a paid mutator transaction binding the contract method 0x27e1f7df.
//
// Solidity: function deleteAdmin(address addr) returns(bool)
func (_Proxy *ProxySession) DeleteAdmin(addr common.Address) (*types.Transaction, error) {
	return _Proxy.Contract.DeleteAdmin(&_Proxy.TransactOpts, addr)
}

// DeleteAdmin is a paid mutator transaction binding the contract method 0x27e1f7df.
//
// Solidity: function deleteAdmin(address addr) returns(bool)
func (_Proxy *ProxyTransactorSession) DeleteAdmin(addr common.Address) (*types.Transaction, error) {
	return _Proxy.Contract.DeleteAdmin(&_Proxy.TransactOpts, addr)
}

// Setcontract is a paid mutator transaction binding the contract method 0x07f8f29a.
//
// Solidity: function setcontract(address addr) returns()
func (_Proxy *ProxyTransactor) Setcontract(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "setcontract", addr)
}

// Setcontract is a paid mutator transaction binding the contract method 0x07f8f29a.
//
// Solidity: function setcontract(address addr) returns()
func (_Proxy *ProxySession) Setcontract(addr common.Address) (*types.Transaction, error) {
	return _Proxy.Contract.Setcontract(&_Proxy.TransactOpts, addr)
}

// Setcontract is a paid mutator transaction binding the contract method 0x07f8f29a.
//
// Solidity: function setcontract(address addr) returns()
func (_Proxy *ProxyTransactorSession) Setcontract(addr common.Address) (*types.Transaction, error) {
	return _Proxy.Contract.Setcontract(&_Proxy.TransactOpts, addr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns(bool)
func (_Proxy *ProxyTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns(bool)
func (_Proxy *ProxySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Proxy.Contract.TransferOwnership(&_Proxy.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns(bool)
func (_Proxy *ProxyTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Proxy.Contract.TransferOwnership(&_Proxy.TransactOpts, newOwner)
}
