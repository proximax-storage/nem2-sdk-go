// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.


package crypto

import "errors"

// KeyPair represent the pair of keys - private & public
type KeyPair struct {
	*PrivateKey
	*PublicKey
}

//NewRandomKeyPair creates a random key pair.
func NewRandomKeyPair() (*KeyPair, error) {
	return NewKeyPairByEngine(CryptoEngines.DefaultEngine)
}

//NewKeyPair The public key is calculated from the private key.
//  The private key must by nil
// if crypto engine is nil - default Engine
func NewKeyPair(privateKey *PrivateKey, publicKey *PublicKey, engine CryptoEngine) (*KeyPair, error) {

	if engine == nil {
		engine = CryptoEngines.DefaultEngine
	}

	if publicKey == nil {
		publicKey = engine.CreateKeyGenerator().DerivePublicKey(privateKey)
	} else if !engine.CreateKeyAnalyzer().IsKeyCompressed(publicKey) {
		return nil, errors.New("publicKey must be in compressed form")
	}
	return &KeyPair{privateKey, publicKey}, nil
}

//NewKeyPairByEngine creates a random key pair that is compatible with the specified engine.
func NewKeyPairByEngine(engine CryptoEngine) (*KeyPair, error) {
	return engine.CreateKeyGenerator().GenerateKeyPair()
}

//HasPrivateKey Determines if the current key pair has a private key.
func (ref *KeyPair) HasPrivateKey() bool {

	return ref.PrivateKey != nil
}
