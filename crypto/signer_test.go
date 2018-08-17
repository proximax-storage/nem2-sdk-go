// Copyright 2017 Author: Ruslan Bikchentaev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crypto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testDataForSigner = []byte("abcdefg")

func TestNewSigner(t *testing.T) {
	keyPair, _ := NewRandomKeyPair()
	sign := NewSigner(keyPair)

	signature, err := sign.Sign(testDataForSigner)

	if err != nil {
		t.Fatal(err)
	}
	//sign.MakeSignatureCanonical()
	res := sign.Verify(testDataForSigner, signature)
	assert.Equal(t, true, res, "sign %v", signature)
}
