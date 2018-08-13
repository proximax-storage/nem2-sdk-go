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

	return Ed25519Curve

} /*  */

// @Override
func (ref *Ed25519CryptoEngine) CreateDsaSigner(keyPair *KeyPair) DsaSigner {

	return NewEd25519DsaSigner(keyPair)
}

// @Override
func (ref *Ed25519CryptoEngine) CreateKeyGenerator() KeyGenerator {

	return NewEd25519KeyGenerator()
}

// @Override
func (ref *Ed25519CryptoEngine) CreateBlockCipher(senderKeyPair *KeyPair, recipientKeyPair *KeyPair) BlockCipher {

	return NewEd25519BlockCipher(senderKeyPair, recipientKeyPair)
}

// @Override
func (ref *Ed25519CryptoEngine) CreateKeyAnalyzer() KeyAnalyzer {

	return NewEd25519KeyAnalyzer()
}

/**
 * Implementation of the block cipher for Ed25519.
 */
type Ed25519BlockCipher struct { /* public  */

	senderKeyPair    *KeyPair      // private
	recipientKeyPair *KeyPair      // private
	random           *SecureRandom // private
	keyLength        int           // private
} /* Ed25519BlockCipher */
func NewEd25519BlockCipher(senderKeyPair *KeyPair, recipientKeyPair *KeyPair) *Ed25519BlockCipher { /* public  */
	ref := &Ed25519BlockCipher{
		senderKeyPair,
		recipientKeyPair,
		NewSecureRandom(),
		len(recipientKeyPair.publicKey.Raw),
	}
	return ref
}
func (ref *Ed25519BlockCipher) setupBlockCipher(sharedKey []byte, ivData []byte, forEncryption bool) *BufferedBlockCipher { /* private  */

	// Setup cipher parameters with key and IV.
	keyParam := NewKeyParameter(sharedKey) //
	params := NewParametersWithIV(keyParam, ivData)
	//
	// Setup AES cipher in CBC mode with PKCS7 padding.
	padding := NewPKCS7Padding()
	//
	cipher := NewPaddedBufferedBlockCipher(NewCBCBlockCipher(NewAESEngine()), padding) //
	cipher.reset()
	cipher.init(forEncryption, params)
	return cipher
}

func (ref *Ed25519BlockCipher) GetSharedKey(privateKey *PrivateKey, publicKey *PublicKey, salt []byte) ([]byte, error) { /* private  */

	senderA := NewEd25519EncodedGroupElement(publicKey.Raw).Decode()
	senderA.PrecomputeForScalarMultiplication()
	sharedKey := senderA.scalarMultiply(PrepareForScalarMultiply(privateKey)).Encode().Raw
	for i := 0; i < ref.keyLength; i++ {
		sharedKey[i] ^= salt[i]
	}

	return HashesSha3_256(sharedKey)
}

// @Override
func (ref *Ed25519BlockCipher) Encrypt(input []byte) []byte { /* public  */

	// Setup salt.
	salt := make([]byte, ref.keyLength)
	ref.random.nextBytes(salt)
	// Derive shared key.
	sharedKey, err := ref.GetSharedKey(ref.senderKeyPair.privateKey, ref.recipientKeyPair.publicKey, salt)
	if err != nil {
		fmt.Println(err)
	}
	// Setup IV.
	ivData := make([]byte, 16)
	ref.random.nextBytes(ivData)
	// Setup block cipher.
	cipher := ref.setupBlockCipher(sharedKey, ivData, true)
	// Encode.
	buf := ref.transform(cipher, input)
	if nil == buf {
		return nil
	}

	result := append(append(salt, ivData...), buf...)

	return result
}

// @Override
func (ref *Ed25519BlockCipher) Decrypt(input []byte) []byte { /* public  */

	if len(input) < 64 {
		return nil
	}

	salt := input[:ref.keyLength]
	ivData := input[ref.keyLength:48]
	encData := input[48:]
	// Derive shared key.
	sharedKey, err := ref.GetSharedKey(ref.recipientKeyPair.privateKey, ref.senderKeyPair.publicKey, salt)
	if err != nil {
		fmt.Println(err)
	}
	// Setup block cipher.
	cipher := ref.setupBlockCipher(sharedKey, ivData, false)
	// Decode.
	return ref.transform(cipher, encData)
}

