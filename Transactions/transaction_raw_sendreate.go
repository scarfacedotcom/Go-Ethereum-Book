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

	rawTx := "f868f866020a825208944592d8f8d7b001e72cb26a73e4fa1806a51ac79d880de0b6b3a7640000802ea0664e67e04852b0304b8cdbdb5df91088aea7a631fb364332a22a47cb4f831f529fc4632ce9049e91449f1abd136d8c39e67b410d651d1c5d28c28f9b4e24d65c"

	rawTxBytes, err := hex.DecodeString(rawTx)

	tx := new(types.Transaction)
	rlp.DecodeBytes(rawTxBytes, &tx)

	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0xc429e5f128387d224ba8bed6885e86525e14bfdc2eb24b5e9c3351a1176fd81f
}
