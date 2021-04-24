package main

import (
	"fmt"
	"strconv"

	"github.com/ebanstev/blockchain/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("Fist block affter Genesis")
	chain.AddBlock("Second block affter Genesis")
	chain.AddBlock("Third block affter Genesis")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("Current hash: %x\n", block.Hash)
		fmt.Printf("Data block: %s\n", block.Data)

		pow := blockchain.NewProof(block)

		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
