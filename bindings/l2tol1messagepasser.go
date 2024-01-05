// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/scroll-tech/go-ethereum"
	"github.com/scroll-tech/go-ethereum/accounts/abi"
	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/event"
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
)

// L2ToL1MessagePasserMetaData contains all meta data concerning the L2ToL1MessagePasser contract.
var L2ToL1MessagePasserMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_l1CrossDomainMessenger\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"MESSAGE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTreeRoot\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initiateWithdrawal\",\"inputs\":[{\"name\":\"_target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"leafNodesCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"messageNonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"messageRoot\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recAddr\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sentMessages\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyMerkleProof\",\"inputs\":[{\"name\":\"leafHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"smtProof\",\"type\":\"bytes32[32]\",\"internalType\":\"bytes32[32]\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"MessagePassed\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"withdrawalHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"rootHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawerBalanceBurnt\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"MerkleTreeFull\",\"inputs\":[]}]",
	Bin: "0x61010060405234801561001157600080fd5b50604051620011c5380380620011c583398101604081905261003291610149565b6001608052600060a081905260c05273111100000000000000000000000000000000111181016001600160a01b031660e05261006c610075565b602d555061018f565b602054600090819081805b6020811015610140578083901c6001166001036100dd57600081602081106100aa576100aa610179565b0154604080516020810192909252810185905260600160405160208183030381529060405280519060200120935061010a565b60408051602081018690529081018390526060016040516020818303038152906040528051906020012093505b604080516020810184905290810183905260600160408051601f1981840301815291905280516020909101209150600101610080565b50919392505050565b60006020828403121561015b57600080fd5b81516001600160a01b038116811461017257600080fd5b9392505050565b634e487b7160e01b600052603260045260246000fd5b60805160a05160c05160e051610ff5620001d0600039600081816102010152610386015260006105a50152600061057c015260006105530152610ff56000f3fe6080604052600436106100b55760003560e01c8063b58343bb11610069578063d4b9f4fa1161004e578063d4b9f4fa146101d9578063e77c99a9146101ef578063ecc704281461024857600080fd5b8063b58343bb146101b0578063c2b3e5ac146101c657600080fd5b806354fd4d501161009a57806354fd4d501461013b57806382e3702d1461015d57806389c09d381461018d57600080fd5b8063340735f7146100de5780633f827a5a1461011357600080fd5b366100d9576100d733620186a06040518060200160405280600081525061025d565b005b600080fd5b3480156100ea57600080fd5b506100fe6100f9366004610ac2565b610482565b60405190151581526020015b60405180910390f35b34801561011f57600080fd5b50610128600181565b60405161ffff909116815260200161010a565b34801561014757600080fd5b5061015061054c565b60405161010a9190610bc9565b34801561016957600080fd5b506100fe610178366004610be3565b602b6020526000908152604090205460ff1681565b34801561019957600080fd5b506101a26105ef565b60405190815260200161010a565b3480156101bc57600080fd5b506101a260205481565b6100d76101d4366004610bfc565b61025d565b3480156101e557600080fd5b506101a2602d5481565b3480156101fb57600080fd5b506102237f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161010a565b34801561025457600080fd5b506101a26106e1565b60006102b16040518060c001604052806102756106e1565b815233602082015273ffffffffffffffffffffffffffffffffffffffff871660408201523460608201526080810186905260a001849052610717565b90506102bc81610764565b6102c46105ef565b602d5573ffffffffffffffffffffffffffffffffffffffff8416336102e76106e1565b7f3034a173cd04c8b47937b2f95d00794df516ce57df60210b71de7ecf6d2fe4f334878787602d54604051610320959493929190610ce4565b60405180910390a4602c80547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808216600101167fffff000000000000000000000000000000000000000000000000000000000000909116179055341561047c5760006103bc7f00000000000000000000000000000000000000000000000000000000000000005a346040518060200160405280600081525061087c565b90508061044f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602860248201527f4c32546f4c314d6573736167655061737365723a20455448207472616e73666560448201527f72206661696c6564000000000000000000000000000000000000000000000000606482015260840160405180910390fd5b60405134907f7967de617a5ac1cc7eba2d6f37570a0135afa950d8bb77cdd35f0d0b4e85a16f90600090a2505b50505050565b600084815b6020811015610540578085901c6001166001036104ed578581602081106104b0576104b0610d17565b6020020151826040516020016104d0929190918252602082015260400190565b604051602081830303815290604052805190602001209150610538565b8186826020811061050057610500610d17565b602002015160405160200161051f929190918252602082015260400190565b6040516020818303038152906040528051906020012091505b600101610487565b50909114949350505050565b60606105777f0000000000000000000000000000000000000000000000000000000000000000610896565b6105a07f0000000000000000000000000000000000000000000000000000000000000000610896565b6105c97f0000000000000000000000000000000000000000000000000000000000000000610896565b6040516020016105db93929190610d46565b604051602081830303815290604052905090565b602054600090819081805b60208110156106d8578083901c600116600103610657576000816020811061062457610624610d17565b01546040805160208101929092528101859052606001604051602081830303815290604052805190602001209350610684565b60408051602081018690529081018390526060016040516020818303038152906040528051906020012093505b6040805160208101849052908101839052606001604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052805160209091012091506001016105fa565b50919392505050565b602c54600090610712907dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff166001610954565b905090565b80516020808301516040808501516060860151608087015160a08801519351600097610747979096959101610dbc565b604051602081830303815290604052805190602001209050919050565b80600161077360206002610f62565b61077d9190610f6e565b602054106107b7576040517fef5ccf6600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006020600081546107c890610f81565b9182905550905060005b602081101561086e578082901c6001166001036108055782600082602081106107fd576107fd610d17565b015550505050565b6000816020811061081857610818610d17565b01546040805160208101929092528101849052606001604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052805160209091012092506001016107d2565b50610877610fb9565b505050565b600080600080845160208601878a8af19695505050505050565b606060006108a383610962565b600101905060008167ffffffffffffffff8111156108c3576108c3610a44565b6040519080825280601f01601f1916602001820160405280156108ed576020820181803683370190505b5090508181016020015b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a85049450846108f757509392505050565b60f081901b82175b92915050565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f01000000000000000083106109ab577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef810000000083106109d7576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc1000083106109f557662386f26fc10000830492506010015b6305f5e1008310610a0d576305f5e100830492506008015b6127108310610a2157612710830492506004015b60648310610a33576064830492506002015b600a831061095c5760010192915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715610aba57610aba610a44565b604052919050565b6000806000806104608587031215610ad957600080fd5b84359350602086603f870112610aee57600080fd5b604051610400810181811067ffffffffffffffff82111715610b1257610b12610a44565b60405280610420880189811115610b2857600080fd5b602089015b81811015610b445780358352918401918401610b2d565b509699919850509435956104400135949350505050565b60005b83811015610b76578181015183820152602001610b5e565b50506000910152565b60008151808452610b97816020860160208601610b5b565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000610bdc6020830184610b7f565b9392505050565b600060208284031215610bf557600080fd5b5035919050565b600080600060608486031215610c1157600080fd5b833573ffffffffffffffffffffffffffffffffffffffff81168114610c3557600080fd5b92506020848101359250604085013567ffffffffffffffff80821115610c5a57600080fd5b818701915087601f830112610c6e57600080fd5b813581811115610c8057610c80610a44565b610cb0847fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601610a73565b91508082528884828501011115610cc657600080fd5b80848401858401376000848284010152508093505050509250925092565b85815284602082015260a060408201526000610d0360a0830186610b7f565b606083019490945250608001529392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60008451610d58818460208901610b5b565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551610d94816001850160208a01610b5b565b60019201918201528351610daf816002840160208801610b5b565b0160020195945050505050565b868152600073ffffffffffffffffffffffffffffffffffffffff808816602084015280871660408401525084606083015283608083015260c060a0830152610e0760c0830184610b7f565b98975050505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600181815b80851115610e9b57817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115610e8157610e81610e13565b80851615610e8e57918102915b93841c9390800290610e47565b509250929050565b600082610eb25750600161095c565b81610ebf5750600061095c565b8160018114610ed55760028114610edf57610efb565b600191505061095c565b60ff841115610ef057610ef0610e13565b50506001821b61095c565b5060208310610133831016604e8410600b8410161715610f1e575081810a61095c565b610f288383610e42565b807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115610f5a57610f5a610e13565b029392505050565b6000610bdc8383610ea3565b8181038181111561095c5761095c610e13565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610fb257610fb2610e13565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fdfea164736f6c6343000817000a",
}

