// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package sdk

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/json-iterator/go"
	"github.com/proximax-storage/nem2-sdk-go/utils"
	"golang.org/x/crypto/sha3"
	"math/big"
	"strings"
	"unsafe"
)

// 	NamespaceId
type NamespaceId struct {
	Id       *big.Int
	FullName string
}

func (n *NamespaceId) String() string {
	return utils.StructToString(
		"NamespaceId",
		utils.NewField("Id", utils.StringPattern, n.Id),
		utils.NewField("FullName", utils.StringPattern, n.FullName),
	)
}

func (n *NamespaceId) toHexString() string {
	return BigIntegerToHex(n.Id)
}

//NewNamespaceId generate new NamespaceId from bigInt
func NewNamespaceId(id *big.Int) (*NamespaceId, error) {

	if id == nil {
		return nil, errNilIdNamespace
	}
	return &NamespaceId{id, ""}, nil
}

//NewNamespaceIdFromName generate Id from namespaceName
func NewNamespaceIdFromName(namespaceName string) (*NamespaceId, error) {

	id, err := generateNamespaceId(namespaceName)
	if err != nil {
		return nil, err
	}
	return &NamespaceId{id, namespaceName}, nil
}

// Equals compares namespaceIds for equality
func (ref *NamespaceId) Equals(nsId NamespaceId) bool {
	return (ref.Id == nsId.Id) && (ref.FullName == nsId.FullName)
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
type NamespaceName struct {
	NamespaceId *NamespaceId
	Name        string
	ParentId    *NamespaceId /* Optional NamespaceId my be nil */
}

func (n *NamespaceName) String() string {
	return fmt.Sprintf(
		`
			"NamespaceId": %s,
			"Name": %s,
			"ParentId": %s,
		`,
		n.NamespaceId,
		n.Name,
		n.ParentId,
	)
}

// NamespaceInfo contains the state information of a Namespace.
type NamespaceInfo struct {
	Active      bool
	Index       int
	MetaId      string
	TypeSpace   NamespaceType
	Depth       int
	Levels      []*NamespaceId
	ParentId    *NamespaceId
	Owner       *PublicAccount
	StartHeight *big.Int
	EndHeight   *big.Int
}

func (ref *NamespaceInfo) String() string {
	return utils.StructToString(
		"NamespaceInfo",
		utils.NewField("Active", utils.BooleanPattern, ref.Active),
		utils.NewField("Index", utils.IntPattern, ref.Index),
		utils.NewField("MetaId", utils.StringPattern, ref.MetaId),
		utils.NewField("TypeSpace", utils.IntPattern, ref.TypeSpace),
		utils.NewField("Depth", utils.IntPattern, ref.Depth),
		utils.NewField("Levels", utils.StringPattern, ref.Levels),
		utils.NewField("ParentId", utils.StringPattern, ref.ParentId),
		utils.NewField("Owner", utils.StringPattern, ref.Owner),
		utils.NewField("StartHeight", utils.StringPattern, ref.StartHeight),
		utils.NewField("EndHeight", utils.StringPattern, ref.EndHeight),
	)
}

// ListNamespaceInfo is a list NamespaceInfo
type ListNamespaceInfo struct {
	List []*NamespaceInfo
}

// generateNamespaceId create NamespaceId from namespace string name (ex: nem or domain.subdom.subdome)
func generateNamespaceId(namespaceName string) (*big.Int, error) {
	list, err := GenerateNamespacePath(namespaceName)

	if err != nil {
		return nil, err
	}

	return list[len(list)-1], nil
}

// GenerateNamespacePath create list NamespaceId from string
func GenerateNamespacePath(name string) ([]*big.Int, error) {
	parts := strings.Split(name, ".")
	path := make([]*big.Int, 0)
	if len(parts) == 0 {
		return nil, errors.New("invalid Namespace Name")
	}

	if len(parts) > 3 {
		return nil, errNamespaceToManyPart
	}

	namespaceId := big.NewInt(0)
	for _, part := range parts {
		if !regValidNamespace.MatchString(part) {
			return nil, errors.New("invalid Namespace name")
		}

		var err error
		namespaceId, err = generateId(part, namespaceId)
		if err != nil {
			return nil, err
		}
		path = append(path, namespaceId)
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
	_, err := result.Write(b)
	if err != nil {
		return nil, err
	}
	_, err = result.Write([]byte(name))
	if err != nil {
		return nil, err
	}
	t := result.Sum(nil)

	return uint64DTO{binary.LittleEndian.Uint32(t[0:4]), binary.LittleEndian.Uint32(t[4:8])}.toBigInt(), nil
}
