// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package sdk

import (
	"bytes"
	"context"
	"fmt"
	"github.com/proximax-storage/proximax-utils-go/net"
	"net/http"
)

type TransactionService service

// Returns transaction information for a given transaction id or hash
func (txs *TransactionService) GetTransaction(ctx context.Context, id string) (Transaction, error) {
	var b bytes.Buffer

	url := net.NewUrl(fmt.Sprintf("/"+mainTransactionRoute+"/%s", id))

	resp, err := txs.client.DoNewRequest(ctx, http.MethodGet, url.Encode(), nil, &b)
	if err != nil {
		return nil, err
	}

	if err = handleResponseStatusCode(resp, map[int]error{404: ErrResourceNotFound, 409: ErrArgumentNotValid}); err != nil {
		return nil, err
	}

	return MapTransaction(&b)
}

// Returns transaction information for a given set of transaction id or hash
func (txs *TransactionService) GetTransactions(ctx context.Context, ids []string) ([]Transaction, error) {
	var b bytes.Buffer
	txIds := &TransactionIdsDTO{
		ids,
	}

	resp, err := txs.client.DoNewRequest(ctx, http.MethodPost, mainTransactionRoute, txIds, &b)
	if err != nil {
		return nil, err
	}

	if err = handleResponseStatusCode(resp, map[int]error{400: ErrInvalidRequest, 409: ErrArgumentNotValid}); err != nil {
		return nil, err
	}

	return MapTransactions(&b)
}

// Announce a transaction to the network
func (txs *TransactionService) Announce(ctx context.Context, tx *SignedTransaction) (string, error) {
	return txs.announceTransaction(ctx, tx, mainTransactionRoute)
}

// Announce a partial transaction to the network
func (txs *TransactionService) AnnounceAggregateBonded(ctx context.Context, tx *SignedTransaction) (string, error) {
	return txs.announceTransaction(ctx, tx, fmt.Sprintf("%s/%s", mainTransactionRoute, announceAggregateRoute))
}

// Announce a cosignature transaction to the network
func (txs *TransactionService) AnnounceAggregateBondedCosignature(ctx context.Context, c *CosignatureSignedTransaction) (string, error) {
	return txs.announceTransaction(ctx, c, fmt.Sprintf("%s/%s", mainTransactionRoute, announceAggregateCosignatureRoute))
}

// Returns transaction status for a given transaction id or hash
func (txs *TransactionService) GetTransactionStatus(ctx context.Context, id string) (*TransactionStatus, error) {
	ts := &transactionStatusDTO{}

	resp, err := txs.client.DoNewRequest(ctx, http.MethodGet, fmt.Sprintf("%s/%s/%s", mainTransactionRoute, id, transactionStatusRoute), nil, ts)
	if err != nil {
		return nil, err
	}

	if err = handleResponseStatusCode(resp, map[int]error{404: ErrResourceNotFound, 409: ErrArgumentNotValid}); err != nil {
		return nil, err
	}

	return ts.toStruct()
}

// Returns transaction information for a given set of transaction hash
func (txs *TransactionService) GetTransactionStatuses(ctx context.Context, hashes []string) ([]*TransactionStatus, error) {
	txIds := &TransactionHashesDTO{
		hashes,
	}

	url := net.NewUrl("/" + mainTransactionRoute + "/" + transactionStatusesRoute)

	dtos := make([]*transactionStatusDTO, len(hashes))
	resp, err := txs.client.DoNewRequest(ctx, http.MethodPost, url.Encode(), txIds, &dtos)
	if err != nil {
		return nil, err
	}

	if err = handleResponseStatusCode(resp, map[int]error{404: ErrResourceNotFound, 409: ErrArgumentNotValid}); err != nil {
		return nil, err
	}

	return transactionStatusDTOsToTransactionStatuses(dtos)
}

func (txs *TransactionService) announceTransaction(ctx context.Context, tx interface{}, path string) (string, error) {
	m := struct {
		Message string `json:"message"`
	}{}

	resp, err := txs.client.DoNewRequest(ctx, http.MethodPut, path, tx, &m)
	if err != nil {
		return "", err
	}

	if err = handleResponseStatusCode(resp, map[int]error{400: ErrInvalidRequest, 409: ErrArgumentNotValid}); err != nil {
		return "", err
	}

	return m.Message, nil
}
