package bcmodels

import (
	"fmt"
	"strings"
)

type Transaction struct {
	SenderBlockchainAddress string `json:"sender_blockchain_address"`
	RecipientBlockchainAddress string `json:"recipient_blockchain_address"`
	Value float32 `json:"value"`
}

func NewTransaction(sender string, recipient string, value float32) *Transaction {
	return &Transaction{SenderBlockchainAddress: sender, RecipientBlockchainAddress: recipient, Value: value}
}

func (t *Transaction) IsRecipient(blockchainAddress string) bool {
	return t.RecipientBlockchainAddress == blockchainAddress
}

func (t *Transaction) IsSender(blockchainAddress string) bool {
	return t.SenderBlockchainAddress == blockchainAddress
}

// TODO: remove
func (t *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf(" sender_blockchain_address      %s\n", t.SenderBlockchainAddress)
	fmt.Printf(" recipient_blockchain_address   %s\n", t.RecipientBlockchainAddress)
	fmt.Printf(" value                          %.1f\n", t.Value)
}