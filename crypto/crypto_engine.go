// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

 package crypto

//CryptoEngine Represents a cryptographic engine that is a factory of crypto-providers.
type CryptoEngine interface {

	// Return The underlying curve.
	GetCurve() Curve
	// Creates a DSA signer.
	CreateDsaSigner(keyPair *KeyPair) DsaSigner
	// Creates a key generator.
	CreateKeyGenerator() KeyGenerator
	//Creates a block cipher.
	CreateBlockCipher(senderKeyPair *KeyPair, recipientKeyPair *KeyPair) BlockCipher
	// Creates a key analyzer.
	CreateKeyAnalyzer() KeyAnalyzer
}

//cryptoEngines Static class that exposes crypto engines.
type cryptoEngines struct {
	Ed25519Engine *Ed25519CryptoEngine
	DefaultEngine *Ed25519CryptoEngine
}

var CryptoEngines = cryptoEngines{
	&Ed25519CryptoEngine{},
	&Ed25519CryptoEngine{},
}
