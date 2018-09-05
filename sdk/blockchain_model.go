package sdk

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
