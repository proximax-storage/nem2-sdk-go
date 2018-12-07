// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package sdk

import (
	"github.com/proximax-storage/proximax-utils-go/str"
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
	return str.StructToString(
		"BlockInfo",
		str.NewField("NetworkType", str.IntPattern, b.NetworkType),
		str.NewField("Hash", str.StringPattern, b.Hash),
		str.NewField("GenerationHash", str.StringPattern, b.GenerationHash),
		str.NewField("TotalFee", str.StringPattern, b.TotalFee),
		str.NewField("NumTransactions", str.IntPattern, b.NumTransactions),
		str.NewField("Signature", str.StringPattern, b.Signature),
		str.NewField("Signer", str.StringPattern, b.Signer),
		str.NewField("Version", str.IntPattern, b.Version),
		str.NewField("Type", str.IntPattern, b.Type),
		str.NewField("Height", str.StringPattern, b.Height),
		str.NewField("Timestamp", str.StringPattern, b.Timestamp),
		str.NewField("Difficulty", str.StringPattern, b.Difficulty),
		str.NewField("PreviousBlockHash", str.StringPattern, b.PreviousBlockHash),
		str.NewField("BlockTransactionsHash", str.StringPattern, b.BlockTransactionsHash),
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

// Blockchain Storage
type BlockchainStorageInfo struct {
	NumBlocks       int `json:"numBlocks"`
	NumTransactions int `json:"numTransactions"`
	NumAccounts     int `json:"numAccounts"`
}

func (b *BlockchainStorageInfo) String() string {
	return str.StructToString(
		"BlockchainStorageInfo",
		str.NewField("NumBlocks", str.IntPattern, b.NumBlocks),
		str.NewField("NumTransactions", str.IntPattern, b.NumTransactions),
		str.NewField("NumAccounts", str.IntPattern, b.NumAccounts),
	)
}

// Chain Score
type chainScoreDTO struct {
	ScoreHigh uint64DTO `json:"scoreHigh"`
	ScoreLow  uint64DTO `json:"scoreLow"`
}

func (dto *chainScoreDTO) toStruct() *big.Int {
	return uint64DTO{uint32(dto.ScoreLow.toBigInt().Uint64()), uint32(dto.ScoreHigh.toBigInt().Uint64())}.toBigInt()
}
