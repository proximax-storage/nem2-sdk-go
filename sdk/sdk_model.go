package sdk

import (
	"encoding/binary"
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

type AccountTransactionsOption struct {
	PageSize int    `url:"pageSize,omitempty"`
	Id       string `url:"id,omitempty"`
}
