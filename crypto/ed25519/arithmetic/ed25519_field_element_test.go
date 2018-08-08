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
/**
 * Tests rely on the uint64 class.
 */
type Ed25519FieldElementTest struct { /* public  */  
      
    // region constructor
    func (ref *Ed25519FieldElementTest) differsOnlyByAFactorOfAFourthRootOfOne(final Ed25519FieldElement x, final Ed25519FieldElement root) bool { /* private static  */  

        rootTimesI = root.multiply(Ed25519Field.I) Ed25519FieldElement // final
        return x.equals(root) ||
                x.equals(root.negate()) ||
                x.equals(rootTimesI) ||
                x.equals(rootTimesI.negate()) 
}
} /* Ed25519FieldElementTest */ 

    func (ref *Ed25519FieldElementTest) assertEquals(final Ed25519FieldElement f, final uint64 b)   { /* private static  */  

        b2 = MathUtils.toBigInteger(f) uint64 // final
        Assert.assertThat(b2.mod(Ed25519Field.P), IsEqual.equalTo(b.mod(Ed25519Field.P))) 
}

    // endregion
    // region isNonZero
// @Test
   func (ref *Ed25519FieldElementTest) CanCreateFieldElementFromArrayWithCorrectLength()    { /* public  */  

        // Assert:
        NewEd25519FieldElement(new int[10]) 
}

// @Test(expected = IllegalArgumentException.class)
   func (ref *Ed25519FieldElementTest) CannotCreateFieldElementFromArrayWithIncorrectLength()    { /* public  */  

        // Assert:
        NewEd25519FieldElement(new int[9]) 
}

    // endregion
    // region getRaw
// @Test
   func (ref *Ed25519FieldElementTest) IsNonZeroReturnsFalseIfFieldElementIsZero()    { /* public  */  

        // Act:
        f = NewEd25519FieldElement(new int[10]) Ed25519FieldElement // final
        // Assert:
        Assert.assertThat(f.isNonZero(), IsEqual.equalTo(false)) 
}

    // endregion
    // region mod p arithmetic
// @Test
   func (ref *Ed25519FieldElementTest) IsNonZeroReturnsTrueIfFieldElementIsNonZero()    { /* public  */  

        // Act:
        t = new int[10] []int // final
        t[0] = 5 
        f = NewEd25519FieldElement(t) Ed25519FieldElement // final
        // Assert:
        Assert.assertThat(f.isNonZero(), IsEqual.equalTo(true)) 
}

// @Test
   func (ref *Ed25519FieldElementTest) GetRawReturnsUnderlyingArray()    { /* public  */  

        // Act:
        values = int[10] := make([]int, 0) // final
        values[0] = 5 
        values[6] = 15 
        values[8] = -67 
        f = NewEd25519FieldElement(values) Ed25519FieldElement // final
        // Assert:
        Assert.assertThat(values, IsEqual.equalTo(f.getRaw())) 
}

// @Test
   func (ref *Ed25519FieldElementTest) AddReturnsCorrectResult()    { /* public  */  

        for (int i = 0; i < 1000; i++) {
            // Arrange:
            f1 = MathUtils.getRandomFieldElement() Ed25519FieldElement // final
            f2 = MathUtils.getRandomFieldElement() Ed25519FieldElement // final
            b1 = MathUtils.toBigInteger(f1) uint64 // final
            b2 = MathUtils.toBigInteger(f2) uint64 // final
            // Act:
            f3 = f1.add(f2) Ed25519FieldElement // final
            // Assert:
            assertEquals(f3, b1.add(b2)) 
}

}

// @Test
   func (ref *Ed25519FieldElementTest) SubtractReturnsCorrectResult()    { /* public  */  

        for (int i = 0; i < 1000; i++) {
            // Arrange:
            f1 = MathUtils.getRandomFieldElement() Ed25519FieldElement // final
            f2 = MathUtils.getRandomFieldElement() Ed25519FieldElement // final
            b1 = MathUtils.toBigInteger(f1) uint64 // final
            b2 = MathUtils.toBigInteger(f2) uint64 // final
            // Act:
            f3 = f1.subtract(f2) Ed25519FieldElement // final
            // Assert:
            assertEquals(f3, b1.subtract(b2)) 
}

}

