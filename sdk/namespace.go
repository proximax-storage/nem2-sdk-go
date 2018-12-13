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

// GetNamespace
// @/namespace/
func (ref *NamespaceService) GetNamespace(ctx context.Context, nsId *NamespaceId) (*NamespaceInfo, error) {
	if nsId == nil {
		return nil, ErrNilNamespaceId
	}

	nsInfoDTO := &namespaceInfoDTO{}

	url := net.NewUrl(fmt.Sprintf(pathNamespace+"/%s", nsId.toHexString()))

	resp, err := ref.client.DoNewRequest(ctx, http.MethodGet, url.Encode(), nil, nsInfoDTO)
	if err != nil {
		return nil, err
	}

	if err = handleResponseStatusCode(resp, map[int]error{404: ErrResourceNotFound, 409: ErrArgumentNotValid}); err != nil {
		return nil, err
	}

	return nsInfoDTO.toStruct()
}

// GetNamespacesFromAccount get required params addresses, other skipped if value < 0
// @/account/%s/namespaces
func (ref *NamespaceService) GetNamespacesFromAccount(ctx context.Context, address *Address, nsId *NamespaceId,
	pageSize int) ([]*NamespaceInfo, error) {
	if address == nil {
		return nil, ErrNilAddress
	}

	url := net.NewUrl(fmt.Sprintf(pathNamespacesFromAccount, address.Address))

	if nsId != nil {
		url.SetParam("id", nsId.toHexString())
	}

	if pageSize > 0 {
		url.SetParam("pageSize", strconv.Itoa(pageSize))
	}

	res := make([]*namespaceInfoDTO, 0)

	resp, err := ref.client.DoNewRequest(ctx, http.MethodGet, url.Encode(), nil, &res)
	if err != nil {
		return nil, err
	}

	if err = handleResponseStatusCode(resp, map[int]error{409: ErrArgumentNotValid}); err != nil {
		return nil, err
	}

	return namespaceDTOsToNamespaceInfos(res)
}

// GetNamespacesFromAccounts get required params addresses, other skipped if value is empty
// @/account/namespaces
func (ref *NamespaceService) GetNamespacesFromAccounts(ctx context.Context, addresses []*Address, nsId *NamespaceId,
	pageSize int) ([]*NamespaceInfo, error) {
	if len(addresses) == 0 {
		return nil, ErrEmptyAddressesIds
	}

	url := net.NewUrl(pathNamespacesFromAccounts)

	if nsId != nil {
		url.AddParam("id", nsId.toHexString())
	}

	if pageSize > 0 {
		url.AddParam("pageSize", strconv.Itoa(pageSize))
	}

	res := make([]*namespaceInfoDTO, 0)

	resp, err := ref.client.DoNewRequest(ctx, http.MethodPost, url.Encode(), &Addresses{List: addresses}, &res)
	if err != nil {
		return nil, err
	}

	if err = handleResponseStatusCode(resp, map[int]error{400: ErrInvalidRequest, 409: ErrArgumentNotValid}); err != nil {
		return nil, err
	}

	return namespaceDTOsToNamespaceInfos(res)
}

// GetNamespaceNames
// @/namespace/names
func (ref *NamespaceService) GetNamespaceNames(ctx context.Context, nsIds []*NamespaceId) ([]*NamespaceName, error) {
	if len(nsIds) == 0 {
		return nil, ErrEmptyNamespaceIds
	}

	res := make([]*namespaceNameDTO, 0)

	resp, err := ref.client.DoNewRequest(ctx, http.MethodPost, pathNamespaceNames, &NamespaceIds{nsIds}, &res)
	if err != nil {
		return nil, err
	}

	if err = handleResponseStatusCode(resp, map[int]error{400: ErrInvalidRequest, 409: ErrArgumentNotValid}); err != nil {
		return nil, err
	}

	return namespaceNameDTOsToNamespaceNames(res)
}
