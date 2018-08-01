package sdk

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
)

type MosaicService service

func NewMosaicService(httpClient *http.Client, conf *Config) *MosaicService {
	ref := &MosaicService{client: NewClient(httpClient, conf)}

	return ref
}

const pathMosaic = "/mosaic"

type MosaicDefinitionDTO struct { /*  */
	namespaceId *uint64DTO
	mosaicId    *uint64DTO
	supply      *uint64DTO
	height      *uint64DTO
	owner       string
	properties  []*uint64DTO
	levy        interface{}
}
type MosaicInfoDTO struct {
	meta   NamespaceMosaicMetaDTO
	mosaic MosaicDefinitionDTO
}
type testT struct {
	MosaicIds []string
}

func (ref *MosaicService) GetMosaic(ctx context.Context, nsId string) (nsInfo *MosaicInfo, resp *http.Response, err error) {

	mosaicTest := testT{MosaicIds: []string{"d525ad41d95fcf29"}}
	nsInfoDTO := &MosaicInfoDTO{}
	resp, err = ref.client.DoNewRequest(ctx, "POST", pathMosaic, mosaicTest, nsInfoDTO)

	if err == nil {
		//nsInfo, err = NamespaceInfoFromDTO(nsInfoDTO)
		//if err == nil {
		return nsInfo, resp, err
		//}
		fmt.Println(mosaicTest)
	}
	//	err occurent
	return nil, resp, err
}
