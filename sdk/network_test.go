package sdk

import "testing"

func TestNetworkService_GetNetworkType(t *testing.T) {
	err := setupTest()
	if err != nil {
		t.Fatal(err)
	}

	serv := serv.client.Network
	netType, resp, err := serv.GetNetworkType(ctx)
	if err != nil {
		t.Error(err)
	} else if resp.StatusCode != 200 {
		t.Error(resp.Status)
		t.Logf("%#v", resp)
	} else if netType != MIJIN_TEST {
		t.Error("%d", netType)
	}

}
