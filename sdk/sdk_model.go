package sdk

import (
	"encoding/binary"
	"github.com/proximax-storage/nem2-sdk-go/utils"
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

func fromBigInt(int *big.Int) []uint32 {
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
