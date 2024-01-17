// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/morph-l2/bindings/solc"
)

const L2ETHGatewayStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L2/gateways/L2ETHGateway.sol:L2ETHGateway\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"contracts/L2/gateways/L2ETHGateway.sol:L2ETHGateway\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"contracts/L2/gateways/L2ETHGateway.sol:L2ETHGateway\",\"label\":\"_status\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_uint256\"},{\"astId\":1003,\"contract\":\"contracts/L2/gateways/L2ETHGateway.sol:L2ETHGateway\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_array(t_uint256)1013_storage\"},{\"astId\":1004,\"contract\":\"contracts/L2/gateways/L2ETHGateway.sol:L2ETHGateway\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_array(t_uint256)1014_storage\"},{\"astId\":1005,\"contract\":\"contracts/L2/gateways/L2ETHGateway.sol:L2ETHGateway\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_address\"},{\"astId\":1006,\"contract\":\"contracts/L2/gateways/L2ETHGateway.sol:L2ETHGateway\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_array(t_uint256)1013_storage\"},{\"astId\":1007,\"contract\":\"contracts/L2/gateways/L2ETHGateway.sol:L2ETHGateway\",\"label\":\"counterpart\",\"offset\":0,\"slot\":\"151\",\"type\":\"t_address\"},{\"astId\":1008,\"contract\":\"contracts/L2/gateways/L2ETHGateway.sol:L2ETHGateway\",\"label\":\"router\",\"offset\":0,\"slot\":\"152\",\"type\":\"t_address\"},{\"astId\":1009,\"contract\":\"contracts/L2/gateways/L2ETHGateway.sol:L2ETHGateway\",\"label\":\"messenger\",\"offset\":0,\"slot\":\"153\",\"type\":\"t_address\"},{\"astId\":1010,\"contract\":\"contracts/L2/gateways/L2ETHGateway.sol:L2ETHGateway\",\"label\":\"__rateLimiter\",\"offset\":0,\"slot\":\"154\",\"type\":\"t_address\"},{\"astId\":1011,\"contract\":\"contracts/L2/gateways/L2ETHGateway.sol:L2ETHGateway\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"155\",\"type\":\"t_array(t_uint256)1012_storage\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)1012_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[46]\",\"numberOfBytes\":\"1472\"},\"t_array(t_uint256)1013_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\"},\"t_array(t_uint256)1014_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L2ETHGatewayStorageLayout = new(solc.StorageLayout)

