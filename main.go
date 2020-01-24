package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	conn, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal("Could not connect to Infura", err)
	}

	ctx := context.Background()
	// Random transaction hash from Etherscan
	txHash := "0x1656a0af683c6cfed1169bf8c459e13f0a8d7f8ce9545784ddeffcfa2a2bddd4"

	tx, pending, err := conn.TransactionByHash(ctx, common.HexToHash(txHash))
	if err != nil {
		log.Fatal("Could not get transaction", err)
	}

	if !pending {
		fmt.Printf("Transaction value: %d wei", tx.Value())
	}
}
