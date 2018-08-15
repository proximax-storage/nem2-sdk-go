package transactions

import (
	"github.com/google/flatbuffers/go"
)

type MosaicSupplyChangeTransactionBuffer struct {
	_tab flatbuffers.Table
}

func GetRootAsMosaicSupplyChangeTransactionBuffer(buf []byte, offset flatbuffers.UOffsetT) *MosaicSupplyChangeTransactionBuffer {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &MosaicSupplyChangeTransactionBuffer{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *MosaicSupplyChangeTransactionBuffer) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *MosaicSupplyChangeTransactionBuffer) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *MosaicSupplyChangeTransactionBuffer) Size() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MosaicSupplyChangeTransactionBuffer) MutateSize(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func (rcv *MosaicSupplyChangeTransactionBuffer) Signature(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *MosaicSupplyChangeTransactionBuffer) SignatureLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *MosaicSupplyChangeTransactionBuffer) SignatureBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *MosaicSupplyChangeTransactionBuffer) Signer(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *MosaicSupplyChangeTransactionBuffer) SignerLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *MosaicSupplyChangeTransactionBuffer) SignerBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *MosaicSupplyChangeTransactionBuffer) Version() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MosaicSupplyChangeTransactionBuffer) MutateVersion(n uint16) bool {
	return rcv._tab.MutateUint16Slot(10, n)
}

func (rcv *MosaicSupplyChangeTransactionBuffer) Type() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MosaicSupplyChangeTransactionBuffer) MutateType(n uint16) bool {
	return rcv._tab.MutateUint16Slot(12, n)
}

func (rcv *MosaicSupplyChangeTransactionBuffer) Fee(j int) uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *MosaicSupplyChangeTransactionBuffer) FeeLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *MosaicSupplyChangeTransactionBuffer) Deadline(j int) uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *MosaicSupplyChangeTransactionBuffer) DeadlineLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *MosaicSupplyChangeTransactionBuffer) MosaicId(j int) uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *MosaicSupplyChangeTransactionBuffer) MosaicIdLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *MosaicSupplyChangeTransactionBuffer) Direction() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MosaicSupplyChangeTransactionBuffer) MutateDirection(n byte) bool {
	return rcv._tab.MutateByteSlot(20, n)
}

func (rcv *MosaicSupplyChangeTransactionBuffer) Delta(j int) uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *MosaicSupplyChangeTransactionBuffer) DeltaLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func MosaicSupplyChangeTransactionBufferAddMosaicId(builder *flatbuffers.Builder, mosaicId flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(7, flatbuffers.UOffsetT(mosaicId), 0)
}
func MosaicSupplyChangeTransactionBufferAddDirection(builder *flatbuffers.Builder, direction uint8) {
	builder.PrependByteSlot(8, byte(direction), 0)
}
func MosaicSupplyChangeTransactionBufferAddDelta(builder *flatbuffers.Builder, delta flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(9, flatbuffers.UOffsetT(delta), 0)
}