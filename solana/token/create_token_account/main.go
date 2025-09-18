package main

import (
	"context"
	"fmt"

	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/common"
	"github.com/blocto/solana-go-sdk/program/associated_token_account"
	"github.com/blocto/solana-go-sdk/types"
)

// FUarP2p5EnxD66vVDL4PWRoWMzA56ZVHG24hpEDFShEz
var feePayer, _ = types.AccountFromBase58("4TMFNY9ntAn3CHzguSAvDNLPRoQTaK3sWbQQXdDXaE6KWRBLufGL6PJdsD2koiEe3gGmMdRK3aAw7sikGNksHJrN")

// 9aE476sH92Vz7DMPyq5WLPkrKWivxeuTKEFKd2sZZcde
var alice, _ = types.AccountFromBase58("4voSPg3tYuWbKzimpQK9EbXHmuyy5fUrtXvpLDMLkmY6TRncaTHAKGD8jUg3maB5Jbrd9CkQg4qjJMyN6sQvnEF2")

var mintPubkey = common.PublicKeyFromString("F6tecPzBMF47yJ2EN6j2aGtE68yR5jehXcZYVZa6ZETo")

func main() {
	c := client.NewClient("https://devnet.helius-rpc.com/?api-key=8a9947bf-2456-4824-b675-98bf7750e9ac")
	//获取派生出的ATA
	ata, _, err := common.FindAssociatedTokenAddress(alice.PublicKey, mintPubkey)

	if err != nil {
		panic(err)
	}
	recentBlockhashResponse, err := c.GetLatestBlockhash(context.Background())
	if err != nil {
		panic(err)
	}
	transaction, err := types.NewTransaction(types.NewTransactionParam{
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer: feePayer.PublicKey,
			Instructions: []types.Instruction{associated_token_account.Create(associated_token_account.CreateParam{
				Funder:                 feePayer.PublicKey,
				Owner:                  alice.PublicKey,
				Mint:                   mintPubkey,
				AssociatedTokenAccount: ata,
			})},
			RecentBlockhash: recentBlockhashResponse.Blockhash,
		}),
		Signers: []types.Account{feePayer},
	})
	if err != nil {
		panic(err)
	}
	txId, err := c.SendTransaction(context.Background(), transaction)
	if err != nil {
		panic(err)
	}
	fmt.Println("txHash:", txId)
}
