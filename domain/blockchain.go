package domain

import (
	"fmt"
	"log"
	"strings"
)

type Blockchain struct {
	transactionPool []*Transaction
	chain           []*Block
	address         string
}

func NewBlockchain(address string) *Blockchain {
	bc := new(Blockchain)
	bc.CreateBlock(0, [32]byte{})
	bc.address = address
	return bc
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	bc.transactionPool = []*Transaction{}
	return b
}

func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}

func (bc *Blockchain) AddTransaction(sender, recipient string, value float32) {
	t := NewTransaction(sender, recipient, value)
	bc.transactionPool = append(bc.transactionPool, t)

}

func (bc *Blockchain) CopyTransactionPool() []*Transaction {
	transactions := make([]*Transaction, 0)
	for _, t := range bc.transactionPool {
		transactions = append(transactions, NewTransaction(t.senderAddress, t.recipientAddress, t.value))
	}
	return transactions
}

func (bc *Blockchain) ValidProof(nonce int, previousHash [32]byte, transactions []*Transaction, difficulty int) bool {
	zeros := strings.Repeat("0", difficulty)
	guessBlock := Block{0, nonce, [32]byte{}, previousHash, transactions}
	guessHashStr := fmt.Sprintf("%x", guessBlock.Hash())
	return guessHashStr[:difficulty] == zeros
}

func (bc *Blockchain) ProofOfWork() int {
	transactions := bc.CopyTransactionPool()
	previousHash := bc.LastBlock().Hash()
	nonce := 0
	for !bc.ValidProof(nonce, previousHash, transactions, MINING_DIFFICULTY) {
		nonce += 1
	}
	return nonce
}

func (bc *Blockchain) Mining() bool {
	bc.AddTransaction(MINING_SENDER, bc.address, MINING_REWARD)
	nonce := bc.ProofOfWork()
	previousHash := bc.LastBlock().Hash()
	bc.CreateBlock(nonce, previousHash)
	log.Println("action=mining, status=success")
	return true
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain - %d %s\n", strings.Repeat("-", 25), i, strings.Repeat("-", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("-", 25))
}
