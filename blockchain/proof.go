package blockchain

import (
	"bytes"
	"encoding/binary"
	"log"
	"math/big"
)

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

// InitData ...
func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.Data,
			ToHex(int64(nonce)),
			ToHex(int64(Difficulty)),
		},
		[]byte{})
	return data
}

// ToHex ...
func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}
