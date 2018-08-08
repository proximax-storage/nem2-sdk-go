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
 * Interface for encryption and decryption of data.
 */
interface BlockCipher { /* public  */  
      
    /**
     * Encrypts an arbitrarily-sized message.
     *
     * @param input The message to encrypt.
     * @return The encrypted message.
     */
    []byte encrypt(input) []byte // final
    /**
     * Decrypts an arbitrarily-sized message.
     *
     * @param input The message to decrypt.
     * @return The decrypted message or nil if decryption failed.
     */
    []byte decrypt(input) []byte // final
}
} /* BlockCipher */ 

