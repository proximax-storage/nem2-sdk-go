package crypto

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

const testNEMPublicKey = "c5247738c3a510fb6c11413331d8a47764f6e78ffcdb02b6878d5dd3b77f38ed"
const testVersion = 68
const testAddress = "NAPRILC6USCTAY7NNXB4COVKQJL427NPCEERGKS6"

func TestGenerateEncodedAddress_NEM(t *testing.T) {
	res, err := GenerateEncodedAddress(testNEMPublicKey, testVersion)
	if err != nil {
		t.Fatal("Error")
	}

	assert.Equal(t, testAddress, res, "Wrong address")

}
func TestGenerateEncodedAddress(t *testing.T) {
	res, err := GenerateEncodedAddress("321DE652C4D3362FC2DDF7800F6582F4A10CFEA134B81F8AB6E4BE78BBA4D18E", 144)
	if err != nil {
		t.Fatal("Error")
	}
	if !reflect.DeepEqual(res, "SBFBW6TUGLEWQIBCMTBMXXQORZKUP3WTVVTOKK5M") {
		t.Errorf("Wrong address %s", res)
	}
}
