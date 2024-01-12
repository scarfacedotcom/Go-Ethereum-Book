package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main2() {
	privateKey, err := crypto.HexToECDSA("7d0adcd321bc7dba076e50a117bcd9da6bca2ddaa907bf128f60db43f3802b8d")
	if err != nil {
		log.Fatal(err)
	}

	data := []byte("SCAR FACE")
	hash := crypto.Keccak256Hash(data)
	fmt.Println(hash.Hex())

	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hexutil.Encode(signature))
}
