// Copyright 2017 Author: Ruslan Bikchentaev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sdk

import (
	"testing"
)

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
	bIntArr := fromBigInt(bInt)
	want := []uint32{1094650402, 17}
	if bIntArr[0] != want[0] || bIntArr[1] != want[1] {
		t.Errorf("wrong result convert DTO {12345, 99999} = %v, expected - %v", bIntArr, want)
	}
	bInt = uint64DTO{1111, 2222}.toBigInt()
	bIntArr = fromBigInt(bInt)
	want = []uint32{1111, 2222}
	if bIntArr[0] != want[0] || bIntArr[1] != want[1] {
		t.Errorf("wrong result convert DTO {1111, 2222} = %v, expected - %v", bInt, want)
	}
}
