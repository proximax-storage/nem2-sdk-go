package sdk

import (
	"golang.org/x/net/context"
	"net/http"
	"strconv"
)

type NamespaceService service

func NewNamespaceService(httpClient *http.Client, conf *Config) *NamespaceService {
	ref := &NamespaceService{client: NewClient(httpClient, conf)}

	return ref
}

type NamespaceDTO struct {
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
} /* NamespaceDTO */
type NamespaceInfoDTO struct {
	Meta      NamespaceMosaicMetaDTO
	Namespace NamespaceDTO
}

const pathNamespace = "/namespace/"

func (ref *NamespaceService) GetNamespace(ctx context.Context, nsId string) (nsInfo *NamespaceInfo, resp *http.Response, err error) {

	nsInfoDTO := &NamespaceInfoDTO{}
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

type NamespaceNameDTO struct {
	namespaceId *uint64DTO
	name        string
	parentId    *uint64DTO
} /* NamespaceNameDTO */
type arrNamespaceNameDTO []*NamespaceNameDTO

func (ref *NamespaceService) GetNamespaceNames(ctx context.Context, nsIds *NamespaceIds) (nsList []*NamespaceName, resp *http.Response, err error) {
	res := make([]*NamespaceNameDTO, 0)
	resp, err = ref.client.DoNewRequest(ctx, "POST", pathNamespacenames, &nsIds, &res)

	if err == nil {
		for _, val := range res {
			nsList = append(nsList, &NamespaceName{
				NewNamespaceId(val.namespaceId, ""),
				val.name,
				NewNamespaceId(val.parentId, "")})
		}
		return nsList, resp, err

	}

	//	err occurent
	return nil, resp, err
}

// GetNamespacesFromAccount get required params addresses, other skipped if value < 0
func (ref *NamespaceService) GetNamespacesFromAccount(ctx context.Context, address *Address, nsId string,
	pageSize int) (nsInfo ListNamespaceInfo, resp *http.Response, err error) {

	addresses := &Addresses{}
	addresses.AddAddress(address)

	return ref.GetNamespacesFromAccounts(ctx, addresses, nsId, pageSize)
}

const pathNamespacesFromAccount = "/account/namespaces"

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

	url = pathNamespacesFromAccount + url

	res := make([]*NamespaceInfoDTO, 0)
	resp, err = ref.client.DoNewRequest(ctx, "POST", url, &addresses, &res)

	if err == nil {

		for _, nsInfoDTO := range res {
			nsInfo, err := NamespaceInfoFromDTO(nsInfoDTO)
			if err != nil {
				return nsList, resp, err
			}
			nsList.list = append(nsList.list, nsInfo)
		}

		if err == nil {
			return nsList, resp, err
		}
	}

	//	err occurent
	return nsList, resp, err
}

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
