package sdk

import (
	"encoding/binary"
)

type schemaAttribute interface {
	serialize(buffer []byte, position uint32, innerObjectPosition uint32) []byte
}

type schema struct {
	schemaDefinition []schemaAttribute
}

func (s *schema) serialize(bytes []byte) []byte {
	var resultBytes []byte
	for i, schemaDefinition := range s.schemaDefinition {
		tmp := schemaDefinition.serialize(bytes, 4+(uint32(i)*2), uint32(bytes[0]))
		resultBytes = append(resultBytes, tmp...)
	}
	return resultBytes
}

type abstractSchemaAttribute struct {
	Name string
}

func (s abstractSchemaAttribute) findParam(innerObjectPosition uint32, position uint32, buffer []byte, size VarSize) []byte {
	offset := s.offset(innerObjectPosition, position, buffer)
	if offset == 0 {
		return []byte{0}
	}
	return buffer[offset+innerObjectPosition : offset+innerObjectPosition+uint32(size)]
}

func (s abstractSchemaAttribute) findVector(innerObjectPosition uint32, position uint32, buffer []byte, size VarSize) []byte {
	offset := s.offset(innerObjectPosition, position, buffer)
	offsetLong := offset + innerObjectPosition
	vecStart := s.vector(offsetLong, buffer)
	vecLength := s.vectorLength(offsetLong, buffer) * uint32(size)
	if offset == 0 {
		return []byte{0}
	}
	return buffer[vecStart : vecStart+vecLength]
}

func (s abstractSchemaAttribute) findObjectStartPosition(innerObjectPosition uint32, position uint32, buffer []byte) uint32 {
	offset := s.offset(innerObjectPosition, position, buffer)
	return s.indirect(offset+innerObjectPosition, buffer)
}

func (s abstractSchemaAttribute) findArrayLength(innerObjectPosition uint32, position uint32, buffer []byte) uint32 {
	offset := s.offset(innerObjectPosition, position, buffer)
	if offset == 0 {
		return 0
	}
	return s.vectorLength(innerObjectPosition+offset, buffer)
}

func (s abstractSchemaAttribute) findObjectArrayElementStartPosition(innerObjectPosition uint32, position uint32, buffer []byte, startPosition uint32) uint32 {
	offset := s.offset(innerObjectPosition, position, buffer)
	vector := s.vector(innerObjectPosition+offset, buffer)
	return s.indirect(vector+startPosition*4, buffer)
}

func (s abstractSchemaAttribute) readUint32(offset uint32, buffer []byte) uint32 {
	return binary.LittleEndian.Uint32(buffer[offset : offset+4])
}

func (s abstractSchemaAttribute) readUint16(offset uint32, buffer []byte) uint16 {
	return binary.LittleEndian.Uint16(buffer[offset : offset+2])
}

func (s abstractSchemaAttribute) offset(innerObjectPosition uint32, position uint32, buffer []byte) uint32 {
	vtable := innerObjectPosition - s.readUint32(innerObjectPosition, buffer)
	if uint16(position) < s.readUint16(vtable, buffer) {
		return uint32(s.readUint16(vtable+position, buffer))
	}
	return 0
}

func (s abstractSchemaAttribute) vectorLength(offset uint32, buffer []byte) uint32 {
	return s.readUint32(offset+s.readUint32(offset, buffer), buffer)
}

func (s abstractSchemaAttribute) indirect(offset uint32, buffer []byte) uint32 {
	return offset + s.readUint32(offset, buffer)
}

func (s abstractSchemaAttribute) vector(offset uint32, buffer []byte) uint32 {
	return offset + s.readUint32(offset, buffer) + 4
}

type arrayAttribute struct {
	abstractSchemaAttribute
	size VarSize
}

func newArrayAttribute(name string, size VarSize) *arrayAttribute {
	return &arrayAttribute{abstractSchemaAttribute{name}, size}
}

func (s arrayAttribute) serialize(buffer []byte, position uint32, innerObjectPosition uint32) []byte {
	return s.findVector(innerObjectPosition, position, buffer, s.size)
}

type scalarAttribute struct {
	abstractSchemaAttribute
	size VarSize
}

func newScalarAttribute(name string, size VarSize) *scalarAttribute {
	return &scalarAttribute{abstractSchemaAttribute{name}, size}
}

func (s scalarAttribute) serialize(buffer []byte, position uint32, innerObjectPosition uint32) []byte {
	return s.findParam(innerObjectPosition, position, buffer, s.size)
}

type tableArrayAttribute struct {
	abstractSchemaAttribute
	schema []schemaAttribute
}

func newTableArrayAttribute(name string, schema []schemaAttribute) *tableArrayAttribute {
	return &tableArrayAttribute{abstractSchemaAttribute{name}, schema}
}

func (s tableArrayAttribute) serialize(buffer []byte, position uint32, innerObjectPosition uint32) []byte {
	var resultBytes []byte
	arrayLength := s.findArrayLength(innerObjectPosition, position, buffer)

	for i := uint32(0); i < arrayLength; i++ {
		startArrayPosition := s.findObjectArrayElementStartPosition(innerObjectPosition, position, buffer, i)
		for j, element := range s.schema {
			tmp := element.serialize(buffer, 4+uint32(j)*2, startArrayPosition)
			resultBytes = append(resultBytes, tmp...)
		}
	}
	return resultBytes
}

type tableAttribute struct {
	abstractSchemaAttribute
	schema []schemaAttribute
}

func newTableAttribute(name string, schema []schemaAttribute) *tableAttribute {
	return &tableAttribute{abstractSchemaAttribute{name}, schema}
}

func (s tableAttribute) serialize(buffer []byte, position uint32, innerObjectPosition uint32) []byte {
	var resultBytes []byte
	var tableStartPosition = s.findObjectStartPosition(innerObjectPosition, position, buffer)

	for j, element := range s.schema {
		tmp := element.serialize(buffer, 4+(uint32(j)*2), tableStartPosition)
		resultBytes = append(resultBytes, tmp...)
	}
	return resultBytes
}

type VarSize uint

const (
	ByteSize  VarSize = 1
	ShortSize VarSize = 2
	IntSize   VarSize = 4
)
