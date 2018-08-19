package transactions

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type CosignatoryModificationBuffer struct {
	_tab flatbuffers.Table
}

func GetRootAsCosignatoryModificationBuffer(buf []byte, offset flatbuffers.UOffsetT) *CosignatoryModificationBuffer {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &CosignatoryModificationBuffer{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *CosignatoryModificationBuffer) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *CosignatoryModificationBuffer) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *CosignatoryModificationBuffer) Type() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *CosignatoryModificationBuffer) MutateType(n byte) bool {
	return rcv._tab.MutateByteSlot(4, n)
}

func (rcv *CosignatoryModificationBuffer) CosignatoryPublicKey(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *CosignatoryModificationBuffer) CosignatoryPublicKeyLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *CosignatoryModificationBuffer) CosignatoryPublicKeyBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func CosignatoryModificationBufferStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func CosignatoryModificationBufferAddType(builder *flatbuffers.Builder, type_ uint8) {
	builder.PrependByteSlot(0, byte(type_), 0)
}
func CosignatoryModificationBufferAddCosignatoryPublicKey(builder *flatbuffers.Builder, cosignatoryPublicKey flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(cosignatoryPublicKey), 0)
}
