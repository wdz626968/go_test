package main

import (
	"context"
	"fmt"

	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/program/token"
)

func main() {
	c := client.NewClient("https://devnet.helius-rpc.com/?api-key=8a9947bf-2456-4824-b675-98bf7750e9ac")

	info, err := c.GetAccountInfo(context.Background(), "HeCBh32JJ8DxcjTyc6q46tirHR8hd2xj3mGoAcQ7eduL")
	if err != nil {
		panic(err)
	}
	fmt.Println("account info:", info)
	tokenAccount, err := token.TokenAccountFromData(info.Data)
	if err != nil {
		panic(err)
	}
	fmt.Println("token account:", tokenAccount)
}
