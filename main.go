package main

import (
	"encoding/json"
	"fmt"
)

// import (
// "fmt"
// "strconv"

// "github.com/nikolasmelui/golang-blockchain/blockchain"
// )

// DocumentOne ...
type DocumentOne struct {
	Name        string
	Description string
}

// DocumentTwo ...
type DocumentTwo struct {
	Name        string
	Description string
}

// PrintInterface ...
type PrintInterface interface {
	Print() string
}

// Print ...
func (d *DocumentOne) Print() string {
	fmt.Printf("Document: \n%s\n", d)
	out, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}
	return string(out)
}

// Print ...
func (d *DocumentTwo) Print() {
	fmt.Printf("Document: \n%s\n", d)
}

func main() {
	// chain := blockchain.InitBlockChain()

	// chain.AddBlock("First block")
	// chain.AddBlock("Second block")
	// chain.AddBlock("Third block")

	// for _, block := range chain.Blocks {
	// 	fmt.Printf("Previous Hash: %x\n", block.PrevHash)
	// 	fmt.Printf("Block Data : %s\n", block.Data)
	// 	fmt.Printf("Block Hash: %x\n", block.Hash)

	// 	pow := blockchain.NewProof(block)
	// 	fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
	// 	fmt.Println()
	// }

	// var docOne PrintInterface = &DocumentOne{"firstdocument", "firstdoc"}
	var docOne PrintInterface = new(DocumentOne)
	// docTwo := &DocumentTwo{"seconddocument", "seconddoc"}
	hello := docOne.Print()
	fmt.Println(hello)
	// docTwo.Print()

}
