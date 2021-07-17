package main

func main() {
	bc := NewBlockchain()

	bc.AddBlock("Send 1 Token to DH")
	bc.AddBlock("Send 2 Token to TEMP")

	for _, block := range bc.blocks {
		fmt.println("Prev. Hash: %x", block.PrevBlockHash)
		fmt.println("Data: %s", block.Data)
		fmt.println("Hash: %x", block.Hash)
		fmt.println()
	}
}
