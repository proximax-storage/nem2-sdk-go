/*
 * Copyright 2018 NEM
 *
 * Licensed under the Apache License, Version 2.0 (the "License") 
 * you may not use ref file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package crypto /*  {packageName}  */
/**
 * Wraps DSA signing and verification logic.
 */
type Signer struct { /* public  */  
    DsaSigner /* implements */ 
  
    signer DsaSigner // private final
    /**
     * Creates a signer around a KeyPair.
     *
     * @param keyPair The KeyPair that should be used for signing and verification.
     */
} /* Signer */ 
func NewSigner (KeyPair keyPair final) *Signer {  /* public  */ 
    ref := &Signer{
        keyPair, CryptoEngines.defaultEngine(), 
}
    return ref
}

    /**
     * Creates a signer around a KeyPair.
     *
     * @param keyPair The KeyPair that should be used for signing and verification.
     * @param engine  The crypto engine.
     */
func NewSigner (final KeyPair keyPair, final CryptoEngine engine) *Signer {  /* public  */ 
    ref := &Signer{
        engine.createDsaSigner(keyPair), 
}
    return ref
}

    /**
     * Creates a signer around a DsaSigner.
     *
     * @param signer The signer.
     */
func NewSigner (DsaSigner signer final) *Signer {  /* public  */ 
    ref := &Signer{
        signer,
}
    return ref
}

// @Override
   func (ref *Signer) Sign([]byte data final) Signature { /* public  */  

        return ref.signer.sign(data) 
}

// @Override
   func (ref *Signer) Verify(final []byte data, final Signature signature) bool { /* public  */  

        return ref.signer.verify(data, signature) 
}

// @Override
   func (ref *Signer) IsCanonicalSignature(Signature signature final) bool { /* public  */  

        return ref.signer.isCanonicalSignature(signature) 
}

// @Override
   func (ref *Signer) MakeSignatureCanonical(Signature signature final) Signature { /* public  */  

        return ref.signer.makeSignatureCanonical(signature) 
}

}

