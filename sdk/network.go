// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package sdk

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
)

// NetworkService provides a set of methods for obtaining information about the Network
type NetworkService service

type networkDTO struct {
	Name        string
	Description string
}

// GetNetworkType return current network type
// @get /network
func (ref *NetworkService) GetNetworkType(ctx context.Context) (NetworkType, error) {
	netDTO := &networkDTO{}

	resp, err := ref.client.DoNewRequest(ctx, http.MethodGet, networkRoute, nil, netDTO)

	if err != nil {
		return 0, err
	}

	if err = handleResponseStatusCode(resp, nil); err != nil {
		return NotSupportedNet, err
	}

	networkType := NetworkTypeFromString(netDTO.Name)

	if networkType == NotSupportedNet {
		err = fmt.Errorf("network %s is not supported yet by the sdk", netDTO.Name)
	}

	return networkType, err
}
