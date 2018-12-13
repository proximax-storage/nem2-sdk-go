// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package sdk

import (
	"fmt"
	"github.com/proximax-storage/proximax-utils-go/mock"
	"github.com/proximax-storage/proximax-utils-go/tests"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
	"time"
)

// Mock response for TestBlockchainService_GetBlockHeight & GetBlockInfo
const (
	blockInfoJSON = `{
	"meta": {
		"hash": "83FB2550BDB72B6F507BDBDE90C265D4A324DF9F1EFEFD9F7BD0FDF6391C30D8",
		"generationHash": "8EC49BBADB3B2FD90810DB9BDACF1FDE999295C594B5FD4B584A0A72F5AAFA59",
		"totalFee": [
			0,
			0
		],
		"numTransactions": 25
	},
	"block": {
		"signature": "0BEAE2B3DCDEC268B43797C7A855EC03FDEE0B4687EC14F250D0EA3588ADDD0B42EBB77E14157EAB168B41457CA28395C1EBAB354B0A20CCB5FC73CFA65A3107",
		"signer": "321DE652C4D3362FC2DDF7800F6582F4A10CFEA134B81F8AB6E4BE78BBA4D18E",
		"version": 36867,
		"type": 32835,
		"height": [
			1,
			0
		],
		"timestamp": [
			0,
			0
		],
		"difficulty": [
			276447232,
			23283
		],
		"previousBlockHash": "0000000000000000000000000000000000000000000000000000000000000000",
		"blockTransactionsHash": "8A77819676852F20EB7ACDE5A18F7CE060C3D1A61A7EF80A99B3346EB9091B19"
	}
}`
	// Mock response for TestBlockchainService_GetBlockTransactions
	blockTransactionsJSON = `[
	{
		"meta": {
			"height": [
				1,
				0
			],
			"hash": "D28F325EDA671D0C98AC9087A8C0568C8C25F75C63F9DBE84EC5FB9F63E82366",
			"merkleComponentHash": "D28F325EDA671D0C98AC9087A8C0568C8C25F75C63F9DBE84EC5FB9F63E82366",
			"index": 0,
			"id": "5B55E02EACCB7B00015DB6D2"
		},
		"transaction": {
			"signature": "AE1558A33F4F595AD5DCEAE4EC11606E815A781E75E3EEC7E9F8BB46BDAF16670C8C36C6815F74FD83487178DDAB8FCE4B4B633875A1549D4FB068ABC5B22A0C",
			"signer": "321DE652C4D3362FC2DDF7800F6582F4A10CFEA134B81F8AB6E4BE78BBA4D18E",
			"version": 36866,
			"type": 16718,
			"fee": [
				0,
				0
			],
			"deadline": [
				1,
				0
			],
			"namespaceType": 0,
			"duration": [
				0,
				0
			],
			"namespaceId": [
				929036875,
				2226345261
			],
			"name": "nem"
		}
	}
]`
)

var (
	blockClient = mockServer.getTestNetClientUnsafe().Blockchain
	testHeight  = big.NewInt(1)
	testLimit   = big.NewInt(100)
)

// Expected value for TestBlockchainService_GetBlockHeight
var wantBlockTransactions []Transaction

// Expected value for TestBlockchainService_GetBlockHeight
var wantBlockInfo *BlockInfo

func init() {
	pubAcc, _ := NewAccountFromPublicKey("321DE652C4D3362FC2DDF7800F6582F4A10CFEA134B81F8AB6E4BE78BBA4D18E", MijinTest)

	wantBlockInfo = &BlockInfo{
		NetworkType:           MijinTest,
		Hash:                  "83FB2550BDB72B6F507BDBDE90C265D4A324DF9F1EFEFD9F7BD0FDF6391C30D8",
		GenerationHash:        "8EC49BBADB3B2FD90810DB9BDACF1FDE999295C594B5FD4B584A0A72F5AAFA59",
		TotalFee:              uint64DTO{0, 0}.toBigInt(),
		NumTransactions:       25,
		Signature:             "0BEAE2B3DCDEC268B43797C7A855EC03FDEE0B4687EC14F250D0EA3588ADDD0B42EBB77E14157EAB168B41457CA28395C1EBAB354B0A20CCB5FC73CFA65A3107",
		Signer:                pubAcc,
		Version:               3,
		Type:                  32835,
		Height:                uint64DTO{1, 0}.toBigInt(),
		Timestamp:             uint64DTO{0, 0}.toBigInt(),
		Difficulty:            uint64DTO{276447232, 23283}.toBigInt(),
		PreviousBlockHash:     "0000000000000000000000000000000000000000000000000000000000000000",
		BlockTransactionsHash: "8A77819676852F20EB7ACDE5A18F7CE060C3D1A61A7EF80A99B3346EB9091B19",
	}

	wantBlockTransactions = append(wantBlockTransactions, &RegisterNamespaceTransaction{
		AbstractTransaction: AbstractTransaction{
			Type:        RegisterNamespace,
			Version:     uint64(2),
			NetworkType: MijinTest,
			Signature:   "AE1558A33F4F595AD5DCEAE4EC11606E815A781E75E3EEC7E9F8BB46BDAF16670C8C36C6815F74FD83487178DDAB8FCE4B4B633875A1549D4FB068ABC5B22A0C",
			Signer:      nil,
			Fee:         uint64DTO{0, 0}.toBigInt(),
			Deadline:    &Deadline{time.Unix(uint64DTO{1, 0}.toBigInt().Int64(), int64(time.Millisecond))},
			TransactionInfo: &TransactionInfo{
				Height:              uint64DTO{1, 0}.toBigInt(),
				Hash:                "D28F325EDA671D0C98AC9087A8C0568C8C25F75C63F9DBE84EC5FB9F63E82366",
				MerkleComponentHash: "D28F325EDA671D0C98AC9087A8C0568C8C25F75C63F9DBE84EC5FB9F63E82366",
				Index:               0,
				Id:                  "5B55E02EACCB7B00015DB6D2",
			},
		},
		NamspaceName: "nem",
		Duration:     uint64DTO{0, 0}.toBigInt(),
	})
}

