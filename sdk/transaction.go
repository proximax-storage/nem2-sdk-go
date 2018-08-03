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

var mainTransactionRoute = "transaction" // TODO We should consider about connecting service with it's route somehow

// Returns transaction information for a given transaction id or hash
func (txs *TransactionService) GetTransaction(ctx context.Context, id string) (Transaction, *http.Response, error) {
	var b bytes.Buffer

	resp, err := txs.client.DoNewRequest(ctx, "GET", fmt.Sprintf("%s/%s", mainTransactionRoute, id), nil, &b)
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
	return txs.announceTransaction(ctx, tx, fmt.Sprintf("%s/partial", mainTransactionRoute))
}

// Announce a cosignature transaction to the network
func (txs *TransactionService) AnnounceAggregateBondedCosignature(ctx context.Context, c *CosignatureSignedTransaction) (string, *http.Response, error) {
	return txs.announceTransaction(ctx, c, fmt.Sprintf("%s/cosignature", mainTransactionRoute))
}

// Returns transaction status for a given transaction id or hash
func (txs *TransactionService) GetTransactionStatus(ctx context.Context, id string) (*TransactionStatus, *http.Response, error) {
	ts := &TransactionStatus{}

	resp, err := txs.client.DoNewRequest(ctx, "GET", fmt.Sprintf("%s/%s/status", mainTransactionRoute, id), nil, ts)
	if err != nil {
		return nil, resp, err
	}

	return ts, resp, nil
}

// Returns transaction information for a given set of transaction hash
func (txs *TransactionService) GetTransactionStatuses(ctx context.Context, hashes []string) ([]*TransactionStatus, *http.Response, error) {
	txIds := &TransactionHashesDTO{
		hashes,
	}

	s := make([]*TransactionStatus, len(hashes))
	resp, err := txs.client.DoNewRequest(ctx, "POST", fmt.Sprintf("%s/statuses", mainTransactionRoute), txIds, &s)
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

func MapTransactions(b *bytes.Buffer) ([]Transaction, error) {
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
			tx[i], err = MapTransaction(bytes.NewBuffer([]byte(t)))
		}(i, t)
	}
	wg.Wait()

	if err != nil {
		return nil, err
	}

	return tx, nil
}

func MapTransaction(b *bytes.Buffer) (Transaction, error) {
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
		dto := mosaicDefinitionTransactionDTO{}

		err := json.Unmarshal(b.Bytes(), &dto)
		if err != nil {
			return nil, err
		}

		tx, err := dto.toStruct()
		if err != nil {
			return nil, err
		}

		return tx, nil
	case MOSAIC_SUPPLY_CHANGE:
		dto := mosaicSupplyChangeTransactionDTO{}

		err := json.Unmarshal(b.Bytes(), &dto)
		if err != nil {
			return nil, err
		}

		tx, err := dto.toStruct()
		if err != nil {
			return nil, err
		}

		return tx, nil
	case MODIFY_MULTISIG_ACCOUNT:
		dto := modifyMultisigAccountTransactionDTO{}

		err := json.Unmarshal(b.Bytes(), &dto)
		if err != nil {
			return nil, err
		}

		tx, err := dto.toStruct()
		if err != nil {
			return nil, err
		}

		return tx, nil
	case REGISTER_NAMESPACE:
		dto := registerNamespaceTransactionDTO{}

		err := json.Unmarshal(b.Bytes(), &dto)
		if err != nil {
			return nil, err
		}

		tx, err := dto.toStruct()
		if err != nil {
			return nil, err
		}

		return tx, nil
	case TRANSFER:
		dto := transferTransactionDTO{}

		err := json.Unmarshal(b.Bytes(), &dto)
		if err != nil {
			return nil, err
		}

		tx, err := dto.toStruct()
		if err != nil {
			return nil, err
		}

		return tx, nil
	case LOCK:
		dto := lockFundsTransactionDTO{}

		err := json.Unmarshal(b.Bytes(), &dto)
		if err != nil {
			return nil, err
		}

		tx, err := dto.toStruct()
		if err != nil {
			return nil, err
		}

		return tx, nil
	case SECRET_LOCK:
		dto := secretLockTransactionDTO{}

		err := json.Unmarshal(b.Bytes(), &dto)
		if err != nil {
			return nil, err
		}

		tx, err := dto.toStruct()
		if err != nil {
			return nil, err
		}

		return tx, nil
	case SECRET_PROOF:
		dto := secretProofTransactionDTO{}

		err := json.Unmarshal(b.Bytes(), &dto)
		if err != nil {
			return nil, err
		}

		tx, err := dto.toStruct()
		if err != nil {
			return nil, err
		}

		return tx, nil
	}

	return nil, nil
}

func mapAggregateTransaction(b *bytes.Buffer, t TransactionType) (*AggregateTransaction, error) {
	dto := aggregateTransactionDTO{}

	err := json.Unmarshal(b.Bytes(), &dto)
	if err != nil {
		return nil, err
	}

	tx, err := dto.toStruct()
	if err != nil {
		return nil, err
	}

	return tx, nil
}
