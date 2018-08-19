package sdk

import (
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"strconv"
)

type MosaicService service

// const routers path for methods MosaicService
const (
	pathMosaic              = "/mosaic/"
	pathMosaicNames         = "/mosaic/names"
	pathMosaicFromNamespace = "/namespaces/%s/mosaic/"
)

type mosaicPropertiesDTO []uint64DTO

// NamespaceMosaicMetaDTO
type NamespaceMosaicMetaDTO struct {
	Active bool
	Index  int
	Id     string
}
type mosaicDefinitionDTO struct {
	NamespaceId uint64DTO
	MosaicId    uint64DTO
	Supply      uint64DTO
	Height      uint64DTO
	Owner       string
	Properties  mosaicPropertiesDTO
	Levy        interface{}
}

// mosaicInfoDTO is temporary struct for reading response & fill MosaicInfo
type mosaicInfoDTO struct {
	Meta   NamespaceMosaicMetaDTO
	Mosaic mosaicDefinitionDTO
}

func (dto mosaicPropertiesDTO) toStruct() *MosaicProperties {
	flags := "00" + dto[0].toBigInt().Text(2)
	bitMapFlags := flags[len(flags)-3:]

	return NewMosaicProperties(bitMapFlags[2] == '1',
		bitMapFlags[1] == '1',
		bitMapFlags[0] == '1',
		dto[1].toBigInt().Int64(),
		dto[2].toBigInt(),
	)
}

func (ref *mosaicInfoDTO) setMosaicInfo() (*MosaicInfo, error) {

	publicAcc, err := NewPublicAccount(ref.Mosaic.Owner, NetworkType(1))
	if err != nil {
		return nil, err
	}
	if len(ref.Mosaic.Properties) < 3 {
		return nil, errors.New("mosaic Properties is not valid")
	}
	mosaicID, err := NewMosaicId(ref.Mosaic.MosaicId.toBigInt(), "")
	if err != nil {
		return nil, err
	}
	return &MosaicInfo{
		ref.Meta.Active,
		ref.Meta.Index,
		ref.Meta.Id,
		NewNamespaceId(ref.Mosaic.NamespaceId.toBigInt(), ""),
		mosaicID,
		ref.Mosaic.Supply.toBigInt(),
		ref.Mosaic.Height.toBigInt(),
		publicAcc,
		ref.Mosaic.Properties.toStruct(),
	}, nil
}

// mosaicInfoDTO is temporary struct for reading response & fill MosaicName
type MosaicNameDTO struct {
	ParentId uint64DTO
	MosaicId uint64DTO
	Name     string
}
type MosaicNamesDTO []*MosaicNameDTO

func (ref MosaicNamesDTO) setMosaicNames() ([]*MosaicName, error) {
	mscNames := make([]*MosaicName, len(ref))
	for i, mscNameDTO := range ref {
		newMscId, err := NewMosaicId(mscNameDTO.MosaicId.toBigInt(), "")
		if err != nil {
			return nil, err
		}
		mscNames[i] = &MosaicName{
			newMscId,
			mscNameDTO.Name,
			NewNamespaceId(mscNameDTO.ParentId.toBigInt(), ""),
		}
	}

	return mscNames, nil
}

// mosaics get mosaics Info
// @get /mosaic/{mosaicId}
func (ref *MosaicService) GetMosaic(ctx context.Context, mosaicId string) (mscInfo *MosaicInfo, resp *http.Response, err error) {

	mscInfoDTO := &mosaicInfoDTO{}
	resp, err = ref.client.DoNewRequest(ctx, "GET", pathMosaic+mosaicId, nil, mscInfoDTO)

	if err != nil {
		return nil, resp, err
	}

	mscInfo, err = mscInfoDTO.setMosaicInfo()
	if err != nil {
		return nil, resp, err
	}

	return mscInfo, resp, nil
}

// GetMosaics get list mosaics Info
// post @/mosaic/
func (ref *MosaicService) GetMosaics(ctx context.Context, mosaicIds MosaicIds) (mscInfoArr MosaicsInfo, resp *http.Response, err error) {

	nsInfosDTO := make([]mosaicInfoDTO, 0)
	resp, err = ref.client.DoNewRequest(ctx, "POST", pathMosaic, &mosaicIds, &nsInfosDTO)

	if err != nil {
		return nil, resp, err
	}

	mscInfoArr = make([]*MosaicInfo, len(nsInfosDTO))
	for i, nsInfoDTO := range nsInfosDTO {
		mscInfoArr[i], err = nsInfoDTO.setMosaicInfo()
		if err != nil {
			return nil, resp, err
		}

	}
	return mscInfoArr, resp, err
}

// GetMosaicNames Get readable names for a set of mosaics
// post @/mosaic/names
func (ref *MosaicService) GetMosaicNames(ctx context.Context, mosaicIds *MosaicIds) (mscNames []*MosaicName, resp *http.Response, err error) {

	mscNamesDTO := make(MosaicNamesDTO, 0)
	resp, err = ref.client.DoNewRequest(ctx, "POST", pathMosaicNames, mosaicIds, &mscNamesDTO)

	if err != nil {
		return nil, resp, err
	}

	mscNames, err = mscNamesDTO.setMosaicNames()
	if err != nil {
		return nil, resp, err
	}

	return mscNames, resp, nil

}

// GetMosaicsFromNamespace Get mosaics information from namespaceId (nsId)
// get @/namespaces/{namespaceId}/mosaic/
func (ref *MosaicService) GetMosaicsFromNamespace(ctx context.Context, namespaceId string, mosaicId string,
	pageSize int) (mscInfo []*MosaicInfo, resp *http.Response, err error) {

	url, comma := "", "?"

	if mosaicId > "" {
		url = comma + "id=" + namespaceId
		comma = "&"
	}

	if pageSize > 0 {
		url += comma + "pageSize=" + strconv.Itoa(pageSize)
	}

	url = fmt.Sprintf(pathMosaicFromNamespace, namespaceId) + url

	mscInfoDTOArr := make([]*mosaicInfoDTO, 0)
	resp, err = ref.client.DoNewRequest(ctx, "GET", url, nil, &mscInfoDTOArr)

	if err != nil {
		return nil, resp, err
	}

	mscInfo = make([]*MosaicInfo, len(mscInfoDTOArr))
	for i, mscInfoDTO := range mscInfoDTOArr {

		mscInfo[i], err = mscInfoDTO.setMosaicInfo()
		if err != nil {
			return nil, resp, err
		}
	}

	return mscInfo, resp, nil
}

type mosaicDTO struct {
	MosaicId uint64DTO `json:"id"`
	Amount   uint64DTO `json:"amount"`
}

func (dto *mosaicDTO) toStruct() (*Mosaic, error) {
	id, err := NewMosaicId(dto.MosaicId.toBigInt(), "")
	if err != nil {
		return nil, err
	}
	return &Mosaic{id, dto.Amount.toBigInt()}, nil
}