// L2ToL1MessagePasserABI is the input ABI used to generate the binding from.
// Deprecated: Use L2ToL1MessagePasserMetaData.ABI instead.
var L2ToL1MessagePasserABI = L2ToL1MessagePasserMetaData.ABI

// L2ToL1MessagePasserBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L2ToL1MessagePasserMetaData.Bin instead.
var L2ToL1MessagePasserBin = L2ToL1MessagePasserMetaData.Bin

// DeployL2ToL1MessagePasser deploys a new Ethereum contract, binding an instance of L2ToL1MessagePasser to it.
func DeployL2ToL1MessagePasser(auth *bind.TransactOpts, backend bind.ContractBackend, _l1CrossDomainMessenger common.Address) (common.Address, *types.Transaction, *L2ToL1MessagePasser, error) {
	parsed, err := L2ToL1MessagePasserMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L2ToL1MessagePasserBin), backend, _l1CrossDomainMessenger)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L2ToL1MessagePasser{L2ToL1MessagePasserCaller: L2ToL1MessagePasserCaller{contract: contract}, L2ToL1MessagePasserTransactor: L2ToL1MessagePasserTransactor{contract: contract}, L2ToL1MessagePasserFilterer: L2ToL1MessagePasserFilterer{contract: contract}}, nil
}

