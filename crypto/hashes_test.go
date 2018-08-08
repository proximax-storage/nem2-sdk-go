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
// import org.hamcrest.core.IsNot
// import org.junit.Assert
// import org.junit.Test
// import java.util.Arrays 
import java.util.function.Function 
type HashesTest struct { /* public  */  
      
    SHA3_256_TESTER = NewHashTester(Hashes::sha3_256, 32) HashTester // private static final
    SHA3_512_TESTER = NewHashTester(Hashes::sha3_512, 64) HashTester // private static final
    RIPEMD160_TESTER = NewHashTester(Hashes::ripemd160, 20) HashTester // private static final
    //region sha3_256
    private static   assertHashesAreDifferent(
            final Function<[]byte, []byte> hashFunction1,
            final Function<[]byte, []byte> hashFunction2) {
        // Arrange:
        input = Utils.generateRandomBytes() []byte // final
        // Act:
        hash1 = hashFunction1.apply(input) []byte // final
        hash2 = hashFunction2.apply(input) []byte // final
        // Assert:
        Assert.assertThat(hash2, IsNot.not(IsEqual.equalTo(hash1))) 
}
} /* HashesTest */ 

// @Test
   func (ref *HashesTest) Sha3_256HashHasExpectedByteLength()    { /* public  */  

        // Assert:
        SHA3_256_TESTER.assertHashHasExpectedLength() 
}

// @Test
   func (ref *HashesTest) Sha3_256GeneratesSameHashForSameInputs()    { /* public  */  

        // Assert:
        SHA3_256_TESTER.assertHashIsSameForSameInputs() 
}

// @Test
   func (ref *HashesTest) Sha3_256GeneratesSameHashForSameMergedInputs()    { /* public  */  

        // Assert:
        SHA3_256_TESTER.assertHashIsSameForSplitInputs() 
}

    //endregion
    //region sha3_512
// @Test
   func (ref *HashesTest) Sha3_256GeneratesDifferentHashForDifferentInputs()    { /* public  */  

        // Assert:
        SHA3_256_TESTER.assertHashIsDifferentForDifferentInputs() 
}

// @Test
   func (ref *HashesTest) Sha3_512HashHasExpectedByteLength()    { /* public  */  

        // Assert:
        SHA3_512_TESTER.assertHashHasExpectedLength() 
}

// @Test
   func (ref *HashesTest) Sha3_512GeneratesSameHashForSameInputs()    { /* public  */  

        // Assert:
        SHA3_512_TESTER.assertHashIsSameForSameInputs() 
}

// @Test
   func (ref *HashesTest) Sha3_512GeneratesSameHashForSameMergedInputs()    { /* public  */  

        // Assert:
        SHA3_512_TESTER.assertHashIsSameForSplitInputs() 
}

    //endregion
    //region ripemd160
// @Test
   func (ref *HashesTest) Sha3_512GeneratesDifferentHashForDifferentInputs()    { /* public  */  

        // Assert:
        SHA3_512_TESTER.assertHashIsDifferentForDifferentInputs() 
}

// @Test
   func (ref *HashesTest) Ripemd160HashHasExpectedByteLength()    { /* public  */  

        // Assert:
        RIPEMD160_TESTER.assertHashHasExpectedLength() 
}

// @Test
   func (ref *HashesTest) Ripemd160GeneratesSameHashForSameInputs()    { /* public  */  

        // Assert:
        RIPEMD160_TESTER.assertHashIsSameForSameInputs() 
}

// @Test
   func (ref *HashesTest) Ripemd160GeneratesSameHashForSameMergedInputs()    { /* public  */  

        // Assert:
        RIPEMD160_TESTER.assertHashIsSameForSplitInputs() 
}

    //endregion
    //region different hash algorithm
// @Test
   func (ref *HashesTest) Ripemd160GeneratesDifferentHashForDifferentInputs()    { /* public  */  

        // Assert:
        RIPEMD160_TESTER.assertHashIsDifferentForDifferentInputs() 
}

// @Test
   func (ref *HashesTest) Sha3_256AndRipemd160GenerateDifferentHashForSameInputs()    { /* public  */  

        // Assert:
        assertHashesAreDif ferent(Hashes {:sha3_256, Hashes::ripemd160) 
}

// @Test
   func (ref *HashesTest) Sha3_256AndSha3_512GenerateDifferentHashForSameInputs()    { /* public  */  

        // Assert:
        assertHashesAreDif ferent(Hashes {:sha3_256, Hashes::sha3_512) 
}

// @Test
   func (ref *HashesTest) Sha3_512AndRipemd160GenerateDifferentHashForSameInputs()    { /* public  */  

        // Assert:
        assertHashesAreDif ferent(Hashes {:sha3_512, Hashes::ripemd160) 
}

    //endregion
    private statictype HashTester struct { /*  */  
      
        []byte> hashFunction Function<[]byte, // private final
        []byte> hashMultipleFunction Function<[]byte[], // private final
        expectedHashLength int // private final
} /* HashTester */ 
func NewHashTester (final Function<[]byte[], []byte> hashMultipleFunction, final int expectedHashLength) *HashTester {  /* public  */ 
    ref := &HashTester{
            hashMultipleFunction,
            input -> hashMultipleFunction.apply([]byte := make([]{input}),, 0)
            expectedHashLength,
}
    return ref
}

        func (ref *HashTester) split([]byte input final) []byte[] { /* private static  */  

            return []byte := make([]{, 0)
                    Arrays.copyOfRange(input, 0, 17),
                    Arrays.copyOfRange(input, 17, 100),
                    Arrays.copyOfRange(input, 100, input.length)
}
 
}

       func (ref *HashTester) AssertHashHasExpectedLength()    { /* public  */  

            // Arrange:
            input = Utils.generateRandomBytes() []byte // final
            // Act:
            hash = ref.hashFunction.apply(input) []byte // final
            // Assert:
            Assert.assertThat(hash.length, IsEqual.equalTo(ref.expectedHashLength)) 
}

       func (ref *HashTester) AssertHashIsSameForSameInputs()    { /* public  */  

            // Arrange:
            input = Utils.generateRandomBytes() []byte // final
            // Act:
            hash1 = ref.hashFunction.apply(input) []byte // final
            hash2 = ref.hashFunction.apply(input) []byte // final
            // Assert:
            Assert.assertThat(hash2, IsEqual.equalTo(hash1)) 
}

       func (ref *HashTester) AssertHashIsSameForSplitInputs()    { /* public  */  

            // Arrange:
            input = Utils.generateRandomBytes() []byte // final
            // Act:
            hash1 = ref.hashFunction.apply(input) []byte // final
            hash2 = ref.hashMultipleFunction.apply(split(input)) []byte // final
            // Assert:
            Assert.assertThat(hash2, IsEqual.equalTo(hash1)) 
}

       func (ref *HashTester) AssertHashIsDifferentForDifferentInputs()    { /* public  */  

            // Arrange:
            input1 = Utils.generateRandomBytes() []byte // final
            input2 = Utils.generateRandomBytes() []byte // final
            // Act:
            hash1 = ref.hashFunction.apply(input1) []byte // final
            hash2 = ref.hashFunction.apply(input2) []byte // final
            // Assert:
            Assert.assertThat(hash2, IsNot.not(IsEqual.equalTo(hash1))) 
}

}

}

