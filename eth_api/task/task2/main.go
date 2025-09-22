package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	counter "go_test"
	"log"
	"math/big"
)

func main() {
	//连接测试网
	client, err := ethclient.Dial("https://eth-sepolia-rpc.publicnode.com")
	if err != nil {
		log.Fatal(err)
	}
	privateKey, err := crypto.HexToECDSA("a66a0d9eaba66ecdd5d3e16025ae3a87d96753a95a0b451d9f0f650bf2f64fb2")
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	gasLimit := uint64(2100000)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice
	deployCounter, transaction, c, err := counter.DeployCounter(auth, client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(deployCounter.Hex())
	fmt.Println(transaction.Hash().Hex())
	_ = c
}
