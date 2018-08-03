// Copyright 2017 Author: Ruslan Bikchentaev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sdk

import "math/big"

type uint64DTO [2]*big.Int

func (ref *uint64DTO) ExtractIntArray() uint64 {
	return ref[1].Uint64()<<4 + ref[0].Uint64()
}
func (ref *uint64DTO) String() string {
	return ref[1].String() + ref[0].String()
}

// NewUint64DTO create uint64DTO from byte sequence
func NewUint64DTO(lowByte, highByte []byte) *uint64DTO {
	low := &big.Int{}
	if lowByte == nil {
		low = nil
	} else {
		low = low.SetBytes(lowByte)
	}
	high := &big.Int{}
	if highByte == nil {
		high = nil
	} else {
		high = high.SetBytes(highByte)
	}
	return &uint64DTO{low, high}
}
func SetUint64DTO(l, h int64) *uint64DTO {
	low := big.NewInt(l)
	high := big.NewInt(h)
	return &uint64DTO{low, high}
}
func NewRootUint64DTO() *uint64DTO {
	low := big.NewInt(0)
	high := big.NewInt(0)
	return &uint64DTO{low, high}
}
