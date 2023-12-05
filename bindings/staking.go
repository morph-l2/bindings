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

// StakingMetaData contains all meta data concerning the Staking contract.
var StakingMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"claimETH\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_sequencerContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_sequencersSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_lock\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"limit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"register\",\"inputs\":[{\"name\":\"tmKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blsKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_minGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"sequencerContract\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sequencers\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sequencersSize\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakeETH\",\"inputs\":[{\"name\":\"_minGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"stakers\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakersNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakings\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tmKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blsKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawETH\",\"inputs\":[{\"name\":\"_minGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdrawals\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"unlock\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"exit\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Claimed\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Registered\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"tmKey\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"blsKey\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"balance\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SequencerUpdated\",\"inputs\":[{\"name\":\"sequencers\",\"type\":\"bytes[]\",\"indexed\":false,\"internalType\":\"bytes[]\"},{\"name\":\"version\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Staked\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawed\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x60806040526006805460ff19169055600060075534801561001f57600080fd5b506128648061002f6000396000f3fe6080604052600436106100dd5760003560e01c8063b21452801161007f578063dc6e13e111610059578063dc6e13e114610278578063f83d08ba146102a8578063fc15799d146102be578063fd5e6dd1146102d157600080fd5b8063b21452801461023a578063cbd679cb14610250578063d702d8aa1461026557600080fd5b80637a9262a2116100bb5780637a9262a21461014f578063a16e0674146101ab578063a4d66daf146101be578063a7044836146101e257600080fd5b80634ec81af1146100e257806367272999146101045780636ba7ccff14610119575b600080fd5b3480156100ee57600080fd5b506101026100fd366004611fc2565b6102f1565b005b34801561011057600080fd5b50610102610675565b34801561012557600080fd5b50610139610134366004611ffb565b6107ed565b6040516101469190612078565b60405180910390f35b34801561015b57600080fd5b5061018e61016a366004612092565b60056020526000908152604090208054600182015460029092015490919060ff1683565b604080519384526020840192909252151590820152606001610146565b6101026101b93660046120c1565b610899565b3480156101ca57600080fd5b506101d460015481565b604051908152602001610146565b3480156101ee57600080fd5b506000546102159062010000900473ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610146565b34801561024657600080fd5b506101d460075481565b34801561025c57600080fd5b506003546101d4565b6101026102733660046120c1565b610cde565b34801561028457600080fd5b50610298610293366004612092565b61113e565b60405161014694939291906120dc565b3480156102b457600080fd5b506101d460025481565b6101026102cc366004612151565b611206565b3480156102dd57600080fd5b506102156102ec366004611ffb565b611945565b600054610100900460ff16158080156103115750600054600160ff909116105b8061032b5750303b15801561032b575060005460ff166001145b6103bc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561041a57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b73ffffffffffffffffffffffffffffffffffffffff8516610497576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f696e76616c69642073657175656e63657220636f6e747261637400000000000060448201526064016103b3565b60008411610527576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f73657175656e6365727353697a65206d7573742067726561746572207468616e60448201527f203000000000000000000000000000000000000000000000000000000000000060648201526084016103b3565b600083116105b7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f7374616b696e67206c696d6974206d7573742067726561746572207468616e2060448201527f300000000000000000000000000000000000000000000000000000000000000060648201526084016103b3565b600080547fffffffffffffffffffff0000000000000000000000000000000000000000ffff166201000073ffffffffffffffffffffffffffffffffffffffff881602179055600784905560018390556002829055801561066e57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b3360009081526005602052604090206002015460ff1680156106a557503360009081526005602052604090205415155b80156106c257503360009081526005602052604090206001015443115b610728576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f696e76616c6964207769746864726177616c000000000000000000000000000060448201526064016103b3565b3360008181526005602052604080822054905181156108fc0292818181858888f1935050505015801561075f573d6000803e3d6000fd5b5033600081815260056020908152604091829020548251938452908301527fd8138f8a3f377c5259ca548e70e4c2de94f129f5a11036a15b69513cba2b426a910160405180910390a1336000908152600560205260408120818155600181019190915560020180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b600881815481106107fd57600080fd5b9060005260206000200160009150905080546108189061223a565b80601f01602080910402602001604051908101604052809291908181526020018280546108449061223a565b80156108915780601f1061086657610100808354040283529160200191610891565b820191906000526020600020905b81548152906001019060200180831161087457829003601f168201915b505050505081565b3360009081526005602052604090206002015460ff1615610916576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f7374616b6572206973206578697465640000000000000000000000000000000060448201526064016103b3565b600080805b600354811015610994573373ffffffffffffffffffffffffffffffffffffffff166003828154811061094f5761094f61228d565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16036109825760019250809150610994565b8061098c816122eb565b91505061091b565b50816109fc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f6f6e207374616b65722063616e2077697468647261770000000000000000000060448201526064016103b3565b60408051606081018252336000908152600460209081529290206003015481526002549091820190610a2e9043612323565b815260016020918201819052336000818152600584526040908190208551808255868601519482019490945594810151600290950180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016951515959095179094558351908152918201527f6cca423c6ffc06e62a0acc433965e074b11c28479b0449250ce3ff65ac9e39fe910160405180910390a1805b600354610ad69060019061233c565b811015610b93576003610aea826001612323565b81548110610afa57610afa61228d565b6000918252602090912001546003805473ffffffffffffffffffffffffffffffffffffffff9092169183908110610b3357610b3361228d565b600091825260209091200180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9290921691909117905580610b8b816122eb565b915050610ac7565b506003805480610ba557610ba561234f565b6000828152602080822083017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff00000000000000000000000000000000000000009081169091559301909355338152600490925260408220805490911681556001810182905590610c2b6002830182611f0b565b506000600391820181905590549003610cc557600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638456cb596040518163ffffffff1660e01b8152600401600060405180830381600087803b158015610ca857600080fd5b505af1158015610cbc573d6000803e3d6000fd5b50505050505050565b600754811015610cd857610cd88361197c565b50505b50565b6000805b600354811015610d58573373ffffffffffffffffffffffffffffffffffffffff1660038281548110610d1657610d1661228d565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1603610d465760019150610d58565b80610d50816122eb565b915050610ce2565b5080610dc0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f7374616b6572206e6f742065786973740000000000000000000000000000000060448201526064016103b3565b600034118015610dee575060015433600090815260046020526040902060030154610dec903490612323565b115b610e54576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f7374616b696e672076616c7565206e6f7420656e6f756768000000000000000060448201526064016103b3565b3360009081526004602052604081206003018054349290610e76908490612323565b909155505033600081815260046020908152604091829020600301548251938452908301527f9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d910160405180910390a1600354600090610ed89060019061233c565b90505b801561110e57600460006003610ef260018561233c565b81548110610f0257610f0261228d565b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600301546004600060038481548110610f8257610f8261228d565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff16835282019290925260400190206003015411156110fc5760006003610fcc60018461233c565b81548110610fdc57610fdc61228d565b6000918252602090912001546003805473ffffffffffffffffffffffffffffffffffffffff909216925090839081106110175761101761228d565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16600361104660018561233c565b815481106110565761105661228d565b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600383815481106110b2576110b261228d565b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505b806111068161237e565b915050610edb565b60065460ff1680156111235750600754600354115b8015611130575060075481105b15610cd857610cd88361197c565b60046020526000908152604090208054600182015460028301805473ffffffffffffffffffffffffffffffffffffffff90931693919261117d9061223a565b80601f01602080910402602001604051908101604052809291908181526020018280546111a99061223a565b80156111f65780601f106111cb576101008083540402835291602001916111f6565b820191906000526020600020905b8154815290600101906020018083116111d957829003601f168201915b5050505050908060030154905084565b6000805b600354811015611280573373ffffffffffffffffffffffffffffffffffffffff166003828154811061123e5761123e61228d565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff160361126e5760019150611280565b80611278816122eb565b91505061120a565b5080156112e9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f616c72656164792072656769737465726564000000000000000000000000000060448201526064016103b3565b3360009081526005602052604090206002015460ff1615611366576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f7374616b6572206973206578697465640000000000000000000000000000000060448201526064016103b3565b6000600754116113f8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f73657175656e6365727353697a65206d7573742067726561746572207468616e60448201527f203000000000000000000000000000000000000000000000000000000000000060648201526084016103b3565b6000849003611463576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f696e76616c69642074656e6465726d696e74207075626b65790000000000000060448201526064016103b3565b8251610100146114cf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f696e76616c696420626c73207075626b6579000000000000000000000000000060448201526064016103b3565b600154341161153a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f7374616b696e672076616c7565206973206e6f7420656e6f756768000000000060448201526064016103b3565b604080516080810182523380825260208083018881528385018881523460608601526000938452600490925293909120825181547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909116178155925160018401555190919060028201906115c59082612401565b506060919091015160039182015580546001810182556000919091527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0180547fffffffffffffffffffffffff000000000000000000000000000000000000000016339081179091556040517fb7f230b53b0f914ccf820ab0618ac8320e984eeb0fb6a740785cf7fdc3b5cee391611662918790879034906120dc565b60405180910390a160035460009061167c9060019061233c565b90505b80156118bb5760046000600361169660018561233c565b815481106116a6576116a661228d565b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206003015460046000600384815481106117265761172661228d565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff16835282019290925260400190206003015411156118a4576000600361177060018461233c565b815481106117805761178061228d565b6000918252602090912001546003805473ffffffffffffffffffffffffffffffffffffffff909216925090839081106117bb576117bb61228d565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1660036117ea60018561233c565b815481106117fa576117fa61228d565b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600383815481106118565761185661228d565b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506118a9565b6118bb565b806118b38161237e565b91505061167f565b60065460ff161580156118d15750600754600354145b1561191057600680547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905561190a8361197c565b5061193f565b60065460ff1680156119315750600754600354111580611931575060075481105b1561066e5761066e8361197c565b50505050565b6003818154811061195557600080fd5b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16905081565b61198860086000611f45565b60075460035481111561199a57506003545b60008167ffffffffffffffff8111156119b5576119b5612122565b604051908082528060200260200182016040528015611a0257816020015b604080516060808201835260008083526020830152918101919091528152602001906001900390816119d35790505b50905060005b82811015611c6e5760086004600060038481548110611a2957611a2961228d565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff1683528281019390935260409091018120835460018101855593825291902090910190611a7c906002018261251b565b5060405180606001604052806004600060038581548110611a9f57611a9f61228d565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff90811684528382019490945260409092018120549092168352600380549390910192600492919086908110611af957611af961228d565b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001015481526020016004600060038581548110611b7e57611b7e61228d565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff16835282019290925260400190206002018054611bbd9061223a565b80601f0160208091040260200160405190810160405280929190818152602001828054611be99061223a565b8015611c365780601f10611c0b57610100808354040283529160200191611c36565b820191906000526020600020905b815481529060010190602001808311611c1957829003601f168201915b5050505050815250828281518110611c5057611c5061228d565b60200260200101819052508080611c66906122eb565b915050611a08565b50600063ad01732f60e01b600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166373452a926040518163ffffffff1660e01b8152600401602060405180830381865afa158015611ce6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d0a9190612648565b611d15906001612323565b83604051602401611d27929190612661565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009094169390931790925260005491517f3135cb9a00000000000000000000000000000000000000000000000000000000815290925073ffffffffffffffffffffffffffffffffffffffff620100009092049190911690633135cb9a90611e0990849060089089906004016127f9565b600060405180830381600087803b158015611e2357600080fd5b505af1158015611e37573d6000803e3d6000fd5b505050507ffccf12f2113375dab7436466a54baaa4c4e5eb61ea2b6964e6716dec4cb8c1986008600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166373452a926040518163ffffffff1660e01b8152600401602060405180830381865afa158015611ecb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611eef9190612648565b604051611efd929190612835565b60405180910390a150505050565b508054611f179061223a565b6000825580601f10611f27575050565b601f016020900490600052602060002090810190610cdb9190611f63565b5080546000825590600052602060002090810190610cdb9190611f7c565b5b80821115611f785760008155600101611f64565b5090565b80821115611f78576000611f908282611f0b565b50600101611f7c565b803573ffffffffffffffffffffffffffffffffffffffff81168114611fbd57600080fd5b919050565b60008060008060808587031215611fd857600080fd5b611fe185611f99565b966020860135965060408601359560600135945092505050565b60006020828403121561200d57600080fd5b5035919050565b6000815180845260005b8181101561203a5760208185018101518683018201520161201e565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b60208152600061208b6020830184612014565b9392505050565b6000602082840312156120a457600080fd5b61208b82611f99565b803563ffffffff81168114611fbd57600080fd5b6000602082840312156120d357600080fd5b61208b826120ad565b73ffffffffffffffffffffffffffffffffffffffff851681528360208201526080604082015260006121116080830185612014565b905082606083015295945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008060006060848603121561216657600080fd5b83359250602084013567ffffffffffffffff8082111561218557600080fd5b818601915086601f83011261219957600080fd5b8135818111156121ab576121ab612122565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f011681019083821181831017156121f1576121f1612122565b8160405282815289602084870101111561220a57600080fd5b826020860160208301376000602084830101528096505050505050612231604085016120ad565b90509250925092565b600181811c9082168061224e57607f821691505b602082108103612287577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361231c5761231c6122bc565b5060010190565b80820180821115612336576123366122bc565b92915050565b81810381811115612336576123366122bc565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b60008161238d5761238d6122bc565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b601f821115610cd857600081815260208120601f850160051c810160208610156123da5750805b601f850160051c820191505b818110156123f9578281556001016123e6565b505050505050565b815167ffffffffffffffff81111561241b5761241b612122565b61242f81612429845461223a565b846123b3565b602080601f831160018114612482576000841561244c5750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b1785556123f9565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b828110156124cf578886015182559484019460019091019084016124b0565b508582101561250b57878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b818103612526575050565b612530825461223a565b67ffffffffffffffff81111561254857612548612122565b61255681612429845461223a565b6000601f8211600181146125a857600083156125725750848201545b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b17845561066e565b6000858152602090207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0841690600086815260209020845b8381101561260057828601548255600195860195909101906020016125e0565b508583101561250b579301547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f8600387901b161c19169092555050600190811b01905550565b60006020828403121561265a57600080fd5b5051919050565b6000604080830185845260208281860152818651808452606093508387019150838160051b88010183890160005b8381101561270a578983037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa00185528151805173ffffffffffffffffffffffffffffffffffffffff16845286810151878501528801518884018890526126f788850182612014565b958701959350509085019060010161268f565b50909a9950505050505050505050565b600081548084526020808501808196506005915083821b81016000878152848120815b878110156127ea578484038b528282546127568161223a565b8087526001828116801561277157600181146127a8576127d3565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0084168c8a01528b8315158c1b8a010194506127d3565b8688528b8820885b848110156127cb5781548b82018f0152908301908d016127b0565b8a018d019550505b509d8a019d9296505050919091019060010161273d565b50919998505050505050505050565b60608152600061280c6060830186612014565b828103602084015261281e818661271a565b91505063ffffffff83166040830152949350505050565b604081526000612848604083018561271a565b9050826020830152939250505056fea164736f6c6343000810000a",
}

// StakingABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingMetaData.ABI instead.
var StakingABI = StakingMetaData.ABI

// StakingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StakingMetaData.Bin instead.
var StakingBin = StakingMetaData.Bin

// DeployStaking deploys a new Ethereum contract, binding an instance of Staking to it.
func DeployStaking(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Staking, error) {
	parsed, err := StakingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StakingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Staking{StakingCaller: StakingCaller{contract: contract}, StakingTransactor: StakingTransactor{contract: contract}, StakingFilterer: StakingFilterer{contract: contract}}, nil
}

// Staking is an auto generated Go binding around an Ethereum contract.
type Staking struct {
	StakingCaller     // Read-only binding to the contract
	StakingTransactor // Write-only binding to the contract
	StakingFilterer   // Log filterer for contract events
}

// StakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingSession struct {
	Contract     *Staking          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingCallerSession struct {
	Contract *StakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingTransactorSession struct {
	Contract     *StakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingRaw struct {
	Contract *Staking // Generic contract binding to access the raw methods on
}

// StakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingCallerRaw struct {
	Contract *StakingCaller // Generic read-only contract binding to access the raw methods on
}

// StakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingTransactorRaw struct {
	Contract *StakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStaking creates a new instance of Staking, bound to a specific deployed contract.
func NewStaking(address common.Address, backend bind.ContractBackend) (*Staking, error) {
	contract, err := bindStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Staking{StakingCaller: StakingCaller{contract: contract}, StakingTransactor: StakingTransactor{contract: contract}, StakingFilterer: StakingFilterer{contract: contract}}, nil
}

// NewStakingCaller creates a new read-only instance of Staking, bound to a specific deployed contract.
func NewStakingCaller(address common.Address, caller bind.ContractCaller) (*StakingCaller, error) {
	contract, err := bindStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingCaller{contract: contract}, nil
}

// NewStakingTransactor creates a new write-only instance of Staking, bound to a specific deployed contract.
func NewStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingTransactor, error) {
	contract, err := bindStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingTransactor{contract: contract}, nil
}

