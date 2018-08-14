package sdk

import (
	"bytes"
	"encoding/binary"
	"math/big"
)

type uint64DTO [2]uint64

func (dto uint64DTO) GetBigInteger() *big.Int { //TODO Needs refactoring and testing
	if dto[0] == 0 && dto[1] == 0 {
		return &big.Int{}
	}
	var int big.Int
	buf := new(bytes.Buffer)
	dto[0], dto[1] = dto[1], dto[0]
	binary.Write(buf, binary.BigEndian, &dto)
	int.SetBytes(buf.Bytes())
	return &int
}

type AccountTransactionsOption struct {
	PageSize int    `url:"pageSize,omitempty"`
	Id       string `url:"id,omitempty"`
}
