// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package sdk

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/proximax-storage/proximax-utils-go/net"
	"golang.org/x/net/context"
	"strconv"
)

type MosaicService service

// mosaics get mosaics Info
// @get /mosaic/{mosaicId}
func (ref *MosaicService) GetMosaic(ctx context.Context, mosaicId *MosaicId) (*MosaicInfo, error) {
	mscInfoDTO := &mosaicInfoDTO{}

	resp, err := ref.client.DoNewRequest(ctx, "GET", pathMosaic+mosaicId.toHexString(), nil, mscInfoDTO)
	if err != nil {
		return nil, err
	}

	if err = handleResponseStatusCode(resp, map[int]error{404: ErrResourceNotFound, 409: ErrArgumentNotValid}); err != nil {
		return nil, err
	}

	mscInfo, err := mscInfoDTO.toStruct(ref.client.config.NetworkType)
	if err != nil {
		return nil, err
	}

	return mscInfo, nil
}

// GetMosaics get list mosaics Info
// post @/mosaic/
func (ref *MosaicService) GetMosaics(ctx context.Context, mosaicIds MosaicIds) (MosaicsInfo, error) {
	if len(mosaicIds.MosaicIds) == 0 {
		return nil, ErrEmptyMosaicIds
	}

	nsInfosDTO := make([]mosaicInfoDTO, 0)

	resp, err := ref.client.DoNewRequest(ctx, "POST", pathMosaic, &mosaicIds, &nsInfosDTO)
	if err != nil {
		return nil, err
	}

	if err = handleResponseStatusCode(resp, map[int]error{400: ErrInvalidRequest, 409: ErrArgumentNotValid}); err != nil {
		return nil, err
	}

	mscInfoArr := make([]*MosaicInfo, len(nsInfosDTO))

	for i, nsInfoDTO := range nsInfosDTO {
		mscInfoArr[i], err = nsInfoDTO.toStruct(ref.client.config.NetworkType)
		if err != nil {
			return nil, err
		}
	}

	return mscInfoArr, nil
}

// GetMosaicNames Get readable names for a set of mosaics
// post @/mosaic/names
/*func (ref *MosaicService) GetMosaicNames(ctx context.Context, mosaicIds MosaicIds) (mscNames []*MosaicName, err error) {
	mscNamesDTO := make(mosaicNamesDTO, 0)

	resp, err := ref.client.DoNewRequest(ctx, "POST", pathMosaicNames, &mosaicIds, &mscNamesDTO)
	if err != nil {
		return nil, err
	}

	if err = handleResponseStatusCode(resp, map[int]error{400: ErrInvalidRequest, 409: ErrArgumentNotValid}); err != nil {
		return nil, err
	}

	mscNames, err = mscNamesDTO.setMosaicNames()

	if err != nil {
		return nil, err
	}

	return mscNames, nil
}*/

// GetMosaicsFromNamespace Get mosaics information from namespaceId (nsId)
func (ref *MosaicService) GetMosaicsFromNamespaceUpToMosaic(ctx context.Context, namespaceId *NamespaceId, mosaicId *MosaicId,
	pageSize int) ([]*MosaicInfo, error) {
	if namespaceId == nil {
		return nil, ErrNilIdNamespace
	}

	url := net.NewUrl(fmt.Sprintf(pathMosaicFromNamespace, namespaceId.toHexString()))

	if pageSize > 0 {
		url.SetParam("pageSize", strconv.Itoa(pageSize))
	}

	if mosaicId != nil {
		url.SetParam("id", mosaicId.toHexString())
	}

	mscInfoDTOArr := make([]*mosaicInfoDTO, 0)

	resp, err := ref.client.DoNewRequest(ctx, "GET", url.Encode(), nil, &mscInfoDTOArr)

	if err != nil {
		return nil, err
	}

	if err = handleResponseStatusCode(resp, map[int]error{409: ErrArgumentNotValid}); err != nil {
		return nil, err
	}

	mscInfo := make([]*MosaicInfo, len(mscInfoDTOArr))
	for i, mscInfoDTO := range mscInfoDTOArr {

		mscInfo[i], err = mscInfoDTO.toStruct(ref.client.config.NetworkType)
		if err != nil {
			return nil, err
		}
	}

	return mscInfo, nil
}

