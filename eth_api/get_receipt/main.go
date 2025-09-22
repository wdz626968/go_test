package main

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
	"math/big"
)

func main() {
	client, err := ethclient.Dial("https://ethereum-sepolia-rpc.publicnode.com")
	if err != nil {
		log.Fatal(err)
	}

	//通过blockHash获取收据
	blockHash := common.HexToHash("0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5")
	blockNumberOrHash := rpc.BlockNumberOrHashWithHash(blockHash, false)
	receipts, err := client.BlockReceipts(context.Background(), blockNumberOrHash)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Block Receipts: %v", receipts)
	//通过blockNumber获取收据列表
	blockNumber := big.NewInt(5671744)
	blockByNum := rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blockNumber.Int64()))
	blockReceipts, err1 := client.BlockReceipts(context.Background(), blockByNum)
	if err1 != nil {
		log.Fatal(err1)
	}
	log.Printf("Block Receipts: %v", blockReceipts)

	//通过交易hash获取收据
	txHash := common.HexToHash("0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5")
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Transaction Receipt: %v", receipt)
}
