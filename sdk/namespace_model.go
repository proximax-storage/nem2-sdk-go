package sdk

import (
	"errors"
	"fmt"
	"github.com/json-iterator/go"
	"github.com/proximax-storage/nem2-sdk-go/utils"
	"golang.org/x/crypto/sha3"
	"math/big"
	"regexp"
	"strings"
	"unsafe"
)

type NamespaceId struct {
	Id       *big.Int
	FullName string
}

type namespaceIdDTO uint64DTO

func (dto *namespaceIdDTO) toStruct() *NamespaceId {
	return NewNamespaceId(uint64DTO(*dto).toBigInt(), "")
}

/* NamespaceId */
func NewNamespaceId(id *big.Int, namespaceName string) *NamespaceId {

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
} /* NamespaceName */
type NamespaceNameDTO struct {
	namespaceId uint64DTO
	name        string
	parentId    uint64DTO
} /* NamespaceNameDTO */
type NamespaceType uint8

const (
	Root NamespaceType = iota
	Sub
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
	StartHeight *big.Int
	EndHeight   *big.Int
} /* NamespaceInfo */
func NamespaceInfoFromDTO(nsInfoDTO *namespaceInfoDTO) (*NamespaceInfo, error) {
	pubAcc, err := NewPublicAccount(nsInfoDTO.Namespace.Owner, NetworkType(nsInfoDTO.Namespace.Type))
	if err != nil {
		return nil, err
	}

	return &NamespaceInfo{
		nsInfoDTO.Meta.Active,
		nsInfoDTO.Meta.Index,
		nsInfoDTO.Meta.Id,
		NamespaceType(nsInfoDTO.Namespace.Type),
		nsInfoDTO.Namespace.Depth,
		nsInfoDTO.extractLevels(),
		NewNamespaceId(nsInfoDTO.Namespace.ParentId.toBigInt(), ""),
		pubAcc,
		nsInfoDTO.Namespace.StartHeight.toBigInt(),
		nsInfoDTO.Namespace.EndHeight.toBigInt(),
	}, nil
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
func generateNamespaceId(namespaceName string) (*big.Int, error) {

	list, err := generateNamespacePath(namespaceName)
	if err != nil {
		return nil, err
	}

	return new(big.Int).SetBytes(list[len(list)-1]), nil
}

// regValidNamespace check namespace on valid symbols
var regValidNamespace = regexp.MustCompile(`^[a-z0-9][a-z0-9-_]*$`)

// generateNamespacePath create list NamespaceId from string
func generateNamespacePath(name string) ([][]byte, error) {

	parts := strings.Split(name, ".")
	path := make([][]byte, 0)
	if len(parts) == 0 {
		return nil, errors.New("invalid Namespace Name")
	}

	if len(parts) > 3 {
		return nil, errors.New("too many parts")
	}

	emptyNamespaceId := make([]byte, 8)
	for i, part := range parts {
		if !regValidNamespace.MatchString(part) {
			return nil, errors.New("invalid Namespace name")
		}

		var err error
		namespaceId, err := generateId(parts[i], emptyNamespaceId)
		if err != nil {
			return nil, err
		}
		path = append(path, namespaceId)
	}

	return path, nil
}

func generateId(name string, parentId []byte) ([]byte, error) {
	utils.ReverseByteArray(parentId)
	result := sha3.New256()
	_, err := result.Write(parentId)
	if err != nil {
		return nil, err
	}
	_, err = result.Write([]byte(name))
	if err != nil {
		return nil, err
	}
	t := result.Sum(nil)
	l := t[:4]
	h := t[4:8]
	utils.ReverseByteArray(l)
	utils.ReverseByteArray(h)
	return append(h, l...), nil
}
