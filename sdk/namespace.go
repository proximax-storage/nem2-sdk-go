// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package sdk

import (
	"fmt"
	"github.com/proximax-storage/proximax-utils-go/net"
	"golang.org/x/net/context"
	"net/http"
	"strconv"
)

// NamespaceService provides a set of methods for obtaining information about the namespace
type NamespaceService service

// GetNamespace return full info about Namespace according to namespace ID
// @/namespace/
func (ref *NamespaceService) GetNamespace(ctx context.Context, nsId *NamespaceId) (*NamespaceInfo, error) {
	if nsId == nil {
		return nil, ErrNilNamespaceId
	}

	nsInfoDTO := &namespaceInfoDTO{}

	url := net.NewUrl(fmt.Sprintf(namespaceRoute, nsId.toHexString()))

	resp, err := ref.client.DoNewRequest(ctx, http.MethodGet, url.Encode(), nil, nsInfoDTO)
	if err != nil {
		return nil, err
	}

	if err = handleResponseStatusCode(resp, map[int]error{404: ErrResourceNotFound, 409: ErrArgumentNotValid}); err != nil {
		return nil, err
	}

	nsInfo, err := nsInfoDTO.toStruct()
	if err != nil {
		return nil, err
	}

	if err = ref.buildNamespaceHierarchy(ctx, nsInfo); err != nil {
		return nil, err
	}

	return nsInfo, nil
}

// GetNamespacesFromAccount get required params addresses, other skipped if value < 0
// @/account/%s/namespaces
func (ref *NamespaceService) GetNamespacesFromAccount(ctx context.Context, address *Address, nsId *NamespaceId,
	pageSize int) ([]*NamespaceInfo, error) {
	if address == nil {
		return nil, ErrNilAddress
	}

	url := net.NewUrl(fmt.Sprintf(namespacesFromAccountRoutes, address.Address))

	if nsId != nil {
		url.SetParam("id", nsId.toHexString())
	}

	if pageSize > 0 {
		url.SetParam("pageSize", strconv.Itoa(pageSize))
	}

	dtos := namespaceInfoDTOs(make([]*namespaceInfoDTO, 0))

	resp, err := ref.client.DoNewRequest(ctx, http.MethodGet, url.Encode(), nil, &dtos)
	if err != nil {
		return nil, err
	}

	if err = handleResponseStatusCode(resp, map[int]error{409: ErrArgumentNotValid}); err != nil {
		return nil, err
	}

	nsInfos, err := dtos.toStruct()
	if err != nil {
		return nil, err
	}

	if err = ref.buildNamespacesHierarchy(ctx, nsInfos); err != nil {
		return nil, err
	}

	return nsInfos, nil
}

// GetNamespacesFromAccounts get required params addresses, other skipped if value is empty
// @/account/namespaces
func (ref *NamespaceService) GetNamespacesFromAccounts(ctx context.Context, addrs []*Address, nsId *NamespaceId,
	pageSize int) ([]*NamespaceInfo, error) {
	if len(addrs) == 0 {
		return nil, ErrEmptyAddressesIds
	}

	url := net.NewUrl(namespacesFromAccountsRoute)

	if nsId != nil {
		url.AddParam("id", nsId.toHexString())
	}

	if pageSize > 0 {
		url.AddParam("pageSize", strconv.Itoa(pageSize))
	}

	dtos := namespaceInfoDTOs(make([]*namespaceInfoDTO, 0))

	resp, err := ref.client.DoNewRequest(ctx, http.MethodPost, url.Encode(), addresses(addrs), &dtos)
	if err != nil {
		return nil, err
	}

	if err = handleResponseStatusCode(resp, map[int]error{400: ErrInvalidRequest, 409: ErrArgumentNotValid}); err != nil {
		return nil, err
	}

	nsInfos, err := dtos.toStruct()
	if err != nil {
		return nil, err
	}

	if err = ref.buildNamespacesHierarchy(ctx, nsInfos); err != nil {
		return nil, err
	}

	return nsInfos, nil
}

// GetNamespaceNames return full info about Namespaces according to slice namespace ID
// @/namespace/names
func (ref *NamespaceService) GetNamespaceNames(ctx context.Context, nsIds []*NamespaceId) ([]*NamespaceName, error) {
	if len(nsIds) == 0 {
		return nil, ErrEmptyNamespaceIds
	}

	dtos := namespaceNameDTOs(make([]*namespaceNameDTO, 0))

	resp, err := ref.client.DoNewRequest(ctx, http.MethodPost, namespaceNamesRoute, &NamespaceIds{nsIds}, &dtos)
	if err != nil {
		return nil, err
	}

	if err = handleResponseStatusCode(resp, map[int]error{400: ErrInvalidRequest, 409: ErrArgumentNotValid}); err != nil {
		return nil, err
	}

	return dtos.toStruct()
}

func (ref *NamespaceService) buildNamespaceHierarchy(ctx context.Context, nsInfo *NamespaceInfo) error {
	if nsInfo == nil || nsInfo.Parent == nil {
		return nil
	}

	if nsInfo.Parent.NamespaceId == nil || namespaceIdToBigInt(nsInfo.Parent.NamespaceId).Int64() == 0 {
		return nil
	}

	parentNsInfo, err := ref.GetNamespace(ctx, nsInfo.Parent.NamespaceId)
	if err != nil {
		return err
	}

	nsInfo.Parent = parentNsInfo

	return ref.buildNamespaceHierarchy(ctx, nsInfo.Parent)
}

func (ref *NamespaceService) buildNamespacesHierarchy(ctx context.Context, nsInfos []*NamespaceInfo) error {
	var err error

	for _, nsInfo := range nsInfos {
		if err = ref.buildNamespaceHierarchy(ctx, nsInfo); err != nil {
			return err
		}
	}

	return nil
}
