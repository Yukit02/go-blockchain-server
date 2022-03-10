package models

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