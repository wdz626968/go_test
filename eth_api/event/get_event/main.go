package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var LockABI = `[{"inputs":[{"internalType":"uint256","name":"_unlockTime","type":"uint256"}],"stateMutability":"payable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"uint256","name":"amount","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"when","type":"uint256"}],"name":"Withdrawal","type":"event"},{"inputs":[],"name":"owner","outputs":[{"internalType":"address payable","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"unlockTime","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"withdraw","outputs":[],"stateMutability":"nonpayable","type":"function"}]`

func main() {
	//连接测试网
	client, err := ethclient.Dial("https://ethereum-sepolia-rpc.publicnode.com")
	if err != nil {
		log.Fatal(err)
	}
	//合约地址
	contractAddress := common.HexToAddress("0x9BFAC4bfe40D9400B0226996912204e6dFbe8D2E")
	query := ethereum.FilterQuery{
		//BlockHash: nil,
		//从哪个区块高度开始检索，不传要从创世区块开始查
		FromBlock: big.NewInt(9079000),
		//从哪个区块高度开始检索，不传要从创世区块开始查
		//ToBlock:   nil,
		//合约地址
		Addresses: []common.Address{contractAddress},
		//主题
		//Topics:    nil,
	}
	//调用eth全节点过滤出相关日志
	logs, err := client.FilterLogs(context.Background(), query)

	if err != nil {
		log.Fatal(err)
	}
	//获取合约ABI
	contractAbi, err := abi.JSON(strings.NewReader(LockABI))
	if err != nil {
		log.Fatal(err)
	}
	for _, vLog := range logs {
		fmt.Println(vLog.BlockHash.Hex())
		fmt.Println(vLog.BlockNumber)
		fmt.Println(vLog.TxHash.Hex())
		event := struct {
			Amount *big.Int
			When   *big.Int
		}{}
		//通过调用ABI获取event返回值
		err := contractAbi.UnpackIntoInterface(&event, "Withdrawal", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(event.Amount)
		fmt.Println(event.When)
		var topics []string
		for i := range vLog.Topics {
			topics = append(topics, vLog.Topics[i].Hex())
		}
		fmt.Println("topics[0]=", topics[0]) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
		if len(topics) > 1 {
			fmt.Println("indexed topics:", topics[1:])
		}
	}
	eventSignature := []byte("ItemSet(bytes32,bytes32)")
	hash := crypto.Keccak256Hash(eventSignature)
	fmt.Println("signature topics=", hash.Hex())
}
