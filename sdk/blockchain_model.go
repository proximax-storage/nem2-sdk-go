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
