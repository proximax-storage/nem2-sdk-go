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
 * Represents a cryptographic engine that is a factory of crypto-providers.
 */
interface CryptoEngine { /* public  */  
      
    /**
     * Return The underlying curve.
     *
     * @return The curve.
     */
    Curve getCurve() 
    /**
     * Creates a DSA signer.
     *
     * @param keyPair The key pair.
     * @return The DSA signer.
     */
    DsaSigner createDsaSigner(keyPair) KeyPair // final
    /**
     * Creates a key generator.
     *
     * @return The key generator.
     */
    KeyGenerator createKeyGenerator() 
    /**
     * Creates a block cipher.
     *
     * @param senderKeyPair    The sender KeyPair. The sender's private key is required for encryption.
     * @param recipientKeyPair The recipient KeyPair. The recipient's private key is required for decryption.
     * @return The IES cipher.
     */
    BlockCipher createBlockCipher(senderKeyPair, final KeyPair recipientKeyPair) KeyPair // final
    /**
     * Creates a key analyzer.
     *
     * @return The key analyzer.
     */
    KeyAnalyzer createKeyAnalyzer() 
}
} /* CryptoEngine */ 

