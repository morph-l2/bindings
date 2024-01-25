// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/morph-l2/bindings/solc"
)

const MintManagerStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/governance/MintManager.sol:MintManager\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":1001,\"contract\":\"contracts/governance/MintManager.sol:MintManager\",\"label\":\"mintPermittedAfter\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_uint256\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"}}}"

var MintManagerStorageLayout = new(solc.StorageLayout)

var MintManagerDeployedBin = "0x608060405234801561001057600080fd5b50600436106100bd5760003560e01c80638da5cb5b1161007657806398f1312e1161005b57806398f1312e14610161578063f2fde38b14610169578063f96dae0a1461017c57600080fd5b80638da5cb5b14610119578063918f86741461015857600080fd5b806340c10f19116100a757806340c10f19146100f3578063715018a61461010657806383ea6e971461010e57600080fd5b8062f8900c146100c25780630900f010146100de575b600080fd5b6100cb60015481565b6040519081526020015b60405180910390f35b6100f16100ec366004610776565b6101a3565b005b6100f1610101366004610798565b6102f7565b6100f161058c565b6100cb6301e1338081565b60005473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100d5565b6100cb6103e881565b6100cb601481565b6100f1610177366004610776565b6105a0565b6101337f000000000000000000000000000000000000000000000000000000000000000081565b6101ab610657565b73ffffffffffffffffffffffffffffffffffffffff8116610253576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603460248201527f4d696e744d616e616765723a206d696e74206d616e616765722063616e6e6f7460448201527f20626520746865207a65726f206164647265737300000000000000000000000060648201526084015b60405180910390fd5b6040517ff2fde38b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063f2fde38b90602401600060405180830381600087803b1580156102dc57600080fd5b505af11580156102f0573d6000803e3d6000fd5b5050505050565b6102ff610657565b600154156104cf57426001541115610399576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4d696e744d616e616765723a206d696e74696e67206e6f74207065726d69747460448201527f6564207965740000000000000000000000000000000000000000000000000000606482015260840161024a565b6103e860147f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610409573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061042d91906107c2565b610437919061080a565b6104419190610847565b8111156104cf576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f4d696e744d616e616765723a206d696e7420616d6f756e74206578636565647360448201527f2063617000000000000000000000000000000000000000000000000000000000606482015260840161024a565b6104dd6301e1338042610882565b6001556040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8381166004830152602482018390527f000000000000000000000000000000000000000000000000000000000000000016906340c10f1990604401600060405180830381600087803b15801561057057600080fd5b505af1158015610584573d6000803e3d6000fd5b505050505050565b610594610657565b61059e60006106d8565b565b6105a8610657565b73ffffffffffffffffffffffffffffffffffffffff811661064b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161024a565b610654816106d8565b50565b60005473ffffffffffffffffffffffffffffffffffffffff16331461059e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161024a565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461077157600080fd5b919050565b60006020828403121561078857600080fd5b6107918261074d565b9392505050565b600080604083850312156107ab57600080fd5b6107b48361074d565b946020939093013593505050565b6000602082840312156107d457600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615610842576108426107db565b500290565b60008261087d577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b80820180821115610895576108956107db565b9291505056fea164736f6c6343000810000a"

func init() {
	if err := json.Unmarshal([]byte(MintManagerStorageLayoutJSON), MintManagerStorageLayout); err != nil {
		panic(err)
	}

	layouts["MintManager"] = MintManagerStorageLayout
	deployedBytecodes["MintManager"] = MintManagerDeployedBin
}
