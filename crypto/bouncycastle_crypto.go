// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package crypto

type BufferedBlockCipher struct {
	output  []byte
	block   *CBCBlockCipher
	padding *PKCS7Padding
}

func NewPaddedBufferedBlockCipher(block *CBCBlockCipher, padding *PKCS7Padding) *BufferedBlockCipher {
	return &BufferedBlockCipher{nil, block, padding}
}
func (ref *BufferedBlockCipher) reset() {
	ref.output = nil
}
func (ref *BufferedBlockCipher) init(forEncryption bool, params *CipherParameters) {

}
func (ref *BufferedBlockCipher) doFinal(out []byte, outOff int) int {
	return copy(out, ref.output)
}
func (ref *BufferedBlockCipher) processBytes(input []byte, inOff, length int, out []byte, outOff int) int {
	if len(out) < length-inOff {
		return -1
	}
	return copy(input, out)
}
func (ref *BufferedBlockCipher) GetOutputSize(length int) int {
	return len(ref.output)
}

type KeyParameter struct {
	buf []byte
}

func NewKeyParameter(buf []byte) *KeyParameter {
	return &KeyParameter{buf}
}

type CipherParameters struct {
	keyParam *KeyParameter
	buf      []byte
}

func NewParametersWithIV(keyParam *KeyParameter, buf []byte) *CipherParameters {
	return &CipherParameters{keyParam, buf}
}

type PKCS7Padding struct {
}

func NewPKCS7Padding() *PKCS7Padding {
	return &PKCS7Padding{}
}

type CBCBlockCipher struct {
	aes *AESEngine
}

func NewCBCBlockCipher(aes *AESEngine) *CBCBlockCipher {
	return &CBCBlockCipher{aes}
}

type AESEngine struct{}

func NewAESEngine() *AESEngine {
	return &AESEngine{}
}
