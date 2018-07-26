package sdk

import (
	"golang.org/x/net/context"
	"testing"
)

func setup() (*Config, error) {
	return LoadTestnetConfig("http://catapult.internal.proximax.io:3000")
}
func TestNamespaceService_GetNamespace(t *testing.T) {

	conf, err := setup()
	if err != nil {
		t.Fatal(err)
	}

	serv := NewNamespaceService(nil, conf)

	ctx := context.TODO()
	nsId := 0
	nsInfo, resp, err := serv.GetNamespace(ctx, nsId)
	if err != nil {
		t.Error(err)
	}

	if nsId != nsInfo.id {
		t.Error("id request & Id responce not equal")
	}

	if resp.Status != "200" {
		t.Error(resp.Status)
	}
}

func TestNamespaceService_GetNamespacesFromAccounts(t *testing.T) {

	conf, err := setup()
	if err != nil {
		t.Fatal(err)
	}

	serv := NewNamespaceService(nil, conf)

	ctx := context.TODO()
	nsId := 0
	pageSize := 1
	addresses := Addresses{}

	nsInfo, resp, err := serv.GetNamespacesFromAccounts(ctx, addresses, nsId, pageSize)
	if err != nil {
		t.Error(err)
	}

	if len(nsInfo) == 0 {
		t.Error("return result must have length > 0")
	}

	if resp.Status != "200" {
		t.Error(resp.Status)
	}

}
