package bcmodels

import (
	"crypto/sha256"
	"encoding/json"
	"time"
)

type Block struct {
	Nonce int `json:"nonce"`
	PreviousHash [32]byte `json:"previous_hash"`
	Timestamp int64 `json:"timestamp"`
	Transactions []*Transaction `json:"transactions"`
}

func NewBlock(nonce int, previousHash [32]byte, transactions []*Transaction) *Block {
	return &Block{
		Nonce: nonce,
		PreviousHash: previousHash,
		Timestamp:  time.Now().UnixNano(),
		Transactions: transactions,
	}
}

func (b *Block) ToHash() [32]byte {
	m, _ := json.Marshal(b)

	return sha256.Sum256([]byte(m))
}