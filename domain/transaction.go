package domain

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Transaction struct {
	senderAddress    string
	recipientAddress string
	value            float32
}

func NewTransaction(sender, recipient string, value float32) *Transaction {
	t := new(Transaction)
	t.senderAddress = sender
	t.recipientAddress = recipient
	t.value = value
	return t
}

func (t *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 25))
	fmt.Printf("		sender_blockchain_address      %s\n", t.senderAddress)
	fmt.Printf("		recipient_blockchain_address   %s\n", t.recipientAddress)
	fmt.Printf("		value                          %.1f\n", t.value)
}

func (t *Transaction) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json:"sender_blockchain_address"`
		Recipient string  `json:"recipient_blockchain_address"`
		Value     float32 `json:"value"`
	}{
		Sender:    t.senderAddress,
		Recipient: t.recipientAddress,
		Value:     t.value,
	})
}
