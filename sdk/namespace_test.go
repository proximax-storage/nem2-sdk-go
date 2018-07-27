package sdk

import (
	"encoding/json"
	"golang.org/x/net/context"
	"testing"
)

var testAddresses = Addresses{
	list: []*Address{
		&Address{Address: "SDRDGFTDLLCB67D4HPGIMIHPNSRYRJRT7DOBGWZY"},
		&Address{Address: "SBCPGZ3S2SCC3YHBBTYDCUZV4ZZEPHM2KGCP4QXX"},
	},
}

func setup() (*Config, error) {
	return LoadTestnetConfig("http://catapult.internal.proximax.io:3000")
}

const testIDs = "84b3552d375ffa4b"
const validResp = `{
  "Meta": {
    "active": true,
    "index": 0,
    "index": "5B55E02EACCB7B00015DB6EB"
  },
  "Namespace": {
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

	conf, err := setup()
	if err != nil {
		t.Fatal(err)
	}

	serv := NewNamespaceService(nil, conf)

	ctx := context.TODO()
	nsInfo, resp, err := serv.GetNamespace(ctx, testIDs)
	if err != nil {
		t.Error(err)
		//} else if nsId != nsInfo.index {
		//	t.Error("index request & Id responce not equal")
	} else if resp.StatusCode != 200 {
		t.Error(resp.Status)
	}
	t.Logf("%#v", nsInfo)
}

func TestNamespaceService_GetNamespacesFromAccounts(t *testing.T) {

	conf, err := setup()
	if err != nil {
		t.Fatal(err)
	}

	serv := NewNamespaceService(nil, conf)

	ctx := context.TODO()
	nsId := -1
	pageSize := -1
	nsInfo, resp, err := serv.GetNamespacesFromAccounts(ctx, &testAddresses, nsId, pageSize)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
	if len(nsInfo) == 0 {
		t.Error("return result must have length > 0")
	}

	if resp.Status != "200" {
		t.Error(resp.Status)
	}

}
