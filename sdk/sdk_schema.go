package sdk

import (
	"encoding/binary"
)

type schemaAttribute interface {
	serialize(buffer []byte, position int, innerObjectPosition int) []byte
}

type schema struct {
	schemaDefinition []schemaAttribute
}

func (s *schema) serialize(bytes []byte) []byte {
	var resultBytes []byte
	for i, schemaDefinition := range s.schemaDefinition {
		tmp := schemaDefinition.serialize(bytes, 4+(i*2), int(bytes[0]))
		resultBytes = append(resultBytes, tmp...)
	}
	return resultBytes
}

type abstractSchemaAttribute struct {
	Name string
}

func (s abstractSchemaAttribute) findParam(innerObjectPosition int, position int, buffer []byte, size VarSize) []byte {
	offset := s.offset(innerObjectPosition, position, buffer)
	if offset == 0 {
		return []byte{0}
	}
	return buffer[offset+innerObjectPosition : offset+innerObjectPosition+int(size)]
}

func (s abstractSchemaAttribute) findVector(innerObjectPosition int, position int, buffer []byte, size VarSize) []byte {
	offset := s.offset(innerObjectPosition, position, buffer)
	offsetLong := offset + innerObjectPosition
	vecStart := s.vector(offsetLong, buffer)
	vecLength := s.vectorLength(offsetLong, buffer) * int(size)
	if offset == 0 {
		return []byte{0}
	}
	return buffer[vecStart : vecStart+vecLength]
}

func (s abstractSchemaAttribute) findObjectStartPosition(innerObjectPosition int, position int, buffer []byte) int {
	offset := s.offset(innerObjectPosition, position, buffer)
	return s.indirect(offset+innerObjectPosition, buffer)
}

func (s abstractSchemaAttribute) findArrayLength(innerObjectPosition int, position int, buffer []byte) int {
	offset := s.offset(innerObjectPosition, position, buffer)
	if offset == 0 {
		return 0
	}
	return s.vectorLength(innerObjectPosition+offset, buffer)
}

func (s abstractSchemaAttribute) findObjectArrayElementStartPosition(innerObjectPosition int, position int, buffer []byte, startPosition int) int {
	offset := s.offset(innerObjectPosition, position, buffer)
	vector := s.vector(innerObjectPosition+offset, buffer)
	return s.indirect(vector+startPosition*4, buffer)
}

func (s abstractSchemaAttribute) readUint32(offset int, buffer []byte) int {
	return int(binary.LittleEndian.Uint32(buffer[offset : offset+4]))
}

func (s abstractSchemaAttribute) readUint16(offset int, buffer []byte) int {
	return int(binary.LittleEndian.Uint16(buffer[offset : offset+2]))
}

func (s abstractSchemaAttribute) offset(innerObjectPosition int, position int, buffer []byte) int {
	vtable := innerObjectPosition - s.readUint32(innerObjectPosition, buffer)
	if position < s.readUint16(vtable, buffer) {
		return int(s.readUint16(vtable+position, buffer))
	}
	return 0
}

func (s abstractSchemaAttribute) vectorLength(offset int, buffer []byte) int {
	return s.readUint32(offset+s.readUint32(offset, buffer), buffer)
}

func (s abstractSchemaAttribute) indirect(offset int, buffer []byte) int {
	return offset + s.readUint32(offset, buffer)
}

func (s abstractSchemaAttribute) vector(offset int, buffer []byte) int {
	return offset + s.readUint32(offset, buffer) + 4
}

type arrayAttribute struct {
	abstractSchemaAttribute
	size VarSize
}

func newArrayAttribute(name string, size VarSize) *arrayAttribute {
	return &arrayAttribute{abstractSchemaAttribute{name}, size}
}

func (s arrayAttribute) serialize(buffer []byte, position int, innerObjectPosition int) []byte {
	return s.findVector(innerObjectPosition, position, buffer, s.size)
}

type scalarAttribute struct {
	abstractSchemaAttribute
	size VarSize
}

func newScalarAttribute(name string, size VarSize) *scalarAttribute {
	return &scalarAttribute{abstractSchemaAttribute{name}, size}
}

func (s scalarAttribute) serialize(buffer []byte, position int, innerObjectPosition int) []byte {
	return s.findParam(innerObjectPosition, position, buffer, s.size)
}

type tableArrayAttribute struct {
	abstractSchemaAttribute
	schema []schemaAttribute
}

func newTableArrayAttribute(name string, schema []schemaAttribute) *tableArrayAttribute {
	return &tableArrayAttribute{abstractSchemaAttribute{name}, schema}
}

func (s tableArrayAttribute) serialize(buffer []byte, position int, innerObjectPosition int) []byte {
	var resultBytes []byte
	arrayLength := s.findArrayLength(innerObjectPosition, position, buffer)

	for i := 1; i <= arrayLength; i++ {
		startArrayPosition := s.findObjectArrayElementStartPosition(innerObjectPosition, position, buffer, i)
		for j, element := range s.schema {
			tmp := element.serialize(buffer, 4+j*2, startArrayPosition)
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

func (s tableAttribute) serialize(buffer []byte, position int, innerObjectPosition int) []byte {
	var resultBytes []byte
	var tableStartPosition = s.findObjectStartPosition(innerObjectPosition, position, buffer)

	for j, element := range s.schema {
		tmp := element.serialize(buffer, 4+(j*2), tableStartPosition)
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
