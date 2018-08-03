package sdk

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
)

// Models

// Mosaic
type Mosaic struct {
	MosaicId MosaicId   `json:"id"`
	Amount   *uint64DTO `json:"amount"`
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
type Mosaics []Mosaic

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
	Id       *uint64DTO
	FullName string
}

func NewMosaicId(id *uint64DTO, name string) (*MosaicId, error) {
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
	Supply      *uint64DTO
	Height      *uint64DTO
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
	Divisibility  int
	Duration      uint
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

type NamespaceMosaicMetaDTO struct {
	Active bool
	Index  int
	Id     string
} /* NamespaceMosaicMetaDTO */
func (ref *NamespaceInfoDTO) extractLevels() []*NamespaceId {

	levels := make([]*NamespaceId, 3)
	var err error

	nsId := NewNamespaceId(ref.Namespace.Level0, "")
	if err == nil {
		levels = append(levels, nsId)
	}

	nsId = NewNamespaceId(ref.Namespace.Level1, "")
	if err == nil {
		levels = append(levels, nsId)
	}

	nsId = NewNamespaceId(ref.Namespace.Level2, "")
	if err == nil {
		levels = append(levels, nsId)
	}
	return levels
}
