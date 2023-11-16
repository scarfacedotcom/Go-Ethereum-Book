package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main1() {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	//client, err := ethclient.Dial("/tmp/geth.ipc")
	if err != nil {
		log.Fatal(err)
	}

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The Latest Block Header is: ", header.Number.String())

	blockNumber := big.NewInt(5677744)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The Block Header Number is: ", block.Number().Uint64())
	fmt.Println("The Block Time is: ", block.Time())
	fmt.Println("The Block Difficulty is: ", block.Difficulty().Uint64())
	fmt.Println("The Block Hash is: ", block.Hash().Hex())
	fmt.Println("The Number of Transactions is: ", len(block.Transactions()))
	fmt.Println("The hash of the first Transaction is: ", block.Transactions()[0].Hash().Hex())
	fmt.Println("The hash of the Second Transaction is: ", block.Transactions()[1].Hash().Hex())
	fmt.Println("The hash of the Third Transaction is: ", block.Transactions()[2].Hash().Hex())
	fmt.Println("The hash of the Fourth Transaction is: ", block.Transactions()[3].Hash().Hex())
	// 11. Gas Limit and Gas Used
	fmt.Println("Gas Limit of the Block: ", block.GasLimit())
	fmt.Println("Gas Used by Transactions in the Block: ", block.GasUsed())

	// 12. Block Nonce
	fmt.Println("Nonce of the Block: ", block.Nonce())

	// 13. Uncle Blocks
	fmt.Println("Number of Uncle Blocks: ", len(block.Uncles()))
	fmt.Println("Hash of Uncle Blocks: ", block.UncleHash().Hex())

	// 14. Coinbase Address
	fmt.Println("Coinbase Address: ", block.Coinbase().Hex())

	// 15. Extra Data
	fmt.Println("Extra Data of the Block: ", string(block.Extra()))

	// 16. Receipts Root and State Root
	//fmt.Println("Receipts Root: ", block.ReceiptsRoot().Hex())
	fmt.Println("State Root: ", block.Root().Hex())

	// 17. Size and Gas Used by the Block
	fmt.Println("Size of the Block (in bytes): ", block.Size())
	fmt.Println("Total Gas Used by Transactions in the Block: ", block.GasUsed())
	//fmt.Println("Cumulative Gas Used in the Block: ", block.CumulativeGasUsed())

	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The Number of Transactions is: ", count)
}
