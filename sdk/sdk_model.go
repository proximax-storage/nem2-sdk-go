package sdk

import (
	"bytes"
	"encoding/binary"
	"math/big"
)

type uint64DTO [2]uint64

func (dto uint64DTO) toStruct() *big.Int { //TODO Needs refactoring and testing
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
