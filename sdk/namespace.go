package sdk

import (
	"bytes"
	"errors"
	"fmt"
	"golang.org/x/crypto/sha3"
	"golang.org/x/net/context"
	"math/big"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type NamespaceService service

func NewNamespaceService(httpClient *http.Client, conf *Config) *NamespaceService {
	ref := &NamespaceService{client: NewClient(httpClient, conf)}

	return ref
}

type uint64DTO [2]*big.Int

func NewUint64TO() *uint64DTO {
	return &uint64DTO{
		nil,
		nil,
	}
}

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
	list []*NamespaceId
}

func (ref *NamespaceIds) MarshalJSON() (buf []byte, err error) {
	buf = []byte(`{"namespaceIds": [`)
	for i, nsId := range ref.list {
		if i > 0 {
			buf = append(buf, ',')
		}
		buf = append(buf, []byte(`"`+nsId.fullName+`"`)...)
	}

	buf = append(buf, ']', '}')
	return
}

type NamespaceName struct {
	namespaceId *NamespaceId

	name string

	parentId *NamespaceId /* Optional NamespaceId my be nil */

}                              /* NamespaceName */
type NamespaceNameDTO struct { /*  */

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
	startHeight *uint64DTO
	endHeight   *uint64DTO
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
		NewNamespaceId(nsInfoDTO.Namespace.ParentId, ""),
		pubAcc,
		nsInfoDTO.Namespace.StartHeight,
		nsInfoDTO.Namespace.EndHeight,
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

type NamespaceMosaicMetaDTO struct {
	Active bool
	Index  int
	Id     string
} /* NamespaceMosaicMetaDTO */
type NamespaceDTO struct {
	Type int

	Depth int

	Level0 *uint64DTO /*DTO*/

	Level1 *uint64DTO /*DTO*/

	Level2 *uint64DTO /*DTO*/

	ParentId *uint64DTO /*DTO*/

	Owner string

	OwnerAddress string

	StartHeight *uint64DTO /*DTO*/

	EndHeight *uint64DTO /*DTO*/

} /* NamespaceDTO */
type NamespaceInfoDTO struct {
	Meta      NamespaceMosaicMetaDTO
	Namespace NamespaceDTO
}

func (ref *NamespaceInfoDTO) extractLevels() []*NamespaceId {

	levels := make([]*NamespaceId, 3)
	var err error

	nsId := NewNamespaceId(ref.Namespace.Level0, "")
	if err == nil {
		levels = append(levels, nsId)
	}

	nsId = NewNamespaceId(ref.Namespace.Level1, "")
	if err == nil {
		levels = append(levels, nsId)
	}

	nsId = NewNamespaceId(ref.Namespace.Level2, "")
	if err == nil {
		levels = append(levels, nsId)
	}
	return levels
}

type ListNamespaceInfo struct {
	list []*NamespaceInfo
}

const pathNamespace = "/namespace/"

func (ref *NamespaceService) GetNamespace(ctx context.Context, nsId string) (nsInfo *NamespaceInfo, resp *http.Response, err error) {

	var req *http.Request
	req, err = ref.client.NewRequest("GET", pathNamespace+nsId, nil)

	if err == nil {
		nsInfoDTO := &NamespaceInfoDTO{}
		resp, err = ref.client.Do(ctx, req, nsInfoDTO)

		if err == nil {
			nsInfo, err = NamespaceInfoFromDTO(nsInfoDTO)
			if err == nil {
				return nsInfo, resp, err
			}

		}
	}
	//	err occurent
	return nil, nil, err
}

const pathNamespacenames = "/namespace/names"

func (ref *NamespaceService) GetNamespaceNames(ctx context.Context, nsIds *NamespaceIds) (nsList []*NamespaceName, resp *http.Response, err error) {
	var req *http.Request
	req, err = ref.client.NewRequest("POST", pathNamespacenames, &nsIds)

	if err == nil {
		res := make([]*NamespaceNameDTO, 0)
		buf := bytes.Buffer{}
		resp, err = ref.client.Do(ctx, req, &buf)

		if err == nil {
			for _, val := range res {
				nsList = append(nsList, &NamespaceName{
					NewNamespaceId(val.namespaceId, ""),
					val.name,
					NewNamespaceId(val.parentId, "")})
			}
			return nsList, resp, err
		}
	}

	//	err occurent
	return nsList, nil, err
}

// GetNamespacesFromAccount get required params addresses, other skipped if value < 0
func (ref *NamespaceService) GetNamespacesFromAccount(ctx context.Context, address *Address, nsId string,
	pageSize int) (nsInfo ListNamespaceInfo, resp *http.Response, err error) {

	addresses := &Addresses{}
	addresses.AddAddress(address)

	return ref.GetNamespacesFromAccounts(ctx, addresses, nsId, pageSize)
}

const pathNamespacesFromAccount = "/account/namespaces"

type ns struct {
	addresses Addresses
}

// GetNamespacesFromAccounts get required params addresses, other skipped if value < 0
func (ref *NamespaceService) GetNamespacesFromAccounts(ctx context.Context, addresses *Addresses, nsId string,
	pageSize int) (nsList ListNamespaceInfo, resp *http.Response, err error) {

	url, comma := "", "?"

	if nsId > "" {
		url = comma + "id=" + nsId
		comma = "&"
	}

	if pageSize >= 0 {
		url += comma + "pageSize=" + strconv.Itoa(pageSize)
	}

	url = pathNamespacesFromAccount + url

	var req *http.Request

	req, err = ref.client.NewRequest("POST", url, &addresses)

	if err == nil {

		res := make([]*NamespaceInfoDTO, 0)
		resp, err = ref.client.Do(ctx, req, &res)

		for _, nsInfoDTO := range res {
			nsInfo, err := NamespaceInfoFromDTO(nsInfoDTO)
			if err != nil {
				return nsList, resp, err
			}
			nsList.list = append(nsList.list, nsInfo)
		}
		return nsList, resp, err

		if err == nil {
			if len(nsList.list) > 0 {
				return nsList, resp, err
			}
			err = errors.New("not result returns")
		}
	}

	//	err occurent
	return nsList, resp, err
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

	namespaceId := NewUint64TO()
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
		low := &big.Int{}
		low = low.SetBytes(t[:4])
		high := &big.Int{}
		high = high.SetBytes(t[4:8])
		return &uint64DTO{low, high}, nil
	}

	return nil, err
}
