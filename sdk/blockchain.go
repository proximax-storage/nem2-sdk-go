// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package sdk

import (
	"bytes"
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"math/big"
	"net/http"
)

type BlockchainService service

// Get Block Height
func (b *BlockchainService) GetBlockByHeight(ctx context.Context, height *big.Int) (*BlockInfo, *http.Response, error) {
	u := fmt.Sprintf(pathBlockByHeight, height.String())

	bDto := &blockInfoDTO{}
	resp, err := b.client.DoNewRequest(ctx, "GET", u, nil, &bDto)
	if err != nil {
		return nil, resp, err
	}

	bInfo, err := bDto.toStruct()
	if err != nil {
		return nil, resp, err
	}

	return bInfo, resp, nil
}

// Get Transactions from a block information
func (b *BlockchainService) GetBlockTransactions(ctx context.Context, height *big.Int) ([]Transaction, *http.Response, error) { // TODO Add params
	u := fmt.Sprintf(pathBlockGetTransaction, height.String())

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

// GetBlocksByHeightWithLimit Returns blocks information for a given block height and limit
func (b *BlockchainService) GetBlocksByHeightWithLimit(ctx context.Context, height, limit *big.Int) ([]*BlockInfo, *http.Response, error) {
	if (height.Int64() == 0) || (limit.Int64() == 0) {
		return nil, nil, errors.New("bad parameters - height, limit must be more then 0")
	}

	url := fmt.Sprintf(pathBlockInfo, height.String(), limit.String())

	var bDtos []blockInfoDTO

	resp, err := b.client.DoNewRequest(ctx, "GET", url, nil, &bDtos)
	if err != nil {
		return nil, resp, err
	}

	bInfos := make([]*BlockInfo, limit.Int64())
	for i, bDto := range bDtos {
		bInfos[i], err = bDto.toStruct()
	}
	if err != nil {
		return nil, resp, err
	}

	return bInfos, resp, nil
}

// Get the Chain Height
func (b *BlockchainService) GetBlockchainHeight(ctx context.Context) (*big.Int, *http.Response, error) {
	bh := &struct {
		Height uint64DTO `json:"height"`
	}{}
	resp, err := b.client.DoNewRequest(ctx, "GET", pathBlockHeight, nil, &bh)
	if err != nil {
		return nil, resp, err
	}

	return bh.Height.toBigInt(), resp, nil
}

// Get the Chain Score
func (b *BlockchainService) GetBlockchainScore(ctx context.Context) (*big.Int, *http.Response, error) {
	cs := &chainScoreDTO{}
	resp, err := b.client.DoNewRequest(ctx, "GET", pathBlockScore, nil, &cs)
	if err != nil {
		return nil, resp, err
	}

	return cs.toStruct(), resp, nil
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
