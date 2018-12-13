// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package sdk

import (
	"fmt"
	"github.com/json-iterator/go"
	"github.com/proximax-storage/proximax-utils-go/mock"
	"github.com/proximax-storage/proximax-utils-go/tests"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func init() {
	jsoniter.RegisterTypeEncoder("*NamespaceIds", testNamespaceIDs)
	jsoniter.RegisterTypeDecoder("*NamespaceIds", ad)

	i, _ := (&big.Int{}).SetString("9562080086528621131", 10)
	testNamespaceId = bigIntToNamespaceId(i)

	namespaceCorr.Levels = []*NamespaceId{testNamespaceId}
	namespaceNameCorr.NamespaceId = testNamespaceId
}

// test data
const (
	pageSize        = 32
	mosaicNamespace = "84b3552d375ffa4b"
)

var (
	namespaceClient = mockServer.getTestNetClientUnsafe().Namespace
	testAddresses   = []*Address{
		{Address: "SDRDGFTDLLCB67D4HPGIMIHPNSRYRJRT7DOBGWZY"},
		{Address: "SBCPGZ3S2SCC3YHBBTYDCUZV4ZZEPHM2KGCP4QXX"},
	}
	testAddress = Address{Address: "SCASIIAPS6BSFEC66V6MU5ZGEVWM53BES5GYBGLE"}

	testNamespaceId  *NamespaceId
	testNamespaceIDs = &NamespaceIds{
		List: []*NamespaceId{
			testNamespaceId,
		},
	}
	ad   = &NamespaceIds{}
	meta = `"meta": {
			"active": true,
			"index": 0,
			"id": "5B55E02EACCB7B00015DB6EB"
			}`
	tplInfo = "{" + meta + `
			  ,
			  "namespace": {
				"namespaceId": [
				  0,
				  0
				], 
				"type": 0,
				"depth": 1,
				"level0": [
				  929036875,
				  2226345261
				],
				"parentId": [
				  0,
				  0
				],
				"owner": "321DE652C4D3362FC2DDF7800F6582F4A10CFEA134B81F8AB6E4BE78BBA4D18E",
				"ownerAddress": "904A1B7A7432C968202264C2CBDE0E8E5547EED3AD66E52BAC",
				"startHeight": [
				  1,
				  0
				],
				"subNamespaces": [
				[
					0,
					0
				]
				],
				"mosaicIds": [
				[
					0,
					0
				]
				],
				"endHeight": [
				  4294967295,
				  4294967295
				]
			  }
			}`

	namespaceCorr = &NamespaceInfo{
		NamespaceId: bigIntToNamespaceId(big.NewInt(0)),
		Active:      true,
		Index:       0,
		MetaId:      "5B55E02EACCB7B00015DB6EB",
		Depth:       1,
		TypeSpace:   Root,
		Owner: &PublicAccount{
			Address: &Address{
				Type:    NotSupportedNet,
				Address: "ABFBW6TUGLEWQIBCMTBMXXQORZKUP3WTVV5DRLO7",
			},
			PublicKey: "321DE652C4D3362FC2DDF7800F6582F4A10CFEA134B81F8AB6E4BE78BBA4D18E",
		},
		SubNamespaceIds: []*NamespaceId{
			bigIntToNamespaceId(big.NewInt(0)),
		},
		MosaicIds: []*MosaicId{
			bigIntToMosaicId(big.NewInt(0)),
		},
		EndHeight:   uint64DTO{4294967295, 4294967295}.toBigInt(),
		StartHeight: big.NewInt(1),
		ParentId:    bigIntToNamespaceId(big.NewInt(0)),
	}

	namespaceNameCorr = &NamespaceName{
		NamespaceId: bigIntToNamespaceId(big.NewInt(0)),
		Name:        "nem",
		ParentId:    bigIntToNamespaceId(big.NewInt(0)),
	}

	tplInfoArr = "[" + tplInfo + "]"
)

func TestNamespaceService_GetNamespace(t *testing.T) {
	mockServer.AddRouter(&mock.Router{
		Path:     fmt.Sprintf("/namespace/%s", testNamespaceId.toHexString()),
		RespBody: tplInfo,
	})

	nsInfo, err := namespaceClient.GetNamespace(ctx, testNamespaceId)

	assert.Nilf(t, err, "NamespaceService.GetNamespace returned error: %s", err)
	tests.ValidateStringers(t, namespaceCorr, nsInfo)
}

func TestNamespaceService_GetNamespacesFromAccount(t *testing.T) {
	mockServer.AddRouter(&mock.Router{
		Path:     fmt.Sprintf(pathNamespacesFromAccount, testAddress.Address),
		RespBody: tplInfoArr,
	})

	nsInfoArr, err := namespaceClient.GetNamespacesFromAccount(ctx, &testAddress, nil, pageSize)

	assert.Nilf(t, err, "NamespaceService.GetNamespacesFromAccount returned error: %s", err)

	fmt.Println(nsInfoArr)

	for _, nsInfo := range nsInfoArr {
		tests.ValidateStringers(t, namespaceCorr, nsInfo)
	}
}

func TestNamespaceService_GetNamespacesFromAccounts(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockServer.AddRouter(&mock.Router{
			Path:     pathNamespacesFromAccounts,
			RespBody: tplInfoArr,
		})

		nsInfoArr, err := namespaceClient.GetNamespacesFromAccounts(ctx, testAddresses, nil, pageSize)

		assert.Nilf(t, err, "NamespaceService.GetNamespacesFromAccounts returned error: %s", err)

		for _, nsInfo := range nsInfoArr {
			tests.ValidateStringers(t, namespaceCorr, nsInfo)
		}
	})

	t.Run("no test addresses", func(t *testing.T) {
		_, err := namespaceClient.GetNamespacesFromAccounts(ctx, nil, nil, pageSize)

		assert.NotNil(t, err, "request with empty Addresses must return error")
	})
}

func TestNamespaceService_GetNamespaceNames(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockServer.AddRouter(&mock.Router{
			Path: pathNamespaceNames,
			RespBody: `[
			  {
				"namespaceId": [
				  929036875,
				  2226345261
				],
				"name": "nem"
			  }
			]`,
		})

		nsInfoArr, err := namespaceClient.GetNamespaceNames(ctx, []*NamespaceId{testNamespaceId})

		assert.Nilf(t, err, "NamespaceService.GetNamespaceNames returned error: %s", err)

		for _, nsInfo := range nsInfoArr {
			tests.ValidateStringers(t, namespaceNameCorr, nsInfo)
		}
	})

	t.Run("empty namespaceIds", func(t *testing.T) {
		_, err := namespaceClient.GetNamespaceNames(ctx, []*NamespaceId{})

		assert.Equal(t, ErrEmptyNamespaceIds, err, "request with empty NamespaceIds must return error")
	})
}
