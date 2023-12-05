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

// GovProposalData is an auto generated low-level Go binding around an user-defined struct.
type GovProposalData struct {
	BatchBlockInterval *big.Int
	BatchMaxBytes      *big.Int
	BatchTimeout       *big.Int
	RollupEpoch        *big.Int
}

// GovMetaData contains all meta data concerning the Gov contract.
var GovMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"L2_SEQUENCER_CONTRACT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"L2_SUBMITTER_CONTRACT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"batchBlockInterval\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"batchMaxBytes\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"batchTimeout\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_proposalInterval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_batchBlockInterval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_batchMaxBytes\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_batchTimeout\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_rollupEpoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"l2Sequencer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposalData\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"batchBlockInterval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"batchMaxBytes\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"batchTimeout\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"rollupEpoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposalInfos\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"endTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"seqsVersion\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"votes\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposalInterval\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposalNumbers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"propose\",\"inputs\":[{\"name\":\"proposal\",\"type\":\"tuple\",\"internalType\":\"structGov.ProposalData\",\"components\":[{\"name\":\"batchBlockInterval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"batchMaxBytes\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"batchTimeout\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"rollupEpoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rollupEpoch\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sequencersVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"vote\",\"inputs\":[{\"name\":\"propID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"votes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false}]",
	Bin: "0x60c060405260006001556000600255600060035560006004556000600555600060065534801561002e57600080fd5b5073530000000000000000000000000000000000000360805273530000000000000000000000000000000000000560a05260805160a0516112206100b56000396000818161018201526108a801526000818161011c015281816101a90152818161034d01528181610502015281816107aa01528181610cae0152610e2f01526112206000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c806396dea93611610097578063dcbab60811610066578063dcbab608146102dd578063e5aec995146102f0578063f59993a4146102f9578063fb1ec34e1461030257600080fd5b806396dea936146101e6578063b511328d1461023b578063c6e36a3214610296578063d441a168146102d457600080fd5b80636cb23707116100d35780636cb23707146101a457806377c79380146101cb57806385963052146101d4578063929a9cbe146101dd57600080fd5b80630121b93f146101055780630c3f35171461011a5780634063a84e146101665780634bbf52521461017d575b600080fd5b610118610113366004610fa8565b610315565b005b7f00000000000000000000000000000000000000000000000000000000000000005b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b61016f600a5481565b60405190815260200161015d565b61013c7f000000000000000000000000000000000000000000000000000000000000000081565b61013c7f000000000000000000000000000000000000000000000000000000000000000081565b61016f60045481565b61016f60025481565b61016f60035481565b61021b6101f4366004610fa8565b60086020526000908152604090208054600182015460028301546003909301549192909184565b60408051948552602085019390935291830152606082015260800161015d565b610274610249366004610fa8565b600960205260009081526040902080546001820154600283015460039093015460ff90921692909184565b604080519415158552602085019390935291830152606082015260800161015d565b6102c46102a4366004610fc1565b600760209081526000928352604080842090915290825290205460ff1681565b604051901515815260200161015d565b61016f60015481565b6101186102eb366004610fe3565b61097d565b61016f60055481565b61016f60065481565b61011861031036600461101e565b610c76565b6040517fd1c55fe3000000000000000000000000000000000000000000000000000000008152600060048201819052336024830152907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063d1c55fe3906044016040805180830381865afa1580156103a8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103cc91906110ab565b5090508061043b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f6f6e6c792073657175656e6365722063616e2070726f706f736500000000000060448201526064015b60405180910390fd5b60008281526009602052604090205460ff166104b3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f70726f706f73616c20696e6163746976650000000000000000000000000000006044820152606401610432565b6040517f342b634500000000000000000000000000000000000000000000000000000000815260006004820181905233602483015290819073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063342b6345906044016040805180830381865afa158015610548573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061056c91906110de565b6000868152600760209081526040808320858452909152902054919350915060ff161561061b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602860248201527f73657175656e63657220616c726561647920766f746520666f7220746869732060448201527f70726f706f73616c0000000000000000000000000000000000000000000000006064820152608401610432565b6000848152600960205260409020600201548114610695576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f76657273696f6e206d69736d61746368000000000000000000000000000000006044820152606401610432565b600084815260096020526040902060010154421115610710576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600860248201527f74696d6520656e640000000000000000000000000000000000000000000000006044820152606401610432565b6000848152600760209081526040808320858452825280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600190811790915587845260099092528220600301805491929091610773908490611131565b90915550506040517f7ad9e3ac000000000000000000000000000000000000000000000000000000008152600060048201819052907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690637ad9e3ac906024016040805180830381865afa158015610805573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061082991906110de565b509050600361083982600261114a565b6108439190611187565b6000868152600960205260409020600301541115610976576000858152600860205260409020600301546005541461091a576005546040517f965fbb9400000000000000000000000000000000000000000000000000000000815260048101919091527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063965fbb9490602401600060405180830381600087803b15801561090157600080fd5b505af1158015610915573d6000803e3d6000fd5b505050505b600085815260086020908152604080832080546002908155600182015460039081559082015460045501546005556009909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690555b5050505050565b600054610100900460ff161580801561099d5750600054600160ff909116105b806109b75750303b1580156109b7575060005460ff166001145b610a43576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610432565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015610aa157600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b60008611610b0b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f696e76616c69642070726f706f73616c20696e74657276616c000000000000006044820152606401610432565b60008211610b75576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f696e76616c696420726f6c6c75702065706f63680000000000000000000000006044820152606401610432565b84151580610b8257508315155b80610b8c57508215155b610bf2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f696e76616c696420626174636820706172616d730000000000000000000000006044820152606401610432565b600a86905560028590556003849055600483905560058290558015610c6e57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050505050565b6040517fd1c55fe3000000000000000000000000000000000000000000000000000000008152600060048201819052336024830152907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063d1c55fe3906044016040805180830381865afa158015610d09573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d2d91906110ab565b50905080610d97576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f6f6e6c792073657175656e6365722063616e2070726f706f73650000000000006044820152606401610432565b8151151580610da95750602082015115155b80610db75750604082015115155b80610dc55750606082015115155b610e2b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f696e76616c696420626174636820706172616d730000000000000000000000006044820152606401610432565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16639d888e866040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e98573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ebc91906111c2565b9050600154811115610ece5760018190555b60068054906000610ede836111db565b9091555050600654600090815260086020908152604091829020855181558186015160018083019190915583870151600283015560608701516003909201919091558251608081019093528252600a5490820190610f3c9042611131565b815260208082019390935260006040918201819052600654815260098452819020825181547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169015151781559282015160018401558101516002830155606001516003909101555050565b600060208284031215610fba57600080fd5b5035919050565b60008060408385031215610fd457600080fd5b50508035926020909101359150565b600080600080600060a08688031215610ffb57600080fd5b505083359560208501359550604085013594606081013594506080013592509050565b60006080828403121561103057600080fd5b6040516080810181811067ffffffffffffffff8211171561107a577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b8060405250823581526020830135602082015260408301356040820152606083013560608201528091505092915050565b600080604083850312156110be57600080fd5b825180151581146110ce57600080fd5b6020939093015192949293505050565b600080604083850312156110f157600080fd5b505080516020909101519092909150565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8082018082111561114457611144611102565b92915050565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561118257611182611102565b500290565b6000826111bd577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b6000602082840312156111d457600080fd5b5051919050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361120c5761120c611102565b506001019056fea164736f6c6343000810000a",
}

