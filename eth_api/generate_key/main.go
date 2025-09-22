package main

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
	"log"
)

func main() {

	privateKey, err2 := crypto.GenerateKey()
	if err2 != nil {
		log.Fatal(err2)
	}
	//转换成私钥字节切片
	ecdsaBytes := crypto.FromECDSA(privateKey)
	fmt.Println("Private Key: %x", hexutil.Encode(ecdsaBytes)[2:])
	//通过私钥获取公钥
	publicKey := privateKey.Public()
	//包装公钥为指针
	ecdsaPublicKeyPtr, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key")
	}
	//转换成公钥字节切片
	publicKeyBytes := crypto.FromECDSAPub(ecdsaPublicKeyPtr)
	fmt.Println("Public Key: %x", hexutil.Encode(publicKeyBytes)[4:])
	//公钥转成Keccak256的20位地址
	address := crypto.PubkeyToAddress(*ecdsaPublicKeyPtr)
	fmt.Println("Address:", address.Hex())
	hash := sha3.NewLegacyKeccak256()

	//手动将公钥转成Keccak256的20位地址
	hash.Write(publicKeyBytes[1:])
	fmt.Println("Hash: %x", hexutil.Encode(hash.Sum(nil))[:])
	fmt.Println("Public Key: %x", hexutil.Encode(hash.Sum(nil)[12:]))
}
