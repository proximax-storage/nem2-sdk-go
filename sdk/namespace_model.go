// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package sdk

import (
	"encoding/binary"
	"errors"
	"github.com/json-iterator/go"
	"github.com/proximax-storage/nem2-sdk-go/utils"
	"github.com/proximax-storage/proximax-utils-go/str"
	"golang.org/x/crypto/sha3"
	"math/big"
	"strings"
	"unsafe"
)

type NamespaceId big.Int

/*type NamespaceId struct {
	Id       *big.Int
	FullName string
}*/

func bigIntToNamespaceId(bigInt *big.Int) *NamespaceId {
	if bigInt == nil {
		return nil
	}

	nsId := NamespaceId(*bigInt)

	return &nsId
}

func namespaceIdToBigInt(nsId *NamespaceId) *big.Int {
	if nsId == nil {
		return nil
	}

	return (*big.Int)(nsId)
}

func (n *NamespaceId) toHexString() string {
	return BigIntegerToHex(namespaceIdToBigInt(n))
}

//NewNamespaceId generate new NamespaceId from bigInt
func NewNamespaceId(id *big.Int) (*NamespaceId, error) {
	if id == nil {
		return nil, ErrNilIdNamespace
	}

	return &NamespaceId{Id: id}, nil
}

//NewNamespaceIdFromName generate Id from namespaceName
func NewNamespaceIdFromName(namespaceName string) (*NamespaceId, error) {
	id, err := generateNamespaceId(namespaceName)
	if err != nil {
		return nil, err
	}

	return bigIntToNamespaceId(id), nil
}

// NamespaceIds is a list of NamespaceId
type NamespaceIds struct {
	List []*NamespaceId
}

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

func (ref *NamespaceIds) IsEmpty(ptr unsafe.Pointer) bool {
	return len((*NamespaceIds)(ptr).List) == 0
}

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

func (ref *NamespaceIds) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	buf, err := (*NamespaceIds)(ptr).MarshalJSON()
	if err == nil {
		_, err = stream.Write(buf)
		//	todo: log error in future
	}

}

// NamespaceName name info structure describes basic information of a namespace and name.

/*
type NamespaceName struct {
	NamespaceId *NamespaceId
	Name        string
	ParentId    *NamespaceId /* Optional NamespaceId my be nil
}

func (n *NamespaceName) String() string {
	return str.StructToString(
		"NamespaceName",
		str.NewField("NamespaceId", str.StringPattern, n.NamespaceId),
		str.NewField("Name", str.StringPattern, n.Name),
		str.NewField("ParentId", str.StringPattern, n.ParentId),
	)
}*/

// NamespaceInfo contains the state information of a Namespace.
type NamespaceInfo struct {
	NamespaceId   *big.Int
	Active        bool
	Index         int
	MetaId        string
	TypeSpace     NamespaceType
	Depth         int
	MosaicIds     []*big.Int
	SubNamespaces []*NamespaceId
	ParentId      NamespaceId
	Owner         *PublicAccount
	StartHeight   *big.Int
	EndHeight     *big.Int
}

func (ref *NamespaceInfo) String() string {
	return str.StructToString(
		"NamespaceInfo",
		str.NewField("Active", str.BooleanPattern, ref.Active),
		str.NewField("Index", str.IntPattern, ref.Index),
		str.NewField("MetaId", str.StringPattern, ref.MetaId),
		str.NewField("TypeSpace", str.IntPattern, ref.TypeSpace),
		str.NewField("Depth", str.IntPattern, ref.Depth),
		//str.NewField("Levels", str.StringPattern, ref.Levels),
		str.NewField("ParentId", str.StringPattern, ref.ParentId),
		str.NewField("Owner", str.StringPattern, ref.Owner),
		str.NewField("StartHeight", str.StringPattern, ref.StartHeight),
		str.NewField("EndHeight", str.StringPattern, ref.EndHeight),
	)
}

// ListNamespaceInfo is a list NamespaceInfo
type ListNamespaceInfo struct {
	List []*NamespaceInfo
}

// generateNamespaceId create NamespaceId from namespace string name (ex: nem or domain.subdom.subdome)
func generateNamespaceId(namespaceName string) (*big.Int, error) {
	if list, err := GenerateNamespacePath(namespaceName); err != nil {
		return nil, err
	} else {
		return list[len(list)-1], nil
	}
}

// GenerateNamespacePath create list NamespaceId from string
func GenerateNamespacePath(name string) ([]*big.Int, error) {
	parts := strings.Split(name, ".")

	if len(parts) == 0 {
		return nil, errors.New("invalid Namespace Name")
	}

	if len(parts) > 3 {
		return nil, ErrNamespaceToManyPart
	}

	namespaceId := big.NewInt(0)
	path := make([]*big.Int, 0)

	for _, part := range parts {
		if !regValidNamespace.MatchString(part) {
			return nil, errors.New("invalid Namespace name")
		}

		if namespaceId, err := generateId(part, namespaceId); err != nil {
			return nil, err
		} else {
			path = append(path, namespaceId)
		}
	}

	return path, nil
}

func generateId(name string, parentId *big.Int) (*big.Int, error) {
	b := make([]byte, 8)

	if parentId.Int64() != 0 {
		b = parentId.Bytes()
	}

	utils.ReverseByteArray(b)

	result := sha3.New256()

	if _, err := result.Write(b); err != nil {
		return nil, err
	}

	if _, err := result.Write([]byte(name)); err != nil {
		return nil, err
	}

	t := result.Sum(nil)

	return uint64DTO{binary.LittleEndian.Uint32(t[0:4]), binary.LittleEndian.Uint32(t[4:8])}.toBigInt(), nil
}
