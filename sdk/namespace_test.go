package sdk

import (
	"bytes"
	"github.com/json-iterator/go"
	"golang.org/x/net/context"
	"testing"
)

var testAddresses = Addresses{
	Addresses: []*Address{
		&Address{Address: "SDRDGFTDLLCB67D4HPGIMIHPNSRYRJRT7DOBGWZY"},
		&Address{Address: "SBCPGZ3S2SCC3YHBBTYDCUZV4ZZEPHM2KGCP4QXX"},
	},
}

func setupCFG() (*Config, error) {
	return LoadTestnetConfig("http://catapult.internal.proximax.io:3000")
}

const testIDs = "84b3552d375ffa4b"
const validResp = `{
  "meta": {
    "active": true,
    "index": 0,
    "id": "5B55E02EACCB7B00015DB6EB"
  },
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

func TestNewNamespaceInfoDTO(t *testing.T) {
	nsDTO := &NamespaceInfoDTO{}
	err := json.Unmarshal([]byte(validResp), &nsDTO)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%#v", nsDTO)
	}
}
func TestNamespaceService_GetNamespace(t *testing.T) {

	conf, err := setupCFG()
	if err != nil {
		t.Fatal(err)
	}

	serv := NewNamespaceService(nil, conf)

	ctx := context.TODO()
	nsInfo, resp, err := serv.GetNamespace(ctx, testIDs)
	if err != nil {
		t.Error(err)
	} else if resp.StatusCode != 200 {
		t.Error(resp.Status)
		t.Logf("%#v", resp)
	} else {
		if !nsInfo.active {
			t.Error("failed Active data Convertion")

		}
		if !(nsInfo.index == 0) {
			t.Error("failed Index data Convertion")

		}
		if !(nsInfo.metaId == "5B55E02EACCB7B00015DB6EB") {
			t.Error("failed Id data Convertion")
		}
		if !(nsInfo.typeSpace == RootNamespace) {
			t.Error("failed Type data Convertion")
		}
		if !(nsInfo.depth == 1) {
			t.Error("failed Depth data Convertion")
		}
		if !(nsInfo.owner.PublicKey == "321DE652C4D3362FC2DDF7800F6582F4A10CFEA134B81F8AB6E4BE78BBA4D18E") {
			t.Error("failed Owner data Convertion")
		}
		if nsId := nsInfo.parentId.id; !(nsId[0].Int64() == 0 && nsId[1].Int64() == 0) {
			t.Error("failed ParentId data Convertion")
		}
		if sH := nsInfo.startHeight; !(sH[0].Int64() == 1 && sH[1].Int64() == 0) {
			t.Error("failed ParentId data Convertion")
		}
		if eH := nsInfo.endHeight; !(eH[0].Int64() == 4294967295 && eH[1].Int64() == 4294967295) {
			t.Error("failed ParentId data Convertion")
		}
	}
	t.Logf("%s", nsInfo)
}

var testNamespaceIDs = &NamespaceIds{
	List: []*NamespaceId{
		{fullName: "84b3552d375ffa4b"},
	},
}
var ad = &NamespaceIds{}

func init() {
	jsoniter.RegisterTypeEncoder("*NamespaceIds", testNamespaceIDs)
	jsoniter.RegisterTypeDecoder("*NamespaceIds", testNamespaceIDs)
	jsoniter.RegisterTypeDecoder("*NamespaceIds", ad)

}
func TestNamespaceIds_MarshalJSON(t *testing.T) {

	b, err := json.Marshal(testNamespaceIDs)
	if err != nil {
		t.Fatal(err)
	}

	b1, err := testNamespaceIDs.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(b, b1) {
		t.Error("not equal")
	}
	t.Log("standart", string(b))
	t.Log("self-made", string(b1))

	err = json.Unmarshal(b1, ad)

	if err != nil {
		t.Error(err)
	} else {
		t.Log(ad)
	}
	err = json.Unmarshal(b, ad)

	if err != nil {
		t.Error(err)
	} else {
		t.Log(ad)
	}

}
func TestNamespaceService_GetNamespaceNames(t *testing.T) {
	conf, err := setupCFG()
	if err != nil {
		t.Fatal(err)
	}

	serv := NewNamespaceService(nil, conf)

	ctx := context.TODO()
	nsInfo, resp, err := serv.GetNamespaceNames(ctx, testNamespaceIDs)
	if err != nil {
		t.Fatal(err)
	} else if resp.StatusCode != 200 {
		t.Error(resp.Status)
		t.Logf("%#v %#v", resp, resp.Body)
	} else if (nsInfo == nil) || (len(nsInfo) == 0) {
		t.Logf("%#v %#v", resp, resp.Body)
	} else if arr0 := (nsInfo)[0]; (arr0.namespaceId == nil) || (arr0.namespaceId.id == nil) {
		t.Logf("%#v", arr0)
	} else {
		if id := arr0.namespaceId.id; !((id[0].Int64() == 929036875) && (id[1].Int64() == 2226345261)) {
			t.Error("failed namespaceName id Convertion")
			t.Logf("%#v", id[0].Int64(), id[1].Int64())
		}
		if arr0.name != "nem" {
			t.Error("failed namespaceName name Convertion")
			t.Logf("%#v", arr0.name)
		}
	}

}
func TestNamespaceService_GetNamespacesFromAccounts(t *testing.T) {

	conf, err := setupCFG()
	if err != nil {
		t.Fatal(err)
	}

	serv := NewNamespaceService(nil, conf)

	ctx := context.TODO()
	nsId := "5B55E02EACCB7B00015DB6EB"
	pageSize := 32
	nsInfo, resp, err := serv.GetNamespacesFromAccounts(ctx, &testAddresses, nsId, pageSize)
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
	} else {
		if len(nsInfo.list) != 0 {
			t.Error("return result must have length = 0")
		}

	}
}
