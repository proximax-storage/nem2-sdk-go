package sdk

import (
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"testing"
)

const validMosaicResp = `{
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

func validateMosaicInfo(nsInfo *MosaicInfo) error {
	if nsInfo == nil {
		return errors.New("return nil structure nsInfo")
	} else if metaId := nsInfo.MetaId; metaId != "5B55E02EACCB7B00015DB6EC" {
		return errors.New(fmt.Sprintf("failed MetaId data Convertion = '%s' (%#v)", metaId, nsInfo))
	} else if fullname := nsInfo.NamespaceId.FullName; fullname != "" {
		return errors.New(fmt.Sprintf("failed namespaseName data Convertion = '%s' (%#v)", fullname, nsInfo))
	} else if !nsInfo.Active {
		return errors.New(fmt.Sprintf("failed Active data Convertion = '%v' (%#v)", nsInfo.Active, nsInfo))
	} else if nsId := nsInfo.NamespaceId.Id; !(nsId[0].Int64() == 929036875 && nsId[1].Int64() == 2226345261) {
		return errors.New(fmt.Sprintf("failed Id data Convertion = '%v' (%#v)", nsId, nsInfo))
	} else if nsId := nsInfo.MosaicId.id; !(nsId[0].Int64() == 3646934825 && nsId[1].Int64() == 3576016193) {
		return errors.New(fmt.Sprintf("failed MosaicId data Convertion = '%v' (%#v)", nsId, nsInfo))
	} else if nsId := nsInfo.Supply; !(nsId[0].Int64() == 3403414400 && nsId[1].Int64() == 2095475) {
		return errors.New(fmt.Sprintf("failed Supply data Convertion = '%v' (%#v)", nsId, nsInfo))
	} else if nsId := nsInfo.Height; !(nsId[0].Int64() == 1 && nsId[1].Int64() == 0) {
		return errors.New(fmt.Sprintf("failed Height data Convertion = '%v' (%#v)", nsId, nsInfo))
	} else if publicKey := nsInfo.Owner.PublicKey; publicKey != "321DE652C4D3362FC2DDF7800F6582F4A10CFEA134B81F8AB6E4BE78BBA4D18E" {
		return errors.New(fmt.Sprintf("failed Owner data Convertion = '%s' (%#v)", publicKey, nsInfo))
	}
	return nil
}
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
	} else if err := validateMosaicInfo(nsInfo); err != nil {
		t.Error(err)
	} else {
		t.Logf("%#v", nsInfo)
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
	} else if err := validateMosaicInfo(nsInfo[0]); err != nil {
		t.Error(err)
	} else {
		t.Log(nsInfo, resp)
	}

}
