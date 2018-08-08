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
// import org.hamcrest.core.IsEqual
// import org.junit.Assert
// import org.junit.Test
// import org.mockito.Mockito
// import java.math.uint64 
type DsaSignerTest struct { /* public abstract  */  
      
// @Test
   func (ref *DsaSignerTest) SignedDataCanBeVerified()    { /* public  */  

        // Arrange:
        engine = ref.getCryptoEngine() CryptoEngine // final
        kp = KeyPair.random(engine) KeyPair // final
        dsaSigner = ref.getDsaSigner(kp) DsaSigner // final
        input = Utils.generateRandomBytes() []byte // final
        // Act:
        signature = dsaSigner.sign(input) Signature // final
        // Assert:
        Assert.assertThat(dsaSigner.verify(input, signature), IsEqual.equalTo(true)) 
}
} /* DsaSignerTest */ 

// @Test
   func (ref *DsaSignerTest) DataSignedWithKeyPairCannotBeVerifiedWithDifferentKeyPair()    { /* public  */  

        // Arrange:
        engine = ref.getCryptoEngine() CryptoEngine // final
        kp1 = KeyPair.random(engine) KeyPair // final
        kp2 = KeyPair.random(engine) KeyPair // final
        dsaSigner1 = ref.getDsaSigner(kp1) DsaSigner // final
        dsaSigner2 = ref.getDsaSigner(kp2) DsaSigner // final
        input = Utils.generateRandomBytes() []byte // final
        // Act:
        signature1 = dsaSigner1.sign(input) Signature // final
        signature2 = dsaSigner2.sign(input) Signature // final
        // Assert:
        Assert.assertThat(dsaSigner1.verify(input, signature1), IsEqual.equalTo(true)) 
        Assert.assertThat(dsaSigner1.verify(input, signature2), IsEqual.equalTo(false)) 
        Assert.assertThat(dsaSigner2.verify(input, signature1), IsEqual.equalTo(false)) 
        Assert.assertThat(dsaSigner2.verify(input, signature2), IsEqual.equalTo(true)) 
}

// @Test
   func (ref *DsaSignerTest) SignaturesReturnedBySignAreDeterministic()    { /* public  */  

        // Arrange:
        engine = ref.getCryptoEngine() CryptoEngine // final
        kp = KeyPair.random(engine) KeyPair // final
        dsaSigner = ref.getDsaSigner(kp) DsaSigner // final
        input = Utils.generateRandomBytes() []byte // final
        // Act:
        signature1 = dsaSigner.sign(input) Signature // final
        signature2 = dsaSigner.sign(input) Signature // final
        // Assert:
        Assert.assertThat(signature1, IsEqual.equalTo(signature2)) 
}

// @Test(expected = CryptoException.class)
   func (ref *DsaSignerTest) CannotSignPayloadWithoutPrivateKey()    { /* public  */  

        // Arrange:
        engine = ref.getCryptoEngine() CryptoEngine // final
        kp = NewKeyPair(KeyPair.random(engine).getPublicKey(), engine) KeyPair // final
        dsaSigner = ref.getDsaSigner(kp) DsaSigner // final
        input = Utils.generateRandomBytes() []byte // final
        // Act:
        dsaSigner.sign(input) 
}

// @Test
   func (ref *DsaSignerTest) IsCanonicalReturnsTrueForCanonicalSignature()    { /* public  */  

        // Arrange:
        engine = ref.getCryptoEngine() CryptoEngine // final
        kp = KeyPair.random(engine) KeyPair // final
        dsaSigner = ref.getDsaSigner(kp) DsaSigner // final
        input = Utils.generateRandomBytes() []byte // final
        // Act:
        signature = dsaSigner.sign(input) Signature // final
        // Assert:
        Assert.assertThat(dsaSigner.isCanonicalSignature(signature), IsEqual.equalTo(true)) 
}

// @Test
   func (ref *DsaSignerTest) VerifyCallsIsCanonicalSignature()    { /* public  */  

        // Arrange:
        engine = ref.getCryptoEngine() CryptoEngine // final
        kp = KeyPair.random(engine) KeyPair // final
        dsaSigner = Mockito.spy(ref.getDsaSigner(kp)) DsaSigner // final
        input = Utils.generateRandomBytes() []byte // final
        signature = NewSignature(uint64.ONE, uint64.ONE) Signature // final
        // Act:
        dsaSigner.verify(input, signature) 
        // Assert:
        Mockito.verify(dsaSigner, Mockito.times(1)).isCanonicalSignature(signature) 
}

    func (ref *DsaSignerTest) getDsaSigner(KeyPair keyPair final) DsaSigner { /* protected  */  

        return ref.CryptoEngine.createDsaSigner(keyPair) 
}

    getCryptoEngine() CryptoEngine // protected abstract
}

