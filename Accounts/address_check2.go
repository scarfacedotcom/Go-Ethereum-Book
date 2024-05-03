package main

import (
	"context"
	"fmt"
	"log"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// validateAddress checks if an address is valid
func validateAddress(address string) bool {
	// Define a regular expression pattern for an Ethereum address
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.MatchString(address)
}

// isSmartContract checks if an Ethereum address is a smart contract address
func isSmartContract(client *ethclient.Client, address common.Address) bool {
	// Get the code at the specified address
	bytecode, err := client.CodeAt(context.Background(), address, nil)
	if err != nil {
		log.Fatalf("Failed to retrieve code at address: %v", err)
	}

	// If the code length is greater than zero, it's a smart contract
	return len(bytecode) > 0
}

func main() {
	// Connect to an Ethereum node using a JSON-RPC endpoint
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum node: %v", err)
	}
	defer client.Close()

	// Define the address you want to validate and check
	address := "0xe41d2489571d322189246dafa5ebde1f4699f498"

	// Validate the address
	if !validateAddress(address) {
		fmt.Println("Invalid Ethereum address")
		return
	}
	fmt.Println("Valid Ethereum address")

	// Check if the address is a smart contract address
	addressCommon := common.HexToAddress(address)
	if isSmartContract(client, addressCommon) {
		fmt.Println("The address is a smart contract address")
	} else {
		fmt.Println("The address is not a smart contract address")
	}
}
