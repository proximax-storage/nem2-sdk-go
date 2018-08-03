// Contains cryptographic procedures for signing and verifying of signatures
package crypto

import (
	"encoding/base32"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
	)
const NUM_CHECKSUM_BYTES = 4

type KeyPair struct {
	PrivateKey
	PublicKey
}

type PrivateKey struct {
	Value uint64
}

type PublicKey struct {
	Value []byte
}

func GenerateEncodedAddress(pKey string, version uint8) (string, error) {
	// step 1: sha3 hash of the public key
	sha3PublicKeyHash := sha3.New256()
	_, err := sha3PublicKeyHash.Write([]byte(pKey))
	if err == nil {

		// step 2: ripemd160 hash of (1)
		ripemd160StepOneHash := ripemd160.New()
		_, err = ripemd160StepOneHash.Write(sha3PublicKeyHash.Sum(nil))
		if err == nil {

			// step 3: add version byte in front of (2)
			versionPrefixedRipemd160Hash := ripemd160StepOneHash.Sum([]byte{version})

			// step 4: get the checksum of (3)
			sha3StepThreeHash := sha3.New256()
			_, err = sha3StepThreeHash.Write(versionPrefixedRipemd160Hash)
			if err == nil {

				stepThreeChecksum := sha3StepThreeHash.Sum(nil)

				// step 5: concatenate (3) and (4)
				concatStepThreeAndStepSix := append(versionPrefixedRipemd160Hash, stepThreeChecksum[:NUM_CHECKSUM_BYTES]...)

				// step 6: base32 encode (5)
				return base32.HexEncoding.EncodeToString(concatStepThreeAndStepSix), nil
			}
		}
	}
	return "", err
}