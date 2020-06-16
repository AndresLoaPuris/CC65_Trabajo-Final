package main

import (
	"Tutorial/blockchain"
	"fmt"
)

func main() {

	chain := blockchain.InitBlockChain()
	Person := blockchain.AddPerson("Andres Loa Puris", []float64{0.0034, 0.0012, 0.0123, 0.0071, 0.0016})
	chain.AddBlock(Person)

	for _, block := range chain.ReturnBlocks() {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Name in Block: %s\n", block.ReturnBlockName())
		fmt.Println("Data in Block:", block.ReturnBlockValues())
		fmt.Printf("Hash: %x\n\n", block.Hash)
	}
}
