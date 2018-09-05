package crypto

//DsaSigner Interface that supports signing and verification of arbitrarily sized message.
type DsaSigner interface {

	// Signs the SHA3 hash of an arbitrarily sized message.
	Sign(mess []byte) (*Signature, error)
	// Verifies that the signature is valid.
	Verify(mess []byte, signature *Signature) bool
	// Determines if the signature is canonical.
	IsCanonicalSignature(signature *Signature) bool
	// Makes ref signature canonical.
	MakeSignatureCanonical(signature *Signature) (*Signature, error)
}
