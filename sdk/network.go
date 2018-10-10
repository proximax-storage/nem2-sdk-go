// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package sdk

import (
	"errors"
	"golang.org/x/net/context"
	"net/http"
	"strings"
)

type NetworkService service

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
		return MijinTest, resp, nil
	}

	return NotSupportedNet, resp, errors.New("network " + netDTO.Name + " is not supported yet by the sdk")
}
