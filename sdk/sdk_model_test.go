// Copyright 2017 Author: Ruslan Bikchentaev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sdk

import (
	"math/big"
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
	bInt = uint64DTO{12345, 99999}.toBigInt()
	want := (&big.Int{}).SetBytes([]byte("429492434645049"))
	if bInt.Uint64() != want.Uint64() {
		t.Errorf("wrong result convert DTO {12345, 99999} = %v, expected - %v", bInt, want)
	}
	bInt = uint64DTO{1111, 2222}.toBigInt()
	want = (&big.Int{}).SetBytes([]byte("9543417332823"))
	if bInt.Uint64() != want.Uint64() {
		t.Errorf("wrong result convert DTO {1111, 2222} = %v, expected - %v", bInt, want)
	}
	bInt = (&big.Int{}).SetBytes([]byte("-8884663987180930485"))
	if bInt.String() != "84b3552d375ffa4b" {
		t.Error("wrong result convert from -8884663987180930485")
	}
	bInt = (&big.Int{}).SetBytes([]byte("-3087871471161192663"))
	if bInt.String() != "d525ad41d95fcf29" {
		t.Error("wrong result convert from -3087871471161192663")
	}
}
