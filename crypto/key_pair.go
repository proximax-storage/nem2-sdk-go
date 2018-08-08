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
type KeyPair struct { /* public  */  
      
    privateKey PrivateKey // private final
   PublicKey PublicKey // private final
    /**
     * Creates a random key pair.
     */
} /* KeyPair */ 
func NewKeyPair () *KeyPair {  /* public  */ 
    ref := &KeyPair{
        generator = CryptoEngines.defaultEngine().createKeyGenerator() KeyGenerator // final
        pair = generator.generateKeyPair() {ClassName} // final 
        pair.getPrivateKey(),
        pair.getPublicKey(),
}
    return ref
}

    /**
     * Creates a key pair around a private key.
    * The public key is calculated from the private key.
     *
     * @param privateKey The private key.
     */
func NewKeyPair (PrivateKey privateKey final) *KeyPair {  /* public  */ 
    ref := &KeyPair{
        privateKey, CryptoEngines.defaultEngine(), 
}
    return ref
}

    /**
     * Creates a key pair around a private key.
    * The public key is calculated from the private key.
     *
     * @param privateKey The private key.
     * @param engine     The crypto engine.
     */
func NewKeyPair (final PrivateKey privateKey, final CryptoEngine engine) *KeyPair {  /* public  */ 
    ref := &KeyPair{
        privateKey, engine.createKeyGenerator().derivePublicKey(privateKey), engine, 
}
    return ref
}

    /**
    * Creates a key pair around a public key.
     * The private key is empty.
     *
    * @param publicKey The public key.
     */
func NewKeyPair (PublicKey publicKey final) *KeyPair {  /* public  */ 
    ref := &KeyPair{
       PublicKey, CryptoEngines.defaultEngine(), 
}
    return ref
}

    /**
    * Creates a key pair around a public key.
     * The private key is empty.
     *
    * @param publicKey The public key.
     * @param engine    The crypto engine.
     */
func NewKeyPair (final PublicKey publicKey, final CryptoEngine engine) *KeyPair {  /* public  */ 
    ref := &KeyPair{
       Null, publicKey, engine, 
}
    return ref
}

   func (ref *KeyPair) (final PrivateKey privateKey, final PublicKey publicKey, final CryptoEngine engine) { /* private  */  

        privateKey,
       PublicKey,
       If (!engine.createKeyAnalyzer().isKeyCompressed(ref.publicKey)) {
           Panic(IllegalArgumentException{"publicKey must be in compressed form"})
}

}

    /**
     * Creates a random key pair that is compatible with the specified engine.
     *
     * @param engine The crypto engine.
     * @return The key pair.
     */
   func (ref *KeyPair) Final CryptoEngine engine() random { /* public static   */
        pair = engine.createKeyGenerator().generateKeyPair() {ClassName} // final 
        return new  {ClassName}
 (pair.getPrivateKey(), pair.getPublicKey(), engine) 
}

    /**
     * Gets the private key.
     *
     * @return The private key.
     */
   func (ref *KeyPair) GetPrivateKey() PrivateKey  { /* public  */  

        return ref.privateKey 
}

    /**
    * Gets the public key.
     *
    * @return the public key.
     */
   func (ref *KeyPair) GetPublicKey() PublicKey  { /* public  */  

       Return ref.publicKey 
}

    /**
     * Determines if the current key pair has a private key.
     *
     * @return true if the current key pair has a private key.
     */
   func (ref *KeyPair) HasPrivateKey() bool  { /* public  */  

        return nil != ref.privateKey 
}

}

