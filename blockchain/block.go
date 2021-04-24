package blockchain

type BlockChain struct {
	Blocks []*Block
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

func CreateBlock(data string, preveHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), preveHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

func (blockChain *BlockChain) AddBlock(data string) {
	preveBlock := blockChain.Blocks[len(blockChain.Blocks)-1]
	block := CreateBlock(data, preveBlock.Hash)
	blockChain.Blocks = append(blockChain.Blocks, block)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
