package wlmodels

import (
	"bcserver/utils"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"log"
)

type Transaction struct {
	SenderPrivateKey *ecdsa.PrivateKey
	SenderPublicKey *ecdsa.PublicKey
	SenderBlockchainAddress string
	RecipientBlockchainAddress string
	Value float32
}

func NewTransaction(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey, senderAdress string, recipientAddress string, value float32) *Transaction {
	return &Transaction {
		SenderPrivateKey: privateKey,
		SenderPublicKey: publicKey,
		SenderBlockchainAddress: senderAdress,
		RecipientBlockchainAddress: recipientAddress,
		Value: value,
	}
}

func (t *Transaction) GenerateSignature() (*utils.Signature, error) {
	m, err := json.Marshal(t)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	h := sha256.Sum256(m)
	r, s, err := ecdsa.Sign(rand.Reader, t.SenderPrivateKey, h[:])
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &utils.Signature{R: r, S: s}, nil
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender string `json:"sender_blockchain_address"`
		Recipient string `json:"recipient_blockchain_address"`
		Value float32 `json:"value"`
	}{
		Sender: t.SenderBlockchainAddress,
		Recipient: t.RecipientBlockchainAddress,
		Value: t.Value,
	})
}