// NewStakingFilterer creates a new log filterer instance of Staking, bound to a specific deployed contract.
func NewStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingFilterer, error) {
	contract, err := bindStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingFilterer{contract: contract}, nil
}

// bindStaking binds a generic wrapper to an already deployed contract.
func bindStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.StakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transact(opts, method, params...)
}

// Limit is a free data retrieval call binding the contract method 0xa4d66daf.
//
// Solidity: function limit() view returns(uint256)
func (_Staking *StakingCaller) Limit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "limit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Limit is a free data retrieval call binding the contract method 0xa4d66daf.
//
// Solidity: function limit() view returns(uint256)
func (_Staking *StakingSession) Limit() (*big.Int, error) {
	return _Staking.Contract.Limit(&_Staking.CallOpts)
}

// Limit is a free data retrieval call binding the contract method 0xa4d66daf.
//
// Solidity: function limit() view returns(uint256)
func (_Staking *StakingCallerSession) Limit() (*big.Int, error) {
	return _Staking.Contract.Limit(&_Staking.CallOpts)
}

// Lock is a free data retrieval call binding the contract method 0xf83d08ba.
//
// Solidity: function lock() view returns(uint256)
func (_Staking *StakingCaller) Lock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "lock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Lock is a free data retrieval call binding the contract method 0xf83d08ba.
//
// Solidity: function lock() view returns(uint256)
func (_Staking *StakingSession) Lock() (*big.Int, error) {
	return _Staking.Contract.Lock(&_Staking.CallOpts)
}

