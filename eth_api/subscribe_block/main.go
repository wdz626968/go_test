package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("wss://ethereum-sepolia-rpc.publicnode.com")
	if err != nil {
		log.Fatal(err)
	}
	//创建通道
	headers := make(chan *types.Header)
	//通道订阅新区块
	head, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}
	//主进程接收通道信息
	for {
		select {
		case err := <-head.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println("NewHead:", header.Hash().Hex())
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(block.Hash().Hex())      // 0x92231be5b9137147f36eeb98e0a16b3f703ba5152e0f46cb5832a35be94678cd
			fmt.Println(block.Number().Uint64()) // 9073862
			fmt.Println(block.Time())            // 1756284408
			fmt.Println(block.Nonce())           // 0
		}
	}
}
