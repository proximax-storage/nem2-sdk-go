// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package crypto

import (
	"crypto/subtle"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
)

func HashesSha3_256(b []byte) ([]byte, error) {
	hash := sha3.New256()
	_, err := hash.Write(b)
	if err != nil {
		return nil, err
	}

	return hash.Sum(nil), nil
}
func HashesSha3_512(inputs ...[]byte) ([]byte, error) {
	hash := sha3.New512()
	for _, b := range inputs {

		_, err := hash.Write(b)
		if err != nil {
			return nil, err
		}
	}

	return hash.Sum(nil), nil
}
func HashesRipemd160(b []byte) ([]byte, error) {
	hash := ripemd160.New()
	_, err := hash.Write(b)
	if err != nil {
		return nil, err
	}

	return hash.Sum(nil), nil

}

func isNegativeConstantTime(b int) int {
	return (b >> 8) & 1
}

func IsConstantTimeByteEq(b, c int) int {

	result := 0
	xor := b ^ c // final
	for i := uint(0); i < 8; i++ {
		result |= xor >> i
	}

	return (result ^ 0x01) & 0x01
}

func isEqualConstantTime(x, y []byte) bool {
	return subtle.ConstantTimeCompare(x, y) == 1
}
