// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package sdk

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"strconv"
)

// NamespaceService provides a set of methods for obtaining information about the namespace
type NamespaceService service

//GetNamespace
// @/namespace/
func (ref *NamespaceService) GetNamespace(ctx context.Context, nsId *NamespaceId) (nsInfo *NamespaceInfo, resp *http.Response, err error) {

	nsInfoDTO := &namespaceInfoDTO{}
	url := pathNamespace + nsId.toHexString()
	resp, err = ref.client.DoNewRequest(ctx, "GET", url, nil, nsInfoDTO)

	if err != nil {
		return nil, resp, err
	}
	nsInfo, err = nsInfoDTO.getNamespaceInfo()
	if err != nil {
		return nil, resp, err
	}

	return nsInfo, resp, err
}

// GetNamespaceNames
//@/namespace/names
func (ref *NamespaceService) GetNamespaceNames(ctx context.Context, nsIds NamespaceIds) (nsList []*NamespaceName, resp *http.Response, err error) {

	if len(nsIds.List) == 0 {
		return nil, nil, errEmptyNamespaceIds
	}
	res := make([]*namespaceNameDTO, 0)
	resp, err = ref.client.DoNewRequest(ctx, "POST", pathNamespacenames, &nsIds, &res)

	if err != nil {
		return nil, resp, err
	}

	for _, dto := range res {
		nsName, err := dto.getNamespaceName()
		if err != nil {
			return nil, resp, err
		}
		nsList = append(nsList, nsName)
	}
	return nsList, resp, err
}

// GetNamespacesFromAccount get required params addresses, other skipped if value < 0
// @/account/%s/namespaces
func (ref *NamespaceService) GetNamespacesFromAccount(ctx context.Context, address *Address, nsId string,
	pageSize int) (nsList ListNamespaceInfo, resp *http.Response, err error) {

	if address == nil {
		return nsList, nil, errNullAddress
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

	if (err != nil) || (listNamespaceInfoFromDTO(res, &nsList) != nil) {
		//	err occurent
		return nsList, resp, err
	}

	return nsList, resp, err
}

// GetNamespacesFromAccounts get required params addresses, other skipped if value is empty
// @/account/namespaces
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

	if (err != nil) || (listNamespaceInfoFromDTO(res, &nsList) != nil) {
		//	err occurent
		return nsList, resp, err
	}

	return nsList, resp, err
}

type namespaceIdDTO uint64DTO

func (dto *namespaceIdDTO) toStruct() (*NamespaceId, error) {
	return NewNamespaceId(uint64DTO(*dto).toBigInt())
}

// namespaceNameDTO temporary struct for reading responce & fill NamespaceName
type namespaceNameDTO struct {
	NamespaceId uint64DTO
	Name        string
	ParentId    uint64DTO
}

func (ref *namespaceNameDTO) getNamespaceName() (*NamespaceName, error) {
	nsId, err := NewNamespaceId(ref.NamespaceId.toBigInt())
	if err != nil {
		return nil, err
	}
	parentId, err := NewNamespaceId(ref.ParentId.toBigInt())
	if err != nil {
		return nil, err
	}

	return &NamespaceName{
		nsId,
		ref.Name,
		parentId,
	}, nil
}

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
	Meta      namespaceMosaicMetaDTO
	Namespace namespaceDTO
}

//getNamespaceInfo create & return new NamespaceInfo from namespaceInfoDTO
func (ref *namespaceInfoDTO) getNamespaceInfo() (*NamespaceInfo, error) {
	pubAcc, err := NewPublicAccount(ref.Namespace.Owner, NetworkType(ref.Namespace.Type))
	if err != nil {
		return nil, err
	}

	parentId, err := NewNamespaceId(ref.Namespace.ParentId.toBigInt())
	if err != nil {
		return nil, err
	}

	levels, err := ref.extractLevels()
	if err != nil {
		return nil, err
	}
	return &NamespaceInfo{
		ref.Meta.Active,
		ref.Meta.Index,
		ref.Meta.Id,
		NamespaceType(ref.Namespace.Type),
		ref.Namespace.Depth,
		levels,
		parentId,
		pubAcc,
		ref.Namespace.StartHeight.toBigInt(),
		ref.Namespace.EndHeight.toBigInt(),
	}, nil
}

func (ref *namespaceInfoDTO) extractLevels() ([]*NamespaceId, error) {

	levels := make([]*NamespaceId, 0)

	if ref.Namespace.Level0 != nil {
		nsName, err := NewNamespaceId(ref.Namespace.Level0.toBigInt())
		if err != nil {
			return nil, err
		}
		levels = append(levels, nsName)
	}

	if ref.Namespace.Level1 != nil {
		nsName, err := NewNamespaceId(ref.Namespace.Level1.toBigInt())
		if err != nil {
			return nil, err
		}
		levels = append(levels, nsName)
	}

	if ref.Namespace.Level2 != nil {
		nsName, err := NewNamespaceId(ref.Namespace.Level2.toBigInt())
		if err != nil {
			return nil, err
		}
		levels = append(levels, nsName)
	}

	return levels, nil
}

func listNamespaceInfoFromDTO(res []*namespaceInfoDTO, nsList *ListNamespaceInfo) error {

	for _, nsInfoDTO := range res {
		nsInfo, err := nsInfoDTO.getNamespaceInfo()
		if err != nil {
			return err
		}
		nsList.list = append(nsList.list, nsInfo)
	}

	return nil
}
