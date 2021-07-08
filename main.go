package main

import (
	"BlockChain-Golang/blocks"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strconv"
)

func main() {
	log.Info("Hi, this is blockchain project")
	chain := blocks.InitBlockChain()

	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")
	chain.AddBlock("Fourth Block")

	for _, block := range chain.Blockchain {
		fmt.Println("Hash", block.Hash)
		fmt.Println("Data", block.Data)
		fmt.Println("PrevHash", block.PreviousHash)

		pow := blocks.NewProof(block)
		fmt.Println("Proof of Work: ", strconv.FormatBool(pow.Validate()))
		fmt.Println("==========================================================")
	}

}
