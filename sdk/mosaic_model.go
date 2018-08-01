package sdk

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
)

// Models

// Mosaics
type Mosaics []Mosaic

// Mosaic
type Mosaic struct {
	MosaicId MosaicId `json:"id"`
	Amount   []uint64 `json:"amount"`
}

func (m *Mosaic) String() string {
	return fmt.Sprintf(
		`
			"MosaicId": %d,
			"Amount": %d 
		`,
		m.MosaicId,
		m.Amount,
	)
}

// MosaicId
type MosaicId struct {
	id       *uint64DTO
	fullName string
}

func NewMosaicFromID(id *uint64DTO) *MosaicId {
	return &MosaicId{id, ""}
}
func NewMosaicId(name string) (*MosaicId, error) {
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

func generateMosaicId(namespaceName string, mosaicName string) (*uint64DTO, error) {

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

type MosaicIds struct {
	MosaicIds []string `json:"mosaicIds"`
}

// MosaicInfo
type MosaicInfo struct {
	active      bool
	index       int
	metaId      string
	namespaceId *NamespaceId
	mosaicId    *MosaicId
	supply      *uint64DTO
	height      *uint64DTO
	owner       *PublicAccount
	properties  *MosaicProperties
}

// MosaicProperties
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

type MosaicSupplyType uint

const (
	DECREASE MosaicSupplyType = 0
	INCREASE MosaicSupplyType = 1
)

func (tx MosaicSupplyType) String() string {
	return fmt.Sprintf("%d", tx)
}
