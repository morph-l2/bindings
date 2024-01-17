// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/morph-l2/bindings/solc"
)

const WrappedEtherStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L2/system/WrappedEther.sol:WrappedEther\",\"label\":\"_balances\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_mapping(t_address,t_uint256)\"},{\"astId\":1001,\"contract\":\"contracts/L2/system/WrappedEther.sol:WrappedEther\",\"label\":\"_allowances\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_mapping(t_address,t_mapping(t_address,t_uint256))\"},{\"astId\":1002,\"contract\":\"contracts/L2/system/WrappedEther.sol:WrappedEther\",\"label\":\"_totalSupply\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_uint256\"},{\"astId\":1003,\"contract\":\"contracts/L2/system/WrappedEther.sol:WrappedEther\",\"label\":\"_name\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_string_storage\"},{\"astId\":1004,\"contract\":\"contracts/L2/system/WrappedEther.sol:WrappedEther\",\"label\":\"_symbol\",\"offset\":0,\"slot\":\"4\",\"type\":\"t_string_storage\"},{\"astId\":1005,\"contract\":\"contracts/L2/system/WrappedEther.sol:WrappedEther\",\"label\":\"_nameFallback\",\"offset\":0,\"slot\":\"5\",\"type\":\"t_string_storage\"},{\"astId\":1006,\"contract\":\"contracts/L2/system/WrappedEther.sol:WrappedEther\",\"label\":\"_versionFallback\",\"offset\":0,\"slot\":\"6\",\"type\":\"t_string_storage\"},{\"astId\":1007,\"contract\":\"contracts/L2/system/WrappedEther.sol:WrappedEther\",\"label\":\"_nonces\",\"offset\":0,\"slot\":\"7\",\"type\":\"t_mapping(t_address,t_struct(Counter)1009_storage)\"},{\"astId\":1008,\"contract\":\"contracts/L2/system/WrappedEther.sol:WrappedEther\",\"label\":\"_PERMIT_TYPEHASH_DEPRECATED_SLOT\",\"offset\":0,\"slot\":\"8\",\"type\":\"t_bytes32\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_mapping(t_address,t_mapping(t_address,t_uint256))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(address =\u003e uint256))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_address,t_uint256)\"},\"t_mapping(t_address,t_struct(Counter)1009_storage)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e struct Counters.Counter)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_struct(Counter)1009_storage\"},\"t_mapping(t_address,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_uint256\"},\"t_string_storage\":{\"encoding\":\"bytes\",\"label\":\"string\",\"numberOfBytes\":\"32\"},\"t_struct(Counter)1009_storage\":{\"encoding\":\"inplace\",\"label\":\"struct Counters.Counter\",\"numberOfBytes\":\"32\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"}}}"

var WrappedEtherStorageLayout = new(solc.StorageLayout)

