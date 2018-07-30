package sdk

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
)

type BlockchainService service

// Get the Chain Height
func (b *BlockchainService) GetChainHeight(ctx context.Context) (*ChainHeight, *http.Response, error) {
	bh := &ChainHeight{}
	resp, err := b.client.DoNewRequest(ctx, "GET", "chain/height", nil, bh)
	if err != nil {
		return nil, nil, err
	}

	return bh, resp, nil
}

// Get the Chain Score
func (b *BlockchainService) GetChainScore(ctx context.Context) (*ChainScore, *http.Response, error) {
	cs := &ChainScore{}
	resp, err := b.client.DoNewRequest(ctx, "GET", "chain/score", nil, cs)
	if err != nil {
		return nil, nil, err
	}

	return cs, resp, nil
}

// Get block height
func (b *BlockchainService) GetBlockHeight(ctx context.Context, height int) (*BlockInfo, *http.Response, error) {
	u := fmt.Sprintf("block/%d", height)
	binfo := &BlockInfo{}

	resp, err := b.client.DoNewRequest(ctx, "GET", u, nil, binfo)
	if err != nil {
		return nil, nil, err
	}

	return binfo, resp, nil
}
