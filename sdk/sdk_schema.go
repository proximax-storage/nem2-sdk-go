package sdk

type schemaAttributeSuper interface {
	serialize(buffer []byte, position int, innerObjectPosition int) []byte
}

type schema struct {
	schemaDefinition []schemaAttributeSuper
}

func (s *schema) serialize(bytes []byte) []byte {
	var resultBytes []byte

	for i, schemaDefinition := range s.schemaDefinition {
		tmp := schemaDefinition.serialize(bytes, 4+(i*2), int(bytes[0]))
		resultBytes = append(resultBytes, tmp...)
	}
	return resultBytes
}

type schemaAttribute struct {
	Name string
}

func newSchemaAttribute(name string) *schemaAttribute {
	return &schemaAttribute{Name: name}
}

func (s schemaAttribute) findParam(innerObjectPosition int, position int, buffer []byte, typeSize int) []byte {
	offset := s.offset(innerObjectPosition, position, buffer)
	if offset == 0 {
		return []byte{0}
	}
	return copyOfRange(buffer, offset+innerObjectPosition, offset+innerObjectPosition+typeSize)
}

func (s schemaAttribute) findVector(innerObjectPosition int, position int, buffer []byte, typeSize int) []byte {
	offset := s.offset(innerObjectPosition, position, buffer)
	offsetLong := offset + innerObjectPosition
	vecStart := s.vector(offsetLong, buffer)
	vecLength := s.vectorLength(offsetLong, buffer) * typeSize
	if offset == 0 {
		return []byte{0}
	}
	return copyOfRange(buffer, vecStart, vecStart+vecLength)
}

func (s schemaAttribute) findObjectStartPosition(innerObjectPosition int, position int, buffer []byte) int {
	offset := s.offset(innerObjectPosition, position, buffer)
	return s.indirect(offset+innerObjectPosition, buffer)
}

func (s schemaAttribute) findArrayLength(innerObjectPosition int, position int, buffer []byte) int {
	offset := s.offset(innerObjectPosition, position, buffer)
	if offset == 0 {
		return 0
	}
	return s.vectorLength(innerObjectPosition+offset, buffer)
}

func (s schemaAttribute) findObjectArrayElementStartPosition(innerObjectPosition int, position int, buffer []byte, startPosition int) int {
	offset := s.offset(innerObjectPosition, position, buffer)
	vector := s.vector(innerObjectPosition+offset, buffer)
	return s.indirect(vector+startPosition*4, buffer)
}

func (s schemaAttribute) readInt32(offset int, buffer []byte) uint32 {
	return uint32(buffer[offset+0]&0xFF) | uint32(buffer[offset+1]&0xFF)<<8 | uint32(buffer[offset+2]&0xFF)<<16 | uint32(buffer[offset+3])<<24
}

func (s schemaAttribute) readInt16(offset int, buffer []byte) uint16 {
	return uint16(buffer[offset+0]&0xFF) | uint16(buffer[offset+1]&0xFF)<<8

}

func (s schemaAttribute) offset(innerObjectPosition int, position int, buffer []byte) int {
	vtable := innerObjectPosition - int(s.readInt32(innerObjectPosition, buffer))
	if position < int(s.readInt16(vtable, buffer)) {
		return int(s.readInt16(vtable+position, buffer))
	}
	return 0
}

func (s schemaAttribute) vectorLength(offset int, buffer []byte) int {
	return int(s.readInt32(offset+int(s.readInt32(offset, buffer)), buffer))
}

func (s schemaAttribute) indirect(offset int, buffer []byte) int {
	return offset + int(s.readInt32(offset, buffer))
}

func (s schemaAttribute) vector(offset int, buffer []byte) int {
	return offset + int(s.readInt32(offset, buffer)) + 4
}

func copyOfRange(src []byte, from, to int) []byte {
	src = src[from:to]
	dst := make([]byte, len(src))
	copy(dst, src)
	return dst
}

type arrayAttribute struct {
	schemaAttribute *schemaAttribute
	name            string
	typeSize        int
}

//
func newArrayAttribute(name string, typeSize int) *arrayAttribute {
	schemaAttribute := newSchemaAttribute(name)
	return &arrayAttribute{schemaAttribute, name, typeSize}
}

func (s arrayAttribute) serialize(buffer []byte, position int, innerObjectPosition int) []byte {
	return s.schemaAttribute.findVector(innerObjectPosition, position, buffer, s.typeSize)
}

type scalarAttribute struct {
	SchemaAttribute *schemaAttribute
	Name            string
	TypeSize        int
}

func newScalarAttribute(name string, typeSize int) *scalarAttribute {
	schemaAttribute := newSchemaAttribute(name)
	return &scalarAttribute{schemaAttribute, name, typeSize}
}

func (s scalarAttribute) serialize(buffer []byte, position int, innerObjectPosition int) []byte {
	return s.SchemaAttribute.findParam(innerObjectPosition, position, buffer, s.TypeSize)
}

type tableArrayAttribute struct {
	schemaAttribute *schemaAttribute
	name            string
	schemaSuper     []schemaAttributeSuper
	schema          schema
}

func newTableArrayAttribute(name string, schemaSuper []schemaAttributeSuper) tableArrayAttribute {
	schemaAttribute := newSchemaAttribute(name)
	schema := schema{schemaSuper}
	return tableArrayAttribute{schemaAttribute, name, schemaSuper, schema}
}

func (s tableArrayAttribute) serialize(buffer []byte, position int, innerObjectPosition int) []byte {
	resultBytes := []byte{0}
	var arrayLength = s.schemaAttribute.findArrayLength(innerObjectPosition, position, buffer)

	for i := 1; i <= arrayLength; i++ {
		var startArrayPosition = s.schemaAttribute.findObjectArrayElementStartPosition(innerObjectPosition, position, buffer, i)
		for index, element := range s.schemaSuper {
			tmp := element.serialize(buffer, 4+index*2, startArrayPosition)
			resultBytes = append(resultBytes, tmp...)
		}
	}
	return resultBytes
}

type tableAttribute struct {
	schemaAttribute *schemaAttribute
	name            string
	schemaSuper     []schemaAttributeSuper
	schema          schema
}

func newTableAttribute(name string, schemaSuper []schemaAttributeSuper) tableAttribute {
	schemaAttribute := newSchemaAttribute(name)
	schema := schema{schemaSuper}
	return tableAttribute{schemaAttribute, name, schemaSuper, schema}
}

func (s tableAttribute) serialize(buffer []byte, position int, innerObjectPosition int) []byte {
	resultBytes := []byte{0}
	var tableStartPosition = s.schemaAttribute.findObjectStartPosition(innerObjectPosition, position, buffer)

	for index, element := range s.schemaSuper {
		tmp := element.serialize(buffer, 4+(index*2), tableStartPosition)
		resultBytes = append(resultBytes, tmp...)
	}
	return resultBytes
}

const (
	ByteSize  = 1
	ShortSize = 2
	IntSize   = 4
)
