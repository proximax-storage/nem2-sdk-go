package utils

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

var (
	testBytesOdd = []byte("ABCD")
	wantBytesOdd = []byte("DCBA")
)

func TestReverseByteArray(t *testing.T) {
	ReverseByteArray(testBytesOdd)
	assert.Equal(t, wantBytesOdd, testBytesOdd)
}

func TestBigIntToByteArray(t *testing.T) {
	v := big.NewInt(0)
	b := BigIntToByteArray(v, 32)

	assert.Equal(t, make([]byte, 32), b)
}
