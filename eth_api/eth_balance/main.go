package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// todo
func main() {
	client, err := ethclient.Dial("https://ethereum-sepolia-rpc.publicnode.com")
	if err != nil {
		log.Fatal(err)
	}
	//直接读取最新的余额
	address := common.HexToAddress("0x50a133dC41Dc07D81338aAF4FD4FAb88a90dc489")
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance) //488535222245672745

	////指定区块高度
	//blockNumber := big.NewInt(9024501)
	//balance1, err := client.BalanceAt(context.Background(), address, blockNumber)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(balance1)

	//将余额转成以eth为单位的小数
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	fmt.Println(ethValue) // 0.488535222245672745

	//tokenAddress := common.HexToAddress("0x80A82BAa2eb8969d14aFCa8089f2d7E0E3eCE2E2")
}
