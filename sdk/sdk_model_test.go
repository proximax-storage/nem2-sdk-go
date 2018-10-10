// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package sdk

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestUint64DTO_big(t *testing.T) {
	bInt, ok := (&big.Int{}).SetString("9543417332823", 10)
	assert.True(t, ok)
	bIntArr := FromBigInt(bInt)
	want := []uint32{1111, 2222}
	assert.Equal(t, want, bIntArr)

	bInt, ok = (&big.Int{}).SetString("429492434645049", 10)
	assert.True(t, ok)
	bIntArr = FromBigInt(bInt)
	want = []uint32{12345, 99999}
	assert.Equal(t, want, bIntArr)
}
func TestUint64DTO_GetBigInteger(t *testing.T) {
	bInt := uint64DTO{0, 0}.toBigInt()
	if bInt.Uint64() != 0 {
		t.Error("wrong result convert nulled DTO")
	}
	bInt = uint64DTO{1, 0}.toBigInt()
	if bInt.Uint64() != 1 {
		t.Error("wrong result convert DTO {1, 0}")
	}
	bInt = uint64DTO{100, 0}.toBigInt()
	if bInt.Uint64() != 100 {
		t.Error("wrong result convert DTO {100, 0}")
	}
	bInt = uint64DTO{1000, 0}.toBigInt()
	if bInt.Uint64() != 1000 {
		t.Error("wrong result convert DTO {1000, 0}")
	}
	bInt = uint64DTO{10000, 0}.toBigInt()
	if bInt.Uint64() != 10000 {
		t.Error("wrong result convert DTO {10000, 0}")
	}

	//todo: check algoritm set BigInteger from string - test don't work
	bInt = uint64DTO{1094650402, 17}.toBigInt()
	bIntArr := FromBigInt(bInt)
	want := []uint32{1094650402, 17}
	if bIntArr[0] != want[0] || bIntArr[1] != want[1] {
		t.Errorf("wrong result convert DTO {12345, 99999} = %v, expected - %v", bIntArr, want)
	}
	bInt = uint64DTO{1111, 2222}.toBigInt()
	bIntArr = FromBigInt(bInt)
	want = []uint32{1111, 2222}
	if bIntArr[0] != want[0] || bIntArr[1] != want[1] {
		t.Errorf("wrong result convert DTO {1111, 2222} = %v, expected - %v", bInt, want)
	}
}
