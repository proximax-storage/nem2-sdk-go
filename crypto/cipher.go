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
 * Wraps IES encryption and decryption logic.
 */
type Cipher struct { /* public  */  
    BlockCipher /* implements */ 
  
    cipher BlockCipher // private final
    /**
     * Creates a cipher around a sender KeyPair and recipient KeyPair.
     *
     * @param senderKeyPair    The sender KeyPair. The sender's private key is required for encryption.
     * @param recipientKeyPair The recipient KeyPair. The recipient's private key is required for decryption.
     */
} /* Cipher */ 
func NewCipher (final KeyPair senderKeyPair, final KeyPair recipientKeyPair) *Cipher {  /* public  */ 
    ref := &Cipher{
        senderKeyPair, recipientKeyPair, CryptoEngines.defaultEngine(), 
}
    return ref
}

    /**
     * Creates a cipher around a sender KeyPair and recipient KeyPair.
     *
     * @param senderKeyPair    The sender KeyPair. The sender's private key is required for encryption.
     * @param recipientKeyPair The recipient KeyPair. The recipient's private key is required for decryption.
     * @param engine           The crypto engine.
     */
func NewCipher (final KeyPair senderKeyPair, final KeyPair recipientKeyPair, final CryptoEngine engine) *Cipher {  /* public  */ 
    ref := &Cipher{
        engine.createBlockCipher(senderKeyPair, recipientKeyPair), 
}
    return ref
}

    /**
     * Creates a cipher around a cipher.
     *
     * @param cipher The cipher.
     */
func NewCipher (BlockCipher cipher final) *Cipher {  /* public  */ 
    ref := &Cipher{
        cipher,
}
    return ref
}

// @Override
   func (ref *Cipher) Encrypt([]byte input final) []byte { /* public  */  

        return ref.cipher.encrypt(input) 
}

// @Override
   func (ref *Cipher) Decrypt([]byte input final) []byte { /* public  */  

        return ref.cipher.decrypt(input) 
}

}

