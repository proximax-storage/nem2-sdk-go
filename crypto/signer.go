package crypto

/**
 * Wraps DSA signing and verification logic.
 */
type Signer struct {
	signer DsaSigner // private
	/**
	 * Creates a signer around a KeyPair.
	 *
	 * @param keyPair The KeyPair that should be used for signing and verification.
	 */
} /* Signer */
/**
 * Creates a signer around a KeyPair.
 *
 * @param keyPair The KeyPair that should be used for signing and verification.
 * @param engine  The crypto engine.
 */
func NewSigner(keyPair *KeyPair, engine CryptoEngine) *Signer {
	if engine == nil {
		engine = CryptoEngines.DefaultEngine
	}
	ref := &Signer{
		engine.createDsaSigner(keyPair)}
	return ref
}

/**
 * Creates a signer around a DsaSigner.
 *
 * @param signer The signer.
 */
func NewSignerFromDsaSigner(signer DsaSigner) *Signer {
	ref := &Signer{
		signer,
	}
	return ref
}

// @Override
func (ref *Signer) Sign(data []byte) Signature {

	return ref.signer.sign(data)
}

// @Override
func (ref *Signer) Verify(data []byte, signature Signature) bool {

	return ref.signer.verify(data, signature)
}

// @Override
func (ref *Signer) IsCanonicalSignature(signature Signature) bool {

	return ref.signer.isCanonicalSignature(signature)
}

// @Override
func (ref *Signer) MakeSignatureCanonical(signature Signature) Signature {

	return ref.signer.makeSignatureCanonical(signature)
}
