package sdk

import "fmt"

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
	id       int64
	fullName string
}

// MosaicInfo
type MosaicInfo struct {
	active      bool
	index       int
	metaId      string
	namespaceId NamespaceId
	mosaicId    MosaicId
	supply      int64
	height      int64
	owner       PublicAccount
	properties  MosaicProperties
}

// MosaicProperties
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
