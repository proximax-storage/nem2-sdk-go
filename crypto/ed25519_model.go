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
func (ref *Ed25519CryptoEngine) CreateDsaSigner(keyPair *KeyPair) DsaSigner {

	return NewEd25519DsaSigner(keyPair)
}

// @Override
func (ref *Ed25519CryptoEngine) CreateKeyGenerator() KeyGenerator {

	return NewEd25519KeyGenerator()
}

// @Override
func (ref *Ed25519CryptoEngine) createBlockCipher(senderKeyPair *KeyPair, recipientKeyPair *KeyPair) BlockCipher {

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

	senderKeyPair *KeyPair // private 
	recipientKeyPair *KeyPair // private 
	random *SecureRandom // private 
	keyLength int // private 
} /* Ed25519BlockCipher */
func NewEd25519BlockCipher (senderKeyPair *KeyPair , recipientKeyPair *KeyPair ) *Ed25519BlockCipher {  /* public  */
	ref := &Ed25519BlockCipher{
		senderKeyPair,
		recipientKeyPair,
		NewSecureRandom(),
		len(recipientKeyPair.publicKey.Raw),
	}
	return ref
}

// @Override
func (ref *Ed25519BlockCipher) Encrypt(input []byte) []byte { /* public  */

	// Setup salt.
	salt := make([] byte, ref.keyLength)
	ref.random.nextBytes(salt)
	// Derive shared key.
	sharedKey := ref.GetSharedKey(ref.senderKeyPair.privateKey, ref.recipientKeyPair.publicKey, salt)
	// Setup IV.
	ivData := new byte[16] []byte //
	ref.random.nextBytes(ivData)
	// Setup block cipher.
	cipher = ref.setupBlockCipher(sharedKey, ivData, true) BufferedBlockCipher // 
	// Encode.
	buf = ref.transform(cipher, input) []byte // 
	if (nil == buf) {
		return nil
	}

	result = new byte[salt.length + ivData.length + buf.length] []byte // 
	System.arraycopy(salt, 0, result, 0, salt.length)
	System.arraycopy(ivData, 0, result, salt.length, ivData.length)
	System.arraycopy(buf, 0, result, salt.length + ivData.length, buf.length)
	return result
}

// @Override
func (ref *Ed25519BlockCipher) Decrypt([]byte input ) []byte { /* public  */

	if (input.length < 64) {
		return nil
	}

	salt = Arrays.copyOfRange(input, 0, ref.keyLength) []byte // 
	ivData = Arrays.copyOfRange(input, ref.keyLength, 48) []byte // 
	encData = Arrays.copyOfRange(input, 48, input.length) []byte // 
	// Derive shared key.
	sharedKey = ref.getSharedKey(ref.recipientKeyPair.getPrivateKey(), ref.senderKeyPair.getPublicKey(), salt) []byte // 
	// Setup block cipher.
	cipher = ref.setupBlockCipher(sharedKey, ivData, false) BufferedBlockCipher // 
	// Decode.
	return ref.transform(cipher, encData)
}

func (ref *Ed25519BlockCipher) transform( BufferedBlockCipher cipher,  []byte data) []byte { /* private  */

	buf = byte[cipher.getOutputSize(data.length)] := make([]byte, 0) // 
	int length = cipher.processBytes(data, 0, data.length, buf, 0)
	defer func() {}// try {
	length += cipher.doFinal(buf, length)
} defer func() {}// catch ( InvalidCipherTextException e) {
return nil
}

return Arrays.copyOf(buf, length)
}

func (ref *Ed25519BlockCipher) setupBlockCipher( []byte sharedKey,  []byte ivData,  bool forEncryption) BufferedBlockCipher { /* private  */

	// Setup cipher parameters with key and IV.
	keyParam = NewKeyParameter(sharedKey) KeyParameter // 
	params = NewParametersWithIV(keyParam, ivData) CipherParameters // 
	// Setup AES cipher in CBC mode with PKCS7 padding.
	padding = NewPKCS7Padding() BlockCipherPadding // 
	cipher = NewPaddedBufferedBlockCipher(NewCBCBlockCipher(NewAESEngine()), padding) BufferedBlockCipher // 
	cipher.reset()
	cipher.init(forEncryption, params)
	return cipher
}

func (ref *Ed25519BlockCipher) GetSharedKey( PrivateKey privateKey,  PublicKey publicKey,  []byte salt) []byte { /* private  */

	SenderA := NewEd25519EncodedGroupElement(publicKey.getRaw()).decode() Ed25519GroupElement //
	senderA.precomputeForScalarMultiplication()
	sharedKey = senderA.scalarMultiply(Ed25519Utils.prepareForScalarMultiply(privateKey)).encode().getRaw() []byte // 
	for (int i = 0; i < ref.keyLength; i++) {
		sharedKey[i] ^= salt[i]
	}

	return Hashes.sha3_256(sharedKey)
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
	privateKey := NewPrivateKey(big.Int{}.SetBytes(seed))
	return NewKeyPair(privateKey, nil, CryptoEngines.Ed25519Engine)
}

// @Override
func (ref *Ed25519KeyGenerator) DerivePublicKey(privateKey *PrivateKey) *PublicKey {

	a := PrepareForScalarMultiply(privateKey)
	// a * base point is the public key.
	pubKey := Ed25519Group.BASE_POINT.scalarMultiply(a)
	// verification of signatures will be about twice as fast when pre-calculating
	// a suitable table of group elements.
	return NewPublicKey(pubKey.Encode().Raw)
}
