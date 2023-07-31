// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/kroma-network/kroma/bindings/solc"
)

const KromaPortalStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L1/KromaPortal.sol:KromaPortal\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"contracts/L1/KromaPortal.sol:KromaPortal\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"contracts/L1/KromaPortal.sol:KromaPortal\",\"label\":\"params\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_struct(ResourceParams)1010_storage\"},{\"astId\":1003,\"contract\":\"contracts/L1/KromaPortal.sol:KromaPortal\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_array(t_uint256)1008_storage\"},{\"astId\":1004,\"contract\":\"contracts/L1/KromaPortal.sol:KromaPortal\",\"label\":\"l2Sender\",\"offset\":0,\"slot\":\"50\",\"type\":\"t_address\"},{\"astId\":1005,\"contract\":\"contracts/L1/KromaPortal.sol:KromaPortal\",\"label\":\"finalizedWithdrawals\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1006,\"contract\":\"contracts/L1/KromaPortal.sol:KromaPortal\",\"label\":\"provenWithdrawals\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_mapping(t_bytes32,t_struct(ProvenWithdrawal)1009_storage)\"},{\"astId\":1007,\"contract\":\"contracts/L1/KromaPortal.sol:KromaPortal\",\"label\":\"paused\",\"offset\":0,\"slot\":\"53\",\"type\":\"t_bool\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)1008_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[48]\",\"numberOfBytes\":\"1536\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_mapping(t_bytes32,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_bool\"},\"t_mapping(t_bytes32,t_struct(ProvenWithdrawal)1009_storage)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e struct KromaPortal.ProvenWithdrawal)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_struct(ProvenWithdrawal)1009_storage\"},\"t_struct(ProvenWithdrawal)1009_storage\":{\"encoding\":\"inplace\",\"label\":\"struct KromaPortal.ProvenWithdrawal\",\"numberOfBytes\":\"64\"},\"t_struct(ResourceParams)1010_storage\":{\"encoding\":\"inplace\",\"label\":\"struct ResourceMetering.ResourceParams\",\"numberOfBytes\":\"32\"},\"t_uint128\":{\"encoding\":\"inplace\",\"label\":\"uint128\",\"numberOfBytes\":\"16\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint64\":{\"encoding\":\"inplace\",\"label\":\"uint64\",\"numberOfBytes\":\"8\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var KromaPortalStorageLayout = new(solc.StorageLayout)

