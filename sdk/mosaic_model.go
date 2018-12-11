// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package sdk

import (
	"errors"
	"fmt"
	"github.com/proximax-storage/nem2-sdk-go/utils"
	"github.com/proximax-storage/proximax-utils-go/str"
	"math/big"
	"regexp"
	"strings"
)

// MosaicId
type Mosaic struct {
	MosaicId MosaicId
	Amount   big.Int
}

func NewMosaic(mosaicId MosaicId, amount big.Int) (*Mosaic, error) {
	if utils.EqualsBigInts(mosaicIdToBigInt(&mosaicId), big.NewInt(0)) {
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

// MosaicId
type MosaicId big.Int

// MosaicId
/*type MosaicId struct {
	Id       *big.Int
	FullName string
}*/

/*func (m *MosaicId) String() string {
	return str.StructToString(
		"MosaicId",
		str.NewField("Id", str.StringPattern, m),
		str.NewField("FullName", str.StringPattern, m.FullName),
	)
}*/

func NewMosaicIdFromName(name string) (*MosaicId, error) {
	if name == "" || strings.Contains(name, " {") {
		return nil, errors.New(name + " is not valid")
	}

	parts := strings.Split(name, ":")
	if len(parts) != 2 {
		return nil, errors.New(name + " is not valid")
	}

	if id, err := generateMosaicId(parts[0], parts[1]); err != nil {
		return nil, err
	} else {
		return bigIntToMosaicId(id), nil
	}
}

func bigIntToMosaicId(bigInt *big.Int) *MosaicId {
	if bigInt == nil {
		return nil
	}

	mscId := MosaicId(*bigInt)

	return &mscId
}

func mosaicIdToBigInt(mscId *MosaicId) *big.Int {
	if mscId == nil {
		return nil
	}

	return (*big.Int)(mscId)
}

func NewMosaicId(id *big.Int) *MosaicId {
	var mscId = MosaicId(*id)

	return &mscId
}

func (m *MosaicId) toHexString() string {
	return BigIntegerToHex(mosaicIdToBigInt(m))
}

var regValidMosaicName = regexp.MustCompile(`^[a-z0-9][a-z0-9-_]*$`)

func generateMosaicId(namespaceName string, mosaicName string) (*big.Int, error) {
	if mosaicName == "" {
		return nil, errors.New(fmt.Sprintf("%s having zero length", mosaicName))
	}

	namespacePath, err := GenerateNamespacePath(namespaceName)
	if err != nil {
		return nil, err
	}

	if !regValidMosaicName.MatchString(mosaicName) {
		return nil, errors.New(mosaicName + "invalid mosaic name")
	}

	return generateId(mosaicName, namespacePath[len(namespacePath)-1])
}

// MosaicIds is a list MosaicId
type MosaicIds struct {
	MosaicIds []*MosaicId `json:"mosaicIds"`
}

func (ref *MosaicIds) MarshalJSON() (buf []byte, err error) {
	buf = []byte(`{"mosaicIds": [`)

	for i, nsId := range ref.MosaicIds {
		if i > 0 {
			buf = append(buf, ',')
		}
		buf = append(buf, []byte(`"`+nsId.toHexString()+`"`)...)
	}

	buf = append(buf, ']', '}')
	return
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
		str.NewField("Active", str.BooleanPattern, m.Active),
		str.NewField("Index", str.IntPattern, m.Index),
		str.NewField("MetaId", str.StringPattern, m.MetaId),
		str.NewField("NamespaceId", str.StringPattern, m.NamespaceId),
		str.NewField("MosaicId", str.StringPattern, m.MosaicId),
		str.NewField("Supply", str.StringPattern, m.Supply),
		str.NewField("Height", str.StringPattern, m.Height),
		str.NewField("Owner", str.StringPattern, m.Owner),
		str.NewField("Properties", str.StringPattern, m.Properties),
	)
}

type MosaicsInfo []*MosaicInfo

func (ref MosaicsInfo) String() string {
	var s string

	for i, mscInfo := range ref {
		if i > 0 {
			s += ", "
		}

		if mscInfo != nil {
			s += fmt.Sprintf("%#v", mscInfo.String())
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
	return &Mosaic{*XemMosaicId, *big.NewInt(amount)}
}

func XemRelative(amount int64) *Mosaic {
	return Xem(big.NewInt(0).Mul(big.NewInt(1000000), big.NewInt(amount)).Int64())
}
