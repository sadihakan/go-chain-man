package main

import (
	"fmt"
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

	fmt.Printf("my_blockchain_address %.1f\n", blockchain.CalculateTotalAmount("my_blockchain_address"))
	fmt.Printf("C %.1f\n", blockchain.CalculateTotalAmount("C"))
	fmt.Printf("D %.1f\n", blockchain.CalculateTotalAmount("D"))
}