var KromaPortalDeployedBin = "0x6080604052600436106101475760003560e01c80638c3152e9116100c0578063cff0ab9611610074578063e965084c11610059578063e965084c14610481578063e9e05c421461050d578063f04987501461052057600080fd5b8063cff0ab96146103c0578063d53a822f1461046157600080fd5b8063a14238e7116100a5578063a14238e71461033c578063b98debbf1461036c578063c30af388146103a057600080fd5b80638c3152e9146102ef5780639bf62d821461030f57600080fd5b80635865b607116101175780636dbffb78116100fc5780636dbffb7814610286578063724c184c146102a65780638456cb59146102da57600080fd5b80635865b607146102285780635c975abb1461025c57600080fd5b80621c2ff6146101735780630757b244146101d15780633f4ba83a146101f157806354fd4d501461020657600080fd5b3661016e5761016c3334620186a0600060405180602001604052806000815250610554565b005b600080fd5b34801561017f57600080fd5b506101a77f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b3480156101dd57600080fd5b5061016c6101ec366004612c17565b61073a565b3480156101fd57600080fd5b5061016c610d3e565b34801561021257600080fd5b5061021b610e47565b6040516101c89190612d6e565b34801561023457600080fd5b506101a77f000000000000000000000000000000000000000000000000000000000000000081565b34801561026857600080fd5b506035546102769060ff1681565b60405190151581526020016101c8565b34801561029257600080fd5b506102766102a1366004612d81565b610eea565b3480156102b257600080fd5b506101a77f000000000000000000000000000000000000000000000000000000000000000081565b3480156102e657600080fd5b5061016c610fc1565b3480156102fb57600080fd5b5061016c61030a366004612d9a565b6110c6565b34801561031b57600080fd5b506032546101a79073ffffffffffffffffffffffffffffffffffffffff1681565b34801561034857600080fd5b50610276610357366004612d81565b60336020526000908152604090205460ff1681565b34801561037857600080fd5b506101a77f000000000000000000000000000000000000000000000000000000000000000081565b3480156103ac57600080fd5b5061016c6103bb366004612de7565b611894565b3480156103cc57600080fd5b50600154610428906fffffffffffffffffffffffffffffffff81169067ffffffffffffffff7001000000000000000000000000000000008204811691780100000000000000000000000000000000000000000000000090041683565b604080516fffffffffffffffffffffffffffffffff909416845267ffffffffffffffff92831660208501529116908201526060016101c8565b34801561046d57600080fd5b5061016c61047c366004612e55565b6119f2565b34801561048d57600080fd5b506104df61049c366004612d81565b603460205260009081526040902080546001909101546fffffffffffffffffffffffffffffffff8082169170010000000000000000000000000000000090041683565b604080519384526fffffffffffffffffffffffffffffffff92831660208501529116908201526060016101c8565b61016c61051b366004612e72565b610554565b34801561052c57600080fd5b506101a77f000000000000000000000000000000000000000000000000000000000000000081565b8260005a905083156105ef5773ffffffffffffffffffffffffffffffffffffffff8716156105ef5760405162461bcd60e51b815260206004820152603d60248201527f4b726f6d61506f7274616c3a206d7573742073656e6420746f2061646472657360448201527f73283029207768656e206372656174696e67206120636f6e747261637400000060648201526084015b60405180910390fd5b6152088567ffffffffffffffff1610156106715760405162461bcd60e51b815260206004820152603560248201527f4b726f6d61506f7274616c3a20676173206c696d6974206d75737420636f766560448201527f7220696e737472696e7369632067617320636f7374000000000000000000000060648201526084016105e6565b33328114610692575033731111000000000000000000000000000000001111015b600034888888886040516020016106ad959493929190612eef565b604051602081830303815290604052905060008973ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fb3813568d9991fc951961fcb4c784893574240a28925604d09fc577c55bb7c328460405161071d9190612d6e565b60405180910390a450506107318282611be1565b50505050505050565b60355460ff161561078d5760405162461bcd60e51b815260206004820152601360248201527f4b726f6d61506f7274616c3a207061757365640000000000000000000000000060448201526064016105e6565b3073ffffffffffffffffffffffffffffffffffffffff16856040015173ffffffffffffffffffffffffffffffffffffffff16036108325760405162461bcd60e51b815260206004820152603c60248201527f4b726f6d61506f7274616c3a20796f752063616e6e6f742073656e64206d657360448201527f736167657320746f2074686520706f7274616c20636f6e74726163740000000060648201526084016105e6565b6040517fa25ae557000000000000000000000000000000000000000000000000000000008152600481018590526000907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063a25ae55790602401608060405180830381865afa1580156108c0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108e49190612f74565b6020015190506109016108fc36869003860186612fec565b611ef4565b81146109755760405162461bcd60e51b815260206004820152602660248201527f4b726f6d61506f7274616c3a20696e76616c6964206f757470757420726f6f7460448201527f2070726f6f66000000000000000000000000000000000000000000000000000060648201526084016105e6565b600061098087611f79565b6000818152603460209081526040918290208251606081018452815481526001909101546fffffffffffffffffffffffffffffffff8082169383018490527001000000000000000000000000000000009091041692810192909252919250901580610ab55750805160408083015190517fa25ae5570000000000000000000000000000000000000000000000000000000081526fffffffffffffffffffffffffffffffff90911660048201527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063a25ae55790602401608060405180830381865afa158015610a8a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610aae9190612f74565b6020015114155b610b275760405162461bcd60e51b815260206004820152603460248201527f4b726f6d61506f7274616c3a207769746864726177616c20686173682068617360448201527f20616c7265616479206265656e2070726f76656e00000000000000000000000060648201526084016105e6565b60408051602080820185905260008284015282518083038401815260608301808552815191909201207f12e64a7200000000000000000000000000000000000000000000000000000000909152917f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16916312e64a7291610bc99185918b918b918e0135906064016130a5565b602060405180830381865afa158015610be6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c0a91906131a6565b610c7c5760405162461bcd60e51b815260206004820152602f60248201527f4b726f6d61506f7274616c3a20696e76616c6964207769746864726177616c2060448201527f696e636c7573696f6e2070726f6f66000000000000000000000000000000000060648201526084016105e6565b604080516060810182528581526fffffffffffffffffffffffffffffffff42811660208084019182528c831684860190815260008981526034835286812095518655925190518416700100000000000000000000000000000000029316929092176001909301929092558b830151908c0151925173ffffffffffffffffffffffffffffffffffffffff918216939091169186917f67a6208cfcc0801d50f6cbe764733f4fddf66ac0b04442061a8a8c0cb6b63f629190a4505050505050505050565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610de95760405162461bcd60e51b815260206004820152602660248201527f4b726f6d61506f7274616c3a206f6e6c7920677561726469616e2063616e207560448201527f6e7061757365000000000000000000000000000000000000000000000000000060648201526084016105e6565b603580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690556040513381527f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa906020015b60405180910390a1565b6060610e727f0000000000000000000000000000000000000000000000000000000000000000611fc6565b610e9b7f0000000000000000000000000000000000000000000000000000000000000000611fc6565b610ec47f0000000000000000000000000000000000000000000000000000000000000000611fc6565b604051602001610ed6939291906131c3565b604051602081830303815290604052905090565b6040517fa25ae55700000000000000000000000000000000000000000000000000000000815260048101829052600090610fbb9073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063a25ae55790602401608060405180830381865afa158015610f7c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fa09190612f74565b604001516fffffffffffffffffffffffffffffffff16612084565b92915050565b3373ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161461106b5760405162461bcd60e51b8152602060048201526024808201527f4b726f6d61506f7274616c3a206f6e6c7920677561726469616e2063616e207060448201527f617573650000000000000000000000000000000000000000000000000000000060648201526084016105e6565b603580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556040513381527f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25890602001610e3d565b60355460ff16156111195760405162461bcd60e51b815260206004820152601360248201527f4b726f6d61506f7274616c3a207061757365640000000000000000000000000060448201526064016105e6565b60325473ffffffffffffffffffffffffffffffffffffffff1661dead146111a85760405162461bcd60e51b815260206004820152603c60248201527f4b726f6d61506f7274616c3a2063616e206f6e6c792074726967676572206f6e60448201527f65207769746864726177616c20706572207472616e73616374696f6e0000000060648201526084016105e6565b60006111b382611f79565b60008181526034602090815260408083208151606081018352815481526001909101546fffffffffffffffffffffffffffffffff808216948301859052700100000000000000000000000000000000909104169181019190915292935090036112845760405162461bcd60e51b815260206004820152602f60248201527f4b726f6d61506f7274616c3a207769746864726177616c20686173206e6f742060448201527f6265656e2070726f76656e20796574000000000000000000000000000000000060648201526084016105e6565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663887862726040518163ffffffff1660e01b8152600401602060405180830381865afa1580156112ef573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113139190613239565b81602001516fffffffffffffffffffffffffffffffff1610156113c45760405162461bcd60e51b815260206004820152604860248201527f4b726f6d61506f7274616c3a207769746864726177616c2074696d657374616d60448201527f70206c657373207468616e204c32204f7261636c65207374617274696e67207460648201527f696d657374616d70000000000000000000000000000000000000000000000000608482015260a4016105e6565b6113e381602001516fffffffffffffffffffffffffffffffff16612084565b61147b5760405162461bcd60e51b815260206004820152604260248201527f4b726f6d61506f7274616c3a2070726f76656e207769746864726177616c206660448201527f696e616c697a6174696f6e20706572696f6420686173206e6f7420656c61707360648201527f6564000000000000000000000000000000000000000000000000000000000000608482015260a4016105e6565b60408181015190517fa25ae5570000000000000000000000000000000000000000000000000000000081526fffffffffffffffffffffffffffffffff90911660048201526000907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063a25ae55790602401608060405180830381865afa158015611520573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115449190612f74565b82516020820151919250146115e75760405162461bcd60e51b815260206004820152604660248201527f4b726f6d61506f7274616c3a206f757470757420726f6f742070726f76656e2060448201527f6973206e6f74207468652073616d652061732063757272656e74206f7574707560648201527f7420726f6f740000000000000000000000000000000000000000000000000000608482015260a4016105e6565b61160681604001516fffffffffffffffffffffffffffffffff16612084565b61169e5760405162461bcd60e51b815260206004820152604260248201527f4b726f6d61506f7274616c3a20636865636b706f696e74206f7574707574206660448201527f696e616c697a6174696f6e20706572696f6420686173206e6f7420656c61707360648201527f6564000000000000000000000000000000000000000000000000000000000000608482015260a4016105e6565b60008381526033602052604090205460ff16156117235760405162461bcd60e51b815260206004820152603260248201527f4b726f6d61506f7274616c3a207769746864726177616c2068617320616c726560448201527f616479206265656e2066696e616c697a6564000000000000000000000000000060648201526084016105e6565b600083815260336020908152604080832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055908601516032805473ffffffffffffffffffffffffffffffffffffffff9092167fffffffffffffffffffffffff00000000000000000000000000000000000000009092169190911790558501516080860151606087015160a08801516117c593929190612127565b603280547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead17905560405190915084907fdb5c7652857aa163daadd670e116628fb42e869d8ac4251ef8971d9e5727df1b9061182a90841515815260200190565b60405180910390a2801580156118405750326001145b1561188d5760405162461bcd60e51b815260206004820152601e60248201527f4b726f6d61506f7274616c3a207769746864726177616c206661696c6564000060448201526064016105e6565b5050505050565b3373ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161461193f5760405162461bcd60e51b815260206004820152603f60248201527f4b726f6d61506f7274616c3a2066756e6374696f6e2063616e206f6e6c79206260448201527f652063616c6c65642066726f6d207468652056616c696461746f72506f6f6c0060648201526084016105e6565b6040513373111100000000000000000000000000000000111101906000906119739082908190879082908890602001612eef565b604051602081830303815290604052905060008573ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fb3813568d9991fc951961fcb4c784893574240a28925604d09fc577c55bb7c32846040516119e39190612d6e565b60405180910390a45050505050565b600054610100900460ff1615808015611a125750600054600160ff909116105b80611a2c5750303b158015611a2c575060005460ff166001145b611a9e5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016105e6565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015611afc57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b603280547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead179055603580548315157fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00909116179055611b5e612185565b8015611bc157600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b600154600090611c17907801000000000000000000000000000000000000000000000000900467ffffffffffffffff1643613281565b90506000611c2361224e565b90506000816020015160ff16826000015163ffffffff16611c4491906132c7565b90508215611d7b57600154600090611c7b908390700100000000000000000000000000000000900467ffffffffffffffff1661332f565b90506000836040015160ff1683611c9291906133a3565b600154611cb29084906fffffffffffffffffffffffffffffffff166133a3565b611cbc91906132c7565b600154909150600090611d0d90611ce69084906fffffffffffffffffffffffffffffffff1661345f565b866060015163ffffffff168760a001516fffffffffffffffffffffffffffffffff16612314565b90506001861115611d3c57611d39611ce682876040015160ff1660018a611d349190613281565b612331565b90505b6fffffffffffffffffffffffffffffffff16780100000000000000000000000000000000000000000000000067ffffffffffffffff4316021760015550505b60018054869190601090611dae908490700100000000000000000000000000000000900467ffffffffffffffff166134d3565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550816000015163ffffffff16600160000160109054906101000a900467ffffffffffffffff1667ffffffffffffffff161315611e775760405162461bcd60e51b815260206004820152603e60248201527f5265736f757263654d65746572696e673a2063616e6e6f7420627579206d6f7260448201527f6520676173207468616e20617661696c61626c6520676173206c696d6974000060648201526084016105e6565b600154600090611ea3906fffffffffffffffffffffffffffffffff1667ffffffffffffffff88166134ff565b90506000611eb548633b9aca00612386565b611ebf908361353c565b905060005a611ece9088613281565b905080821115611eea57611eea611ee58284613281565b61239e565b5050505050505050565b8051600090611f0657610fbb826123cc565b60405162461bcd60e51b815260206004820152602a60248201527f48617368696e673a20756e6b6e6f776e206f757470757420726f6f742070726f60448201527f6f662076657273696f6e0000000000000000000000000000000000000000000060648201526084016105e6565b919050565b80516020808301516040808501516060860151608087015160a08801519351600097611fa9979096959101613550565b604051602081830303815290604052805190602001209050919050565b60606000611fd38361240f565b600101905060008167ffffffffffffffff811115611ff357611ff3612a69565b6040519080825280601f01601f19166020018201604052801561201d576020820181803683370190505b5090508181016020015b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a850494508461202757509392505050565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663f4daa2916040518163ffffffff1660e01b8152600401602060405180830381865afa1580156120f1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121159190613239565b61211f90836135a7565b421192915050565b60008060006121378660006124f1565b90508061216d576308c379a06000526020805278185361666543616c6c3a204e6f7420656e6f756768206761736058526064601cfd5b600080855160208701888b5af1979650505050505050565b600054610100900460ff166122025760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016105e6565b60408051606081018252633b9aca00808252600060208301524367ffffffffffffffff169190920181905278010000000000000000000000000000000000000000000000000217600155565b6040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a08101919091527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663cc731b026040518163ffffffff1660e01b815260040160c060405180830381865afa1580156122eb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061230f91906135e4565b905090565b6000612329612323858561250f565b8361251e565b949350505050565b6000670de0b6b3a764000061237261234985836132c7565b61235b90670de0b6b3a764000061332f565b61236d85670de0b6b3a76400006133a3565b61252d565b61237c90866133a3565b61232991906132c7565b60008183116123955781612397565b825b9392505050565b6000805a90505b825a6123b19083613281565b10156123c7576123c082613683565b91506123a5565b505050565b80516020808301516040808501516060808701516080808901518551978801989098529386019490945284015282015260a081019190915260009060c001611fa9565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f0100000000000000008310612458577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef81000000008310612484576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc1000083106124a257662386f26fc10000830492506010015b6305f5e10083106124ba576305f5e100830492506008015b61271083106124ce57612710830492506004015b606483106124e0576064830492506002015b600a8310610fbb5760010192915050565b600080603f83619c4001026040850201603f5a021015949350505050565b60008183136123955781612397565b60008183126123955781612397565b6000612397670de0b6b3a7640000836125458661255e565b61254f91906133a3565b61255991906132c7565b612788565b60008082136125af5760405162461bcd60e51b815260206004820152600960248201527f554e444546494e4544000000000000000000000000000000000000000000000060448201526064016105e6565b600060606125bc846129ad565b03609f8181039490941b90931c6c465772b2bbbb5f824b15207a3081018102606090811d6d0388eaa27412d5aca026815d636e018202811d6d0df99ac502031bf953eff472fdcc018202811d6d13cdffb29d51d99322bdff5f2211018202811d6d0a0f742023def783a307a986912e018202811d6d01920d8043ca89b5239253284e42018202811d6c0b7a86d7375468fac667a0a527016c29508e458543d8aa4df2abee7883018302821d6d0139601a2efabe717e604cbb4894018302821d6d02247f7a7b6594320649aa03aba1018302821d7fffffffffffffffffffffffffffffffffffffff73c0c716a594e00d54e3c4cbc9018302821d7ffffffffffffffffffffffffffffffffffffffdc7b88c420e53a9890533129f6f01830290911d7fffffffffffffffffffffffffffffffffffffff465fda27eb4d63ded474e5f832019091027ffffffffffffffff5f6af8f7b3396644f18e157960000000000000000000000000105711340daa0d5f769dba1915cef59f0815a5506027d0267a36c0c95b3975ab3ee5b203a7614a3f75373f047d803ae7b6687f2b393909302929092017d57115e47018c7177eebf7cd370a3356a1b7863008a5ae8028c72b88642840160ae1d92915050565b60007ffffffffffffffffffffffffffffffffffffffffffffffffdb731c958f34d94c182136127b957506000919050565b680755bf798b4a1bf1e582126128115760405162461bcd60e51b815260206004820152600c60248201527f4558505f4f564552464c4f57000000000000000000000000000000000000000060448201526064016105e6565b6503782dace9d9604e83901b059150600060606bb17217f7d1cf79abc9e3b39884821b056b80000000000000000000000001901d6bb17217f7d1cf79abc9e3b39881029093037fffffffffffffffffffffffffffffffffffffffdbf3ccf1604d263450f02a550481018102606090811d6d0277594991cfc85f6e2461837cd9018202811d7fffffffffffffffffffffffffffffffffffffe5adedaa1cb095af9e4da10e363c018202811d6db1bbb201f443cf962f1a1d3db4a5018202811d7ffffffffffffffffffffffffffffffffffffd38dc772608b0ae56cce01296c0eb018202811d6e05180bb14799ab47a8a8cb2a527d57016d02d16720577bd19bf614176fe9ea6c10fe68e7fd37d0007b713f765084018402831d9081019084017ffffffffffffffffffffffffffffffffffffffe2c69812cf03b0763fd454a8f7e010290911d6e0587f503bb6ea29d25fcb7401964500190910279d835ebba824c98fb31b83b2ca45c000000000000000000000000010574029d9dc38563c32e5c2f6dc192ee70ef65f9978af30260c3939093039290921c92915050565b60008082116129fe5760405162461bcd60e51b815260206004820152600960248201527f554e444546494e4544000000000000000000000000000000000000000000000060448201526064016105e6565b5060016fffffffffffffffffffffffffffffffff821160071b82811c67ffffffffffffffff1060061b1782811c63ffffffff1060051b1782811c61ffff1060041b1782811c60ff10600390811b90911783811c600f1060021b1783811c909110821b1791821c111790565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff81168114612aba57600080fd5b50565b600082601f830112612ace57600080fd5b813567ffffffffffffffff80821115612ae957612ae9612a69565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908282118183101715612b2f57612b2f612a69565b81604052838152866020858801011115612b4857600080fd5b836020870160208301376000602085830101528094505050505092915050565b600060c08284031215612b7a57600080fd5b60405160c0810167ffffffffffffffff8282108183111715612b9e57612b9e612a69565b816040528293508435835260208501359150612bb982612a98565b81602084015260408501359150612bcf82612a98565b816040840152606085013560608401526080850135608084015260a0850135915080821115612bfd57600080fd5b50612c0a85828601612abd565b60a0830152505092915050565b6000806000806000858703610100811215612c3157600080fd5b863567ffffffffffffffff80821115612c4957600080fd5b612c558a838b01612b68565b97506020890135965060a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc084011215612c8e57600080fd5b60408901955060e0890135925080831115612ca857600080fd5b828901925089601f840112612cbc57600080fd5b8235915080821115612ccd57600080fd5b508860208260051b8401011115612ce357600080fd5b959894975092955050506020019190565b60005b83811015612d0f578181015183820152602001612cf7565b83811115612d1e576000848401525b50505050565b60008151808452612d3c816020860160208601612cf4565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006123976020830184612d24565b600060208284031215612d9357600080fd5b5035919050565b600060208284031215612dac57600080fd5b813567ffffffffffffffff811115612dc357600080fd5b61232984828501612b68565b803567ffffffffffffffff81168114611f7457600080fd5b600080600060608486031215612dfc57600080fd5b8335612e0781612a98565b9250612e1560208501612dcf565b9150604084013567ffffffffffffffff811115612e3157600080fd5b612e3d86828701612abd565b9150509250925092565b8015158114612aba57600080fd5b600060208284031215612e6757600080fd5b813561239781612e47565b600080600080600060a08688031215612e8a57600080fd5b8535612e9581612a98565b945060208601359350612eaa60408701612dcf565b92506060860135612eba81612e47565b9150608086013567ffffffffffffffff811115612ed657600080fd5b612ee288828901612abd565b9150509295509295909350565b8581528460208201527fffffffffffffffff0000000000000000000000000000000000000000000000008460c01b16604082015282151560f81b604882015260008251612f43816049850160208701612cf4565b919091016049019695505050505050565b80516fffffffffffffffffffffffffffffffff81168114611f7457600080fd5b600060808284031215612f8657600080fd5b6040516080810181811067ffffffffffffffff82111715612fa957612fa9612a69565b6040528251612fb781612a98565b815260208381015190820152612fcf60408401612f54565b6040820152612fe060608401612f54565b60608201529392505050565b600060a08284031215612ffe57600080fd5b60405160a0810181811067ffffffffffffffff8211171561302157613021612a69565b806040525082358152602083013560208201526040830135604082015260608301356060820152608083013560808201528091505092915050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b8481526000602060808184015280608084015260018060a085015260c0840160c060408601528087825260e08601905060e08860051b87010191508860005b8981101561318e577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff2088850301835281357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18c360301811261314557600080fd5b8b01868101903567ffffffffffffffff81111561316157600080fd5b80360382131561317057600080fd5b61317b86828461305c565b95505050918501919085019084016130e4565b50505080935050505082606083015295945050505050565b6000602082840312156131b857600080fd5b815161239781612e47565b600084516131d5818460208901612cf4565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551613211816001850160208a01612cf4565b6001920191820152835161322c816002840160208801612cf4565b0160020195945050505050565b60006020828403121561324b57600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008282101561329357613293613252565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000826132d6576132d6613298565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff83147f80000000000000000000000000000000000000000000000000000000000000008314161561332a5761332a613252565b500590565b6000808312837f80000000000000000000000000000000000000000000000000000000000000000183128115161561336957613369613252565b837f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01831381161561339d5761339d613252565b50500390565b60007f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6000841360008413858304851182821616156133e4576133e4613252565b7f8000000000000000000000000000000000000000000000000000000000000000600087128682058812818416161561341f5761341f613252565b6000871292508782058712848416161561343b5761343b613252565b8785058712818416161561345157613451613252565b505050929093029392505050565b6000808212827f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0384138115161561349957613499613252565b827f80000000000000000000000000000000000000000000000000000000000000000384128116156134cd576134cd613252565b50500190565b600067ffffffffffffffff8083168185168083038211156134f6576134f6613252565b01949350505050565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561353757613537613252565b500290565b60008261354b5761354b613298565b500490565b868152600073ffffffffffffffffffffffffffffffffffffffff808816602084015280871660408401525084606083015283608083015260c060a083015261359b60c0830184612d24565b98975050505050505050565b600082198211156135ba576135ba613252565b500190565b805163ffffffff81168114611f7457600080fd5b805160ff81168114611f7457600080fd5b600060c082840312156135f657600080fd5b60405160c0810181811067ffffffffffffffff8211171561361957613619612a69565b604052613625836135bf565b8152613633602084016135d3565b6020820152613644604084016135d3565b6040820152613655606084016135bf565b6060820152613666608084016135bf565b608082015261367760a08401612f54565b60a08201529392505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036136b4576136b4613252565b506001019056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(KromaPortalStorageLayoutJSON), KromaPortalStorageLayout); err != nil {
		panic(err)
	}

	layouts["KromaPortal"] = KromaPortalStorageLayout
	deployedBytecodes["KromaPortal"] = KromaPortalDeployedBin
}
