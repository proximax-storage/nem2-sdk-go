package sdk

import (
	"fmt"
	"testing"
)

var mosaicTest = MosaicIds{MosaicIds: []string{"d525ad41d95fcf29"}}

const testMosaicID = "d525ad41d95fcf29"
const testMosaicFromNamesaceId = "5B55E02EACCB7B00015DB6EC"

func validateMosaicInfo(nsInfo *MosaicInfo, t *testing.T) bool {
	result := true

	if nsInfo == nil {
		t.Error("return nil structure nsInfo")
		result = false
	} else if metaId := nsInfo.MetaId; metaId != "5B55E02EACCB7B00015DB6EC" {
		t.Error(fmt.Sprintf("failed MetaId data Convertion = '%s' (%#v)", metaId, nsInfo))
		result = false
	} else if fullname := nsInfo.NamespaceId.FullName; fullname != "" {
		t.Error(fmt.Sprintf("failed namespaseName data Convertion = '%s' (%#v)", fullname, nsInfo))
		result = false
	} else if !nsInfo.Active {
		t.Error(fmt.Sprintf("failed Active data Convertion = '%v' (%#v)", nsInfo.Active, nsInfo))
		result = false
	//} else if nsId := nsInfo.NamespaceId.Id; !(nsId[0].Int64() == 929036875 && nsId[1].Int64() == 2226345261) {
	//	t.Error(fmt.Sprintf("failed Id data Convertion = '%v' (%#v)", nsId, nsInfo))
	//	result = false
	//} else if nsId := nsInfo.MosaicId.Id; !(nsId[0].Int64() == 3646934825 && nsId[1].Int64() == 3576016193) {
	//	t.Error(fmt.Sprintf("failed MosaicId data Convertion = '%v' (%#v)", nsId, nsInfo))
	//	result = false
	//} else if nsId := nsInfo.Supply; !(nsId[0].Int64() == 3403414400 && nsId[1].Int64() == 2095475) {
	//	t.Error(fmt.Sprintf("failed Supply data Convertion = '%v' (%#v)", nsId, nsInfo))
	//	result = false
	//} else if nsId := nsInfo.Height; !(nsId[0].Int64() == 1 && nsId[1].Int64() == 0) {
	//	t.Error(fmt.Sprintf("failed Height data Convertion = '%v' (%#v)", nsId, nsInfo))
	//	result = false
	} else if publicKey := nsInfo.Owner.PublicKey; publicKey != "321DE652C4D3362FC2DDF7800F6582F4A10CFEA134B81F8AB6E4BE78BBA4D18E" {
		t.Error(fmt.Sprintf("failed Owner data Convertion = '%s' (%#v)", publicKey, nsInfo))
		result = false
	}
	return result
}
func TestMosaicService_GetMosaic(t *testing.T) {
	err := setupTest()
	if err != nil {
		t.Fatal(err)
	}

	serv := serv.client.Mosaic
	mscInfo, resp, err := serv.GetMosaic(ctx, testMosaicID)
	if err != nil {
		t.Error(err)
	} else if resp.StatusCode != 200 {
		t.Error(resp.Status)
		t.Logf("%#v", resp)
	} else if validateMosaicInfo(mscInfo, t) {
		t.Logf("%#v", mscInfo)
	}

}
func TestMosaicService_GetMosaics(t *testing.T) {
	err := setupTest()
	if err != nil {
		t.Fatal(err)
	}

	serv := serv.client.Mosaic
	mscInfoArr, resp, err := serv.GetMosaics(ctx, mosaicTest)
	if err != nil {
		t.Error(err)
	} else if resp.StatusCode != 200 {
		t.Error(resp.Status)
		t.Logf("%#v", resp)
	} else {
		isValid := true
		for _, mscInfo := range mscInfoArr {
			isValid = isValid && validateMosaicInfo(mscInfo, t)
		}
		if isValid {
			t.Logf("%s", mscInfoArr)
		}
	}

}

var testMosaicIds = &MosaicIds{
	[]string{
		"d525ad41d95fcf29",
	}}

func TestMosaicService_GetMosaicNames(t *testing.T) {
	err := setupTest()
	if err != nil {
		t.Fatal(err)
	}

	serv := serv.client.Mosaic
	mscInfoArr, resp, err := serv.GetMosaicNames(ctx, testMosaicIds)
	if err != nil {
		t.Error(err)
	} else if resp.StatusCode != 200 {
		t.Error(resp.Status)
		t.Logf("%#v", resp)
	} else {
		t.Logf("%s", mscInfoArr)

	}
}
func TestMosaicService_GetMosaicsFromNamespace(t *testing.T) {
	err := setupTest()
	if err != nil {
		t.Fatal(err)
	}

	serv := serv.client.Mosaic
	mscInfoArr, resp, err := serv.GetMosaicsFromNamespace(ctx, mosaicNamespace, testMosaicFromNamesaceId, pageSize)
	if err != nil {
		t.Error(err)
	} else if resp.StatusCode != 200 {
		t.Error(resp.Status)
		t.Logf("%#v", resp)
	} else {
		t.Logf("%v", mscInfoArr)

	}

}
