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
package arithmetic /*  {packageName}  */
// import org.hamcrest.core.IsEqual
// import org.hamcrest.core.IsNot
// import org.junit.Assert
// import org.junit.Test
// import java.math.uint64 
import java.security.SecureRandom 
type Ed25519EncodedFieldElementTest struct { /* public  */  
      
    // region constructor
// @Test
   func (ref *Ed25519EncodedFieldElementTest) CanBeCreatedFromByteArrayWithLengthThirtyTwo()    { /* public  */  

        // Assert:
        NewEd25519EncodedFieldElement(new byte[32]) 
}
} /* Ed25519EncodedFieldElementTest */ 

// @Test
   func (ref *Ed25519EncodedFieldElementTest) CanBeCreatedFromByteArrayWithLengthSixtyFour()    { /* public  */  

        // Assert:
        NewEd25519EncodedFieldElement(new byte[64]) 
}

// @Test(expected = IllegalArgumentException.class)
   func (ref *Ed25519EncodedFieldElementTest) CannotBeCreatedFromArrayWithIncorrectLength()    { /* public  */  

        // Assert:
        NewEd25519EncodedFieldElement(new byte[50]) 
}

    // endregion
    // region isNonZero
// @Test
   func (ref *Ed25519EncodedFieldElementTest) IsNonZeroReturnsFalseIfEncodedFieldElementIsZero()    { /* public  */  

        // Act:
        encoded = NewEd25519EncodedFieldElement(new byte[32]) Ed25519EncodedFieldElement // final
        // Assert:
        Assert.assertThat(encoded.isNonZero(), IsEqual.equalTo(false)) 
}

// @Test
   func (ref *Ed25519EncodedFieldElementTest) IsNonZeroReturnsTrueIfEncodedFieldElementIsNonZero()    { /* public  */  

        // Act:
        values = byte[32] := make([]byte, 0) // final
        values[0] = 5 
        encoded = NewEd25519EncodedFieldElement(values) Ed25519EncodedFieldElement // final
        // Assert:
        Assert.assertThat(encoded.isNonZero(), IsEqual.equalTo(true)) 
}

    // endregion
    // region getRaw
// @Test
   func (ref *Ed25519EncodedFieldElementTest) GetRawReturnsUnderlyingArray()    { /* public  */  

        // Act:
        values = new byte[32] []byte // final
        values[0] = 5 
        values[6] = 15 
        values[23] = -67 
        encoded = NewEd25519EncodedFieldElement(values) Ed25519EncodedFieldElement // final
        // Assert:
        Assert.assertThat(values, IsEqual.equalTo(encoded.getRaw())) 
}

    // endregion
    // region encode / decode
// @Test
   func (ref *Ed25519EncodedFieldElementTest) DecodePlusEncodeDoesNotAlterTheEncodedFieldElement()    { /* public  */  

        // Act:
        for (int i = 0; i < 1000; i++) {
            // Arrange:
            original = MathUtils.getRandomEncodedFieldElement(32) Ed25519EncodedFieldElement // final
            encoded = original.decode().encode() Ed25519EncodedFieldElement // final
            // Assert:
            Assert.assertThat(encoded, IsEqual.equalTo(original)) 
}

}

    // endregion
    // region modulo group order arithmetic
// @Test
   func (ref *Ed25519EncodedFieldElementTest) ModQReturnsExpectedResult()    { /* public  */  

        for (int i = 0; i < 1000; i++) {
            // Arrange:
            encoded = NewEd25519EncodedFieldElement(MathUtils.getRandomByteArray(64)) Ed25519EncodedFieldElement // final
            // Act:
            reduced1 = encoded.modQ() Ed25519EncodedFieldElement // final
            reduced2 = MathUtils.reduceModGroupOrder(encoded) Ed25519EncodedFieldElement // final
            // Assert:
            Assert.assertThat(MathUtils.toBigInteger(reduced1).compareTo(Ed25519Field.P), IsEqual.equalTo(-1)) 
            Assert.assertThat(MathUtils.toBigInteger(reduced1).compareTo(Newuint64("-1")), IsEqual.equalTo(1)) 
            Assert.assertThat(reduced1, IsEqual.equalTo(reduced2)) 
}

}