var WrappedEtherDeployedBin = "0x6080604052600436106101125760003560e01c806370a08231116100a5578063a457c2d711610074578063d0e30db011610059578063d0e30db014610311578063d505accf14610319578063dd62ed3e1461033957600080fd5b8063a457c2d7146102d1578063a9059cbb146102f157600080fd5b806370a08231146102315780637ecebe001461027457806384b0196e1461029457806395d89b41146102bc57600080fd5b80632e1a7d4d116100e15780632e1a7d4d146101c0578063313ce567146101e05780633644e515146101fc578063395093511461021157600080fd5b806306fdde0314610126578063095ea7b31461015157806318160ddd1461018157806323b872dd146101a057600080fd5b366101215761011f61038c565b005b600080fd5b34801561013257600080fd5b5061013b6103ea565b6040516101489190611804565b60405180910390f35b34801561015d57600080fd5b5061017161016c366004611847565b61047c565b6040519015158152602001610148565b34801561018d57600080fd5b506002545b604051908152602001610148565b3480156101ac57600080fd5b506101716101bb366004611871565b610496565b3480156101cc57600080fd5b5061011f6101db3660046118ad565b6104ba565b3480156101ec57600080fd5b5060405160128152602001610148565b34801561020857600080fd5b506101926105e9565b34801561021d57600080fd5b5061017161022c366004611847565b6105f8565b34801561023d57600080fd5b5061019261024c3660046118c6565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b34801561028057600080fd5b5061019261028f3660046118c6565b610644565b3480156102a057600080fd5b506102a961066f565b60405161014897969594939291906118e1565b3480156102c857600080fd5b5061013b610714565b3480156102dd57600080fd5b506101716102ec366004611847565b610723565b3480156102fd57600080fd5b5061017161030c366004611847565b6107f4565b61011f61038c565b34801561032557600080fd5b5061011f6103343660046119a0565b610802565b34801561034557600080fd5b50610192610354366004611a13565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b3361039781346109f5565b8073ffffffffffffffffffffffffffffffffffffffff167fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c346040516103df91815260200190565b60405180910390a250565b6060600380546103f990611a46565b80601f016020809104026020016040519081016040528092919081815260200182805461042590611a46565b80156104725780601f1061044757610100808354040283529160200191610472565b820191906000526020600020905b81548152906001019060200180831161045557829003601f168201915b5050505050905090565b60003361048a818585610ae8565b60019150505b92915050565b6000336104a4858285610c9c565b6104af858585610d73565b506001949350505050565b336104c58183610fe2565b60008173ffffffffffffffffffffffffffffffffffffffff168360405160006040518083038185875af1925050503d806000811461051f576040519150601f19603f3d011682016040523d82523d6000602084013e610524565b606091505b5050905080610594576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f776974686472617720455448206661696c65640000000000000000000000000060448201526064015b60405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff167f7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65846040516105dc91815260200190565b60405180910390a2505050565b60006105f36111a3565b905090565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716845290915281205490919061048a908290869061063f908790611a93565b610ae8565b73ffffffffffffffffffffffffffffffffffffffff8116600090815260076020526040812054610490565b6000606080828080836106a37f000000000000000000000000000000000000000000000000000000000000000060056112db565b6106ce7f000000000000000000000000000000000000000000000000000000000000000060066112db565b604080516000808252602082019092527f0f000000000000000000000000000000000000000000000000000000000000009b939a50919850469750309650945092509050565b6060600480546103f990611a46565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff87168452909152812054909190838110156107e7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f000000000000000000000000000000000000000000000000000000606482015260840161058b565b6104af8286868403610ae8565b60003361048a818585610d73565b8342111561086c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f45524332305065726d69743a206578706972656420646561646c696e65000000604482015260640161058b565b60007f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c988888861089b8c61137f565b60408051602081019690965273ffffffffffffffffffffffffffffffffffffffff94851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090506000610903826113b4565b90506000610913828787876113fc565b90508973ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146109aa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f45524332305065726d69743a20696e76616c6964207369676e61747572650000604482015260640161058b565b6109b58a8a8a610ae8565b50505050505050505050565b60006020835110156109dd576109d683611424565b9050610490565b816109e88482611b4a565b5060ff9050610490565b90565b73ffffffffffffffffffffffffffffffffffffffff8216610a72576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015260640161058b565b8060026000828254610a849190611a93565b909155505073ffffffffffffffffffffffffffffffffffffffff8216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b73ffffffffffffffffffffffffffffffffffffffff8316610b8a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f7265737300000000000000000000000000000000000000000000000000000000606482015260840161058b565b73ffffffffffffffffffffffffffffffffffffffff8216610c2d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f7373000000000000000000000000000000000000000000000000000000000000606482015260840161058b565b73ffffffffffffffffffffffffffffffffffffffff83811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610d6d5781811015610d60576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e6365000000604482015260640161058b565b610d6d8484848403610ae8565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610e16576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f6472657373000000000000000000000000000000000000000000000000000000606482015260840161058b565b73ffffffffffffffffffffffffffffffffffffffff8216610eb9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f6573730000000000000000000000000000000000000000000000000000000000606482015260840161058b565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020819052604090205481811015610f6f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e63650000000000000000000000000000000000000000000000000000606482015260840161058b565b73ffffffffffffffffffffffffffffffffffffffff848116600081815260208181526040808320878703905593871680835291849020805487019055925185815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a3610d6d565b73ffffffffffffffffffffffffffffffffffffffff8216611085576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360448201527f7300000000000000000000000000000000000000000000000000000000000000606482015260840161058b565b73ffffffffffffffffffffffffffffffffffffffff82166000908152602081905260409020548181101561113b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60448201527f6365000000000000000000000000000000000000000000000000000000000000606482015260840161058b565b73ffffffffffffffffffffffffffffffffffffffff83166000818152602081815260408083208686039055600280548790039055518581529192917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9101610c8f565b505050565b60003073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614801561120957507f000000000000000000000000000000000000000000000000000000000000000046145b1561123357507f000000000000000000000000000000000000000000000000000000000000000090565b6105f3604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201527f0000000000000000000000000000000000000000000000000000000000000000918101919091527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a082015260009060c00160405160208183030381529060405280519060200120905090565b606060ff83146112ee576109d68361147b565b8180546112fa90611a46565b80601f016020809104026020016040519081016040528092919081815260200182805461132690611a46565b80156113735780601f1061134857610100808354040283529160200191611373565b820191906000526020600020905b81548152906001019060200180831161135657829003601f168201915b50505050509050610490565b73ffffffffffffffffffffffffffffffffffffffff811660009081526007602052604090208054600181018255905b50919050565b60006104906113c16111a3565b836040517f19010000000000000000000000000000000000000000000000000000000000008152600281019290925260228201526042902090565b600080600061140d878787876114ba565b9150915061141a816115a9565b5095945050505050565b600080829050601f8151111561146857826040517f305a27a900000000000000000000000000000000000000000000000000000000815260040161058b9190611804565b805161147382611c64565b179392505050565b606060006114888361175f565b604080516020808252818301909252919250600091906020820181803683375050509182525060208101929092525090565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156114f157506000905060036115a0565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015611545573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff8116611599576000600192509250506115a0565b9150600090505b94509492505050565b60008160048111156115bd576115bd611ca6565b036115c55750565b60018160048111156115d9576115d9611ca6565b03611640576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604482015260640161058b565b600281600481111561165457611654611ca6565b036116bb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604482015260640161058b565b60038160048111156116cf576116cf611ca6565b0361175c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f7565000000000000000000000000000000000000000000000000000000000000606482015260840161058b565b50565b600060ff8216601f811115610490576040517fb3512b0c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000815180845260005b818110156117c6576020818501810151868301820152016117aa565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b60208152600061181760208301846117a0565b9392505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461184257600080fd5b919050565b6000806040838503121561185a57600080fd5b6118638361181e565b946020939093013593505050565b60008060006060848603121561188657600080fd5b61188f8461181e565b925061189d6020850161181e565b9150604084013590509250925092565b6000602082840312156118bf57600080fd5b5035919050565b6000602082840312156118d857600080fd5b6118178261181e565b7fff00000000000000000000000000000000000000000000000000000000000000881681526000602060e08184015261191d60e084018a6117a0565b838103604085015261192f818a6117a0565b6060850189905273ffffffffffffffffffffffffffffffffffffffff8816608086015260a0850187905284810360c0860152855180825283870192509083019060005b8181101561198e57835183529284019291840191600101611972565b50909c9b505050505050505050505050565b600080600080600080600060e0888a0312156119bb57600080fd5b6119c48861181e565b96506119d26020890161181e565b95506040880135945060608801359350608088013560ff811681146119f657600080fd5b9699959850939692959460a0840135945060c09093013592915050565b60008060408385031215611a2657600080fd5b611a2f8361181e565b9150611a3d6020840161181e565b90509250929050565b600181811c90821680611a5a57607f821691505b6020821081036113ae577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b80820180821115610490577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b601f82111561119e57600081815260208120601f850160051c81016020861015611b235750805b601f850160051c820191505b81811015611b4257828155600101611b2f565b505050505050565b815167ffffffffffffffff811115611b6457611b64611acd565b611b7881611b728454611a46565b84611afc565b602080601f831160018114611bcb5760008415611b955750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555611b42565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b82811015611c1857888601518255948401946001909101908401611bf9565b5085821015611c5457878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b805160208083015191908110156113ae577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60209190910360031b1b16919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fdfea164736f6c6343000810000a"

func init() {
	if err := json.Unmarshal([]byte(WrappedEtherStorageLayoutJSON), WrappedEtherStorageLayout); err != nil {
		panic(err)
	}

	layouts["WrappedEther"] = WrappedEtherStorageLayout
	deployedBytecodes["WrappedEther"] = WrappedEtherDeployedBin
}