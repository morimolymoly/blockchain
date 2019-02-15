package main

import (
	"fmt"
	"strconv"

	b "github.com/morimolymoly/blockchain/block"
)

func main() {
	fmt.Println("Hello Blockchain")
	bc, err := b.NewBlockchain()
	if err != nil {
		panic(err)
	}

	if err := bc.AddBlock("Send 1 BTC to Ivan"); err != nil {
		panic(err)
	}
	if err := bc.AddBlock("Send 2 more BTC to Ivan"); err != nil {
		panic(err)
	}

	i := bc.Iterator()
	for {
		block, err := i.Next()
		if err != nil {
			panic(err)
		}
		if i.Final {
			break
		}
		pow := b.NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
