package main

import "fmt"

func main() {
	bc := NewBlockchain()

	bc.AddBlock("Send 1 Token to DH")
	bc.AddBlock("Send 2 Token to TEMP")

	for _, block := range bc.blocks {
		fmt.Println("Prev. Hash: %x", block.PrevBlockHash)
		fmt.Println("Data: %s", block.Data)
		fmt.Println("Hash: %x", block.Hash)
		fmt.Println()
	}
}
