package crypto

// Signer Wraps DSA signing and verification logic.
type Signer struct {
	signer DsaSigner
}

// NewSigner creates a signer around a KeyPair.
func NewSigner(keyPair *KeyPair) *Signer {
	return NewSignerByEngine(keyPair, CryptoEngines.DefaultEngine)
}

// NewSignerByEngine creates a signer around a KeyPair.
func NewSignerByEngine(keyPair *KeyPair, engine CryptoEngine) *Signer {
	if engine == nil {
		engine = CryptoEngines.DefaultEngine
	}
	ref := &Signer{
		engine.CreateDsaSigner(keyPair)}
	return ref
}

// NewSignerFromDsaSigner Creates a signer around a DsaSigner.
func NewSignerFromDsaSigner(signer DsaSigner) *Signer {
	ref := &Signer{
		signer,
	}
	return ref
}

func (ref *Signer) Sign(data []byte) (*Signature, error) {

	return ref.signer.Sign(data)
}

func (ref *Signer) Verify(data []byte, signature *Signature) bool {

	return ref.signer.Verify(data, signature)
}

func (ref *Signer) IsCanonicalSignature(signature *Signature) bool {

	return ref.signer.IsCanonicalSignature(signature)
}

func (ref *Signer) MakeSignatureCanonical(signature *Signature) (*Signature, error) {

	return ref.signer.MakeSignatureCanonical(signature)
}
