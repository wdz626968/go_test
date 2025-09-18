package main

import (
	"context"
	"fmt"

	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/common"
	"github.com/blocto/solana-go-sdk/program/token"
)

var mintPubkey = common.PublicKeyFromString("F6tecPzBMF47yJ2EN6j2aGtE68yR5jehXcZYVZa6ZETo")

func main() {
	c := client.NewClient("https://devnet.helius-rpc.com/?api-key=8a9947bf-2456-4824-b675-98bf7750e9ac")

	info, err := c.GetAccountInfo(context.Background(), mintPubkey.ToBase58())

	if err != nil {
		panic(err)
	}
	mintAccount, err := token.MintAccountFromData(info.Data)
	if err != nil {
		panic(err)
	}
	fmt.Println("mint account:", mintAccount)
}
