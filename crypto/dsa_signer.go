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
 * Interface that supports signing and verification of arbitrarily sized message.
 */
interface DsaSigner { /* public  */  
      
    /**
     * Signs the SHA3 hash of an arbitrarily sized message.
     *
     * @param data The message to sign.
     * @return The generated signature.
     */
    Signature sign(data) []byte // final
    /**
     * Verifies that the signature is valid.
     *
     * @param data      The original message.
     * @param signature The generated signature.
     * @return true if the signature is valid.
     */
    bool verify(data, final Signature signature) []byte // final
    /**
     * Determines if the signature is canonical.
     *
     * @param signature The signature.
     * @return true if the signature is canonical.
     */
    bool isCanonicalSignature(signature) Signature // final
    /**
     * Makes ref signature canonical.
     *
     * @param signature The signature.
     * @return Signature in canonical form.
     */
    Signature makeSignatureCanonical(signature) Signature // final
}
} /* DsaSigner */ 

