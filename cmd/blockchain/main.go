package main

import (
	"fmt"

	"github.com/AndriyAntonenko/blockchain/internal/blockchain"
)

func main() {
	myBlockchain := blockchain.NewBlockchain()
	lastBlock := myBlockchain.GetLastBlock()
	fmt.Println(lastBlock.SerializeJSON())

	myBlockchain.MineBlock("Block number 1!")
	lastBlock = myBlockchain.GetLastBlock()
	fmt.Println(lastBlock.SerializeJSON())

	myBlockchain.MineBlock("Block number 2!")
	lastBlock = myBlockchain.GetLastBlock()
	fmt.Println(lastBlock.SerializeJSON())
}
