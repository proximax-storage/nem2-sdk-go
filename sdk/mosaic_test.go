package sdk

import (
	"golang.org/x/net/context"
	"testing"
)

func TestMosaicService_GetMosaic(t *testing.T) {
	conf, err := setupCFG()
	if err != nil {
		t.Fatal(err)
	}

	serv := NewMosaicService(nil, conf)
	ctx := context.TODO()
	nsInfo, resp, err := serv.GetMosaic(ctx, testIDs)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(nsInfo, resp)
	}

}
