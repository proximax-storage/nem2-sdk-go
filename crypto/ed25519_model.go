// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package crypto

import (
	"crypto/rand"
	rand2 "crypto/rand"
	"errors"
	"fmt"
	"io"
	"math/big"
)

//Ed25519CryptoEngine wraps a cryptographic engine ed25519
type Ed25519CryptoEngine struct {
}

// GetCurve implemented interface CryptoEngine method
func (ref *Ed25519CryptoEngine) GetCurve() Curve {

	return Ed25519Curve

}

// CreateDsaSigner implemented interface CryptoEngine method
func (ref *Ed25519CryptoEngine) CreateDsaSigner(keyPair *KeyPair) DsaSigner {

	return NewEd25519DsaSigner(keyPair)
}

// CreateKeyGenerator implemented interface CryptoEngine method
func (ref *Ed25519CryptoEngine) CreateKeyGenerator() KeyGenerator {

	return NewEd25519KeyGenerator()
}

// CreateBlockCipher implemented interface CryptoEngine method
func (ref *Ed25519CryptoEngine) CreateBlockCipher(senderKeyPair *KeyPair, recipientKeyPair *KeyPair) BlockCipher {

	return NewEd25519BlockCipher(senderKeyPair, recipientKeyPair)
}

// CreateKeyAnalyzer implemented interface CryptoEngine method
func (ref *Ed25519CryptoEngine) CreateKeyAnalyzer() KeyAnalyzer {

	return NewEd25519KeyAnalyzer()
}

// Ed25519BlockCipher Implementation of the block cipher for Ed25519.
type Ed25519BlockCipher struct {
	senderKeyPair    *KeyPair
	recipientKeyPair *KeyPair
	keyLength        int
}

func NewEd25519BlockCipher(senderKeyPair *KeyPair, recipientKeyPair *KeyPair) *Ed25519BlockCipher {
	ref := &Ed25519BlockCipher{
		senderKeyPair,
		recipientKeyPair,
		len(recipientKeyPair.PublicKey.Raw),
	}
	return ref
}

