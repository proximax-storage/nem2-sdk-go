package sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const testNEMPublicKey = "b4f12e7c9f6946091e2cb8b6d3a12b50d17ccbbf646386ea27ce2946a7423dcf"

var testAddressesForEncoded = map[NetworkType]string{
	MIJIN_TEST: "SARNASAS2BIAB6LMFA3FPMGBPGIJGK6IJETM3ZSP",
	MIJIN:      "MARNASAS2BIAB6LMFA3FPMGBPGIJGK6IJE5K5RYU",
	TEST_NET:   "TARNASAS2BIAB6LMFA3FPMGBPGIJGK6IJE47FYR3",
	MAIN_NET:   "NARNASAS2BIAB6LMFA3FPMGBPGIJGK6IJFJKUV32",
}

func TestGenerateEncodedAddress_NEM(t *testing.T) {

	for nType, testAddress := range testAddressesForEncoded {

		res, err := generateEncodedAddress(testNEMPublicKey, nType)
		if err != nil {
			t.Fatal("Error")
		}

		assert.Equal(t, testAddress, res, "Wrong address")
	}

}
func TestGenerateEncodedAddress(t *testing.T) {
	res, err := generateEncodedAddress("321DE652C4D3362FC2DDF7800F6582F4A10CFEA134B81F8AB6E4BE78BBA4D18E", 144)
	if err != nil {
		t.Fatal("Error")
	}

	assert.Equal(t, "SBFBW6TUGLEWQIBCMTBMXXQORZKUP3WTVVTOKK5M", res, "Wrong address %s", res)

}
