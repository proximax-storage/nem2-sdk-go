package sdk

import (
	"bytes"
	"context"
	jsonLib "encoding/json"
	"fmt"
	"net/http"
	"sync"
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

	tx, err := txs.mapTransaction(&b)
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

	tx, err := txs.mapTransactions(&b)
	if err != nil {
		return nil, resp, err
	}

	return tx, resp, nil
}

// Announce a transaction to the network
func (txs *TransactionService) Announce(ctx context.Context, tx SignedTransaction) (string, *http.Response, error) {
	return txs.announceTransaction(ctx, tx, mainRoute)
}

// Announce a partial transaction to the network
func (txs *TransactionService) AnnounceAggregateBonded(ctx context.Context, tx SignedTransaction) (string, *http.Response, error) {
	return txs.announceTransaction(ctx, tx, fmt.Sprintf("%s/partial", mainRoute))
}

// Announce a cosignature transaction to the network
func (txs *TransactionService) AnnounceAggregateBondedCosignature(ctx context.Context, c CosignatureSignedTransaction) (string, *http.Response, error) {
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

func (txs *TransactionService) announceTransaction(ctx context.Context, tx Transaction, path string) (string, *http.Response, error) {
	var m string
	resp, err := txs.client.DoNewRequest(ctx, "PUT", path, tx, m)
	if err != nil {
		return "", nil, err
	}

	return m, resp, nil
}

func (txs *TransactionService) mapTransactions(b *bytes.Buffer) ([]Transaction, error) {
	var wg sync.WaitGroup
	var err error

	m := []jsonLib.RawMessage{}

	json.Unmarshal(b.Bytes(), &m)

	tx := make([]Transaction, len(m))
	for i, t := range m {
		wg.Add(1)
		go func(i int, t jsonLib.RawMessage) {
			defer wg.Done()
			json.Marshal(t)
			tx[i], err = txs.mapTransaction(bytes.NewBuffer([]byte(t)))
		}(i, t)
	}
	wg.Wait()

	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (txs *TransactionService) mapTransaction(b *bytes.Buffer) (Transaction, error) {
	rawT := struct {
		Transaction struct {
			Type uint32
		}
	}{}
	err := json.Unmarshal(b.Bytes(), &rawT)
	if err != nil {
		return nil, err
	}

	t, err := TransactionTypeFromRaw(rawT.Transaction.Type)
	if err != nil {
		return nil, err
	}

	switch t {
	case AGGREGATE_BONDED:
		mapAggregateTransaction(b, AGGREGATE_BONDED)
	case AGGREGATE_COMPLETE:
		mapAggregateTransaction(b, AGGREGATE_COMPLETE)
	case MOSAIC_DEFINITION:
		rawTx := struct {
			Tx struct {
				MosaicDefinitionTransaction
			} `json:"transaction"`
		}{}

		err := json.Unmarshal(b.Bytes(), &rawTx)
		if err != nil {
			return nil, err
		}

		tx := rawTx.Tx.MosaicDefinitionTransaction

		aTx, err := mapAbstractTransaction(b, MOSAIC_DEFINITION)
		if err != nil {
			return nil, err
		}

		tx.AbstractTransaction = *aTx

		return &tx, nil
	case MOSAIC_SUPPLY_CHANGE:
		rawTx := struct {
			Tx struct {
				MosaicSupplyChangeTransaction
			} `json:"transaction"`
		}{}

		err := json.Unmarshal(b.Bytes(), &rawTx)
		if err != nil {
			return nil, err
		}

		tx := rawTx.Tx.MosaicSupplyChangeTransaction

		aTx, err := mapAbstractTransaction(b, MOSAIC_SUPPLY_CHANGE)
		if err != nil {
			return nil, err
		}

		tx.AbstractTransaction = *aTx

		return &tx, nil
	case MODIFY_MULTISIG_ACCOUNT:
		rawTx := struct {
			Tx struct {
				ModifyMultisigAccountTransaction
			} `json:"transaction"`
		}{}

		err := json.Unmarshal(b.Bytes(), &rawTx)
		if err != nil {
			return nil, err
		}

		tx := rawTx.Tx.ModifyMultisigAccountTransaction

		aTx, err := mapAbstractTransaction(b, MODIFY_MULTISIG_ACCOUNT)
		if err != nil {
			return nil, err
		}

		tx.AbstractTransaction = *aTx

		return &tx, nil
	case REGISTER_NAMESPACE:
		rawTx := struct {
			Tx struct {
				RegisterNamespaceTransaction
			} `json:"transaction"`
		}{}

		err := json.Unmarshal(b.Bytes(), &rawTx)
		if err != nil {
			return nil, err
		}

		tx := rawTx.Tx.RegisterNamespaceTransaction

		aTx, err := mapAbstractTransaction(b, REGISTER_NAMESPACE)
		if err != nil {
			return nil, err
		}

		tx.AbstractTransaction = *aTx

		return &tx, nil
	case TRANSFER:
		rawTx := struct {
			Tx struct {
				TransferTransaction
			} `json:"transaction"`
		}{}

		err := json.Unmarshal(b.Bytes(), &rawTx)
		if err != nil {
			return nil, err
		}

		tx := rawTx.Tx.TransferTransaction

		aTx, err := mapAbstractTransaction(b, TRANSFER)
		if err != nil {
			return nil, err
		}

		tx.AbstractTransaction = *aTx

		return &tx, nil
	case LOCK:
		rawTx := struct {
			Tx struct {
				LockFundsTransaction
			} `json:"transaction"`
		}{}

		err := json.Unmarshal(b.Bytes(), &rawTx)
		if err != nil {
			return nil, err
		}

		tx := rawTx.Tx.LockFundsTransaction

		aTx, err := mapAbstractTransaction(b, LOCK)
		if err != nil {
			return nil, err
		}

		tx.AbstractTransaction = *aTx

		return &tx, nil
	case SECRET_LOCK:
		rawTx := struct {
			Tx struct {
				SecretLockTransaction
			} `json:"transaction"`
		}{}

		err := json.Unmarshal(b.Bytes(), &rawTx)
		if err != nil {
			return nil, err
		}

		tx := rawTx.Tx.SecretLockTransaction

		aTx, err := mapAbstractTransaction(b, SECRET_LOCK)
		if err != nil {
			return nil, err
		}

		tx.AbstractTransaction = *aTx

		return &tx, nil
	case SECRET_PROOF:
		rawTx := struct {
			Tx struct {
				SecretProofTransaction
			} `json:"transaction"`
		}{}

		err := json.Unmarshal(b.Bytes(), &rawTx)
		if err != nil {
			return nil, err
		}

		tx := rawTx.Tx.SecretProofTransaction

		aTx, err := mapAbstractTransaction(b, SECRET_PROOF)
		if err != nil {
			return nil, err
		}

		tx.AbstractTransaction = *aTx

		return &tx, nil
	}

	return nil, nil
}

func mapAbstractTransaction(b *bytes.Buffer, t TransactionType) (*AbstractTransaction, error) {
	rawTx := struct {
		Tx struct {
			AbstractTransaction
		} `json:"transaction"`
		TransactionInfo `json:"meta"`
	}{}

	err := json.Unmarshal(b.Bytes(), &rawTx)
	if err != nil {
		return nil, err
	}

	aTx := rawTx.Tx.AbstractTransaction
	aTx.Type = t
	aTx.TransactionInfo = rawTx.TransactionInfo

	//nt, err := extractNetworkType(aTx.Version)
	//if err != nil {
	//	return nil, err
	//}
	//
	//aTx.NetworkType = nt
	return &aTx, nil
}

func mapAggregateTransaction(b *bytes.Buffer, t TransactionType) (*AggregateTransaction, error) {
	rawTx := struct {
		Tx struct {
			AggregateTransaction
		} `json:"transaction"`
	}{}

	err := json.Unmarshal(b.Bytes(), &rawTx)
	if err != nil {
		return nil, err
	}

	tx := rawTx.Tx.AggregateTransaction

	aTx, err := mapAbstractTransaction(b, t)
	if err != nil {
		return nil, err
	}

	tx.AbstractTransaction = *aTx

	return &tx, nil
}
