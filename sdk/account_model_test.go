// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	testNEMPublicKey    = "b4f12e7c9f6946091e2cb8b6d3a12b50d17ccbbf646386ea27ce2946a7423dcf"
	testPublicKey1      = "321DE652C4D3362FC2DDF7800F6582F4A10CFEA134B81F8AB6E4BE78BBA4D18E"
	testEncodedAddress1 = "SBFBW6TUGLEWQIBCMTBMXXQORZKUP3WTVVTOKK5M"
)

var testAddressesForEncoded = map[NetworkType]string{
	MijinTest: "SARNASAS2BIAB6LMFA3FPMGBPGIJGK6IJETM3ZSP",
	Mijin:     "MARNASAS2BIAB6LMFA3FPMGBPGIJGK6IJE5K5RYU",
	TestNet:   "TARNASAS2BIAB6LMFA3FPMGBPGIJGK6IJE47FYR3",
	MainNet:   "NARNASAS2BIAB6LMFA3FPMGBPGIJGK6IJFJKUV32",
}

func TestGenerateNewAccount_NEM(t *testing.T) {
	acc, err := NewAccount(MijinTest)

	assert.Nilf(t, err, "NewAccount returned error: %s", err)
	assert.NotNil(t, acc.KeyPair.PrivateKey.String())
}

func TestGenerateEncodedAddress_NEM(t *testing.T) {
	for nType, testAddress := range testAddressesForEncoded {
		res, err := generateEncodedAddress(testNEMPublicKey, nType)

		assert.Nilf(t, err, "generateEncodedAddress returned error: %s", err)
		assert.Equal(t, testAddress, res)
	}
}

func TestGenerateEncodedAddress(t *testing.T) {
	res, err := generateEncodedAddress(testPublicKey1, MijinTest)

	assert.Nilf(t, err, "generateEncodedAddress returned error: %s", err)
	assert.Equal(t, testEncodedAddress1, res)
}
