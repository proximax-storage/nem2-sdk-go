package sdk

import (
	"errors"
	"fmt"
	"strconv"
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
	Version               *uint64  `json:"version"`
	Type                  *uint64  `json:"type"`
	Height                []uint64 `json:"height"`
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
	NumTransactions *uint64  `json:"numTransactions"`
	MerkleTree      []string `json:"merkleTree"`
}

// Blockchain Storage
type BlockchainStorageInfo struct {
	NumBlocks       *int `json:"numBlocks"`
	NumTransactions *int `json:"numTransactions"`
	NumAccounts     *int `json:"numAccounts"`
}

type NetworkType uint8

// NetworkType enums
const (
	MAIN_NET                  NetworkType = 104
	TEST_NET                  NetworkType = 152
	MIJIN                     NetworkType = 96
	MIJIN_TEST                NetworkType = 144
	NOT_SUPPORTED_NETWORKTYPE NetworkType = 0
)

func (nt NetworkType) String() string {
	return fmt.Sprintf("%d", nt)
}

// Network error
var networkTypeError = errors.New("wrong raw NetworkType value")

// Get NetworkType by raw value
func NetworkTypeFromRaw(value uint32) (NetworkType, error) {
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

func ExtractNetworkType(version uint64) (NetworkType, error) {
	res, err := strconv.ParseUint(strconv.FormatUint(version, 16)[:2], 16, 32)
	if err != nil {
		return 0, err
	}

	t, err := NetworkTypeFromRaw(uint32(res))
	if err != nil {
		return 0, err
	}
	return t, nil
}
