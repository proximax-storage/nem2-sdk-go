// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package sdk

import (
	"encoding/binary"
	"github.com/proximax-storage/nem2-sdk-go/utils"
	"math/big"
	"strconv"
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

// analog JAVA Uint64.bigIntegerToHex
func BigIntegerToHex(id *big.Int) string {

	u := FromBigInt(id)

	return strconv.FormatInt(int64(u[1]), 16) + strconv.FormatInt(int64(u[0]), 16)

}
func FromBigInt(int *big.Int) []uint32 {
	b := int.Bytes()
	ln := len(b)
	utils.ReverseByteArray(b)
	l, h, s := uint32(0), uint32(0), 4
	if ln < 4 {
		s = ln
	}
	lb := make([]byte, 4)
	copy(lb[:s], b[:s])
	l = binary.LittleEndian.Uint32(lb)
	if ln > 4 {
		if ln-4 < 4 {
			s = ln - 4
		}
		hb := make([]byte, 4)
		copy(hb[:s], b[4:])
		h = binary.LittleEndian.Uint32(hb)
	}
	return []uint32{l, h}
}

type AccountTransactionsOption struct {
	PageSize int    `url:"pageSize,omitempty"`
	Id       string `url:"id,omitempty"`
}
