package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	// store合约的字节码
	contractBytecode = "60806040526040516105ac3803806105ac833981810160405281019061002591906100ea565b804210610067576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161005e90610195565b60405180910390fd5b805f819055503360015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506101b3565b5f5ffd5b5f819050919050565b6100c9816100b7565b81146100d3575f5ffd5b50565b5f815190506100e4816100c0565b92915050565b5f602082840312156100ff576100fe6100b3565b5b5f61010c848285016100d6565b91505092915050565b5f82825260208201905092915050565b7f556e6c6f636b2074696d652073686f756c6420626520696e20746865206675745f8201527f7572650000000000000000000000000000000000000000000000000000000000602082015250565b5f61017f602383610115565b915061018a82610125565b604082019050919050565b5f6020820190508181035f8301526101ac81610173565b9050919050565b6103ec806101c05f395ff3fe608060405234801561000f575f5ffd5b506004361061003f575f3560e01c8063251c1aa3146100435780633ccfd60b146100615780638da5cb5b1461006b575b5f5ffd5b61004b610089565b604051610058919061023e565b60405180910390f35b61006961008e565b005b610073610201565b6040516100809190610296565b60405180910390f35b5f5481565b5f544210156100d2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100c990610309565b60405180910390fd5b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610161576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161015890610371565b60405180910390fd5b7fbf2ed60bd5b5965d685680c01195c9514e4382e28e3a5a2d2d5244bf59411b93474260405161019292919061038f565b60405180910390a160015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc4790811502906040515f60405180830381858888f193505050501580156101fe573d5f5f3e3d5ffd5b50565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f819050919050565b61023881610226565b82525050565b5f6020820190506102515f83018461022f565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61028082610257565b9050919050565b61029081610276565b82525050565b5f6020820190506102a95f830184610287565b92915050565b5f82825260208201905092915050565b7f596f752063616e277420776974686472617720796574000000000000000000005f82015250565b5f6102f36016836102af565b91506102fe826102bf565b602082019050919050565b5f6020820190508181035f830152610320816102e7565b9050919050565b7f596f75206172656e277420746865206f776e65720000000000000000000000005f82015250565b5f61035b6014836102af565b915061036682610327565b602082019050919050565b5f6020820190508181035f8301526103888161034f565b9050919050565b5f6040820190506103a25f83018561022f565b6103af602083018461022f565b939250505056fea2646970667358221220c05681039d8c36fb468fd0060a14604f1877e642f8eaa1b7221912faa418831564736f6c634300081e0033"
)

func main() {
	// 连接到以太坊网络（这里使用 Goerli 测试网络作为示例）
	client, err := ethclient.Dial("https://ethereum-sepolia-rpc.publicnode.com")
	if err != nil {
		log.Fatal(err)
	}

	// 创建私钥（在实际应用中，您应该使用更安全的方式来管理私钥）
	privateKey, err := crypto.HexToECDSA("d4f92103da1106a9eac579281458f51a541e0525253993246d8e08f440b28e77")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 获取nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	// 获取建议的gas价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	data, err := hex.DecodeString(contractBytecode)
	if err != nil {
		log.Fatal(err)
	}
	// 创建交易
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      uint64(1000000),
		Value:    big.NewInt(0),
		//没有to变量
		//data是合约的字节码
		Data: data,
	})

	// 签名交易
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Transaction sent: %s\n", signedTx.Hash().Hex())

	// 等待交易被打包
	receipt, err := waitForReceipt(client, signedTx.Hash())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Contract deployed at: %s\n", receipt.ContractAddress.Hex())
}

func waitForReceipt(client *ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
	for {
		receipt, err := client.TransactionReceipt(context.Background(), txHash)
		if err == nil {
			return receipt, nil
		}
		if err != ethereum.NotFound {
			return nil, err
		}
		// 等待一段时间后再次查询
		time.Sleep(1 * time.Second)
	}
}
