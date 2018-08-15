package sdk

import (
	"bytes"
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
)

type BlockchainService service

// const routers path for methods BlockchainService
const (
	pathBlockHeight         = "/chain/height"
	pathBlockByHeight       = "/block/%d"
	pathBlockScore          = "/chain/score"
	pathBlockGetTransaction = "/block/%d/transactions"
	pathBlockInfo           = "/blocks/%d/limit/%d"
	pathBlockStorage        = "/diagnostic/storage"
)

// Get the Chain Height
func (b *BlockchainService) GetBlockchainHeight(ctx context.Context) (*ChainHeight, *http.Response, error) {

	bh := &ChainHeight{}
	resp, err := b.client.DoNewRequest(ctx, "GET", pathBlockHeight, nil, &bh)
	if err != nil {
		return nil, resp, err
	}

	return bh, resp, nil
}

// Get the Chain Score
func (b *BlockchainService) GetBlockchainScore(ctx context.Context) (*ChainScore, *http.Response, error) {
	cs := &ChainScore{}
	resp, err := b.client.DoNewRequest(ctx, "GET", pathBlockScore, nil, &cs)
	if err != nil {
		return nil, resp, err
	}

	return cs, resp, nil
}

// Get Block Height
func (b *BlockchainService) GetBlockByHeight(ctx context.Context, height int) (*BlockInfo, *http.Response, error) {
	u := fmt.Sprintf(pathBlockByHeight, height)

	binfo := &BlockInfo{}
	resp, err := b.client.DoNewRequest(ctx, "GET", u, nil, &binfo)
	if err != nil {
		return nil, resp, err
	}

	return binfo, resp, nil
}

// Get Transactions from a block information
func (b *BlockchainService) GetBlockTransactions(ctx context.Context, height int) ([]Transaction, *http.Response, error) {
	u := fmt.Sprintf(pathBlockGetTransaction, height)

	var data bytes.Buffer
	resp, err := b.client.DoNewRequest(ctx, "GET", u, nil, &data)
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
	bstorage := &BlockchainStorageInfo{}
	resp, err := b.client.DoNewRequest(ctx, "GET", pathBlockStorage, nil, &bstorage)
	if err != nil {
		return nil, resp, err
	}

	return bstorage, resp, nil
}

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