var L2ETHGatewayDeployedBin = "0x6080604052600436106100bc5760003560e01c8063797594b011610074578063c7cdea371161004e578063c7cdea37146101df578063f2fde38b146101f2578063f887ea401461021257600080fd5b8063797594b0146101675780638da5cb5b14610194578063c0c53b8b146101bf57600080fd5b80633cb747bf116100a55780633cb747bf146100e95780636dc241831461013f578063715018a61461015257600080fd5b8063232e8748146100c15780632fcc29fa146100d6575b600080fd5b6100d46100cf366004611192565b61023f565b005b6100d46100e4366004611231565b6105af565b3480156100f557600080fd5b506099546101169073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b6100d461014d36600461132a565b6105ed565b34801561015e57600080fd5b506100d46105ff565b34801561017357600080fd5b506097546101169073ffffffffffffffffffffffffffffffffffffffff1681565b3480156101a057600080fd5b5060655473ffffffffffffffffffffffffffffffffffffffff16610116565b3480156101cb57600080fd5b506100d46101da3660046113ce565b610613565b6100d46101ed366004611419565b610827565b3480156101fe57600080fd5b506100d461020d36600461143b565b610837565b34801561021e57600080fd5b506098546101169073ffffffffffffffffffffffffffffffffffffffff1681565b60995473ffffffffffffffffffffffffffffffffffffffff163381146102c6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f6f6e6c79206d657373656e6765722063616e2063616c6c00000000000000000060448201526064015b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa158015610311573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610335919061145f565b60975473ffffffffffffffffffffffffffffffffffffffff9081169116146103b9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6f6e6c792063616c6c20627920636f756e74657270617274000000000000000060448201526064016102bd565b6103c16108ee565b83341461042a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f6d73672e76616c7565206d69736d61746368000000000000000000000000000060448201526064016102bd565b60008573ffffffffffffffffffffffffffffffffffffffff168560405160006040518083038185875af1925050503d8060008114610484576040519150601f19603f3d011682016040523d82523d6000602084013e610489565b606091505b50509050806104f4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f455448207472616e73666572206661696c65640000000000000000000000000060448201526064016102bd565b6105348685858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061096192505050565b8573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167f9e86c356e14e24e26e3ce769bf8b87de38e0faa0ed0ca946fa09659aa606bd2d8787876040516105959392919061147c565b60405180910390a3506105a760018055565b505050505050565b6105e8838360005b6040519080825280601f01601f1916602001820160405280156105e1576020820181803683370190505b5084610a14565b505050565b6105f984848484610a14565b50505050565b610607610c58565b6106116000610cd9565b565b600054610100900460ff16158080156106335750600054600160ff909116105b8061064d5750303b15801561064d575060005460ff166001145b6106d9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016102bd565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561073757600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b73ffffffffffffffffffffffffffffffffffffffff83166107b4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f7a65726f20726f7574657220616464726573730000000000000000000000000060448201526064016102bd565b6107bf848484610d50565b80156105f957600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a150505050565b610833338360006105b7565b5050565b61083f610c58565b73ffffffffffffffffffffffffffffffffffffffff81166108e2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016102bd565b6108eb81610cd9565b50565b60026001540361095a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016102bd565b6002600155565b60008151118015610989575060008273ffffffffffffffffffffffffffffffffffffffff163b115b15610833576040517f444b281f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83169063444b281f906109e090849060040161153e565b600060405180830381600087803b1580156109fa57600080fd5b505af11580156105a7573d6000803e3d6000fd5b60018055565b610a1c6108ee565b60003411610a86576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f7769746864726177207a65726f2065746800000000000000000000000000000060448201526064016102bd565b609854339073ffffffffffffffffffffffffffffffffffffffff16819003610ac15782806020019051810190610abc9190611551565b935090505b600081868686604051602401610ada94939291906115de565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f8eaac8a30000000000000000000000000000000000000000000000000000000017905260995460975491517fb2267a7b00000000000000000000000000000000000000000000000000000000815292935073ffffffffffffffffffffffffffffffffffffffff9081169263b2267a7b923492610bb4929116908a9087908a90600401611627565b6000604051808303818588803b158015610bcd57600080fd5b505af1158015610be1573d6000803e3d6000fd5b50505050508573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fd8ed6eaa9a7a8980d7901e911fde6686810b989d3082182d1d3a3df6306ce20e8787604051610c4592919061166d565b60405180910390a350506105f960018055565b60655473ffffffffffffffffffffffffffffffffffffffff163314610611576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102bd565b6065805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b73ffffffffffffffffffffffffffffffffffffffff8316610dcd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f7a65726f20636f756e746572706172742061646472657373000000000000000060448201526064016102bd565b73ffffffffffffffffffffffffffffffffffffffff8116610e4a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f7a65726f206d657373656e67657220616464726573730000000000000000000060448201526064016102bd565b610e52610efb565b610e5a610f9a565b6097805473ffffffffffffffffffffffffffffffffffffffff8086167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316179092556099805484841692169190911790558216156105e8576098805473ffffffffffffffffffffffffffffffffffffffff84167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116179055505050565b600054610100900460ff16610f92576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016102bd565b610611611039565b600054610100900460ff16611031576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016102bd565b6106116110d0565b600054610100900460ff16610a0e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016102bd565b600054610100900460ff16611167576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016102bd565b61061133610cd9565b73ffffffffffffffffffffffffffffffffffffffff811681146108eb57600080fd5b6000806000806000608086880312156111aa57600080fd5b85356111b581611170565b945060208601356111c581611170565b935060408601359250606086013567ffffffffffffffff808211156111e957600080fd5b818801915088601f8301126111fd57600080fd5b81358181111561120c57600080fd5b89602082850101111561121e57600080fd5b9699959850939650602001949392505050565b60008060006060848603121561124657600080fd5b833561125181611170565b95602085013595506040909401359392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156112dc576112dc611266565b604052919050565b600067ffffffffffffffff8211156112fe576112fe611266565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b6000806000806080858703121561134057600080fd5b843561134b81611170565b935060208501359250604085013567ffffffffffffffff81111561136e57600080fd5b8501601f8101871361137f57600080fd5b803561139261138d826112e4565b611295565b8181528860208385010111156113a757600080fd5b81602084016020830137600091810160200191909152949793965093946060013593505050565b6000806000606084860312156113e357600080fd5b83356113ee81611170565b925060208401356113fe81611170565b9150604084013561140e81611170565b809150509250925092565b6000806040838503121561142c57600080fd5b50508035926020909101359150565b60006020828403121561144d57600080fd5b813561145881611170565b9392505050565b60006020828403121561147157600080fd5b815161145881611170565b83815260406020820152816040820152818360608301376000818301606090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016010192915050565b60005b838110156114eb5781810151838201526020016114d3565b50506000910152565b6000815180845261150c8160208601602086016114d0565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061145860208301846114f4565b6000806040838503121561156457600080fd5b825161156f81611170565b602084015190925067ffffffffffffffff81111561158c57600080fd5b8301601f8101851361159d57600080fd5b80516115ab61138d826112e4565b8181528660208385010111156115c057600080fd5b6115d18260208301602086016114d0565b8093505050509250929050565b600073ffffffffffffffffffffffffffffffffffffffff80871683528086166020840152508360408301526080606083015261161d60808301846114f4565b9695505050505050565b73ffffffffffffffffffffffffffffffffffffffff8516815283602082015260806040820152600061165c60808301856114f4565b905082606083015295945050505050565b82815260406020820152600061168660408301846114f4565b94935050505056fea164736f6c6343000810000a"

func init() {
	if err := json.Unmarshal([]byte(L2ETHGatewayStorageLayoutJSON), L2ETHGatewayStorageLayout); err != nil {
		panic(err)
	}

	layouts["L2ETHGateway"] = L2ETHGatewayStorageLayout
	deployedBytecodes["L2ETHGateway"] = L2ETHGatewayDeployedBin
}
