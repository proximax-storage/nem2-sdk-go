package sdk

import (
	"fmt"
	"github.com/json-iterator/go"
	"golang.org/x/net/context"
	"net/http"
	"testing"
	"time"
)

var (
	testAddresses = Addresses{
		Addresses: []*Address{
			&Address{Address: "SDRDGFTDLLCB67D4HPGIMIHPNSRYRJRT7DOBGWZY"},
			&Address{Address: "SBCPGZ3S2SCC3YHBBTYDCUZV4ZZEPHM2KGCP4QXX"},
		},
	}
	testAddress = Address{Address: "SCASIIAPS6BSFEC66V6MU5ZGEVWM53BES5GYBGLE"}
)

const pageSize = 32

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
	routers    = map[string]string{
		pathNamespace: tplInfo,
		pathNamespacenames: `[
			  {
				"namespaceId": [
				  929036875,
				  2226345261
				],
				"name": "nem"
			  }
			]`,
		pathNamespacesFromAccounts:                                  tplInfoArr,
		fmt.Sprintf(pathNamespacesFromAccount, testAddress.Address): tplInfoArr,
	}
)

// const for test routing
var (
	serv *NamespaceService
	ctx  = context.TODO()
)

func setupTest() error {
	if serv != nil {
		return nil
	}
	client, mux, _, teardown, err := setupMockServer()
	if err != nil {
		return err
	}
	time.AfterFunc(time.Minute*5, teardown)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Mock JSON response
		w.Write([]byte("unknow route"))
	})
	for path, resp := range routers {
		resp := []byte(resp)
		mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			// Mock JSON response
			w.Write(resp)
		})

	}

	serv = client.Namespace
	return nil
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
	if nsId := nsInfo.ParentId.Id; !(nsId[0].Int64() == 0 && nsId[1].Int64() == 0) {
		t.Error("failed ParentId data Convertion")
		result = false
	}
	if sH := nsInfo.StartHeight; !(sH[0].Int64() == 1 && sH[1].Int64() == 0) {
		t.Error("failed ParentId data Convertion")
		result = false
	}
	if eH := nsInfo.EndHeight; !(eH[0].Int64() == 4294967295 && eH[1].Int64() == 4294967295) {
		t.Error("failed ParentId data Convertion")
		result = false
	}

	return result
}

const testIDs = "84b3552d375ffa4b"

func TestNamespaceService_GetNamespace(t *testing.T) {

	err := setupTest()
	if err != nil {
		t.Fatal(err)
	}

	nsInfo, resp, err := serv.GetNamespace(ctx, testIDs)
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

	err := setupTest()
	if err != nil {
		t.Fatal(err)
	}

	nsInfoArr, resp, err := serv.GetNamespacesFromAccount(ctx, &testAddress, testNamespaceID, pageSize)
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
			t.Logf("%s", nsInfoArr)
		}
	}
}
func TestNamespaceService_GetNamespacesFromAccounts(t *testing.T) {

	err := setupTest()
	if err != nil {
		t.Fatal(err)
	}

	nsInfoArr, resp, err := serv.GetNamespacesFromAccounts(ctx, &testAddresses, testNamespaceID, pageSize)
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
			t.Logf("%s", nsInfoArr)
		}
	}
}

var testNamespaceIDs = &NamespaceIds{
	List: []*NamespaceId{
		{FullName: "84b3552d375ffa4b"},
	},
}
var ad = &NamespaceIds{}

func init() {
	jsoniter.RegisterTypeEncoder("*NamespaceIds", testNamespaceIDs)
	jsoniter.RegisterTypeDecoder("*NamespaceIds", testNamespaceIDs)
	jsoniter.RegisterTypeDecoder("*NamespaceIds", ad)

}
func TestNamespaceService_GetNamespaceNames(t *testing.T) {
	err := setupTest()
	if err != nil {
		t.Fatal(err)
	}

	nsInfo, resp, err := serv.GetNamespaceNames(ctx, testNamespaceIDs)
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
		if id := arr0.NamespaceId.Id; !((id[0].Int64() == 929036875) && (id[1].Int64() == 2226345261)) {
			t.Error("failed namespaceName id Convertion")
			t.Logf("%d %d", id[0].Int64(), id[1].Int64())
		}
		if arr0.Name != "nem" {
			t.Error("failed namespaceName Name Convertion")
			t.Logf("%#v", arr0.Name)
		}
	}
	t.Logf("%#v", nsInfo)

}
