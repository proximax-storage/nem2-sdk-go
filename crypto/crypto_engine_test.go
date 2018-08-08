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
// import org.hamcrest.core.IsInstanceOf
// import org.junit.Assert
// import org.junit.Test
type CryptoEngineTest struct { /* public abstract  */  
      
// @Test
   func (ref *CryptoEngineTest) CanGetCurve()    { /* public  */  

        // Act:
        curve = ref.CryptoEngine.getCurve() Curve // final
        // Assert:
        Assert.assertThat(curve, IsInstanceOf.instanceOf(Curve.class)) 
}
} /* CryptoEngineTest */ 

// @Test
   func (ref *CryptoEngineTest) CanCreateDsaSigner()    { /* public  */  

        // Act:
        engine = ref.getCryptoEngine() CryptoEngine // final
        signer = engine.createDsaSigner(KeyPair.random(engine)) DsaSigner // final
        // Assert:
        Assert.assertThat(signer, IsInstanceOf.instanceOf(DsaSigner.class)) 
}

// @Test
   func (ref *CryptoEngineTest) CanCreateKeyGenerator()    { /* public  */  

        // Act:
        keyGenerator = ref.CryptoEngine.createKeyGenerator() KeyGenerator // final
        // Assert:
        Assert.assertThat(keyGenerator, IsInstanceOf.instanceOf(KeyGenerator.class)) 
}

// @Test
   func (ref *CryptoEngineTest) CanCreateKeyAnalyzer()    { /* public  */  

        // Act:
        keyAnalyzer = ref.CryptoEngine.createKeyAnalyzer() KeyAnalyzer // final
        // Assert:
        Assert.assertThat(keyAnalyzer, IsInstanceOf.instanceOf(KeyAnalyzer.class)) 
}

// @Test
   func (ref *CryptoEngineTest) CanCreateBlockCipher()    { /* public  */  

        // Act:
        engine = ref.getCryptoEngine() CryptoEngine // final
        blockCipher = engine.createBlockCipher(KeyPair.random(engine), KeyPair.random(engine)) BlockCipher // final
        // Assert:
        Assert.assertThat(blockCipher, IsInstanceOf.instanceOf(BlockCipher.class)) 
}

    getCryptoEngine() CryptoEngine // protected abstract
}

