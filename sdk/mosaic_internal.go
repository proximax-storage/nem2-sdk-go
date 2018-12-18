// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package sdk

import (
	"errors"
	"math/big"
)

func bigIntToMosaicId(bigInt *big.Int) *MosaicId {
	if bigInt == nil {
		return nil
	}

	mscId := MosaicId(*bigInt)

	return &mscId
}

func mosaicIdToBigInt(mscId *MosaicId) *big.Int {
	if mscId == nil {
		return nil
	}

	return (*big.Int)(mscId)
}

func generateMosaicId(namespaceName string, mosaicName string) (*MosaicId, error) {
	if mosaicName == "" {
		return nil, ErrInvalidMosaicName
	}

	namespacePath, err := GenerateNamespacePath(namespaceName)
	if err != nil {
		return nil, err
	}

	if !regValidMosaicName.MatchString(mosaicName) {
		return nil, ErrInvalidMosaicName
	}

	bigInt, err := generateId(mosaicName, namespacePath[len(namespacePath)-1])
	if err != nil {
		return nil, err
	}

	return bigIntToMosaicId(bigInt), nil
}

// mosaicInfoDTO is temporary struct for reading response & fill MosaicName
type mosaicNameDTO struct {
	ParentId uint64DTO
	MosaicId uint64DTO
	Name     string
}

func (m *mosaicNameDTO) toStruct() (*MosaicName, error) {
	mosaicId, err := NewMosaicId(m.MosaicId.toBigInt())
	if err != nil {
		return nil, err
	}

	parentId, err := NewNamespaceId(m.ParentId.toBigInt())
	if err != nil {
		return nil, err
	}

	return &MosaicName{
		MosaicId: mosaicId,
		Name:     m.Name,
		ParentId: parentId,
	}, nil
}

type mosaicNameDTOs []*mosaicNameDTO

func (m *mosaicNameDTOs) toStruct() ([]*MosaicName, error) {
	dtos := *m
	mscNames := make([]*MosaicName, 0, len(dtos))

	for _, dto := range dtos {
		mscName, err := dto.toStruct()
		if err != nil {
			return nil, err
		}

		mscNames = append(mscNames, mscName)
	}

	return mscNames, nil
}

type mosaicDTO struct {
	MosaicId uint64DTO `json:"id"`
	Amount   uint64DTO `json:"amount"`
}

func (dto *mosaicDTO) toStruct() (*Mosaic, error) {
	mosaicId, err := NewMosaicId(dto.MosaicId.toBigInt())
	if err != nil {
		return nil, err
	}

	return &Mosaic{mosaicId, dto.Amount.toBigInt()}, nil
}

type mosaicPropertiesDTO []uint64DTO

// namespaceMosaicMetaDTO
type namespaceMosaicMetaDTO struct {
	Active bool
	Index  int
	Id     string
}

type mosaicDefinitionDTO struct {
	MosaicId    uint64DTO
	NamespaceId uint64DTO
	Name        string
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
	flags := "00" + (*dto)[0].toBigInt().Text(2)
	bitMapFlags := flags[len(flags)-3:]

	return NewMosaicProperties(bitMapFlags[2] == '1',
		bitMapFlags[1] == '1',
		bitMapFlags[0] == '1',
		(*dto)[1].toBigInt().Int64(),
		(*dto)[2].toBigInt(),
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

	nsId, err := NewNamespaceId(ref.Mosaic.NamespaceId.toBigInt())
	if err != nil {
		return nil, err
	}

	mosaicId, err := NewMosaicId(ref.Mosaic.MosaicId.toBigInt())

	mscInfo := &MosaicInfo{
		Active:     ref.Meta.Active,
		Index:      ref.Meta.Index,
		FullName:   ref.Mosaic.Name,
		MetaId:     ref.Meta.Id,
		MosaicId:   mosaicId,
		Supply:     ref.Mosaic.Supply.toBigInt(),
		Height:     ref.Mosaic.Height.toBigInt(),
		Owner:      publicAcc,
		Properties: ref.Mosaic.Properties.toStruct(),
	}

	if nsId != nil && namespaceIdToBigInt(nsId).Int64() != 0 {
		mscInfo.Namespace = &NamespaceInfo{NamespaceId: nsId}
	}

	return mscInfo, nil
}

type mosaicInfoDTOs []*mosaicInfoDTO

func (m *mosaicInfoDTOs) toStruct(networkType NetworkType) ([]*MosaicInfo, error) {
	dtos := *m

	mscInfos := make([]*MosaicInfo, 0, len(dtos))

	for _, dto := range dtos {
		mscInfo, err := dto.toStruct(networkType)
		if err != nil {
			return nil, err
		}

		mscInfos = append(mscInfos, mscInfo)
	}

	return mscInfos, nil
}

type mosaicIds struct {
	MosaicIds []*MosaicId `json:"mosaicIds"`
}

func (ref *mosaicIds) MarshalJSON() ([]byte, error) {
	buf := []byte(`{"mosaicIds": [`)

	for i, nsId := range ref.MosaicIds {
		if i > 0 {
			buf = append(buf, ',')
		}

		buf = append(buf, []byte(`"`+nsId.toHexString()+`"`)...)
	}

	buf = append(buf, ']', '}')

	return buf, nil
}
