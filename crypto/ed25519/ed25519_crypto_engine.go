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
/**
 * Class that wraps the Ed25519 specific implementation.
 */
type Ed25519CryptoEngine struct { /* public  */  
    CryptoEngine /* implements */ 
  
// @Override
   func (ref *Ed25519CryptoEngine) GetCurve() Curve  { /* public  */  

        return Ed25519Curve. {package ed25519 /* Name} () */
}
} /* Ed25519CryptoEngine */ 

// @Override
   func (ref *Ed25519CryptoEngine) CreateDsaSigner(KeyPair keyPair final) DsaSigner { /* public  */  

        return NewEd25519DsaSigner(keyPair) 
}

// @Override
   func (ref *Ed25519CryptoEngine) CreateKeyGenerator() KeyGenerator  { /* public  */  

        return NewEd25519KeyGenerator() 
}

// @Override
   func (ref *Ed25519CryptoEngine) CreateBlockCipher(final KeyPair senderKeyPair, final KeyPair recipientKeyPair) BlockCipher { /* public  */  

        return NewEd25519BlockCipher(senderKeyPair, recipientKeyPair) 
}

// @Override
   func (ref *Ed25519CryptoEngine) CreateKeyAnalyzer() KeyAnalyzer  { /* public  */  

        return NewEd25519KeyAnalyzer() 
}

}

