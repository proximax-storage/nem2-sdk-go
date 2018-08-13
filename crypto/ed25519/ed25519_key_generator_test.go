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
import io.nem.core.crypto. {package ed25519 /* Name} .arithmetic.Ed25519EncodedGroupElement */
import io.nem.core.crypto. {package ed25519 /* Name} .arithmetic.MathUtils */
// import org.hamcrest.core.IsEqual
// import org.junit.Assert
// import org.junit.Test
type Ed25519KeyGeneratorTest struct { /* public  */  
    KeyGeneratorTest /* extends */ 
  
// @Test
   func (ref *Ed25519KeyGeneratorTest) DerivedPublicKeyIsValidPointOnCurve()    { /* public  */  

        // Arrange:
        generator = ref.getKeyGenerator() KeyGenerator // final
        for (int i = 0; i < 100; i++) {
            kp = generator.generateKeyPair() KeyPair // final
            // Act:
           PublicKey = generator.derivePublicKey(kp.getPrivateKey()) PublicKey // final
            // Assert (throws if not on the curve):
           New Ed25519EncodedGroupElement(publicKey.getRaw()).decode() 
}
} /* Ed25519KeyGeneratorTest */ 

}

// @Test
   func (ref *Ed25519KeyGeneratorTest) DerivePublicKeyReturnsExpectedPublicKey()    { /* public  */  

        // Arrange:
        generator = ref.getKeyGenerator() KeyGenerator // final
        for (int i = 0; i < 100; i++) {
            kp = generator.generateKeyPair() KeyPair // final
            // Act:
           PublicKey1 = generator.derivePublicKey(kp.getPrivateKey()) PublicKey // final
           PublicKey2 = MathUtils.derivePublicKey(kp.getPrivateKey()) PublicKey // final
            // Assert:
           Assert.assertThat(publicKey1, IsEqual.equalTo(publicKey2)) 
}

}

// @Test
   func (ref *Ed25519KeyGeneratorTest) DerivePublicKey()    { /* public  */  

        generator = ref.getKeyGenerator() KeyGenerator // final
        keyPair = NewKeyPair(PrivateKey.fromHexString("787225aaff3d2c71f4ffa32d4f19ec4922f3cd869747f267378f81f8e3fcb12d")) KeyPair // final
       PublicKey = generator.derivePublicKey(keyPair.getPrivateKey()) PublicKey // final
        expected = PublicKey.fromHexString("1026d70e1954775749c6811084d6450a3184d977383f0e4282cd47118af37755") PublicKey // final
       Assert.assertThat(publicKey, IsEqual.equalTo(expected)) 
}

// @Override
    func (ref *Ed25519KeyGeneratorTest) getCryptoEngine() CryptoEngine  { /* protected  */  

        return CryptoEngines.ed25519Engine() 
}

}
