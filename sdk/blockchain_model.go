package sdk

import (
	"errors"
)

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

type NetworkType uint8

// NetworkType enums
const (
	MAIN_NET NetworkType = 104
	TEST_NET NetworkType = 152
	MIJIN NetworkType = 96
	MIJIN_TEST NetworkType = 144
)

// Network error
var networkTypeError = errors.New("wrong raw NetworkType int")

// Get NetworkType by raw value
func NetworkTypeFromRaw(value int) (NetworkType, error){
	switch value {
	case 104:
		return MAIN_NET, nil
	case 152:
		return TEST_NET, nil
	case 96:
		return MIJIN, nil
	case 144:
		return MIJIN_TEST, nil
	default:
		return 0, networkTypeError
	}
}
