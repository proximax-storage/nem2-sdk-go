// Copyright 2017 Author: Ruslan Bikchentaev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sdk

import (
	"errors"
	"fmt"
	"github.com/json-iterator/go"
	"golang.org/x/crypto/sha3"
	"regexp"
	"strings"
	"unsafe"
)

type NamespaceId struct {
	id       *uint64DTO
	fullName string
} /* NamespaceId */
func NewNamespaceId(id *uint64DTO, namespaceName string) *NamespaceId {

	if namespaceName == "" {
		return &NamespaceId{id, ""}
	}

	id, err := generateNamespaceId(namespaceName)
	if err != nil {
		return nil
	}
	return &NamespaceId{id, namespaceName}
}

type NamespaceIds struct {
	List []*NamespaceId
}

func (ref *NamespaceIds) MarshalJSON() (buf []byte, err error) {
	buf = []byte(`{"namespaceIds": [`)
	for i, nsId := range ref.List {
		if i > 0 {
			buf = append(buf, ',')
		}
		buf = append(buf, []byte(`"`+nsId.fullName+`"`)...)
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
		stream.Write(buf)
	}

}

type NamespaceName struct {
	namespaceId *NamespaceId
	name        string
	parentId    *NamespaceId /* Optional NamespaceId my be nil */
} /* NamespaceName */

type NamespaceType int

const (
	RootNamespace NamespaceType = iota
	SubNamespace
) /* NamespaceType */
// NamespaceInfo contains the state information of a Namespace.
type NamespaceInfo struct {
	active      bool
	index       int
	metaId      string
	typeSpace   NamespaceType
	depth       int
	levels      []*NamespaceId
	parentId    *NamespaceId
	owner       *PublicAccount
	startHeight *uint64DTO
	endHeight   *uint64DTO
} /* NamespaceInfo */

const templNamespaceInfo = `"active": %v,
    "index": %d,
    "id": "%s",
	"type": %d,
    "depth": %d,
    "levels": [
      %v
    ],
    "parentId": [
      %v
    ],
    "owner": "%v",
    "ownerAddress": "%s",
    "startHeight": [
      %v
    ],
    "endHeight": [
      %v
    ]
  }
`

func (ref *NamespaceInfo) String() string {
	return fmt.Sprintf(templNamespaceInfo,
		ref.active,
		ref.index,
		ref.metaId,
		ref.typeSpace,
		ref.depth,
		ref.levels,
		ref.parentId,
		ref.owner,
		ref.owner.Address.Address,
		ref.startHeight,
		ref.endHeight,
	)
}

type ListNamespaceInfo struct {
	list []*NamespaceInfo
}

func generateNamespaceId(namespaceName string) (*uint64DTO, error) {

	list, err := generateNamespacePath(namespaceName)
	if err != nil {
		return nil, err
	}

	return list[len(list)-1], nil
}

var regValidNamespace = regexp.MustCompile(`^[a-z0-9][a-z0-9-_]*$`)

func generateNamespacePath(name string) ([]*uint64DTO, error) {

	parts := strings.Split(name, ".")
	path := make([]*uint64DTO, 0)
	if len(parts) == 0 {
		return nil, errors.New("invalid Namespace name")
	}

	if len(parts) > 3 {
		return nil, errors.New("too many parts")
	}

	namespaceId := NewRootUint64DTO()
	for i, part := range parts {
		if !regValidNamespace.MatchString(part) {
			return nil, errors.New("invalid Namespace name")
		}

		var err error
		namespaceId, err = generateId(parts[i], namespaceId)
		if err != nil {
			return nil, err
		}
		path = append(path, namespaceId)
	}

	return path, nil
}
func generateId(name string, parentId *uint64DTO) (*uint64DTO, error) {

	result := sha3.New256()
	_, err := result.Write(append(parentId[1].Bytes(), parentId[0].Bytes()...))
	if err == nil {
		t := result.Sum([]byte(name))
		return NewUint64DTO(t[:4], t[4:8]), nil
	}

	return nil, err
}
