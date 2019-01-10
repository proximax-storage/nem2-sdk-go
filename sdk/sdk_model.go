// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package sdk

import (
	"encoding/binary"
	"fmt"
	"math/big"
)

type uint64DTO [2]uint32

func (dto uint64DTO) toBigInt() *big.Int {
	if dto[0] == 0 && dto[1] == 0 {
		return &big.Int{}
	}
	var int big.Int
	b := make([]byte, len(dto)*4)
	binary.BigEndian.PutUint32(b[:len(dto)*2], dto[1])
	binary.BigEndian.PutUint32(b[len(dto)*2:], dto[0])
	int.SetBytes(b)
	return &int
}

func IntToHex(u uint32) string {
	s := fmt.Sprintf("%x", u)

	if len(s)%2 == 1 {
		return "0" + s
	} else {
		return s
	}
}

// analog JAVA Uint64.bigIntegerToHex
func BigIntegerToHex(id *big.Int) string {
	u := FromBigInt(id)
	return IntToHex(u[1]) + IntToHex(u[0])
}

func FromBigInt(int *big.Int) []uint32 {
	var u64 uint64 = uint64(int.Int64())
	l := uint32(u64 & 0xFFFFFFFF)
	r := uint32(u64 >> 32)
	return []uint32{l, r}
}

type AccountTransactionsOption struct {
	PageSize int    `url:"pageSize,omitempty"`
	Id       string `url:"id,omitempty"`
}
