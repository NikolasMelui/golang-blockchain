package blockchain

import "fmt"

// BlockChain ...
type BlockChain struct {
	Blocks []*Block
}

// Block ...
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

// CreateBlock ...
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

// AddBlock ...
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	fmt.Println(prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

// Genesis ...
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// InitBlockChain ...
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