// L2ToL1MessagePasser is an auto generated Go binding around an Ethereum contract.
type L2ToL1MessagePasser struct {
	L2ToL1MessagePasserCaller     // Read-only binding to the contract
	L2ToL1MessagePasserTransactor // Write-only binding to the contract
	L2ToL1MessagePasserFilterer   // Log filterer for contract events
}

// L2ToL1MessagePasserCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2ToL1MessagePasserCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ToL1MessagePasserTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2ToL1MessagePasserTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ToL1MessagePasserFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2ToL1MessagePasserFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ToL1MessagePasserSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2ToL1MessagePasserSession struct {
	Contract     *L2ToL1MessagePasser // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// L2ToL1MessagePasserCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2ToL1MessagePasserCallerSession struct {
	Contract *L2ToL1MessagePasserCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// L2ToL1MessagePasserTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2ToL1MessagePasserTransactorSession struct {
	Contract     *L2ToL1MessagePasserTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// L2ToL1MessagePasserRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2ToL1MessagePasserRaw struct {
	Contract *L2ToL1MessagePasser // Generic contract binding to access the raw methods on
}

// L2ToL1MessagePasserCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2ToL1MessagePasserCallerRaw struct {
	Contract *L2ToL1MessagePasserCaller // Generic read-only contract binding to access the raw methods on
}

// L2ToL1MessagePasserTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2ToL1MessagePasserTransactorRaw struct {
	Contract *L2ToL1MessagePasserTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2ToL1MessagePasser creates a new instance of L2ToL1MessagePasser, bound to a specific deployed contract.
