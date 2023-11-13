package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func main2() {
	address := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")

	// Print the hexadecimal representation of the address
	fmt.Println(address.Hex()) // 0x71C7656EC7ab88b098defB751B7401B5f6d8976F

	// Calculate the hash of the address using Keccak256
	addressHash := crypto.Keccak256Hash(address.Bytes())
	fmt.Println(addressHash.Hex()) // 0x00000000000000000000000071c7656ec7ab88b098defb751b7401b5f6d8976f

	// Print the byte representation of the address
	fmt.Println(address.Bytes()) // [113 199 101 110 199 171 136 176 152 222 251 117 27 116 1 181 246 216 151 111]
}
