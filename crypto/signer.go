// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package crypto

// Signer wraps DSA signing and verification logic.
type Signer struct {
	signer DsaSigner
}

// NewSignerFromDsaSigner Creates a signer around a DsaSigner.
func NewSigner(signer DsaSigner) *Signer {
	return &Signer{signer}
}

// NewSigner creates a signer around a KeyPair.
func NewSignerFromKeyPair(keyPair *KeyPair, engine CryptoEngine) *Signer {
	if engine == nil {
		engine = CryptoEngines.DefaultEngine
	}
	return NewSigner(engine.CreateDsaSigner(keyPair))
}

// Sign implemented interface DsaSigner method
func (ref *Signer) Sign(data []byte) (*Signature, error) {

	return ref.signer.Sign(data)
}

// Verify implemented interface DsaSigner method
func (ref *Signer) Verify(data []byte, signature *Signature) bool {

	return ref.signer.Verify(data, signature)
}

// IsCanonicalSignature implemented interface DsaSigner method
func (ref *Signer) IsCanonicalSignature(signature *Signature) bool {

	return ref.signer.IsCanonicalSignature(signature)
}

// MakeSignatureCanonical implemented interface DsaSigner method
func (ref *Signer) MakeSignatureCanonical(signature *Signature) (*Signature, error) {

	return ref.signer.MakeSignatureCanonical(signature)
}
