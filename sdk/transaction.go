package sdk

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
)

type TransactionService service

const (
	mainTransactionRoute               = "transaction"
	announceAggreagateRoute            = "partial"
	announceAggreagateCosignatureRoute = "cosignature"
	transactionStatusRoute             = "status"
	transactionStatusesRoute           = "statuses"
)

// Returns transaction information for a given transaction id or hash
func (txs *TransactionService) GetTransaction(ctx context.Context, id string) (Transaction, *http.Response, error) {
	var b bytes.Buffer

	resp, err := txs.client.DoNewRequest(ctx, "GET", fmt.Sprintf("%s/%s", mainTransactionRoute, id), nil, &b)
	if err != nil {
		return nil, resp, err
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
	txIds := &TransactionIdsDTO{
		ids,
	}

	resp, err := txs.client.DoNewRequest(ctx, "POST", mainTransactionRoute, txIds, &b)
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
	return txs.announceTransaction(ctx, tx, mainTransactionRoute)
}

// Announce a partial transaction to the network
func (txs *TransactionService) AnnounceAggregateBonded(ctx context.Context, tx *SignedTransaction) (string, *http.Response, error) {
	return txs.announceTransaction(ctx, tx, fmt.Sprintf("%s/%s", mainTransactionRoute, announceAggreagateRoute))
}

// Announce a cosignature transaction to the network
func (txs *TransactionService) AnnounceAggregateBondedCosignature(ctx context.Context, c *CosignatureSignedTransaction) (string, *http.Response, error) {
	return txs.announceTransaction(ctx, c, fmt.Sprintf("%s/%s", mainTransactionRoute, announceAggreagateCosignatureRoute))
}

// Returns transaction status for a given transaction id or hash
func (txs *TransactionService) GetTransactionStatus(ctx context.Context, id string) (*TransactionStatus, *http.Response, error) {
	ts := &transactionStatusDTO{}

	resp, err := txs.client.DoNewRequest(ctx, "GET", fmt.Sprintf("%s/%s/%s", mainTransactionRoute, id, transactionStatusRoute), nil, ts)
	if err != nil {
		return nil, resp, err
	}

	tx, err := ts.toStruct()
	if err != nil {
		return nil, resp, err
	}

	return tx, resp, nil
}

// Returns transaction information for a given set of transaction hash
func (txs *TransactionService) GetTransactionStatuses(ctx context.Context, hashes []string) ([]*TransactionStatus, *http.Response, error) {
	txIds := &TransactionHashesDTO{
		hashes,
	}

	dtos := make([]*transactionStatusDTO, len(hashes))
	resp, err := txs.client.DoNewRequest(ctx, "POST", fmt.Sprintf("%s/%s", mainTransactionRoute, transactionStatusesRoute), txIds, &dtos)
	if err != nil {
		return nil, resp, err
	}

	tss := make([]*TransactionStatus, len(dtos))
	for i, ts := range dtos {
		tss[i], err = ts.toStruct()
	}
	if err != nil {
		return nil, resp, err
	}

	return tss, resp, nil
}

func (txs *TransactionService) announceTransaction(ctx context.Context, tx Signed, path string) (string, *http.Response, error) {
	var m string
	resp, err := txs.client.DoNewRequest(ctx, "PUT", path, tx, m)
	if err != nil {
		return "", resp, err
	}

	return m, resp, nil
}
