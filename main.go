package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type BlockChain struct {
	blocks []*Block
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, preveHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), preveHash}
	block.DeriveHash()
	return block
}

func (blockChain *BlockChain) AddBlock(data string) {
	preveBlock := blockChain.blocks[len(blockChain.blocks)-1]
	block := CreateBlock(data, preveBlock.Hash)
	blockChain.blocks = append(blockChain.blocks, block)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {
	chain := InitBlockChain()

	chain.AddBlock("Fist block affter Genesis")
	chain.AddBlock("Second block affter Genesis")
	chain.AddBlock("Third block affter Genesis")

	for _, block := range chain.blocks {
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("Current hash: %x\n", block.Hash)
		fmt.Printf("Data block: %s\n", block.Data)
	}
}
