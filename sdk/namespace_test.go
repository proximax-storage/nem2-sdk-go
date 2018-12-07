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
	"net/http"
	"testing"
)

func init() {
	jsoniter.RegisterTypeEncoder("*NamespaceIds", testNamespaceIDs)
	jsoniter.RegisterTypeDecoder("*NamespaceIds", ad)

	i, _ := (&big.Int{}).SetString("9562080086528621131", 10)
	testNamespaceId.Id = i

	namespaceCorr.Levels = []*NamespaceId{{Id: i}}
	namespaceNameCorr.NamespaceId.Id = i
}

// test data
const (
	pageSize        = 32
	mosaicNamespace = "84b3552d375ffa4b"
	testNamespaceID = "5B55E02EACCB7B00015DB6EB"
)

var (
	namespaceClient = mockServer.getTestNetClientUnsafe().Namespace
	testAddresses   = Addresses{
		List: []*Address{
			{Address: "SDRDGFTDLLCB67D4HPGIMIHPNSRYRJRT7DOBGWZY"},
			{Address: "SBCPGZ3S2SCC3YHBBTYDCUZV4ZZEPHM2KGCP4QXX"},
		},
	}
	testAddress = Address{Address: "SCASIIAPS6BSFEC66V6MU5ZGEVWM53BES5GYBGLE"}

	testNamespaceId  = &NamespaceId{}
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
				"endHeight": [
				  4294967295,
				  4294967295
				]
			  }
			}`

	namespaceCorr = &NamespaceInfo{
		Active:    true,
		Index:     0,
		MetaId:    "5B55E02EACCB7B00015DB6EB",
		Depth:     1,
		TypeSpace: Root,
		Owner: &PublicAccount{
			Address: &Address{
				Type:    NotSupportedNet,
				Address: "ABFBW6TUGLEWQIBCMTBMXXQORZKUP3WTVV5DRLO7",
			},
			PublicKey: "321DE652C4D3362FC2DDF7800F6582F4A10CFEA134B81F8AB6E4BE78BBA4D18E",
		},
		EndHeight:   uint64DTO{4294967295, 4294967295}.toBigInt(),
		StartHeight: big.NewInt(1),
		ParentId: &NamespaceId{
			Id: big.NewInt(0),
		},
	}

	namespaceNameCorr = &NamespaceName{
		NamespaceId: &NamespaceId{},
		Name:        "nem",
		ParentId: &NamespaceId{
			Id: big.NewInt(0),
		},
	}

	tplInfoArr = "[" + tplInfo + "]"
)

func TestNamespaceService_GetNamespace(t *testing.T) {
	mockServer.AddRouter(&mock.Router{
		Path:     fmt.Sprintf("/namespace/%s", testNamespaceId.toHexString()),
		RespBody: tplInfo,
	})

	nsInfo, resp, err := namespaceClient.GetNamespace(ctx, testNamespaceId)

	assert.Nilf(t, err, "NamespaceService.GetNamespace returned error: %s", err)

	if tests.IsOkResponse(t, resp) {
		tests.ValidateStringers(t, namespaceCorr, nsInfo)
	}
}

func TestNamespaceService_GetNamespacesFromAccount(t *testing.T) {
	mockServer.AddRouter(&mock.Router{
		Path:     fmt.Sprintf(pathNamespacesFromAccount, testAddress.Address),
		RespBody: tplInfoArr,
	})

	nsInfoArr, resp, err := namespaceClient.GetNamespacesFromAccount(ctx, &testAddress, testNamespaceID, pageSize)

	assert.Nilf(t, err, "NamespaceService.GetNamespacesFromAccount returned error: %s", err)
	if tests.IsOkResponse(t, resp) {
		for _, nsInfo := range nsInfoArr.List {
			tests.ValidateStringers(t, namespaceCorr, nsInfo)
		}
	}
}

func TestNamespaceService_GetNamespacesFromAccounts(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockServer.AddRouter(&mock.Router{
			Path:     pathNamespacesFromAccounts,
			RespBody: tplInfoArr,
		})

		nsInfoArr, resp, err := namespaceClient.GetNamespacesFromAccounts(ctx, &testAddresses, testNamespaceID, pageSize)

		assert.Nilf(t, err, "NamespaceService.GetNamespacesFromAccounts returned error: %s", err)

		if tests.IsOkResponse(t, resp) {
			for _, nsInfo := range nsInfoArr.List {
				tests.ValidateStringers(t, namespaceCorr, nsInfo)
			}
		}
	})

	t.Run("no test addresses", func(t *testing.T) {
		_, resp, err := namespaceClient.GetNamespacesFromAccounts(ctx, nil, testNamespaceID, pageSize)

		assert.NotNil(t, err, "request with empty Addresses must return error")

		if resp != nil {
			assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		}
	})
}

func TestNamespaceService_GetNamespaceNames(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockServer.AddRouter(&mock.Router{
			Path: pathNamespacenames,
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

		nsInfoArr, resp, err := namespaceClient.GetNamespaceNames(ctx, *testNamespaceIDs)

		assert.Nilf(t, err, "NamespaceService.GetNamespaceNames returned error: %s", err)
		if tests.IsOkResponse(t, resp) {
			for _, nsInfo := range nsInfoArr {
				tests.ValidateStringers(t, namespaceNameCorr, nsInfo)
			}
		}
	})

	t.Run("empty namespaceIds", func(t *testing.T) {
		_, resp, err := namespaceClient.GetNamespaceNames(ctx, NamespaceIds{})

		assert.Equal(t, errEmptyNamespaceIds, err, "request with empty NamespaceIds must return error")

		if resp != nil {
			assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		}
	})
}
