package sdk

import (
	"golang.org/x/net/context"
	"testing"
)

func TestNamespaceService_GetNameSpaceInfo(t *testing.T) {
	serv := NamespaceService{}

	ctx := context.TODO()
	nsId := 0
	nsInfo, resp, err := serv.GetNameSpaceInfo(ctx, nsId)
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

func TestNamespaceService_GetAccountNameSpaceInfo(t *testing.T) {
	serv := NamespaceService{}

	ctx := context.TODO()
	nsId := 0
	pageSize := 1
	addresses := Addresses{}

	nsInfo, resp, err := serv.GetAccountNameSpaceInfo(ctx, nsId, pageSize, addresses)
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
