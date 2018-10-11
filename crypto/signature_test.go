// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package crypto

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"strconv"
	"testing"
)

const (
	testSignatureR = "99512345"
	testSignatureS = "12351234"
)

func TestNewSignatureFromBigInt(t *testing.T) {

	rInt, err := strconv.ParseInt(testSignatureR, 10, 64)
	if err != nil {
		t.Fatal(err)
	}
	sInt, err := strconv.ParseInt(testSignatureS, 10, 64)
	if err != nil {
		t.Fatal(err)
	}

	r := big.NewInt(rInt)
	s := big.NewInt(sInt)

	signature, err := NewSignatureFromBigInt(r, s)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, r.Uint64(), signature.GetR().Uint64(), `signature.getR() and r must by equal !`)
	assert.Equal(t, s.Uint64(), signature.GetS().Uint64(), `signature.getS() and s must by equal (%v = %v)`,
		s.Bytes(), signature.S)

}
func TestNewSignatureFromBigInt_BadParams(t *testing.T) {
	_, err := NewSignatureFromBigInt(nil, nil)
	assert.Error(t, err, "we must get error - %s", errBadParamNewSignatureBigInt)
}
func TestNewSignatureFromBytes_Fail(t *testing.T) {
	_, err := NewSignatureFromBytes([]byte{1})
	assert.Error(t, err, "we must get error - %s", errBadParamNewSignatureFromBytes)
}
func TestNewSignatureFromBytes(t *testing.T) {

	originalSignature, err := createSignature(testSignatureR, testSignatureS)
	if err != nil {
		t.Fatal(err)
	}
	signature, err := NewSignatureFromBytes(originalSignature.Bytes())
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, originalSignature.R, signature.R, `signature.getR() and r must by not equal !`)
	assert.Equal(t, originalSignature.S, signature.S, `signature.getS() and s must by not equal !`)

}
func TestNewSignature(t *testing.T) {

	originalSignature, err := createSignature(testSignatureR, testSignatureS)
	if err != nil {
		t.Fatal(err)
	}
	signature, err := NewSignature(originalSignature.R, originalSignature.S)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, originalSignature.R, signature.R, `signature.getR() and r must by not equal !`)
	assert.Equal(t, originalSignature.S, signature.S, `signature.getS() and s must by not equal !`)

}
func TestNewSignature_Fail(t *testing.T) {
	_, err := NewSignature([]byte{0}, []byte{1})
	assert.Error(t, err, "we must get error - %s", errBadParamNewSignature)
}

const expectedSignature = "0c000000000000000000000000000000000000000000000000000000000000000102000000000000000000000000000000000000000000000000000000000000"

func TestSignature_String_HexRepresentation(t *testing.T) {

	signature, err := createSignature("12", "513")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expectedSignature, signature.String(), `signature.toString() and expectedSignature must by not equal !`)
}

// createSignature create Signature from two string values
func createSignature(strR, strS string) (*Signature, error) {
	rInt, err := strconv.ParseInt(strR, 10, 64)
	if err != nil {
		return nil, err
	}
	sInt, err := strconv.ParseInt(strS, 10, 64)
	if err != nil {
		return nil, err
	}

	r := big.NewInt(rInt)
	s := big.NewInt(sInt)
	signature, err := NewSignatureFromBigInt(r, s)
	if err != nil {
		return nil, err
	}

	return signature, nil
}
