package crypto

import (
	"testing"
	"reflect"
)

func TestGenerateEncodedAddress(t *testing.T) {
	res, err := GenerateEncodedAddress("321DE652C4D3362FC2DDF7800F6582F4A10CFEA134B81F8AB6E4BE78BBA4D18E", 144)
	if err != nil {
		t.Error("Error")
	}

	if !reflect.DeepEqual(res, "SBFBW6TUGLEWQIBCMTBMXXQORZKUP3WTVVTOKK5M") {
		t.Errorf("Wrong address %s", res)
	}
}