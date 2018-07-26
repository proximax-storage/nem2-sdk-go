package sdk

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
)

type BlockchainService service

// Get the Chain Height
func (b *BlockchainService) GetChainHeight(ctx context.Context) (*ChainHeight, *http.Response, error) {
	req, err := b.client.NewRequest("GET", "chain/height", nil)
	if err != nil {
		return nil, nil, err
	}

	bh := &ChainHeight{}
	resp, err := b.client.Do(ctx, req, bh)
	if err != nil {
		return nil, resp, err
	}

	return bh, resp, nil
}

// Get the Chain Score
func (b *BlockchainService) GetChainScore(ctx context.Context) (*ChainScore, *http.Response, error) {
	req, err := b.client.NewRequest("GET", "chain/score", nil)
	if err != nil {
		return nil, nil, err
	}

	cs := &ChainScore{}
	resp, err := b.client.Do(ctx, req, cs)
	if err != nil {
		return nil, resp, err
	}

	return cs, resp, nil
}

// Get block height
func (b *BlockchainService) GetBlockHeight(ctx context.Context, height int) (*BlockInfo, *http.Response, error) {
	u := fmt.Sprintf("block/%d", height)

	req, err := b.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	binfo := &BlockInfo{}
	resp, err := b.client.Do(ctx, req, binfo)
	if err != nil {
		return nil, resp, err
	}

	return binfo, resp, nil
}
