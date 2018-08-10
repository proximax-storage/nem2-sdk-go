// Copyright 2017 Author: Ruslan Bikchentaev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crypto

import (
	"errors"
	//ed25519 "github.com/proximax-storage/nem2-sdk-go/crypto/ed25519/arithmetic"
	"bytes"
	"fmt"
	"math/big"
)

//Ed25519CryptoEngine
type Ed25519CryptoEngine struct {
}

// @Override
func (ref *Ed25519CryptoEngine) GetCurve() Curve {

	return ED25519Ed25519Curve

} /*  */

// @Override
func (ref *Ed25519CryptoEngine) CreateDsaSigner(keyPair *KeyPair) *DsaSigner {

	return NewEd25519DsaSigner(keyPair)
}

// @Override
func (ref *Ed25519CryptoEngine) CreateKeyGenerator() *KeyGenerator {

	return NewEd25519KeyGenerator()
}

// @Override
func (ref *Ed25519CryptoEngine) CreateBlockCipher(senderKeyPair KeyPair, recipientKeyPair KeyPair) BlockCipher {

	return NewEd25519BlockCipher(senderKeyPair, recipientKeyPair)
}

// @Override
func (ref *Ed25519CryptoEngine) CreateKeyAnalyzer() KeyAnalyzer {

	return NewEd25519KeyAnalyzer()
}

// Ed25519DsaSigner
type Ed25519DsaSigner struct {
	KeyPair *KeyPair // private
	/**
	 * Creates a Ed25519 DSA signer.
	 *
	 * @param keyPair The key pair to use.
	 */
} /*  */
func NewEd25519DsaSigner(keyPair *KeyPair) *Ed25519DsaSigner {
	ref := &Ed25519DsaSigner{
		keyPair,
	}
	return ref
}

// @Override
func (ref *Ed25519DsaSigner) Sign(data []byte) (*Signature, error) {

	if !ref.KeyPair.HasPrivateKey() {
		return nil, errors.New("cannot sign without private key")
	}

	// Hash the private key to improve randomness.
	hash, err := HashesSha3_512(ref.KeyPair.privateKey.getBytes())
	if err != nil {
		return nil, err
	}
	// R = H(hash_b,...,hash_2b-1, data) where b=256.
	//r := ed25519.NewEd25519EncodedFieldElement(HashesSha3_512( hash[32: 64]), // only include the last 32 bytes of the private key hash
	//	data)
	//// Reduce size of R since we are calculating mod group order anyway
	//rModQ := r.modQ()
	//// R = rModQ * base point.
	//R := Ed25519Group.BASE_POINT.scalarMultiply(rModQ)
	//encodedR := R.encode()
	//// S = (R + H(encodedR, encodedA, data) * a) mod group order where
	//// encodedR and encodedA are the little endian encodings of the group element R and the public key A and
	//// a is the lower 32 bytes of hash after clamping.
	//h := NewEd25519EncodedFieldElement(Hashes.sha3_512(
	//	encodedR.Raw,
	//	ref.KeyPair.publicKey.Raw,
	//	data))
	//hModQ := h.modQ()
	//encodedS := hModQ.multiplyAndAddModQ(
	//	Ed25519Utils.prepareForScalarMultiply(ref.KeyPair.privateKey),
	//	rModQ)
	//// Signature is (encodedR, encodedS)
	//signature := NewSignature(encodedR.Raw, encodedS.Raw)
	signature := &Signature{}
	if !ref.IsCanonicalSignature(signature) {
		return nil, errors.New("Generated signature is not canonical")
	}

	return signature, nil
}

// @Override
func (ref *Ed25519DsaSigner) Verify(data []byte, signature *Signature) bool {

	if !ref.IsCanonicalSignature(signature) {
		return false
	}

	//todo: recheck this equvalents
	if b := make([]byte, 32); bytes.Equal(ref.KeyPair.publicKey.Raw, b) {
		return false
	}

	// h = H(encodedR, encodedA, data).
	rawEncodedR := signature.R
	rawEncodedA := ref.KeyPair.publicKey.Raw

	values, err := HashesSha3_512(
		rawEncodedR,
		rawEncodedA,
		data)
	if err != nil {
		fmt.Print(err)
		return false
	}
	h, err := NewEd25519EncodedFieldElement(values)
	if err != nil {
		fmt.Print(err)
		return false
	}
	// hReduced = h mod group order
	hModQ := h.modQ()
	// Must compute A.
	A := NewEd25519EncodedGroupElement(rawEncodedA).Decode()

	A.precomputeForDoubleScalarMultiplication()
	// R = encodedS * B - H(encodedR, encodedA, data) * A

	calculatedR := Ed25519Group.BASE_POINT.float64ScalarMultiplyVariableTime(
		A,
		hModQ,
		NewEd25519EncodedFieldElement(signature.getBinaryS()))
	// Compare calculated R to given R.
	encodedCalculatedR := calculatedR.encode().Raw
	return bytes.Equal(encodedCalculatedR, rawEncodedR)
}

// @Override
func (ref *Ed25519DsaSigner) IsCanonicalSignature(signature *Signature) bool {

	return -1 == signature.S.compareTo(Ed25519Group.GROUP_ORDER) &&
		1 == signature.S.compareTo(uint64.ZERO)
}

// @Override
func (ref *Ed25519DsaSigner) MakeSignatureCanonical(signature *Signature) Signature {

	s := NewEd25519EncodedFieldElement(Arrays.copyOf(signature.getBinaryS(), 64))
	Ed25519EncodedFieldElement
	sModQ = s.modQ()
	Ed25519EncodedFieldElement
	return NewSignature(signature.getBinaryR(), sModQ.Raw)
}

/**
 * Class that wraps the elliptic curve Ed25519.
 */type ed25519Curve struct {
	Curve
}

var Ed25519Curve = &ed25519Curve{}

// @Override
func (ref *ed25519Curve) GetGroupOrder() uint64 {
	return Ed25519Group.GROUP_ORDER
}

// @Override
func (ref *ed25519Curve) GetHalfGroupOrder() uint64 {
	return Ed25519Group.GROUP_ORDER.shiftRight(1)
}

// dummy class
type SecureRandom struct{}

func NewSecureRandom() *SecureRandom {
	return &SecureRandom{}
}
func (ref *SecureRandom) nextBytes([]byte) {

}

//Ed25519KeyGenerator Implementation of the key generator for Ed25519.
type Ed25519KeyGenerator struct {
	random *SecureRandom // private
} /* Ed25519KeyGenerator */
func NewEd25519KeyGenerator() *Ed25519KeyGenerator {
	ref := &Ed25519KeyGenerator{
		NewSecureRandom(),
	}
	return ref
}

// @Override
func (ref *Ed25519KeyGenerator) GenerateKeyPair() (*KeyPair, error) {

	seed := make([]byte, 32)
	ref.random.nextBytes(seed)
	// seed is the private key.
	privateKey := NewPrivateKey(&big.Int{}.SetBytes(seed))
	return NewKeyPair(privateKey, CryptoEngines.ed25519Engine())
}

// @Override
func (ref *Ed25519KeyGenerator) DerivePublicKey(privateKey *PrivateKey) *PublicKey {

	a := PrepareForScalarMultiply(privateKey)
	// a * base point is the public key.
	pubKey := Ed25519Group.BASE_POINT.scalarMultiply(a)
	Ed25519GroupElement
	// verification of signatures will be about twice as fast when pre-calculating
	// a suitable table of group elements.
	return NewPublicKey(pubKey.Encode().Raw)
}
