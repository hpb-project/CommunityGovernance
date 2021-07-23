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
const ContractsABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"deleteAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAdmins\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"getProposalThreshold\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"voteProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"},{\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"resetProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"getValue\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newThreshold\",\"type\":\"uint256\"}],\"name\":\"setThreshold\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"},{\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"addProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getThreshold\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AddAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"DelAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"keywords\",\"type\":\"string\"}],\"name\":\"PassedThreshold\",\"type\":\"event\"}]"

// ContractsBin is the compiled bytecode used for deploying new contracts.
var ContractsBin = "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600180819055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f342827c97908e5e2f71151c08502a66d44b6f758e3ac2f1de95f02eb95f0a73560405160405180910390a3611fc2806100e36000396000f3fe6080604052600436106100b4576000357c01000000000000000000000000000000000000000000000000000000009004806327e1f7df146100b957806331ae450b1461010a578063481a59a21461017657806355fd06f114610252578063704802751461031a578063893d20e81461036b5780639390ea46146103c2578063960384a014610494578063960bfe041461052e578063a6f9dae114610569578063c208abdc146105ba578063e75235b81461068c575b600080fd5b3480156100c557600080fd5b50610108600480360360208110156100dc57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506106b7565b005b34801561011657600080fd5b5061011f610a2f565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b83811015610162578082015181840152602081019050610147565b505050509050019250505060405180910390f35b34801561018257600080fd5b5061023c6004803603602081101561019957600080fd5b81019080803590602001906401000000008111156101b657600080fd5b8201836020820111156101c857600080fd5b803590602001918460018302840111640100000000831117156101ea57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610abd565b6040518082815260200191505060405180910390f35b34801561025e57600080fd5b506103186004803603602081101561027557600080fd5b810190808035906020019064010000000081111561029257600080fd5b8201836020820111156102a457600080fd5b803590602001918460018302840111640100000000831117156102c657600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610c1a565b005b34801561032657600080fd5b506103696004803603602081101561033d57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506111be565b005b34801561037757600080fd5b506103806113f1565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156103ce57600080fd5b50610492600480360360408110156103e557600080fd5b810190808035906020019064010000000081111561040257600080fd5b82018360208201111561041457600080fd5b8035906020019184600183028401116401000000008311171561043657600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019092919050505061141a565b005b3480156104a057600080fd5b50610518600480360360208110156104b757600080fd5b81019080803590602001906401000000008111156104d457600080fd5b8201836020820111156104e657600080fd5b8035906020019184600183028401116401000000008311171561050857600080fd5b9091929391929390505050611659565b6040518082815260200191505060405180910390f35b34801561053a57600080fd5b506105676004803603602081101561055157600080fd5b81019080803590602001909291905050506116d1565b005b34801561057557600080fd5b506105b86004803603602081101561058c57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611849565b005b3480156105c657600080fd5b5061068a600480360360408110156105dd57600080fd5b81019080803590602001906401000000008111156105fa57600080fd5b82018360208201111561060c57600080fd5b8035906020019184600183028401116401000000008311171561062e57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803590602001909291905050506119cb565b005b34801561069857600080fd5b506106a1611e27565b6040518082815260200191505060405180910390f35b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561077b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f43616c6c6572206973206e6f74206f776e65720000000000000000000000000081525060200191505060405180910390fd5b600560008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615610a2c5760008090505b6003805490508110156109fb578173ffffffffffffffffffffffffffffffffffffffff1660038281548110151561080657fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614156109ee57600360016003805490500381548110151561086457fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660038281548110151561089e57fe5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060036001600380549050038154811015156108fd57fe5b9060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81549060ff0219169055600380548091906001900361098e9190611e31565b508173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fff210541965015a61784a98072bef6701ce8c6087580100ed42c39322432905160405160405180910390a36109fb565b80806001019150506107d3565b506003805490506001541115610a2b5760038054905060018190555060006001541415610a2a57600180819055505b5b5b50565b60606003805480602002602001604051908101604052809291908181526020018280548015610ab357602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610a69575b5050505050905090565b6000806002836040518082805190602001908083835b602083101515610af85780518252602082019150602081019050602083039250610ad3565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060020154111515610ba4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f70726f706f736f6c206e6f74206578697374000000000000000000000000000081525060200191505060405180910390fd5b6002826040518082805190602001908083835b602083101515610bdc5780518252602082019150602081019050602083039250610bb7565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020600201549050919050565b600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1680610cbe57506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b1515610d32576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f43616c6c6572206973206e6f74206f776e657220616e642061646d696e00000081525060200191505060405180910390fd5b60006002826040518082805190602001908083835b602083101515610d6c5780518252602082019150602081019050602083039250610d47565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060020154111515610e18576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f4e6f7420666f756e642070726f706f73616c000000000000000000000000000081525060200191505060405180910390fd5b60008090505b6006826040518082805190602001908083835b602083101515610e565780518252602082019150602081019050602083039250610e31565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902080549050811015610f7d573373ffffffffffffffffffffffffffffffffffffffff166006836040518082805190602001908083835b602083101515610ee35780518252602082019150602081019050602083039250610ebe565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902082815481101515610f2357fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415610f7057506111bb565b8080600101915050610e1e565b506006816040518082805190602001908083835b602083101515610fb65780518252602082019150602081019050602083039250610f91565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390203390806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506002816040518082805190602001908083835b6020831015156110855780518252602082019150602081019050602083039250611060565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020600201546006826040518082805190602001908083835b6020831015156110f457805182526020820191506020810190506020830392506110cf565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020805490501015156111ba5760016002826040518082805190602001908083835b60208310151561116c5780518252602082019150602081019050602083039250611147565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060030160006101000a81548160ff0219169083151502179055505b5b50565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611282576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f43616c6c6572206973206e6f74206f776e65720000000000000000000000000081525060200191505060405180910390fd5b600560008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615156113ee576001600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555060038190806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550508073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f99ad91c72e4fcc0a37d0481b982aec5c0c78491d30d39e65338bec15fa9cc4b760405160405180910390a35b50565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16806114be57506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b1515611532576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f43616c6c6572206973206e6f74206f776e657220616e642061646d696e00000081525060200191505060405180910390fd5b6002826040518082805190602001908083835b60208310151561156a5780518252602082019150602081019050602083039250611545565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020600080820160006115ae9190611e5d565b600182016000905560028201600090556003820160006101000a81549060ff021916905550506006826040518082805190602001908083835b60208310151561160c57805182526020820191506020810190506020830392506115e7565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020600061164b9190611ea5565b61165582826119cb565b5050565b600060028383604051808383808284378083019250505092505050908152602001604051809103902060030160009054906101000a900460ff16151561169e57600080fd5b60028383604051808383808284378083019250505092505050908152602001604051809103902060010154905092915050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611795576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f43616c6c6572206973206e6f74206f776e65720000000000000000000000000081525060200191505060405180910390fd5b60008111151561180d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f7468726573686f6c64203d3d203000000000000000000000000000000000000081525060200191505060405180910390fd5b60038054905081111561182b57600380549050600181905550611833565b806001819055505b6000600154141561184657600180819055505b50565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561190d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f43616c6c6572206973206e6f74206f776e65720000000000000000000000000081525060200191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f342827c97908e5e2f71151c08502a66d44b6f758e3ac2f1de95f02eb95f0a73560405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1680611a6f57506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b1515611ae3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f43616c6c6572206973206e6f74206f776e657220616e642061646d696e00000081525060200191505060405180910390fd5b60006002836040518082805190602001908083835b602083101515611b1d5780518252602082019150602081019050602083039250611af8565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060020154141515611bc9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f4e6f74206e65772050726f706f73616c0000000000000000000000000000000081525060200191505060405180910390fd5b611bd1611ec6565b828160000181905250600154816040018181525050818160200181815250506006836040518082805190602001908083835b602083101515611c285780518252602082019150602081019050602083039250611c03565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390203390806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505080604001516006846040518082805190602001908083835b602083101515611cfc5780518252602082019150602081019050602083039250611cd7565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902080549050101515611d50576001816060019015159081151581525050611d62565b60008160600190151590811515815250505b806002846040518082805190602001908083835b602083101515611d9b5780518252602082019150602081019050602083039250611d76565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390206000820151816000019080519060200190611dea929190611ef1565b50602082015181600101556040820151816002015560608201518160030160006101000a81548160ff021916908315150217905550905050505050565b6000600154905090565b815481835581811115611e5857818360005260206000209182019101611e579190611f71565b5b505050565b50805460018160011615610100020316600290046000825580601f10611e835750611ea2565b601f016020900490600052602060002090810190611ea19190611f71565b5b50565b5080546000825590600052602060002090810190611ec39190611f71565b50565b6080604051908101604052806060815260200160008152602001600081526020016000151581525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611f3257805160ff1916838001178555611f60565b82800160010185558215611f60579182015b82811115611f5f578251825591602001919060010190611f44565b5b509050611f6d9190611f71565b5090565b611f9391905b80821115611f8f576000816000905550600101611f77565b5090565b9056fea165627a7a72305820292823d82b20ad2926c5593dafc32c88fda94be46ae42a5db763c360f428f6780029"

// DeployContracts deploys a new Ethereum contract, binding an instance of Contracts to it.
func DeployContracts(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contracts, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ContractsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

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

// GetProposalThreshold is a free data retrieval call binding the contract method 0x481a59a2.
//
// Solidity: function getProposalThreshold(string key) view returns(uint256)
func (_Contracts *ContractsCaller) GetProposalThreshold(opts *bind.CallOpts, key string) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getProposalThreshold", key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProposalThreshold is a free data retrieval call binding the contract method 0x481a59a2.
//
// Solidity: function getProposalThreshold(string key) view returns(uint256)
func (_Contracts *ContractsSession) GetProposalThreshold(key string) (*big.Int, error) {
	return _Contracts.Contract.GetProposalThreshold(&_Contracts.CallOpts, key)
}

// GetProposalThreshold is a free data retrieval call binding the contract method 0x481a59a2.
//
// Solidity: function getProposalThreshold(string key) view returns(uint256)
func (_Contracts *ContractsCallerSession) GetProposalThreshold(key string) (*big.Int, error) {
	return _Contracts.Contract.GetProposalThreshold(&_Contracts.CallOpts, key)
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
