// Copyright 2017 Author: Ruslan Bikchentaev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crypto

import (
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

const testHexPrivatKeyValue = "227F"
const testHexPrivatKeyOdd = "ABC"
const testHexPrivatKeyNegative = "8000"
const testHexPrivatKeyMalformed = "22G75"

func getBigIntFromHex(hStr string) (*big.Int, error) {
	b, err := hexDecodeString(hStr)
	if err != nil {
		return nil, err
	}
	return (&big.Int{}).SetBytes(b), nil
}
func TestNewPrivatKeyfromHexString(t *testing.T) {
	key, err := NewPrivatKeyfromHexString(testHexPrivatKeyValue)

	assert.NoError(t, err, `NewPrivateKeyfromDecimalString("2275") must to return no error`)

	val, err := getBigIntFromHex(testHexPrivatKeyValue)
	if err != nil {
		t.Fatal(err)
	}
	assertPrivateKey(t, key, val)
}
func TestNewPrivatKeyfromHexString_OddLength(t *testing.T) {
	key, err := NewPrivatKeyfromHexString(testHexPrivatKeyOdd)
	if err != nil {
		t.Fatal(err)
	}

	b := []byte{0x0A, 0xBC, 0, 0}
	val := (&big.Int{}).SetBytes(b)
	assertPrivateKey(t, key, val)
}
func TestNewPrivatKeyfromHexString_Negative(t *testing.T) {
	key, err := NewPrivatKeyfromHexString(testHexPrivatKeyNegative)
	if err != nil {
		t.Fatal(err)
	}

	b := []byte{0x80, 0x00, 0, 0}
	val := (&big.Int{}).SetBytes(b)
	assertPrivateKey(t, key, val)
}
func TestNewPrivatKeyfromHexString_Malformed(t *testing.T) {
	_, err := NewPrivatKeyfromHexString(testHexPrivatKeyMalformed)

	assert.Error(t, err)
}
