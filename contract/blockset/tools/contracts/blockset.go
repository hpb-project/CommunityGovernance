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

// BlockSetABI is the input ABI used to generate the binding from.
const BlockSetABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"getVoter\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"deleteAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAdmins\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"getProposalThreshold\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"voteProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"},{\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"resetProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"getValue\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newThreshold\",\"type\":\"uint256\"}],\"name\":\"setThreshold\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"},{\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"addProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getThreshold\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AddAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"DelAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"keywords\",\"type\":\"string\"}],\"name\":\"PassedThreshold\",\"type\":\"event\"}]"

// BlockSetFuncSigs maps the 4-byte function signature to its string representation.
var BlockSetFuncSigs = map[string]string{
	"70480275": "addAdmin(address)",
	"c208abdc": "addProposal(string,uint256)",
	"a6f9dae1": "changeOwner(address)",
	"27e1f7df": "deleteAdmin(address)",
	"31ae450b": "getAdmins()",
	"893d20e8": "getOwner()",
	"481a59a2": "getProposalThreshold(string)",
	"e75235b8": "getThreshold()",
	"960384a0": "getValue(string)",
	"14a0069a": "getVoter(string)",
	"9390ea46": "resetProposal(string,uint256)",
	"960bfe04": "setThreshold(uint256)",
	"55fd06f1": "voteProposal(string)",
}

// BlockSetBin is the compiled bytecode used for deploying new contracts.
var BlockSetBin = "0x608060405234801561001057600080fd5b5060008054600160a060020a031916331780825560018055604051600160a060020a039190911691907f342827c97908e5e2f71151c08502a66d44b6f758e3ac2f1de95f02eb95f0a735908290a3611c268061006d6000396000f3fe6080604052600436106100be577c0100000000000000000000000000000000000000000000000000000000600035046314a0069a81146100c357806327e1f7df146101c657806331ae450b146101fb578063481a59a21461021057806355fd06f1146102d55780637048027514610388578063893d20e8146103bb5780639390ea46146103ec578063960384a0146104a1578063960bfe041461051e578063a6f9dae114610548578063c208abdc1461057b578063e75235b814610630575b600080fd5b3480156100cf57600080fd5b50610176600480360360208110156100e657600080fd5b81019060208101813564010000000081111561010157600080fd5b82018360208201111561011357600080fd5b8035906020019184600183028401116401000000008311171561013557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610645945050505050565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156101b257818101518382015260200161019a565b505050509050019250505060405180910390f35b3480156101d257600080fd5b506101f9600480360360208110156101e957600080fd5b5035600160a060020a03166106fd565b005b34801561020757600080fd5b506101766108ed565b34801561021c57600080fd5b506102c36004803603602081101561023357600080fd5b81019060208101813564010000000081111561024e57600080fd5b82018360208201111561026057600080fd5b8035906020019184600183028401116401000000008311171561028257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610950945050505050565b60408051918252519081900360200190f35b3480156102e157600080fd5b506101f9600480360360208110156102f857600080fd5b81019060208101813564010000000081111561031357600080fd5b82018360208201111561032557600080fd5b8035906020019184600183028401116401000000008311171561034757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610a79945050505050565b34801561039457600080fd5b506101f9600480360360208110156103ab57600080fd5b5035600160a060020a0316610f48565b3480156103c757600080fd5b506103d0611054565b60408051600160a060020a039092168252519081900360200190f35b3480156103f857600080fd5b506101f96004803603604081101561040f57600080fd5b81019060208101813564010000000081111561042a57600080fd5b82018360208201111561043c57600080fd5b8035906020019184600183028401116401000000008311171561045e57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250611063915050565b3480156104ad57600080fd5b506102c3600480360360208110156104c457600080fd5b8101906020810181356401000000008111156104df57600080fd5b8201836020820111156104f157600080fd5b8035906020019184600183028401116401000000008311171561051357600080fd5b50909250905061136d565b34801561052a57600080fd5b506101f96004803603602081101561054157600080fd5b50356114c2565b34801561055457600080fd5b506101f96004803603602081101561056b57600080fd5b5035600160a060020a0316611595565b34801561058757600080fd5b506101f96004803603604081101561059e57600080fd5b8101906020810181356401000000008111156105b957600080fd5b8201836020820111156105cb57600080fd5b803590602001918460018302840111640100000000831117156105ed57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550509135925061164d915050565b34801561063c57600080fd5b506102c3611a16565b60606007826040518082805190602001908083835b602083106106795780518252601f19909201916020918201910161065a565b51815160209384036101000a6000190180199092169116179052920194855250604080519485900382018520805480840287018401909252818652935091508301828280156106f157602002820191906000526020600020905b8154600160a060020a031681526001909101906020018083116106d3575b50505050509050919050565b600054600160a060020a0316331461074d576040805160e560020a62461bcd0281526020600482015260136024820152600080516020611bdb833981519152604482015290519081900360640190fd5b600160a060020a03811660009081526006602052604090205460ff16156108ea5760005b6004548110156108ca5781600160a060020a031660048281548110151561079457fe5b600091825260209091200154600160a060020a031614156108c2576004805460001981019081106107c157fe5b60009182526020909120015460048054600160a060020a0390921691839081106107e757fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039290921691909117905560048054600019810190811061082f57fe5b60009182526020808320909101805473ffffffffffffffffffffffffffffffffffffffff19169055600160a060020a03841682526006905260409020805460ff191690556004805490610886906000198301611a1c565b50604051600160a060020a0383169033907fff210541965015a61784a98072bef6701ce8c6087580100ed42c39322432905190600090a36108ca565b600101610771565b5060045460015411156108ea57600454600181905515156108ea57600180555b50565b6060600480548060200260200160405190810160405280929190818152602001828054801561094557602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610927575b505050505090505b90565b6000806002836040518082805190602001908083835b602083106109855780518252601f199092019160209182019101610966565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922060020154929092119150610a109050576040805160e560020a62461bcd02815260206004820152601260248201527f70726f706f736f6c206e6f742065786973740000000000000000000000000000604482015290519081900360640190fd5b6002826040518082805190602001908083835b60208310610a425780518252601f199092019160209182019101610a23565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922060020154949350505050565b3360009081526006602052604090205460ff1680610aa15750600054600160a060020a031633145b1515610af7576040805160e560020a62461bcd02815260206004820152601d60248201527f43616c6c6572206973206e6f74206f776e657220616e642061646d696e000000604482015290519081900360640190fd5b60006002826040518082805190602001908083835b60208310610b2b5780518252601f199092019160209182019101610b0c565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922060020154929092119150610bb69050576040805160e560020a62461bcd02815260206004820152601260248201527f4e6f7420666f756e642070726f706f73616c0000000000000000000000000000604482015290519081900360640190fd5b60005b6007826040518082805190602001908083835b60208310610beb5780518252601f199092019160209182019101610bcc565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220548310159150610cc290505733600160a060020a03166007836040518082805190602001908083835b60208310610c5e5780518252601f199092019160209182019101610c3f565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922080549092508491508110610c9957fe5b600091825260209091200154600160a060020a03161415610cba57506108ea565b600101610bb9565b506007816040518082805190602001908083835b60208310610cf55780518252601f199092019160209182019101610cd6565b51815160001960209485036101000a019081169019919091161790529201948552506040519384900381018420805460018101825560009182529082902001805473ffffffffffffffffffffffffffffffffffffffff1916331790558451600294869450925082918401908083835b60208310610d835780518252601f199092019160209182019101610d64565b51815160209384036101000a60001901801990921691161790529201948552506040519384900381018420600201548551909460079450869350918291908401908083835b60208310610de75780518252601f199092019160209182019101610dc8565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220549290921091506108ea90505760016002826040518082805190602001908083835b60208310610e535780518252601f199092019160209182019101610e34565b51815160209384036101000a600019018019909216911617905292019485525060405193849003810184206003908101805460ff191696151596909617909555855186949350839250908401908083835b60208310610ec35780518252601f199092019160209182019101610ea4565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092206003015460ff16151591506108ea90505760058054600181018083556000929092528251610f43917f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db001906020850190611a40565b505050565b600054600160a060020a03163314610f98576040805160e560020a62461bcd0281526020600482015260136024820152600080516020611bdb833981519152604482015290519081900360640190fd5b600160a060020a03811660009081526006602052604090205460ff1615156108ea57600160a060020a038116600081815260066020526040808220805460ff1916600190811790915560048054918201815583527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b01805473ffffffffffffffffffffffffffffffffffffffff1916841790555133917f99ad91c72e4fcc0a37d0481b982aec5c0c78491d30d39e65338bec15fa9cc4b791a350565b600054600160a060020a031690565b3360009081526006602052604090205460ff168061108b5750600054600160a060020a031633145b15156110e1576040805160e560020a62461bcd02815260206004820152601d60248201527f43616c6c6572206973206e6f74206f776e657220616e642061646d696e000000604482015290519081900360640190fd5b6002826040518082805190602001908083835b602083106111135780518252601f1990920191602091820191016110f4565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092206003015460ff1615915061126b9050576002826040518082805190602001908083835b602083106111805780518252601f199092019160209182019101611161565b51815160209384036101000a600019018019909216911617905292019485525060405193849003810184208651909460039450879350918291908401908083835b602083106111e05780518252601f1990920191602091820191016111c1565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390206000820181600001908054600181600116156101000203166002900461123a929190611abe565b5060018281015490820155600280830154908201556003918201549101805460ff191660ff90921615159190911790555b6002826040518082805190602001908083835b6020831061129d5780518252601f19909201916020918201910161127e565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092209150600090506112d88282611b33565b506000600182018190556002820155600301805460ff1916905560405182516007918491819060208401908083835b602083106113265780518252601f199092019160209182019101611307565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922061135f925090506000611b77565b611369828261164d565b5050565b60006002838360405180838380828437919091019485525050604051928390036020019092206003015460ff1691508190506113d357506003838360405180838380828437919091019485525050604051928390036020019092206003015460ff169150505b1515611429576040805160e560020a62461bcd02815260206004820152600960248201527f6e6f742076616c69640000000000000000000000000000000000000000000000604482015290519081900360640190fd5b6002838360405180838380828437919091019485525050604051928390036020019092206003015460ff1615915061148f9050576002838360405180838380828437808301925050509250505090815260200160405180910390206001015490506114bc565b60038383604051808383808284379190910194855250506040519283900360200190922060010154925050505b92915050565b600054600160a060020a03163314611512576040805160e560020a62461bcd0281526020600482015260136024820152600080516020611bdb833981519152604482015290519081900360640190fd5b6000811161156a576040805160e560020a62461bcd02815260206004820152600e60248201527f7468726573686f6c64203d3d2030000000000000000000000000000000000000604482015290519081900360640190fd5b60045481111561157f57600454600155611585565b60018190555b60015415156108ea576001805550565b600054600160a060020a031633146115e5576040805160e560020a62461bcd0281526020600482015260136024820152600080516020611bdb833981519152604482015290519081900360640190fd5b60008054604051600160a060020a03808516939216917f342827c97908e5e2f71151c08502a66d44b6f758e3ac2f1de95f02eb95f0a73591a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b3360009081526006602052604090205460ff16806116755750600054600160a060020a031633145b15156116cb576040805160e560020a62461bcd02815260206004820152601d60248201527f43616c6c6572206973206e6f74206f776e657220616e642061646d696e000000604482015290519081900360640190fd5b6002826040518082805190602001908083835b602083106116fd5780518252601f1990920191602091820191016116de565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220600201541591506117859050576040805160e560020a62461bcd02815260206004820152601060248201527f4e6f74206e65772050726f706f73616c00000000000000000000000000000000604482015290519081900360640190fd5b61178d611b95565b82815260015460408083019190915260208083018490529051845160079286929182918401908083835b602083106117d65780518252601f1990920191602091820191016117b7565b51815160001960209485036101000a01908116901991909116179052920194855250604080519485900382018520805460018101825560009182529083902001805473ffffffffffffffffffffffffffffffffffffffff1916331790558501518751909460079450889350918291908401908083835b6020831061186b5780518252601f19909201916020918201910161184c565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220549290921091506119669050576001606082015260405183516003918591819060208401908083835b602083106118de5780518252601f1990920191602091820191016118bf565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092206003015460ff16151591506119619050576005805460018101808355600092909252845161195e917f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db001906020870190611a40565b50505b61196e565b600060608201525b806002846040518082805190602001908083835b602083106119a15780518252601f199092019160209182019101611982565b51815160209384036101000a60001901801990921691161790529201948552506040519384900381019093208451805191946119e294508593500190611a40565b5060208201516001820155604082015160028201556060909101516003909101805460ff1916911515919091179055505050565b60015490565b815481835581811115610f4357600083815260209020610f43918101908301611bc0565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611a8157805160ff1916838001178555611aae565b82800160010185558215611aae579182015b82811115611aae578251825591602001919060010190611a93565b50611aba929150611bc0565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611af75780548555611aae565b82800160010185558215611aae57600052602060002091601f016020900482015b82811115611aae578254825591600101919060010190611b18565b50805460018160011615610100020316600290046000825580601f10611b5957506108ea565b601f0160209004906000526020600020908101906108ea9190611bc0565b50805460008255906000526020600020908101906108ea9190611bc0565b6080604051908101604052806060815260200160008152602001600081526020016000151581525090565b61094d91905b80821115611aba5760008155600101611bc656fe43616c6c6572206973206e6f74206f776e657200000000000000000000000000a165627a7a7230582018f2d3e87d90f9762b90f7fc333e6e15fe888e089c5cc6d41b951876752c45070029"

// DeployBlockSet deploys a new Ethereum contract, binding an instance of BlockSet to it.
func DeployBlockSet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BlockSet, error) {
	parsed, err := abi.JSON(strings.NewReader(BlockSetABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BlockSetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BlockSet{BlockSetCaller: BlockSetCaller{contract: contract}, BlockSetTransactor: BlockSetTransactor{contract: contract}, BlockSetFilterer: BlockSetFilterer{contract: contract}}, nil
}

// BlockSet is an auto generated Go binding around an Ethereum contract.
type BlockSet struct {
	BlockSetCaller     // Read-only binding to the contract
	BlockSetTransactor // Write-only binding to the contract
	BlockSetFilterer   // Log filterer for contract events
}

// BlockSetCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlockSetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockSetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlockSetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockSetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlockSetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockSetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlockSetSession struct {
	Contract     *BlockSet         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BlockSetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlockSetCallerSession struct {
	Contract *BlockSetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BlockSetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlockSetTransactorSession struct {
	Contract     *BlockSetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BlockSetRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlockSetRaw struct {
	Contract *BlockSet // Generic contract binding to access the raw methods on
}

// BlockSetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlockSetCallerRaw struct {
	Contract *BlockSetCaller // Generic read-only contract binding to access the raw methods on
}

// BlockSetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlockSetTransactorRaw struct {
	Contract *BlockSetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlockSet creates a new instance of BlockSet, bound to a specific deployed contract.
func NewBlockSet(address common.Address, backend bind.ContractBackend) (*BlockSet, error) {
	contract, err := bindBlockSet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BlockSet{BlockSetCaller: BlockSetCaller{contract: contract}, BlockSetTransactor: BlockSetTransactor{contract: contract}, BlockSetFilterer: BlockSetFilterer{contract: contract}}, nil
}

// NewBlockSetCaller creates a new read-only instance of BlockSet, bound to a specific deployed contract.
func NewBlockSetCaller(address common.Address, caller bind.ContractCaller) (*BlockSetCaller, error) {
	contract, err := bindBlockSet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlockSetCaller{contract: contract}, nil
}

// NewBlockSetTransactor creates a new write-only instance of BlockSet, bound to a specific deployed contract.
func NewBlockSetTransactor(address common.Address, transactor bind.ContractTransactor) (*BlockSetTransactor, error) {
	contract, err := bindBlockSet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlockSetTransactor{contract: contract}, nil
}

// NewBlockSetFilterer creates a new log filterer instance of BlockSet, bound to a specific deployed contract.
func NewBlockSetFilterer(address common.Address, filterer bind.ContractFilterer) (*BlockSetFilterer, error) {
	contract, err := bindBlockSet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlockSetFilterer{contract: contract}, nil
}

// bindBlockSet binds a generic wrapper to an already deployed contract.
func bindBlockSet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BlockSetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlockSet *BlockSetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlockSet.Contract.BlockSetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlockSet *BlockSetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockSet.Contract.BlockSetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlockSet *BlockSetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlockSet.Contract.BlockSetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlockSet *BlockSetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlockSet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlockSet *BlockSetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockSet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlockSet *BlockSetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlockSet.Contract.contract.Transact(opts, method, params...)
}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() view returns(address[])
func (_BlockSet *BlockSetCaller) GetAdmins(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BlockSet.contract.Call(opts, &out, "getAdmins")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() view returns(address[])
func (_BlockSet *BlockSetSession) GetAdmins() ([]common.Address, error) {
	return _BlockSet.Contract.GetAdmins(&_BlockSet.CallOpts)
}

// GetAdmins is a free data retrieval call binding the contract method 0x31ae450b.
//
// Solidity: function getAdmins() view returns(address[])
func (_BlockSet *BlockSetCallerSession) GetAdmins() ([]common.Address, error) {
	return _BlockSet.Contract.GetAdmins(&_BlockSet.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_BlockSet *BlockSetCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BlockSet.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_BlockSet *BlockSetSession) GetOwner() (common.Address, error) {
	return _BlockSet.Contract.GetOwner(&_BlockSet.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_BlockSet *BlockSetCallerSession) GetOwner() (common.Address, error) {
	return _BlockSet.Contract.GetOwner(&_BlockSet.CallOpts)
}

// GetProposalThreshold is a free data retrieval call binding the contract method 0x481a59a2.
//
// Solidity: function getProposalThreshold(string key) view returns(uint256)
func (_BlockSet *BlockSetCaller) GetProposalThreshold(opts *bind.CallOpts, key string) (*big.Int, error) {
	var out []interface{}
	err := _BlockSet.contract.Call(opts, &out, "getProposalThreshold", key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProposalThreshold is a free data retrieval call binding the contract method 0x481a59a2.
//
// Solidity: function getProposalThreshold(string key) view returns(uint256)
func (_BlockSet *BlockSetSession) GetProposalThreshold(key string) (*big.Int, error) {
	return _BlockSet.Contract.GetProposalThreshold(&_BlockSet.CallOpts, key)
}

// GetProposalThreshold is a free data retrieval call binding the contract method 0x481a59a2.
//
// Solidity: function getProposalThreshold(string key) view returns(uint256)
func (_BlockSet *BlockSetCallerSession) GetProposalThreshold(key string) (*big.Int, error) {
	return _BlockSet.Contract.GetProposalThreshold(&_BlockSet.CallOpts, key)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (_BlockSet *BlockSetCaller) GetThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BlockSet.contract.Call(opts, &out, "getThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (_BlockSet *BlockSetSession) GetThreshold() (*big.Int, error) {
	return _BlockSet.Contract.GetThreshold(&_BlockSet.CallOpts)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (_BlockSet *BlockSetCallerSession) GetThreshold() (*big.Int, error) {
	return _BlockSet.Contract.GetThreshold(&_BlockSet.CallOpts)
}

// GetValue is a free data retrieval call binding the contract method 0x960384a0.
//
// Solidity: function getValue(string key) view returns(uint256)
func (_BlockSet *BlockSetCaller) GetValue(opts *bind.CallOpts, key string) (*big.Int, error) {
	var out []interface{}
	err := _BlockSet.contract.Call(opts, &out, "getValue", key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValue is a free data retrieval call binding the contract method 0x960384a0.
//
// Solidity: function getValue(string key) view returns(uint256)
func (_BlockSet *BlockSetSession) GetValue(key string) (*big.Int, error) {
	return _BlockSet.Contract.GetValue(&_BlockSet.CallOpts, key)
}

// GetValue is a free data retrieval call binding the contract method 0x960384a0.
//
// Solidity: function getValue(string key) view returns(uint256)
func (_BlockSet *BlockSetCallerSession) GetValue(key string) (*big.Int, error) {
	return _BlockSet.Contract.GetValue(&_BlockSet.CallOpts, key)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address admin) returns()
func (_BlockSet *BlockSetTransactor) AddAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _BlockSet.contract.Transact(opts, "addAdmin", admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address admin) returns()
func (_BlockSet *BlockSetSession) AddAdmin(admin common.Address) (*types.Transaction, error) {
	return _BlockSet.Contract.AddAdmin(&_BlockSet.TransactOpts, admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address admin) returns()
func (_BlockSet *BlockSetTransactorSession) AddAdmin(admin common.Address) (*types.Transaction, error) {
	return _BlockSet.Contract.AddAdmin(&_BlockSet.TransactOpts, admin)
}

// AddProposal is a paid mutator transaction binding the contract method 0xc208abdc.
//
// Solidity: function addProposal(string key, uint256 number) returns()
func (_BlockSet *BlockSetTransactor) AddProposal(opts *bind.TransactOpts, key string, number *big.Int) (*types.Transaction, error) {
	return _BlockSet.contract.Transact(opts, "addProposal", key, number)
}

// AddProposal is a paid mutator transaction binding the contract method 0xc208abdc.
//
// Solidity: function addProposal(string key, uint256 number) returns()
func (_BlockSet *BlockSetSession) AddProposal(key string, number *big.Int) (*types.Transaction, error) {
	return _BlockSet.Contract.AddProposal(&_BlockSet.TransactOpts, key, number)
}

// AddProposal is a paid mutator transaction binding the contract method 0xc208abdc.
//
// Solidity: function addProposal(string key, uint256 number) returns()
func (_BlockSet *BlockSetTransactorSession) AddProposal(key string, number *big.Int) (*types.Transaction, error) {
	return _BlockSet.Contract.AddProposal(&_BlockSet.TransactOpts, key, number)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_BlockSet *BlockSetTransactor) ChangeOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BlockSet.contract.Transact(opts, "changeOwner", newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_BlockSet *BlockSetSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _BlockSet.Contract.ChangeOwner(&_BlockSet.TransactOpts, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_BlockSet *BlockSetTransactorSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _BlockSet.Contract.ChangeOwner(&_BlockSet.TransactOpts, newOwner)
}

// DeleteAdmin is a paid mutator transaction binding the contract method 0x27e1f7df.
//
// Solidity: function deleteAdmin(address admin) returns()
func (_BlockSet *BlockSetTransactor) DeleteAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _BlockSet.contract.Transact(opts, "deleteAdmin", admin)
}

// DeleteAdmin is a paid mutator transaction binding the contract method 0x27e1f7df.
//
// Solidity: function deleteAdmin(address admin) returns()
func (_BlockSet *BlockSetSession) DeleteAdmin(admin common.Address) (*types.Transaction, error) {
	return _BlockSet.Contract.DeleteAdmin(&_BlockSet.TransactOpts, admin)
}

// DeleteAdmin is a paid mutator transaction binding the contract method 0x27e1f7df.
//
// Solidity: function deleteAdmin(address admin) returns()
func (_BlockSet *BlockSetTransactorSession) DeleteAdmin(admin common.Address) (*types.Transaction, error) {
	return _BlockSet.Contract.DeleteAdmin(&_BlockSet.TransactOpts, admin)
}

// GetVoter is a paid mutator transaction binding the contract method 0x14a0069a.
//
// Solidity: function getVoter(string key) returns(address[])
func (_BlockSet *BlockSetTransactor) GetVoter(opts *bind.TransactOpts, key string) (*types.Transaction, error) {
	return _BlockSet.contract.Transact(opts, "getVoter", key)
}

// GetVoter is a paid mutator transaction binding the contract method 0x14a0069a.
//
// Solidity: function getVoter(string key) returns(address[])
func (_BlockSet *BlockSetSession) GetVoter(key string) (*types.Transaction, error) {
	return _BlockSet.Contract.GetVoter(&_BlockSet.TransactOpts, key)
}

// GetVoter is a paid mutator transaction binding the contract method 0x14a0069a.
//
// Solidity: function getVoter(string key) returns(address[])
func (_BlockSet *BlockSetTransactorSession) GetVoter(key string) (*types.Transaction, error) {
	return _BlockSet.Contract.GetVoter(&_BlockSet.TransactOpts, key)
}

// ResetProposal is a paid mutator transaction binding the contract method 0x9390ea46.
//
// Solidity: function resetProposal(string key, uint256 number) returns()
func (_BlockSet *BlockSetTransactor) ResetProposal(opts *bind.TransactOpts, key string, number *big.Int) (*types.Transaction, error) {
	return _BlockSet.contract.Transact(opts, "resetProposal", key, number)
}

// ResetProposal is a paid mutator transaction binding the contract method 0x9390ea46.
//
// Solidity: function resetProposal(string key, uint256 number) returns()
func (_BlockSet *BlockSetSession) ResetProposal(key string, number *big.Int) (*types.Transaction, error) {
	return _BlockSet.Contract.ResetProposal(&_BlockSet.TransactOpts, key, number)
}

// ResetProposal is a paid mutator transaction binding the contract method 0x9390ea46.
//
// Solidity: function resetProposal(string key, uint256 number) returns()
func (_BlockSet *BlockSetTransactorSession) ResetProposal(key string, number *big.Int) (*types.Transaction, error) {
	return _BlockSet.Contract.ResetProposal(&_BlockSet.TransactOpts, key, number)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x960bfe04.
//
// Solidity: function setThreshold(uint256 newThreshold) returns()
func (_BlockSet *BlockSetTransactor) SetThreshold(opts *bind.TransactOpts, newThreshold *big.Int) (*types.Transaction, error) {
	return _BlockSet.contract.Transact(opts, "setThreshold", newThreshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x960bfe04.
//
// Solidity: function setThreshold(uint256 newThreshold) returns()
func (_BlockSet *BlockSetSession) SetThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _BlockSet.Contract.SetThreshold(&_BlockSet.TransactOpts, newThreshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x960bfe04.
//
// Solidity: function setThreshold(uint256 newThreshold) returns()
func (_BlockSet *BlockSetTransactorSession) SetThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _BlockSet.Contract.SetThreshold(&_BlockSet.TransactOpts, newThreshold)
}

// VoteProposal is a paid mutator transaction binding the contract method 0x55fd06f1.
//
// Solidity: function voteProposal(string key) returns()
func (_BlockSet *BlockSetTransactor) VoteProposal(opts *bind.TransactOpts, key string) (*types.Transaction, error) {
	return _BlockSet.contract.Transact(opts, "voteProposal", key)
}

// VoteProposal is a paid mutator transaction binding the contract method 0x55fd06f1.
//
// Solidity: function voteProposal(string key) returns()
func (_BlockSet *BlockSetSession) VoteProposal(key string) (*types.Transaction, error) {
	return _BlockSet.Contract.VoteProposal(&_BlockSet.TransactOpts, key)
}

// VoteProposal is a paid mutator transaction binding the contract method 0x55fd06f1.
//
// Solidity: function voteProposal(string key) returns()
func (_BlockSet *BlockSetTransactorSession) VoteProposal(key string) (*types.Transaction, error) {
	return _BlockSet.Contract.VoteProposal(&_BlockSet.TransactOpts, key)
}

// BlockSetAddAdminIterator is returned from FilterAddAdmin and is used to iterate over the raw logs and unpacked data for AddAdmin events raised by the BlockSet contract.
type BlockSetAddAdminIterator struct {
	Event *BlockSetAddAdmin // Event containing the contract specifics and raw log

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
func (it *BlockSetAddAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockSetAddAdmin)
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
		it.Event = new(BlockSetAddAdmin)
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
func (it *BlockSetAddAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockSetAddAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockSetAddAdmin represents a AddAdmin event raised by the BlockSet contract.
type BlockSetAddAdmin struct {
	User     common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAddAdmin is a free log retrieval operation binding the contract event 0x99ad91c72e4fcc0a37d0481b982aec5c0c78491d30d39e65338bec15fa9cc4b7.
//
// Solidity: event AddAdmin(address indexed user, address indexed newAdmin)
func (_BlockSet *BlockSetFilterer) FilterAddAdmin(opts *bind.FilterOpts, user []common.Address, newAdmin []common.Address) (*BlockSetAddAdminIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _BlockSet.contract.FilterLogs(opts, "AddAdmin", userRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return &BlockSetAddAdminIterator{contract: _BlockSet.contract, event: "AddAdmin", logs: logs, sub: sub}, nil
}

// WatchAddAdmin is a free log subscription operation binding the contract event 0x99ad91c72e4fcc0a37d0481b982aec5c0c78491d30d39e65338bec15fa9cc4b7.
//
// Solidity: event AddAdmin(address indexed user, address indexed newAdmin)
func (_BlockSet *BlockSetFilterer) WatchAddAdmin(opts *bind.WatchOpts, sink chan<- *BlockSetAddAdmin, user []common.Address, newAdmin []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _BlockSet.contract.WatchLogs(opts, "AddAdmin", userRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockSetAddAdmin)
				if err := _BlockSet.contract.UnpackLog(event, "AddAdmin", log); err != nil {
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
func (_BlockSet *BlockSetFilterer) ParseAddAdmin(log types.Log) (*BlockSetAddAdmin, error) {
	event := new(BlockSetAddAdmin)
	if err := _BlockSet.contract.UnpackLog(event, "AddAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockSetDelAdminIterator is returned from FilterDelAdmin and is used to iterate over the raw logs and unpacked data for DelAdmin events raised by the BlockSet contract.
type BlockSetDelAdminIterator struct {
	Event *BlockSetDelAdmin // Event containing the contract specifics and raw log

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
func (it *BlockSetDelAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockSetDelAdmin)
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
		it.Event = new(BlockSetDelAdmin)
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
func (it *BlockSetDelAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockSetDelAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockSetDelAdmin represents a DelAdmin event raised by the BlockSet contract.
type BlockSetDelAdmin struct {
	User  common.Address
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDelAdmin is a free log retrieval operation binding the contract event 0xff210541965015a61784a98072bef6701ce8c6087580100ed42c393224329051.
//
// Solidity: event DelAdmin(address indexed user, address indexed admin)
func (_BlockSet *BlockSetFilterer) FilterDelAdmin(opts *bind.FilterOpts, user []common.Address, admin []common.Address) (*BlockSetDelAdminIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _BlockSet.contract.FilterLogs(opts, "DelAdmin", userRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &BlockSetDelAdminIterator{contract: _BlockSet.contract, event: "DelAdmin", logs: logs, sub: sub}, nil
}

// WatchDelAdmin is a free log subscription operation binding the contract event 0xff210541965015a61784a98072bef6701ce8c6087580100ed42c393224329051.
//
// Solidity: event DelAdmin(address indexed user, address indexed admin)
func (_BlockSet *BlockSetFilterer) WatchDelAdmin(opts *bind.WatchOpts, sink chan<- *BlockSetDelAdmin, user []common.Address, admin []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _BlockSet.contract.WatchLogs(opts, "DelAdmin", userRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockSetDelAdmin)
				if err := _BlockSet.contract.UnpackLog(event, "DelAdmin", log); err != nil {
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
func (_BlockSet *BlockSetFilterer) ParseDelAdmin(log types.Log) (*BlockSetDelAdmin, error) {
	event := new(BlockSetDelAdmin)
	if err := _BlockSet.contract.UnpackLog(event, "DelAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockSetOwnerSetIterator is returned from FilterOwnerSet and is used to iterate over the raw logs and unpacked data for OwnerSet events raised by the BlockSet contract.
type BlockSetOwnerSetIterator struct {
	Event *BlockSetOwnerSet // Event containing the contract specifics and raw log

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
func (it *BlockSetOwnerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockSetOwnerSet)
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
		it.Event = new(BlockSetOwnerSet)
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
func (it *BlockSetOwnerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockSetOwnerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockSetOwnerSet represents a OwnerSet event raised by the BlockSet contract.
type BlockSetOwnerSet struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerSet is a free log retrieval operation binding the contract event 0x342827c97908e5e2f71151c08502a66d44b6f758e3ac2f1de95f02eb95f0a735.
//
// Solidity: event OwnerSet(address indexed oldOwner, address indexed newOwner)
func (_BlockSet *BlockSetFilterer) FilterOwnerSet(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*BlockSetOwnerSetIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BlockSet.contract.FilterLogs(opts, "OwnerSet", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BlockSetOwnerSetIterator{contract: _BlockSet.contract, event: "OwnerSet", logs: logs, sub: sub}, nil
}

// WatchOwnerSet is a free log subscription operation binding the contract event 0x342827c97908e5e2f71151c08502a66d44b6f758e3ac2f1de95f02eb95f0a735.
//
// Solidity: event OwnerSet(address indexed oldOwner, address indexed newOwner)
func (_BlockSet *BlockSetFilterer) WatchOwnerSet(opts *bind.WatchOpts, sink chan<- *BlockSetOwnerSet, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BlockSet.contract.WatchLogs(opts, "OwnerSet", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockSetOwnerSet)
				if err := _BlockSet.contract.UnpackLog(event, "OwnerSet", log); err != nil {
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
func (_BlockSet *BlockSetFilterer) ParseOwnerSet(log types.Log) (*BlockSetOwnerSet, error) {
	event := new(BlockSetOwnerSet)
	if err := _BlockSet.contract.UnpackLog(event, "OwnerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockSetPassedThresholdIterator is returned from FilterPassedThreshold and is used to iterate over the raw logs and unpacked data for PassedThreshold events raised by the BlockSet contract.
type BlockSetPassedThresholdIterator struct {
	Event *BlockSetPassedThreshold // Event containing the contract specifics and raw log

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
func (it *BlockSetPassedThresholdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockSetPassedThreshold)
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
		it.Event = new(BlockSetPassedThreshold)
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
func (it *BlockSetPassedThresholdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockSetPassedThresholdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockSetPassedThreshold represents a PassedThreshold event raised by the BlockSet contract.
type BlockSetPassedThreshold struct {
	Keywords common.Hash
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPassedThreshold is a free log retrieval operation binding the contract event 0xcd5e25901a3200193fd4e05510166dd567108536fd9d3b5309934183539298e1.
//
// Solidity: event PassedThreshold(string indexed keywords)
func (_BlockSet *BlockSetFilterer) FilterPassedThreshold(opts *bind.FilterOpts, keywords []string) (*BlockSetPassedThresholdIterator, error) {

	var keywordsRule []interface{}
	for _, keywordsItem := range keywords {
		keywordsRule = append(keywordsRule, keywordsItem)
	}

	logs, sub, err := _BlockSet.contract.FilterLogs(opts, "PassedThreshold", keywordsRule)
	if err != nil {
		return nil, err
	}
	return &BlockSetPassedThresholdIterator{contract: _BlockSet.contract, event: "PassedThreshold", logs: logs, sub: sub}, nil
}

// WatchPassedThreshold is a free log subscription operation binding the contract event 0xcd5e25901a3200193fd4e05510166dd567108536fd9d3b5309934183539298e1.
//
// Solidity: event PassedThreshold(string indexed keywords)
func (_BlockSet *BlockSetFilterer) WatchPassedThreshold(opts *bind.WatchOpts, sink chan<- *BlockSetPassedThreshold, keywords []string) (event.Subscription, error) {

	var keywordsRule []interface{}
	for _, keywordsItem := range keywords {
		keywordsRule = append(keywordsRule, keywordsItem)
	}

	logs, sub, err := _BlockSet.contract.WatchLogs(opts, "PassedThreshold", keywordsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockSetPassedThreshold)
				if err := _BlockSet.contract.UnpackLog(event, "PassedThreshold", log); err != nil {
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
func (_BlockSet *BlockSetFilterer) ParsePassedThreshold(log types.Log) (*BlockSetPassedThreshold, error) {
	event := new(BlockSetPassedThreshold)
	if err := _BlockSet.contract.UnpackLog(event, "PassedThreshold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
