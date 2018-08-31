package sdk

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"regexp"
	"strings"
)

// Models

// Mosaic
type Mosaic struct {
	MosaicId *MosaicId
	Amount   *big.Int
}

func (m *Mosaic) String() string {
	return fmt.Sprintf(
		`
			"MosaicId": %v,
			"Amount": %d 
		`,
		m.MosaicId,
		m.Amount,
	)
}

// Mosaics
type Mosaics []*Mosaic

func (ref Mosaics) String() string {
	s := "["
	for i, mosaic := range ref {
		if i > 0 {
			s += ", "
		}
		s += mosaic.String()
	}
	return s + "]"
}

// MosaicId
type MosaicId struct {
	Id       *big.Int
	FullName string
}

func NewMosaicId(id *big.Int, name string) (*MosaicId, error) {
	if id != nil {
		return &MosaicId{id, ""}, nil
	}

	if (name == "") || strings.Contains(name, " {") {
		return nil, errors.New(name + " is not valid")
	}
	parts := strings.Split(name, ":")
	if len(parts) != 2 {
		return nil, errors.New(name + " is not valid")

	}
	id, err := generateMosaicId(parts[0], parts[1])
	if err != nil {
		return nil, err
	}
	return &MosaicId{id, name}, nil
}

var regValidMosaicName = regexp.MustCompile(`^[a-z0-9][a-z0-9-_]*$`)

func generateMosaicId(namespaceName string, mosaicName string) (*big.Int, error) {

	if mosaicName == "" {
		return nil, errors.New(mosaicName + " having zero length")
	}

	namespacePath, err := generateNamespacePath(namespaceName)
	if err != nil {
		return nil, err
	}

	if !regValidMosaicName.MatchString(mosaicName) {
		return nil, errors.New(mosaicName + "invalid mosaic name")
	}

	b, err := generateId(mosaicName, namespacePath[len(namespacePath)-1])
	return new(big.Int).SetBytes(b), err
}

// MosaicIds is a list MosaicId
type MosaicIds struct {
	MosaicIds []string `json:"mosaicIds"`
}

// MosaicInfo info structure contains its properties, the owner and the namespace to which it belongs to.
type MosaicInfo struct {
	Active      bool
	Index       int
	MetaId      string
	NamespaceId *NamespaceId
	MosaicId    *MosaicId
	Supply      *big.Int
	Height      *big.Int
	Owner       *PublicAccount
	Properties  *MosaicProperties
}
type MosaicsInfo []*MosaicInfo

func (ref MosaicsInfo) String() string {
	s := "["
	for i, mscInfo := range ref {
		if i > 0 {
			s += ", "
		}
		s += fmt.Sprintf("%#v", mscInfo)

	}

	return s
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
	return fmt.Sprintf(
		`
			"SupplyMutable": %t,
			"Transferable": %t,
			"LevyMutable": %t,
			"Divisibility": %d,
			"Duration": %d
		`,
		mp.SupplyMutable,
		mp.Transferable,
		mp.LevyMutable,
		mp.Divisibility,
		mp.Duration,
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

type MosaicName struct {
	MosaicId *MosaicId
	Name     string
	ParentId *NamespaceId
}

var XemMosaicId, _ = NewMosaicId(nil, "nem:xem") // NewMosaicId(big.NewInt(0).SetBytes([]byte{213, 37, 173, 65, 217, 95, 207, 41}), "")

func Xem(amount int64) *Mosaic {
	return &Mosaic{XemMosaicId, big.NewInt(amount)}
}

func XemRelative(amount int64) *Mosaic {
	i := big.NewFloat(math.Pow10(6))
	i2, _ := i.Int64()
	return Xem(big.NewInt(0).Mul(big.NewInt(i2), big.NewInt(amount)).Int64())
}
