package sdk

import (
	"testing"
	"net/http"
	"fmt"
	"context"
	"reflect"
)

func TestBlockchainService_GetChainHeight(t *testing.T) {
	client, mux, _, teardown := setupMockServer()
	defer teardown()

	mux.HandleFunc("/chain/height", func(w http.ResponseWriter, r *http.Request) {
		// Mock JSON response
		fmt.Fprint(w, `{"height":[11235,0]}`)
	})

	got, _, err := client.Blockchain.GetChainHeight(context.Background())
	if err != nil {
		t.Errorf("Blockchain.GetChainHeight returned error: %v", err)
	}

	want := &ChainHeight{Height: []int64{11235, 0}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Blockchain.GetChainHeight returned %+v, want %+v", got, want)
	}
}

func TestBlockchainService_GetChainScore(t *testing.T) {
	client, mux, _, teardown := setupMockServer()
	defer teardown()

	mux.HandleFunc("/chain/score", func(w http.ResponseWriter, r *http.Request) {
		// Mock JSON response
		fmt.Fprint(w, `{"scoreHigh": [0,0],"scoreLow": [3999308498,121398739]}`)
	})

	got, _, err := client.Blockchain.GetChainScore(context.Background())
	if err != nil {
		t.Errorf("Blockchain.GetChainScore returned error: %v", err)
	}

	want := &ChainScore{ScoreHigh: []int64{0, 0}, ScoreLow: []int64{3999308498, 121398739}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Blockchain.GetChainScore returned %+v, want %+v", got, want)
	}
}

func TestBlockchainService_GetBlockHeight(t *testing.T) {
	client, mux, _, teardown := setupMockServer()
	defer teardown()

	mux.HandleFunc("/block/1", func(w http.ResponseWriter, r *http.Request) {
		// Mock JSON response
		w.WriteHeader(http.StatusOK)
		w.Write(blockInfoJSON)
	})

	got, _, err := client.Blockchain.GetBlockHeight(context.Background(), 1)
	if err != nil {
		t.Errorf("Blockchain.GetBlockHeight returned error: %v", err)
	}

	if want := wantBlockInfo; !reflect.DeepEqual(got, want) {
		t.Errorf("Blockchain.GetBlockHeight returned %+v, want %+v", got, want)
	}
}

var blockInfoJSON = []byte(`{
	"meta": {
		"hash": "83FB2550BDB72B6F507BDBDE90C265D4A324DF9F1EFEFD9F7BD0FDF6391C30D8",
		"generationHash": "8EC49BBADB3B2FD90810DB9BDACF1FDE999295C594B5FD4B584A0A72F5AAFA59",
		"totalFee": [
			0,
			0
		],
		"numTransactions": 25,
		"merkleTree": [
			"0o8yXtpnHQyYrJCHqMBWjIwl91xj+dvoTsX7n2PoI2Y=",
			"OWfIFOfMGCAd5BTpcAHbJCS+KLrDaEW4yDzC0OgBVzc=",
			"wALVcIaWFKQ8P9PDiffHSDr1l5TGlmVNqrQgy2PrXIE=",
			"0TiPIIoTC0RjwXv7nOmGSg//tIgKBotImfsEy61cJjE=",
			"fR6tYj1KLsmgIpdZCuyE+RPUdGW0xSu4t0LRVw1O0eI=",
			"YaEnbHOFKXKcxvrIeK+LaLt9Z1Z6rgk8hkDEt3IEa6k=",
			"tEHMyTVy/R6i+meP61PszSA99lyCnc2/v2qPexMudSk=",
			"fTVOBWoQ562sZnQdECGw55pXmY6tfhcZiCEUHOh89j8=",
			"l5YwqzEediC+pPT93Xp1Ywvwb+G6Ut37sM+uYk6D3+Q=",
			"sBMlE6t3tOc8hXmKzLDJfH6oOoqyvP7xlOsOPMsvsAY="
		]
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
}`)

var wantBlockInfo = &BlockInfo{
	&Block{
		Signature:             String("0BEAE2B3DCDEC268B43797C7A855EC03FDEE0B4687EC14F250D0EA3588ADDD0B42EBB77E14157EAB168B41457CA28395C1EBAB354B0A20CCB5FC73CFA65A3107"),
		Signer:                String("321DE652C4D3362FC2DDF7800F6582F4A10CFEA134B81F8AB6E4BE78BBA4D18E"),
		Version:               Int64(36867),
		Type:                  Int64(32835),
		Height:                []int64{1, 0},
		Timestamp:             []int64{0, 0},
		Difficulty:            []int64{276447232, 23283},
		PreviousBlockHash:     String("0000000000000000000000000000000000000000000000000000000000000000"),
		BlockTransactionsHash: String("8A77819676852F20EB7ACDE5A18F7CE060C3D1A61A7EF80A99B3346EB9091B19"),
	},
	&BlockMeta{
		Hash:            String("83FB2550BDB72B6F507BDBDE90C265D4A324DF9F1EFEFD9F7BD0FDF6391C30D8"),
		GenerationHash:  String("8EC49BBADB3B2FD90810DB9BDACF1FDE999295C594B5FD4B584A0A72F5AAFA59"),
		TotalFee:        []int64{0, 0},
		NumTransactions: Int64(25),
		MerkleTree: []string{
			"0o8yXtpnHQyYrJCHqMBWjIwl91xj+dvoTsX7n2PoI2Y=",
			"OWfIFOfMGCAd5BTpcAHbJCS+KLrDaEW4yDzC0OgBVzc=",
			"wALVcIaWFKQ8P9PDiffHSDr1l5TGlmVNqrQgy2PrXIE=",
			"0TiPIIoTC0RjwXv7nOmGSg//tIgKBotImfsEy61cJjE=",
			"fR6tYj1KLsmgIpdZCuyE+RPUdGW0xSu4t0LRVw1O0eI=",
			"YaEnbHOFKXKcxvrIeK+LaLt9Z1Z6rgk8hkDEt3IEa6k=",
			"tEHMyTVy/R6i+meP61PszSA99lyCnc2/v2qPexMudSk=",
			"fTVOBWoQ562sZnQdECGw55pXmY6tfhcZiCEUHOh89j8=",
			"l5YwqzEediC+pPT93Xp1Ywvwb+G6Ut37sM+uYk6D3+Q=",
			"sBMlE6t3tOc8hXmKzLDJfH6oOoqyvP7xlOsOPMsvsAY=",
		},
	},
}
