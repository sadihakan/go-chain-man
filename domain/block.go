package domain

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

const (
	MINING_DIFFICULTY = 3
	MINING_SENDER     = "THE BLOCKCHAIN"
	MINING_REWARD     = 1.0
)

type Block struct {
	timestamp    int64
	nonce        int
	hash         [32]byte
	previousHash [32]byte
	transactions []*Transaction
}

func NewBlock(nonce int, previousHash [32]byte, transaction []*Transaction) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	b.hash = b.Hash()
	b.transactions = transaction
	return b
}

func (b *Block) Hash() [32]byte {
	m, _ := b.ToJSON()
	return sha256.Sum256([]byte(m))
}

func (b *Block) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64          `json:"timestamp"`
		Nonce        int            `json:"nonce"`
		PreviousHash [32]byte       `json:"previous_hash"`
		Transactions []*Transaction `json:"transactions"`
	}{
		Timestamp:    b.timestamp,
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Transactions: b.transactions,
	})
}

func (b *Block) Print() {
	fmt.Printf("	timestamp        %d\n", b.timestamp)
	fmt.Printf("	nonce            %d\n", b.nonce)
	fmt.Printf("	hash             %x\n", b.hash)
	fmt.Printf("	previous_hash    %x\n", b.previousHash)
	for _, t := range b.transactions {
		t.Print()
	}
}
