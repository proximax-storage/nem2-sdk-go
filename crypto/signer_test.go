// Copyright 2017 Author: Ruslan Bikchentaev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crypto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testDataForSigner = []byte("abcdefg")
var (
	keyPair, _          = NewRandomKeyPair()
	contextSignature, _ = NewSignatureFromBigInt(BigInteger_ONE(), BigInteger_ONE())
	contextDsaSigner    = CryptoEngines.DefaultEngine.CreateDsaSigner(keyPair)
)

func TestNewSignerFromKeyPair(t *testing.T) {
	sign := NewSignerFromKeyPair(keyPair, nil)

	signature, err := sign.Sign(testDataForSigner)

	if err != nil {
		t.Fatal(err)
	}
	sign1 := NewSigner(contextDsaSigner)
	signature1, err := sign1.Sign(testDataForSigner)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, signature1.R, signature.R, `signature.getR() and r must by not equal !`)
	assert.Equal(t, signature1.S, signature.S, `signature.getS() and s must by not equal !`)
}
func TestNewSigner(t *testing.T) {
	sign := NewSignerFromKeyPair(keyPair, nil)

	signature, err := sign.Sign(testDataForSigner)

	if err != nil {
		t.Fatal(err)
	}
	//sign.MakeSignatureCanonical()
	res := sign.Verify(testDataForSigner, signature)
	assert.Equal(t, true, res, "sign %v", signature)
}
func TestIsCanonicalSignatureDelegatesToDsaSigner(t *testing.T) {

	signer := NewSigner(contextDsaSigner)
	assert.Equal(t, true, signer.IsCanonicalSignature(contextSignature), " must by canonical")
}

func TestMakeSignatureCanonicalDelegatesToDsaSigner(t *testing.T) {

	signer := NewSigner(contextDsaSigner)
	signature, err := signer.MakeSignatureCanonical(contextSignature)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, contextSignature, signature, " must by canonical")
}
