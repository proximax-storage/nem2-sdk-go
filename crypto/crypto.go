// Contains cryptographic procedures for signing and verifying of signatures
package crypto

import (
	"encoding/hex"
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
func hexDecodeString(str string) ([]byte, error) {
	return hexDecode([]byte(str))
}
func hexDecode(src []byte) ([]byte, error) {
	if len(src)%2 != 0 {
		src = append([]byte{'0'}, src...)
	}
	dst := make([]byte, len(src))
	_, err := hex.Decode(dst, src)
	if err != nil {
		return nil, err
	}
	return dst, nil
}

func isNegativeConstantTime(b int) int {
	return (b >> 8) & 1
}

func IsEqualConstantTime(b, c int) int {

	result := 0
	xor := b ^ c // final
	for i := uint(0); i < 8; i++ {
		result |= xor >> i
	}

	return (result ^ 0x01) & 0x01
}
