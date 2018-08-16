package transactions

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type MosaicBuffer struct {
	_tab flatbuffers.Table
}

func GetRootAsMosaicBuffer(buf []byte, offset flatbuffers.UOffsetT) *MosaicBuffer {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &MosaicBuffer{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *MosaicBuffer) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *MosaicBuffer) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *MosaicBuffer) Id(j int) uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *MosaicBuffer) IdLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *MosaicBuffer) Amount(j int) uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *MosaicBuffer) AmountLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func MosaicBufferStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func MosaicBufferAddId(builder *flatbuffers.Builder, id flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(id), 0)
}
func MosaicBufferCreateIdVector(builder *flatbuffers.Builder, data []uint32) flatbuffers.UOffsetT {
	builder.StartVector(4, len(data), 4)
	for i := len(data) - 1; i >= 0; i-- {
		builder.PrependUint32(data[i])
	}
	return builder.EndVector(len(data))
}
func MosaicBufferAddAmount(builder *flatbuffers.Builder, amount flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(amount), 0)
}
func MosaicBufferCreateAmountVector(builder *flatbuffers.Builder, data []uint32) flatbuffers.UOffsetT {
	builder.StartVector(4, len(data), 4)
	for i := len(data) - 1; i >= 0; i-- {
		builder.PrependUint32(data[i])
	}
	return builder.EndVector(len(data))
}
func MosaicBufferEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
