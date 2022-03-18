package main

import (
	"go-chain-man/domain"
)

func main() {
	myBlockchainAddress := "my_blockchain_address"
	blockchain := domain.NewBlockchain(myBlockchainAddress)

	blockchain.AddTransaction("A", "B", 1.0)
	blockchain.Mining()

	blockchain.AddTransaction("C", "B", 15.0)
	blockchain.AddTransaction("B", "C", 12.0)
	blockchain.Mining()

	blockchain.Print()
}
