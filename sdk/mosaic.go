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

type MosaicService service

// GetMosaic returns
// @get /mosaic/{mosaicId}
func (ref *MosaicService) GetMosaic(ctx context.Context, mosaicId *MosaicId) (*MosaicInfo, error) {
	if mosaicId == nil {
		return nil, ErrNilMosaicId
	}

	url := net.NewUrl(fmt.Sprintf(pathMosaic+"/%s", mosaicId.toHexString()))

	mscInfoDTO := &mosaicInfoDTO{}

	resp, err := ref.client.DoNewRequest(ctx, http.MethodGet, url.Encode(), nil, mscInfoDTO)
	if err != nil {
		return nil, err
	}

	if err = handleResponseStatusCode(resp, map[int]error{404: ErrResourceNotFound, 409: ErrArgumentNotValid}); err != nil {
		return nil, err
	}

	return mscInfoDTO.toStruct(ref.client.config.NetworkType)
}

// GetMosaics get list mosaics Info
// post @/mosaic/
func (ref *MosaicService) GetMosaics(ctx context.Context, mosaicIds []*MosaicId) ([]*MosaicInfo, error) {
	if len(mosaicIds) == 0 {
		return nil, ErrEmptyMosaicIds
	}

	nsInfosDTO := make([]*mosaicInfoDTO, 0)

	resp, err := ref.client.DoNewRequest(ctx, http.MethodPost, pathMosaic, &MosaicIds{mosaicIds}, &nsInfosDTO)
	if err != nil {
		return nil, err
	}

	if err = handleResponseStatusCode(resp, map[int]error{400: ErrInvalidRequest, 409: ErrArgumentNotValid}); err != nil {
		return nil, err
	}

	return mosaicInfoDTOsToMosaicInfos(nsInfosDTO, ref.client.config.NetworkType)
}

// GetMosaicsFromNamespace Get mosaics information from namespaceId (nsId)
func (ref *MosaicService) GetMosaicsFromNamespaceUpToMosaic(ctx context.Context, namespaceId *NamespaceId, mosaicId *MosaicId,
	pageSize int) ([]*MosaicInfo, error) {
	if namespaceId == nil {
		return nil, ErrNilNamespaceId
	}

	url := net.NewUrl(fmt.Sprintf(pathMosaicFromNamespace, namespaceId.toHexString()))

	if pageSize > 0 {
		url.SetParam("pageSize", strconv.Itoa(pageSize))
	}

	if mosaicId != nil {
		url.SetParam("id", mosaicId.toHexString())
	}

	mscInfoDTOArr := make([]*mosaicInfoDTO, 0)

	resp, err := ref.client.DoNewRequest(ctx, http.MethodGet, url.Encode(), nil, &mscInfoDTOArr)
	if err != nil {
		return nil, err
	}

	if err = handleResponseStatusCode(resp, map[int]error{409: ErrArgumentNotValid}); err != nil {
		return nil, err
	}

	return mosaicInfoDTOsToMosaicInfos(mscInfoDTOArr, ref.client.config.NetworkType)
}

// GetMosaicsFromNamespace Get mosaics information from namespaceId (nsId)
func (ref *MosaicService) GetMosaicsFromNamespace(ctx context.Context, namespaceId *NamespaceId, pageSize int) (mscInfo []*MosaicInfo, err error) {
	return ref.GetMosaicsFromNamespaceUpToMosaic(ctx, namespaceId, nil, pageSize)
}

// GetMosaicNames Get readable names for a set of mosaics
// post @/mosaic/names
func (ref *MosaicService) GetMosaicNames(ctx context.Context, mosaicIds []*MosaicId) ([]*MosaicName, error) {
	if len(mosaicIds) == 0 {
		return nil, ErrEmptyMosaicIds
	}

	mscNamesDTO := make([]*mosaicNameDTO, 0)

	resp, err := ref.client.DoNewRequest(ctx, http.MethodPost, pathMosaicNames, &MosaicIds{mosaicIds}, &mscNamesDTO)
	if err != nil {
		return nil, err
	}

	if err = handleResponseStatusCode(resp, map[int]error{400: ErrInvalidRequest, 409: ErrArgumentNotValid}); err != nil {
		return nil, err
	}

	return mosaicNameDTOsToMosaicNames(mscNamesDTO)
}
