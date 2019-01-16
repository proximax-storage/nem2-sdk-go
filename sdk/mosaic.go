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

// MosaicService provides a set of methods for obtaining information about the mosaics
type MosaicService service

// GetMosaic returns
// @get /mosaic/{mosaicId}
func (ref *MosaicService) GetMosaic(ctx context.Context, mosaicId *MosaicId) (*MosaicInfo, error) {
	if mosaicId == nil {
		return nil, ErrNilMosaicId
	}

	url := net.NewUrl(fmt.Sprintf(mosaicRoute, mosaicId.toHexString()))

	dto := &mosaicInfoDTO{}

	resp, err := ref.client.DoNewRequest(ctx, http.MethodGet, url.Encode(), nil, dto)
	if err != nil {
		return nil, err
	}

	if err = handleResponseStatusCode(resp, map[int]error{404: ErrResourceNotFound, 409: ErrArgumentNotValid}); err != nil {
		return nil, err
	}

	mscInfo, err := dto.toStruct(ref.client.config.NetworkType)
	if err != nil {
		return nil, err
	}

	if err = ref.buildMosaicHierarchy(ctx, mscInfo); err != nil {
		return nil, err
	}

	return mscInfo, nil
}

// GetMosaics get list mosaics Info
// post @/mosaic/
func (ref *MosaicService) GetMosaics(ctx context.Context, mscIds []*MosaicId) ([]*MosaicInfo, error) {
	if len(mscIds) == 0 {
		return nil, ErrEmptyMosaicIds
	}

	dtos := mosaicInfoDTOs(make([]*mosaicInfoDTO, 0))

	resp, err := ref.client.DoNewRequest(ctx, http.MethodPost, mosaicsRoute, &mosaicIds{mscIds}, &dtos)
	if err != nil {
		return nil, err
	}

	if err = handleResponseStatusCode(resp, map[int]error{400: ErrInvalidRequest, 409: ErrArgumentNotValid}); err != nil {
		return nil, err
	}

	mscInfos, err := dtos.toStruct(ref.client.config.NetworkType)
	if err != nil {
		return nil, err
	}

	if err = ref.buildMosaicsHierarchy(ctx, mscInfos); err != nil {
		return nil, err
	}

	return mscInfos, nil
}

// GetMosaicsFromNamespaceUpToMosaic get mosaics information according to namespace ID & mosaic ID
func (ref *MosaicService) GetMosaicsFromNamespaceUpToMosaic(ctx context.Context, namespaceId *NamespaceId, mosaicId *MosaicId,
	pageSize int) ([]*MosaicInfo, error) {
	if namespaceId == nil {
		return nil, ErrNilNamespaceId
	}

	url := net.NewUrl(fmt.Sprintf(mosaicsFromNamespaceRoute, namespaceId.toHexString()))

	if pageSize > 0 {
		url.SetParam("pageSize", strconv.Itoa(pageSize))
	}

	if mosaicId != nil {
		url.SetParam("id", mosaicId.toHexString())
	}

	dtos := mosaicInfoDTOs(make([]*mosaicInfoDTO, 0))

	resp, err := ref.client.DoNewRequest(ctx, http.MethodGet, url.Encode(), nil, &dtos)
	if err != nil {
		return nil, err
	}

	if err = handleResponseStatusCode(resp, map[int]error{409: ErrArgumentNotValid}); err != nil {
		return nil, err
	}

	mscInfos, err := dtos.toStruct(ref.client.config.NetworkType)
	if err != nil {
		return nil, err
	}

	if err = ref.buildMosaicsHierarchy(ctx, mscInfos); err != nil {
		return nil, err
	}

	return mscInfos, nil
}

// GetMosaicsFromNamespace Get mosaics information from namespaceId (nsId)
func (ref *MosaicService) GetMosaicsFromNamespace(ctx context.Context, namespaceId *NamespaceId, pageSize int) (mscInfo []*MosaicInfo, err error) {
	return ref.GetMosaicsFromNamespaceUpToMosaic(ctx, namespaceId, nil, pageSize)
}

// GetMosaicNames Get readable names for a set of mosaics
// post @/mosaic/names
func (ref *MosaicService) GetMosaicNames(ctx context.Context, mscIds []*MosaicId) ([]*MosaicName, error) {
	if len(mscIds) == 0 {
		return nil, ErrEmptyMosaicIds
	}

	dtos := mosaicNameDTOs(make([]*mosaicNameDTO, 0))

	resp, err := ref.client.DoNewRequest(ctx, http.MethodPost, mosaicNamesRoute, &mosaicIds{mscIds}, &dtos)
	if err != nil {
		return nil, err
	}

	if err = handleResponseStatusCode(resp, map[int]error{400: ErrInvalidRequest, 409: ErrArgumentNotValid}); err != nil {
		return nil, err
	}

	return dtos.toStruct()
}

func (ref *MosaicService) buildMosaicHierarchy(ctx context.Context, mscInfo *MosaicInfo) error {
	if mscInfo == nil || mscInfo.Namespace == nil {
		return nil
	}

	if mscInfo.Namespace.NamespaceId == nil || namespaceIdToBigInt(mscInfo.Namespace.NamespaceId).Int64() == 0 {
		return nil
	}

	nsInfo, err := ref.client.Namespace.GetNamespace(ctx, mscInfo.Namespace.NamespaceId)
	if err != nil {
		return err
	}

	mscInfo.Namespace = nsInfo

	return ref.client.Namespace.buildNamespaceHierarchy(ctx, nsInfo)
}

func (ref *MosaicService) buildMosaicsHierarchy(ctx context.Context, mscInfos []*MosaicInfo) error {
	var err error

	for _, mscInfo := range mscInfos {
		if err = ref.buildMosaicHierarchy(ctx, mscInfo); err != nil {
			return err
		}
	}

	return nil
}
