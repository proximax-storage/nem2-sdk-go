// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package sdk

import (
	"github.com/json-iterator/go"
	"github.com/proximax-storage/proximax-utils-go/str"
	"math/big"
	"strings"
	"unsafe"
)

// NamespaceId encapsulates namespace id operations
type NamespaceId big.Int

//NewNamespaceId generate new NamespaceId from bigInt
func NewNamespaceId(id *big.Int) (*NamespaceId, error) {
	if id == nil {
		return nil, ErrNilNamespaceId
	}

	return bigIntToNamespaceId(id), nil
}

//NewNamespaceIdFromName generate Id from namespaceName
func NewNamespaceIdFromName(namespaceName string) (*NamespaceId, error) {
	id, err := generateNamespaceId(namespaceName)
	if err != nil {
		return nil, err
	}

	return bigIntToNamespaceId(id), nil
}

func (n *NamespaceId) String() string {
	return (*big.Int)(n).String()
}

func (n *NamespaceId) toHexString() string {
	return BigIntegerToHex(namespaceIdToBigInt(n))
}

// NamespaceIds is a list of NamespaceId
type NamespaceIds struct {
	List []*NamespaceId
}

// MarshalJSON serialize namespaceId list
func (ref *NamespaceIds) MarshalJSON() (buf []byte, err error) {
	buf = []byte(`{"namespaceIds": [`)

	for i, nsId := range ref.List {
		if i > 0 {
			buf = append(buf, ',')
		}

		buf = append(buf, []byte(`"`+nsId.toHexString()+`"`)...)
	}

	buf = append(buf, ']', '}')

	return
}

// IsEmpty returns true is list not empty for JSON operation
func (ref *NamespaceIds) IsEmpty(ptr unsafe.Pointer) bool {
	return len((*NamespaceIds)(ptr).List) == 0
}

// Decode get namespace list from JSON stream
func (ref *NamespaceIds) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	if (*NamespaceIds)(ptr) == nil {
		ptr = (unsafe.Pointer)(&NamespaceIds{})
	}

	if iter.ReadNil() {
		*((*unsafe.Pointer)(ptr)) = nil
	} else {
		if iter.WhatIsNext() == jsoniter.ArrayValue {
			iter.Skip()
			newIter := iter.Pool().BorrowIterator([]byte("{}"))
			defer iter.Pool().ReturnIterator(newIter)
			v := newIter.Read()
			list := make([]*NamespaceId, 0)
			for _, val := range v.([]*NamespaceId) {
				list = append(list, val)
			}
			(*NamespaceIds)(ptr).List = list
		}
	}
}

// Encode serialize NamespaceIds to JSON stream
func (ref *NamespaceIds) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	//	todo: log error in future
	buf, err := (*NamespaceIds)(ptr).MarshalJSON()
	if err == nil {
		_, err = stream.Write(buf)
	}

}

// NamespaceInfo contains the state information of a Namespace.
type NamespaceInfo struct {
	NamespaceId *NamespaceId
	FullName    string
	Active      bool
	Index       int
	MetaId      string
	TypeSpace   NamespaceType
	Depth       int
	Levels      []*NamespaceId
	Parent      *NamespaceInfo
	Owner       *PublicAccount
	StartHeight *big.Int
	EndHeight   *big.Int
}

func (ref *NamespaceInfo) String() string {
	return str.StructToString(
		"NamespaceInfo",
		str.NewField("NamespaceId", str.StringPattern, ref.NamespaceId),
		str.NewField("FullName", str.StringPattern, ref.FullName),
		str.NewField("Active", str.BooleanPattern, ref.Active),
		str.NewField("Index", str.IntPattern, ref.Index),
		str.NewField("MetaId", str.StringPattern, ref.MetaId),
		str.NewField("TypeSpace", str.IntPattern, ref.TypeSpace),
		str.NewField("Depth", str.IntPattern, ref.Depth),
		str.NewField("Levels", str.StringPattern, ref.Levels),
		str.NewField("Parent", str.StringPattern, ref.Parent),
		str.NewField("Owner", str.StringPattern, ref.Owner),
		str.NewField("StartHeight", str.StringPattern, ref.StartHeight),
		str.NewField("EndHeight", str.StringPattern, ref.EndHeight),
	)
}

// generateNamespaceId create NamespaceId from namespace string name (ex: nem or domain.subdom.subdome)
func generateNamespaceId(namespaceName string) (*big.Int, error) {
	list, err := GenerateNamespacePath(namespaceName)
	if err != nil {
		return nil, err
	}

	return list[len(list)-1], nil
}

// NamespaceName name info structure describes basic information of a namespace and name.
type NamespaceName struct {
	NamespaceId *NamespaceId
	Name        string
	ParentId    *NamespaceId /* Optional NamespaceId my be nil */
}

func (n *NamespaceName) String() string {
	return str.StructToString(
		"NamespaceName",
		str.NewField("NamespaceId", str.StringPattern, n.NamespaceId),
		str.NewField("Name", str.StringPattern, n.Name),
		str.NewField("ParentId", str.StringPattern, n.ParentId),
	)
}

// GenerateNamespacePath create list NamespaceId from string
func GenerateNamespacePath(name string) ([]*big.Int, error) {
	parts := strings.Split(name, ".")

	if len(parts) == 0 {
		return nil, ErrInvalidNamespaceName
	}

	if len(parts) > 3 {
		return nil, ErrNamespaceTooManyPart
	}

	var (
		namespaceId = big.NewInt(0)
		path        = make([]*big.Int, 0)
		err         error
	)

	for _, part := range parts {
		if !regValidNamespace.MatchString(part) {
			return nil, ErrInvalidNamespaceName
		}

		if namespaceId, err = generateId(part, (*big.Int)(namespaceId)); err != nil {
			return nil, err
		}

		path = append(path, namespaceId)
	}

	return path, nil
}
