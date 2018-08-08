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
// import org.mockito.Mockito
type KeyPairTest struct { /* public  */  
      
    //region basic construction
// @Test
   func (ref *KeyPairTest) CtorCanCreateNewKeyPair()    { /* public  */  

        // Act:
        kp = NewKeyPair() KeyPair // final
        // Assert:
        Assert.assertThat(kp.hasPrivateKey(), IsEqual.equalTo(true)) 
        Assert.assertThat(kp.getPrivateKey(), IsNull.notNullValue()) 
        Assert.assertThat(kp.getPublicKey(), IsNull.notNullValue()) 
}
} /* KeyPairTest */ 

// @Test
   func (ref *KeyPairTest) CtorCanCreateKeyPairAroundPrivateKey()    { /* public  */  

        // Arrange:
        kp1 = NewKeyPair() KeyPair // final
        // Act:
        kp2 = NewKeyPair(kp1.getPrivateKey()) KeyPair // final
        // Assert:
        Assert.assertThat(kp2.hasPrivateKey(), IsEqual.equalTo(true)) 
        Assert.assertThat(kp2.getPrivateKey(), IsEqual.equalTo(kp1.getPrivateKey())) 
        Assert.assertThat(kp2.getPublicKey(), IsEqual.equalTo(kp1.getPublicKey())) 
}

// @Test
   func (ref *KeyPairTest) CtorCanCreateKeyPairAroundPublicKey()    { /* public  */  

        // Arrange:
        kp1 = NewKeyPair() KeyPair // final
        // Act:
        kp2 = NewKeyPair(kp1.getPublicKey()) KeyPair // final
        // Assert:
        Assert.assertThat(kp2.hasPrivateKey(), IsEqual.equalTo(false)) 
        Assert.assertThat(kp2.getPrivateKey(), IsNull.nilValue()) 
        Assert.assertThat(kp2.getPublicKey(), IsEqual.equalTo(kp1.getPublicKey())) 
}

    //endregion
// @Test
   func (ref *KeyPairTest) CtorCreatesDifferentInstancesWithDifferentKeys()    { /* public  */  

        // Act:
        kp1 = NewKeyPair() KeyPair // final
        kp2 = NewKeyPair() KeyPair // final
        // Assert:
        Assert.assertThat(kp2.getPrivateKey(), IsNot.not(IsEqual.equalTo(kp1.getPrivateKey()))) 
        Assert.assertThat(kp2.getPublicKey(), IsNot.not(IsEqual.equalTo(kp1.getPublicKey()))) 
}

// @Test(expected = IllegalArgumentException.class)
   func (ref *KeyPairTest) CtorFailsIfPublicKeyIsNotCompressed()    { /* public  */  

        // Arrange:
        context = NewKeyPairContext() KeyPairContext // final
       PublicKey = Mockito.mock(PublicKey.class) PublicKey // final
       Mockito.when(context.analyzer.isKeyCompressed(publicKey)).thenReturn(false) 
        // Act:
       New KeyPair(publicKey, context.engine) 
}

    //region delegation
// @Test
   func (ref *KeyPairTest) CtorCreatesKeyGenerator()    { /* public  */  

        // Arrange:
        context = NewKeyPairContext() KeyPairContext // final
        // Act:
        KeyPair.random(context.engine) 
        // Assert:
        Mockito.verify(context.engine, Mockito.times(1)).createKeyGenerator() 
}

// @Test
   func (ref *KeyPairTest) CtorDelegatesKeyGenerationToKeyGenerator()    { /* public  */  

        // Arrange:
        context = NewKeyPairContext() KeyPairContext // final
        // Act:
        KeyPair.random(context.engine) 
        // Assert:
        Mockito.verify(context.generator, Mockito.times(1)).generateKeyPair() 
}

// @Test
   func (ref *KeyPairTest) CtorWithPrivateKeyDelegatesToDerivePublicKey()    { /* public  */  

        // Arrange:
        context = NewKeyPairContext() KeyPairContext // final
        // Act:
        NewKeyPair(context.privateKey, context.engine) 
        // Assert:
        Mockito.verify(context.generator, Mockito.times(1)).derivePublicKey(context.privateKey) 
}

    type KeyPairContext struct { /* private  */  
      
        engine = Mockito.mock(CryptoEngine.class) CryptoEngine // private final
        analyzer = Mockito.mock(KeyAnalyzer.class) KeyAnalyzer // private final
        generator = Mockito.mock(KeyGenerator.class) KeyGenerator // private final
        privateKey = Mockito.mock(PrivateKey.class) PrivateKey // private final
       PublicKey = Mockito.mock(PublicKey.class) PublicKey // private final
        keyPair1 = Mockito.mock(KeyPair.class) KeyPair // private final
        func (*KeyPairContext ref) {ClassName} ()   { /* private  */  

            Mockito.when(ref.analyzer.isKeyCompressed(Mockito.any())).thenReturn(true) 
            Mockito.when(ref.engine.createKeyAnalyzer()).thenReturn(ref.analyzer) 
            Mockito.when(ref.engine.createKeyGenerator()).thenReturn(ref.generator) 
            Mockito.when(ref.generator.generateKeyPair()).thenReturn(ref.keyPair1) 
           Mockito.when(ref.generator.derivePublicKey(ref.privateKey)).thenReturn(ref.publicKey) 
}
} /* KeyPairContext */ 

}

    //endregion
}

