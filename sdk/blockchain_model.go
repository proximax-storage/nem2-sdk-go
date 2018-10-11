// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package sdk

import (
	"fmt"
	"math/big"
)

// Models
// Block
type BlockInfo struct {
	NetworkType
	Hash                  string
	GenerationHash        string
	TotalFee              *big.Int
	NumTransactions       uint64
	Signature             string
	Signer                *PublicAccount
	Version               uint64
	Type                  uint64
	Height                *big.Int
	Timestamp             *big.Int
	Difficulty            *big.Int
	PreviousBlockHash     string
	BlockTransactionsHash string
}

func (b *BlockInfo) String() string {
	return fmt.Sprintf(
		`
			"NetworkType:" %d,
			"Hash": %s,
			"GenerationHash": %s,
			"TotalFee": %s,
			"NumTransactions": %d,
			"Signature": %s,
			"Signer": %s,
			"Version": %d,
			"Type": %d,
			"Height": %s,
			"Timestamp": %s,
			"Difficulty": %s,
			"PreviousBlockHash": %s,
			"BlockTransactionHash": %s
		`,
		b.NetworkType,
		b.Hash,
		b.GenerationHash,
		b.TotalFee,
		b.NumTransactions,
		b.Signature,
		b.Signer,
		b.Version,
		b.Type,
		b.Height,
		b.Timestamp,
		b.Difficulty,
		b.PreviousBlockHash,
		b.BlockTransactionsHash,
	)
}

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

	pa, err := NewPublicAccount(dto.Block.Signer, nt)
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

// Blockchain Storage
type BlockchainStorageInfo struct {
	NumBlocks       int `json:"numBlocks"`
	NumTransactions int `json:"numTransactions"`
	NumAccounts     int `json:"numAccounts"`
}

// Chain Score
type chainScoreDTO struct {
	ScoreHigh uint64DTO `json:"scoreHigh"`
	ScoreLow  uint64DTO `json:"scoreLow"`
}

func (dto *chainScoreDTO) toStruct() *big.Int {
	return uint64DTO{uint32(dto.ScoreLow.toBigInt().Uint64()), uint32(dto.ScoreHigh.toBigInt().Uint64())}.toBigInt()
}
