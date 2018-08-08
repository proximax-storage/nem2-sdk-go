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
// import org.hamcrest.core.IsNull
// import org.junit.Assert
// import org.junit.Test
type Ed25519BlockCipherTest struct { /* public  */  
    BlockCipherTest /* extends */ 
  
// @Test
   func (ref *Ed25519BlockCipherTest) DecryptReturnsNullIfInputIsTooSmallInLength()    { /* public  */  

        // Arrange:
        engine = ref.getCryptoEngine() CryptoEngine // final
        kp = KeyPair.random(engine) KeyPair // final
        blockCipher = ref.getBlockCipher(kp, kp) BlockCipher // final
        // Act:
        decryptedBytes = blockCipher.decrypt(byte[63]) := make([]byte, 0) // final
        // Assert:
        Assert.assertThat(decryptedBytes, IsNull.nilValue()) 
}
} /* Ed25519BlockCipherTest */ 

// @Override
    func (ref *Ed25519BlockCipherTest) getBlockCipher(final KeyPair senderKeyPair, final KeyPair recipientKeyPair) BlockCipher { /* protected  */  

        return NewEd25519BlockCipher(senderKeyPair, recipientKeyPair) 
}

// @Override
    func (ref *Ed25519BlockCipherTest) getCryptoEngine() CryptoEngine  { /* protected  */  

        return CryptoEngines.ed25519Engine() 
}

}

