package sdk

import (
	"errors"
	"golang.org/x/net/context"
	"net/http"
	"time"
)

type MosaicService service

func NewMosaicService(httpClient *http.Client, conf *Config) *MosaicService {
	ref := &MosaicService{client: NewClient(httpClient, conf)}

	return ref
}

// NamespaceMosaicMetaDTO
type NamespaceMosaicMetaDTO struct {
	Active bool
	Index  int
	Id     string
}                                 /*  */
type mosaicDefinitionDTO struct { /*  */
	NamespaceId *uint64DTO
	MosaicId    *uint64DTO
	Supply      *uint64DTO
	Height      *uint64DTO
	Owner       string
	Properties  []*uint64DTO
	Levy        interface{}
}
type mosaicInfoDTO struct {
	Meta   NamespaceMosaicMetaDTO
	Mosaic mosaicDefinitionDTO
}

func (ref *mosaicInfoDTO) extractMosaicProperties() *MosaicProperties {

	flags := []byte("00" + ref.Mosaic.Properties[0].String())
	bitMapFlags := flags[len(flags)-3:]

	return NewMosaicProperties(bitMapFlags[2] == '1',
		bitMapFlags[1] == '1',
		bitMapFlags[0] == '1',
		ref.Mosaic.Properties[1][0].Int64(),
		time.Duration(ref.Mosaic.Properties[2].ExtractIntArray()))

}
func (ref *mosaicInfoDTO) getMosaicInfo() (*MosaicInfo, error) {

	publicAcc, err := NewPublicAccount(ref.Mosaic.Owner, NetworkType(1))
	if err != nil {
		return nil, err
	}
	if len(ref.Mosaic.Properties) < 3 {
		return nil, errors.New("Mosaic Properties is not valid")
	}
	return &MosaicInfo{
		ref.Meta.Active,
		ref.Meta.Index,
		ref.Meta.Id,
		NewNamespaceId(ref.Mosaic.NamespaceId, ""),
		NewMosaicFromID(ref.Mosaic.MosaicId),
		ref.Mosaic.Supply,
		ref.Mosaic.Height,
		publicAcc,
		ref.extractMosaicProperties(),
	}, nil
}

const pathMosaic = "/mosaic/"

func (ref *MosaicService) GetMosaic(ctx context.Context, mosaicId string) (nsInfo *MosaicInfo, resp *http.Response, err error) {

	nsInfoDTO := &mosaicInfoDTO{}
	resp, err = ref.client.DoNewRequest(ctx, "GET", pathMosaic+mosaicId, nil, nsInfoDTO)

	if err == nil {
		nsInfo, err = nsInfoDTO.getMosaicInfo()
		if err == nil {
			return nsInfo, resp, nil
		}
	}
	//	err occurent
	return nil, resp, err
}
func (ref *MosaicService) GetMosaics(ctx context.Context, mosaicId MosaicIds) (nsInfo []*MosaicInfo, resp *http.Response, err error) {

	nsInfosDTO := make([]mosaicInfoDTO, 0)
	resp, err = ref.client.DoNewRequest(ctx, "POST", pathMosaic, &mosaicId, &nsInfosDTO)

	if err == nil {
		nsInfo = make([]*MosaicInfo, len(nsInfosDTO))
		for i, nsInfoDTO := range nsInfosDTO {
			nsInfo[i], err = nsInfoDTO.getMosaicInfo()
			if err != nil {
				return nil, resp, err
			}

		}
		return nsInfo, resp, err
	}
	//	err occurent
	return nil, resp, err
}
