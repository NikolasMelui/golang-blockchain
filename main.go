package main

import (
	"bytes"
	"crypto/sha256"
	"log"
)

// Block ...
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

// DeriveHash ...
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func main() {
	log.Println("First blockchain message")
}
