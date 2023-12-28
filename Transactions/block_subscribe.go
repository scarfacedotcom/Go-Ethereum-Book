package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main99() {
	client, err := ethclient.Dial("wss://goerli.infura.io/ws/v3/4048b34a30a34c099bd9280862364f8d")
	if err != nil {
		log.Fatal(err)
	}

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println("The Hash of the New Block Header is: ", header.Hash().Hex())

			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("---------------------- N E W    B L O C K---------------------- ")
			fmt.Println("The Hash of this block is: ", block.Hash().Hex())
			fmt.Println("The Block Number is: ", block.Number().Uint64())
			fmt.Println("The Block Time is: ", block.Time())
			fmt.Println("The Block Nonce is: ", block.Nonce())
			fmt.Println("The Number of Transactions in this Block is: ", len(block.Transactions()))
		}
	}
}
