package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The Block Header Number is: ", header.Number.String()) // 5671744

	blockNumber := big.NewInt(5671744)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The Block Header Number is: ", block.Number().Uint64())
	fmt.Println("The Block Time is: ", block.Time())
	fmt.Println("The Block Difficulty is: ", block.Difficulty().Uint64())
	fmt.Println("The Block Hash is: ", block.Hash().Hex())
	fmt.Println("The Number of Transactions is: ", len(block.Transactions()))

	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The Number of Transactions is: ", count)
}
