package sdk

import (
	"bytes"
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
)

type BlockchainService service

// Get the Chain Height
func (b *BlockchainService) GetBlockchainHeight(ctx context.Context) (*ChainHeight, *http.Response, error) {
	req, err := b.client.NewRequest("GET", "chain/height", nil)
	if err != nil {
		return nil, nil, err
	}

	bh := &ChainHeight{}
	resp, err := b.client.Do(ctx, req, &bh)
	if err != nil {
		return nil, resp, err
	}

	return bh, resp, nil
}

// Get the Chain Score
func (b *BlockchainService) GetBlockchainScore(ctx context.Context) (*ChainScore, *http.Response, error) {
	req, err := b.client.NewRequest("GET", "chain/score", nil)
	if err != nil {
		return nil, nil, err
	}

	cs := &ChainScore{}
	resp, err := b.client.Do(ctx, req, &cs)
	if err != nil {
		return nil, resp, err
	}

	return cs, resp, nil
}

const pathBlockByHeight = "/block/%d"

// Get Block Height
func (b *BlockchainService) GetBlockByHeight(ctx context.Context, height int) (*BlockInfo, *http.Response, error) {
	u := fmt.Sprintf(pathBlockByHeight, height)

	req, err := b.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	binfo := &BlockInfo{}
	resp, err := b.client.Do(ctx, req, &binfo)
	if err != nil {
		return nil, resp, err
	}

	return binfo, resp, nil
}

const pathBlockGetTransaction = "block/%d/transactions"

// Get Transactions from a block information
func (b *BlockchainService) GetBlockTransactions(ctx context.Context, height int) ([]Transaction, *http.Response, error) {
	u := fmt.Sprintf(pathBlockGetTransaction, height)

	req, err := b.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var data bytes.Buffer
	resp, err := b.client.Do(ctx, req, &data)
	if err != nil {
		return nil, resp, err
	}

	tx, err := MapTransactions(&data)
	if err != nil {
		return nil, resp, err
	}

	return tx, resp, nil
}

// Get the Storage Information
func (b *BlockchainService) GetBlockchainStorage(ctx context.Context) (*BlockchainStorageInfo, *http.Response, error) {
	req, err := b.client.NewRequest("GET", "diagnostic/storage", nil)
	if err != nil {
		return nil, nil, err
	}

	bstorage := &BlockchainStorageInfo{}
	resp, err := b.client.Do(ctx, req, &bstorage)
	if err != nil {
		return nil, resp, err
	}

	return bstorage, resp, nil
}

const pathBlockInfo = "/blocks/%d/limit/%d"

//GetBlockchainInfo Returns blocks information for a given block height and limit
func (b *BlockchainService) GetBlockchainInfo(ctx context.Context, height, limit int) (*BlockInfo, *http.Response, error) {

	if (height < 0) || (limit < 0) {
		return nil, nil, errors.New("bad parameters - height, limit must be more then 0")
	}

	url := fmt.Sprintf(pathBlockInfo, height, limit)
	bcInfo := &BlockInfo{}

	resp, err := b.client.DoNewRequest(ctx, "GET", url, nil, &bcInfo)
	if err != nil {
		return nil, nil, err
	}

	return bcInfo, resp, nil
}
