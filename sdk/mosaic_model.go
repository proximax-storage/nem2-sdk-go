package sdk

import (
	"errors"
	"fmt"
	"math/big"
	"regexp"
	"strings"
	"time"
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
	if networkTypeError != nil {
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
	namespaceId := namespacePath[len(namespacePath)-1]
	if !regValidMosaicName.MatchString(mosaicName) {
		return nil, errors.New(mosaicName + "invalid mosaic name")
	}

	return generateId(mosaicName, namespaceId)
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
	Duration      time.Duration
}

func NewMosaicProperties(supplyMutable bool, transferable bool, levyMutable bool, divisibility int64, duration time.Duration) *MosaicProperties {
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

//MosaicSupplyType mosaic supply type :
// Decrease the supply - DECREASE.
//Increase the supply - INCREASE.
type MosaicSupplyType uint

const (
	DECREASE MosaicSupplyType = iota
	INCREASE
)

func (tx MosaicSupplyType) String() string {
	return fmt.Sprintf("%d", tx)
}

type MosaicName struct {
	MosaicId *MosaicId
	Name     string
	ParentId *NamespaceId
}
