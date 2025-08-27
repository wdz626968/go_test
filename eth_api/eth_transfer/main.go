package main

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	//连接测试网
	client, err := ethclient.Dial("https://ethereum-sepolia-rpc.publicnode.com")
	if err != nil {
		log.Fatal(err)
	}

	//加载私钥
	privateKey, err := crypto.HexToECDSA("30a01b6328ba5b31310aaa91afd95c95eebe1f853194a2dcd9f4dba36e75bf50")
	if err != nil {
		log.Fatal(err)
	}
	//通过私钥获取公钥
	publicKey := privateKey.Public()
	pulicKeyPtr := publicKey.(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*pulicKeyPtr)
	//获取账户交易随机数
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	//通过公钥获取交易地址对象
	toAddress := common.HexToAddress("0x4ffDe62cE898639329850eC872A5535F64cb8181")
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	gasLimit := uint64(21000) // in units
	//包装转账信息
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      gasLimit,
		To:       &toAddress,
		Value:    big.NewInt(1000000000000000),
		Data:     nil,
	})
	//获取链id
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	//交易签名
	signTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	//发起交易
	err = client.SendTransaction(context.Background(), signTx)
	if err != nil {
		log.Fatal(err)
	}

}
