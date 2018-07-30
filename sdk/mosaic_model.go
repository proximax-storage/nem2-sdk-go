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