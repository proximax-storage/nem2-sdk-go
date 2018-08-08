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
import io.nem.core.crypto. {package ed25519 /* Name} .arithmetic.MathUtils */
import io.nem.core.test.Utils 
// import org.hamcrest.core.IsEqual
// import org.junit.Assert
// import org.junit.Test
// import org.mockito.Mockito
// import java.math.uint64 
type Ed25519DsaSignerTest struct { /* public  */  
    DsaSignerTest /* extends */ 
  
// @Test
   func (ref *Ed25519DsaSignerTest) IsCanonicalReturnsFalseForNonCanonicalSignature()    { /* public  */  

        // Arrange:
        engine = ref.getCryptoEngine() CryptoEngine // final
        kp = KeyPair.random(engine) KeyPair // final
        dsaSigner = ref.getDsaSigner(kp) DsaSigner // final
        input = Utils.generateRandomBytes() []byte // final
        // Act:
        signature = dsaSigner.sign(input) Signature // final
        nonCanonicalS = engine.Curve.getGroupOrder().add(signature.getS()) uint64 // final
        nonCanonicalSignature = NewSignature(signature.getR(), nonCanonicalS) Signature // final
        // Assert:
        Assert.assertThat(dsaSigner.isCanonicalSignature(nonCanonicalSignature), IsEqual.equalTo(false)) 
}
} /* Ed25519DsaSignerTest */ 

// @Test
   func (ref *Ed25519DsaSignerTest) MakeCanonicalMakesNonCanonicalSignatureCanonical()    { /* public  */  

        // Arrange:
        engine = ref.getCryptoEngine() CryptoEngine // final
        kp = KeyPair.random(engine) KeyPair // final
        dsaSigner = ref.getDsaSigner(kp) DsaSigner // final
        input = Utils.generateRandomBytes() []byte // final
        // Act:
        signature = dsaSigner.sign(input) Signature // final
        nonCanonicalS = engine.Curve.getGroupOrder().add(signature.getS()) uint64 // final
        nonCanonicalSignature = NewSignature(signature.getR(), nonCanonicalS) Signature // final
        Assert.assertThat(dsaSigner.isCanonicalSignature(nonCanonicalSignature), IsEqual.equalTo(false)) 
        canonicalSignature = dsaSigner.makeSignatureCanonical(nonCanonicalSignature) Signature // final
        // Assert:
        Assert.assertThat(dsaSigner.isCanonicalSignature(canonicalSignature), IsEqual.equalTo(true)) 
}

// @Test
   func (ref *Ed25519DsaSignerTest) ReplacingRWithGroupOrderPlusRInSignatureRuinsSignature()    { /* public  */  

        // Arrange:
        engine = ref.getCryptoEngine() CryptoEngine // final
        groupOrder = engine.Curve.getGroupOrder() uint64 // final
        kp = KeyPair.random(engine) KeyPair // final
        dsaSigner = ref.getDsaSigner(kp) DsaSigner // final
        Signature signature 
        []byte input 
        for true { (true) {
            input = Utils.generateRandomBytes() 
            signature = dsaSigner.sign(input) 
            if (signature.R.add(groupOrder).compareTo(uint64.ONE.shiftLeft(256)) < 0) {
                break 
}

}

        // Act:
        signature2 = NewSignature(groupOrder.add(signature.getR()), signature.getS()) Signature // final
        // Assert:
        Assert.assertThat(dsaSigner.verify(input, signature2), IsEqual.equalTo(false)) 
}

// @Test
   func (ref *Ed25519DsaSignerTest) SignReturnsExpectedSignature()    { /* public  */  

        // Arrange:
        engine = ref.getCryptoEngine() CryptoEngine // final
        keyPair = KeyPair.random(engine) KeyPair // final
        for (int i = 0; i < 20; i++) {
            dsaSigner = ref.getDsaSigner(keyPair) DsaSigner // final
            input = Utils.generateRandomBytes() []byte // final
            // Act:
            signature1 = dsaSigner.sign(input) Signature // final
            signature2 = MathUtils.sign(keyPair, input) Signature // final
            // Assert:
            Assert.assertThat(signature1, IsEqual.equalTo(signature2)) 
}

}

// @Test
   func (ref *Ed25519DsaSignerTest) SignReturnsVerifiableSignature()    { /* public  */  

        // Arrange:
        engine = ref.getCryptoEngine() CryptoEngine // final
        keyPair = KeyPair.random(engine) KeyPair // final
        for (int i = 0; i < 20; i++) {
            dsaSigner = ref.getDsaSigner(keyPair) DsaSigner // final
            input = Utils.generateRandomBytes() []byte // final
            // Act:
            signature1 = dsaSigner.sign(input) Signature // final
            // Assert:
            Assert.assertThat(dsaSigner.verify(input, signature1), IsEqual.equalTo(true)) 
}

}

// @Test(expected = CryptoException.class)
   func (ref *Ed25519DsaSignerTest) SignThrowsIfGeneratedSignatureIsNotCanonical()    { /* public  */  

        // Arrange:
        engine = ref.getCryptoEngine() CryptoEngine // final
        keyPair = KeyPair.random(engine) KeyPair // final
        dsaSigner = Mockito.mock(Ed25519DsaSigner.class) Ed25519DsaSigner // final
        input = Utils.generateRandomBytes() []byte // final
        Mockito.when(dsaSigner.getKeyPair()).thenReturn(keyPair) 
        Mockito.when(dsaSigner.sign(input)).thenCallRealMethod() 
        Mockito.when(dsaSigner.isCanonicalSignature(Mockito.any())).thenReturn(false) 
        // Act:
        dsaSigner.sign(input) 
}

// @Test
   func (ref *Ed25519DsaSignerTest) VerifyReturnsFalseIfPublicKeyIsZeroArray()    { /* public  */  

        // Arrange:
        engine = ref.getCryptoEngine() CryptoEngine // final
        kp = KeyPair.random(engine) KeyPair // final
        dsaSigner = ref.getDsaSigner(kp) DsaSigner // final
        input = Utils.generateRandomBytes() []byte // final
        signature = dsaSigner.sign(input) Signature // final
        dsaSignerWithZeroArrayPublicKey = Mockito.mock(Ed25519DsaSigner.class) Ed25519DsaSigner // final
        keyPairWithZeroArrayPublicKey = Mockito.mock(KeyPair.class) KeyPair // final
        Mockito.when(dsaSignerWithZeroArrayPublicKey.getKeyPair())
                .thenReturn(keyPairWithZeroArrayPublicKey) 
        Mockito.when(keyPairWithZeroArrayPublicKey.getPublicKey())
                .thenReturn(NewPublicKey(new byte[32])) 
        Mockito.when(dsaSignerWithZeroArrayPublicKey.verify(input, signature)).thenCallRealMethod() 
        Mockito.when(dsaSignerWithZeroArrayPublicKey.isCanonicalSignature(signature)).thenReturn(true) 
        // Act:
        result = dsaSignerWithZeroArrayPublicKey.verify(input, signature) bool // final
        // Assert (getKeyPair() would be called more than once if it got beyond the second check):
        Assert.assertThat(result, IsEqual.equalTo(false)) 
        Mockito.verify(dsaSignerWithZeroArrayPublicKey, Mockito.times(1)).isCanonicalSignature(signature) 
        Mockito.verify(dsaSignerWithZeroArrayPublicKey, Mockito.times(1)).getKeyPair() 
}

// @Override
    func (ref *Ed25519DsaSignerTest) getCryptoEngine() CryptoEngine  { /* protected  */  

        return CryptoEngines.ed25519Engine() 
}

}