func NewL2ToL1MessagePasser(address common.Address, backend bind.ContractBackend) (*L2ToL1MessagePasser, error) {
	contract, err := bindL2ToL1MessagePasser(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2ToL1MessagePasser{L2ToL1MessagePasserCaller: L2ToL1MessagePasserCaller{contract: contract}, L2ToL1MessagePasserTransactor: L2ToL1MessagePasserTransactor{contract: contract}, L2ToL1MessagePasserFilterer: L2ToL1MessagePasserFilterer{contract: contract}}, nil
}

// NewL2ToL1MessagePasserCaller creates a new read-only instance of L2ToL1MessagePasser, bound to a specific deployed contract.
func NewL2ToL1MessagePasserCaller(address common.Address, caller bind.ContractCaller) (*L2ToL1MessagePasserCaller, error) {
	contract, err := bindL2ToL1MessagePasser(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2ToL1MessagePasserCaller{contract: contract}, nil
}

// NewL2ToL1MessagePasserTransactor creates a new write-only instance of L2ToL1MessagePasser, bound to a specific deployed contract.
func NewL2ToL1MessagePasserTransactor(address common.Address, transactor bind.ContractTransactor) (*L2ToL1MessagePasserTransactor, error) {
	contract, err := bindL2ToL1MessagePasser(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2ToL1MessagePasserTransactor{contract: contract}, nil
}

// NewL2ToL1MessagePasserFilterer creates a new log filterer instance of L2ToL1MessagePasser, bound to a specific deployed contract.
func NewL2ToL1MessagePasserFilterer(address common.Address, filterer bind.ContractFilterer) (*L2ToL1MessagePasserFilterer, error) {
	contract, err := bindL2ToL1MessagePasser(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2ToL1MessagePasserFilterer{contract: contract}, nil
}

// bindL2ToL1MessagePasser binds a generic wrapper to an already deployed contract.
func bindL2ToL1MessagePasser(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L2ToL1MessagePasserABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2ToL1MessagePasser *L2ToL1MessagePasserRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2ToL1MessagePasser.Contract.L2ToL1MessagePasserCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2ToL1MessagePasser *L2ToL1MessagePasserRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.L2ToL1MessagePasserTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2ToL1MessagePasser *L2ToL1MessagePasserRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.L2ToL1MessagePasserTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2ToL1MessagePasser.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2ToL1MessagePasser *L2ToL1MessagePasserTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2ToL1MessagePasser *L2ToL1MessagePasserTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.contract.Transact(opts, method, params...)
}

// MESSAGEVERSION is a free data retrieval call binding the contract method 0x3f827a5a.
//
// Solidity: function MESSAGE_VERSION() view returns(uint16)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCaller) MESSAGEVERSION(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _L2ToL1MessagePasser.contract.Call(opts, &out, "MESSAGE_VERSION")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MESSAGEVERSION is a free data retrieval call binding the contract method 0x3f827a5a.
//
// Solidity: function MESSAGE_VERSION() view returns(uint16)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserSession) MESSAGEVERSION() (uint16, error) {
	return _L2ToL1MessagePasser.Contract.MESSAGEVERSION(&_L2ToL1MessagePasser.CallOpts)
}

// MESSAGEVERSION is a free data retrieval call binding the contract method 0x3f827a5a.
//
// Solidity: function MESSAGE_VERSION() view returns(uint16)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCallerSession) MESSAGEVERSION() (uint16, error) {
	return _L2ToL1MessagePasser.Contract.MESSAGEVERSION(&_L2ToL1MessagePasser.CallOpts)
}

// GetTreeRoot is a free data retrieval call binding the contract method 0x89c09d38.
//
// Solidity: function getTreeRoot() view returns(bytes32)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCaller) GetTreeRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _L2ToL1MessagePasser.contract.Call(opts, &out, "getTreeRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTreeRoot is a free data retrieval call binding the contract method 0x89c09d38.
//
// Solidity: function getTreeRoot() view returns(bytes32)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserSession) GetTreeRoot() ([32]byte, error) {
	return _L2ToL1MessagePasser.Contract.GetTreeRoot(&_L2ToL1MessagePasser.CallOpts)
}

// GetTreeRoot is a free data retrieval call binding the contract method 0x89c09d38.
//
// Solidity: function getTreeRoot() view returns(bytes32)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCallerSession) GetTreeRoot() ([32]byte, error) {
	return _L2ToL1MessagePasser.Contract.GetTreeRoot(&_L2ToL1MessagePasser.CallOpts)
}

