package models

import (
	"fmt"
	"strings"
)

type Blockchain struct {
	transactionPool []*Transaction
	chain []*Block
	blockchainAddress string
}

const (
	MINING_DIFFICULTY = 3
	MINING_SENDER = "THE BLOCKCHAIN NODE"
	MINING_REWARD = 1.0
)

func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash, bc.transactionPool)
	bc.chain = append(bc.chain, b)

	bc.transactionPool = []*Transaction{}

	return b
}

func NewBlockchain(blockchainAddress string) *Blockchain {
	b := &Block{}
	bc := &Blockchain{blockchainAddress: blockchainAddress}
	bc.CreateBlock(0, b.ToHash())

	return bc
}

func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain) - 1]
}

func (bc *Blockchain) AddTransaction(sender string, recipient string, value float32) {
	t := NewTransaction(sender, recipient, value)
	bc.transactionPool = append(bc.transactionPool, t)
}

func (bc *Blockchain) IsValidProof(nonce int, previousHash [32]byte, transactions []*Transaction, difficulty int) bool {
	zeros := strings.Repeat("0", difficulty)
	guessBlock := Block{Nonce: nonce, PreviousHash: previousHash, Timestamp: 0, Transactions: transactions}
	guessHashStr := fmt.Sprintf("%x", guessBlock.ToHash())

	return guessHashStr[:difficulty] == zeros
}

func (bc *Blockchain) ProofOfWork() int {
	nonce := 0
	previousHash := bc.LastBlock().ToHash()

	for !bc.IsValidProof(nonce, previousHash, bc.transactionPool, MINING_DIFFICULTY) {
		nonce += 1
	}

	return nonce
}

func (bc *Blockchain) Mining() bool {
	bc.AddTransaction(MINING_SENDER, bc.blockchainAddress, MINING_REWARD)

	nonce := bc.ProofOfWork()
	previousHash := bc.LastBlock().ToHash()
	bc.CreateBlock(nonce, previousHash)

	return true
}

func (bc *Blockchain) CalcTotalAmount(blockchainAddress string) float32 {
	var totalAmount float32 = 0.0

	for _, b := range bc.chain {
		for _, t := range b.Transactions {
			if t.IsRecipient(blockchainAddress) {
				totalAmount += t.Value
			}

			if t.IsSender(blockchainAddress) {
				totalAmount -= t.Value
			}
		}
	}

	return totalAmount
}