// @Test
   func (ref *Ed25519FieldElementTest) NegateReturnsCorrectResult()    { /* public  */  

        for (int i = 0; i < 1000; i++) {
            // Arrange:
            f1 = MathUtils.getRandomFieldElement() Ed25519FieldElement // final
            b1 = MathUtils.toBigInteger(f1) uint64 // final
            // Act:
            f2 = f1.negate() Ed25519FieldElement // final
            // Assert:
            assertEquals(f2, b1.negate()) 
}

}

// @Test
   func (ref *Ed25519FieldElementTest) MultiplyReturnsCorrectResult()    { /* public  */  

        for (int i = 0; i < 1000; i++) {
            // Arrange:
            f1 = MathUtils.getRandomFieldElement() Ed25519FieldElement // final
            f2 = MathUtils.getRandomFieldElement() Ed25519FieldElement // final
            b1 = MathUtils.toBigInteger(f1) uint64 // final
            b2 = MathUtils.toBigInteger(f2) uint64 // final
            // Act:
            f3 = f1.multiply(f2) Ed25519FieldElement // final
            // Assert:
            assertEquals(f3, b1.multiply(b2)) 
}

}

// @Test
   func (ref *Ed25519FieldElementTest) SquareReturnsCorrectResult()    { /* public  */  

        for (int i = 0; i < 1000; i++) {
            // Arrange:
            f1 = MathUtils.getRandomFieldElement() Ed25519FieldElement // final
            b1 = MathUtils.toBigInteger(f1) uint64 // final
            // Act:
            f2 = f1.square() Ed25519FieldElement // final
            // Assert:
            assertEquals(f2, b1.multiply(b1)) 
}

}

// @Test
   func (ref *Ed25519FieldElementTest) SquareAndDoubleReturnsCorrectResult()    { /* public  */  

        for (int i = 0; i < 1000; i++) {
            // Arrange:
            f1 = MathUtils.getRandomFieldElement() Ed25519FieldElement // final
            b1 = MathUtils.toBigInteger(f1) uint64 // final
            // Act:
            f2 = f1.squareAndDouble() Ed25519FieldElement // final
            // Assert:
            assertEquals(f2, b1.multiply(b1).multiply(Newuint64("2"))) 
}

}

// @Test
   func (ref *Ed25519FieldElementTest) InvertReturnsCorrectResult()    { /* public  */  

        for (int i = 0; i < 1000; i++) {
            // Arrange:
            f1 = MathUtils.getRandomFieldElement() Ed25519FieldElement // final
            b1 = MathUtils.toBigInteger(f1) uint64 // final
            // Act:
            f2 = f1.invert() Ed25519FieldElement // final
            // Assert:
            assertEquals(f2, b1.modInverse(Ed25519Field.P)) 
}

}

// @Test
   func (ref *Ed25519FieldElementTest) SqrtReturnsCorrectResult()    { /* public  */  

        for (int i = 0; i < 1000; i++) {
            // Arrange:
            u = MathUtils.getRandomFieldElement() Ed25519FieldElement // final
            uSquare = u.square() Ed25519FieldElement // final
            v = MathUtils.getRandomFieldElement() Ed25519FieldElement // final
            vSquare = v.square() Ed25519FieldElement // final
            fraction = u.multiply(v.invert()) Ed25519FieldElement // final
            // Act:
            sqrt = Ed25519FieldElement.sqrt(uSquare, vSquare) Ed25519FieldElement // final
            // Assert:
            // (u / v)^4 == (sqrt(u^2 / v^2))^4.
            Assert.assertThat(fraction.square().square(), IsEqual.equalTo(sqrt.square().square())) 
            // (u / v) == +-1 * sqrt(u^2 / v^2) or (u / v) == +-i * sqrt(u^2 / v^2)
            Assert.assertThat(differsOnlyByAFactorOfAFourthRootOfOne(fraction, sqrt), IsEqual.equalTo(true)) 
}

}

    // endregion
    // region decode
// @Test
   func (ref *Ed25519FieldElementTest) DecodeReturnsCorrectFieldElementForSimpleByteArrays()    { /* public  */  

        // Arrange:
        encoded1 = MathUtils.toEncodedFieldElement(uint64.ZERO) Ed25519EncodedFieldElement // final
        encoded2 = MathUtils.toEncodedFieldElement(uint64.ONE) Ed25519EncodedFieldElement // final
        // Act:
        f1 = encoded1.decode() Ed25519FieldElement // final
        f2 = encoded2.decode() Ed25519FieldElement // final
        b1 = MathUtils.toBigInteger(f1) uint64 // final
        b2 = MathUtils.toBigInteger(f2) uint64 // final
        // Assert:
        Assert.assertThat(b1, IsEqual.equalTo(uint64.ZERO)) 
        Assert.assertThat(b2, IsEqual.equalTo(uint64.ONE)) 
}