// LeafNodesCount is a free data retrieval call binding the contract method 0xb58343bb.
//
// Solidity: function leafNodesCount() view returns(uint256)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCaller) LeafNodesCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2ToL1MessagePasser.contract.Call(opts, &out, "leafNodesCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LeafNodesCount is a free data retrieval call binding the contract method 0xb58343bb.
//
// Solidity: function leafNodesCount() view returns(uint256)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserSession) LeafNodesCount() (*big.Int, error) {
	return _L2ToL1MessagePasser.Contract.LeafNodesCount(&_L2ToL1MessagePasser.CallOpts)
}

// LeafNodesCount is a free data retrieval call binding the contract method 0xb58343bb.
//
// Solidity: function leafNodesCount() view returns(uint256)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCallerSession) LeafNodesCount() (*big.Int, error) {
	return _L2ToL1MessagePasser.Contract.LeafNodesCount(&_L2ToL1MessagePasser.CallOpts)
}

// MessageNonce is a free data retrieval call binding the contract method 0xecc70428.
//
// Solidity: function messageNonce() view returns(uint256)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCaller) MessageNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2ToL1MessagePasser.contract.Call(opts, &out, "messageNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MessageNonce is a free data retrieval call binding the contract method 0xecc70428.
//
// Solidity: function messageNonce() view returns(uint256)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserSession) MessageNonce() (*big.Int, error) {
	return _L2ToL1MessagePasser.Contract.MessageNonce(&_L2ToL1MessagePasser.CallOpts)
}

// MessageNonce is a free data retrieval call binding the contract method 0xecc70428.
//
// Solidity: function messageNonce() view returns(uint256)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCallerSession) MessageNonce() (*big.Int, error) {
	return _L2ToL1MessagePasser.Contract.MessageNonce(&_L2ToL1MessagePasser.CallOpts)
}

// MessageRoot is a free data retrieval call binding the contract method 0xd4b9f4fa.
//
// Solidity: function messageRoot() view returns(bytes32)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCaller) MessageRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _L2ToL1MessagePasser.contract.Call(opts, &out, "messageRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MessageRoot is a free data retrieval call binding the contract method 0xd4b9f4fa.
//
// Solidity: function messageRoot() view returns(bytes32)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserSession) MessageRoot() ([32]byte, error) {
	return _L2ToL1MessagePasser.Contract.MessageRoot(&_L2ToL1MessagePasser.CallOpts)
}

// MessageRoot is a free data retrieval call binding the contract method 0xd4b9f4fa.
//
// Solidity: function messageRoot() view returns(bytes32)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCallerSession) MessageRoot() ([32]byte, error) {
	return _L2ToL1MessagePasser.Contract.MessageRoot(&_L2ToL1MessagePasser.CallOpts)
}

// RecAddr is a free data retrieval call binding the contract method 0xe77c99a9.
//
// Solidity: function recAddr() view returns(address)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCaller) RecAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2ToL1MessagePasser.contract.Call(opts, &out, "recAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RecAddr is a free data retrieval call binding the contract method 0xe77c99a9.
//
// Solidity: function recAddr() view returns(address)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserSession) RecAddr() (common.Address, error) {
	return _L2ToL1MessagePasser.Contract.RecAddr(&_L2ToL1MessagePasser.CallOpts)
}

// RecAddr is a free data retrieval call binding the contract method 0xe77c99a9.
//
// Solidity: function recAddr() view returns(address)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCallerSession) RecAddr() (common.Address, error) {
	return _L2ToL1MessagePasser.Contract.RecAddr(&_L2ToL1MessagePasser.CallOpts)
}

