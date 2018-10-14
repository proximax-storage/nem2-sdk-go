// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package sdk

import (
	"fmt"
	"github.com/json-iterator/go"
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
}

var (
	testAddresses = Addresses{
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
	ad = &NamespaceIds{}
)

const pageSize = 32
const mosaicNamespace = "84b3552d375ffa4b"

var (
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
	tplInfoArr = "[" + tplInfo + "]"

	nsRouters = map[string]sRouting{
		pathNamespace: {tplInfo, nil},
		pathNamespacenames: {`[
			  {
				"namespaceId": [
				  929036875,
				  2226345261
				],
				"name": "nem"
			  }
			]`, routeNeedBody},
		pathNamespacesFromAccounts:                                  {tplInfoArr, routeNeedBody},
		fmt.Sprintf(pathNamespacesFromAccount, testAddress.Address): {tplInfoArr, nil},
	}
)

func init() {
	addRouters(nsRouters)
}

const testIDs = "84b3552d375ffa4b"

func TestNamespaceService_GetNamespace(t *testing.T) {

	nsInfo, resp, err := serv.Namespace.GetNamespace(ctx, testNamespaceId)
	if err != nil {
		t.Error(err)
	} else if validateResp(resp, t) && validateNamespaceInfo(nsInfo, t) {
		t.Logf("%s", nsInfo)
	}
}

const testNamespaceID = "5B55E02EACCB7B00015DB6EB"

func TestNamespaceService_GetNamespacesFromAccount(t *testing.T) {

	nsInfoArr, resp, err := serv.Namespace.GetNamespacesFromAccount(ctx, &testAddress, testNamespaceID, pageSize)
	if err != nil {
		t.Error(err)
	} else if validateResp(resp, t) {
		if len(nsInfoArr.List) != 1 {
			t.Error("return result must have length = 1")
		} else {
			isValid := true
			for _, nsInfo := range nsInfoArr.List {
				isValid = isValid && validateNamespaceInfo(nsInfo, t)
			}
			if isValid {
				t.Logf("%v", nsInfoArr)
			}
		}
	}

	nsInfoArr, resp, err = serv.Namespace.GetNamespacesFromAccount(ctx, nil, testNamespaceID, pageSize)
	assert.NotNil(t, err, "request with empty Address must return error")
}
func TestNamespaceService_GetNamespacesFromAccounts(t *testing.T) {

	nsInfoArr, resp, err := serv.Namespace.GetNamespacesFromAccounts(ctx, &testAddresses, testNamespaceID, pageSize)
	if err != nil {
		t.Error(err)
	} else if validateResp(resp, t) {
		if len(nsInfoArr.List) != 1 {
			t.Error("return result must have length = 1")
		} else {
			isValid := true
			for _, nsInfo := range nsInfoArr.List {
				isValid = isValid && validateNamespaceInfo(nsInfo, t)
			}
			if isValid {
				t.Logf("%v", nsInfoArr)
			}
		}
	}

	nsInfoArr, resp, err = serv.Namespace.GetNamespacesFromAccounts(ctx, nil, testNamespaceID, pageSize)
	assert.NotNil(t, err, "request with empty Addresses must return error")
	if resp != nil {
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	}
}

func TestNamespaceService_GetNamespaceNames(t *testing.T) {

	nsInfo, resp, err := serv.Namespace.GetNamespaceNames(ctx, *testNamespaceIDs)
	if err != nil {
		t.Fatal(err)
	} else if validateResp(resp, t) {
		if (nsInfo == nil) || (len(nsInfo) == 0) {
			t.Logf("%#v %#v", resp, resp.Body)
		} else if arr0 := (nsInfo)[0]; (arr0.NamespaceId == nil) || (arr0.NamespaceId.Id == nil) {
			t.Logf("%#v", arr0)
		} else {
			if id := arr0.NamespaceId.Id; !(id.Uint64() == uint64DTO{929036875, 2226345261}.toBigInt().Uint64()) {
				t.Error("failed namespaceName id Convertion")
				t.Logf("%s", id)
			}
			if arr0.Name != "nem" {
				t.Error("failed namespaceName Name Convertion")
				t.Logf("%#v", arr0.Name)
			}
		}
	}
	t.Logf("%#v", nsInfo)

	nsInfo, resp, err = serv.Namespace.GetNamespaceNames(ctx, NamespaceIds{})
	assert.Equal(t, errEmptyNamespaceIds, err, "request with empty NamespaceIds must return error")
	if resp != nil {
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	}
}

func validateNamespaceInfo(nsInfo *NamespaceInfo, t *testing.T) bool {
	result := true
	if !nsInfo.Active {
		t.Error("failed Active data Convertion")
		result = false
	}
	if !(nsInfo.Index == 0) {
		t.Error("failed Index data Convertion")
		result = false
	}
	if !(nsInfo.MetaId == "5B55E02EACCB7B00015DB6EB") {
		t.Error("failed Id data Convertion")
		result = false
	}
	if !(nsInfo.TypeSpace == Root) {
		t.Error("failed Type data Convertion")
		result = false
	}
	if !(nsInfo.Depth == 1) {
		t.Error("failed Depth data Convertion")
		result = false
	}
	if !(nsInfo.Owner.PublicKey == "321DE652C4D3362FC2DDF7800F6582F4A10CFEA134B81F8AB6E4BE78BBA4D18E") {
		t.Error("failed Owner data Convertion")
		result = false
	}
	if nsId := nsInfo.ParentId.Id; !(nsId.Uint64() == 0) {
		t.Error("failed ParentId data Convertion")
		result = false
	}
	if sH := nsInfo.StartHeight; !(sH.Uint64() == 1) {
		t.Error("failed ParentId data Convertion")
		result = false
	}
	if eH := nsInfo.EndHeight; !(eH.Uint64() == uint64DTO{4294967295, 4294967295}.toBigInt().Uint64()) {
		t.Error("failed ParentId data Convertion")
		result = false
	}

	return result
}
