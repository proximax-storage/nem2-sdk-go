package sdk

import "fmt"

// Models

// Mosaic
type Mosaic struct {
	MosaicId []uint64 `json:"id"`
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