// SentMessages is a free data retrieval call binding the contract method 0x82e3702d.
//
// Solidity: function sentMessages(bytes32 ) view returns(bool)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCaller) SentMessages(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _L2ToL1MessagePasser.contract.Call(opts, &out, "sentMessages", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SentMessages is a free data retrieval call binding the contract method 0x82e3702d.
//
// Solidity: function sentMessages(bytes32 ) view returns(bool)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserSession) SentMessages(arg0 [32]byte) (bool, error) {
	return _L2ToL1MessagePasser.Contract.SentMessages(&_L2ToL1MessagePasser.CallOpts, arg0)
}

// SentMessages is a free data retrieval call binding the contract method 0x82e3702d.
//
// Solidity: function sentMessages(bytes32 ) view returns(bool)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCallerSession) SentMessages(arg0 [32]byte) (bool, error) {
	return _L2ToL1MessagePasser.Contract.SentMessages(&_L2ToL1MessagePasser.CallOpts, arg0)
}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0x340735f7.
//
// Solidity: function verifyMerkleProof(bytes32 leafHash, bytes32[32] smtProof, uint256 index, bytes32 root) pure returns(bool)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCaller) VerifyMerkleProof(opts *bind.CallOpts, leafHash [32]byte, smtProof [32][32]byte, index *big.Int, root [32]byte) (bool, error) {
	var out []interface{}
	err := _L2ToL1MessagePasser.contract.Call(opts, &out, "verifyMerkleProof", leafHash, smtProof, index, root)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0x340735f7.
