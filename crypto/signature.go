// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.
package crypto

import (
	"encoding/hex"
	"errors"
	"github.com/proximax-storage/nem2-sdk-go/utils"
	"math"
	"math/big"
)

//Signature
type Signature struct {
	R []byte
	S []byte
}

var (
	errBadParamNewSignature          = errors.New("binary signature representation of r and s must both have 32 bytes length")
	errBadParamNewSignatureBigInt    = errors.New("bad parameters NewSignatureFromBigInt")
	errBadParamNewSignatureFromBytes = errors.New("binary signature representation must be 64 bytes")
)

// NewSignature R and S must fit into 32 bytes
func NewSignature(r []byte, s []byte) (*Signature, error) {
	if (len(r) != 32) || (len(s) != 32) {
		return nil, errBadParamNewSignature
	}
	ref := &Signature{r, s}
	return ref, nil
}
func NewSignatureFromBigInt(rInt, sInt *big.Int) (*Signature, error) {
	if (rInt == nil) || (sInt == nil) ||
		(rInt.Uint64() > math.MaxInt32) ||
		(sInt.Uint64() > math.MaxInt32) {
		return nil, errBadParamNewSignatureBigInt
	}

	r := utils.BigIntToByteArray(rInt, 32)
	s := utils.BigIntToByteArray(sInt, 32)

	return NewSignature(r, s)
}

//NewSignatureFromBytes Creates a new signature from bytes array 64
func NewSignatureFromBytes(b []byte) (*Signature, error) {
	if len(b) < 64 {
		return nil, errBadParamNewSignatureFromBytes
	}
	return NewSignature(b[:32], b[32:])
}

/**
 * Gets the R-part of the signature.
 *
 * @return The R-part of the signature.
 */
func (ref *Signature) GetR() *big.Int {

	return utils.BytesToBigInteger(ref.R)
}

//GetS Gets the S-part of the signature.
func (ref *Signature) GetS() *big.Int {

	return utils.BytesToBigInteger(ref.S)
}

//Bytes Gets a little-endian 64-byte representation of the signature.
func (ref *Signature) Bytes() []byte {

	return append(ref.R, ref.S...)
}

func (ref *Signature) String() string {

	return hex.EncodeToString(ref.Bytes())
}
