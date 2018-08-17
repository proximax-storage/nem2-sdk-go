// Copyright 2017 Author: Ruslan Bikchentaev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crypto

import (
	"bytes"
	"crypto/rand"
	"errors"
	"fmt"
	"golang.org/x/crypto/curve25519"
	"golang.org/x/crypto/ed25519"
	"io"
	"math/big"
)

//Ed25519CryptoEngine
type Ed25519CryptoEngine struct {
}

func (ref *Ed25519CryptoEngine) GetCurve() Curve {

	return Ed25519Curve

}

func (ref *Ed25519CryptoEngine) CreateDsaSigner(keyPair *KeyPair) DsaSigner {

	return NewEd25519DsaSigner(keyPair)
}

func (ref *Ed25519CryptoEngine) CreateKeyGenerator() KeyGenerator {

	return NewEd25519KeyGenerator()
}

func (ref *Ed25519CryptoEngine) CreateBlockCipher(senderKeyPair *KeyPair, recipientKeyPair *KeyPair) BlockCipher {

	return NewEd25519BlockCipher(senderKeyPair, recipientKeyPair)
}

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
		len(recipientKeyPair.publicKey.Raw),
	}
	return ref
}

//todo: change methods
// today he use java library - I use dummy struct instead
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

func (ref *Ed25519BlockCipher) Encrypt(input []byte) []byte {

	// Setup salt.
	salt := make([]byte, ref.keyLength)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil
	}
	// Derive shared key.
	sharedKey, err := ref.GetSharedKey(ref.senderKeyPair.privateKey, ref.recipientKeyPair.publicKey, salt)
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

// Ed25519DsaSigner implement DSasigned interface with Ed25519 algo
type Ed25519DsaSigner struct {
	KeyPair *KeyPair
}

//NewEd25519DsaSigner Creates a Ed25519 DSA signer.
func NewEd25519DsaSigner(keyPair *KeyPair) *Ed25519DsaSigner {
	return &Ed25519DsaSigner{
		keyPair,
	}
}

func (ref *Ed25519DsaSigner) Sign(data []byte) (*Signature, error) {

	if !ref.KeyPair.HasPrivateKey() {
		return nil, errors.New("cannot sign without private key")
	}

	// Signature is (encodedR, encodedS)
	signature, err := NewSignatureFromBytes(ed25519.Sign(ref.KeyPair.PrivateKey(), data))
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
func (ref *Ed25519DsaSigner) Verify(data []byte, signature *Signature) (res bool) {

	if !ref.IsCanonicalSignature(signature) || (len(ref.KeyPair.PublicKey()) != ed25519.PublicKeySize) {
		return false
	}

	if b := make([]byte, 32); bytes.Equal(ref.KeyPair.PrivateKey(), b) {
		return false
	}

	defer func() {
		err := recover()
		if err != nil {
			fmt.Errorf("%v", err)
			res = false
		}
	}()
	return ed25519.Verify(ref.KeyPair.PublicKey(), data, signature.Bytes())
}

func (ref *Ed25519DsaSigner) IsCanonicalSignature(signature *Signature) bool {

	sgnS := signature.GetS()
	return uint64(sgnS) != Ed25519Group.GROUP_ORDER.Uint64() && sgnS > 0
}

func (ref *Ed25519DsaSigner) MakeSignatureCanonical(signature *Signature) (*Signature, error) {

	s, err := NewEd25519EncodedFieldElement(signature.S[:64])
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

	//ignore publicKey
	_, prKey, err := ed25519.GenerateKey(nil)
	if err != nil {
		return nil, err
	}
	// seed is the private key.
	privateKey := NewPrivateKey(prKey)
	publicKey := ref.DerivePublicKey(privateKey)
	return NewKeyPair(privateKey, publicKey, CryptoEngines.Ed25519Engine)
}

func (ref *Ed25519KeyGenerator) DerivePublicKey(privateKey *PrivateKey) *PublicKey {

	a := PrepareForScalarMultiply(privateKey)
	// a * base point is the public key.
	var pubKey, prKey, base [32]byte
	copy(prKey[:], privateKey.Raw)
	copy(base[:], a.Raw)
	curve25519.ScalarMult(&pubKey, &prKey, &base)
	// verification of signatures will be about twice as fast when pre-calculating
	// a suitable table of group elements.
	return NewPublicKey(pubKey[:])
}
