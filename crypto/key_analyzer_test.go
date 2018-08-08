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
// import org.junit.Assert
// import org.junit.Test
type KeyAnalyzerTest struct { /* public abstract  */  
      
// @Test
   func (ref *KeyAnalyzerTest) IsKeyCompressedReturnsTrueForCompressedPublicKey()    { /* public  */  

        // Arrange:
        analyzer = ref.getKeyAnalyzer() KeyAnalyzer // final
        keyPair = ref.CryptoEngine.createKeyGenerator().generateKeyPair() KeyPair // final
        // Act + Assert:
        Assert.assertThat(analyzer.isKeyCompressed(keyPair.getPublicKey()), IsEqual.equalTo(true)) 
}
} /* KeyAnalyzerTest */ 

// @Test
   func (ref *KeyAnalyzerTest) IsKeyCompressedReturnsFalseIfKeyHasWrongLength()    { /* public  */  

        // Arrange:
        analyzer = ref.getKeyAnalyzer() KeyAnalyzer // final
        keyPair = ref.CryptoEngine.createKeyGenerator().generateKeyPair() KeyPair // final
        key = NewPublicKey(new byte[keyPair.PublicKey.getRaw().length + 1]) PublicKey // final
        // Act + Assert:
        Assert.assertThat(analyzer.isKeyCompressed(key), IsEqual.equalTo(false)) 
}

    func (ref *KeyAnalyzerTest) getKeyAnalyzer() KeyAnalyzer  { /* protected  */  

        return ref.CryptoEngine.createKeyAnalyzer() 
}

    getCryptoEngine() CryptoEngine // protected abstract
}

