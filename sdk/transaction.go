package sdk

import (
	"context"
	"net/http"
	"fmt"
	"encoding/json"
)

type TransactionService service

var mainRoute = "transaction" // TODO We should consider about connecting service with it's route somehow

// Returns transaction information for a given transaction id or hash
func (txs *TransactionService) GetTransaction(ctx context.Context, id string) (*Transaction, *http.Response, error) {
	req, err := txs.client.NewRequest("GET", fmt.Sprintf("%s/%s", mainRoute, id), nil)
	if err != nil {
		return nil, nil, err
	}
	tx := &Transaction{}
	resp, err := txs.client.Do(ctx, req, tx)
	if err != nil {
		return nil, resp, err
	}

	return tx, resp, nil
}

// Returns transaction information for a given set of transaction id or hash
func (txs *TransactionService) GetTransactions(ctx context.Context, ids []string) ([]*Transaction, *http.Response, error) {
	txIds := &TransactionIds{
		ids,
	}

	idsJson, err := json.Marshal(txIds)
	if err != nil {
		return nil, nil, err
	}

	req, err := txs.client.NewRequest("POST", mainRoute, string(idsJson))
	if err != nil {
		return nil, nil, err
	}

	i := make([]*Transaction, len(ids))
	resp, err := txs.client.Do(ctx, req, i)
	if err != nil {
		return nil, resp, err
	}

	return i, resp, nil
}

// Announce a transaction to the network
func (txs *TransactionService) Announce(ctx context.Context, tx SignedTransaction) (string, *http.Response, error) {
	return txs.announceTransaction(ctx, tx, mainRoute)
}

// Announce a partial transaction to the network
func (txs *TransactionService) AnnounceAggregateBonded(ctx context.Context, tx SignedTransaction) (string, *http.Response, error) {
	return txs.announceTransaction(ctx, tx, fmt.Sprintf("%s/partial", mainRoute))
}

// Returns transaction status for a given transaction id or hash
func (txs *TransactionService) GetTransactionStatus(ctx context.Context, id string) (*TransactionStatus, *http.Response, error) {
	req, err := txs.client.NewRequest("GET", fmt.Sprintf("%s/%s/status", mainRoute, id), nil)
	if err != nil {
		return nil, nil, err
	}

	ts := &TransactionStatus{}
	resp, err := txs.client.Do(ctx, req, ts)
	if err != nil {
		return nil, resp, err
	}

	return ts, resp, nil
}

// Returns transaction information for a given set of transaction id or hash
func (txs *TransactionService) GetTransactionStatuses(ctx context.Context, ids []string) ([]*TransactionStatus, *http.Response, error) {
	txIds := &TransactionIds{
		ids,
	}

	idsJson, err := json.Marshal(txIds)
	if err != nil {
		return nil, nil, err
	}

	req, err := txs.client.NewRequest("POST", fmt.Sprintf("%s/statuses", mainRoute), string(idsJson))
	if err != nil {
		return nil, nil, err
	}

	s := make([]*TransactionStatus, len(ids))
	resp, err := txs.client.Do(ctx, req, s)
	if err != nil {
		return nil, resp, err
	}

	return s, resp, nil
}

func (txs *TransactionService) announceTransaction(ctx context.Context, tx SignedTransaction, path string) (string, *http.Response, error) {
	jsonTx, err := json.Marshal(tx)
	if err != nil {
		return "", nil, err
	}

	req, err := txs.client.NewRequest("PUT", path, string(jsonTx))
	if err != nil {
		return "", nil, err
	}

	var message string
	resp, err := txs.client.Do(ctx, req, message)
	if err != nil {
		return "", nil, err
	}

	return message, resp, nil
}