package transactions

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type RegisterNamespaceTransactionBuffer struct {
	_tab flatbuffers.Table
}

func GetRootAsRegisterNamespaceTransactionBuffer(buf []byte, offset flatbuffers.UOffsetT) *RegisterNamespaceTransactionBuffer {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &RegisterNamespaceTransactionBuffer{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *RegisterNamespaceTransactionBuffer) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *RegisterNamespaceTransactionBuffer) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *RegisterNamespaceTransactionBuffer) Size() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *RegisterNamespaceTransactionBuffer) MutateSize(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func (rcv *RegisterNamespaceTransactionBuffer) Signature(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *RegisterNamespaceTransactionBuffer) SignatureLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *RegisterNamespaceTransactionBuffer) SignatureBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *RegisterNamespaceTransactionBuffer) Signer(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *RegisterNamespaceTransactionBuffer) SignerLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *RegisterNamespaceTransactionBuffer) SignerBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *RegisterNamespaceTransactionBuffer) Version() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *RegisterNamespaceTransactionBuffer) MutateVersion(n uint16) bool {
	return rcv._tab.MutateUint16Slot(10, n)
}

func (rcv *RegisterNamespaceTransactionBuffer) Type() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *RegisterNamespaceTransactionBuffer) MutateType(n uint16) bool {
	return rcv._tab.MutateUint16Slot(12, n)
}

func (rcv *RegisterNamespaceTransactionBuffer) Fee(j int) uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *RegisterNamespaceTransactionBuffer) FeeLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *RegisterNamespaceTransactionBuffer) Deadline(j int) uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *RegisterNamespaceTransactionBuffer) DeadlineLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *RegisterNamespaceTransactionBuffer) NamespaceType() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *RegisterNamespaceTransactionBuffer) MutateNamespaceType(n byte) bool {
	return rcv._tab.MutateByteSlot(18, n)
}

func (rcv *RegisterNamespaceTransactionBuffer) DurationParentId(j int) uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *RegisterNamespaceTransactionBuffer) DurationParentIdLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *RegisterNamespaceTransactionBuffer) NamespaceId(j int) uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *RegisterNamespaceTransactionBuffer) NamespaceIdLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *RegisterNamespaceTransactionBuffer) NamespaceNameSize() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *RegisterNamespaceTransactionBuffer) MutateNamespaceNameSize(n byte) bool {
	return rcv._tab.MutateByteSlot(24, n)
}

func (rcv *RegisterNamespaceTransactionBuffer) NamespaceName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func RegisterNamespaceTransactionBufferStart(builder *flatbuffers.Builder) {
	builder.StartObject(12)
}
func RegisterNamespaceTransactionBufferAddNamespaceType(builder *flatbuffers.Builder, namespaceType uint8) {
	builder.PrependByteSlot(7, byte(namespaceType), 0)
}
func RegisterNamespaceTransactionBufferAddDurationParentId(builder *flatbuffers.Builder, durationParentId flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(8, flatbuffers.UOffsetT(durationParentId), 0)
}
func RegisterNamespaceTransactionBufferAddNamespaceId(builder *flatbuffers.Builder, namespaceId flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(9, flatbuffers.UOffsetT(namespaceId), 0)
}
func RegisterNamespaceTransactionBufferAddNamespaceNameSize(builder *flatbuffers.Builder, namespaceNameSize int) {
	builder.PrependByteSlot(10, byte(namespaceNameSize), 0)
}
func RegisterNamespaceTransactionBufferAddNamespaceName(builder *flatbuffers.Builder, namespaceName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(11, flatbuffers.UOffsetT(namespaceName), 0)
}
