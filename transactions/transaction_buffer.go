package transactions

import "github.com/google/flatbuffers/go"

func TransactionBufferAddSize(builder *flatbuffers.Builder, size int) {
	builder.PrependUint32Slot(0, uint32(size), 0)
}
func TransactionBufferAddSignature(builder *flatbuffers.Builder, signature flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(signature), 0)
}
func TransactionBufferAddSigner(builder *flatbuffers.Builder, signer flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(signer), 0)
}
func TransactionBufferAddVersion(builder *flatbuffers.Builder, version uint64) {
	builder.PrependUint16Slot(3, uint16(version), 0)
}
func TransactionBufferAddType(builder *flatbuffers.Builder, type_ uint16) {
	builder.PrependUint16Slot(4, type_, 0)
}
func TransactionBufferAddFee(builder *flatbuffers.Builder, fee flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(fee), 0)
}
func TransactionBufferAddDeadline(builder *flatbuffers.Builder, deadline flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(deadline), 0)
}
func TransactionBufferEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
func TransactionBufferCreateUint32Vector(builder *flatbuffers.Builder, data []uint32) flatbuffers.UOffsetT {
	builder.StartVector(4, len(data), 4)
	for i := len(data) - 1; i >= 0; i-- {
		builder.PrependUint32(data[i])
	}
	return builder.EndVector(len(data))
}
func TransactionBufferCreateByteVector(builder *flatbuffers.Builder, data []byte) flatbuffers.UOffsetT {
	builder.StartVector(1, len(data), 1)
	for i := len(data) - 1; i >= 0; i-- {
		builder.PrependByte(data[i])
	}
	return builder.EndVector(len(data))
}
func TransactionBufferCreateUOffsetVector(builder *flatbuffers.Builder, data []flatbuffers.UOffsetT) flatbuffers.UOffsetT {
	TransferTransactionBufferStartMosaicsVector(builder, len(data))
	for i := len(data) - 1; i >= 0; i-- {
		builder.PrependUOffsetT(data[i])
	}
	return builder.EndVector(len(data))
}