func TestBlockchainService_GetBlocksByHeightWithLimit(t *testing.T) {
	mockServer.AddRouter(&mock.Router{
		Path:     fmt.Sprintf(pathBlockInfo, testHeight, testLimit),
		RespBody: "[" + blockInfoJSON + "]",
	})

	bcInfo, err := blockClient.GetBlocksByHeightWithLimit(ctx, testHeight, testLimit)

	assert.Nilf(t, err, "GetBlocksByHeightWithLimit returned error: %s", err)

	tests.ValidateStringers(t, wantBlockInfo, bcInfo[0])
}

func TestBlockchainService_GetBlockchainHeight(t *testing.T) {
	want := uint64DTO{11235, 0}.toBigInt()

	mockServer.AddRouter(&mock.Router{
		Path:     pathBlockHeight,
		RespBody: `{"height":[11235,0]}`,
	})

	got, err := blockClient.GetBlockchainHeight(ctx)

	assert.Nilf(t, err, "GetBlockchainHeight returned error: %s", err)

	tests.ValidateStringers(t, want, got)
}

func TestBlockchainService_GetBlockchainStorage(t *testing.T) {
	want := &BlockchainStorageInfo{NumBlocks: 62094, NumTransactions: 56, NumAccounts: 25}

	mockServer.AddRouter(&mock.Router{
		Path:     pathBlockStorage,
		RespBody: `{"numBlocks":62094,"numTransactions":56,"numAccounts":25}`,
	})

	got, err := blockClient.GetBlockchainStorage(ctx)

	assert.Nilf(t, err, "GetBlockchainStorage returned error: %s", err)

	tests.ValidateStringers(t, want, got)
}

func TestBlockchainService_GetBlockchainScore(t *testing.T) {
	dto := chainScoreDTO{ScoreHigh: uint64DTO{0, 0}, ScoreLow: uint64DTO{3999308498, 121398739}}

	mockServer.AddRouter(&mock.Router{
		Path:     pathBlockScore,
		RespBody: `{"scoreHigh": [0,0],"scoreLow": [3999308498,121398739]}`,
	})

	got, err := blockClient.GetBlockchainScore(ctx)

	assert.Nilf(t, err, "GetBlockchainScore returned error: %s", err)

	tests.ValidateStringers(t, dto.toStruct(), got)
}

func TestBlockchainService_GetBlockByHeight(t *testing.T) {
	mockServer.AddRouter(&mock.Router{
		Path:     fmt.Sprintf(pathBlockByHeight, testHeight.String()),
		RespBody: blockInfoJSON,
	})

	got, err := blockClient.GetBlockByHeight(ctx, testHeight)

	assert.Nilf(t, err, "GetBlockByHeight returned error: %s", err)

	tests.ValidateStringers(t, wantBlockInfo, got)
}

func TestBlockchainService_GetBlockTransactions(t *testing.T) {
	mockServer.AddRouter(&mock.Router{
		Path:     fmt.Sprintf(pathBlockGetTransaction, testHeight.String()),
		RespBody: blockTransactionsJSON,
	})

	got, err := blockClient.GetBlockTransactions(ctx, testHeight)

	assert.Nilf(t, err, "GetBlockByHeight returned error: %s", err)

	for key, transaction := range got {
		assert.Equal(t, wantBlockTransactions[key].GetAbstractTransaction().Signature, transaction.GetAbstractTransaction().Signature)
	}
}
