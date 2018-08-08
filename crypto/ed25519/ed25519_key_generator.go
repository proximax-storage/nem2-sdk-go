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
package ed25519 /*  {packageName}  */
import "github.com/proximax/nem2-go-sdk/sdk/core/crypto" //*"
import io.nem.core.crypto. {package ed25519 /* Name} .arithmetic.Ed25519EncodedFieldElement */
import io.nem.core.crypto. {package ed25519 /* Name} .arithmetic.Ed25519Group */
import io.nem.core.crypto. {package ed25519 /* Name} .arithmetic.Ed25519GroupElement */
import io.nem.core.utils.ArrayUtils 
import java.security.SecureRandom 
/**
 * Implementation of the key generator for Ed25519.
 */
type Ed25519KeyGenerator struct { /* public  */  
    KeyGenerator /* implements */ 
  
    random SecureRandom // private final
} /* Ed25519KeyGenerator */ 
func NewEd25519KeyGenerator () *Ed25519KeyGenerator {  /* public  */ 
    ref := &Ed25519KeyGenerator{
        NewSecureRandom(),
}
    return ref
}

// @Override
   func (ref *Ed25519KeyGenerator) GenerateKeyPair() KeyPair  { /* public  */  

        seed = new byte[32] []byte // final
        ref.random.nextBytes(seed) 
        // seed is the private key.
        privateKey = NewPrivateKey(ArrayUtils.toBigInteger(seed)) PrivateKey // final
        return NewKeyPair(privateKey, CryptoEngines.ed25519Engine()) 
}

// @Override
   func (ref *Ed25519KeyGenerator) DerivePublicKey(PrivateKey privateKey final) PublicKey { /* public  */  

        a = Ed25519Utils.prepareForScalarMultiply(privateKey) Ed25519EncodedFieldElement // final
       // a * base point is the public key.
        pubKey = Ed25519Group.BASE_POINT.scalarMultiply(a) Ed25519GroupElement // final
        // verification of signatures will be about twice as fast when pre-calculating
        // a suitable table of group elements.
        return NewPublicKey(pubKey.encode().getRaw()) 
}

}

