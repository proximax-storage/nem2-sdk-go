// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package utils

import (
	"encoding/hex"
	"math/big"
)

func ReverseByteArray(a []byte) {
	//for i := len(a)/2 - 1; i >= 0; i-- {
	//	opp := len(a) - 1 - i
	//	a[i], a[opp] = a[opp], a[i]
	//}
	lenA := len(a)
	j := lenA
	for i := range a[lenA/2:] {
		j--
		a[i], a[j] = a[j], a[i]
	}
}
func MustHexDecodeString(s string) []byte {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return b
}
func HexDecodeStringOdd(s string) ([]byte, error) {
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

	if len(bigIntegerBytes) == 0 {
		return outputBytes
	}
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

// BytesToBigInteger converts a little endian byte array to a BigInteger.
func BytesToBigInteger(bytes []byte) *big.Int {

	bigEndianBytes := make([]byte, len(bytes)+1)
	//reverse & copy
	for i := range bytes {
		bigEndianBytes[i+1] = bytes[len(bytes)-i-1]
	}

	return (&big.Int{}).SetBytes(bigEndianBytes)
}

func HexDecode(src []byte) ([]byte, error) {
	dst := make([]byte, len(src))
	_, err := hex.Decode(dst, src)
	if err != nil {
		return nil, err
	}
	return dst, nil
}

func EqualsBigInts(first, second *big.Int) bool {
	if first == nil && second == nil {
		return true
	}

	if first != nil {
		return first.Cmp(second) == 0
	}

	return second.Cmp(first) == 0
}
