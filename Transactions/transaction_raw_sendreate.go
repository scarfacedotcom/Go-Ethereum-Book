package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
)

func main11() {
	client, err := ethclient.Dial("https://goerli.infura.io/v3/4048b34a30a34c099bd9280862364f8d")
	if err != nil {
		log.Fatal(err)
	}

	rawTx := "f869f8670436825208944592d8f8d7b001e72cb26a73e4fa1806a51ac79d880de0b6b3a7640000802da09e2a3ce88f77f6d2db39708c4eb0b2af88885264990abbfb19a5279735f9a15da029d369b6e382f4f12028d59806d2884b15ead1e3876bec6ab14aa8ee69b7fdf5"

	rawTxBytes, err := hex.DecodeString(rawTx)

	tx := new(types.Transaction)
	rlp.DecodeBytes(rawTxBytes, &tx)

	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0xc429e5f128387d224ba8bed6885e86525e14bfdc2eb24b5e9c3351a1176fd81f
}
