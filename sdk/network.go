package sdk

import (
	"errors"
	"golang.org/x/net/context"
	"net/http"
	"strings"
)

type NetworkService service

// const routers path for methods MosaicService
const (
	pathNetwork = "/network"
)

type networkDTO struct {
	Name        string
	Description string
}

// mosaics get mosaics Info
// @get /network
func (ref *NetworkService) GetNetworkType(ctx context.Context) (mscInfo NetworkType, resp *http.Response, err error) {

	netDTO := &networkDTO{}
	resp, err = ref.client.DoNewRequest(ctx, "GET", pathNetwork, nil, netDTO)

	if err != nil {
		return 0, resp, err
	}

	if strings.ToLower(netDTO.Name) == "mijintest" {
		return MIJIN_TEST, resp, nil
	}

	return networkTypeNOT_SUPPORTED, resp, errors.New("network " + netDTO.Name + " is not supported yet by the sdk")
}
