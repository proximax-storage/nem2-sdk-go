package crypto

import (
	"encoding/hex"
	"math/big"
	"strconv"
)

//KeyAnalyzer Interface to analyze keys.
type KeyAnalyzer interface {

	// Gets a Value indicating whether or not the public key is compressed.
	IsKeyCompressed(publicKey *PublicKey) bool
}

//KeyGenerator Interface for generating keys.
type KeyGenerator interface {
	// Creates a random key pair.
	GenerateKeyPair() (*KeyPair, error)
	// Derives a public key from a private key.
	DerivePublicKey(privateKey *PrivateKey) *PublicKey
}

//PrivateKey Represents a private key.
type PrivateKey struct {
	// I have kept this field for compatibility
	value *big.Int
	Raw   []byte
}

// NewPrivateKey creates a new private key from []byte
func NewPrivateKey(raw []byte) *PrivateKey {
	return &PrivateKey{(&big.Int{}).SetBytes(raw), raw}
}

// NewPrivateKey creates a new private key from []byte
func NewPrivateKeyfromBigInt(val *big.Int) *PrivateKey {
	return &PrivateKey{val, val.Bytes()}
}

//PrivatKeyfromHexString creates a private key from a hex strings.
func NewPrivatKeyfromHexString(sHex string) (*PrivateKey, error) {
	raw, err := hexDecodeString(sHex)
	if err != nil {
		return nil, err
	}

	return NewPrivateKey(raw), nil
}

//PrivateKeyfromDecimalString creates a private key from a decimal strings.
func NewPrivateKeyfromDecimalString(decimal string) (*PrivateKey, error) {
	u, err := strconv.ParseInt(decimal, 10, 64)
	if err != nil {
		return nil, err
	}

	return NewPrivateKeyfromBigInt(big.NewInt(u)), nil

}

func (ref *PrivateKey) String() string {

	return string(ref.Raw)
}

//PublicKey  Represents a public key.
type PublicKey struct {
	Raw []byte
}

//NewPublicKey creates a new public key.
func NewPublicKey(raw []byte) *PublicKey {
	return &PublicKey{raw}
}

func NewPublicKeyfromHex(hStr string) (*PublicKey, error) {
	raw, err := hexDecodeString(hStr)
	if err != nil {
		return nil, err
	}

	return &PublicKey{raw}, nil
}

// Creates a public key from a hex strings.
func (ref *PublicKey) hex() string {
	return string(hex.EncodeToString(ref.Raw))
}

func (ref *PublicKey) String() string {

	return hex.EncodeToString(ref.Raw)
}
