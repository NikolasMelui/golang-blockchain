package blockchain

import "math/big"

// Difficulty ...
const Difficulty = 12

// ProofOfWork ...
type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

// NewProof ...
func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))
	pow := &ProofOfWork{b, target}
	return pow
}
