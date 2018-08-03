package sdk

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
)

type TransactionService service

var mainRoute = "transaction" // TODO We should consider about connecting service with it's route somehow

// Returns transaction information for a given transaction id or hash
func (txs *TransactionService) GetTransaction(ctx context.Context, id string) (Transaction, *http.Response, error) {
	var b bytes.Buffer

	resp, err := txs.client.DoNewRequest(ctx, "GET", fmt.Sprintf("%s/%s", mainRoute, id), nil, &b)
	if err != nil {
		return nil, nil, err
	}

	tx, err := MapTransaction(&b)
	if err != nil {
		return nil, resp, err
	}

	return tx, resp, nil
}

// Returns transaction information for a given set of transaction id or hash
func (txs *TransactionService) GetTransactions(ctx context.Context, ids []string) ([]Transaction, *http.Response, error) {
	var b bytes.Buffer
	txIds := &TransactionIds{
		ids,
	}

	resp, err := txs.client.DoNewRequest(ctx, "POST", mainRoute, txIds, &b)
	if err != nil {
		return nil, resp, err
	}

	tx, err := MapTransactions(&b)
	if err != nil {
		return nil, resp, err
	}

	return tx, resp, nil
}

// Announce a transaction to the network
func (txs *TransactionService) Announce(ctx context.Context, tx *SignedTransaction) (string, *http.Response, error) {
	return txs.announceTransaction(ctx, tx, mainRoute)
}

// Announce a partial transaction to the network
func (txs *TransactionService) AnnounceAggregateBonded(ctx context.Context, tx *SignedTransaction) (string, *http.Response, error) {
	return txs.announceTransaction(ctx, tx, fmt.Sprintf("%s/partial", mainRoute))
}

// Announce a cosignature transaction to the network
func (txs *TransactionService) AnnounceAggregateBondedCosignature(ctx context.Context, c *CosignatureSignedTransaction) (string, *http.Response, error) {
	return txs.announceTransaction(ctx, c, fmt.Sprintf("%s/cosignature", mainRoute))
}

// Returns transaction status for a given transaction id or hash
func (txs *TransactionService) GetTransactionStatus(ctx context.Context, id string) (*TransactionStatus, *http.Response, error) {
	ts := &TransactionStatus{}

	resp, err := txs.client.DoNewRequest(ctx, "GET", fmt.Sprintf("%s/%s/status", mainRoute, id), nil, ts)
	if err != nil {
		return nil, resp, err
	}

	return ts, resp, nil
}

// Returns transaction information for a given set of transaction hash
func (txs *TransactionService) GetTransactionStatuses(ctx context.Context, hashes []string) ([]*TransactionStatus, *http.Response, error) {
	txIds := &TransactionHashes{
		hashes,
	}

	s := make([]*TransactionStatus, len(hashes))
	resp, err := txs.client.DoNewRequest(ctx, "POST", fmt.Sprintf("%s/statuses", mainRoute), txIds, &s)
	if err != nil {
		return nil, resp, err
	}

	return s, resp, nil
}

func (txs *TransactionService) announceTransaction(ctx context.Context, tx Signed, path string) (string, *http.Response, error) {
	var m string
	resp, err := txs.client.DoNewRequest(ctx, "PUT", path, tx, m)
	if err != nil {
		return "", nil, err
	}

	return m, resp, nil
}
