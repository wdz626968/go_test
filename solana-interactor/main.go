package main

import (
	"context"
	"fmt"
	"github.com/blocto/solana-go-sdk/types"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/rpc"
	confirm "github.com/gagliardetto/solana-go/rpc/sendAndConfirmTransaction"
	"github.com/gagliardetto/solana-go/rpc/ws"
	"golang.org/x/time/rate"
	"log"
	"time"
)

func main() {
	c := rpc.NewWithCustomRPCClient(rpc.NewWithLimiter(rpc.DevNet_RPC, rate.Every(time.Second), 5))
	wsClient, err := ws.Connect(context.Background(), rpc.DevNet_WS)

	//slot, err := c.GetSlot(context.Background(), rpc.CommitmentConfirmed)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//// 获取最新区块
	//recentBlock, err := c.GetBlock(context.Background(), slot) // 0表示最新区块
	//
	//if err != nil {
	//	log.Fatalf("获取最新插槽失败: %v", err)
	//}
	//
	//if err != nil {
	//	panic("查询失败: " + err.Error())
	//}
	//fmt.Println("slot:", slot)
	//fmt.Printf("区块高度: %d\n", recentBlock.BlockHeight)
	//fmt.Printf("交易数量: %d\n", len(recentBlock.Transactions))

	// 生成发送方密钥对（实际项目应从钱包加载）
	payer := types.NewAccount()
	receipt := types.NewAccount()

	blockhash, err := c.GetLatestBlockhash(context.Background(), rpc.CommitmentFinalized)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("区块hash", blockhash)
	//创建转账指令
	instruction := system.NewTransferInstruction(1_000_000, solana.PublicKey(payer.PublicKey), solana.PublicKey(receipt.PublicKey))

	//构建交易
	transaction, err := solana.NewTransaction([]solana.Instruction{instruction.Build()}, blockhash.Value.Blockhash, solana.TransactionPayer(solana.PublicKey(payer.PublicKey)))

	if err != nil {
		log.Fatal(err)
	}
	//发送交易
	signature, err := confirm.SendAndConfirmTransaction(context.Background(), c, wsClient, transaction)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("交易已提交，哈希: %s\n", signature)

}