func (ref *Ed25519BlockCipher) transform(cipher *BufferedBlockCipher, data []byte) []byte { /* private  */

	buf := make([]byte, cipher.GetOutputSize(len(data)))
	length := cipher.processBytes(data, 0, len(data), buf, 0)
	length += cipher.doFinal(buf, length)

	return buf
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
	//R = H(hash_b,...,hash_2b-1, data) where b=256.
	b, err := HashesSha3_512(hash[32:64], data) // only include the last 32 bytes of the private key hash
	if err != nil {
		return nil, err
	}
	r, err := NewEd25519EncodedFieldElement(b)
	if err != nil {
		return nil, err
	}

	// Reduce size of R since we are calculating mod group order anyway
	rModQ := r.modQ()
	// R = rModQ * base point.
	R := Ed25519Group.BASE_POINT.scalarMultiply(rModQ)
	encodedR := R.Encode()
	// S = (R + H(encodedR, encodedA, data) * a) mod group order where
	// encodedR and encodedA are the little endian encodings of the group element R and the public key A and
	// a is the lower 32 bytes of hash after clamping.
	b, err = HashesSha3_512(encodedR.Raw, ref.KeyPair.publicKey.Raw, data)
	if err != nil {
		return nil, err
	}
	h, err := NewEd25519EncodedFieldElement(b)
	if err != nil {
		return nil, err
	}
	hModQ := h.modQ()
	encodedS := hModQ.multiplyAndAddModQ(PrepareForScalarMultiply(ref.KeyPair.privateKey), rModQ)
	// Signature is (encodedR, encodedS)
	signature, err := NewSignature(encodedR.Raw, encodedS.Raw)
	if err != nil {
		return nil, err
	}
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

	A.PrecomputeForDoubleScalarMultiplication()
	// R = encodedS * B - H(encodedR, encodedA, data) * A

	b, err := NewEd25519EncodedFieldElement(signature.S)
	//todo: may by add arameter error ?
	if err != nil {
		fmt.Println(err)
		return false
	}
	calculatedR := Ed25519Group.BASE_POINT.doubleScalarMultiplyVariableTime(
		A,
		hModQ,
		b)
	// Compare calculated R to given R.
	encodedCalculatedR := calculatedR.Encode().Raw
	return bytes.Equal(encodedCalculatedR, rawEncodedR)
}

// @Override
func (ref *Ed25519DsaSigner) IsCanonicalSignature(signature *Signature) bool {

	return uint64(signature.GetS()) != Ed25519Group.GROUP_ORDER.Uint64() &&
		signature.GetS() == 0 //uint64.ZERO
}

// @Override
func (ref *Ed25519DsaSigner) MakeSignatureCanonical(signature *Signature) (*Signature, error) {

	s, err := NewEd25519EncodedFieldElement(signature.S[:64])
	if err != nil {
		fmt.Println(err)
	}
	sModQ := s.modQ()
	return NewSignature(signature.R, sModQ.Raw)
}

/**
 * Class that wraps the elliptic curve Ed25519.
 */type ed25519Curve struct {
	Curve
}

var Ed25519Curve = &ed25519Curve{}

// @Override
func (ref *ed25519Curve) GetGroupOrder() *big.Int {
	return Ed25519Group.GROUP_ORDER
}

// @Override
func (ref *ed25519Curve) GetHalfGroupOrder() uint64 {
	return Ed25519Group.GROUP_ORDER.Uint64() >> 1
}

// dummy class
type SecureRandom struct{}

//todo: implement in future
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
	privateKey := NewPrivateKey((&big.Int{}).SetBytes(seed))
	return NewKeyPair(privateKey, nil, CryptoEngines.Ed25519Engine)
}

// @Override
func (ref *Ed25519KeyGenerator) DerivePublicKey(privateKey *PrivateKey) *PublicKey {

	a := PrepareForScalarMultiply(privateKey)
	// a * base point is the public key.
	pubKey := Ed25519Group.BASE_POINT.scalarMultiply(a)
	// verification of signatures will be about twice as fast when pre-calculating
	// a suitable table of group elements.
	return NewPublicKey(string(pubKey.Encode().Raw))
}
