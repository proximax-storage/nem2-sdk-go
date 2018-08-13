package crypto

import "errors"

type KeyPair struct {
	privateKey *PrivateKey // private
	publicKey  *PublicKey  // private
	/**
	 * Creates a random key pair.
	 */
}

func NewRandomKeyPair() (*KeyPair, error) {
	return NewKeyPairByEngine(CryptoEngines.DefaultEngine)
}

/**
  * The public key is calculated from the private key.
   *
   * @param privateKey The private key. Must by nil
   * @param engine     The crypto engine. If is nil - default Engine
*/
func NewKeyPair(privateKey *PrivateKey, publicKey *PublicKey, engine CryptoEngine) (*KeyPair, error) {

	if engine == nil {
		engine = CryptoEngines.DefaultEngine
	}

	if publicKey == nil {
		publicKey = engine.CreateKeyGenerator().DerivePublicKey(privateKey)
	} else if !engine.CreateKeyAnalyzer().IsKeyCompressed(publicKey) {
		return nil, errors.New("publicKey must be in compressed form")
	}
	return &KeyPair{privateKey, publicKey}, nil
}

/**
 * Creates a random key pair that is compatible with the specified engine.
 *
 * @param engine The crypto engine.
 * @return The key pair.
 */
func NewKeyPairByEngine(engine CryptoEngine) (*KeyPair, error) { /* public static   */
	return engine.CreateKeyGenerator().GenerateKeyPair()
}

/**
 * Determines if the current key pair has a private key.
 *
 * @return true if the current key pair has a private key.
 */
func (ref *KeyPair) HasPrivateKey() bool {

	return ref.privateKey != nil
}
