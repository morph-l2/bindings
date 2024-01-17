// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/morph-l2/bindings/solc"
)

const GasPriceOracleStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L2/system/GasPriceOracle.sol:GasPriceOracle\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":1001,\"contract\":\"contracts/L2/system/GasPriceOracle.sol:GasPriceOracle\",\"label\":\"l1BaseFee\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_uint256\"},{\"astId\":1002,\"contract\":\"contracts/L2/system/GasPriceOracle.sol:GasPriceOracle\",\"label\":\"overhead\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_uint256\"},{\"astId\":1003,\"contract\":\"contracts/L2/system/GasPriceOracle.sol:GasPriceOracle\",\"label\":\"scalar\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_uint256\"},{\"astId\":1004,\"contract\":\"contracts/L2/system/GasPriceOracle.sol:GasPriceOracle\",\"label\":\"allowListEnabled\",\"offset\":0,\"slot\":\"4\",\"type\":\"t_bool\"},{\"astId\":1005,\"contract\":\"contracts/L2/system/GasPriceOracle.sol:GasPriceOracle\",\"label\":\"isAllowed\",\"offset\":0,\"slot\":\"5\",\"type\":\"t_mapping(t_address,t_bool)\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_mapping(t_address,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_bool\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"}}}"

var GasPriceOracleStorageLayout = new(solc.StorageLayout)

var GasPriceOracleDeployedBin = "0x608060405234801561001057600080fd5b50600436106100df5760003560e01c80638da5cb5b1161008c578063e3de72a511610066578063e3de72a5146101b4578063efeadb6d146101c7578063f2fde38b146101da578063f45e65d8146101ed57600080fd5b80638da5cb5b14610156578063babcc5391461017e578063bede39b5146101a157600080fd5b8063519b4bd3116100bd578063519b4bd314610132578063704655971461013b578063715018a61461014e57600080fd5b80630c18c162146100e457806322bd5c1c146101005780633577afc51461011d575b600080fd5b6100ed60025481565b6040519081526020015b60405180910390f35b60045461010d9060ff1681565b60405190151581526020016100f7565b61013061012b36600461092d565b6101f6565b005b6100ed60015481565b61013061014936600461092d565b6102fc565b6101306103f6565b60005460405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100f7565b61010d61018c36600461096f565b60056020526000908152604090205460ff1681565b6101306101af36600461092d565b61040a565b6101306101c2366004610ab5565b610504565b6101306101d5366004610b75565b6106a7565b6101306101e836600461096f565b610780565b6100ed60035481565b3361021660005473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff16148061023b575060045460ff16155b8061025557503360009081526005602052604090205460ff165b6102c0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f6e6f7420616c6c6f77656400000000000000000000000000000000000000000060448201526064015b60405180910390fd5b60028190556040518181527f32740b35c0ea213650f60d44366b4fb211c9033b50714e4a1d34e65d5beb9bb4906020015b60405180910390a150565b3361031c60005473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff161480610341575060045460ff16155b8061035b57503360009081526005602052604090205460ff165b6103c1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f6e6f7420616c6c6f77656400000000000000000000000000000000000000000060448201526064016102b7565b60038190556040518181527f3336cd9708eaf2769a0f0dc0679f30e80f15dcd88d1921b5a16858e8b85c591a906020016102f1565b6103fe610837565b61040860006108b8565b565b3361042a60005473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff16148061044f575060045460ff16155b8061046957503360009081526005602052604090205460ff165b6104cf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f6e6f7420616c6c6f77656400000000000000000000000000000000000000000060448201526064016102b7565b60018190556040518181527f351fb23757bb5ea0546c85b7996ddd7155f96b939ebaa5ff7bc49c75f27f2c44906020016102f1565b61050c610837565b8051825114610577576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f494e56414c49445f494e5055540000000000000000000000000000000000000060448201526064016102b7565b60005b82518110156106a25781818151811061059557610595610b90565b6020026020010151600560008584815181106105b3576105b3610b90565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555082818151811061061e5761061e610b90565b602002602001015173ffffffffffffffffffffffffffffffffffffffff167fd9739f45a01ce092c5cdb3d68f63d63d21676b1c6c0b4f9cbc6be4cf5449595a83838151811061066f5761066f610b90565b6020026020010151604051610688911515815260200190565b60405180910390a28061069a81610bbf565b91505061057a565b505050565b6106af610837565b60045460ff16151581151503610721576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f414c52454144595f53455400000000000000000000000000000000000000000060448201526064016102b7565b600480547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00168215159081179091556040519081527f16435b45f7482047f839a6a19d291442627200f52cad2803c595150d0d440eb3906020016102f1565b610788610837565b73ffffffffffffffffffffffffffffffffffffffff811661082b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016102b7565b610834816108b8565b50565b60005473ffffffffffffffffffffffffffffffffffffffff163314610408576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102b7565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b60006020828403121561093f57600080fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461096a57600080fd5b919050565b60006020828403121561098157600080fd5b61098a82610946565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715610a0757610a07610991565b604052919050565b600067ffffffffffffffff821115610a2957610a29610991565b5060051b60200190565b8035801515811461096a57600080fd5b600082601f830112610a5457600080fd5b81356020610a69610a6483610a0f565b6109c0565b82815260059290921b84018101918181019086841115610a8857600080fd5b8286015b84811015610aaa57610a9d81610a33565b8352918301918301610a8c565b509695505050505050565b60008060408385031215610ac857600080fd5b823567ffffffffffffffff80821115610ae057600080fd5b818501915085601f830112610af457600080fd5b81356020610b04610a6483610a0f565b82815260059290921b84018101918181019089841115610b2357600080fd5b948201945b83861015610b4857610b3986610946565b82529482019490820190610b28565b96505086013592505080821115610b5e57600080fd5b50610b6b85828601610a43565b9150509250929050565b600060208284031215610b8757600080fd5b61098a82610a33565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610c17577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b506001019056fea164736f6c6343000810000a"

func init() {
	if err := json.Unmarshal([]byte(GasPriceOracleStorageLayoutJSON), GasPriceOracleStorageLayout); err != nil {
		panic(err)
	}

	layouts["GasPriceOracle"] = GasPriceOracleStorageLayout
	deployedBytecodes["GasPriceOracle"] = GasPriceOracleDeployedBin
}
