// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

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
