// Copyright 2017 Author: Ruslan Bikchentaev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sdk

import (
	"errors"
	"fmt"
	"github.com/json-iterator/go"
	"golang.org/x/crypto/sha3"
	"net/http"
	"regexp"
	"strings"
	"unsafe"
	"math/big"
)

type NamespaceService service

func NewNamespaceService(httpClient *http.Client, conf *Config) *NamespaceService {
	ref := &NamespaceService{client: NewClient(httpClient, conf)}

	return ref
}

type NamespaceId struct {
	id       *big.Int
	fullName string
} /* NamespaceId */
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
type NamespaceNameDTO struct {
	namespaceId *uint64DTO
	name        string
	parentId    *uint64DTO
} /* NamespaceNameDTO */
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
	startHeight *big.Int
	endHeight   *big.Int
} /* NamespaceInfo */

func NamespaceInfoFromDTO(nsInfoDTO *NamespaceInfoDTO) (*NamespaceInfo, error) {
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
		NewNamespaceId(nsInfoDTO.Namespace.ParentId.toStruct(), ""),
		pubAcc,
		nsInfoDTO.Namespace.StartHeight.toStruct(),
		nsInfoDTO.Namespace.EndHeight.toStruct(),
	}, nil
}

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

func generateNamespaceId(namespaceName string) (*big.Int, error) {

	list, err := generateNamespacePath(namespaceName)
	if err != nil {
		return nil, err
	}

	return list[len(list)-1], nil
}

var regValidNamespace = regexp.MustCompile(`^[a-z0-9][a-z0-9-_]*$`)

func generateNamespacePath(name string) ([]*big.Int, error) {

	parts := strings.Split(name, ".")
	path := make([]*big.Int, 0)
	if len(parts) == 0 {
		return nil, errors.New("invalid Namespace name")
	}

	if len(parts) > 3 {
		return nil, errors.New("too many parts")
	}

	namespaceId := big.NewInt(0)
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

func generateId(name string, parentId *big.Int) (*big.Int, error) {
	var int big.Int
	result := sha3.New256()
	_, err := result.Write(parentId.Bytes())
	if err == nil {
		t := result.Sum([]byte(name))
		return int.SetBytes(append(t[:4], t[4:8]...)), nil
	}
	return nil, err
}