// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package crypto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testDataForSigner = []byte("abcdefg")
var (
	keyPair          *KeyPair
	contextSignature *Signature
	contextDsaSigner DsaSigner
)

func init() {

	keyPair, _ = NewRandomKeyPair()
	contextSignature, _ = NewSignatureFromBigInt(BigInteger_ONE(), BigInteger_ONE())
	contextDsaSigner = CryptoEngines.DefaultEngine.CreateDsaSigner(keyPair)
}

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
	for i := 0; i < numIter; i++ {
		keyPair, err := NewRandomKeyPair()
		assert.Nil(t, err)
		sign := NewSignerFromKeyPair(keyPair, nil)

		signature, err := sign.Sign(testDataForSigner)

		if err != nil {
			t.Fatal(err)
		}
		res := sign.Verify(testDataForSigner, signature)
		assert.Truef(t, res, "iter=%d, sign %v", i+1, signature)
	}
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
