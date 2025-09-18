package main

import (
	"context"
	"log"

	"github.com/davecgh/go-spew/spew"

	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/common"
	"github.com/blocto/solana-go-sdk/program/metaplex/token_metadata"
)

func main() {
	// NFT in solana is a normal mint but only mint 1.
	// If you want to get its metadata, you need to know where it stored.
	// and you can use `tokenmeta.GetTokenMetaPubkey` to get the metadata account key
	// here I take a random Degenerate Ape Academy as an example
	mint := common.PublicKeyFromString("GphF2vTuzhwhLWBWWvD8y5QLCPp1aQC5EnzrWsnbiWPx")
	metadataAccount, err := token_metadata.GetTokenMetaPubkey(mint)
	if err != nil {
		log.Fatalf("faield to get metadata account, err: %v", err)
	}

	// new a client
	c := client.NewClient("https://devnet.helius-rpc.com/?api-key=8a9947bf-2456-4824-b675-98bf7750e9ac")

	// get data which stored in metadataAccount
	accountInfo, err := c.GetAccountInfo(context.Background(), metadataAccount.ToBase58())
	if err != nil {
		log.Fatalf("failed to get accountInfo, err: %v", err)
	}

	log.Println("metadata account: ", metadataAccount.ToBase58())
	log.Println("metadata account info: ", accountInfo)
	// parse it
	metadata, err := token_metadata.MetadataDeserialize(accountInfo.Data)
	if err != nil {
		log.Fatalf("failed to parse metaAccount, err: %v", err)
	}
	spew.Dump(metadata)
}
