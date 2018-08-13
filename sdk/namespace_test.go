package sdk

import (
	"fmt"
	"github.com/json-iterator/go"
	"net/http"
	"testing"
)

var (
	testAddresses = Addresses{
		list: []*Address{
			{Address: "SDRDGFTDLLCB67D4HPGIMIHPNSRYRJRT7DOBGWZY"},
			{Address: "SBCPGZ3S2SCC3YHBBTYDCUZV4ZZEPHM2KGCP4QXX"},
		},
	}
	testAddress = Address{Address: "SCASIIAPS6BSFEC66V6MU5ZGEVWM53BES5GYBGLE"}

	testNamespaceIDs = &NamespaceIds{
		List: []*NamespaceId{
			{FullName: "84b3552d375ffa4b"},
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
	if !(nsInfo.TypeSpace == RootNamespace) {
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
	if nsId := nsInfo.ParentId.Id; !(nsId == uint64DTO{0, 0}.toStruct()) {
		t.Error("failed ParentId data Convertion")
		result = false
	}
	if sH := nsInfo.StartHeight; !(sH == uint64DTO{1, 0}.toStruct()) {
		t.Error("failed ParentId data Convertion")
		result = false
	}
	if eH := nsInfo.EndHeight; !(eH == uint64DTO{4294967295, 4294967295}.toStruct()) {
		t.Error("failed ParentId data Convertion")
		result = false
	}

	return result
}

const testIDs = "84b3552d375ffa4b"

func TestNamespaceService_GetNamespace(t *testing.T) {

	nsInfo, resp, err := serv.Namespace.GetNamespace(ctx, testIDs)
	if err != nil {
		t.Error(err)
	} else if resp.StatusCode != 200 {
		t.Error(resp.Status)
		t.Logf("%#v", resp)
	} else if validateNamespaceInfo(nsInfo, t) {
		t.Logf("%s", nsInfo)
	}
}

const testNamespaceID = "5B55E02EACCB7B00015DB6EB"

func TestNamespaceService_GetNamespacesFromAccount(t *testing.T) {

	nsInfoArr, resp, err := serv.Namespace.GetNamespacesFromAccount(ctx, &testAddress, testNamespaceID, pageSize)
	if err != nil {
		t.Error(err)
	} else if resp.StatusCode != 200 {
		t.Error(resp.Status)
		t.Logf("%#v %#v", resp, resp.Body)

		b := make([]byte, resp.ContentLength)
		if _, err := resp.Body.Read(b); err == nil {
			t.Logf("%s", b)
		} else {
			t.Error(err)
		}
	} else if len(nsInfoArr.list) != 1 {
		t.Error("return result must have length = 1")
	} else {
		isValid := true
		for _, nsInfo := range nsInfoArr.list {
			isValid = isValid && validateNamespaceInfo(nsInfo, t)
		}
		if isValid {
			t.Logf("%v", nsInfoArr)
		}
	}

	nsInfoArr, resp, err = serv.Namespace.GetNamespacesFromAccount(ctx, nil, testNamespaceID, pageSize)
	if err == nil {
		t.Error("addrees is null - method must return error")
	} else if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Error responce status code = %d", resp.StatusCode)
	}

}
func TestNamespaceService_GetNamespacesFromAccounts(t *testing.T) {

	nsInfoArr, resp, err := serv.Namespace.GetNamespacesFromAccounts(ctx, &testAddresses, testNamespaceID, pageSize)
	if err != nil {
		t.Error(err)
	} else if resp.StatusCode != 200 {
		t.Error(resp.Status)
		t.Logf("%#v %#v", resp, resp.Body)

		b := make([]byte, resp.ContentLength)
		if _, err := resp.Body.Read(b); err == nil {
			t.Logf("%s", b)
		} else {
			t.Error(err)
		}
	} else if len(nsInfoArr.list) != 1 {
		t.Error("return result must have length = 1")
	} else {
		isValid := true
		for _, nsInfo := range nsInfoArr.list {
			isValid = isValid && validateNamespaceInfo(nsInfo, t)
		}
		if isValid {
			t.Logf("%v", nsInfoArr)
		}
	}
	nsInfoArr, resp, err = serv.Namespace.GetNamespacesFromAccounts(ctx, nil, testNamespaceID, pageSize)
	if err != nil {
		t.Error(err)
	} else if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Error responce status code = %d", resp.StatusCode)
	}
}

func init() {
	jsoniter.RegisterTypeEncoder("*NamespaceIds", testNamespaceIDs)
	jsoniter.RegisterTypeDecoder("*NamespaceIds", testNamespaceIDs)
	jsoniter.RegisterTypeDecoder("*NamespaceIds", ad)

}
func TestNamespaceService_GetNamespaceNames(t *testing.T) {

	nsInfo, resp, err := serv.Namespace.GetNamespaceNames(ctx, testNamespaceIDs)
	if err != nil {
		t.Fatal(err)
	} else if resp.StatusCode != 200 {
		t.Error(resp.Status)
		t.Logf("%#v %#v", resp, resp.Body)
	} else if (nsInfo == nil) || (len(nsInfo) == 0) {
		t.Logf("%#v %#v", resp, resp.Body)
	} else if arr0 := (nsInfo)[0]; (arr0.NamespaceId == nil) || (arr0.NamespaceId.Id == nil) {
		t.Logf("%#v", arr0)
	} else {
		if id := arr0.NamespaceId.Id; !(id == uint64DTO{929036875, 2226345261}.toStruct()) {
			t.Error("failed namespaceName id Convertion")
			t.Logf("%s", id)
		}
		if arr0.Name != "nem" {
			t.Error("failed namespaceName Name Convertion")
			t.Logf("%#v", arr0.Name)
		}
	}
	t.Logf("%#v", nsInfo)

}