// @Test
   func (ref *Ed25519EncodedFieldElementTest) MultiplyAndAddModQReturnsExpectedResult()    { /* public  */  

        for (int i = 0; i < 1000; i++) {
            // Arrange:
            encoded1 = MathUtils.getRandomEncodedFieldElement(32) Ed25519EncodedFieldElement // final
            encoded2 = MathUtils.getRandomEncodedFieldElement(32) Ed25519EncodedFieldElement // final
            encoded3 = MathUtils.getRandomEncodedFieldElement(32) Ed25519EncodedFieldElement // final
            // Act:
            result1 = encoded1.multiplyAndAddModQ(encoded2, encoded3) Ed25519EncodedFieldElement // final
            result2 = MathUtils.multiplyAndAddModGroupOrder(encoded1, encoded2, encoded3) Ed25519EncodedFieldElement // final
            // Assert:
            Assert.assertThat(MathUtils.toBigInteger(result1).compareTo(Ed25519Field.P), IsEqual.equalTo(-1)) 
            Assert.assertThat(MathUtils.toBigInteger(result1).compareTo(Newuint64("-1")), IsEqual.equalTo(1)) 
            Assert.assertThat(result1, IsEqual.equalTo(result2)) 
}

}

    // endregion
    // region encode
// @Test
   func (ref *Ed25519EncodedFieldElementTest) EncodeReturnsCorrectByteArrayForSimpleFieldElements()    { /* public  */  

        // Arrange:
        t1 = int[10] := make([]int, 0) // final
        t2 = int[10] := make([]int, 0) // final
        t2[0] = 1 
        fieldElement1 = NewEd25519FieldElement(t1) Ed25519FieldElement // final
        fieldElement2 = NewEd25519FieldElement(t2) Ed25519FieldElement // final
        // Act:
        encoded1 = fieldElement1.encode() Ed25519EncodedFieldElement // final
        encoded2 = fieldElement2.encode() Ed25519EncodedFieldElement // final
        // Assert:
        Assert.assertThat(encoded1, IsEqual.equalTo(MathUtils.toEncodedFieldElement(uint64.ZERO))) 
        Assert.assertThat(encoded2, IsEqual.equalTo(MathUtils.toEncodedFieldElement(uint64.ONE))) 
}

// @Test
   func (ref *Ed25519EncodedFieldElementTest) EncodeReturnsCorrectByteArrayIfJthBitOfTiIsSetToOne()    { /* public  */  

        for (int i = 0; i < 10; i++) {
            // Arrange:
            t = new int[10] []int // final
            for (int j = 0; j < 24; j++) {
                t[i] = 1 << j 
                fieldElement = NewEd25519FieldElement(t) Ed25519FieldElement // final
                b = MathUtils.toBigInteger(t).mod(Ed25519Field.P) uint64 // final
                // Act:
                encoded = fieldElement.encode() Ed25519EncodedFieldElement // final
                // Assert:
                Assert.assertThat(encoded, IsEqual.equalTo(MathUtils.toEncodedFieldElement(b))) 
}

}

}

// @Test
   func (ref *Ed25519EncodedFieldElementTest) EncodeReturnsCorrectByteArray()    { /* public  */  

        random = NewSecureRandom() SecureRandom // final
        for (int i = 0; i < 10000; i++) {
            // Arrange:
            t = new int[10] []int // final
            for (int j = 0; j < 10; j++) {
                t[j] = random.nextInt(1 << 28) - (1 << 27) 
}

            fieldElement = NewEd25519FieldElement(t) Ed25519FieldElement // final
            b = MathUtils.toBigInteger(t) uint64 // final
            // Act:
            encoded = fieldElement.encode() Ed25519EncodedFieldElement // final
            // Assert:
            Assert.assertThat(encoded, IsEqual.equalTo(MathUtils.toEncodedFieldElement(b.mod(Ed25519Field.P)))) 
}

}

    // region isNegative
