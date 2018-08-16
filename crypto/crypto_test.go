package crypto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateEncodedAddress(t *testing.T) {
	res, err := GenerateEncodedAddress("321DE652C4D3362FC2DDF7800F6582F4A10CFEA134B81F8AB6E4BE78BBA4D18E", 144)
	if err != nil {
		t.Fatal("Error")
	}

	assert.Equal(t, "SBFBW6TUGLEWQIBCMTBMXXQORZKUP3WTVVTOKK5M", res, "Wrong address %s", res)

}
