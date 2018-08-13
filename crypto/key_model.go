package crypto

import (
	"encoding/hex"
	"math/big"
	"strconv"
)

//KeyAnalyzer Interface to analyze keys.
type KeyAnalyzer interface {

	/**
	  * Gets a Value indicating whether or not the public key is compressed.
	   *
	  * @param publicKey The public key.
	  * @return true if the public key is compressed, false otherwise.
	*/
	IsKeyCompressed(publicKey *PublicKey) bool
}

//KeyGenerator Interface for generating keys.
type KeyGenerator interface {
	/**
	 * Creates a random key pair.
	 *
	 * @return The key pair.
	 */
	GenerateKeyPair() (*KeyPair, error)
	/**
	* Derives a public key from a private key.
	 *
	 * @param privateKey the private key.
	* @return The public key.
	*/
	DerivePublicKey(privateKey *PrivateKey) *PublicKey
}

//PrivateKey Represents a private key.
type PrivateKey struct {
	Value *big.Int
	/**
	 * Creates a new private key.
	 *
	 * @param Value The  private key Value.
	 */
}

func NewPrivateKey(value *big.Int) *PrivateKey {
	ref := &PrivateKey{value}
	return ref
}

//PrivatKeyfromHexString creates a private key from a hex strings.
func PrivatKeyfromHexString(sHex string) (*PrivateKey, error) {
	value, err := hex.DecodeString(sHex)
	if err != nil {
		return nil, err
	}

	return &PrivateKey{(&big.Int{}).SetBytes(value)}, nil
}

//PrivateKeyfromDecimalString creates a private key from a decimal strings.
func PrivateKeyfromDecimalString(decimal string) (*PrivateKey, error) {
	u, err := strconv.ParseInt(decimal, 10, 64)
	if err != nil {
		return nil, err
	}
	ref := &PrivateKey{big.NewInt(u)}
	return ref, nil

}

func (ref *PrivateKey) getBytes() []byte {

	return ref.Value.Bytes()
}

func (ref *PrivateKey) String() string {

	return string(ref.getBytes())
}

//PublicKey  Represents a public key.
type PublicKey struct {
	Raw []byte
}

//NewPublicKey creates a new public key.
func NewPublicKey(hex string) *PublicKey {
	ref := &PublicKey{[]byte(hex)}
	return ref
}

// Creates a public key from a hex strings.
func (ref *PublicKey) hex() string {
	return string(ref.Raw)
}

func (ref *PublicKey) String() string {

	return string(ref.Raw)
}