// Lock is a free data retrieval call binding the contract method 0xf83d08ba.
//
// Solidity: function lock() view returns(uint256)
func (_Staking *StakingCallerSession) Lock() (*big.Int, error) {
	return _Staking.Contract.Lock(&_Staking.CallOpts)
}

// SequencerContract is a free data retrieval call binding the contract method 0xa7044836.
//
// Solidity: function sequencerContract() view returns(address)
func (_Staking *StakingCaller) SequencerContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "sequencerContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SequencerContract is a free data retrieval call binding the contract method 0xa7044836.
//
// Solidity: function sequencerContract() view returns(address)
func (_Staking *StakingSession) SequencerContract() (common.Address, error) {
	return _Staking.Contract.SequencerContract(&_Staking.CallOpts)
}

// SequencerContract is a free data retrieval call binding the contract method 0xa7044836.
//
// Solidity: function sequencerContract() view returns(address)
func (_Staking *StakingCallerSession) SequencerContract() (common.Address, error) {
	return _Staking.Contract.SequencerContract(&_Staking.CallOpts)
}

// Sequencers is a free data retrieval call binding the contract method 0x6ba7ccff.
//
// Solidity: function sequencers(uint256 ) view returns(bytes)
func (_Staking *StakingCaller) Sequencers(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "sequencers", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Sequencers is a free data retrieval call binding the contract method 0x6ba7ccff.
//
// Solidity: function sequencers(uint256 ) view returns(bytes)
func (_Staking *StakingSession) Sequencers(arg0 *big.Int) ([]byte, error) {
	return _Staking.Contract.Sequencers(&_Staking.CallOpts, arg0)
}

// Sequencers is a free data retrieval call binding the contract method 0x6ba7ccff.
//
// Solidity: function sequencers(uint256 ) view returns(bytes)
func (_Staking *StakingCallerSession) Sequencers(arg0 *big.Int) ([]byte, error) {
	return _Staking.Contract.Sequencers(&_Staking.CallOpts, arg0)
}

// SequencersSize is a free data retrieval call binding the contract method 0xb2145280.
//
// Solidity: function sequencersSize() view returns(uint256)
func (_Staking *StakingCaller) SequencersSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "sequencersSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SequencersSize is a free data retrieval call binding the contract method 0xb2145280.
//
// Solidity: function sequencersSize() view returns(uint256)
func (_Staking *StakingSession) SequencersSize() (*big.Int, error) {
	return _Staking.Contract.SequencersSize(&_Staking.CallOpts)
}

// SequencersSize is a free data retrieval call binding the contract method 0xb2145280.
//
// Solidity: function sequencersSize() view returns(uint256)
func (_Staking *StakingCallerSession) SequencersSize() (*big.Int, error) {
	return _Staking.Contract.SequencersSize(&_Staking.CallOpts)
}

// Stakers is a free data retrieval call binding the contract method 0xfd5e6dd1.
//
// Solidity: function stakers(uint256 ) view returns(address)
func (_Staking *StakingCaller) Stakers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "stakers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Stakers is a free data retrieval call binding the contract method 0xfd5e6dd1.
//
// Solidity: function stakers(uint256 ) view returns(address)
func (_Staking *StakingSession) Stakers(arg0 *big.Int) (common.Address, error) {
	return _Staking.Contract.Stakers(&_Staking.CallOpts, arg0)
}

// Stakers is a free data retrieval call binding the contract method 0xfd5e6dd1.
//
// Solidity: function stakers(uint256 ) view returns(address)
func (_Staking *StakingCallerSession) Stakers(arg0 *big.Int) (common.Address, error) {
	return _Staking.Contract.Stakers(&_Staking.CallOpts, arg0)
}

// StakersNumber is a free data retrieval call binding the contract method 0xcbd679cb.
//
// Solidity: function stakersNumber() view returns(uint256)
func (_Staking *StakingCaller) StakersNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "stakersNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakersNumber is a free data retrieval call binding the contract method 0xcbd679cb.
//
// Solidity: function stakersNumber() view returns(uint256)
func (_Staking *StakingSession) StakersNumber() (*big.Int, error) {
	return _Staking.Contract.StakersNumber(&_Staking.CallOpts)
}

// StakersNumber is a free data retrieval call binding the contract method 0xcbd679cb.
//
// Solidity: function stakersNumber() view returns(uint256)
func (_Staking *StakingCallerSession) StakersNumber() (*big.Int, error) {
	return _Staking.Contract.StakersNumber(&_Staking.CallOpts)
}

// Stakings is a free data retrieval call binding the contract method 0xdc6e13e1.
//
// Solidity: function stakings(address ) view returns(address addr, bytes32 tmKey, bytes blsKey, uint256 balance)
func (_Staking *StakingCaller) Stakings(opts *bind.CallOpts, arg0 common.Address) (struct {
	Addr    common.Address
	TmKey   [32]byte
	BlsKey  []byte
	Balance *big.Int
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "stakings", arg0)

	outstruct := new(struct {
		Addr    common.Address
		TmKey   [32]byte
		BlsKey  []byte
		Balance *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Addr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TmKey = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.BlsKey = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.Balance = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Stakings is a free data retrieval call binding the contract method 0xdc6e13e1.
//
// Solidity: function stakings(address ) view returns(address addr, bytes32 tmKey, bytes blsKey, uint256 balance)
func (_Staking *StakingSession) Stakings(arg0 common.Address) (struct {
	Addr    common.Address
	TmKey   [32]byte
	BlsKey  []byte
	Balance *big.Int
}, error) {
	return _Staking.Contract.Stakings(&_Staking.CallOpts, arg0)
}

// Stakings is a free data retrieval call binding the contract method 0xdc6e13e1.
//
// Solidity: function stakings(address ) view returns(address addr, bytes32 tmKey, bytes blsKey, uint256 balance)
func (_Staking *StakingCallerSession) Stakings(arg0 common.Address) (struct {
	Addr    common.Address
	TmKey   [32]byte
	BlsKey  []byte
	Balance *big.Int
}, error) {
	return _Staking.Contract.Stakings(&_Staking.CallOpts, arg0)
}

// Withdrawals is a free data retrieval call binding the contract method 0x7a9262a2.
//
// Solidity: function withdrawals(address ) view returns(uint256 balance, uint256 unlock, bool exit)
func (_Staking *StakingCaller) Withdrawals(opts *bind.CallOpts, arg0 common.Address) (struct {
	Balance *big.Int
	Unlock  *big.Int
	Exit    bool
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "withdrawals", arg0)

	outstruct := new(struct {
		Balance *big.Int
		Unlock  *big.Int
		Exit    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Balance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Unlock = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Exit = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// Withdrawals is a free data retrieval call binding the contract method 0x7a9262a2.
//
// Solidity: function withdrawals(address ) view returns(uint256 balance, uint256 unlock, bool exit)
func (_Staking *StakingSession) Withdrawals(arg0 common.Address) (struct {
	Balance *big.Int
	Unlock  *big.Int
	Exit    bool
}, error) {
	return _Staking.Contract.Withdrawals(&_Staking.CallOpts, arg0)
}

// Withdrawals is a free data retrieval call binding the contract method 0x7a9262a2.
//
// Solidity: function withdrawals(address ) view returns(uint256 balance, uint256 unlock, bool exit)
func (_Staking *StakingCallerSession) Withdrawals(arg0 common.Address) (struct {
	Balance *big.Int
	Unlock  *big.Int
	Exit    bool
}, error) {
	return _Staking.Contract.Withdrawals(&_Staking.CallOpts, arg0)
}

// ClaimETH is a paid mutator transaction binding the contract method 0x67272999.
//
// Solidity: function claimETH() returns()
func (_Staking *StakingTransactor) ClaimETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "claimETH")
}

// ClaimETH is a paid mutator transaction binding the contract method 0x67272999.
//
// Solidity: function claimETH() returns()
func (_Staking *StakingSession) ClaimETH() (*types.Transaction, error) {
	return _Staking.Contract.ClaimETH(&_Staking.TransactOpts)
}

// ClaimETH is a paid mutator transaction binding the contract method 0x67272999.
//
// Solidity: function claimETH() returns()
func (_Staking *StakingTransactorSession) ClaimETH() (*types.Transaction, error) {
	return _Staking.Contract.ClaimETH(&_Staking.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x4ec81af1.
//
// Solidity: function initialize(address _sequencerContract, uint256 _sequencersSize, uint256 _limit, uint256 _lock) returns()
func (_Staking *StakingTransactor) Initialize(opts *bind.TransactOpts, _sequencerContract common.Address, _sequencersSize *big.Int, _limit *big.Int, _lock *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "initialize", _sequencerContract, _sequencersSize, _limit, _lock)
}

// Initialize is a paid mutator transaction binding the contract method 0x4ec81af1.
//
// Solidity: function initialize(address _sequencerContract, uint256 _sequencersSize, uint256 _limit, uint256 _lock) returns()
func (_Staking *StakingSession) Initialize(_sequencerContract common.Address, _sequencersSize *big.Int, _limit *big.Int, _lock *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Initialize(&_Staking.TransactOpts, _sequencerContract, _sequencersSize, _limit, _lock)
}

// Initialize is a paid mutator transaction binding the contract method 0x4ec81af1.
//
// Solidity: function initialize(address _sequencerContract, uint256 _sequencersSize, uint256 _limit, uint256 _lock) returns()
func (_Staking *StakingTransactorSession) Initialize(_sequencerContract common.Address, _sequencersSize *big.Int, _limit *big.Int, _lock *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Initialize(&_Staking.TransactOpts, _sequencerContract, _sequencersSize, _limit, _lock)
}

// Register is a paid mutator transaction binding the contract method 0xfc15799d.
//
// Solidity: function register(bytes32 tmKey, bytes blsKey, uint32 _minGasLimit) payable returns()
func (_Staking *StakingTransactor) Register(opts *bind.TransactOpts, tmKey [32]byte, blsKey []byte, _minGasLimit uint32) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "register", tmKey, blsKey, _minGasLimit)
}

// Register is a paid mutator transaction binding the contract method 0xfc15799d.
//
// Solidity: function register(bytes32 tmKey, bytes blsKey, uint32 _minGasLimit) payable returns()
func (_Staking *StakingSession) Register(tmKey [32]byte, blsKey []byte, _minGasLimit uint32) (*types.Transaction, error) {
	return _Staking.Contract.Register(&_Staking.TransactOpts, tmKey, blsKey, _minGasLimit)
}

// Register is a paid mutator transaction binding the contract method 0xfc15799d.
//
// Solidity: function register(bytes32 tmKey, bytes blsKey, uint32 _minGasLimit) payable returns()
func (_Staking *StakingTransactorSession) Register(tmKey [32]byte, blsKey []byte, _minGasLimit uint32) (*types.Transaction, error) {
	return _Staking.Contract.Register(&_Staking.TransactOpts, tmKey, blsKey, _minGasLimit)
}

// StakeETH is a paid mutator transaction binding the contract method 0xd702d8aa.
//
// Solidity: function stakeETH(uint32 _minGasLimit) payable returns()
func (_Staking *StakingTransactor) StakeETH(opts *bind.TransactOpts, _minGasLimit uint32) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "stakeETH", _minGasLimit)
}

// StakeETH is a paid mutator transaction binding the contract method 0xd702d8aa.
//
// Solidity: function stakeETH(uint32 _minGasLimit) payable returns()
func (_Staking *StakingSession) StakeETH(_minGasLimit uint32) (*types.Transaction, error) {
	return _Staking.Contract.StakeETH(&_Staking.TransactOpts, _minGasLimit)
}

// StakeETH is a paid mutator transaction binding the contract method 0xd702d8aa.
//
// Solidity: function stakeETH(uint32 _minGasLimit) payable returns()
func (_Staking *StakingTransactorSession) StakeETH(_minGasLimit uint32) (*types.Transaction, error) {
	return _Staking.Contract.StakeETH(&_Staking.TransactOpts, _minGasLimit)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xa16e0674.
//
// Solidity: function withdrawETH(uint32 _minGasLimit) payable returns()
func (_Staking *StakingTransactor) WithdrawETH(opts *bind.TransactOpts, _minGasLimit uint32) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "withdrawETH", _minGasLimit)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xa16e0674.
//
// Solidity: function withdrawETH(uint32 _minGasLimit) payable returns()
func (_Staking *StakingSession) WithdrawETH(_minGasLimit uint32) (*types.Transaction, error) {
	return _Staking.Contract.WithdrawETH(&_Staking.TransactOpts, _minGasLimit)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xa16e0674.
//
// Solidity: function withdrawETH(uint32 _minGasLimit) payable returns()
func (_Staking *StakingTransactorSession) WithdrawETH(_minGasLimit uint32) (*types.Transaction, error) {
	return _Staking.Contract.WithdrawETH(&_Staking.TransactOpts, _minGasLimit)
}

// StakingClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the Staking contract.
type StakingClaimedIterator struct {
	Event *StakingClaimed // Event containing the contract specifics and raw log

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
func (it *StakingClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingClaimed)
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
		it.Event = new(StakingClaimed)
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
func (it *StakingClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingClaimed represents a Claimed event raised by the Staking contract.
type StakingClaimed struct {
	Addr    common.Address
	Balance *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0xd8138f8a3f377c5259ca548e70e4c2de94f129f5a11036a15b69513cba2b426a.
//
// Solidity: event Claimed(address addr, uint256 balance)
func (_Staking *StakingFilterer) FilterClaimed(opts *bind.FilterOpts) (*StakingClaimedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return &StakingClaimedIterator{contract: _Staking.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0xd8138f8a3f377c5259ca548e70e4c2de94f129f5a11036a15b69513cba2b426a.
//
// Solidity: event Claimed(address addr, uint256 balance)
func (_Staking *StakingFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *StakingClaimed) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingClaimed)
				if err := _Staking.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0xd8138f8a3f377c5259ca548e70e4c2de94f129f5a11036a15b69513cba2b426a.
//
// Solidity: event Claimed(address addr, uint256 balance)
func (_Staking *StakingFilterer) ParseClaimed(log types.Log) (*StakingClaimed, error) {
	event := new(StakingClaimed)
	if err := _Staking.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Staking contract.
type StakingInitializedIterator struct {
	Event *StakingInitialized // Event containing the contract specifics and raw log

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
func (it *StakingInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingInitialized)
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
		it.Event = new(StakingInitialized)
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
func (it *StakingInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingInitialized represents a Initialized event raised by the Staking contract.
type StakingInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Staking *StakingFilterer) FilterInitialized(opts *bind.FilterOpts) (*StakingInitializedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StakingInitializedIterator{contract: _Staking.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Staking *StakingFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StakingInitialized) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingInitialized)
				if err := _Staking.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Staking *StakingFilterer) ParseInitialized(log types.Log) (*StakingInitialized, error) {
	event := new(StakingInitialized)
	if err := _Staking.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingRegisteredIterator is returned from FilterRegistered and is used to iterate over the raw logs and unpacked data for Registered events raised by the Staking contract.
type StakingRegisteredIterator struct {
	Event *StakingRegistered // Event containing the contract specifics and raw log

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
func (it *StakingRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingRegistered)
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
		it.Event = new(StakingRegistered)
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
func (it *StakingRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingRegistered represents a Registered event raised by the Staking contract.
type StakingRegistered struct {
	Addr    common.Address
	TmKey   [32]byte
	BlsKey  []byte
	Balance *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRegistered is a free log retrieval operation binding the contract event 0xb7f230b53b0f914ccf820ab0618ac8320e984eeb0fb6a740785cf7fdc3b5cee3.
//
// Solidity: event Registered(address addr, bytes32 tmKey, bytes blsKey, uint256 balance)
func (_Staking *StakingFilterer) FilterRegistered(opts *bind.FilterOpts) (*StakingRegisteredIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Registered")
	if err != nil {
		return nil, err
	}
	return &StakingRegisteredIterator{contract: _Staking.contract, event: "Registered", logs: logs, sub: sub}, nil
}

// WatchRegistered is a free log subscription operation binding the contract event 0xb7f230b53b0f914ccf820ab0618ac8320e984eeb0fb6a740785cf7fdc3b5cee3.
//
// Solidity: event Registered(address addr, bytes32 tmKey, bytes blsKey, uint256 balance)
func (_Staking *StakingFilterer) WatchRegistered(opts *bind.WatchOpts, sink chan<- *StakingRegistered) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Registered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingRegistered)
				if err := _Staking.contract.UnpackLog(event, "Registered", log); err != nil {
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

// ParseRegistered is a log parse operation binding the contract event 0xb7f230b53b0f914ccf820ab0618ac8320e984eeb0fb6a740785cf7fdc3b5cee3.
//
// Solidity: event Registered(address addr, bytes32 tmKey, bytes blsKey, uint256 balance)
func (_Staking *StakingFilterer) ParseRegistered(log types.Log) (*StakingRegistered, error) {
	event := new(StakingRegistered)
	if err := _Staking.contract.UnpackLog(event, "Registered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingSequencerUpdatedIterator is returned from FilterSequencerUpdated and is used to iterate over the raw logs and unpacked data for SequencerUpdated events raised by the Staking contract.
type StakingSequencerUpdatedIterator struct {
	Event *StakingSequencerUpdated // Event containing the contract specifics and raw log

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
func (it *StakingSequencerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingSequencerUpdated)
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
		it.Event = new(StakingSequencerUpdated)
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
func (it *StakingSequencerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingSequencerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingSequencerUpdated represents a SequencerUpdated event raised by the Staking contract.
type StakingSequencerUpdated struct {
	Sequencers [][]byte
	Version    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSequencerUpdated is a free log retrieval operation binding the contract event 0xfccf12f2113375dab7436466a54baaa4c4e5eb61ea2b6964e6716dec4cb8c198.
//
// Solidity: event SequencerUpdated(bytes[] sequencers, uint256 version)
func (_Staking *StakingFilterer) FilterSequencerUpdated(opts *bind.FilterOpts) (*StakingSequencerUpdatedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "SequencerUpdated")
	if err != nil {
		return nil, err
	}
	return &StakingSequencerUpdatedIterator{contract: _Staking.contract, event: "SequencerUpdated", logs: logs, sub: sub}, nil
}

// WatchSequencerUpdated is a free log subscription operation binding the contract event 0xfccf12f2113375dab7436466a54baaa4c4e5eb61ea2b6964e6716dec4cb8c198.
//
// Solidity: event SequencerUpdated(bytes[] sequencers, uint256 version)
func (_Staking *StakingFilterer) WatchSequencerUpdated(opts *bind.WatchOpts, sink chan<- *StakingSequencerUpdated) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "SequencerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingSequencerUpdated)
				if err := _Staking.contract.UnpackLog(event, "SequencerUpdated", log); err != nil {
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

// ParseSequencerUpdated is a log parse operation binding the contract event 0xfccf12f2113375dab7436466a54baaa4c4e5eb61ea2b6964e6716dec4cb8c198.
//
// Solidity: event SequencerUpdated(bytes[] sequencers, uint256 version)
func (_Staking *StakingFilterer) ParseSequencerUpdated(log types.Log) (*StakingSequencerUpdated, error) {
	event := new(StakingSequencerUpdated)
	if err := _Staking.contract.UnpackLog(event, "SequencerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the Staking contract.
type StakingStakedIterator struct {
	Event *StakingStaked // Event containing the contract specifics and raw log

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
func (it *StakingStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingStaked)
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
		it.Event = new(StakingStaked)
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
func (it *StakingStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingStaked represents a Staked event raised by the Staking contract.
type StakingStaked struct {
	Addr    common.Address
	Balance *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address addr, uint256 balance)
func (_Staking *StakingFilterer) FilterStaked(opts *bind.FilterOpts) (*StakingStakedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Staked")
	if err != nil {
		return nil, err
	}
	return &StakingStakedIterator{contract: _Staking.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address addr, uint256 balance)
func (_Staking *StakingFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *StakingStaked) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Staked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingStaked)
				if err := _Staking.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address addr, uint256 balance)
func (_Staking *StakingFilterer) ParseStaked(log types.Log) (*StakingStaked, error) {
	event := new(StakingStaked)
	if err := _Staking.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingWithdrawedIterator is returned from FilterWithdrawed and is used to iterate over the raw logs and unpacked data for Withdrawed events raised by the Staking contract.
type StakingWithdrawedIterator struct {
	Event *StakingWithdrawed // Event containing the contract specifics and raw log

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
func (it *StakingWithdrawedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWithdrawed)
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
		it.Event = new(StakingWithdrawed)
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
func (it *StakingWithdrawedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWithdrawedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWithdrawed represents a Withdrawed event raised by the Staking contract.
type StakingWithdrawed struct {
	Addr    common.Address
	Balance *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawed is a free log retrieval operation binding the contract event 0x6cca423c6ffc06e62a0acc433965e074b11c28479b0449250ce3ff65ac9e39fe.
//
// Solidity: event Withdrawed(address addr, uint256 balance)
func (_Staking *StakingFilterer) FilterWithdrawed(opts *bind.FilterOpts) (*StakingWithdrawedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Withdrawed")
	if err != nil {
		return nil, err
	}
	return &StakingWithdrawedIterator{contract: _Staking.contract, event: "Withdrawed", logs: logs, sub: sub}, nil
}

// WatchWithdrawed is a free log subscription operation binding the contract event 0x6cca423c6ffc06e62a0acc433965e074b11c28479b0449250ce3ff65ac9e39fe.
//
// Solidity: event Withdrawed(address addr, uint256 balance)
func (_Staking *StakingFilterer) WatchWithdrawed(opts *bind.WatchOpts, sink chan<- *StakingWithdrawed) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Withdrawed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWithdrawed)
				if err := _Staking.contract.UnpackLog(event, "Withdrawed", log); err != nil {
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

// ParseWithdrawed is a log parse operation binding the contract event 0x6cca423c6ffc06e62a0acc433965e074b11c28479b0449250ce3ff65ac9e39fe.
//
// Solidity: event Withdrawed(address addr, uint256 balance)
func (_Staking *StakingFilterer) ParseWithdrawed(log types.Log) (*StakingWithdrawed, error) {
	event := new(StakingWithdrawed)
	if err := _Staking.contract.UnpackLog(event, "Withdrawed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
