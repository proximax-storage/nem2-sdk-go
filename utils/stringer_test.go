package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	testStructStr = "testStruct [One=Hello, Two=5432, Three=3.14, Four=true, Five=[0 0], Six=<nil>]"
)

type testStruct struct {
	One   string
	Two   int
	Three float32
	Four  bool
	Five  []int
	Six   interface{}
}

func (t *testStruct) String() string {
	return StructToString(
		"testStruct",
		NewField("One", StringPattern, t.One),
		NewField("Two", IntPattern, t.Two),
		NewField("Three", FloatPattern, t.Three),
		NewField("Four", BooleanPattern, t.Four),
		NewField("Five", ValuePattern, t.Five),
		NewField("Six", ValuePattern, t.Six),
	)
}

func TestStructToString(t *testing.T) {
	a := testStruct{One: "Hello", Two: 5432, Three: 3.14, Four: true, Five: make([]int, 2), Six: nil}

	assert.Equal(t, testStructStr, a.String())
}
