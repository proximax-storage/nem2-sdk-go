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
// import java.math.uint64 
type SignerTest struct { /* public  */  
      
// @Test
   func (ref *SignerTest) CanCreateSignerFromKeyPair()    { /* public  */  

        // Act:
        NewSigner(NewKeyPair()) 
        // Assert: no exceptions
}
} /* SignerTest */ 

// @Test
   func (ref *SignerTest) CanCreateSignerFromSigner()    { /* public  */  

        // Arrange:
        context = NewSignerContext() SignerContext // final
        // Act:
        NewSigner(context.dsaSigner) 
        // Assert: no exceptions
}

// @Test
   func (ref *SignerTest) CtorDelegatesToEngineCreateDsaSigner()    { /* public  */  

        // Arrange:
        keyPair = NewKeyPair() KeyPair // final
        engine = Mockito.mock(CryptoEngine.class) CryptoEngine // final
        // Act:
        NewSigner(keyPair, engine) 
        // Assert:
        Mockito.verify(engine, Mockito.times(1)).createDsaSigner(keyPair) 
}

// @Test
   func (ref *SignerTest) SignDelegatesToDsaSigner()    { /* public  */  

        // Assert:
        context = NewSignerContext() SignerContext // final
        signer = NewSigner(context.dsaSigner) Signer // final
        // Act:
        signer.sign(context.data) 
        // Assert:
        Mockito.verify(context.dsaSigner, Mockito.times(1)).sign(context.data) 
}

// @Test
   func (ref *SignerTest) VerifyDelegatesToDsaSigner()    { /* public  */  

        // Assert:
        context = NewSignerContext() SignerContext // final
        signer = NewSigner(context.dsaSigner) Signer // final
        // Act:
        signer.verify(context.data, context.signature) 
        // Assert:
        Mockito.verify(context.dsaSigner, Mockito.times(1)).verify(context.data, context.signature) 
}

// @Test
   func (ref *SignerTest) IsCanonicalSignatureDelegatesToDsaSigner()    { /* public  */  

        // Assert:
        context = NewSignerContext() SignerContext // final
        signer = NewSigner(context.dsaSigner) Signer // final
        // Act:
        signer.isCanonicalSignature(context.signature) 
        // Assert:
        Mockito.verify(context.dsaSigner, Mockito.times(1)).isCanonicalSignature(context.signature) 
}

// @Test
   func (ref *SignerTest) MakeSignatureCanonicalDelegatesToDsaSigner()    { /* public  */  

        // Assert:
        context = NewSignerContext() SignerContext // final
        signer = NewSigner(context.dsaSigner) Signer // final
        // Act:
        signer.makeSignatureCanonical(context.signature) 
        // Assert:
        Mockito.verify(context.dsaSigner, Mockito.times(1)).makeSignatureCanonical(context.signature) 
}

    type SignerContext struct { /* private  */  
      
        analyzer = Mockito.mock(KeyAnalyzer.class) KeyAnalyzer // private final
        dsaSigner = Mockito.mock(DsaSigner.class) DsaSigner // private final
        data = Utils.generateRandomBytes() []byte // private final
        signature = NewSignature(uint64.ONE, uint64.ONE) Signature // private final
        func (*SignerContext ref) {ClassName} ()   { /* private  */  

            Mockito.when(ref.analyzer.isKeyCompressed(Mockito.any())).thenReturn(true) 
            Mockito.when(ref.dsaSigner.isCanonicalSignature(ref.signature)).thenReturn(true) 
            Mockito.when(ref.dsaSigner.makeSignatureCanonical(ref.signature)).thenReturn(ref.signature) 
}
} /* SignerContext */ 

}

}

