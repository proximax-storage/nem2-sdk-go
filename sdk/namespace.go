package sdk

import (
	"errors"
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

// const routers path for methods NamespaceService
const (
	pathNamespacesFromAccounts = "/account/namespaces"
	pathNamespace              = "/namespace/"
	pathNamespacenames         = "/namespace/names"
	pathNamespacesFromAccount  = "/account/%s/namespaces"
)

// namespaceDTO temporary struct for reading responce & fill NamespaceInfo
type namespaceDTO struct {
	Type         int
	Depth        int
	Level0       *uint64DTO
	Level1       *uint64DTO
	Level2       *uint64DTO
	ParentId     uint64DTO
	Owner        string
	OwnerAddress string
	StartHeight  uint64DTO
	EndHeight    uint64DTO
}

// namespaceInfoDTO temporary struct for reading responce & fill NamespaceInfo
type namespaceInfoDTO struct {
	Meta      NamespaceMosaicMetaDTO
	Namespace namespaceDTO
}

//setNamespaceInfo create & return new NamespaceInfo from namespaceInfoDTO
func (ref *namespaceInfoDTO) setNamespaceInfo() (*NamespaceInfo, error) {
	pubAcc, err := NewPublicAccount(ref.Namespace.Owner, NetworkType(ref.Namespace.Type))
	if err != nil {
		return nil, err
	}

	return &NamespaceInfo{
		ref.Meta.Active,
		ref.Meta.Index,
		ref.Meta.Id,
		NamespaceType(ref.Namespace.Type),
		ref.Namespace.Depth,
		ref.extractLevels(),
		NewNamespaceId(ref.Namespace.ParentId.GetBigInteger(), ""),
		pubAcc,
		ref.Namespace.StartHeight.GetBigInteger(),
		ref.Namespace.EndHeight.GetBigInteger(),
	}, nil
}

func (ref *namespaceInfoDTO) extractLevels() []*NamespaceId {

	levels := make([]*NamespaceId, 0)

	if ref.Namespace.Level0 != nil {
		levels = append(levels, NewNamespaceId(ref.Namespace.Level0.GetBigInteger(), ""))
	}

	if ref.Namespace.Level1 != nil {
		levels = append(levels, NewNamespaceId(ref.Namespace.Level1.GetBigInteger(), ""))
	}

	if ref.Namespace.Level2 != nil {
		levels = append(levels, NewNamespaceId(ref.Namespace.Level2.GetBigInteger(), ""))
	}
	return levels
}

func (ref *NamespaceService) GetNamespace(ctx context.Context, nsId string) (nsInfo *NamespaceInfo, resp *http.Response, err error) {

	nsInfoDTO := &namespaceInfoDTO{}
	resp, err = ref.client.DoNewRequest(ctx, "GET", pathNamespace+nsId, nil, nsInfoDTO)

	if err != nil {
		return nil, resp, err
	}
	nsInfo, err = nsInfoDTO.setNamespaceInfo()
	if err != nil {
		return nil, resp, err
	}

	return nsInfo, resp, err
}

// namespaceNameDTO temporary struct for reading responce & fill NamespaceName
type namespaceNameDTO struct {
	NamespaceId uint64DTO
	Name        string
	ParentId    uint64DTO
}

func (ref *namespaceNameDTO) getNamespaceName() *NamespaceName {
	return &NamespaceName{
		NewNamespaceId(ref.NamespaceId.GetBigInteger(), ""),
		ref.Name,
		NewNamespaceId(ref.ParentId.GetBigInteger(), "")}
}

// GetNamespaceNames
func (ref *NamespaceService) GetNamespaceNames(ctx context.Context, nsIds *NamespaceIds) (nsList []*NamespaceName, resp *http.Response, err error) {
	res := make([]*namespaceNameDTO, 0)
	resp, err = ref.client.DoNewRequest(ctx, "POST", pathNamespacenames, &nsIds, &res)

	if err != nil {
		return nil, resp, err
	}

	for _, dto := range res {
		nsList = append(nsList, dto.getNamespaceName())
	}
	return nsList, resp, err
}

// GetNamespacesFromAccount get required params addresses, other skipped if value < 0
func (ref *NamespaceService) GetNamespacesFromAccount(ctx context.Context, address *Address, nsId string,
	pageSize int) (nsList ListNamespaceInfo, resp *http.Response, err error) {

	if address == nil {
		return nsList, nil, errors.New("address is null")
	}

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

	if (err != nil) || (ListNamespaceInfoFromDTO(res, &nsList) != nil) {
		//	err occurent
		return nsList, resp, err
	}

	return nsList, resp, err
}

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

	if (err != nil) || (ListNamespaceInfoFromDTO(res, &nsList) != nil) {
		//	err occurent
		return nsList, resp, err
	}

	return nsList, resp, err
}
func ListNamespaceInfoFromDTO(res []*namespaceInfoDTO, nsList *ListNamespaceInfo) error {

	for _, nsInfoDTO := range res {
		nsInfo, err := nsInfoDTO.setNamespaceInfo()
		if err != nil {
			return err
		}
		nsList.list = append(nsList.list, nsInfo)
	}

	return nil
}
