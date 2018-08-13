package crypto

//CryptoEngine Represents a cryptographic engine that is a factory of crypto-providers.
type CryptoEngine interface {

	/**
	 * Return The underlying curve.
	 *
	 * @return The curve.
	 */
	GetCurve() Curve
	/**
	 * Creates a DSA signer.
	 *
	 * @param keyPair The key pair.
	 * @return The DSA signer.
	 */
	CreateDsaSigner(keyPair *KeyPair) DsaSigner
	/**
	 * Creates a key generator.
	 *
	 * @return The key generator.
	 */
	CreateKeyGenerator() KeyGenerator
	/**
	 * Creates a block cipher.
	 *
	 * @param senderKeyPair    The sender KeyPair. The sender'S private key is required for encryption.
	 * @param recipientKeyPair The recipient KeyPair. The recipient'S private key is required for decryption.
	 * @return The IES cipher.
	 */
	CreateBlockCipher(senderKeyPair *KeyPair, recipientKeyPair *KeyPair) BlockCipher
	/**
	 * Creates a key analyzer.
	 *
	 * @return The key analyzer.
	 */
	CreateKeyAnalyzer() KeyAnalyzer
}

//cryptoEngines Static class that exposes crypto engines.
type cryptoEngines struct {
	Ed25519Engine *Ed25519CryptoEngine
	DefaultEngine *Ed25519CryptoEngine
}

var CryptoEngines = cryptoEngines{
	&Ed25519CryptoEngine{},
	&Ed25519CryptoEngine{},
}
