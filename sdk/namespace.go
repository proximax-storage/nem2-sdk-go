package sdk

import (
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"strconv"
	"strings"
)

type NamespaceService service

func NewNamespaceService(httpClient *http.Client, conf *Config) *NamespaceService {
	ref := &NamespaceService{client: NewClient(httpClient, conf)}

	return ref
}

type uint64DTO [2]uint64
type NamespaceId struct {
	id       uint64DTO
	fullName string
} /* NamespaceId */
func NewNamespaceId(id uint64DTO, namespaceName string) *NamespaceId {

	if namespaceName == "" {
		return &NamespaceId{id, ""}
	}

	return &NamespaceId{generateNamespaceId(namespaceName), namespaceName}
}

type NamespaceType int

const (
	RootNamespace NamespaceType = iota
	SubNamespace
) /* NamespaceType */
// NamespaceInfo contains the state information of a Namespace.
type NamespaceInfo struct {
	active bool

	index int

	metaId string

	typeSpace NamespaceType

	depth int

	levels []*NamespaceId

	parentId *NamespaceId

	owner PublicAccount

	startHeight uint64DTO

	endHeight uint64DTO
} /* NamespaceInfo */
type NamespaceMosaicMetaDTO struct {
	Active bool
	Index  int
	Id     string
} /* NamespaceMosaicMetaDTO */
type NamespaceDTO struct {
	Type int

	Depth int

	Level0 uint64DTO /*DTO*/

	Level1 uint64DTO /*DTO*/

	Level2 uint64DTO /*DTO*/

	ParentId uint64DTO /*DTO*/

	Owner string

	OwnerAddress string

	StartHeight uint64DTO /*DTO*/

	EndHeight uint64DTO /*DTO*/

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

type ListNamespaceInfo []*NamespaceInfo

const pathNamespace = "/namespace/"

func (ref *NamespaceService) GetNamespace(ctx context.Context, nsId string) (nsInfo *NamespaceInfo, resp *http.Response, err error) {
	var req *http.Request

	req, err = ref.client.NewRequest("GET", pathNamespace+nsId, nil)

	if err == nil {
		nsInfoDTO := &NamespaceInfoDTO{}
		resp, err = ref.client.Do(ctx, req, nsInfoDTO)

		fmt.Printf("%#v", nsInfoDTO)
		nsInfo = &NamespaceInfo{nsInfoDTO.Meta.Active,
			nsInfoDTO.Meta.Index,
			nsInfoDTO.Meta.Id,
			NamespaceType(nsInfoDTO.Namespace.Type),
			nsInfoDTO.Namespace.Depth,
			nsInfoDTO.extractLevels(),
			NewNamespaceId(nsInfoDTO.Namespace.ParentId, ""),
			PublicAccount{createFromPublicKey(nsInfoDTO.Namespace.Owner, NetworkType(nsInfoDTO.Namespace.Type)), nsInfoDTO.Namespace.Owner},
			nsInfoDTO.Namespace.StartHeight,
			nsInfoDTO.Namespace.EndHeight}

		if err == nil {
			return nsInfo, resp, err
		}
	}
	//	err occurent
	return nil, nil, err
}
func (ref *NamespaceService) GetNamespaces(ctx context.Context, nsIds []string) (nsInfo ListNamespaceInfo, resp *http.Response, err error) {
	var req *http.Request

	req, err = ref.client.NewRequest("GET", pathNamespace+strings.Join(nsIds, ","), nil)

	if err == nil {
		nsInfo = make(ListNamespaceInfo, 0)
		resp, err = ref.client.Do(ctx, req, nsInfo)

		if err == nil {
			//if nsInfo[0].index == nsIds[0] {
			return nsInfo, resp, err
			//}
			err = errors.New("nod valid index returns")
		}
	}
	//	err occurent
	return nil, nil, err
}

// GetNamespacesFromAccount get required params addresses, other skipped if value < 0
func (ref *NamespaceService) GetNamespacesFromAccount(ctx context.Context, address *Address, nsId int,
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
func (ref *NamespaceService) GetNamespacesFromAccounts(ctx context.Context, addresses *Addresses, nsId int,
	pageSize int) (nsInfo ListNamespaceInfo, resp *http.Response, err error) {

	//body, err := json.Marshal(addresses)
	body := []byte(`{"addresses": ["SDRDGFTDLLCB67D4HPGIMIHPNSRYRJRT7DOBGWZY","SBCPGZ3S2SCC3YHBBTYDCUZV4ZZEPHM2KGCP4QXX"]}`)

	url, comma := "", "?"

	if nsId >= 0 {
		url = comma + strconv.Itoa(nsId)
		comma = "&"
	}

	if pageSize >= 0 {
		url = comma + strconv.Itoa(pageSize)
	}

	url = pathNamespacesFromAccount + url

	var req *http.Request
	req, err = ref.client.NewRequest("POST", pathNamespacesFromAccount, &body)

	if err == nil {

		var ns ns
		resp, err = ref.client.Do(ctx, req, &ns)
		return nsInfo, resp, err

		if err == nil {
			if len(nsInfo) > 0 {
				return nsInfo, resp, err
			}
			err = errors.New("nod result returns")
		}
	}
	//	err occurent
	return nsInfo, resp, err
}
func generateNamespaceId(namespaceName string) uint64DTO { /* public static  */

	return generateNamespacePath(namespaceName)
}
func generateId(name string, parentId int64) int64 { /* public static  */

	//var parentIdBytes byte[8]
	//bytes.Buffer.wrap(parentIdBytes).put(parentId.toByteArray()); // GO
	//ArrayUtils.reverse(parentIdBytes)
	//[]byte bytes = name.getBytes()
	//[]byte result = Hashes.sha3_256(parentIdBytes, bytes)
	//[]byte low = Arrays.copyOfRange(result, 0, 4)
	//[]byte high = Arrays.copyOfRange(result, 4, 8)
	//ArrayUtils.reverse(low)
	//ArrayUtils.reverse(high)
	//[]byte last = ArrayUtils.addAll(high, low)
	return -1 //NewBigInteger(last)
}
func generateNamespacePath(name string) uint64DTO { /* public static  */

	//[]string parts = name.split(Pattern.quote("."))
	//[]int64 path = new Array[]int64()
	//if (parts.length == 0) {
	//	panic(IllegalIdentifierException{"invalid Namespace name"})
	//}
	//else if (parts.length > 3) {
	//	panic(IllegalIdentifierException{"too many parts"})
	//}
	//
	//int64 namespaceId = int64.valueOf(0)
	//for (int i = 0; i < parts.length; i++) {
	//	if (!parts[i].matches("^[a-z0-9][a-z0-9-_]*$")) {
	//		panic(IllegalIdentifierException{"invalid Namespace name"})
	//	}
	//
	//	namespaceId = generateId(parts[i], namespaceId)
	//	path.add(namespaceId)
	//}

	return uint64DTO{0, 1}
}
