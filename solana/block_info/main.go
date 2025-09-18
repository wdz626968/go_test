package main

import (
	"context"
	"fmt"
	"log"

	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/rpc"
)

func main() {
	client := client.NewClient(rpc.MainnetRPCEndpoint)

	slot, err := client.GetSlot(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	block, err := client.GetBlock(context.Background(), slot)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("区块高度: %d\n", block.BlockHeight)
	fmt.Printf("交易数量: %d\n", len(block.Transactions))
	for i, tx := range block.Transactions {
		fmt.Printf("交易%d: %s\n", i, tx.Transaction.Message)
	}

	balance, err := client.GetBalance(
		context.Background(),
		"9qeP9DmjXAmKQc4wy133XZrQ3Fo4ejsYteA7X4YFJ3an",
	)
	if err != nil {
		log.Fatalf("failed to request airdrop, err: %v", err)
	}
	fmt.Println(balance)
}
