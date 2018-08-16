// Contains cryptographic procedures for signing and verifying of signatures
package crypto

import (
	"encoding/base32"
	"encoding/hex"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
)

const NUM_CHECKSUM_BYTES = 4

// GenerateEncodedAddress convert publicKey to address
func GenerateEncodedAddress(pKey string, version uint8) (string, error) {
	// step 1: sha3 hash of the public key
	pKeyD, err := hexDecodeString(pKey)
	if err != nil {
		return "", err
	}
	sha3PublicKeyHash := sha3.New256()
	_, err = sha3PublicKeyHash.Write(pKeyD)
	if err != nil {
		return "", err
	}

	// step 2: ripemd160 hash of (1)
	ripemd160StepOneHash := ripemd160.New()
	ripemd160StepOneHash.Write(sha3PublicKeyHash.Sum(nil))

	// step 3: add version byte in front of (2)
	versionPrefixedRipemd160Hash := append([]byte{version}, ripemd160StepOneHash.Sum(nil)...)

	// step 4: get the checksum of (3)
	stepThreeChecksum, err := GenerateChecksum(versionPrefixedRipemd160Hash)
	if err != nil {
		return "", err
	}

	// step 5: concatenate (3) and (4)
	concatStepThreeAndStepSix := append(versionPrefixedRipemd160Hash, stepThreeChecksum...)

	// step 6: base32 encode (5)
	return base32.StdEncoding.EncodeToString(concatStepThreeAndStepSix), nil
}

func GenerateChecksum(b []byte) ([]byte, error) {
	// step 1: sha3 hash of (input
	sha3StepThreeHash := sha3.New256()
	_, err := sha3StepThreeHash.Write(b)
	if err != nil {
		return nil, err
	}

	// step 2: get the first X bytes of (1)
	return sha3StepThreeHash.Sum(nil)[:NUM_CHECKSUM_BYTES], nil
}

// todo: need check the three following methods
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
