package sdk

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"strconv"
)

// NamespaceService provides a set of methods for obtaining information about the namespace
type NamespaceService service

func NewNamespaceService(httpClient *http.Client, conf *Config) *NamespaceService {
	ref := &NamespaceService{client: NewClient(httpClient, conf)}

	return ref
}

// namespaceDTO temporary struct for reading responce & fill NamespaceInfo
type namespaceDTO struct {
	Type         int
	Depth        int
	Level0       *uint64DTO
	Level1       *uint64DTO
	Level2       *uint64DTO
	ParentId     *uint64DTO
	Owner        string
	OwnerAddress string
	StartHeight  *uint64DTO
	EndHeight    *uint64DTO
}

// namespaceInfoDTO temporary struct for reading responce & fill NamespaceInfo
type namespaceInfoDTO struct {
	Meta      NamespaceMosaicMetaDTO
	Namespace namespaceDTO
}

func (ref *namespaceInfoDTO) extractLevels() []*NamespaceId {

	levels := make([]*NamespaceId, 0)

	if ref.Namespace.Level0 != nil {
		levels = append(levels, NewNamespaceId(ref.Namespace.Level0, ""))
	}

	if ref.Namespace.Level1 != nil {
		levels = append(levels, NewNamespaceId(ref.Namespace.Level1, ""))
	}

	if ref.Namespace.Level2 != nil {
		levels = append(levels, NewNamespaceId(ref.Namespace.Level2, ""))
	}
	return levels
}

const pathNamespace = "/namespace/"

func (ref *NamespaceService) GetNamespace(ctx context.Context, nsId string) (nsInfo *NamespaceInfo, resp *http.Response, err error) {

	nsInfoDTO := &namespaceInfoDTO{}
	resp, err = ref.client.DoNewRequest(ctx, "GET", pathNamespace+nsId, nil, nsInfoDTO)

	if err == nil {
		nsInfo, err = NamespaceInfoFromDTO(nsInfoDTO)
		if err == nil {
			return nsInfo, resp, err
		}

	}
	//	err occurent
	return nil, resp, err
}

const pathNamespacenames = "/namespace/names"

// namespaceNameDTO temporary struct for reading responce & fill NamespaceName
type namespaceNameDTO struct {
	NamespaceId *uint64DTO
	Name        string
	ParentId    *uint64DTO
}

// GetNamespaceNames
func (ref *NamespaceService) GetNamespaceNames(ctx context.Context, nsIds *NamespaceIds) (nsList []*NamespaceName, resp *http.Response, err error) {
	res := make([]*namespaceNameDTO, 0)
	resp, err = ref.client.DoNewRequest(ctx, "POST", pathNamespacenames, &nsIds, &res)

	if err == nil {
		for _, val := range res {
			nsList = append(nsList, &NamespaceName{
				NewNamespaceId(val.NamespaceId, ""),
				val.Name,
				NewNamespaceId(val.ParentId, "")})
		}
		return nsList, resp, err

	}

	//	err occurent
	return nil, resp, err
}

const pathNamespacesFromAccount = "/account/%s/namespaces"

// GetNamespacesFromAccount get required params addresses, other skipped if value < 0
func (ref *NamespaceService) GetNamespacesFromAccount(ctx context.Context, address *Address, nsId string,
	pageSize int) (nsList ListNamespaceInfo, resp *http.Response, err error) {

	url, comma := "", "?"

	if nsId > "" {
		url = comma + "id=" + nsId
		comma = "&"
	}

	if pageSize > 0 {
		url += comma + "pageSize=" + strconv.Itoa(pageSize)
	}

	url = fmt.Sprintf(pathNamespacesFromAccount, address.Address) + url

	res := make([]*namespaceInfoDTO, 0)
	resp, err = ref.client.DoNewRequest(ctx, "GET", url, nil, &res)

	if err == nil {

		err = ListNamespaceInfoFromDTO(res, &nsList)

		if err == nil {
			return nsList, resp, err
		}
	}

	//	err occurent
	return nsList, resp, err
}

const pathNamespacesFromAccounts = "/account/namespaces"

// GetNamespacesFromAccounts get required params addresses, other skipped if value is empty
func (ref *NamespaceService) GetNamespacesFromAccounts(ctx context.Context, addresses *Addresses, nsId string,
	pageSize int) (nsList ListNamespaceInfo, resp *http.Response, err error) {

	url, comma := "", "?"

	if nsId > "" {
		url = comma + "id=" + nsId
		comma = "&"
	}

	if pageSize > 0 {
		url += comma + "pageSize=" + strconv.Itoa(pageSize)
	}

	url = pathNamespacesFromAccounts + url

	res := make([]*namespaceInfoDTO, 0)
	resp, err = ref.client.DoNewRequest(ctx, "POST", url, &addresses, &res)

	if err == nil {

		err = ListNamespaceInfoFromDTO(res, &nsList)
		if err == nil {
			return nsList, resp, err
		}
	}

	//	err occurent
	return nsList, resp, err
}
func ListNamespaceInfoFromDTO(res []*namespaceInfoDTO, nsList *ListNamespaceInfo) error {

	for _, nsInfoDTO := range res {
		nsInfo, err := NamespaceInfoFromDTO(nsInfoDTO)
		if err != nil {
			return err
		}
		nsList.list = append(nsList.list, nsInfo)
	}

	return nil
}

//NamespaceInfoFromDTO create & return new NamespaceInfo from namespaceInfoDTO
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
		NewNamespaceId(nsInfoDTO.Namespace.ParentId, ""),
		pubAcc,
		nsInfoDTO.Namespace.StartHeight,
		nsInfoDTO.Namespace.EndHeight,
	}, nil
}
