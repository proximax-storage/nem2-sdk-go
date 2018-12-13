// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package sdk

import (
	"fmt"
	"github.com/proximax-storage/nem2-sdk-go/utils"
	"github.com/proximax-storage/proximax-utils-go/str"
	"math/big"
	"strings"
)

// MosaicId
type MosaicId big.Int

func (m *MosaicId) String() string {
	return (*big.Int)(m).String()
}

func NewMosaicIdFromName(name string) (*MosaicId, error) {
	if len(name) == 0 || strings.Contains(name, " {") {
		return nil, ErrInvalidMosaicName
	}

	parts := strings.Split(name, ":")

	if len(parts) != 2 {
		return nil, ErrInvalidMosaicName
	}

	return generateMosaicId(parts[0], parts[1])
}

func NewMosaicId(id *big.Int) (*MosaicId, error) {
	if id == nil {
		return nil, ErrNilMosaicId
	}

	return bigIntToMosaicId(id), nil
}

func (m *MosaicId) toHexString() string {
	return BigIntegerToHex(mosaicIdToBigInt(m))
}

// MosaicId
type Mosaic struct {
	MosaicId *MosaicId
	Amount   *big.Int
}

func NewMosaic(mosaicId *MosaicId, amount *big.Int) (*Mosaic, error) {
	if mosaicId == nil {
		return nil, ErrNilMosaicId
	}

	if amount == nil {
		return nil, ErrNilMosaicAmount
	}

	if utils.EqualsBigInts(amount, big.NewInt(0)) {
		return nil, ErrNilMosaicAmount
	}

	return &Mosaic{
		MosaicId: mosaicId,
		Amount:   amount,
	}, nil
}

func (m *Mosaic) String() string {
	return str.StructToString(
		"MosaicId",
		str.NewField("MosaicId", str.StringPattern, m.MosaicId),
		str.NewField("Amount", str.IntPattern, m.Amount),
	)
}

// Mosaics
type Mosaics []*Mosaic

func (ref Mosaics) String() string {
	var s string

	for i, mosaic := range ref {
		if i > 0 {
			s += ", "
		}

		if mosaic != nil {
			s += mosaic.String()
		}
	}

	return "[" + s + "]"
}

// MosaicIds is a list MosaicId
type MosaicIds struct {
	MosaicIds []*MosaicId `json:"mosaicIds"`
}

func (ref *MosaicIds) MarshalJSON() ([]byte, error) {
	buf := []byte(`{"mosaicIds": [`)

	for i, nsId := range ref.MosaicIds {
		if i > 0 {
			buf = append(buf, ',')
		}

		buf = append(buf, []byte(`"`+nsId.toHexString()+`"`)...)
	}

	buf = append(buf, ']', '}')

	return buf, nil
}

// MosaicInfo info structure contains its properties, the owner and the namespace to which it belongs to.
type MosaicInfo struct {
	MosaicId    *MosaicId
	FullName    string
	Active      bool
	Index       int
	MetaId      string
	NamespaceId *NamespaceId
	Supply      *big.Int
	Height      *big.Int
	Owner       *PublicAccount
	Properties  *MosaicProperties
}

func (m *MosaicInfo) String() string {
	return str.StructToString(
		"MosaicInfo",
		str.NewField("MosaicId", str.StringPattern, m.MosaicId),
		str.NewField("FullName", str.StringPattern, m.FullName),
		str.NewField("Active", str.BooleanPattern, m.Active),
		str.NewField("Index", str.IntPattern, m.Index),
		str.NewField("MetaId", str.StringPattern, m.MetaId),
		str.NewField("NamespaceId", str.StringPattern, m.NamespaceId),
		str.NewField("Supply", str.StringPattern, m.Supply),
		str.NewField("Height", str.StringPattern, m.Height),
		str.NewField("Owner", str.StringPattern, m.Owner),
		str.NewField("Properties", str.StringPattern, m.Properties),
	)
}

func (m *MosaicInfo) ShortName() string {
	if lastIndex := strings.LastIndex(m.FullName, ":"); lastIndex == -1 || lastIndex+1 >= len(m.FullName) {
		return ""
	} else {
		return m.FullName[lastIndex+1:]
	}
}

type MosaicsInfo []*MosaicInfo

func (ref MosaicsInfo) String() string {
	var s string

	for i, mscInfo := range ref {
		if i > 0 {
			s += ", "
		}

		if mscInfo != nil {
			s += mscInfo.String()
		}
	}

	return "[" + s + "]"
}

// MosaicProperties  structure describes mosaic properties.
type MosaicProperties struct {
	SupplyMutable bool
	Transferable  bool
	LevyMutable   bool
	Divisibility  int64
	Duration      *big.Int
}

func NewMosaicProperties(supplyMutable bool, transferable bool, levyMutable bool, divisibility int64, duration *big.Int) *MosaicProperties {
	ref := &MosaicProperties{
		supplyMutable,
		transferable,
		levyMutable,
		divisibility,
		duration,
	}

	return ref
}

func (mp *MosaicProperties) String() string {
	return str.StructToString(
		"MosaicProperties",
		str.NewField("SupplyMutable", str.BooleanPattern, mp.SupplyMutable),
		str.NewField("Transferable", str.BooleanPattern, mp.Transferable),
		str.NewField("LevyMutable", str.BooleanPattern, mp.LevyMutable),
		str.NewField("Divisibility", str.IntPattern, mp.Divisibility),
		str.NewField("Duration", str.StringPattern, mp.Duration),
	)
}

type MosaicName struct {
	MosaicId *MosaicId
	Name     string
	ParentId *NamespaceId
}

func (m *MosaicName) String() string {
	return str.StructToString(
		"MosaicName",
		str.NewField("MosaicId", str.StringPattern, m.MosaicId),
		str.NewField("Name", str.StringPattern, m.Name),
		str.NewField("ParentId", str.StringPattern, m.ParentId),
	)
}

// MosaicSupplyType mosaic supply type :
// Decrease the supply - DECREASE.
// Increase the supply - INCREASE.
type MosaicSupplyType uint8

const (
	Decrease MosaicSupplyType = iota
	Increase
)

func (tx MosaicSupplyType) String() string {
	return fmt.Sprintf("%d", tx)
}

// Create xem with using xem as unit
func Xem(amount int64) *Mosaic {
	return &Mosaic{XemMosaicId, big.NewInt(amount)}
}

func XemRelative(amount int64) *Mosaic {
	return Xem(big.NewInt(0).Mul(big.NewInt(1000000), big.NewInt(amount)).Int64())
}
