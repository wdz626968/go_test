package main

import (
	"context"
	"github.com/gagliardetto/solana-go/rpc"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc/ws"
)

func main() {
	ctx := context.Background()
	client, err := ws.Connect(context.Background(), "wss://devnet.helius-rpc.com/?api-key=8a9947bf-2456-4824-b675-98bf7750e9ac")
	if err != nil {
		panic(err)
	}
	program := solana.MustPublicKeyFromBase58("9xQeWvG816bUx9EPjHmaT23yvVM2ZWbrrpZb9PusVFin") // serum

	{
		sub, err := client.LogsSubscribe(
			ws.LogsSubscribeFilterAll,
			rpc.CommitmentRecent,
		)
		if err != nil {
			panic(err)
		}
		defer sub.Unsubscribe()

		for {
			got, err := sub.Recv(ctx)
			if err != nil {
				panic(err)
			}
			spew.Dump(got)
		}
	}
	if false {
		sub, err := client.AccountSubscribeWithOpts(
			program,
			"",
			// You can specify the data encoding of the returned accounts:
			solana.EncodingBase64,
		)
		if err != nil {
			panic(err)
		}
		defer sub.Unsubscribe()

		for {
			got, err := sub.Recv(ctx)
			if err != nil {
				panic(err)
			}
			spew.Dump(got)
		}
	}
}
