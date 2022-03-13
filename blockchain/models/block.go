package bcmodels

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

type Block struct {
	Nonce int
	PreviousHash [32]byte
	Timestamp int64
	Transactions []*Transaction
}

func NewBlock(nonce int, previousHash [32]byte, transactions []*Transaction) *Block {
	return &Block{
		Nonce: nonce,
		PreviousHash: previousHash,
		Timestamp:  time.Now().UnixNano(),
		Transactions: transactions,
	}
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp int64 `json:"timestamp"`
		Nonce int `json:"nonce"`
		PreviousHash string `json:"previous_hash"`
		Transactions []*Transaction `json:"transactions"`
	}{
		Timestamp: b.Timestamp,
		Nonce: b.Nonce,
		PreviousHash: fmt.Sprintf("%x", b.PreviousHash),
		Transactions: b.Transactions,
	})
}

func (b *Block) ToHash() [32]byte {
	m, _ := json.Marshal(b)

	return sha256.Sum256([]byte(m))
}

// TODO: remove
func (b *Block) Print() {
	fmt.Printf("timestamp       %d\n", b.Timestamp)
	fmt.Printf("nonce           %d\n", b.Nonce)
	fmt.Printf("previous_hash   %x\n", b.PreviousHash)
	for _, t := range b.Transactions {
		t.Print()
	}
}