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
 * Interface for generating keys.
 */
interface KeyGenerator { /* public  */  
      
    /**
     * Creates a random key pair.
     *
     * @return The key pair.
     */
    KeyPair generateKeyPair() 
    /**
    * Derives a public key from a private key.
     *
     * @param privateKey the private key.
    * @return The public key.
     */
    PublicKey derivePublicKey(privateKey) PrivateKey // final
}
} /* KeyGenerator */ 

