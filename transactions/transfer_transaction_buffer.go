package transactions

import (
	"github.com/google/flatbuffers/go"
)

type TransferTransactionBuffer struct {
	_tab flatbuffers.Table
}

func GetRootAsTransferTransactionBuffer(buf []byte, offset flatbuffers.UOffsetT) *TransferTransactionBuffer {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &TransferTransactionBuffer{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *TransferTransactionBuffer) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *TransferTransactionBuffer) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *TransferTransactionBuffer) Size() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *TransferTransactionBuffer) MutateSize(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func (rcv *TransferTransactionBuffer) Signature(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *TransferTransactionBuffer) SignatureLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *TransferTransactionBuffer) SignatureBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *TransferTransactionBuffer) Signer(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *TransferTransactionBuffer) SignerLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *TransferTransactionBuffer) SignerBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *TransferTransactionBuffer) Version() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *TransferTransactionBuffer) MutateVersion(n uint16) bool {
	return rcv._tab.MutateUint16Slot(10, n)
}

func (rcv *TransferTransactionBuffer) Type() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *TransferTransactionBuffer) MutateType(n uint16) bool {
	return rcv._tab.MutateUint16Slot(12, n)
}

func (rcv *TransferTransactionBuffer) Fee(j int) uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *TransferTransactionBuffer) FeeLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *TransferTransactionBuffer) Deadline(j int) uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *TransferTransactionBuffer) DeadlineLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *TransferTransactionBuffer) Recipient(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *TransferTransactionBuffer) RecipientLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *TransferTransactionBuffer) RecipientBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *TransferTransactionBuffer) MessageSize() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *TransferTransactionBuffer) MutateMessageSize(n uint16) bool {
	return rcv._tab.MutateUint16Slot(20, n)
}

func (rcv *TransferTransactionBuffer) NumMosaics() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *TransferTransactionBuffer) MutateNumMosaics(n byte) bool {
	return rcv._tab.MutateByteSlot(22, n)
}

func (rcv *TransferTransactionBuffer) Message(obj *MessageBuffer) *MessageBuffer {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(MessageBuffer)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *TransferTransactionBuffer) Mosaics(obj *MosaicBuffer, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *TransferTransactionBuffer) MosaicsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func TransferTransactionBufferAddRecipient(builder *flatbuffers.Builder, recipient flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(7, flatbuffers.UOffsetT(recipient), 0)
}
func TransferTransactionBufferAddMessageSize(builder *flatbuffers.Builder, messageSize int) {
	builder.PrependUint16Slot(8, uint16(messageSize), 0)
}
func TransferTransactionBufferAddNumMosaics(builder *flatbuffers.Builder, numMosaics int) {
	builder.PrependByteSlot(9, byte(numMosaics), 0)
}
func TransferTransactionBufferAddMessage(builder *flatbuffers.Builder, message flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(10, flatbuffers.UOffsetT(message), 0)
}
func TransferTransactionBufferAddMosaics(builder *flatbuffers.Builder, mosaics flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(11, flatbuffers.UOffsetT(mosaics), 0)
}
func TransferTransactionBufferStartMosaicsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}