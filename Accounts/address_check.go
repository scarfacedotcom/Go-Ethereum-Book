package main

import (
	"context"
	"fmt"
	"log"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

	fmt.Printf("is valid: %v\n", re.MatchString("0xe41d2489571d322189246dafabde1f4699f498"))
	fmt.Printf("is valid: %v\n", re.MatchString("0xe41d2489571d322189246dafa5ebde1f4699f498"))

	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	// 0x Protocol Token (ZRX) smart contract address
	address := common.HexToAddress("0xe41d2489571d322189246dafa5ebde1f4699f498")
	bytecode, err := client.CodeAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}

	isContract := len(bytecode) > 0

	fmt.Printf("is it a Smart contract Account: %v\n", isContract)

	// Scar Faces account address
	address = common.HexToAddress("0x8e215d06ea7ec1fdb4fc5fd21768f4b34ee92ef4")
	bytecode, err = client.CodeAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}

	isContract = len(bytecode) > 0

	fmt.Printf("is it a Smart contract Account: %v\n", isContract)
}
