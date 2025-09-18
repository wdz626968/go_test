package main

import (
	"context"
	"fmt"

	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/common"
	"github.com/blocto/solana-go-sdk/program/system"
	"github.com/blocto/solana-go-sdk/program/token"
	"github.com/blocto/solana-go-sdk/types"
)

var feePayer, _ = types.AccountFromBase58("4TMFNY9ntAn3CHzguSAvDNLPRoQTaK3sWbQQXdDXaE6KWRBLufGL6PJdsD2koiEe3gGmMdRK3aAw7sikGNksHJrN")

var alice, _ = types.AccountFromBase58("4voSPg3tYuWbKzimpQK9EbXHmuyy5fUrtXvpLDMLkmY6TRncaTHAKGD8jUg3maB5Jbrd9CkQg4qjJMyN6sQvnEF2")

func main() {
	c := client.NewClient("https://devnet.helius-rpc.com/?api-key=8a9947bf-2456-4824-b675-98bf7750e9ac")
	recentBlockhashResponse, err := c.GetLatestBlockhash(context.Background())
	if err != nil {
		panic(err)
	}
	mint := types.NewAccount()
	fmt.Println("mintAddress:", mint)
	// 获取空账户的rent
	rentExemptionBalance, err := c.GetMinimumBalanceForRentExemption(
		context.Background(),
		token.MintAccountSize,
	)
	transaction, err := types.NewTransaction(types.NewTransactionParam{
		Message: types.NewMessage(
			types.NewMessageParam{
				FeePayer: feePayer.PublicKey,
				Instructions: []types.Instruction{
					system.CreateAccount(system.CreateAccountParam{
						From:     feePayer.PublicKey,
						New:      mint.PublicKey,
						Owner:    common.TokenProgramID,
						Lamports: rentExemptionBalance,
						Space:    token.MintAccountSize,
					}),
					token.InitializeMint(token.InitializeMintParam{
						Decimals:   8,
						Mint:       mint.PublicKey,
						MintAuth:   alice.PublicKey,
						FreezeAuth: nil,
					}),
				},
				RecentBlockhash: recentBlockhashResponse.Blockhash,
			},
		),
		Signers: []types.Account{feePayer, mint},
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
