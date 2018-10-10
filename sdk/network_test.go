package sdk

import "testing"

func init() {
	addRouters(map[string]sRouting{pathNetwork: {`{
  "name": "mijinTest",
  "description": "catapult development network"
}`, nil}})

}
func TestNetworkService_GetNetworkType(t *testing.T) {

	netType, resp, err := serv.Network.GetNetworkType(ctx)
	if err != nil {
		t.Error(err)
	} else if resp.StatusCode != 200 {
		t.Error(resp.Status)
		t.Logf("%#v", resp)
	} else if netType != MijinTest {
		t.Errorf("%d", netType)
	}

}
func TestExtractNetworkType(t *testing.T) {
	i := uint64(36888)

	nt := ExtractNetworkType(i)
	if nt != MijinTest {
		t.Errorf("wrong convert %d (%d - %d)", i, nt, MijinTest)
	}

}
