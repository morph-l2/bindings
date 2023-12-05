// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/morph-l2/bindings/solc"
)

const SequencerFeeVaultStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L2/SequencerFeeVault.sol:SequencerFeeVault\",\"label\":\"totalProcessed\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint256\"}],\"types\":{\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"}}}"

var SequencerFeeVaultStorageLayout = new(solc.StorageLayout)

var SequencerFeeVaultDeployedBin = "0x6080604052600436106100695760003560e01c806384411d651161004357806384411d651461010c578063d3e5792b14610130578063d4ff92181461016457600080fd5b80630d9019e1146100755780633ccfd60b146100d357806354fd4d50146100ea57600080fd5b3661007057005b600080fd5b34801561008157600080fd5b506100a97f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b3480156100df57600080fd5b506100e8610197565b005b3480156100f657600080fd5b506100ff6103b8565b6040516100ca919061066a565b34801561011857600080fd5b5061012260005481565b6040519081526020016100ca565b34801561013c57600080fd5b506101227f000000000000000000000000000000000000000000000000000000000000000081565b34801561017057600080fd5b507f00000000000000000000000000000000000000000000000000000000000000006100a9565b7f0000000000000000000000000000000000000000000000000000000000000000471015610271576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604a60248201527f4665655661756c743a207769746864726177616c20616d6f756e74206d75737460448201527f2062652067726561746572207468616e206d696e696d756d207769746864726160648201527f77616c20616d6f756e7400000000000000000000000000000000000000000000608482015260a40160405180910390fd5b6000479050806000808282546102879190610684565b9091555050604080518281527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166020820152338183015290517fc8a211cc64b6ed1b50595a9fcb1932b6d1e5a6e8ef15b60e5b1f988ea9086bba9181900360600190a1604080516020810182526000815290517fe11013dd0000000000000000000000000000000000000000000000000000000081527342000000000000000000000000000000000000109163e11013dd918491610383917f0000000000000000000000000000000000000000000000000000000000000000916188b8916004016106be565b6000604051808303818588803b15801561039c57600080fd5b505af11580156103b0573d6000803e3d6000fd5b505050505050565b60606103e37f000000000000000000000000000000000000000000000000000000000000000061045b565b61040c7f000000000000000000000000000000000000000000000000000000000000000061045b565b6104357f000000000000000000000000000000000000000000000000000000000000000061045b565b60405160200161044793929190610702565b604051602081830303815290604052905090565b6060600061046883610519565b600101905060008167ffffffffffffffff81111561048857610488610778565b6040519080825280601f01601f1916602001820160405280156104b2576020820181803683370190505b5090508181016020015b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a85049450846104bc57509392505050565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f0100000000000000008310610562577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef8100000000831061058e576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc1000083106105ac57662386f26fc10000830492506010015b6305f5e10083106105c4576305f5e100830492506008015b61271083106105d857612710830492506004015b606483106105ea576064830492506002015b600a83106105f6576001015b92915050565b60005b838110156106175781810151838201526020016105ff565b50506000910152565b600081518084526106388160208601602086016105fc565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061067d6020830184610620565b9392505050565b808201808211156105f6577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff8416815263ffffffff831660208201526060604082015260006106f96060830184610620565b95945050505050565b600084516107148184602089016105fc565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551610750816001850160208a016105fc565b6001920191820152835161076b8160028401602088016105fc565b0160020195945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fdfea164736f6c6343000810000a"

func init() {
	if err := json.Unmarshal([]byte(SequencerFeeVaultStorageLayoutJSON), SequencerFeeVaultStorageLayout); err != nil {
		panic(err)
	}

	layouts["SequencerFeeVault"] = SequencerFeeVaultStorageLayout
	deployedBytecodes["SequencerFeeVault"] = SequencerFeeVaultDeployedBin
}
