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

// NamespaceId id structure describes namespace id
type NamespaceId struct {
	Id       *uint64DTO
	FullName string
}

// NewNamespaceId create NamespaceId from namespace string name if he present
// other create NamespaceId from biginteger id
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
		buf = append(buf, []byte(`"`+nsId.FullName+`"`)...)
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

// NamespaceName name info structure describes basic information of a namespace and name.
type NamespaceName struct {
	NamespaceId *NamespaceId
	Name        string
	ParentId    *NamespaceId /* Optional NamespaceId my be nil */
}

// NamespaceType containing namespace supply type.
type NamespaceType int

const (
	RootNamespace NamespaceType = iota
	SubNamespace
)

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
	StartHeight *uint64DTO
	EndHeight   *uint64DTO
}

const tplNamespaceInfo = `"active": %v,
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
	return fmt.Sprintf(tplNamespaceInfo,
		ref.Active,
		ref.Index,
		ref.MetaId,
		ref.TypeSpace,
		ref.Depth,
		ref.Levels,
		ref.ParentId,
		ref.Owner,
		ref.Owner.Address.Address,
		ref.StartHeight,
		ref.EndHeight,
	)
}

// ListNamespaceInfo is a list NamespaceInfo
type ListNamespaceInfo struct {
	list []*NamespaceInfo
}

// generateNamespaceId create NamespaceId from namespace string name (ex: nem or domain.subdom.subdome)
func generateNamespaceId(namespaceName string) (*uint64DTO, error) {

	list, err := generateNamespacePath(namespaceName)
	if err != nil {
		return nil, err
	}

	return list[len(list)-1], nil
}

// regValidNamespace check namespace on valid symbols
var regValidNamespace = regexp.MustCompile(`^[a-z0-9][a-z0-9-_]*$`)

// generateNamespacePath create list NamespaceId from string
func generateNamespacePath(name string) ([]*uint64DTO, error) {

	parts := strings.Split(name, ".")
	path := make([]*uint64DTO, 0)
	if len(parts) == 0 {
		return nil, errors.New("invalid Namespace Name")
	}

	if len(parts) > 3 {
		return nil, errors.New("too many parts")
	}

	namespaceId := NewRootUint64DTO()
	for i, part := range parts {
		if !regValidNamespace.MatchString(part) {
			return nil, errors.New("invalid Namespace Name")
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
