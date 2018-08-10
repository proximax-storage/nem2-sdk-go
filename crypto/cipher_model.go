package crypto

/**
 * Wraps IES encryption and decryption logic.
 */
type Cipher struct {
	cipher BlockCipher // private
	/**
	 * Creates a cipher around a sender KeyPair and recipient KeyPair.
	 *
	 * @param senderKeyPair    The sender KeyPair. The sender'S private key is required for encryption.
	 * @param recipientKeyPair The recipient KeyPair. The recipient'S private key is required for decryption.
	 */
}

/**
 * Creates a cipher around a sender KeyPair and recipient KeyPair.
 *
 * @param senderKeyPair    The sender KeyPair. The sender'S private key is required for encryption.
 * @param recipientKeyPair The recipient KeyPair. The recipient'S private key is required for decryption.
 * @param engine           The crypto engine.
 */
func NewCipher(senderKeyPair *KeyPair, recipientKeyPair *KeyPair, engine CryptoEngine) *Cipher {
	if engine == nil {
		engine = CryptoEngines.DefaultEngine
	}
	ref := &Cipher{
		engine.createBlockCipher(senderKeyPair, recipientKeyPair),
	}
	return ref
}

/**
 * Creates a cipher around a cipher.
 *
 * @param cipher The cipher.
 */
func NewCipherFromCipher(cipher BlockCipher) *Cipher {
	ref := &Cipher{
		cipher,
	}
	return ref
}

// @Override
func (ref *Cipher) Encrypt(input []byte) []byte {

	return ref.cipher.Encrypt(input)
}

// @Override
func (ref *Cipher) Decrypt(input []byte) []byte {

	return ref.cipher.Decrypt(input)
}

/**
 * BlockCipher Interface for encryption and decryption of data.
 */type BlockCipher interface {
	/**
	 * Encrypts an arbitrarily-sized message.
	 *
	 * @param input The message to encrypt.
	 * @return The encrypted message.
	 */
	Encrypt(input []byte) []byte
	/**
	 * Decrypts an arbitrarily-sized message.
	 *
	 * @param input The message to decrypt.
	 * @return The decrypted message or nil if decryption failed.
	 */
	Decrypt(input []byte) []byte
}
