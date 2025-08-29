package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	store "go_test"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	//连接测试网
	client, err := ethclient.Dial("https://ethereum-sepolia-rpc.publicnode.com")
	if err != nil {
		log.Fatal(err)
	}
	privateKey, err := crypto.HexToECDSA("d4f92103da1106a9eac579281458f51a541e0525253993246d8e08f440b28e77")
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public().(*ecdsa.PublicKey)

	fromAddress := crypto.PubkeyToAddress(*publicKey)
	//toAddress := common.HexToAddress("0x4ffDe62cE898639329850eC872A5535F64cb8181")

	if err != nil {
		log.Fatal(err)
	}

	nounce, err := client.NonceAt(context.Background(), fromAddress, nil)

	gasPrice, err := client.SuggestGasPrice(context.Background())

	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nounce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice
	input := big.NewInt(1)
	address, tx, instance, err := store.DeployStore(auth, client, input)
	if err != nil {
		log.Fatal(err)
	}
	_ = instance
	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())
}
