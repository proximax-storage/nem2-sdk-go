package sdk

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/big"
	"net/http"
	"testing"
)

func init() {
	addRouters(mscRouters)
	i, _ := (&big.Int{}).SetString("15358872602548358953", 10)
	testMosaicId.Id = i
}

var (
	testMosaicId  = &MosaicId{}
	testMosaicIds = MosaicIds{MosaicIds: []*MosaicId{testMosaicId}}
)

const testMosaicPathID = "d525ad41d95fcf29"
const testMosaicFromNamesaceId = "5B55E02EACCB7B00015DB6EC"

var (
	tplMosaic = `{
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
	mscRouters = map[string]sRouting{
		pathMosaic + testMosaicPathID: {tplMosaic, nil},
		pathMosaic:                    {"[" + tplMosaic + "]", routeNeedBody},
		pathMosaicNames: {`[
						  {
							"mosaicId": [
							  3646934825,
							  3576016193
							],
							"name": "xem",
							"parentId": [
							  929036875,
							  2226345261
							]
						  }
						]`, routeNeedBody},
		fmt.Sprintf(pathMosaicFromNamespace, mosaicNamespace): {"[" + tplMosaic + "]", nil},
	}
)

func TestMosaicService_GetMosaic(t *testing.T) {

	mscInfo, resp, err := serv.Mosaic.GetMosaic(ctx, *testMosaicId)
	if err != nil {
		t.Error(err)
	} else if validateResp(resp, t) && validateMosaicInfo(mscInfo, t) {
		t.Logf("%#v", mscInfo)
	}

}
func TestMosaicService_GetMosaics(t *testing.T) {

	mscInfoArr, resp, err := serv.Mosaic.GetMosaics(ctx, testMosaicIds)
	if err != nil {
		t.Error(err)
	} else if validateResp(resp, t) {
		isValid := true
		for _, mscInfo := range mscInfoArr {
			isValid = isValid && validateMosaicInfo(mscInfo, t)
		}
		if isValid {
			t.Logf("%s", mscInfoArr)
		}
	}

	mscInfoArr, resp, err = serv.Mosaic.GetMosaics(ctx, MosaicIds{})

	assert.NotNil(t, err, "request with empty MosaicIds must return error")
	if resp != nil {
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	}

}

func TestMosaicService_GetMosaicNames(t *testing.T) {

	mscInfoArr, resp, err := serv.Mosaic.GetMosaicNames(ctx, testMosaicIds)
	if err != nil {
		t.Error(err)
	} else if validateResp(resp, t) {
		t.Logf("%s", mscInfoArr)

	}
}

func TestMosaicService_GetMosaicsFromNamespace(t *testing.T) {

	mscInfoArr, resp, err := serv.Mosaic.GetMosaicsFromNamespace(ctx, testNamespaceId, testMosaicId, pageSize)
	if err != nil {
		t.Error(err)
	} else if validateResp(resp, t) {
		t.Logf("%v", mscInfoArr)

	}

}

func validateMosaicInfo(mscInfo *MosaicInfo, t *testing.T) bool {
	result := true

	if !assert.NotNil(t, mscInfo) {
		result = false
	} else if metaId := mscInfo.MetaId; metaId != "5B55E02EACCB7B00015DB6EC" {
		t.Error(fmt.Sprintf("failed MetaId data Convertion = '%s' (%#v)", metaId, mscInfo))
		result = false
	} else if fullname := mscInfo.NamespaceId.FullName; fullname != "" {
		t.Error(fmt.Sprintf("failed namespaseName data Convertion = '%s' (%#v)", fullname, mscInfo))
		result = false
	} else if !mscInfo.Active {
		t.Error(fmt.Sprintf("failed Active data Convertion = '%v' (%#v)", mscInfo.Active, mscInfo))
		result = false
	} else if nsId := mscInfo.NamespaceId.Id; !(nsId.Uint64() == uint64DTO{929036875, 2226345261}.toBigInt().Uint64()) {
		t.Error(fmt.Sprintf("failed NamespaceId data Convertion = '%v' (%#v)", nsId, mscInfo))
		result = false
	} else if mscId := mscInfo.MosaicId.Id; !(mscId.Uint64() == uint64DTO{3646934825, 3576016193}.toBigInt().Uint64()) {
		t.Error(fmt.Sprintf("failed MosaicId data Convertion = '%v' (%#v)", mscId, mscInfo))
		result = false
	} else if nsId := mscInfo.Supply; !(nsId.Uint64() == uint64DTO{3403414400, 2095475}.toBigInt().Uint64()) {
		t.Error(fmt.Sprintf("failed Supply data Convertion = '%v' (%#v)", nsId, mscInfo))
		result = false
	} else if nsId := mscInfo.Height; !(nsId.Uint64() == 1) {
		t.Error(fmt.Sprintf("failed Height data Convertion = '%v' (%#v)", nsId, mscInfo))
		result = false
	} else if publicKey := mscInfo.Owner.PublicKey; publicKey != "321DE652C4D3362FC2DDF7800F6582F4A10CFEA134B81F8AB6E4BE78BBA4D18E" {
		t.Error(fmt.Sprintf("failed Owner data Convertion = '%s' (%#v)", publicKey, mscInfo))
		result = false
	}
	return result
}
