// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package sdk

import (
	"github.com/proximax-storage/proximax-utils-go/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	testNetRoute = `{
  			"name": "TEST_NET",
  			"description": "catapult development network"
  	}`
	notSupportedRoute = `{
			"name": "",
			"description": "catapult development network"
	}`
)

func TestNetworkService_GetNetworkType(t *testing.T) {
	t.Run("TEST_NET", func(t *testing.T) {
		mockServ := newSdkMockWithRouter(&mock.Router{
			Path:     pathNetwork,
			RespBody: testNetRoute,
		})

		defer mockServ.Close()

		netType, err := mockServ.getTestNetClientUnsafe().Network.GetNetworkType(ctx)

		assert.Nilf(t, err, "NetworkService.GetNetworkType returned error=%s", err)

		assert.Equal(t, netType, TestNet)
	})

	t.Run("NotSupportedNet", func(t *testing.T) {
		mock := newSdkMockWithRouter(&mock.Router{
			Path:     pathNetwork,
			RespBody: notSupportedRoute,
		})

		defer mock.Close()

		netType, err := mock.getTestNetClientUnsafe().Network.GetNetworkType(ctx)

		assert.NotNil(t, err, "NetworkService.GetNetworkType should return error")
		assert.Equal(t, netType, NotSupportedNet)
	})
}

func TestExtractNetworkType(t *testing.T) {
	i := uint64(36888)

	nt := ExtractNetworkType(i)

	assert.Equal(t, MijinTest, nt)
}