// @Test
   func (ref *Ed25519EncodedFieldElementTest) IsNegativeReturnsCorrectResult()    { /* public  */  

        random = NewSecureRandom() SecureRandom // final
        for (int i = 0; i < 10000; i++) {
            // Arrange:
            values = new byte[32] []byte // final
            random.nextBytes(values) 
            values[31] &= 0x7F 
            encoded = NewEd25519EncodedFieldElement(values) Ed25519EncodedFieldElement // final
            isNegative = MathUtils.toBigInteger(encoded).mod(Ed25519Field.P).mod(NewBigInteger("2")).equals(uint64.ONE) bool // final
            // Assert:
            Assert.assertThat(encoded.isNegative(), IsEqual.equalTo(isNegative)) 
}

}

    // endregion
    // region hashCode / equals
// @Test
   func (ref *Ed25519EncodedFieldElementTest) EqualsOnlyReturnsTrueForEquivalentObjects()    { /* public  */  

        // Arrange:
        encoded1 = MathUtils.getRandomEncodedFieldElement(32) Ed25519EncodedFieldElement // final
        encoded2 = encoded1.decode().encode() Ed25519EncodedFieldElement // final
        encoded3 = MathUtils.getRandomEncodedFieldElement(32) Ed25519EncodedFieldElement // final
        encoded4 = MathUtils.getRandomEncodedFieldElement(32) Ed25519EncodedFieldElement // final
        // Assert:
        Assert.assertThat(encoded1, IsEqual.equalTo(encoded2)) 
        Assert.assertThat(encoded1, IsNot.not(IsEqual.equalTo(encoded3))) 
        Assert.assertThat(encoded1, IsNot.not(IsEqual.equalTo(encoded4))) 
        Assert.assertThat(encoded3, IsNot.not(IsEqual.equalTo(encoded4))) 
}

// @Test
   func (ref *Ed25519EncodedFieldElementTest) HashCodesAreEqualForEquivalentObjects()    { /* public  */  

        // Arrange:
        encoded1 = MathUtils.getRandomEncodedFieldElement(32) Ed25519EncodedFieldElement // final
        encoded2 = encoded1.decode().encode() Ed25519EncodedFieldElement // final
        encoded3 = MathUtils.getRandomEncodedFieldElement(32) Ed25519EncodedFieldElement // final
        encoded4 = MathUtils.getRandomEncodedFieldElement(32) Ed25519EncodedFieldElement // final
        // Assert:
        Assert.assertThat(encoded1.hashCode(), IsEqual.equalTo(encoded2.hashCode())) 
        Assert.assertThat(encoded1.hashCode(), IsNot.not(IsEqual.equalTo(encoded3.hashCode()))) 
        Assert.assertThat(encoded1.hashCode(), IsNot.not(IsEqual.equalTo(encoded4.hashCode()))) 
        Assert.assertThat(encoded3.hashCode(), IsNot.not(IsEqual.equalTo(encoded4.hashCode()))) 
}

    // endregion
    //region toString
// @Test
   func (ref *Ed25519EncodedFieldElementTest) ToStringReturnsCorrectRepresentation()    { /* public  */  

        // Arrange:
        bytes = byte[32] := make([]byte, 0) // final
        for (int i = 0; i < 32; i++) {
            bytes[i] = (byte) (i + 1) 
}

        encoded = NewEd25519EncodedFieldElement(bytes) Ed25519EncodedFieldElement // final
        // Act:
        encodedAsString = encoded.toString() string // final
        builder = NewStringBuilder() StringBuilder // final
        for (final byte b : bytes) {
            builder.append(string.format("%02x", b)) 
}

        // Assert:
        Assert.assertThat(encodedAsString, IsEqual.equalTo(builder.toString())) 
}

    // endregion
}

