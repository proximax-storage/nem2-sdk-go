package utils

import (
	"github.com/stretchr/testify/assert"
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
