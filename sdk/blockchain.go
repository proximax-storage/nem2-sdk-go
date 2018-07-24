package sdk

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
)

type BlockchainService service

// Models
// Chain Height
type ChainHeight struct {
	Height []uint64 `json:"height"`
}

// Chain Score
type ChainScore struct {
	ScoreHigh []uint64 `json:"scoreHigh"`
	ScoreLow  []uint64 `json:"scoreLow"`
}

// Block Info
type BlockInfo struct {
	Block     *Block     `json:"block"`
	BlockMeta *BlockMeta `json:"meta"`
}

// Block
type Block struct {
	Signature             *string  `json:"signature"`
	Signer                *string  `json:"signer"`
	Version               *uint64  `json:"version"` // TODO: Java BigDecimal equivalent? big.Rat has no unmarshall?
	Type                  *uint64  `json:"type"`    // TODO: Java BigDecimal equivalent? big.Rat has no unmarshall?
	Height                []uint64 `json:"version"`
	Timestamp             []uint64 `json:"timestamp"`
	Difficulty            []uint64 `json:"difficulty"`
	PreviousBlockHash     *string  `json:"previousBlockHash"`
	BlockTransactionsHash *string  `json:"blockTransactionsHash"`
}

// Block Meta
type BlockMeta struct {
	Hash            *string  `json:"hash"`
	GenerationHash  *string  `json:"generationHash"`
	TotalFee        []uint64 `json:"totalFee"`
	NumTransactions *uint64  `json:"numTransactions"` // TODO: Java BigDecimal equivalent? big.Rat has no unmarshall?
}

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
