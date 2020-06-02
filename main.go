package main

import "log"

// Block ...
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

func main() {
	log.Println("First blockchain message")
}
