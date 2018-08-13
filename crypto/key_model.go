package crypto

import (
	"encoding/hex"
	"math/big"
	"strconv"
)

/**
 * Interface to analyze keys.
 */
type KeyAnalyzer interface {

	/**
	  * Gets a Value indicating whether or not the public key is compressed.
	   *
	  * @param publicKey The public key.
	  * @return true if the public key is compressed, false otherwise.
	*/
	IsKeyCompressed(publicKey *PublicKey) bool
}

/**
 * Interface for generating keys.
 */type KeyGenerator interface {
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

/**
 * Represents a private key.
 */type PrivateKey struct {
	Value *big.Int // private
	/**
	 * Creates a new private key.
	 *
	 * @param Value The  private key Value.
	 */
} /* PrivateKey */
func NewPrivateKey(value *big.Int) *PrivateKey {
	ref := &PrivateKey{value}
	return ref
}

/**
 * Creates a private key from a hex strings.
 *
 * @param hex The hex strings.
 * @return The new private key.
 */
func PrivatKeyfromHexString(sHex string) (*PrivateKey, error) { /* public static   */
	value, err := hex.DecodeString(sHex)
	if err != nil {
		return nil, err
	}

	return &PrivateKey{(&big.Int{}).SetBytes(value)}, nil
}

/**
 * Creates a private key from a decimal strings.
 *
 * @param decimal The decimal strings.
 * @return The new private key.
 */func PrivateKeyfromDecimalString(decimal string) (*PrivateKey, error) { /* public static   */
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

// @Override
func (ref *PrivateKey) String() string {

	return string(ref.getBytes())
}

// import java.util.Arrays
/**
* Represents a public key.
 */type PublicKey struct {
	Raw []byte // private
	/**
	* Creates a new public key.
	 *
	* @param bytes The raw public key Value.
	*/
} /* PublicKey */
func NewPublicKey(hex string) *PublicKey {
	ref := &PublicKey{[]byte(hex)}
	return ref
}

/**
  * Creates a public key from a hex strings.
   *
   * @param hex The hex strings.
  * @return The new public key.
*/func (ref *PublicKey) hex() string { /* public static   */
	return string(ref.Raw)
}

func (ref *PublicKey) String() string {

	return string(ref.Raw)
}
