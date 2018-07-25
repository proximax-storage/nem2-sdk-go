package sdk

import (
	"context"
	"net/http"
	"fmt"
)

type TransactionService service

// Models
// Transaction Status
type TransactionStatus struct {
	Group    string   `json:"group"`
	Status   string   `json:"status"`
	Hash     string   `json:"hash"`
	Deadline []uint64 `json:"deadline"`
	Height   []uint64 `json:"height"`
}

// Transaction
type Transaction struct {
	Type            TransactionType `json:"transaction_type"`
	NetworkType     NetworkType     `json:"network_type"`
	Version         *uint64         `json:"version"`
	Fee             *uint64         `json:"fee"`
	Deadline        []uint64        `json:"deadline"`
	Signature       string          `json:"signature"`
	Signer          PublicAccount   `json:"signer"`
	TransactionInfo TransactionInfo `json:"transaction_info"`
}

// Transaction Info
type TransactionInfo struct {
	Height              *uint64 `json:"height"`
	Index               *uint32 `json:"index"`
	Id                  string  `json:"id"`
	Hash                string  `json:"hash"`
	MerkleComponentHash string  `json:"merkle_component_hash"`
	AggregateHash       string  `json:"aggregate_hash"`
	AggregateId         string  `json:"aggregate_id"`
}

//Get the status of the transaction by id
func (tx *TransactionService) GetTransactionStatus(ctx context.Context, hash string) (*TransactionStatus, *http.Response, error) {
	req, err := tx.client.NewRequest("GET", fmt.Sprintf("transaction/%s", hash), nil)
	if err != nil {
		return nil, nil, err
	}

	ts := &TransactionStatus{}
	resp, err := tx.client.Do(ctx, req, ts)
	if err != nil {
		return nil, resp, err
	}

	return ts, resp, nil
}