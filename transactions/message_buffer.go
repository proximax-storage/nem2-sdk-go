package transactions

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type MessageBuffer struct {
	_tab flatbuffers.Table
}

func GetRootAsMessageBuffer(buf []byte, offset flatbuffers.UOffsetT) *MessageBuffer {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &MessageBuffer{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *MessageBuffer) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *MessageBuffer) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *MessageBuffer) Type() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MessageBuffer) MutateType(n byte) bool {
	return rcv._tab.MutateByteSlot(4, n)
}

func (rcv *MessageBuffer) Payload(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *MessageBuffer) PayloadLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *MessageBuffer) PayloadBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func MessageBufferStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func MessageBufferAddType(builder *flatbuffers.Builder, type_ uint8) {
	builder.PrependByteSlot(0, byte(type_), 0)
}
func MessageBufferAddPayload(builder *flatbuffers.Builder, payload flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(payload), 0)
}
func MessageBufferEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
