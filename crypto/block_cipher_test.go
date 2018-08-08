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
type BlockCipherTest struct { /* public abstract  */
      
// @Test
   func (ref *BlockCipherTest) EncryptedDataCanBeDecrypted()    { /* public  */  

        // Arrange:
        engine = ref.getCryptoEngine() CryptoEngine // final
        kp = KeyPair.random(engine) KeyPair // final
        blockCipher = ref.getBlockCipher(kp, kp) BlockCipher // final
        input = Utils.generateRandomBytes() []byte // final
        // Act:
        encryptedBytes = blockCipher.encrypt(input) []byte // final
        decryptedBytes = blockCipher.decrypt(encryptedBytes) []byte // final
        // Assert:
        Assert.assertThat(encryptedBytes, IsNot.not(IsEqual.equalTo(decryptedBytes))) 
        Assert.assertThat(decryptedBytes, IsEqual.equalTo(input)) 
}
} /* BlockCipherTest */ 

// @Test
   func (ref *BlockCipherTest) DataCanBeEncryptedWithSenderPrivateKeyAndRecipientPublicKey()    { /* public  */  

        // Arrange:
        engine = ref.getCryptoEngine() CryptoEngine // final
        skp = KeyPair.random(engine) KeyPair // final
        rkp = KeyPair.random(engine) KeyPair // final
        blockCipher = ref.getBlockCipher(skp, NewKeyPair(rkp.getPublicKey(), engine)) BlockCipher // final
        input = Utils.generateRandomBytes() []byte // final
        // Act:
        encryptedBytes = blockCipher.encrypt(input) []byte // final
        // Assert:
        Assert.assertThat(encryptedBytes, IsNot.not(IsEqual.equalTo(input))) 
}

// @Test
   func (ref *BlockCipherTest) DataCanBeDecryptedWithSenderPublicKeyAndRecipientPrivateKey()    { /* public  */  

        // Arrange:
        engine = ref.getCryptoEngine() CryptoEngine // final
        skp = KeyPair.random(engine) KeyPair // final
        rkp = KeyPair.random(engine) KeyPair // final
        blockCipher1 = ref.getBlockCipher(skp, NewKeyPair(rkp.getPublicKey(), engine)) BlockCipher // final
        blockCipher2 = ref.getBlockCipher(NewKeyPair(skp.getPublicKey(), engine), rkp) BlockCipher // final
        input = Utils.generateRandomBytes() []byte // final
        // Act:
        encryptedBytes = blockCipher1.encrypt(input) []byte // final
        decryptedBytes = blockCipher2.decrypt(encryptedBytes) []byte // final
        // Assert:
        Assert.assertThat(decryptedBytes, IsEqual.equalTo(input)) 
}

// @Test
   func (ref *BlockCipherTest) DataCanBeDecryptedWithSenderPrivateKeyAndRecipientPublicKey()    { /* public  */  

        // Arrange:
        engine = ref.getCryptoEngine() CryptoEngine // final
        skp = KeyPair.random(engine) KeyPair // final
        rkp = KeyPair.random(engine) KeyPair // final
        blockCipher1 = ref.getBlockCipher(skp, NewKeyPair(rkp.getPublicKey(), engine)) BlockCipher // final
        blockCipher2 = ref.getBlockCipher(NewKeyPair(rkp.getPublicKey(), engine), skp) BlockCipher // final
        input = Utils.generateRandomBytes() []byte // final
        // Act:
        encryptedBytes = blockCipher1.encrypt(input) []byte // final
        decryptedBytes = blockCipher2.decrypt(encryptedBytes) []byte // final
        // Assert:
        Assert.assertThat(decryptedBytes, IsEqual.equalTo(input)) 
}

// @Test
   func (ref *BlockCipherTest) DataEncryptedWithPrivateKeyCanOnlyBeDecryptedByMatchingPublicKey()    { /* public  */  

        // Arrange:
        engine = ref.getCryptoEngine() CryptoEngine // final
        blockCipher1 = ref.getBlockCipher(KeyPair.random(engine), KeyPair.random(engine)) BlockCipher // final
        blockCipher2 = ref.getBlockCipher(KeyPair.random(engine), KeyPair.random(engine)) BlockCipher // final
        input = Utils.generateRandomBytes() []byte // final
        // Act:
        encryptedBytes1 = blockCipher1.encrypt(input) []byte // final
        encryptedBytes2 = blockCipher2.encrypt(input) []byte // final
        // Assert:
        Assert.assertThat(blockCipher1.decrypt(encryptedBytes1), IsEqual.equalTo(input)) 
        Assert.assertThat(blockCipher1.decrypt(encryptedBytes2), IsNot.not(IsEqual.equalTo(input))) 
        Assert.assertThat(blockCipher2.decrypt(encryptedBytes1), IsNot.not(IsEqual.equalTo(input))) 
        Assert.assertThat(blockCipher2.decrypt(encryptedBytes2), IsEqual.equalTo(input)) 
}

    func (ref *BlockCipherTest) getBlockCipher(final KeyPair senderKeyPair, final KeyPair recipientKeyPair) BlockCipher { /* protected  */  

        return ref.CryptoEngine.createBlockCipher(senderKeyPair, recipientKeyPair) 
}

    getCryptoEngine() CryptoEngine // protected abstract
}

