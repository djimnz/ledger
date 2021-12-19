package core

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

// Transactions struct
type Transactions struct {
	Transactions []Transaction `json:"transactions" binding:"required,dive"`
}

// Transaction struct
type Transaction struct {
	ID        int64    `json:"txid"`
	Postings  Postings `json:"postings" binding:"required,dive"`
	Reference string   `json:"reference"`
	Timestamp string   `json:"timestamp"`
	Hash      string   `json:"hash" swaggerignore:"true"`
	Metadata  Metadata `json:"metadata" swaggertype:"object"`
}

// AppendPosting
func (t *Transaction) AppendPosting(p Posting) {
	t.Postings = append(t.Postings, p)
}

// Reverse
func (t *Transaction) Reverse() Transaction {
	postings := t.Postings
	postings.Reverse()

	return Transaction{
		Postings:  postings,
		Reference: "revert_" + t.Reference,
	}
}

// Hash
func Hash(t1 *Transaction, t2 *Transaction) string {
	b1, _ := json.Marshal(t1)
	b2, _ := json.Marshal(t2)

	h := sha256.New()
	h.Write(b1)
	h.Write(b2)

	return fmt.Sprintf("%x", h.Sum(nil))
}
