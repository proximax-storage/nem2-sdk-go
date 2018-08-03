package sdk

import (
	"math/big"
	"bytes"
	"encoding/binary"
)

type uint64DTO [2]uint64

func (dto uint64DTO) toStruct() *big.Int {
	var int big.Int
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, &dto)
	int.SetBytes(buf.Bytes())
	return &int
}

type AccountTransactionsOption struct {
	PageSize int `url:"pageSize,omitempty"`
	Id string `url:"id,omitempty"`
}



