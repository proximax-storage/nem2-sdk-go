// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.


package crypto

//Cipher Wraps IES encryption and decryption logic.
type Cipher struct {
	cipher BlockCipher
}

// NewCipher creates a cipher around a sender KeyPair and recipient KeyPair.
// if engine not present - use CryptoEngines.DefaultEngine insend
// The sender KeyPair. The sender'S private key is required for encryption.
// The recipient KeyPair. The recipient'S private key is required for decryption.
func NewCipher(senderKeyPair *KeyPair, recipientKeyPair *KeyPair, engine CryptoEngine) *Cipher {
	if engine == nil {
		engine = CryptoEngines.DefaultEngine
	}
	ref := &Cipher{
		engine.CreateBlockCipher(senderKeyPair, recipientKeyPair),
	}
	return ref
}

//NewCipherFromCipher creates a cipher around a cipher.
func NewCipherFromCipher(cipher BlockCipher) *Cipher {
	return &Cipher{
		cipher,
	}
}

func (ref *Cipher) Encrypt(input []byte) []byte {

	return ref.cipher.Encrypt(input)
}

func (ref *Cipher) Decrypt(input []byte) []byte {

	return ref.cipher.Decrypt(input)
}

// BlockCipher Interface for encryption and decryption of data.
type BlockCipher interface {
	// Encrypts an arbitrarily-sized message (input).
	Encrypt(input []byte) []byte
	// Decrypts an arbitrarily-sized message.
	//return The decrypted message or nil if decryption failed.
	Decrypt(input []byte) []byte
}
