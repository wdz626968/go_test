package main

import (
	"fmt"
	store "go_test"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	//连接测试网
	client, err := ethclient.Dial("https://ethereum-sepolia-rpc.publicnode.com")
	if err != nil {
		log.Fatal(err)
	}
	// 创建私钥（在实际应用中，您应该使用更安全的方式来管理私钥）
	privateKey, err := crypto.HexToECDSA("d4f92103da1106a9eac579281458f51a541e0525253993246d8e08f440b28e77")
	if err != nil {
		log.Fatal(err)
	}
	storeContract, err := store.NewStore(common.HexToAddress("0xB5BeF50C0A41446968f6cB89Bb80e75B95cfc4C8"), client)
	if err != nil {
		log.Fatal(err)
	}
	opt, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))

	withdraw, err := storeContract.Withdraw(opt)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(withdraw)
}
