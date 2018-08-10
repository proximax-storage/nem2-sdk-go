package ed25519

import "github.com/proximax-storage/nem2-sdk-go/crypto"

/**
 * Class that wraps the Ed25519 specific implementation.
 */
type Ed25519CryptoEngine struct {
}
// @Override
   func (ref *Ed25519CryptoEngine) GetCurve() crypto.ed25519Curve {

        return crypto.ed25519Curve. {package ed25519 /* Name} () */
}
} /* Ed25519CryptoEngine */ 

// @Override
   func (ref *Ed25519CryptoEngine) CreateDsaSigner(KeyPair keyPair ) DsaSigner {

        return crypto.NewEd25519DsaSigner(keyPair)
}

// @Override
   func (ref *Ed25519CryptoEngine) CreateKeyGenerator() *KeyGenerator  {

        return crypto.NewEd25519KeyGenerator()
}

// @Override
   func (ref *Ed25519CryptoEngine) CreateBlockCipher( KeyPair senderKeyPair,  KeyPair recipientKeyPair) BlockCipher {

        return NewEd25519BlockCipher(senderKeyPair, recipientKeyPair) 
}

// @Override
   func (ref *Ed25519CryptoEngine) CreateKeyAnalyzer() KeyAnalyzer  {

        return NewEd25519KeyAnalyzer() 
}

}

