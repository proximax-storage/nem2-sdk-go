package sdk

import "math/big"

type blockInfoDTO struct {
	BlockMeta struct {
		Hash            string    `json:"hash"`
		GenerationHash  string    `json:"generationHash"`
		TotalFee        uint64DTO `json:"totalFee"`
		NumTransactions uint64    `json:"numTransactions"`
		// MerkleTree      uint64DTO `json:"merkleTree"` is needed?
	} `json:"meta"`
	Block struct {
		Signature             string    `json:"signature"`
		Signer                string    `json:"signer"`
		Version               uint64    `json:"version"`
		Type                  uint64    `json:"type"`
		Height                uint64DTO `json:"height"`
		Timestamp             uint64DTO `json:"timestamp"`
		Difficulty            uint64DTO `json:"difficulty"`
		PreviousBlockHash     string    `json:"previousBlockHash"`
		BlockTransactionsHash string    `json:"blockTransactionsHash"`
	} `json:"block"`
}

func (dto *blockInfoDTO) toStruct() (*BlockInfo, error) {
	nt := ExtractNetworkType(dto.Block.Version)

	pa, err := NewAccountFromPublicKey(dto.Block.Signer, nt)
	if err != nil {
		return nil, err
	}

	v, err := ExtractVersion(dto.Block.Version)
	if err != nil {
		return nil, err
	}

	return &BlockInfo{
		NetworkType:           nt,
		Hash:                  dto.BlockMeta.Hash,
		GenerationHash:        dto.BlockMeta.GenerationHash,
		TotalFee:              dto.BlockMeta.TotalFee.toBigInt(),
		NumTransactions:       dto.BlockMeta.NumTransactions,
		Signature:             dto.Block.Signature,
		Signer:                pa,
		Version:               v,
		Type:                  dto.Block.Type,
		Height:                dto.Block.Height.toBigInt(),
		Timestamp:             dto.Block.Timestamp.toBigInt(),
		Difficulty:            dto.Block.Difficulty.toBigInt(),
		PreviousBlockHash:     dto.Block.PreviousBlockHash,
		BlockTransactionsHash: dto.Block.BlockTransactionsHash,
	}, nil
}

// Chain Score
type chainScoreDTO struct {
	ScoreHigh uint64DTO `json:"scoreHigh"`
	ScoreLow  uint64DTO `json:"scoreLow"`
}

func (dto *chainScoreDTO) toStruct() *big.Int {
	return uint64DTO{uint32(dto.ScoreLow.toBigInt().Uint64()), uint32(dto.ScoreHigh.toBigInt().Uint64())}.toBigInt()
}

type blockInfoDTOs []*blockInfoDTO

func (b *blockInfoDTOs) toStruct() ([]*BlockInfo, error) {
	dtos := *b
	blocks := make([]*BlockInfo, 0, len(dtos))

	for _, dto := range dtos {
		block, err := dto.toStruct()
		if err != nil {
			return nil, err
		}

		blocks = append(blocks, block)
	}

	return blocks, nil
}
