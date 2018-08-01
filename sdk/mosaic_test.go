package sdk

import (
	"golang.org/x/net/context"
	"testing"
)

const valkMosaicResp = `{
  "meta": {
    "active": true,
    "index": 0,
    "id": "5B55E02EACCB7B00015DB6EC"
  },
  "mosaic": {
    "namespaceId": [
      929036875,
      2226345261
    ],
    "mosaicId": [
      3646934825,
      3576016193
    ],
    "supply": [
      3403414400,
      2095475
    ],
    "height": [
      1,
      0
    ],
    "owner": "321DE652C4D3362FC2DDF7800F6582F4A10CFEA134B81F8AB6E4BE78BBA4D18E",
    "properties": [
      [
        2,
        0
      ],
      [
        6,
        0
      ],
      [
        0,
        0
      ]
    ],
    "levy": {}
  }
}`

var mosaicTest = MosaicIds{MosaicIds: []string{"d525ad41d95fcf29"}}

const testMosaicID = "d525ad41d95fcf29"

func TestMosaicService_GetMosaic(t *testing.T) {
	conf, err := setupCFG()
	if err != nil {
		t.Fatal(err)
	}

	serv := NewMosaicService(nil, conf)
	ctx := context.TODO()
	nsInfo, resp, err := serv.GetMosaic(ctx, testMosaicID)
	if err != nil {
		t.Error(err)
	} else if resp.StatusCode != 200 {
		t.Error(resp.Status)
		t.Logf("%#v", resp)
	} else {
		t.Log(nsInfo, resp)
	}

}
func TestMosaicService_GetMosaics(t *testing.T) {
	conf, err := setupCFG()
	if err != nil {
		t.Fatal(err)
	}

	serv := NewMosaicService(nil, conf)
	ctx := context.TODO()
	nsInfo, resp, err := serv.GetMosaics(ctx, mosaicTest)
	if err != nil {
		t.Error(err)
	} else if resp.StatusCode != 200 {
		t.Error(resp.Status)
		t.Logf("%#v", resp)
	} else {
		t.Log(nsInfo, resp)
	}

}
