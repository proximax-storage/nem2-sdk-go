package crypto

/**
 * Interface that supports signing and verification of arbitrarily sized message.
 */
type DsaSigner interface {

	/**
	 * Signs the SHA3 hash of an arbitrarily sized message.
	 *
	 * @param data The message to sign.
	 * @return The generated signature.
	 */
	Sign(data []byte) (*Signature, error)
	/**
	 * Verifies that the signature is valid.
	 *
	 * @param data      The original message.
	 * @param signature The generated signature.
	 * @return true if the signature is valid.
	 */
	Verify(data []byte, signature *Signature) bool
	/**
	 * Determines if the signature is canonical.
	 *
	 * @param signature The signature.
	 * @return true if the signature is canonical.
	 */
	IsCanonicalSignature(signature *Signature) bool
	/**
	 * Makes ref signature canonical.
	 *
	 * @param signature The signature.
	 * @return Signature in canonical form.
	 */
	MakeSignatureCanonical(signature *Signature) *Signature
}
