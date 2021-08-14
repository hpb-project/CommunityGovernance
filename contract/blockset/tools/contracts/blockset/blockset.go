// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blockset

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

// BlocksetABI is the input ABI used to generate the binding from.
const BlocksetABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"getVoter\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"deleteAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAdmins\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"getProposalThreshold\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"voteProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"},{\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"resetProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"getValue\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newThreshold\",\"type\":\"uint256\"}],\"name\":\"setThreshold\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"},{\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"addProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getThreshold\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AddAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"DelAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"keywords\",\"type\":\"string\"}],\"name\":\"PassedThreshold\",\"type\":\"event\"}]"

// BlocksetBin is the compiled bytecode used for deploying new contracts.
var BlocksetBin = "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600180819055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f342827c97908e5e2f71151c08502a66d44b6f758e3ac2f1de95f02eb95f0a73560405160405180910390a361274a806100e36000396000f3fe6080604052600436106100bf576000357c01000000000000000000000000000000000000000000000000000000009004806314a0069a146100c457806327e1f7df146101e157806331ae450b14610232578063481a59a21461029e57806355fd06f11461037a5780637048027514610442578063893d20e8146104935780639390ea46146104ea578063960384a0146105bc578063960bfe0414610656578063a6f9dae114610691578063c208abdc146106e2578063e75235b8146107b4575b600080fd5b3480156100d057600080fd5b5061018a600480360360208110156100e757600080fd5b810190808035906020019064010000000081111561010457600080fd5b82018360208201111561011657600080fd5b8035906020019184600183028401116401000000008311171561013857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192905050506107df565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156101cd5780820151818401526020810190506101b2565b505050509050019250505060405180910390f35b3480156101ed57600080fd5b506102306004803603602081101561020457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506108d8565b005b34801561023e57600080fd5b50610247610c50565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561028a57808201518184015260208101905061026f565b505050509050019250505060405180910390f35b3480156102aa57600080fd5b50610364600480360360208110156102c157600080fd5b81019080803590602001906401000000008111156102de57600080fd5b8201836020820111156102f057600080fd5b8035906020019184600183028401116401000000008311171561031257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610cde565b6040518082815260200191505060405180910390f35b34801561038657600080fd5b506104406004803603602081101561039d57600080fd5b81019080803590602001906401000000008111156103ba57600080fd5b8201836020820111156103cc57600080fd5b803590602001918460018302840111640100000000831117156103ee57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610e3b565b005b34801561044e57600080fd5b506104916004803603602081101561046557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506114a1565b005b34801561049f57600080fd5b506104a86116d4565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156104f657600080fd5b506105ba6004803603604081101561050d57600080fd5b810190808035906020019064010000000081111561052a57600080fd5b82018360208201111561053c57600080fd5b8035906020019184600183028401116401000000008311171561055e57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803590602001909291905050506116fd565b005b3480156105c857600080fd5b50610640600480360360208110156105df57600080fd5b81019080803590602001906401000000008111156105fc57600080fd5b82018360208201111561060e57600080fd5b8035906020019184600183028401116401000000008311171561063057600080fd5b9091929391929390505050611b00565b6040518082815260200191505060405180910390f35b34801561066257600080fd5b5061068f6004803603602081101561067957600080fd5b8101908080359060200190929190505050611c90565b005b34801561069d57600080fd5b506106e0600480360360208110156106b457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611e08565b005b3480156106ee57600080fd5b506107b26004803603604081101561070557600080fd5b810190808035906020019064010000000081111561072257600080fd5b82018360208201111561073457600080fd5b8035906020019184600183028401116401000000008311171561075657600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929080359060200190929190505050611f8a565b005b3480156107c057600080fd5b506107c96124a8565b6040518082815260200191505060405180910390f35b60606007826040518082805190602001908083835b60208310151561081957805182526020820191506020810190506020830392506107f4565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390208054806020026020016040519081016040528092919081815260200182805480156108cc57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610882575b50505050509050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561099c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f43616c6c6572206973206e6f74206f776e65720000000000000000000000000081525060200191505060405180910390fd5b600660008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615610c4d5760008090505b600480549050811015610c1c578173ffffffffffffffffffffffffffffffffffffffff16600482815481101515610a2757fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415610c0f576004600160048054905003815481101515610a8557fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600482815481101515610abf57fe5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506004600160048054905003815481101515610b1e57fe5b9060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81549060ff02191690556004805480919060019003610baf91906124b2565b508173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fff210541965015a61784a98072bef6701ce8c6087580100ed42c39322432905160405160405180910390a3610c1c565b80806001019150506109f4565b506004805490506001541115610c4c5760048054905060018190555060006001541415610c4b57600180819055505b5b5b50565b60606004805480602002602001604051908101604052809291908181526020018280548015610cd457602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610c8a575b5050505050905090565b6000806002836040518082805190602001908083835b602083101515610d195780518252602082019150602081019050602083039250610cf4565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060020154111515610dc5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f70726f706f736f6c206e6f74206578697374000000000000000000000000000081525060200191505060405180910390fd5b6002826040518082805190602001908083835b602083101515610dfd5780518252602082019150602081019050602083039250610dd8565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020600201549050919050565b600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1680610edf57506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b1515610f53576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f43616c6c6572206973206e6f74206f776e657220616e642061646d696e00000081525060200191505060405180910390fd5b60006002826040518082805190602001908083835b602083101515610f8d5780518252602082019150602081019050602083039250610f68565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060020154111515611039576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f4e6f7420666f756e642070726f706f73616c000000000000000000000000000081525060200191505060405180910390fd5b60008090505b6007826040518082805190602001908083835b6020831015156110775780518252602082019150602081019050602083039250611052565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390208054905081101561119e573373ffffffffffffffffffffffffffffffffffffffff166007836040518082805190602001908083835b60208310151561110457805182526020820191506020810190506020830392506110df565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390208281548110151561114457fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415611191575061149e565b808060010191505061103f565b506007816040518082805190602001908083835b6020831015156111d757805182526020820191506020810190506020830392506111b2565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390203390806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506002816040518082805190602001908083835b6020831015156112a65780518252602082019150602081019050602083039250611281565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020600201546007826040518082805190602001908083835b60208310151561131557805182526020820191506020810190506020830392506112f0565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390208054905010151561149d5760016002826040518082805190602001908083835b60208310151561138d5780518252602082019150602081019050602083039250611368565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060030160006101000a81548160ff0219169083151502179055506003816040518082805190602001908083835b60208310151561141257805182526020820191506020810190506020830392506113ed565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060030160009054906101000a900460ff16151561149c5760058190806001815401808255809150509060018203906000526020600020016000909192909190915090805190602001906114999291906124de565b50505b5b5b50565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611565576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f43616c6c6572206973206e6f74206f776e65720000000000000000000000000081525060200191505060405180910390fd5b600660008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615156116d1576001600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555060048190806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550508073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f99ad91c72e4fcc0a37d0481b982aec5c0c78491d30d39e65338bec15fa9cc4b760405160405180910390a35b50565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16806117a157506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b1515611815576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f43616c6c6572206973206e6f74206f776e657220616e642061646d696e00000081525060200191505060405180910390fd5b6002826040518082805190602001908083835b60208310151561184d5780518252602082019150602081019050602083039250611828565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060030160009054906101000a900460ff16156119d9576002826040518082805190602001908083835b6020831015156118ce57805182526020820191506020810190506020830392506118a9565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390206003836040518082805190602001908083835b6020831015156119395780518252602082019150602081019050602083039250611914565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390206000820181600001908054600181600116156101000203166002900461199392919061255e565b5060018201548160010155600282015481600201556003820160009054906101000a900460ff168160030160006101000a81548160ff0219169083151502179055509050505b6002826040518082805190602001908083835b602083101515611a1157805182526020820191506020810190506020830392506119ec565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060008082016000611a5591906125e5565b600182016000905560028201600090556003820160006101000a81549060ff021916905550506007826040518082805190602001908083835b602083101515611ab35780518252602082019150602081019050602083039250611a8e565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390206000611af2919061262d565b611afc8282611f8a565b5050565b600060028383604051808383808284378083019250505092505050908152602001604051809103902060030160009054906101000a900460ff1680611b79575060038383604051808383808284378083019250505092505050908152602001604051809103902060030160009054906101000a900460ff165b1515611bed576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260098152602001807f6e6f742076616c6964000000000000000000000000000000000000000000000081525060200191505060405180910390fd5b60028383604051808383808284378083019250505092505050908152602001604051809103902060030160009054906101000a900460ff1615611c5c57600283836040518083838082843780830192505050925050509081526020016040518091039020600101549050611c8a565b6003838360405180838380828437808301925050509250505090815260200160405180910390206001015490505b92915050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611d54576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f43616c6c6572206973206e6f74206f776e65720000000000000000000000000081525060200191505060405180910390fd5b600081111515611dcc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f7468726573686f6c64203d3d203000000000000000000000000000000000000081525060200191505060405180910390fd5b600480549050811115611dea57600480549050600181905550611df2565b806001819055505b60006001541415611e0557600180819055505b50565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611ecc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f43616c6c6572206973206e6f74206f776e65720000000000000000000000000081525060200191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f342827c97908e5e2f71151c08502a66d44b6f758e3ac2f1de95f02eb95f0a73560405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff168061202e57506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b15156120a2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f43616c6c6572206973206e6f74206f776e657220616e642061646d696e00000081525060200191505060405180910390fd5b60006002836040518082805190602001908083835b6020831015156120dc57805182526020820191506020810190506020830392506120b7565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060020154141515612188576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f4e6f74206e65772050726f706f73616c0000000000000000000000000000000081525060200191505060405180910390fd5b61219061264e565b828160000181905250600154816040018181525050818160200181815250506007836040518082805190602001908083835b6020831015156121e757805182526020820191506020810190506020830392506121c2565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390203390806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505080604001516007846040518082805190602001908083835b6020831015156122bb5780518252602082019150602081019050602083039250612296565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020805490501015156123d15760018160600190151590811515815250506003836040518082805190602001908083835b602083101515612342578051825260208201915060208101905060208303925061231d565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060030160009054906101000a900460ff1615156123cc5760058390806001815401808255809150509060018203906000526020600020016000909192909190915090805190602001906123c99291906124de565b50505b6123e3565b60008160600190151590811515815250505b806002846040518082805190602001908083835b60208310151561241c57805182526020820191506020810190506020830392506123f7565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020600082015181600001908051906020019061246b929190612679565b50602082015181600101556040820151816002015560608201518160030160006101000a81548160ff021916908315150217905550905050505050565b6000600154905090565b8154818355818111156124d9578183600052602060002091820191016124d891906126f9565b5b505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061251f57805160ff191683800117855561254d565b8280016001018555821561254d579182015b8281111561254c578251825591602001919060010190612531565b5b50905061255a91906126f9565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061259757805485556125d4565b828001600101855582156125d457600052602060002091601f016020900482015b828111156125d35782548255916001019190600101906125b8565b5b5090506125e191906126f9565b5090565b50805460018160011615610100020316600290046000825580601f1061260b575061262a565b601f01602090049060005260206000209081019061262991906126f9565b5b50565b508054600082559060005260206000209081019061264b91906126f9565b50565b6080604051908101604052806060815260200160008152602001600081526020016000151581525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106126ba57805160ff19168380011785556126e8565b828001600101855582156126e8579182015b828111156126e75782518255916020019190600101906126cc565b5b5090506126f591906126f9565b5090565b61271b91905b808211156127175760008160009055506001016126ff565b5090565b9056fea165627a7a72305820dd56127b18f9e76d2e47b4f7a48cd721054ed61ec0a159e40fb3a21aa78b625d0029"

// DeployBlockset deploys a new Ethereum contract, binding an instance of Blockset to it.
func DeployBlockset(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Blockset, error) {
	parsed, err := abi.JSON(strings.NewReader(BlocksetABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BlocksetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Blockset{BlocksetCaller: BlocksetCaller{contract: contract}, BlocksetTransactor: BlocksetTransactor{contract: contract}, BlocksetFilterer: BlocksetFilterer{contract: contract}}, nil
}

// Blockset is an auto generated Go binding around an Ethereum contract.
type Blockset struct {
	BlocksetCaller     // Read-only binding to the contract
	BlocksetTransactor // Write-only binding to the contract
	BlocksetFilterer   // Log filterer for contract events
}

// BlocksetCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlocksetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlocksetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlocksetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlocksetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlocksetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlocksetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlocksetSession struct {
	Contract     *Blockset         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BlocksetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlocksetCallerSession struct {
	Contract *BlocksetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BlocksetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlocksetTransactorSession struct {
	Contract     *BlocksetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BlocksetRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlocksetRaw struct {
	Contract *Blockset // Generic contract binding to access the raw methods on
}

// BlocksetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlocksetCallerRaw struct {
	Contract *BlocksetCaller // Generic read-only contract binding to access the raw methods on
}

// BlocksetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlocksetTransactorRaw struct {
	Contract *BlocksetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlockset creates a new instance of Blockset, bound to a specific deployed contract.
func NewBlockset(address common.Address, backend bind.ContractBackend) (*Blockset, error) {
	contract, err := bindBlockset(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Blockset{BlocksetCaller: BlocksetCaller{contract: contract}, BlocksetTransactor: BlocksetTransactor{contract: contract}, BlocksetFilterer: BlocksetFilterer{contract: contract}}, nil
}

// NewBlocksetCaller creates a new read-only instance of Blockset, bound to a specific deployed contract.
func NewBlocksetCaller(address common.Address, caller bind.ContractCaller) (*BlocksetCaller, error) {
	contract, err := bindBlockset(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlocksetCaller{contract: contract}, nil
}

// NewBlocksetTransactor creates a new write-only instance of Blockset, bound to a specific deployed contract.
func NewBlocksetTransactor(address common.Address, transactor bind.ContractTransactor) (*BlocksetTransactor, error) {
	contract, err := bindBlockset(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlocksetTransactor{contract: contract}, nil
}

// NewBlocksetFilterer creates a new log filterer instance of Blockset, bound to a specific deployed contract.
func NewBlocksetFilterer(address common.Address, filterer bind.ContractFilterer) (*BlocksetFilterer, error) {
	contract, err := bindBlockset(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlocksetFilterer{contract: contract}, nil
}

// bindBlockset binds a generic wrapper to an already deployed contract.
func bindBlockset(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BlocksetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blockset *BlocksetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Blockset.Contract.BlocksetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blockset *BlocksetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blockset.Contract.BlocksetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blockset *BlocksetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Blockset.Contract.BlocksetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blockset *BlocksetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Blockset.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blockset *BlocksetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blockset.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blockset *BlocksetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Blockset.Contract.contract.Transact(opts, method, params...)
}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() view returns(address[])
func (_Blockset *BlocksetCaller) GetAdmins(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Blockset.contract.Call(opts, &out, "getAdmins")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() view returns(address[])
func (_Blockset *BlocksetSession) GetAdmins() ([]common.Address, error) {
	return _Blockset.Contract.GetAdmins(&_Blockset.CallOpts)
}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() view returns(address[])
func (_Blockset *BlocksetCallerSession) GetAdmins() ([]common.Address, error) {
	return _Blockset.Contract.GetAdmins(&_Blockset.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Blockset *BlocksetCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Blockset.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Blockset *BlocksetSession) GetOwner() (common.Address, error) {
	return _Blockset.Contract.GetOwner(&_Blockset.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Blockset *BlocksetCallerSession) GetOwner() (common.Address, error) {
	return _Blockset.Contract.GetOwner(&_Blockset.CallOpts)
}

// GetProposalThreshold is a free data retrieval call binding the contract method 0x481a59a2.
//
// Solidity: function getProposalThreshold(string key) view returns(uint256)
func (_Blockset *BlocksetCaller) GetProposalThreshold(opts *bind.CallOpts, key string) (*big.Int, error) {
	var out []interface{}
	err := _Blockset.contract.Call(opts, &out, "getProposalThreshold", key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProposalThreshold is a free data retrieval call binding the contract method 0x481a59a2.
//
// Solidity: function getProposalThreshold(string key) view returns(uint256)
func (_Blockset *BlocksetSession) GetProposalThreshold(key string) (*big.Int, error) {
	return _Blockset.Contract.GetProposalThreshold(&_Blockset.CallOpts, key)
}

// GetProposalThreshold is a free data retrieval call binding the contract method 0x481a59a2.
//
// Solidity: function getProposalThreshold(string key) view returns(uint256)
func (_Blockset *BlocksetCallerSession) GetProposalThreshold(key string) (*big.Int, error) {
	return _Blockset.Contract.GetProposalThreshold(&_Blockset.CallOpts, key)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (_Blockset *BlocksetCaller) GetThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Blockset.contract.Call(opts, &out, "getThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (_Blockset *BlocksetSession) GetThreshold() (*big.Int, error) {
	return _Blockset.Contract.GetThreshold(&_Blockset.CallOpts)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (_Blockset *BlocksetCallerSession) GetThreshold() (*big.Int, error) {
	return _Blockset.Contract.GetThreshold(&_Blockset.CallOpts)
}

// GetValue is a free data retrieval call binding the contract method 0x960384a0.
//
// Solidity: function getValue(string key) view returns(uint256)
func (_Blockset *BlocksetCaller) GetValue(opts *bind.CallOpts, key string) (*big.Int, error) {
	var out []interface{}
	err := _Blockset.contract.Call(opts, &out, "getValue", key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValue is a free data retrieval call binding the contract method 0x960384a0.
//
// Solidity: function getValue(string key) view returns(uint256)
func (_Blockset *BlocksetSession) GetValue(key string) (*big.Int, error) {
	return _Blockset.Contract.GetValue(&_Blockset.CallOpts, key)
}

// GetValue is a free data retrieval call binding the contract method 0x960384a0.
//
// Solidity: function getValue(string key) view returns(uint256)
func (_Blockset *BlocksetCallerSession) GetValue(key string) (*big.Int, error) {
	return _Blockset.Contract.GetValue(&_Blockset.CallOpts, key)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address admin) returns()
func (_Blockset *BlocksetTransactor) AddAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _Blockset.contract.Transact(opts, "addAdmin", admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address admin) returns()
func (_Blockset *BlocksetSession) AddAdmin(admin common.Address) (*types.Transaction, error) {
	return _Blockset.Contract.AddAdmin(&_Blockset.TransactOpts, admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address admin) returns()
func (_Blockset *BlocksetTransactorSession) AddAdmin(admin common.Address) (*types.Transaction, error) {
	return _Blockset.Contract.AddAdmin(&_Blockset.TransactOpts, admin)
}

// AddProposal is a paid mutator transaction binding the contract method 0xc208abdc.
//
// Solidity: function addProposal(string key, uint256 number) returns()
func (_Blockset *BlocksetTransactor) AddProposal(opts *bind.TransactOpts, key string, number *big.Int) (*types.Transaction, error) {
	return _Blockset.contract.Transact(opts, "addProposal", key, number)
}

// AddProposal is a paid mutator transaction binding the contract method 0xc208abdc.
//
// Solidity: function addProposal(string key, uint256 number) returns()
func (_Blockset *BlocksetSession) AddProposal(key string, number *big.Int) (*types.Transaction, error) {
	return _Blockset.Contract.AddProposal(&_Blockset.TransactOpts, key, number)
}

// AddProposal is a paid mutator transaction binding the contract method 0xc208abdc.
//
// Solidity: function addProposal(string key, uint256 number) returns()
func (_Blockset *BlocksetTransactorSession) AddProposal(key string, number *big.Int) (*types.Transaction, error) {
	return _Blockset.Contract.AddProposal(&_Blockset.TransactOpts, key, number)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_Blockset *BlocksetTransactor) ChangeOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Blockset.contract.Transact(opts, "changeOwner", newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_Blockset *BlocksetSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Blockset.Contract.ChangeOwner(&_Blockset.TransactOpts, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_Blockset *BlocksetTransactorSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Blockset.Contract.ChangeOwner(&_Blockset.TransactOpts, newOwner)
}

// DeleteAdmin is a paid mutator transaction binding the contract method 0x27e1f7df.
//
// Solidity: function deleteAdmin(address admin) returns()
func (_Blockset *BlocksetTransactor) DeleteAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _Blockset.contract.Transact(opts, "deleteAdmin", admin)
}

// DeleteAdmin is a paid mutator transaction binding the contract method 0x27e1f7df.
//
// Solidity: function deleteAdmin(address admin) returns()
func (_Blockset *BlocksetSession) DeleteAdmin(admin common.Address) (*types.Transaction, error) {
	return _Blockset.Contract.DeleteAdmin(&_Blockset.TransactOpts, admin)
}

// DeleteAdmin is a paid mutator transaction binding the contract method 0x27e1f7df.
//
// Solidity: function deleteAdmin(address admin) returns()
func (_Blockset *BlocksetTransactorSession) DeleteAdmin(admin common.Address) (*types.Transaction, error) {
	return _Blockset.Contract.DeleteAdmin(&_Blockset.TransactOpts, admin)
}

// GetVoter is a paid mutator transaction binding the contract method 0x14a0069a.
//
// Solidity: function getVoter(string key) returns(address[])
func (_Blockset *BlocksetTransactor) GetVoter(opts *bind.TransactOpts, key string) (*types.Transaction, error) {
	return _Blockset.contract.Transact(opts, "getVoter", key)
}

// GetVoter is a paid mutator transaction binding the contract method 0x14a0069a.
//
// Solidity: function getVoter(string key) returns(address[])
func (_Blockset *BlocksetSession) GetVoter(key string) (*types.Transaction, error) {
	return _Blockset.Contract.GetVoter(&_Blockset.TransactOpts, key)
}

// GetVoter is a paid mutator transaction binding the contract method 0x14a0069a.
//
// Solidity: function getVoter(string key) returns(address[])
func (_Blockset *BlocksetTransactorSession) GetVoter(key string) (*types.Transaction, error) {
	return _Blockset.Contract.GetVoter(&_Blockset.TransactOpts, key)
}

// ResetProposal is a paid mutator transaction binding the contract method 0x9390ea46.
//
// Solidity: function resetProposal(string key, uint256 number) returns()
func (_Blockset *BlocksetTransactor) ResetProposal(opts *bind.TransactOpts, key string, number *big.Int) (*types.Transaction, error) {
	return _Blockset.contract.Transact(opts, "resetProposal", key, number)
}

// ResetProposal is a paid mutator transaction binding the contract method 0x9390ea46.
//
// Solidity: function resetProposal(string key, uint256 number) returns()
func (_Blockset *BlocksetSession) ResetProposal(key string, number *big.Int) (*types.Transaction, error) {
	return _Blockset.Contract.ResetProposal(&_Blockset.TransactOpts, key, number)
}

// ResetProposal is a paid mutator transaction binding the contract method 0x9390ea46.
//
// Solidity: function resetProposal(string key, uint256 number) returns()
func (_Blockset *BlocksetTransactorSession) ResetProposal(key string, number *big.Int) (*types.Transaction, error) {
	return _Blockset.Contract.ResetProposal(&_Blockset.TransactOpts, key, number)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x960bfe04.
//
// Solidity: function setThreshold(uint256 newThreshold) returns()
func (_Blockset *BlocksetTransactor) SetThreshold(opts *bind.TransactOpts, newThreshold *big.Int) (*types.Transaction, error) {
	return _Blockset.contract.Transact(opts, "setThreshold", newThreshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x960bfe04.
//
// Solidity: function setThreshold(uint256 newThreshold) returns()
func (_Blockset *BlocksetSession) SetThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _Blockset.Contract.SetThreshold(&_Blockset.TransactOpts, newThreshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x960bfe04.
//
// Solidity: function setThreshold(uint256 newThreshold) returns()
func (_Blockset *BlocksetTransactorSession) SetThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _Blockset.Contract.SetThreshold(&_Blockset.TransactOpts, newThreshold)
}

// VoteProposal is a paid mutator transaction binding the contract method 0x55fd06f1.
//
// Solidity: function voteProposal(string key) returns()
func (_Blockset *BlocksetTransactor) VoteProposal(opts *bind.TransactOpts, key string) (*types.Transaction, error) {
	return _Blockset.contract.Transact(opts, "voteProposal", key)
}

// VoteProposal is a paid mutator transaction binding the contract method 0x55fd06f1.
//
// Solidity: function voteProposal(string key) returns()
func (_Blockset *BlocksetSession) VoteProposal(key string) (*types.Transaction, error) {
	return _Blockset.Contract.VoteProposal(&_Blockset.TransactOpts, key)
}

// VoteProposal is a paid mutator transaction binding the contract method 0x55fd06f1.
//
// Solidity: function voteProposal(string key) returns()
func (_Blockset *BlocksetTransactorSession) VoteProposal(key string) (*types.Transaction, error) {
	return _Blockset.Contract.VoteProposal(&_Blockset.TransactOpts, key)
}

// BlocksetAddAdminIterator is returned from FilterAddAdmin and is used to iterate over the raw logs and unpacked data for AddAdmin events raised by the Blockset contract.
type BlocksetAddAdminIterator struct {
	Event *BlocksetAddAdmin // Event containing the contract specifics and raw log

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
func (it *BlocksetAddAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlocksetAddAdmin)
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
		it.Event = new(BlocksetAddAdmin)
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
func (it *BlocksetAddAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlocksetAddAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlocksetAddAdmin represents a AddAdmin event raised by the Blockset contract.
type BlocksetAddAdmin struct {
	User     common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAddAdmin is a free log retrieval operation binding the contract event 0x99ad91c72e4fcc0a37d0481b982aec5c0c78491d30d39e65338bec15fa9cc4b7.
//
// Solidity: event AddAdmin(address indexed user, address indexed newAdmin)
func (_Blockset *BlocksetFilterer) FilterAddAdmin(opts *bind.FilterOpts, user []common.Address, newAdmin []common.Address) (*BlocksetAddAdminIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _Blockset.contract.FilterLogs(opts, "AddAdmin", userRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return &BlocksetAddAdminIterator{contract: _Blockset.contract, event: "AddAdmin", logs: logs, sub: sub}, nil
}

// WatchAddAdmin is a free log subscription operation binding the contract event 0x99ad91c72e4fcc0a37d0481b982aec5c0c78491d30d39e65338bec15fa9cc4b7.
//
// Solidity: event AddAdmin(address indexed user, address indexed newAdmin)
func (_Blockset *BlocksetFilterer) WatchAddAdmin(opts *bind.WatchOpts, sink chan<- *BlocksetAddAdmin, user []common.Address, newAdmin []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _Blockset.contract.WatchLogs(opts, "AddAdmin", userRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlocksetAddAdmin)
				if err := _Blockset.contract.UnpackLog(event, "AddAdmin", log); err != nil {
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
func (_Blockset *BlocksetFilterer) ParseAddAdmin(log types.Log) (*BlocksetAddAdmin, error) {
	event := new(BlocksetAddAdmin)
	if err := _Blockset.contract.UnpackLog(event, "AddAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlocksetDelAdminIterator is returned from FilterDelAdmin and is used to iterate over the raw logs and unpacked data for DelAdmin events raised by the Blockset contract.
type BlocksetDelAdminIterator struct {
	Event *BlocksetDelAdmin // Event containing the contract specifics and raw log

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
func (it *BlocksetDelAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlocksetDelAdmin)
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
		it.Event = new(BlocksetDelAdmin)
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
func (it *BlocksetDelAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlocksetDelAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlocksetDelAdmin represents a DelAdmin event raised by the Blockset contract.
type BlocksetDelAdmin struct {
	User  common.Address
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDelAdmin is a free log retrieval operation binding the contract event 0xff210541965015a61784a98072bef6701ce8c6087580100ed42c393224329051.
//
// Solidity: event DelAdmin(address indexed user, address indexed admin)
func (_Blockset *BlocksetFilterer) FilterDelAdmin(opts *bind.FilterOpts, user []common.Address, admin []common.Address) (*BlocksetDelAdminIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Blockset.contract.FilterLogs(opts, "DelAdmin", userRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &BlocksetDelAdminIterator{contract: _Blockset.contract, event: "DelAdmin", logs: logs, sub: sub}, nil
}

// WatchDelAdmin is a free log subscription operation binding the contract event 0xff210541965015a61784a98072bef6701ce8c6087580100ed42c393224329051.
//
// Solidity: event DelAdmin(address indexed user, address indexed admin)
func (_Blockset *BlocksetFilterer) WatchDelAdmin(opts *bind.WatchOpts, sink chan<- *BlocksetDelAdmin, user []common.Address, admin []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Blockset.contract.WatchLogs(opts, "DelAdmin", userRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlocksetDelAdmin)
				if err := _Blockset.contract.UnpackLog(event, "DelAdmin", log); err != nil {
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
func (_Blockset *BlocksetFilterer) ParseDelAdmin(log types.Log) (*BlocksetDelAdmin, error) {
	event := new(BlocksetDelAdmin)
	if err := _Blockset.contract.UnpackLog(event, "DelAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlocksetOwnerSetIterator is returned from FilterOwnerSet and is used to iterate over the raw logs and unpacked data for OwnerSet events raised by the Blockset contract.
type BlocksetOwnerSetIterator struct {
	Event *BlocksetOwnerSet // Event containing the contract specifics and raw log

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
func (it *BlocksetOwnerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlocksetOwnerSet)
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
		it.Event = new(BlocksetOwnerSet)
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
func (it *BlocksetOwnerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlocksetOwnerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlocksetOwnerSet represents a OwnerSet event raised by the Blockset contract.
type BlocksetOwnerSet struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerSet is a free log retrieval operation binding the contract event 0x342827c97908e5e2f71151c08502a66d44b6f758e3ac2f1de95f02eb95f0a735.
//
// Solidity: event OwnerSet(address indexed oldOwner, address indexed newOwner)
func (_Blockset *BlocksetFilterer) FilterOwnerSet(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*BlocksetOwnerSetIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Blockset.contract.FilterLogs(opts, "OwnerSet", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BlocksetOwnerSetIterator{contract: _Blockset.contract, event: "OwnerSet", logs: logs, sub: sub}, nil
}

// WatchOwnerSet is a free log subscription operation binding the contract event 0x342827c97908e5e2f71151c08502a66d44b6f758e3ac2f1de95f02eb95f0a735.
//
// Solidity: event OwnerSet(address indexed oldOwner, address indexed newOwner)
func (_Blockset *BlocksetFilterer) WatchOwnerSet(opts *bind.WatchOpts, sink chan<- *BlocksetOwnerSet, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Blockset.contract.WatchLogs(opts, "OwnerSet", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlocksetOwnerSet)
				if err := _Blockset.contract.UnpackLog(event, "OwnerSet", log); err != nil {
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
func (_Blockset *BlocksetFilterer) ParseOwnerSet(log types.Log) (*BlocksetOwnerSet, error) {
	event := new(BlocksetOwnerSet)
	if err := _Blockset.contract.UnpackLog(event, "OwnerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlocksetPassedThresholdIterator is returned from FilterPassedThreshold and is used to iterate over the raw logs and unpacked data for PassedThreshold events raised by the Blockset contract.
type BlocksetPassedThresholdIterator struct {
	Event *BlocksetPassedThreshold // Event containing the contract specifics and raw log

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
func (it *BlocksetPassedThresholdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlocksetPassedThreshold)
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
		it.Event = new(BlocksetPassedThreshold)
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
func (it *BlocksetPassedThresholdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlocksetPassedThresholdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlocksetPassedThreshold represents a PassedThreshold event raised by the Blockset contract.
type BlocksetPassedThreshold struct {
	Keywords common.Hash
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPassedThreshold is a free log retrieval operation binding the contract event 0xcd5e25901a3200193fd4e05510166dd567108536fd9d3b5309934183539298e1.
//
// Solidity: event PassedThreshold(string indexed keywords)
func (_Blockset *BlocksetFilterer) FilterPassedThreshold(opts *bind.FilterOpts, keywords []string) (*BlocksetPassedThresholdIterator, error) {

	var keywordsRule []interface{}
	for _, keywordsItem := range keywords {
		keywordsRule = append(keywordsRule, keywordsItem)
	}

	logs, sub, err := _Blockset.contract.FilterLogs(opts, "PassedThreshold", keywordsRule)
	if err != nil {
		return nil, err
	}
	return &BlocksetPassedThresholdIterator{contract: _Blockset.contract, event: "PassedThreshold", logs: logs, sub: sub}, nil
}

// WatchPassedThreshold is a free log subscription operation binding the contract event 0xcd5e25901a3200193fd4e05510166dd567108536fd9d3b5309934183539298e1.
//
// Solidity: event PassedThreshold(string indexed keywords)
func (_Blockset *BlocksetFilterer) WatchPassedThreshold(opts *bind.WatchOpts, sink chan<- *BlocksetPassedThreshold, keywords []string) (event.Subscription, error) {

	var keywordsRule []interface{}
	for _, keywordsItem := range keywords {
		keywordsRule = append(keywordsRule, keywordsItem)
	}

	logs, sub, err := _Blockset.contract.WatchLogs(opts, "PassedThreshold", keywordsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlocksetPassedThreshold)
				if err := _Blockset.contract.UnpackLog(event, "PassedThreshold", log); err != nil {
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
func (_Blockset *BlocksetFilterer) ParsePassedThreshold(log types.Log) (*BlocksetPassedThreshold, error) {
	event := new(BlocksetPassedThreshold)
	if err := _Blockset.contract.UnpackLog(event, "PassedThreshold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
