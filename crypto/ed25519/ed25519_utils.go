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
import "github.com/proximax/nem2-go-sdk/sdk/core/crypto" //hashes"
import "github.com/proximax/nem2-go-sdk/sdk/core/crypto" //privatekey"
import io.nem.core.crypto. {package ed25519 /* Name} .arithmetic.Ed25519EncodedFieldElement */
// import java.util.Arrays 
/**
 * Utility methods for Ed25519.
 */
type Ed25519Utils struct { /* public  */  
      
    /**
     * Prepares a private key's raw value for scalar multiplication.
     * The hashing is for achieving better randomness and the clamping prevents small subgroup attacks.
     *
     * @param key The private key.
     * @return The prepared encoded field element.
     */
   func (ref *Ed25519Utils) PrepareForScalarMultiply(PrivateKey key final) Ed25519EncodedFieldElement { /* public static  */  

        hash = Hashes.sha3_512(key.getBytes()) []byte // final
        a = Arrays.copyOfRange(hash, 0, 32) []byte // final
        a[31] &= 0x7F 
        a[31] |= 0x40 
        a[0] &= 0xF8 
        return NewEd25519EncodedFieldElement(a) 
}
} /* Ed25519Utils */ 

}

