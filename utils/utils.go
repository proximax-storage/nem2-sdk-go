package utils

import (
	"encoding/hex"
	"math/big"
)

func ReverseByteArray(a []byte) {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
}

func HexDecode(s string) ([]byte, error) {
	if len(s)%2 != 0 {
		s = "0" + s
	}
	return hex.DecodeString(s)
}

//BigIntToByteArray converts a BigInteger to a little endian byte array.
func BigIntToByteArray(value *big.Int, numBytes int) []byte {

	// output must tohave lenght NumBytes!
	outputBytes := make([]byte, numBytes)
	bigIntegerBytes := value.Bytes()
	copyStartIndex := 0
	if 0x00 == bigIntegerBytes[0] {
		copyStartIndex = 1
	}
	numBytesToCopy := len(bigIntegerBytes) - copyStartIndex
	if numBytesToCopy > numBytes {
		copyStartIndex += numBytesToCopy - numBytes
		numBytesToCopy = numBytes
	}

	//reverse & copy
	for i := 0; i < numBytesToCopy; i++ {
		outputBytes[i] = bigIntegerBytes[copyStartIndex+numBytesToCopy-i-1]
	}

	return outputBytes
}

//BytesToBigInteger converts a little endian byte array to a BigInteger.
func BytesToBigInteger(bytes []byte) *big.Int {

	bigEndianBytes := make([]byte, len(bytes)+1)
	//reverse & copy
	for i := range bytes {
		bigEndianBytes[i+1] = bytes[len(bytes)-i-1]
	}

	return (&big.Int{}).SetBytes(bigEndianBytes)
}
