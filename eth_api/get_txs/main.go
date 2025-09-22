package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {
	// 连接到以太坊Sepolia测试网节点
	client, err := ethclient.Dial("https://ethereum-sepolia-rpc.publicnode.com")
	if err != nil {
		log.Fatal(err) // 连接失败时终止程序
	}

	// 获取当前链ID
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err) // 获取失败时终止程序
	}

	// 指定要查询的区块号
	blockNumber := big.NewInt(5671744)
	// 获取指定区块的信息
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err) // 获取失败时终止程序
	}

	// 遍历区块中的所有交易
	for _, tx := range block.Transactions() {
		// 打印交易哈希
		fmt.Println(tx.Hash().Hex()) // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
		// 打印交易金额(wei)
		fmt.Println(tx.Value().String()) // 100000000000000000
		// 打印交易消耗的gas
		fmt.Println(tx.Gas()) // 21000
		// 打印gas价格(wei)
		fmt.Println(tx.GasPrice().Uint64()) // 100000000000
		// 打印交易nonce
		fmt.Println(tx.Nonce()) // 245132
		// 打印交易数据(输入数据)
		fmt.Println(tx.Data()) // []
		// 打印交易接收地址
		fmt.Println(tx.To().Hex()) // 0x8F9aFd209339088Ced7Bc0f57Fe08566ADda3587

		// 获取交易发送者地址(使用EIP155签名器)
		if sender, err := types.Sender(types.NewEIP155Signer(chainID), tx); err == nil {
			fmt.Println("sender", sender.Hex()) // 0x2CdA41645F2dBffB852a605E92B185501801FC28
		} else {
			log.Fatal(err) // 获取失败时终止程序
		}

		// 获取交易收据
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err) // 获取失败时终止程序
		}

		// 打印交易状态(1表示成功)
		fmt.Println(receipt.Status) // 1
		// 打印交易日志
		fmt.Println(receipt.Logs) // []
		break                     // 只处理第一条交易
	}

	// 通过区块哈希获取交易数量
	blockHash := common.HexToHash("0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5")
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		log.Fatal(err) // 获取失败时终止程序
	}

	// 遍历区块中的所有交易(通过区块哈希和索引)
	for idx := uint(0); idx < count; idx++ {
		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
		if err != nil {
			log.Fatal(err) // 获取失败时终止程序
		}

		fmt.Println(tx.Hash().Hex()) // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
		break                        // 只处理第一条交易
	}

	// 通过交易哈希获取交易详情
	txHash := common.HexToHash("0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5")
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err) // 获取失败时终止程序
	}
	// 打印交易是否待处理
	fmt.Println(isPending) // false
	// 打印交易哈希
	fmt.Println(tx.Hash().Hex()) // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
}