// GovABI is the input ABI used to generate the binding from.
// Deprecated: Use GovMetaData.ABI instead.
var GovABI = GovMetaData.ABI

// GovBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GovMetaData.Bin instead.
var GovBin = GovMetaData.Bin

// DeployGov deploys a new Ethereum contract, binding an instance of Gov to it.
func DeployGov(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Gov, error) {
	parsed, err := GovMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GovBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Gov{GovCaller: GovCaller{contract: contract}, GovTransactor: GovTransactor{contract: contract}, GovFilterer: GovFilterer{contract: contract}}, nil
}

// Gov is an auto generated Go binding around an Ethereum contract.
type Gov struct {
	GovCaller     // Read-only binding to the contract
	GovTransactor // Write-only binding to the contract
	GovFilterer   // Log filterer for contract events
}

// GovCaller is an auto generated read-only Go binding around an Ethereum contract.
type GovCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GovTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GovFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GovSession struct {
	Contract     *Gov              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GovCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GovCallerSession struct {
	Contract *GovCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// GovTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GovTransactorSession struct {
	Contract     *GovTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GovRaw is an auto generated low-level Go binding around an Ethereum contract.
type GovRaw struct {
	Contract *Gov // Generic contract binding to access the raw methods on
}

// GovCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GovCallerRaw struct {
	Contract *GovCaller // Generic read-only contract binding to access the raw methods on
}

// GovTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GovTransactorRaw struct {
	Contract *GovTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGov creates a new instance of Gov, bound to a specific deployed contract.
func NewGov(address common.Address, backend bind.ContractBackend) (*Gov, error) {
	contract, err := bindGov(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Gov{GovCaller: GovCaller{contract: contract}, GovTransactor: GovTransactor{contract: contract}, GovFilterer: GovFilterer{contract: contract}}, nil
}

// NewGovCaller creates a new read-only instance of Gov, bound to a specific deployed contract.
func NewGovCaller(address common.Address, caller bind.ContractCaller) (*GovCaller, error) {
	contract, err := bindGov(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GovCaller{contract: contract}, nil
}

// NewGovTransactor creates a new write-only instance of Gov, bound to a specific deployed contract.
func NewGovTransactor(address common.Address, transactor bind.ContractTransactor) (*GovTransactor, error) {
	contract, err := bindGov(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GovTransactor{contract: contract}, nil
}

// NewGovFilterer creates a new log filterer instance of Gov, bound to a specific deployed contract.
func NewGovFilterer(address common.Address, filterer bind.ContractFilterer) (*GovFilterer, error) {
	contract, err := bindGov(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GovFilterer{contract: contract}, nil
}

// bindGov binds a generic wrapper to an already deployed contract.
func bindGov(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GovABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gov *GovRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gov.Contract.GovCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gov *GovRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gov.Contract.GovTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gov *GovRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gov.Contract.GovTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gov *GovCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gov.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gov *GovTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gov.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gov *GovTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gov.Contract.contract.Transact(opts, method, params...)
}

// L2SEQUENCERCONTRACT is a free data retrieval call binding the contract method 0x6cb23707.
//
// Solidity: function L2_SEQUENCER_CONTRACT() view returns(address)
func (_Gov *GovCaller) L2SEQUENCERCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gov.contract.Call(opts, &out, "L2_SEQUENCER_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2SEQUENCERCONTRACT is a free data retrieval call binding the contract method 0x6cb23707.
//
// Solidity: function L2_SEQUENCER_CONTRACT() view returns(address)
func (_Gov *GovSession) L2SEQUENCERCONTRACT() (common.Address, error) {
	return _Gov.Contract.L2SEQUENCERCONTRACT(&_Gov.CallOpts)
}

// L2SEQUENCERCONTRACT is a free data retrieval call binding the contract method 0x6cb23707.
//
// Solidity: function L2_SEQUENCER_CONTRACT() view returns(address)
func (_Gov *GovCallerSession) L2SEQUENCERCONTRACT() (common.Address, error) {
	return _Gov.Contract.L2SEQUENCERCONTRACT(&_Gov.CallOpts)
}

// L2SUBMITTERCONTRACT is a free data retrieval call binding the contract method 0x4bbf5252.
//
// Solidity: function L2_SUBMITTER_CONTRACT() view returns(address)
func (_Gov *GovCaller) L2SUBMITTERCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gov.contract.Call(opts, &out, "L2_SUBMITTER_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2SUBMITTERCONTRACT is a free data retrieval call binding the contract method 0x4bbf5252.
//
// Solidity: function L2_SUBMITTER_CONTRACT() view returns(address)
func (_Gov *GovSession) L2SUBMITTERCONTRACT() (common.Address, error) {
	return _Gov.Contract.L2SUBMITTERCONTRACT(&_Gov.CallOpts)
}

// L2SUBMITTERCONTRACT is a free data retrieval call binding the contract method 0x4bbf5252.
//
// Solidity: function L2_SUBMITTER_CONTRACT() view returns(address)
func (_Gov *GovCallerSession) L2SUBMITTERCONTRACT() (common.Address, error) {
	return _Gov.Contract.L2SUBMITTERCONTRACT(&_Gov.CallOpts)
}

// BatchBlockInterval is a free data retrieval call binding the contract method 0x85963052.
//
// Solidity: function batchBlockInterval() view returns(uint256)
func (_Gov *GovCaller) BatchBlockInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gov.contract.Call(opts, &out, "batchBlockInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BatchBlockInterval is a free data retrieval call binding the contract method 0x85963052.
//
// Solidity: function batchBlockInterval() view returns(uint256)
func (_Gov *GovSession) BatchBlockInterval() (*big.Int, error) {
	return _Gov.Contract.BatchBlockInterval(&_Gov.CallOpts)
}

// BatchBlockInterval is a free data retrieval call binding the contract method 0x85963052.
//
// Solidity: function batchBlockInterval() view returns(uint256)
func (_Gov *GovCallerSession) BatchBlockInterval() (*big.Int, error) {
	return _Gov.Contract.BatchBlockInterval(&_Gov.CallOpts)
}

// BatchMaxBytes is a free data retrieval call binding the contract method 0x929a9cbe.
//
// Solidity: function batchMaxBytes() view returns(uint256)
func (_Gov *GovCaller) BatchMaxBytes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gov.contract.Call(opts, &out, "batchMaxBytes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BatchMaxBytes is a free data retrieval call binding the contract method 0x929a9cbe.
//
// Solidity: function batchMaxBytes() view returns(uint256)
func (_Gov *GovSession) BatchMaxBytes() (*big.Int, error) {
	return _Gov.Contract.BatchMaxBytes(&_Gov.CallOpts)
}

// BatchMaxBytes is a free data retrieval call binding the contract method 0x929a9cbe.
//
// Solidity: function batchMaxBytes() view returns(uint256)
func (_Gov *GovCallerSession) BatchMaxBytes() (*big.Int, error) {
	return _Gov.Contract.BatchMaxBytes(&_Gov.CallOpts)
}

// BatchTimeout is a free data retrieval call binding the contract method 0x77c79380.
//
// Solidity: function batchTimeout() view returns(uint256)
func (_Gov *GovCaller) BatchTimeout(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gov.contract.Call(opts, &out, "batchTimeout")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BatchTimeout is a free data retrieval call binding the contract method 0x77c79380.
//
// Solidity: function batchTimeout() view returns(uint256)
func (_Gov *GovSession) BatchTimeout() (*big.Int, error) {
	return _Gov.Contract.BatchTimeout(&_Gov.CallOpts)
}

// BatchTimeout is a free data retrieval call binding the contract method 0x77c79380.
//
// Solidity: function batchTimeout() view returns(uint256)
func (_Gov *GovCallerSession) BatchTimeout() (*big.Int, error) {
	return _Gov.Contract.BatchTimeout(&_Gov.CallOpts)
}

// L2Sequencer is a free data retrieval call binding the contract method 0x0c3f3517.
//
// Solidity: function l2Sequencer() view returns(address)
func (_Gov *GovCaller) L2Sequencer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gov.contract.Call(opts, &out, "l2Sequencer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2Sequencer is a free data retrieval call binding the contract method 0x0c3f3517.
//
// Solidity: function l2Sequencer() view returns(address)
func (_Gov *GovSession) L2Sequencer() (common.Address, error) {
	return _Gov.Contract.L2Sequencer(&_Gov.CallOpts)
}

// L2Sequencer is a free data retrieval call binding the contract method 0x0c3f3517.
//
// Solidity: function l2Sequencer() view returns(address)
func (_Gov *GovCallerSession) L2Sequencer() (common.Address, error) {
	return _Gov.Contract.L2Sequencer(&_Gov.CallOpts)
}

// ProposalData is a free data retrieval call binding the contract method 0x96dea936.
//
// Solidity: function proposalData(uint256 ) view returns(uint256 batchBlockInterval, uint256 batchMaxBytes, uint256 batchTimeout, uint256 rollupEpoch)
func (_Gov *GovCaller) ProposalData(opts *bind.CallOpts, arg0 *big.Int) (struct {
	BatchBlockInterval *big.Int
	BatchMaxBytes      *big.Int
	BatchTimeout       *big.Int
	RollupEpoch        *big.Int
}, error) {
	var out []interface{}
	err := _Gov.contract.Call(opts, &out, "proposalData", arg0)

	outstruct := new(struct {
		BatchBlockInterval *big.Int
		BatchMaxBytes      *big.Int
		BatchTimeout       *big.Int
		RollupEpoch        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BatchBlockInterval = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BatchMaxBytes = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BatchTimeout = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.RollupEpoch = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ProposalData is a free data retrieval call binding the contract method 0x96dea936.
//
// Solidity: function proposalData(uint256 ) view returns(uint256 batchBlockInterval, uint256 batchMaxBytes, uint256 batchTimeout, uint256 rollupEpoch)
func (_Gov *GovSession) ProposalData(arg0 *big.Int) (struct {
	BatchBlockInterval *big.Int
	BatchMaxBytes      *big.Int
	BatchTimeout       *big.Int
	RollupEpoch        *big.Int
}, error) {
	return _Gov.Contract.ProposalData(&_Gov.CallOpts, arg0)
}

// ProposalData is a free data retrieval call binding the contract method 0x96dea936.
//
// Solidity: function proposalData(uint256 ) view returns(uint256 batchBlockInterval, uint256 batchMaxBytes, uint256 batchTimeout, uint256 rollupEpoch)
func (_Gov *GovCallerSession) ProposalData(arg0 *big.Int) (struct {
	BatchBlockInterval *big.Int
	BatchMaxBytes      *big.Int
	BatchTimeout       *big.Int
	RollupEpoch        *big.Int
}, error) {
	return _Gov.Contract.ProposalData(&_Gov.CallOpts, arg0)
}

// ProposalInfos is a free data retrieval call binding the contract method 0xb511328d.
//
// Solidity: function proposalInfos(uint256 ) view returns(bool active, uint256 endTime, uint256 seqsVersion, uint256 votes)
func (_Gov *GovCaller) ProposalInfos(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Active      bool
	EndTime     *big.Int
	SeqsVersion *big.Int
	Votes       *big.Int
}, error) {
	var out []interface{}
	err := _Gov.contract.Call(opts, &out, "proposalInfos", arg0)

	outstruct := new(struct {
		Active      bool
		EndTime     *big.Int
		SeqsVersion *big.Int
		Votes       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Active = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.EndTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.SeqsVersion = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Votes = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ProposalInfos is a free data retrieval call binding the contract method 0xb511328d.
//
// Solidity: function proposalInfos(uint256 ) view returns(bool active, uint256 endTime, uint256 seqsVersion, uint256 votes)
func (_Gov *GovSession) ProposalInfos(arg0 *big.Int) (struct {
	Active      bool
	EndTime     *big.Int
	SeqsVersion *big.Int
	Votes       *big.Int
}, error) {
	return _Gov.Contract.ProposalInfos(&_Gov.CallOpts, arg0)
}

// ProposalInfos is a free data retrieval call binding the contract method 0xb511328d.
//
// Solidity: function proposalInfos(uint256 ) view returns(bool active, uint256 endTime, uint256 seqsVersion, uint256 votes)
func (_Gov *GovCallerSession) ProposalInfos(arg0 *big.Int) (struct {
	Active      bool
	EndTime     *big.Int
	SeqsVersion *big.Int
	Votes       *big.Int
}, error) {
	return _Gov.Contract.ProposalInfos(&_Gov.CallOpts, arg0)
}

// ProposalInterval is a free data retrieval call binding the contract method 0x4063a84e.
//
// Solidity: function proposalInterval() view returns(uint256)
func (_Gov *GovCaller) ProposalInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gov.contract.Call(opts, &out, "proposalInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalInterval is a free data retrieval call binding the contract method 0x4063a84e.
//
// Solidity: function proposalInterval() view returns(uint256)
func (_Gov *GovSession) ProposalInterval() (*big.Int, error) {
	return _Gov.Contract.ProposalInterval(&_Gov.CallOpts)
}

// ProposalInterval is a free data retrieval call binding the contract method 0x4063a84e.
//
// Solidity: function proposalInterval() view returns(uint256)
func (_Gov *GovCallerSession) ProposalInterval() (*big.Int, error) {
	return _Gov.Contract.ProposalInterval(&_Gov.CallOpts)
}

// ProposalNumbers is a free data retrieval call binding the contract method 0xf59993a4.
//
// Solidity: function proposalNumbers() view returns(uint256)
func (_Gov *GovCaller) ProposalNumbers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gov.contract.Call(opts, &out, "proposalNumbers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalNumbers is a free data retrieval call binding the contract method 0xf59993a4.
//
// Solidity: function proposalNumbers() view returns(uint256)
func (_Gov *GovSession) ProposalNumbers() (*big.Int, error) {
	return _Gov.Contract.ProposalNumbers(&_Gov.CallOpts)
}

// ProposalNumbers is a free data retrieval call binding the contract method 0xf59993a4.
//
// Solidity: function proposalNumbers() view returns(uint256)
func (_Gov *GovCallerSession) ProposalNumbers() (*big.Int, error) {
	return _Gov.Contract.ProposalNumbers(&_Gov.CallOpts)
}

// RollupEpoch is a free data retrieval call binding the contract method 0xe5aec995.
//
// Solidity: function rollupEpoch() view returns(uint256)
func (_Gov *GovCaller) RollupEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gov.contract.Call(opts, &out, "rollupEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RollupEpoch is a free data retrieval call binding the contract method 0xe5aec995.
//
// Solidity: function rollupEpoch() view returns(uint256)
func (_Gov *GovSession) RollupEpoch() (*big.Int, error) {
	return _Gov.Contract.RollupEpoch(&_Gov.CallOpts)
}

// RollupEpoch is a free data retrieval call binding the contract method 0xe5aec995.
//
// Solidity: function rollupEpoch() view returns(uint256)
func (_Gov *GovCallerSession) RollupEpoch() (*big.Int, error) {
	return _Gov.Contract.RollupEpoch(&_Gov.CallOpts)
}

// SequencersVersion is a free data retrieval call binding the contract method 0xd441a168.
//
// Solidity: function sequencersVersion() view returns(uint256)
func (_Gov *GovCaller) SequencersVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gov.contract.Call(opts, &out, "sequencersVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SequencersVersion is a free data retrieval call binding the contract method 0xd441a168.
//
// Solidity: function sequencersVersion() view returns(uint256)
func (_Gov *GovSession) SequencersVersion() (*big.Int, error) {
	return _Gov.Contract.SequencersVersion(&_Gov.CallOpts)
}

// SequencersVersion is a free data retrieval call binding the contract method 0xd441a168.
//
// Solidity: function sequencersVersion() view returns(uint256)
func (_Gov *GovCallerSession) SequencersVersion() (*big.Int, error) {
	return _Gov.Contract.SequencersVersion(&_Gov.CallOpts)
}

// Votes is a free data retrieval call binding the contract method 0xc6e36a32.
//
// Solidity: function votes(uint256 , uint256 ) view returns(bool)
func (_Gov *GovCaller) Votes(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _Gov.contract.Call(opts, &out, "votes", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Votes is a free data retrieval call binding the contract method 0xc6e36a32.
//
// Solidity: function votes(uint256 , uint256 ) view returns(bool)
func (_Gov *GovSession) Votes(arg0 *big.Int, arg1 *big.Int) (bool, error) {
	return _Gov.Contract.Votes(&_Gov.CallOpts, arg0, arg1)
}

// Votes is a free data retrieval call binding the contract method 0xc6e36a32.
//
// Solidity: function votes(uint256 , uint256 ) view returns(bool)
func (_Gov *GovCallerSession) Votes(arg0 *big.Int, arg1 *big.Int) (bool, error) {
	return _Gov.Contract.Votes(&_Gov.CallOpts, arg0, arg1)
}

// Initialize is a paid mutator transaction binding the contract method 0xdcbab608.
//
// Solidity: function initialize(uint256 _proposalInterval, uint256 _batchBlockInterval, uint256 _batchMaxBytes, uint256 _batchTimeout, uint256 _rollupEpoch) returns()
func (_Gov *GovTransactor) Initialize(opts *bind.TransactOpts, _proposalInterval *big.Int, _batchBlockInterval *big.Int, _batchMaxBytes *big.Int, _batchTimeout *big.Int, _rollupEpoch *big.Int) (*types.Transaction, error) {
	return _Gov.contract.Transact(opts, "initialize", _proposalInterval, _batchBlockInterval, _batchMaxBytes, _batchTimeout, _rollupEpoch)
}

// Initialize is a paid mutator transaction binding the contract method 0xdcbab608.
//
// Solidity: function initialize(uint256 _proposalInterval, uint256 _batchBlockInterval, uint256 _batchMaxBytes, uint256 _batchTimeout, uint256 _rollupEpoch) returns()
func (_Gov *GovSession) Initialize(_proposalInterval *big.Int, _batchBlockInterval *big.Int, _batchMaxBytes *big.Int, _batchTimeout *big.Int, _rollupEpoch *big.Int) (*types.Transaction, error) {
	return _Gov.Contract.Initialize(&_Gov.TransactOpts, _proposalInterval, _batchBlockInterval, _batchMaxBytes, _batchTimeout, _rollupEpoch)
}

// Initialize is a paid mutator transaction binding the contract method 0xdcbab608.
//
// Solidity: function initialize(uint256 _proposalInterval, uint256 _batchBlockInterval, uint256 _batchMaxBytes, uint256 _batchTimeout, uint256 _rollupEpoch) returns()
func (_Gov *GovTransactorSession) Initialize(_proposalInterval *big.Int, _batchBlockInterval *big.Int, _batchMaxBytes *big.Int, _batchTimeout *big.Int, _rollupEpoch *big.Int) (*types.Transaction, error) {
	return _Gov.Contract.Initialize(&_Gov.TransactOpts, _proposalInterval, _batchBlockInterval, _batchMaxBytes, _batchTimeout, _rollupEpoch)
}

// Propose is a paid mutator transaction binding the contract method 0xfb1ec34e.
//
// Solidity: function propose((uint256,uint256,uint256,uint256) proposal) returns()
func (_Gov *GovTransactor) Propose(opts *bind.TransactOpts, proposal GovProposalData) (*types.Transaction, error) {
	return _Gov.contract.Transact(opts, "propose", proposal)
}

// Propose is a paid mutator transaction binding the contract method 0xfb1ec34e.
//
// Solidity: function propose((uint256,uint256,uint256,uint256) proposal) returns()
func (_Gov *GovSession) Propose(proposal GovProposalData) (*types.Transaction, error) {
	return _Gov.Contract.Propose(&_Gov.TransactOpts, proposal)
}

// Propose is a paid mutator transaction binding the contract method 0xfb1ec34e.
//
// Solidity: function propose((uint256,uint256,uint256,uint256) proposal) returns()
func (_Gov *GovTransactorSession) Propose(proposal GovProposalData) (*types.Transaction, error) {
	return _Gov.Contract.Propose(&_Gov.TransactOpts, proposal)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 propID) returns()
func (_Gov *GovTransactor) Vote(opts *bind.TransactOpts, propID *big.Int) (*types.Transaction, error) {
	return _Gov.contract.Transact(opts, "vote", propID)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 propID) returns()
func (_Gov *GovSession) Vote(propID *big.Int) (*types.Transaction, error) {
	return _Gov.Contract.Vote(&_Gov.TransactOpts, propID)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 propID) returns()
func (_Gov *GovTransactorSession) Vote(propID *big.Int) (*types.Transaction, error) {
	return _Gov.Contract.Vote(&_Gov.TransactOpts, propID)
}

// GovInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Gov contract.
type GovInitializedIterator struct {
	Event *GovInitialized // Event containing the contract specifics and raw log

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
func (it *GovInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovInitialized)
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
		it.Event = new(GovInitialized)
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
func (it *GovInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovInitialized represents a Initialized event raised by the Gov contract.
type GovInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Gov *GovFilterer) FilterInitialized(opts *bind.FilterOpts) (*GovInitializedIterator, error) {

	logs, sub, err := _Gov.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &GovInitializedIterator{contract: _Gov.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Gov *GovFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *GovInitialized) (event.Subscription, error) {

	logs, sub, err := _Gov.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovInitialized)
				if err := _Gov.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Gov *GovFilterer) ParseInitialized(log types.Log) (*GovInitialized, error) {
	event := new(GovInitialized)
	if err := _Gov.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