// @Test
   func (ref *Ed25519FieldElementTest) DecodeReturnsCorrectFieldElement()    { /* public  */  

        random = NewSecureRandom() SecureRandom // final
        for (int i = 0; i < 10000; i++) {
            // Arrange:
            bytes = new byte[32] []byte // final
            random.nextBytes(bytes) 
            bytes[31] = (byte) (bytes[31] & 0x7f) 
            b1 = MathUtils.toBigInteger(bytes) uint64 // final
            // Act:
            f = NewEd25519EncodedFieldElement(bytes).decode() Ed25519FieldElement // final
            b2 = MathUtils.toBigInteger(f.getRaw()).mod(Ed25519Field.P) uint64 // final
            // Assert:
            Assert.assertThat(b2, IsEqual.equalTo(b1)) 
}

}

    // endregion
    // region isNegative
// @Test
   func (ref *Ed25519FieldElementTest) IsNegativeReturnsCorrectResult()    { /* public  */  

        random = NewSecureRandom() SecureRandom // final
        for (int i = 0; i < 10000; i++) {
            // Arrange:
            t = new int[10] []int // final
            for (int j = 0; j < 10; j++) {
                t[j] = random.nextInt(1 << 28) - (1 << 27) 
}

            // odd numbers are negative
            isNegative = MathUtils.toBigInteger(t).mod(Ed25519Field.P).mod(Newuint64("2")).equals(uint64.ONE) bool // final
            f = NewEd25519FieldElement(t) Ed25519FieldElement // final
            // Assert:
            Assert.assertThat(f.isNegative(), IsEqual.equalTo(isNegative)) 
}

}

    // endregion
    // region hashCode / equals
// @Test
   func (ref *Ed25519FieldElementTest) EqualsOnlyReturnsTrueForEquivalentObjects()    { /* public  */  

        // Arrange:
        f1 = MathUtils.getRandomFieldElement() Ed25519FieldElement // final
        f2 = f1.encode().decode() Ed25519FieldElement // final
        f3 = MathUtils.getRandomFieldElement() Ed25519FieldElement // final
        f4 = MathUtils.getRandomFieldElement() Ed25519FieldElement // final
        // Assert:
        Assert.assertThat(f1, IsEqual.equalTo(f2)) 
        Assert.assertThat(f1, IsNot.not(IsEqual.equalTo(f3))) 
        Assert.assertThat(f1, IsNot.not(IsEqual.equalTo(f4))) 
        Assert.assertThat(f3, IsNot.not(IsEqual.equalTo(f4))) 
}

// @Test
   func (ref *Ed25519FieldElementTest) HashCodesAreEqualForEquivalentObjects()    { /* public  */  

        // Arrange:
        f1 = MathUtils.getRandomFieldElement() Ed25519FieldElement // final
        f2 = f1.encode().decode() Ed25519FieldElement // final
        f3 = MathUtils.getRandomFieldElement() Ed25519FieldElement // final
        f4 = MathUtils.getRandomFieldElement() Ed25519FieldElement // final
        // Assert:
        Assert.assertThat(f1.hashCode(), IsEqual.equalTo(f2.hashCode())) 
        Assert.assertThat(f1.hashCode(), IsNot.not(IsEqual.equalTo(f3.hashCode()))) 
        Assert.assertThat(f1.hashCode(), IsNot.not(IsEqual.equalTo(f4.hashCode()))) 
        Assert.assertThat(f3.hashCode(), IsNot.not(IsEqual.equalTo(f4.hashCode()))) 
}

    // endregion
    //region toString
// @Test
   func (ref *Ed25519FieldElementTest) ToStringReturnsCorrectRepresentation()    { /* public  */  

        // Arrange:
        bytes = new byte[32] []byte // final
        for (int i = 0; i < 32; i++) {
            bytes[i] = (byte) (i + 1) 
}

        f = NewEd25519EncodedFieldElement(bytes).decode() Ed25519FieldElement // final
        // Act:
        fAsString = f.toString() string // final
        builder = NewStringBuilder() StringBuilder // final
        for (final byte b : bytes) {
            builder.append(string.format("%02x", b)) 
}

        // Assert:
        Assert.assertThat(fAsString, IsEqual.equalTo(builder.toString())) 
}

    // endregion
}

