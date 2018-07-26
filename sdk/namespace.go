package sdk

import (
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"strconv"
)

type NamespaceService service

func NewNamespaceService(httpClient *http.Client, conf *Config) *NamespaceService {
	ref := &NamespaceService{client: NewClient(httpClient, conf)}

	return ref
}

type NamespaceId struct {
	id       int64
	fullName string
} /* NamespaceId */
func NewNamespaceId(id int64, namespaceName string) (ref *NamespaceId, err error) {
	if namespaceName == "" {
		return nil, errors.New("nameSpace is empty")
	}
	if id < 0 {
		id = generateNamespaceId(namespaceName)
	}
	ref = &NamespaceId{id, namespaceName}
	return
}

type NamespaceType int

const (
	RootNamespace NamespaceType = iota
	SubNamespace
) /* NamespaceType */
// NamespaceInfo contains the state information of a namespace.
type NamespaceInfo struct {
	active bool

	id int

	metaId string

	typeSpace NamespaceType

	depth int

	levels []NamespaceId

	parentId NamespaceId

	owner PublicAccount

	startHeight int64

	endHeight int64
} /* NamespaceInfo */

type ListNamespaceInfo []*NamespaceInfo

func (ref *NamespaceService) GetNamespace(ctx context.Context, nsId int) (nsInfo *NamespaceInfo, resp *http.Response, err error) {
	var req *http.Request

	req, err = ref.client.NewRequest("GET", fmt.Sprintf("/namespace/%d", nsId), nil)

	if err == nil {
		nsInfo = &NamespaceInfo{}
		resp, err = ref.client.Do(ctx, req, nsInfo)

		if err == nil {
			if nsInfo.id == nsId {
				return nsInfo, resp, err
			}
			err = errors.New("nod valid id returns")
		}
	}
	//	err occurent
	return nil, nil, err
}
func (ref *NamespaceService) GetNamespaces(ctx context.Context, nsIds []int) (nsInfo ListNamespaceInfo, resp *http.Response, err error) {
	var req *http.Request

	req, err = ref.client.NewRequest("GET", fmt.Sprintf("/namespace/%d", nsIds), nil)

	if err == nil {
		nsInfo = make(ListNamespaceInfo, 0)
		resp, err = ref.client.Do(ctx, req, nsInfo)

		if err == nil {
			if nsInfo[0].id == nsIds[0] {
				return nsInfo, resp, err
			}
			err = errors.New("nod valid id returns")
		}
	}
	//	err occurent
	return nil, nil, err
}

// GetNamespacesFromAccount get required params addresses, other skipped if value < 0
func (ref *NamespaceService) GetNamespacesFromAccount(ctx context.Context, address Address, nsId int,
	pageSize int) (nsInfo ListNamespaceInfo, resp *http.Response, err error) {

	addresses := make(Addresses, 1)
	addresses[0] = address

	return ref.GetNamespacesFromAccounts(ctx, addresses, nsId, pageSize)
}

// GetNamespacesFromAccounts get required params addresses, other skipped if value < 0
func (ref *NamespaceService) GetNamespacesFromAccounts(ctx context.Context, addresses Addresses, nsId int,
	pageSize int) (nsInfo ListNamespaceInfo, resp *http.Response, err error) {

	body, _ := addresses.MarshalJSON()

	url, comma := "", "?"

	if nsId >= 0 {
		url = comma + strconv.Itoa(nsId)
		comma = "&"
	}

	if pageSize >= 0 {
		url = comma + strconv.Itoa(pageSize)
	}

	url = "account/namespace" + url

	var req *http.Request
	req, err = ref.client.NewRequest("POST", url, &body)

	if err == nil {

		resp, err = ref.client.Do(ctx, req, nsInfo)

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
func generateNamespaceId(namespaceName string) int64 { /* public static  */

	namespacePath := generateNamespacePath(namespaceName)
	return namespacePath[len(namespacePath)-1]
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
func generateNamespacePath(name string) []int64 { /* public static  */

	//[]string parts = name.split(Pattern.quote("."))
	//[]int64 path = new Array[]int64()
	//if (parts.length == 0) {
	//	panic(IllegalIdentifierException{"invalid namespace name"})
	//}
	//else if (parts.length > 3) {
	//	panic(IllegalIdentifierException{"too many parts"})
	//}
	//
	//int64 namespaceId = int64.valueOf(0)
	//for (int i = 0; i < parts.length; i++) {
	//	if (!parts[i].matches("^[a-z0-9][a-z0-9-_]*$")) {
	//		panic(IllegalIdentifierException{"invalid namespace name"})
	//	}
	//
	//	namespaceId = generateId(parts[i], namespaceId)
	//	path.add(namespaceId)
	//}

	return []int64{0, 1}
}
