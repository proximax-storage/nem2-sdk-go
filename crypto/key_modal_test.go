// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package crypto

import (
	"github.com/proximax-storage/nem2-sdk-go/utils"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

var testPrivatKeyValue = big.NewInt(2275)

func TestNewPrivateKey(t *testing.T) {
	val := testPrivatKeyValue
	key := NewPrivateKey(val.Bytes())

	assertPrivateKey(t, key, val)
}

func assertPrivateKey(t *testing.T, key *PrivateKey, val *big.Int) {
	assert.Equal(t, val, key.value, `key.Raw and NewBigInteger("%d") must by equal !`, val.Int64())
	assert.Equal(t, val.Bytes(), key.Raw, `key.Raw and NewBigInteger("%d").Bytes must by equal !`, val.Int64())
}

func TestNewPrivateKeyfromDecimalString(t *testing.T) {
	val := testPrivatKeyValue
	str := val.String()
	key, err := NewPrivateKeyfromDecimalString(str)

	assert.NoError(t, err, `NewPrivateKeyfromDecimalString("2275") must to return no error`)
	assertPrivateKey(t, key, val)
}
func TestNewPrivateKeyfromDecimalString_Negative(t *testing.T) {
	val := big.NewInt(-2275)
	str := val.String()
	key, err := NewPrivateKeyfromDecimalString(str)

	if err != nil {
		t.Fatal(err)
	}
	assertPrivateKey(t, key, val)
}

const testDecimalPrivatKeyMalformed = "22A75"

func TestNewPrivatKeyfromDecimalString_Malformed(t *testing.T) {
	_, err := NewPrivateKeyfromDecimalString(testDecimalPrivatKeyMalformed)

	assert.Error(t, err)
}

const testHexKeyValue = "227F"
const testHexPrivatKeyOdd = "ABC"
const testHexPrivatKeyNegative = "8000"
const testHexKeyMalformed = "22G75"
const testKeyStringResult = "22ab71"

func getBigIntFromHex(hStr string) (*big.Int, error) {
	b, err := utils.HexDecodeStringOdd(hStr)
	if err != nil {
		return nil, err
	}
	return (&big.Int{}).SetBytes(b), nil
}
func TestNewPrivatKeyfromHexString(t *testing.T) {
	key, err := NewPrivateKeyfromHexString(testHexKeyValue)

	assert.NoError(t, err, `NewPrivateKeyfromDecimalString("2275") must to return no error`)

	val, err := getBigIntFromHex(testHexKeyValue)
	if err != nil {
		t.Fatal(err)
	}
	assertPrivateKey(t, key, val)
}
func TestNewPrivatKeyfromHexString_OddLength(t *testing.T) {
	key, err := NewPrivateKeyfromHexString(testHexPrivatKeyOdd)
	if err != nil {
		t.Fatal(err)
	}

	b := []byte{0x0A, 0xBC}
	val := (&big.Int{}).SetBytes(b)
	assertPrivateKey(t, key, val)
}
func TestNewPrivatKeyfromHexString_Negative(t *testing.T) {
	key, err := NewPrivateKeyfromHexString(testHexPrivatKeyNegative)
	if err != nil {
		t.Fatal(err)
	}

	b := []byte{0x80, 0x00}
	val := (&big.Int{}).SetBytes(b)
	assertPrivateKey(t, key, val)
}
func TestNewPrivatKeyfromHexString_Malformed(t *testing.T) {
	_, err := NewPrivateKeyfromHexString(testHexKeyMalformed)

	assert.Error(t, err)
}

// publicKey tests
var (
	testBytes         = []byte{0x22, 0xAB, 0x71}
	modifiedTestBytes = []byte{0x22, 0xAB, 0x72}
	testHexBytes      = []byte{0x22, 0x7F}
)

func TestNewPublicKey(t *testing.T) {
	key := NewPublicKey(testBytes)

	assert.Equal(t, testBytes, key.Raw, "not equal")
}

func TestNewPublicKeyfromHex(t *testing.T) {
	key, err := NewPublicKeyfromHex(testHexKeyValue)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, testHexBytes, key.Raw, "not equal")
}
func TestNewPublicKeyfromHex_Malformed(t *testing.T) {
	_, err := NewPublicKeyfromHex(testHexKeyMalformed)

	assert.Error(t, err)
}

func TestPublicKey_String(t *testing.T) {
	key := NewPublicKey(testBytes)

	assert.Equal(t, testKeyStringResult, key.String(), "wrong string")
}