// GetMosaicsFromNamespace Get mosaics information from namespaceId (nsId)
func (ref *MosaicService) GetMosaicsFromNamespace(ctx context.Context, namespaceId *NamespaceId, pageSize int) (mscInfo []*MosaicInfo, err error) {
	if namespaceId == nil {
		return nil, ErrNilIdNamespace
	}

	url := net.NewUrl(fmt.Sprintf(pathMosaicFromNamespace, namespaceId.toHexString()))
	url.SetParam("pageSize", strconv.Itoa(pageSize))

	mscInfoDTOArr := make([]*mosaicInfoDTO, 0)

	resp, err := ref.client.DoNewRequest(ctx, "GET", url.Encode(), nil, &mscInfoDTOArr)

	if err != nil {
		return nil, err
	}

	if err = handleResponseStatusCode(resp, map[int]error{409: ErrArgumentNotValid}); err != nil {
		return nil, err
	}

	mscInfo = make([]*MosaicInfo, len(mscInfoDTOArr))
	for i, mscInfoDTO := range mscInfoDTOArr {

		mscInfo[i], err = mscInfoDTO.toStruct(ref.client.config.NetworkType)
		if err != nil {
			return nil, err
		}
	}

	return mscInfo, nil
}

type mosaicPropertiesDTO []uint64DTO

// namespaceMosaicMetaDTO
type namespaceMosaicMetaDTO struct {
	Active bool
	Index  int
	Id     string
}

type mosaicDefinitionDTO struct {
	NamespaceId uint64DTO
	MosaicId    uint64DTO
	FullName    string
	Supply      uint64DTO
	Height      uint64DTO
	Owner       string
	Properties  mosaicPropertiesDTO
	Levy        interface{}
}

// mosaicInfoDTO is temporary struct for reading response & fill MosaicInfo
type mosaicInfoDTO struct {
	Meta   namespaceMosaicMetaDTO
	Mosaic mosaicDefinitionDTO
}

func (dto *mosaicPropertiesDTO) toStruct() *MosaicProperties {
	flags := "00" + dto[0].toBigInt().Text(2)
	bitMapFlags := flags[len(flags)-3:]

	return NewMosaicProperties(bitMapFlags[2] == '1',
		bitMapFlags[1] == '1',
		bitMapFlags[0] == '1',
		dto[1].toBigInt().Int64(),
		dto[2].toBigInt(),
	)
}

func (ref *mosaicInfoDTO) toStruct(networkType NetworkType) (*MosaicInfo, error) {
	publicAcc, err := NewAccountFromPublicKey(ref.Mosaic.Owner, networkType)

	if err != nil {
		return nil, err
	}

	if len(ref.Mosaic.Properties) < 3 {
		return nil, errors.New("mosaic Properties is not valid")
	}

	nsName, err := NewNamespaceId(ref.Mosaic.NamespaceId.toBigInt())

	if err != nil {
		return nil, err
	}

	return &MosaicInfo{
		Active:      ref.Meta.Active,
		Index:       ref.Meta.Index,
		MetaId:      ref.Meta.Id,
		NamespaceId: nsName,
		MosaicId:    NewMosaicId(ref.Mosaic.MosaicId.toBigInt()),
		Supply:      ref.Mosaic.Supply.toBigInt(),
		Height:      ref.Mosaic.Height.toBigInt(),
		Owner:       publicAcc,
		Properties:  ref.Mosaic.Properties.toStruct(),
	}, nil
}

// mosaicInfoDTO is temporary struct for reading response & fill MosaicName
type mosaicNameDTO struct {
	ParentId uint64DTO
	MosaicId uint64DTO
	Name     string
}

type mosaicNamesDTO []*mosaicNameDTO

func (ref mosaicNamesDTO) setMosaicNames() ([]*MosaicName, error) {
	mscNames := make([]*MosaicName, len(ref))
	for i, mscNameDTO := range ref {
		parentId, err := NewNamespaceId(mscNameDTO.ParentId.toBigInt())
		if err != nil {
			return nil, err
		}

		mscNames[i] = &MosaicName{
			NewMosaicId(mscNameDTO.MosaicId.toBigInt()),
			mscNameDTO.Name,
			parentId,
		}
	}

	return mscNames, nil
}

type mosaicDTO struct {
	MosaicId uint64DTO `json:"id"`
	Amount   uint64DTO `json:"amount"`
}

func (dto *mosaicDTO) toStruct() *Mosaic {
	return &Mosaic{NewMosaicId(dto.MosaicId.toBigInt()), dto.Amount.toBigInt()}
}