//
// Solidity: function verifyMerkleProof(bytes32 leafHash, bytes32[32] smtProof, uint256 index, bytes32 root) pure returns(bool)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserSession) VerifyMerkleProof(leafHash [32]byte, smtProof [32][32]byte, index *big.Int, root [32]byte) (bool, error) {
	return _L2ToL1MessagePasser.Contract.VerifyMerkleProof(&_L2ToL1MessagePasser.CallOpts, leafHash, smtProof, index, root)
}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0x340735f7.
//
// Solidity: function verifyMerkleProof(bytes32 leafHash, bytes32[32] smtProof, uint256 index, bytes32 root) pure returns(bool)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCallerSession) VerifyMerkleProof(leafHash [32]byte, smtProof [32][32]byte, index *big.Int, root [32]byte) (bool, error) {
	return _L2ToL1MessagePasser.Contract.VerifyMerkleProof(&_L2ToL1MessagePasser.CallOpts, leafHash, smtProof, index, root)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _L2ToL1MessagePasser.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserSession) Version() (string, error) {
	return _L2ToL1MessagePasser.Contract.Version(&_L2ToL1MessagePasser.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCallerSession) Version() (string, error) {
	return _L2ToL1MessagePasser.Contract.Version(&_L2ToL1MessagePasser.CallOpts)
}

// InitiateWithdrawal is a paid mutator transaction binding the contract method 0xc2b3e5ac.
//
// Solidity: function initiateWithdrawal(address _target, uint256 _gasLimit, bytes _data) payable returns()
func (_L2ToL1MessagePasser *L2ToL1MessagePasserTransactor) InitiateWithdrawal(opts *bind.TransactOpts, _target common.Address, _gasLimit *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2ToL1MessagePasser.contract.Transact(opts, "initiateWithdrawal", _target, _gasLimit, _data)
}

// InitiateWithdrawal is a paid mutator transaction binding the contract method 0xc2b3e5ac.
//
// Solidity: function initiateWithdrawal(address _target, uint256 _gasLimit, bytes _data) payable returns()
func (_L2ToL1MessagePasser *L2ToL1MessagePasserSession) InitiateWithdrawal(_target common.Address, _gasLimit *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.InitiateWithdrawal(&_L2ToL1MessagePasser.TransactOpts, _target, _gasLimit, _data)
}

// InitiateWithdrawal is a paid mutator transaction binding the contract method 0xc2b3e5ac.
//
// Solidity: function initiateWithdrawal(address _target, uint256 _gasLimit, bytes _data) payable returns()
func (_L2ToL1MessagePasser *L2ToL1MessagePasserTransactorSession) InitiateWithdrawal(_target common.Address, _gasLimit *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.InitiateWithdrawal(&_L2ToL1MessagePasser.TransactOpts, _target, _gasLimit, _data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L2ToL1MessagePasser *L2ToL1MessagePasserTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2ToL1MessagePasser.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L2ToL1MessagePasser *L2ToL1MessagePasserSession) Receive() (*types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.Receive(&_L2ToL1MessagePasser.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L2ToL1MessagePasser *L2ToL1MessagePasserTransactorSession) Receive() (*types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.Receive(&_L2ToL1MessagePasser.TransactOpts)
}

// L2ToL1MessagePasserMessagePassedIterator is returned from FilterMessagePassed and is used to iterate over the raw logs and unpacked data for MessagePassed events raised by the L2ToL1MessagePasser contract.
type L2ToL1MessagePasserMessagePassedIterator struct {
	Event *L2ToL1MessagePasserMessagePassed // Event containing the contract specifics and raw log

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
func (it *L2ToL1MessagePasserMessagePassedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ToL1MessagePasserMessagePassed)
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
		it.Event = new(L2ToL1MessagePasserMessagePassed)
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
func (it *L2ToL1MessagePasserMessagePassedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ToL1MessagePasserMessagePassedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ToL1MessagePasserMessagePassed represents a MessagePassed event raised by the L2ToL1MessagePasser contract.
type L2ToL1MessagePasserMessagePassed struct {
	Nonce          *big.Int
	Sender         common.Address
	Target         common.Address
	Value          *big.Int
	GasLimit       *big.Int
	Data           []byte
	WithdrawalHash [32]byte
	RootHash       [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMessagePassed is a free log retrieval operation binding the contract event 0x3034a173cd04c8b47937b2f95d00794df516ce57df60210b71de7ecf6d2fe4f3.
//
// Solidity: event MessagePassed(uint256 indexed nonce, address indexed sender, address indexed target, uint256 value, uint256 gasLimit, bytes data, bytes32 withdrawalHash, bytes32 rootHash)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserFilterer) FilterMessagePassed(opts *bind.FilterOpts, nonce []*big.Int, sender []common.Address, target []common.Address) (*L2ToL1MessagePasserMessagePassedIterator, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _L2ToL1MessagePasser.contract.FilterLogs(opts, "MessagePassed", nonceRule, senderRule, targetRule)
	if err != nil {
		return nil, err
	}
	return &L2ToL1MessagePasserMessagePassedIterator{contract: _L2ToL1MessagePasser.contract, event: "MessagePassed", logs: logs, sub: sub}, nil
}

// WatchMessagePassed is a free log subscription operation binding the contract event 0x3034a173cd04c8b47937b2f95d00794df516ce57df60210b71de7ecf6d2fe4f3.
//
// Solidity: event MessagePassed(uint256 indexed nonce, address indexed sender, address indexed target, uint256 value, uint256 gasLimit, bytes data, bytes32 withdrawalHash, bytes32 rootHash)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserFilterer) WatchMessagePassed(opts *bind.WatchOpts, sink chan<- *L2ToL1MessagePasserMessagePassed, nonce []*big.Int, sender []common.Address, target []common.Address) (event.Subscription, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _L2ToL1MessagePasser.contract.WatchLogs(opts, "MessagePassed", nonceRule, senderRule, targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ToL1MessagePasserMessagePassed)
				if err := _L2ToL1MessagePasser.contract.UnpackLog(event, "MessagePassed", log); err != nil {
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

// ParseMessagePassed is a log parse operation binding the contract event 0x3034a173cd04c8b47937b2f95d00794df516ce57df60210b71de7ecf6d2fe4f3.
//
// Solidity: event MessagePassed(uint256 indexed nonce, address indexed sender, address indexed target, uint256 value, uint256 gasLimit, bytes data, bytes32 withdrawalHash, bytes32 rootHash)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserFilterer) ParseMessagePassed(log types.Log) (*L2ToL1MessagePasserMessagePassed, error) {
	event := new(L2ToL1MessagePasserMessagePassed)
	if err := _L2ToL1MessagePasser.contract.UnpackLog(event, "MessagePassed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2ToL1MessagePasserWithdrawerBalanceBurntIterator is returned from FilterWithdrawerBalanceBurnt and is used to iterate over the raw logs and unpacked data for WithdrawerBalanceBurnt events raised by the L2ToL1MessagePasser contract.
type L2ToL1MessagePasserWithdrawerBalanceBurntIterator struct {
	Event *L2ToL1MessagePasserWithdrawerBalanceBurnt // Event containing the contract specifics and raw log

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
func (it *L2ToL1MessagePasserWithdrawerBalanceBurntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ToL1MessagePasserWithdrawerBalanceBurnt)
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
		it.Event = new(L2ToL1MessagePasserWithdrawerBalanceBurnt)
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
func (it *L2ToL1MessagePasserWithdrawerBalanceBurntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ToL1MessagePasserWithdrawerBalanceBurntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ToL1MessagePasserWithdrawerBalanceBurnt represents a WithdrawerBalanceBurnt event raised by the L2ToL1MessagePasser contract.
type L2ToL1MessagePasserWithdrawerBalanceBurnt struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawerBalanceBurnt is a free log retrieval operation binding the contract event 0x7967de617a5ac1cc7eba2d6f37570a0135afa950d8bb77cdd35f0d0b4e85a16f.
//
// Solidity: event WithdrawerBalanceBurnt(uint256 indexed amount)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserFilterer) FilterWithdrawerBalanceBurnt(opts *bind.FilterOpts, amount []*big.Int) (*L2ToL1MessagePasserWithdrawerBalanceBurntIterator, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _L2ToL1MessagePasser.contract.FilterLogs(opts, "WithdrawerBalanceBurnt", amountRule)
	if err != nil {
		return nil, err
	}
	return &L2ToL1MessagePasserWithdrawerBalanceBurntIterator{contract: _L2ToL1MessagePasser.contract, event: "WithdrawerBalanceBurnt", logs: logs, sub: sub}, nil
}

// WatchWithdrawerBalanceBurnt is a free log subscription operation binding the contract event 0x7967de617a5ac1cc7eba2d6f37570a0135afa950d8bb77cdd35f0d0b4e85a16f.
//
// Solidity: event WithdrawerBalanceBurnt(uint256 indexed amount)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserFilterer) WatchWithdrawerBalanceBurnt(opts *bind.WatchOpts, sink chan<- *L2ToL1MessagePasserWithdrawerBalanceBurnt, amount []*big.Int) (event.Subscription, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _L2ToL1MessagePasser.contract.WatchLogs(opts, "WithdrawerBalanceBurnt", amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ToL1MessagePasserWithdrawerBalanceBurnt)
				if err := _L2ToL1MessagePasser.contract.UnpackLog(event, "WithdrawerBalanceBurnt", log); err != nil {
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

// ParseWithdrawerBalanceBurnt is a log parse operation binding the contract event 0x7967de617a5ac1cc7eba2d6f37570a0135afa950d8bb77cdd35f0d0b4e85a16f.
//
// Solidity: event WithdrawerBalanceBurnt(uint256 indexed amount)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserFilterer) ParseWithdrawerBalanceBurnt(log types.Log) (*L2ToL1MessagePasserWithdrawerBalanceBurnt, error) {
	event := new(L2ToL1MessagePasserWithdrawerBalanceBurnt)
	if err := _L2ToL1MessagePasser.contract.UnpackLog(event, "WithdrawerBalanceBurnt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
