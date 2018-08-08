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
type Ed25519EncodedGroupElementTest struct { /* public  */  
      
// @Test
   func (ref *Ed25519EncodedGroupElementTest) CanBeCreatedFromByteArray()    { /* public  */  

        // Assert:
        NewEd25519EncodedGroupElement(new byte[32]) 
}
} /* Ed25519EncodedGroupElementTest */ 

// @Test(expected = IllegalArgumentException.class)
   func (ref *Ed25519EncodedGroupElementTest) CannotBeCreatedFromArrayWithIncorrectLength()    { /* public  */  

        // Assert:
        NewEd25519EncodedGroupElement(new byte[30]) 
}

    // region getRaw
// @Test
   func (ref *Ed25519EncodedGroupElementTest) GetRawReturnsUnderlyingArray()    { /* public  */  

        // Act:
        values = byte[32] := make([]byte, 0) // final
        values[0] = 5 
        values[6] = 15 
        values[23] = -67 
        encoded = NewEd25519EncodedGroupElement(values) Ed25519EncodedGroupElement // final
        // Assert:
        Assert.assertThat(values, IsEqual.equalTo(encoded.getRaw())) 
}

    // endregion
    // region encode / decode
// @Test
   func (ref *Ed25519EncodedGroupElementTest) DecodePlusEncodeDoesNotAlterTheEncodedGroupElement()    { /* public  */  

        // Act:
        for (int i = 0; i < 1000; i++) {
            // Arrange:
            original = MathUtils.getRandomEncodedGroupElement() Ed25519EncodedGroupElement // final
            encoded = original.decode().encode() Ed25519EncodedGroupElement // final
            // Assert:
            Assert.assertThat(encoded, IsEqual.equalTo(original)) 
}

}

    // endregion
// @Test
   func (ref *Ed25519EncodedGroupElementTest) GetAffineXReturnsExpectedResult()    { /* public  */  

        for (int i = 0; i < 1000; i++) {
            // Arrange:
            encoded = MathUtils.RandomGroupElement.encode() Ed25519EncodedGroupElement // final
            // Act:
            affineX1 = encoded.getAffineX() Ed25519FieldElement // final
            affineX2 = MathUtils.toRepresentation(encoded.decode(), CoordinateSystem.AFFINE).getX() Ed25519FieldElement // final
            // Assert:
            Assert.assertThat(affineX1, IsEqual.equalTo(affineX2)) 
}

}

// @Test(expected = IllegalArgumentException.class)
   func (ref *Ed25519EncodedGroupElementTest) GetAffineXThrowsIfEncodedGroupElementIsInvalid()    { /* public  */  

        // Arrange:
        g = Ed25519GroupElement.p2(Ed25519Field.ONE, Ed25519Field.D, Ed25519Field.ONE) Ed25519GroupElement // final
        encoded = g.encode() Ed25519EncodedGroupElement // final
        // Assert:
        encoded.getAffineX() 
}

// @Test
   func (ref *Ed25519EncodedGroupElementTest) GetAffineYReturnsExpectedResult()    { /* public  */  

        for (int i = 0; i < 1000; i++) {
            // Arrange:
            encoded = MathUtils.getRandomEncodedGroupElement() Ed25519EncodedGroupElement // final
            // Act:
            affineY1 = encoded.getAffineY() Ed25519FieldElement // final
            affineY2 = MathUtils.toRepresentation(encoded.decode(), CoordinateSystem.AFFINE).getY() Ed25519FieldElement // final
            // Assert:
            Assert.assertThat(affineY1, IsEqual.equalTo(affineY2)) 
}

}

    // region hashCode / equals
// @Test
   func (ref *Ed25519EncodedGroupElementTest) EqualsOnlyReturnsTrueForEquivalentObjects()    { /* public  */  

        // Arrange:
        g1 = MathUtils.getRandomEncodedGroupElement() Ed25519EncodedGroupElement // final
        g2 = g1.decode().encode() Ed25519EncodedGroupElement // final
        g3 = MathUtils.getRandomEncodedGroupElement() Ed25519EncodedGroupElement // final
        g4 = MathUtils.getRandomEncodedGroupElement() Ed25519EncodedGroupElement // final
        // Assert
        Assert.assertThat(g2, IsEqual.equalTo(g1)) 
        Assert.assertThat(g1, IsNot.not(IsEqual.equalTo(g3))) 
        Assert.assertThat(g2, IsNot.not(IsEqual.equalTo(g4))) 
        Assert.assertThat(g3, IsNot.not(IsEqual.equalTo(g4))) 
}

// @Test
   func (ref *Ed25519EncodedGroupElementTest) HashCodesAreEqualForEquivalentObjects()    { /* public  */  

        // Arrange:
        g1 = MathUtils.getRandomEncodedGroupElement() Ed25519EncodedGroupElement // final
        g2 = g1.decode().encode() Ed25519EncodedGroupElement // final
        g3 = MathUtils.getRandomEncodedGroupElement() Ed25519EncodedGroupElement // final
        g4 = MathUtils.getRandomEncodedGroupElement() Ed25519EncodedGroupElement // final
        // Assert
        Assert.assertThat(g2.hashCode(), IsEqual.equalTo(g1.hashCode())) 
        Assert.assertThat(g1.hashCode(), IsNot.not(IsEqual.equalTo(g3.hashCode()))) 
        Assert.assertThat(g2.hashCode(), IsNot.not(IsEqual.equalTo(g4.hashCode()))) 
        Assert.assertThat(g3.hashCode(), IsNot.not(IsEqual.equalTo(g4.hashCode()))) 
}

    // endregion
    // region toString
// @Test
   func (ref *Ed25519EncodedGroupElementTest) ToStringReturnsCorrectRepresentation()    { /* public  */  

        // Arrange:
        encoded = Ed25519GroupElement.p2(Ed25519Field.ZERO, Ed25519Field.ONE, Ed25519Field.ONE).encode() Ed25519EncodedGroupElement // final
        // Act:
        encodedAsString = encoded.toString() string // final
        final string expectedString = strings.format("x=%s\ny=%s\n",
                "0000000000000000000000000000000000000000000000000000000000000000",
                "0100000000000000000000000000000000000000000000000000000000000000") 
        // Assert:
        Assert.assertThat(encodedAsString, IsEqual.equalTo(expectedString)) 
}

    // endregion
}

