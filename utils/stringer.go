package utils

import (
	"fmt"
	"strings"
)

type FieldPattern string

const (
	StringPattern  FieldPattern = "s"
	IntPattern     FieldPattern = "d"
	BooleanPattern FieldPattern = "t"
	FloatPattern   FieldPattern = "g"
	ValuePattern   FieldPattern = "v"
)

type Field struct {
	fieldName string
	pattern   FieldPattern
	value     interface{}
}

func NewField(fieldName string, pattern FieldPattern, val interface{}) *Field {
	return &Field{
		fieldName: fieldName,
		pattern:   pattern,
		value:     val,
	}
}

func (f *Field) String() string {
	return fmt.Sprintf(f.fieldName+"=%"+string(f.pattern), f.value)
}

func StructToString(structName string, fields ...*Field) string {
	if len(fields) == 0 {
		return ""
	}

	values := make([]string, len(fields))

	for idx, field := range fields {
		values[idx] = field.String()
	}

	return fmt.Sprintf("%s [%s]", structName, strings.Join(values, ", "))
}
