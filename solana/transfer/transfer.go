package main

import (
	"context"
	"fmt"
	"log"

	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/common"
	"github.com/blocto/solana-go-sdk/program/system"
	"github.com/blocto/solana-go-sdk/types"
)

var feePayer, _ = types.AccountFromBase58("4TMFNY9ntAn3CHzguSAvDNLPRoQTaK3sWbQQXdDXaE6KWRBLufGL6PJdsD2koiEe3gGmMdRK3aAw7sikGNksHJrN")

var alice, _ = types.AccountFromBase58("4voSPg3tYuWbKzimpQK9EbXHmuyy5fUrtXvpLDMLkmY6TRncaTHAKGD8jUg3maB5Jbrd9CkQg4qjJMyN6sQvnEF2")

func main() {
	c := client.NewClient("https://devnet.helius-rpc.com/?api-key=8a9947bf-2456-4824-b675-98bf7750e9ac")
	recentBlockhashResponse, err := c.GetLatestBlockhash(context.Background())

	balance, err := c.GetBalance(context.Background(), "9qeP9DmjXAmKQc4wy133XZrQ3Fo4ejsYteA7X4YFJ3an")
	if err != nil {
		panic(err)
	}
	log.Println(balance)

	feePayerBalance, err := c.GetBalance(context.Background(), "9qeP9DmjXAmKQc4wy133XZrQ3Fo4ejsYteA7X4YFJ3an")
	if err != nil {
		panic(err)
	}
	log.Println(feePayerBalance)
	//sig, err := c.RequestAirdrop(
	//	context.TODO(),
	//	"9qeP9DmjXAmKQc4wy133XZrQ3Fo4ejsYteA7X4YFJ3an", // address
	//	1e8, // lamports (1 SOL = 10^9 lamports)
	//)
	//if err != nil {
	//	log.Fatalf("failed to request airdrop, err: %v", err)
	//}
	//fmt.Println(sig)
	tx, err := types.NewTransaction(
		types.NewTransactionParam{
			Signers: []types.Account{feePayer, alice},
			Message: types.NewMessage(
				types.NewMessageParam{
					FeePayer:        feePayer.PublicKey,
					RecentBlockhash: recentBlockhashResponse.Blockhash,
					Instructions: []types.Instruction{
						system.Transfer(
							system.TransferParam{
								From:   alice.PublicKey,
								To:     common.PublicKeyFromString("2xNweLHLqrbx4zo1waDvgWJHgsUpPj8Y8icbAFeR4a8i"),
								Amount: 1e2,
							},
						),
					},
				},
			),
		},
	)
	if err != nil {
		panic(err)
	}
	txId, err := c.SendTransaction(context.Background(), tx)
	if err != nil {
		panic(err)
	}
	fmt.Println("txHash:", txId)
}
