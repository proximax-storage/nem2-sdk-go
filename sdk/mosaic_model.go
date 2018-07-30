package sdk

// Models

// Mosaics
type Mosaics []Mosaic

// Mosaic
type Mosaic struct {
	MosaicId []uint64 `json:"id"`
	Amount   []uint64 `json:"amount"`
}

// MosaicProperties
type MosaicProperties struct {
	SupplyMutable bool
	Transferable  bool
	LevyMutable   bool
	Divisibility  int
	Duration      uint
}

type MosaicSupplyType uint

const (
	DECREASE MosaicSupplyType = 0
	INCREASE MosaicSupplyType = 1
)

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
