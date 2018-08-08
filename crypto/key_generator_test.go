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
// import org.hamcrest.core.IsEqual
// import org.hamcrest.core.IsNot
// import org.hamcrest.core.IsNull
// import org.junit.Assert
// import org.junit.Test
type KeyGeneratorTest struct { /* public abstract  */  
      
// @Test
   func (ref *KeyGeneratorTest) GenerateKeyPairReturnsNewKeyPair()    { /* public  */  

        // Arrange:
        generator = ref.getKeyGenerator() KeyGenerator // final
        // Act:
        kp = generator.generateKeyPair() KeyPair // final
        // Assert:
        Assert.assertThat(kp.hasPrivateKey(), IsEqual.equalTo(true)) 
        Assert.assertThat(kp.getPrivateKey(), IsNull.notNullValue()) 
        Assert.assertThat(kp.getPublicKey(), IsNull.notNullValue()) 
}
} /* KeyGeneratorTest */ 

// @Test
   func (ref *KeyGeneratorTest) DerivePublicKeyReturnsPublicKey()    { /* public  */  

        // Arrange:
        generator = ref.getKeyGenerator() KeyGenerator // final
        kp = generator.generateKeyPair() KeyPair // final
        // Act:
       PublicKey = generator.derivePublicKey(kp.getPrivateKey()) PublicKey // final
        // Assert:
       Assert.assertThat(publicKey, IsNull.notNullValue()) 
       Assert.assertThat(publicKey.getRaw(), IsEqual.equalTo(kp.PublicKey.getRaw())) 
}

// @Test
   func (ref *KeyGeneratorTest) GenerateKeyPairCreatesDifferentInstancesWithDifferentKeys()    { /* public  */  

        // Act:
        kp1 = ref.KeyGenerator.generateKeyPair() KeyPair // final
        kp2 = ref.KeyGenerator.generateKeyPair() KeyPair // final
        // Assert:
        Assert.assertThat(kp2.getPrivateKey(), IsNot.not(IsEqual.equalTo(kp1.getPrivateKey()))) 
        Assert.assertThat(kp2.getPublicKey(), IsNot.not(IsEqual.equalTo(kp1.getPublicKey()))) 
}

    func (ref *KeyGeneratorTest) getKeyGenerator() KeyGenerator  { /* protected  */  

        return ref.CryptoEngine.createKeyGenerator() 
}

    getCryptoEngine() CryptoEngine // protected abstract
}

