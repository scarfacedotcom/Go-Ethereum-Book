package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/rlp"
)

type MyStruct struct {
	Field1 string
	Field2 uint64
}

func main() {
	// Encoding data
	data := MyStruct{"Hello", 42}
	encodedData, err := rlp.EncodeToBytes(data)
	if err != nil {
		fmt.Println("Error encoding:", err)
		return
	}

	fmt.Printf("Encoded data: %x\n", encodedData)

	// Decoding data
	var decodedData MyStruct
	err = rlp.DecodeBytes(encodedData, &decodedData)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return
	}

	fmt.Printf("Decoded data: %+v\n", decodedData)
}
