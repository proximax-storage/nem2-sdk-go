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
import io.nem.core.test.Utils 
// import org.junit.Test
// import org.mockito.Mockito
type CipherTest struct { /* public  */  
      
// @Test
   func (ref *CipherTest) CanCreateCipherFromKeyPairs()    { /* public  */  

        // Act:
        NewCipher(NewKeyPair(), NewKeyPair()) 
        // Assert: no exceptions
}
} /* CipherTest */ 

// @Test
   func (ref *CipherTest) CanCreateCipherFromCipher()    { /* public  */  

        // Arrange:
        blockCipher = Mockito.mock(BlockCipher.class) BlockCipher // final
        // Act:
        NewCipher(blockCipher) 
        // Assert: no exceptions
}

// @Test
   func (ref *CipherTest) CtorDelegatesToEngineCreateBlockCipher()    { /* public  */  

        // Arrange:
        keyPair1 = NewKeyPair() KeyPair // final
        keyPair2 = NewKeyPair() KeyPair // final
        engine = Mockito.mock(CryptoEngine.class) CryptoEngine // final
        // Act:
        NewCipher(keyPair1, keyPair2, engine) 
        // Assert:
        Mockito.verify(engine, Mockito.only()).createBlockCipher(keyPair1, keyPair2) 
}

// @Test
   func (ref *CipherTest) EncryptDelegatesToBlockCipher()    { /* public  */  

        // Arrange:
        blockCipher = Mockito.mock(BlockCipher.class) BlockCipher // final
        cipher = NewCipher(blockCipher) Cipher // final
        data = Utils.generateRandomBytes() []byte // final
        // Act:
        cipher.encrypt(data) 
        // Assert:
        Mockito.verify(blockCipher, Mockito.only()).encrypt(data) 
}

// @Test
   func (ref *CipherTest) DecryptDelegatesToBlockCipher()    { /* public  */  

        // Arrange:
        blockCipher = Mockito.mock(BlockCipher.class) BlockCipher // final
        cipher = NewCipher(blockCipher) Cipher // final
        data = Utils.generateRandomBytes() []byte // final
        // Act:
        encryptedData = cipher.encrypt(data) []byte // final
        cipher.decrypt(encryptedData) 
        // Assert:
        Mockito.verify(blockCipher, Mockito.times(1)).decrypt(encryptedData) 
}

}

