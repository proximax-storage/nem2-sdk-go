package utils

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

var (
	testBytesOdd  = []byte("AB0CD")
	wantBytesOdd  = []byte("DC0BA")
	testBytesOdd1 = []byte("AB0CD1")
	wantBytesOdd1 = []byte("1DC0BA")
)

func TestReverseByteArray(t *testing.T) {
	ReverseByteArray(testBytesOdd)
	assert.Equal(t, wantBytesOdd, testBytesOdd)
	ReverseByteArray(testBytesOdd1)
	assert.Equal(t, wantBytesOdd1, testBytesOdd1)
}

func TestBigIntToByteArray(t *testing.T) {
	v := big.NewInt(0)
	b := BigIntToByteArray(v, 32)

	assert.Equal(t, make([]byte, 32), b)
}

//using different numbers from original javs sdk because of signed and unsigned transformation
//ex. uint64(-8884663987180930485) = 9562080086528621131
func TestBigIntegerToHex_bigIntegerNEMAndXEMToHex(t *testing.T) {
	testBigInt(t, "9562080086528621131", "84b3552d375ffa4b")
	testBigInt(t, "15358872602548358953", "d525ad41d95fcf29")
}
func testBigInt(t *testing.T, str, hexStr string) {
	i, ok := (&big.Int{}).SetString(str, 10)
	assert.True(t, ok)
	result := BigIntegerToHex(i)
	assert.Equal(t, hexStr, result)

}