//todo: change methods
// today he use java library - I use dummy struct instead
func (ref *Ed25519BlockCipher) setupBlockCipher(sharedKey []byte, ivData []byte, forEncryption bool) *BufferedBlockCipher {

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

func (ref *Ed25519BlockCipher) GetSharedKey(privateKey *PrivateKey, publicKey *PublicKey, salt []byte) ([]byte, error) {

	grA, err := NewEd25519EncodedGroupElement(publicKey.Raw)
	if err != nil {
		return nil, err
	}
	senderA, err := grA.Decode()
	if err != nil {
		return nil, err
	}
	senderA.PrecomputeForScalarMultiplication()
	el, err := senderA.scalarMultiply(PrepareForScalarMultiply(privateKey))
	if err != nil {
		return nil, err
	}
	sharedKey, err := el.Encode()
	if err != nil {
		return nil, err
	}
	for i := 0; i < ref.keyLength; i++ {
		sharedKey.Raw[i] ^= salt[i]
	}

	return HashesSha3_256(sharedKey.Raw)
}

func (ref *Ed25519BlockCipher) Encrypt(input []byte) []byte {

	// Setup salt.
	salt := make([]byte, ref.keyLength)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil
	}
	// Derive shared key.
	sharedKey, err := ref.GetSharedKey(ref.senderKeyPair.PrivateKey, ref.recipientKeyPair.PublicKey, salt)
	if err != nil {
		fmt.Println(err)
	}
	// Setup IV.
	ivData := make([]byte, 16)
	_, err = io.ReadFull(rand.Reader, ivData)
	if err != nil {
		return nil
	}
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

func (ref *Ed25519BlockCipher) Decrypt(input []byte) []byte {

	if len(input) < 64 {
		return nil
	}

	salt := input[:ref.keyLength]
	ivData := input[ref.keyLength:48]
	encData := input[48:]
	// Derive shared key.
	sharedKey, err := ref.GetSharedKey(ref.recipientKeyPair.PrivateKey, ref.senderKeyPair.PublicKey, salt)
	if err != nil {
		fmt.Println(err)
	}
	// Setup block cipher.
	cipher := ref.setupBlockCipher(sharedKey, ivData, false)
	// Decode.
	return ref.transform(cipher, encData)
}

func (ref *Ed25519BlockCipher) transform(cipher *BufferedBlockCipher, data []byte) []byte {

	buf := make([]byte, cipher.GetOutputSize(len(data)))
	length := cipher.processBytes(data, 0, len(data), buf, 0)
	length += cipher.doFinal(buf, length)

	return buf
}

// Ed25519DsaSigner implement DSasigned interface with Ed25519 algo
type Ed25519DsaSigner struct {
	KeyPair *KeyPair
}

//NewEd25519DsaSigner creates a Ed25519 DSA signer.
func NewEd25519DsaSigner(keyPair *KeyPair) *Ed25519DsaSigner {
	return &Ed25519DsaSigner{keyPair}
}

func (ref *Ed25519DsaSigner) Sign(mess []byte) (*Signature, error) {

	if !ref.KeyPair.HasPrivateKey() {
		return nil, errors.New("cannot sign without private key")
	}

	// Hash the private key to improve randomness.
	hash, err := HashesSha3_512(ref.KeyPair.PrivateKey.Raw)
	if err != nil {
		return nil, err
	}
	// r = H(hash_b,...,hash_2b-1, data) where b=256.
	hashR, err := HashesSha3_512(
		hash[32:], // only include the last 32 bytes of the private key hash
		mess)
	if err != nil {
		return nil, err
	}
	r, err := NewEd25519EncodedFieldElement(hashR)
	if err != nil {
		return nil, err
	}
	// Reduce size of r since we are calculating mod group order anyway
	rModQ := r.modQ()
	// R = rModQ * base point.
	R, err := Ed25519Group.BASE_POINT().scalarMultiply(rModQ)
	if err != nil {
		return nil, err
	}
	encodedR, err := R.Encode()
	if err != nil {
		return nil, err
	}
	// S = (r + H(encodedR, encodedA, data) * a) mod group order where
	// encodedR and encodedA are the little endian encodings of the group element R and the public key A and
	// a is the lower 32 bytes of hash after clamping.
	hashH, err := HashesSha3_512(
		encodedR.Raw,
		ref.KeyPair.PublicKey.Raw,
		mess)
	if err != nil {
		return nil, err
	}
	h, err := NewEd25519EncodedFieldElement(hashH)
	if err != nil {
		return nil, err
	}
	hModQ := h.modQ()
	encodedS := hModQ.multiplyAndAddModQ(PrepareForScalarMultiply(ref.KeyPair.PrivateKey),
		rModQ)
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

// Verify reports whether sig is a valid signature of message 'data' by publicKey. It
// prevent  panic inside ed25519.Verify
func (ref *Ed25519DsaSigner) Verify(mess []byte, signature *Signature) (res bool) {

	if !ref.IsCanonicalSignature(signature) {
		return false
	}

	if isEqualConstantTime(ref.KeyPair.PublicKey.Raw, make([]byte, 32)) {
		return false
	}

	// h = H(encodedR, encodedA, data).
	rawEncodedR := signature.R
	rawEncodedA := ref.KeyPair.PublicKey.Raw
	hashR, err := HashesSha3_512(
		rawEncodedR,
		rawEncodedA,
		mess)
	if err != nil {
		panic(err)
		return false
	}
	h, err := NewEd25519EncodedFieldElement(hashR)
	if err != nil {
		panic(err)
		return false
	}
	// hReduced = h mod group order
	hModQ := h.modQ()
	// Must compute A.
	A, err := (&Ed25519EncodedGroupElement{rawEncodedA}).Decode()
	if err != nil {
		fmt.Println(err)
		return false
	}
	A.PrecomputeForDoubleScalarMultiplication()
	// R = encodedS * B - H(encodedR, encodedA, data) * A
	calculatedR, err := Ed25519Group.BASE_POINT().doubleScalarMultiplyVariableTime(
		A,
		hModQ,
		&Ed25519EncodedFieldElement{Ed25519Field_ZERO_SHORT(), signature.S})
	if err != nil {
		panic(err)
		return false
	}
	// Compare calculated R to given R.
	encodedCalculatedR, err := calculatedR.Encode()
	if err != nil {
		panic(err)
		return false
	}

	return isEqualConstantTime(encodedCalculatedR.Raw, rawEncodedR)
}

func (ref *Ed25519DsaSigner) IsCanonicalSignature(signature *Signature) bool {

	sgnS := signature.GetS().Uint64()
	return sgnS != Ed25519Group.GROUP_ORDER.Uint64() && sgnS > 0
}

func (ref *Ed25519DsaSigner) MakeSignatureCanonical(signature *Signature) (*Signature, error) {

	sign := make([]byte, 64)
	copy(sign, signature.S)
	s, err := NewEd25519EncodedFieldElement(sign)
	if err != nil {
		return nil, err
	}
	sModQ := s.modQ()
	return NewSignature(signature.R, sModQ.Raw)
}

//ed25519Curve Class that wraps the elliptic curve Ed25519.
type ed25519Curve struct {
}

var Ed25519Curve = &ed25519Curve{}

func (ref *ed25519Curve) GetName() string {
	return "Ed25519Curve"
}
func (ref *ed25519Curve) GetGroupOrder() *big.Int {
	return Ed25519Group.GROUP_ORDER
}

func (ref *ed25519Curve) GetHalfGroupOrder() uint64 {
	return Ed25519Group.GROUP_ORDER.Uint64() >> 1
}

//Ed25519KeyGenerator Implementation of the key generator for Ed25519.
type Ed25519KeyGenerator struct {
}

func NewEd25519KeyGenerator() *Ed25519KeyGenerator {
	return &Ed25519KeyGenerator{}
}

// GenerateKeyPair generate key pair use ed25519.GenerateKey
func (ref *Ed25519KeyGenerator) GenerateKeyPair() (*KeyPair, error) {

	seed := make([]byte, 32)
	rand := rand2.Reader
	_, err := io.ReadFull(rand, seed[:])
	if err != nil {
		return nil, err
	} // seed is the private key.

	// seed is the private key.
	privateKey := NewPrivateKey(seed)
	publicKey := ref.DerivePublicKey(privateKey)
	return NewKeyPair(privateKey, publicKey, CryptoEngines.Ed25519Engine)
}

func (ref *Ed25519KeyGenerator) DerivePublicKey(privateKey *PrivateKey) *PublicKey {

	a := PrepareForScalarMultiply(privateKey)
	// a * base point is the public key.
	pubKey, err := Ed25519Group.BASE_POINT().scalarMultiply(a)
	if err != nil {
		panic(err)
	}
	el, _ := pubKey.Encode()
	return NewPublicKey(el.Raw)
}
