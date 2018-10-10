// Copyright 2018 ProximaX Limited. All rights reserved. // Use of this source code is governed by the Apache 2.0 // license that can be found in the LICENSE file.  package sdk

import (
	"reflect"
	"testing"
)

var b = []byte{32, 0, 0, 0, 28, 0, 44, 0, 40, 0, 36, 0, 32, 0, 30, 0, 28, 0, 24, 0, 20, 0, 16, 0, 12, 0, 15, 0, 8, 0, 4, 0, 28, 0, 0, 0,
	200, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 28, 0, 0, 0, 68, 0, 0, 0, 52, 0, 0, 0, 1, 65, 3, 144, 68, 0, 0, 0, 100, 0, 0, 0, 166, 0, 0,
	0, 25, 0, 0, 0, 144, 232, 254, 189, 103, 29, 212, 27, 238, 148, 236, 59,
	165, 131, 28, 182, 8, 163, 18, 194, 242, 3, 186, 132, 172, 0, 0, 0, 2,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 69, 164, 14, 203, 10, 0, 0, 0, 32, 0, 0, 0, 154, 73, 54, 100, 6, 172,
	169, 82, 184, 139, 173, 245, 241, 233, 190, 108, 228, 150,
	129, 65, 3, 90, 96, 190, 80, 50, 115, 234, 101, 69, 107, 36, 64, 0, 0, 0, 38, 167, 193, 210,
	7, 30, 251, 149, 236, 15, 91, 233, 73, 174, 79, 86, 20, 133, 161, 167, 112,
	102, 42, 244, 246, 239, 78, 29, 104, 150, 190, 48, 230, 111, 129, 164, 66,
	29, 244, 75, 46, 150, 68, 242, 76, 26, 69, 205, 205, 122, 253, 219, 142,
	171, 28, 217, 139, 28, 133, 247, 59, 100, 161, 14, 1, 0, 0, 0, 12, 0, 0, 0, 8, 0, 12, 0, 8, 0, 4,
	0, 8, 0, 0, 0, 8, 0, 0, 0, 16, 0, 0, 0, 2, 0, 0, 0, 128, 150, 152, 0, 0, 0, 0, 0, 2, 0, 0, 0, 41, 207, 95, 217,
	65, 173, 37, 213, 8, 0, 8, 0, 0, 0, 4, 0, 8, 0, 0, 0, 4, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0}

type fakeSchemaAttribute struct {
	abstractSchemaAttribute
}

func TestSchemaReadInt32(t *testing.T) {
	attr := &fakeSchemaAttribute{}
	r := attr.readUint32(3, b)
	want := uint32(738204672)

	if !reflect.DeepEqual(r, want) {
		t.Errorf("Schema returned %d, want %d", r, want)
	}
}

func TestSchemaReadInt16(t *testing.T) {
	attr := &fakeSchemaAttribute{}
	r := attr.readUint16(3, b)
	want := uint32(7168)

	if !reflect.DeepEqual(r, want) {
		t.Errorf("Schema returned %d, want %d", r, want)
	}
}

func TestSchemaOffset(t *testing.T) {
	attr := &fakeSchemaAttribute{}
	r := attr.offset(32, 4, b)
	want := uint32(40)

	if !reflect.DeepEqual(r, want) {
		t.Errorf("Schema returned %d, want %d", r, want)
	}
}

func TestSchemaVectorLength(t *testing.T) {
	attr := &fakeSchemaAttribute{}
	r := attr.vectorLength(68, b)
	want := uint32(64)

	if !reflect.DeepEqual(r, want) {
		t.Errorf("Schema returned %d, want %d", r, want)
	}
}

func TestSchemaIndirect(t *testing.T) {
	attr := &fakeSchemaAttribute{}
	r := attr.indirect(68, b)
	want := uint32(168)

	if !reflect.DeepEqual(r, want) {
		t.Errorf("Schema returned %d, want %d", r, want)
	}
}

func TestSchemaFindParam(t *testing.T) {
	attr := &fakeSchemaAttribute{}
	r := attr.findParam(32, 14, b, IntSize)
	want := byte(52)

	if !reflect.DeepEqual(r[0], want) {
		t.Errorf("Schema returned %d, want %d", r[0], want)
	}
}

func TestSchemaFindVector(t *testing.T) {
	attr := &fakeSchemaAttribute{}
	r := attr.findVector(32, 18, b, ByteSize)
	want := []byte{144, 232, 254, 189, 103, 29, 212, 27, 238, 148, 236, 59,
		165, 131, 28, 182, 8, 163, 18, 194, 242, 3, 186, 132, 172}

	if !reflect.DeepEqual(r, want) {
		t.Errorf("Schema returned %d, want %d", r, want)
	}
}

func TestSchemaVectorStartPosition(t *testing.T) {
	attr := &fakeSchemaAttribute{}
	p := attr.findObjectStartPosition(32, 24, b)
	r := attr.findVector(p, 6, b, ByteSize)
	want := byte(0)

	if !reflect.DeepEqual(r[0], want) {
		t.Errorf("Schema returned %d, want %d", r[0], want)
	}
}

func TestSchemaStartPosition(t *testing.T) {
	attr := &fakeSchemaAttribute{}
	r := attr.findObjectStartPosition(32, 24, b)
	want := uint32(296)

	if !reflect.DeepEqual(r, want) {
		t.Errorf("Schema returned %d, want %d", r, want)
	}
}
