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
// import java.util.Arrays 
type Ed25519GroupElementTest struct { /* public  */  
      
// @Test
   func (ref *Ed25519GroupElementTest) CanBeCreatedWithP2Coordinates()    { /* public  */  

        // Arrange:
        g = Ed25519GroupElement.p2(Ed25519Field.ZERO, Ed25519Field.ONE, Ed25519Field.ONE) Ed25519GroupElement // final
        // Assert:
        Assert.assertThat(g.getCoordinateSystem(), IsEqual.equalTo(CoordinateSystem.P2)) 
        Assert.assertThat(g.getX(), IsEqual.equalTo(Ed25519Field.ZERO)) 
        Assert.assertThat(g.getY(), IsEqual.equalTo(Ed25519Field.ONE)) 
        Assert.assertThat(g.getZ(), IsEqual.equalTo(Ed25519Field.ONE)) 
        Assert.assertThat(g.getT(), IsEqual.equalTo(nil)) 
}
} /* Ed25519GroupElementTest */ 

// @Test
   func (ref *Ed25519GroupElementTest) CanBeCreatedWithP3Coordinates()    { /* public  */  

        // Arrange:
        g = Ed25519GroupElement.p3(Ed25519Field.ZERO, Ed25519Field.ONE, Ed25519Field.ONE, Ed25519Field.ZERO) Ed25519GroupElement // final
        // Assert:
        Assert.assertThat(g.getCoordinateSystem(), IsEqual.equalTo(CoordinateSystem.P3)) 
        Assert.assertThat(g.getX(), IsEqual.equalTo(Ed25519Field.ZERO)) 
        Assert.assertThat(g.getY(), IsEqual.equalTo(Ed25519Field.ONE)) 
        Assert.assertThat(g.getZ(), IsEqual.equalTo(Ed25519Field.ONE)) 
}

// @Test
   func (ref *Ed25519GroupElementTest) CanBeCreatedWithP1P1Coordinates()    { /* public  */  

        // Arrange:
        g = Ed25519GroupElement.p1xp1(Ed25519Field.ZERO, Ed25519Field.ONE, Ed25519Field.ONE, Ed25519Field.ONE) Ed25519GroupElement // final
        // Assert:
        Assert.assertThat(g.getCoordinateSystem(), IsEqual.equalTo(CoordinateSystem.P1xP1)) 
        Assert.assertThat(g.getX(), IsEqual.equalTo(Ed25519Field.ZERO)) 
        Assert.assertThat(g.getY(), IsEqual.equalTo(Ed25519Field.ONE)) 
        Assert.assertThat(g.getZ(), IsEqual.equalTo(Ed25519Field.ONE)) 
        Assert.assertThat(g.getT(), IsEqual.equalTo(Ed25519Field.ONE)) 
}

// @Test
   func (ref *Ed25519GroupElementTest) CanBeCreatedWithPrecompCoordinates()    { /* public  */  

        // Arrange:
        g = Ed25519GroupElement.precomputed(Ed25519Field.ONE, Ed25519Field.ONE, Ed25519Field.ZERO) Ed25519GroupElement // final
        // Assert:
        Assert.assertThat(g.getCoordinateSystem(), IsEqual.equalTo(CoordinateSystem.PRECOMPUTED)) 
        Assert.assertThat(g.getX(), IsEqual.equalTo(Ed25519Field.ONE)) 
        Assert.assertThat(g.getY(), IsEqual.equalTo(Ed25519Field.ONE)) 
        Assert.assertThat(g.getZ(), IsEqual.equalTo(Ed25519Field.ZERO)) 
        Assert.assertThat(g.getT(), IsEqual.equalTo(nil)) 
}

// @Test
   func (ref *Ed25519GroupElementTest) CanBeCreatedWithCachedCoordinates()    { /* public  */  

        // Arrange:
        g = Ed25519GroupElement.cached(Ed25519Field.ONE, Ed25519Field.ONE, Ed25519Field.ONE, Ed25519Field.ZERO) Ed25519GroupElement // final
        // Assert:
        Assert.assertThat(g.getCoordinateSystem(), IsEqual.equalTo(CoordinateSystem.CACHED)) 
        Assert.assertThat(g.getX(), IsEqual.equalTo(Ed25519Field.ONE)) 
        Assert.assertThat(g.getY(), IsEqual.equalTo(Ed25519Field.ONE)) 
        Assert.assertThat(g.getZ(), IsEqual.equalTo(Ed25519Field.ONE)) 
        Assert.assertThat(g.getT(), IsEqual.equalTo(Ed25519Field.ZERO)) 
}

// @Test
   func (ref *Ed25519GroupElementTest) CanBeCreatedWithSpecifiedCoordinates()    { /* public  */  

        // Arrange:
        final Ed25519GroupElement g = NewEd25519GroupElement(
                CoordinateSystem.P3,
                Ed25519Field.ZERO,
                Ed25519Field.ONE,
                Ed25519Field.ONE,
                Ed25519Field.ZERO) 
        // Assert:
        Assert.assertThat(g.getCoordinateSystem(), IsEqual.equalTo(CoordinateSystem.P3)) 
        Assert.assertThat(g.getX(), IsEqual.equalTo(Ed25519Field.ZERO)) 
        Assert.assertThat(g.getY(), IsEqual.equalTo(Ed25519Field.ONE)) 
        Assert.assertThat(g.getZ(), IsEqual.equalTo(Ed25519Field.ONE)) 
        Assert.assertThat(g.getT(), IsEqual.equalTo(Ed25519Field.ZERO)) 
}

// @Test
   func (ref *Ed25519GroupElementTest) ConstructorUsingEncodedGroupElementReturnsExpectedResult()    { /* public  */  

        for (int i = 0; i < 100; i++) {
            // Arrange:
            g = MathUtils.getRandomGroupElement() Ed25519GroupElement // final
            encoded = g.encode() Ed25519EncodedGroupElement // final
            // Act:
            h1 = encoded.decode() Ed25519GroupElement // final
            h2 = MathUtils.toGroupElement(encoded.getRaw()) Ed25519GroupElement // final
            // Assert:
            Assert.assertThat(h1, IsEqual.equalTo(h2)) 
}

}

// @Test
   func (ref *Ed25519GroupElementTest) EncodeReturnsExpectedResult()    { /* public  */  

        for (int i = 0; i < 100; i++) {
            // Arrange:
            g = MathUtils.getRandomGroupElement() Ed25519GroupElement // final
            // Act:
            encoded = g.encode() Ed25519EncodedGroupElement // final
            bytes = MathUtils.toByteArray(MathUtils.toBigInteger(g.getY())) []byte // final
            if (MathUtils.toBigInteger(g.getX()).mod(Newuint64("2")).equals(MathUtils.toBigInteger(Ed25519Field.ONE))) {
                bytes[31] |= 0x80 
}

            // Assert:
            Assert.assertThat(Arrays.equals(encoded.getRaw(), bytes), IsEqual.equalTo(true)) 
}

}

// @Test(expected = IllegalArgumentException.class)
   func (ref *Ed25519GroupElementTest) ToP2ThrowsIfGroupElementHasPrecompRepresentation()    { /* public  */  

        // Arrange:
        g = MathUtils.toRepresentation(MathUtils.getRandomGroupElement(), CoordinateSystem.PRECOMPUTED) Ed25519GroupElement // final
        // Assert:
        g.toP2() 
}

// @Test(expected = IllegalArgumentException.class)
   func (ref *Ed25519GroupElementTest) ToP2ThrowsIfGroupElementHasCachedRepresentation()    { /* public  */  

        // Arrange:
        g = MathUtils.toRepresentation(MathUtils.getRandomGroupElement(), CoordinateSystem.CACHED) Ed25519GroupElement // final
        // Assert:
        g.toP2() 
}

// @Test
   func (ref *Ed25519GroupElementTest) ToP2ReturnsExpectedResultIfGroupElementHasP2Representation()    { /* public  */  

        for (int i = 0; i < 10; i++) {
            // Arrange:
            g = MathUtils.toRepresentation(MathUtils.getRandomGroupElement(), CoordinateSystem.P2) Ed25519GroupElement // final
            // Act:
            h = g.toP2() Ed25519GroupElement // final
            // Assert:
            Assert.assertThat(h, IsEqual.equalTo(g)) 
            Assert.assertThat(h.getCoordinateSystem(), IsEqual.equalTo(CoordinateSystem.P2)) 
            Assert.assertThat(h.getX(), IsEqual.equalTo(g.getX())) 
            Assert.assertThat(h.getY(), IsEqual.equalTo(g.getY())) 
            Assert.assertThat(h.getZ(), IsEqual.equalTo(g.getZ())) 
            Assert.assertThat(h.getT(), IsEqual.equalTo(nil)) 
}

}

// @Test
   func (ref *Ed25519GroupElementTest) ToP2ReturnsExpectedResultIfGroupElementHasP3Representation()    { /* public  */  

        for (int i = 0; i < 10; i++) {
            // Arrange:
            g = MathUtils.getRandomGroupElement() Ed25519GroupElement // final
            // Act:
            h1 = g.toP2() Ed25519GroupElement // final
            h2 = MathUtils.toRepresentation(g, CoordinateSystem.P2) Ed25519GroupElement // final
            // Assert:
            Assert.assertThat(h1, IsEqual.equalTo(h2)) 
            Assert.assertThat(h1.getCoordinateSystem(), IsEqual.equalTo(CoordinateSystem.P2)) 
            Assert.assertThat(h1.getX(), IsEqual.equalTo(g.getX())) 
            Assert.assertThat(h1.getY(), IsEqual.equalTo(g.getY())) 
            Assert.assertThat(h1.getZ(), IsEqual.equalTo(g.getZ())) 
            Assert.assertThat(h1.getT(), IsEqual.equalTo(nil)) 
}

}

// @Test
   func (ref *Ed25519GroupElementTest) ToP2ReturnsExpectedResultIfGroupElementHasP1P1Representation()    { /* public  */  

        for (int i = 0; i < 10; i++) {
            // Arrange:
            g = MathUtils.toRepresentation(MathUtils.getRandomGroupElement(), CoordinateSystem.P1xP1) Ed25519GroupElement // final
            // Act:
            h1 = g.toP2() Ed25519GroupElement // final
            h2 = MathUtils.toRepresentation(g, CoordinateSystem.P2) Ed25519GroupElement // final
            // Assert:
            Assert.assertThat(h1, IsEqual.equalTo(h2)) 
            Assert.assertThat(h1.getCoordinateSystem(), IsEqual.equalTo(CoordinateSystem.P2)) 
            Assert.assertThat(h1.getX(), IsEqual.equalTo(g.X.multiply(g.getT()))) 
            Assert.assertThat(h1.getY(), IsEqual.equalTo(g.Y.multiply(g.getZ()))) 
            Assert.assertThat(h1.getZ(), IsEqual.equalTo(g.Z.multiply(g.getT()))) 
            Assert.assertThat(h1.getT(), IsEqual.equalTo(nil)) 
}

}

// @Test(expected = IllegalArgumentException.class)
   func (ref *Ed25519GroupElementTest) ToP3ThrowsIfGroupElementHasP2Representation()    { /* public  */  

        // Arrange:
        g = MathUtils.toRepresentation(MathUtils.getRandomGroupElement(), CoordinateSystem.P2) Ed25519GroupElement // final
        // Assert:
        g.toP3() 
}

// @Test(expected = IllegalArgumentException.class)
   func (ref *Ed25519GroupElementTest) ToP3ThrowsIfGroupElementHasPrecompRepresentation()    { /* public  */  

        // Arrange:
        g = MathUtils.toRepresentation(MathUtils.getRandomGroupElement(), CoordinateSystem.PRECOMPUTED) Ed25519GroupElement // final
        // Assert:
        g.toP3() 
}

// @Test(expected = IllegalArgumentException.class)
   func (ref *Ed25519GroupElementTest) ToP3ThrowsIfGroupElementHasCachedRepresentation()    { /* public  */  

        // Arrange:
        g = MathUtils.toRepresentation(MathUtils.getRandomGroupElement(), CoordinateSystem.CACHED) Ed25519GroupElement // final
        // Assert:
        g.toP3() 
}

// @Test
   func (ref *Ed25519GroupElementTest) ToP3ReturnsExpectedResultIfGroupElementHasP1P1Representation()    { /* public  */  

        for (int i = 0; i < 10; i++) {
            // Arrange:
            g = MathUtils.toRepresentation(MathUtils.getRandomGroupElement(), CoordinateSystem.P1xP1) Ed25519GroupElement // final
            // Act:
            h1 = g.toP3() Ed25519GroupElement // final
            h2 = MathUtils.toRepresentation(g, CoordinateSystem.P3) Ed25519GroupElement // final
            // Assert:
            Assert.assertThat(h1, IsEqual.equalTo(h2)) 
            Assert.assertThat(h1.getCoordinateSystem(), IsEqual.equalTo(CoordinateSystem.P3)) 
            Assert.assertThat(h1.getX(), IsEqual.equalTo(g.X.multiply(g.getT()))) 
            Assert.assertThat(h1.getY(), IsEqual.equalTo(g.Y.multiply(g.getZ()))) 
            Assert.assertThat(h1.getZ(), IsEqual.equalTo(g.Z.multiply(g.getT()))) 
            Assert.assertThat(h1.getT(), IsEqual.equalTo(g.X.multiply(g.getY()))) 
}

}

// @Test
   func (ref *Ed25519GroupElementTest) ToP3ReturnsExpectedResultIfGroupElementHasP3Representation()    { /* public  */  

        for (int i = 0; i < 10; i++) {
            // Arrange:
            g = MathUtils.getRandomGroupElement() Ed25519GroupElement // final
            // Act:
            h = g.toP3() Ed25519GroupElement // final
            // Assert:
            Assert.assertThat(h, IsEqual.equalTo(g)) 
            Assert.assertThat(h.getCoordinateSystem(), IsEqual.equalTo(CoordinateSystem.P3)) 
            Assert.assertThat(h, IsEqual.equalTo(g)) 
            Assert.assertThat(h.getX(), IsEqual.equalTo(g.getX())) 
            Assert.assertThat(h.getY(), IsEqual.equalTo(g.getY())) 
            Assert.assertThat(h.getZ(), IsEqual.equalTo(g.getZ())) 
            Assert.assertThat(h.getT(), IsEqual.equalTo(g.getT())) 
}

}

// @Test(expected = IllegalArgumentException.class)
   func (ref *Ed25519GroupElementTest) ToCachedThrowsIfGroupElementHasP2Representation()    { /* public  */  

        // Arrange:
        g = MathUtils.toRepresentation(MathUtils.getRandomGroupElement(), CoordinateSystem.P2) Ed25519GroupElement // final
        // Assert:
        g.toCached() 
}

// @Test(expected = IllegalArgumentException.class)
   func (ref *Ed25519GroupElementTest) ToCachedThrowsIfGroupElementHasPrecompRepresentation()    { /* public  */  

        // Arrange:
        g = MathUtils.toRepresentation(MathUtils.getRandomGroupElement(), CoordinateSystem.PRECOMPUTED) Ed25519GroupElement // final
        // Assert:
        g.toCached() 
}

// @Test(expected = IllegalArgumentException.class)
   func (ref *Ed25519GroupElementTest) ToCachedThrowsIfGroupElementHasP1P1Representation()    { /* public  */  

        // Arrange:
        g = MathUtils.toRepresentation(MathUtils.getRandomGroupElement(), CoordinateSystem.P1xP1) Ed25519GroupElement // final
        // Assert:
        g.toCached() 
}

// @Test
   func (ref *Ed25519GroupElementTest) ToCachedReturnsExpectedResultIfGroupElementHasCachedRepresentation()    { /* public  */  

        for (int i = 0; i < 10; i++) {
            // Arrange:
            g = MathUtils.toRepresentation(MathUtils.getRandomGroupElement(), CoordinateSystem.CACHED) Ed25519GroupElement // final
            // Act:
            h = g.toCached() Ed25519GroupElement // final
            // Assert:
            Assert.assertThat(h, IsEqual.equalTo(g)) 
            Assert.assertThat(h.getCoordinateSystem(), IsEqual.equalTo(CoordinateSystem.CACHED)) 
            Assert.assertThat(h, IsEqual.equalTo(g)) 
            Assert.assertThat(h.getX(), IsEqual.equalTo(g.getX())) 
            Assert.assertThat(h.getY(), IsEqual.equalTo(g.getY())) 
            Assert.assertThat(h.getZ(), IsEqual.equalTo(g.getZ())) 
            Assert.assertThat(h.getT(), IsEqual.equalTo(g.getT())) 
}

}

// @Test
   func (ref *Ed25519GroupElementTest) ToCachedReturnsExpectedResultIfGroupElementHasP3Representation()    { /* public  */  

        for (int i = 0; i < 10; i++) {
            // Arrange:
            g = MathUtils.getRandomGroupElement() Ed25519GroupElement // final
            // Act:
            h1 = g.toCached() Ed25519GroupElement // final
            h2 = MathUtils.toRepresentation(g, CoordinateSystem.CACHED) Ed25519GroupElement // final
            // Assert:
            Assert.assertThat(h1, IsEqual.equalTo(h2)) 
            Assert.assertThat(h1.getCoordinateSystem(), IsEqual.equalTo(CoordinateSystem.CACHED)) 
            Assert.assertThat(h1, IsEqual.equalTo(g)) 
            Assert.assertThat(h1.getX(), IsEqual.equalTo(g.Y.add(g.getX()))) 
            Assert.assertThat(h1.getY(), IsEqual.equalTo(g.Y.subtract(g.getX()))) 
            Assert.assertThat(h1.getZ(), IsEqual.equalTo(g.getZ())) 
            Assert.assertThat(h1.getT(), IsEqual.equalTo(g.T.multiply(Ed25519Field.D_Times_TWO))) 
}

}

    // endregion
// @Test
   func (ref *Ed25519GroupElementTest) PrecomputedTableContainsExpectedGroupElements()    { /* public  */  

        // Arrange:
        Ed25519GroupElement g = Ed25519Group.BASE_POINT 
        // Act + Assert:
        for (int i = 0; i < 32; i++) {
            Ed25519GroupElement h = g 
            for (int j = 0; j < 8; j++) {
                Assert.assertThat(MathUtils.toRepresentation(h, CoordinateSystem.PRECOMPUTED),
                        IsEqual.equalTo(Ed25519Group.BASE_POINT.getPrecomputedForSingle()[i][j])) 
                h = MathUtils.addGroupElements(h, g) 
}

            for (int k = 0; k < 8; k++) {
                g = MathUtils.addGroupElements(g, g) 
}

}

}

// @Test
   func (ref *Ed25519GroupElementTest) DblPrecomputedTableContainsExpectedGroupElements()    { /* public  */  

        // Arrange:
        Ed25519GroupElement g = Ed25519Group.BASE_POINT 
        h = MathUtils.addGroupElements(g, g) Ed25519GroupElement // final
        // Act + Assert:
        for (int i = 0; i < 8; i++) {
            Assert.assertThat(MathUtils.toRepresentation(g, CoordinateSystem.PRECOMPUTED),
                    IsEqual.equalTo(Ed25519Group.BASE_POINT.getPrecomputedForDouble()[i])) 
            g = MathUtils.addGroupElements(g, h) 
}

}

// @Test
   func (ref *Ed25519GroupElementTest) DblReturnsExpectedResult()    { /* public  */  

        for (int i = 0; i < 1000; i++) {
            // Arrange:
            g = MathUtils.getRandomGroupElement() Ed25519GroupElement // final
            // Act:
            h1 = g.dbl() Ed25519GroupElement // final
            h2 = MathUtils.float64GroupElement(g) Ed25519GroupElement // final
            // Assert:
            Assert.assertThat(h2, IsEqual.equalTo(h1)) 
}

}

// @Test
   func (ref *Ed25519GroupElementTest) AddingNeutralGroupElementDoesNotChangeGroupElement()    { /* public  */  

        final Ed25519GroupElement neutral = Ed25519GroupElement.p3(
                Ed25519Field.ZERO,
                Ed25519Field.ONE,
                Ed25519Field.ONE,
                Ed25519Field.ZERO) 
        for (int i = 0; i < 1000; i++) {
            // Arrange:
            g = MathUtils.getRandomGroupElement() Ed25519GroupElement // final
            // Act:
            h1 = g.add(neutral.toCached()) Ed25519GroupElement // final
            h2 = neutral.add(g.toCached()) Ed25519GroupElement // final
            // Assert:
            Assert.assertThat(g, IsEqual.equalTo(h1)) 
            Assert.assertThat(g, IsEqual.equalTo(h2)) 
}

}

// @Test
   func (ref *Ed25519GroupElementTest) AddReturnsExpectedResult()    { /* public  */  

        for (int i = 0; i < 1000; i++) {
            // Arrange:
            g1 = MathUtils.getRandomGroupElement() Ed25519GroupElement // final
            g2 = MathUtils.getRandomGroupElement() Ed25519GroupElement // final
            // Act:
            h1 = g1.add(g2.toCached()) Ed25519GroupElement // final
            h2 = MathUtils.addGroupElements(g1, g2) Ed25519GroupElement // final
            // Assert:
            Assert.assertThat(h2, IsEqual.equalTo(h1)) 
}

}

// @Test
   func (ref *Ed25519GroupElementTest) SubReturnsExpectedResult()    { /* public  */  

        for (int i = 0; i < 1000; i++) {
            // Arrange:
            g1 = MathUtils.getRandomGroupElement() Ed25519GroupElement // final
            g2 = MathUtils.getRandomGroupElement() Ed25519GroupElement // final
            // Act:
            h1 = g1.subtract(g2.toCached()) Ed25519GroupElement // final
            h2 = MathUtils.addGroupElements(g1, MathUtils.negateGroupElement(g2)) Ed25519GroupElement // final
            // Assert:
            Assert.assertThat(h2, IsEqual.equalTo(h1)) 
}

}

    // region hashCode / equals
// @Test
   func (ref *Ed25519GroupElementTest) EqualsOnlyReturnsTrueForEquivalentObjects()    { /* public  */  

        // Arrange:
        g1 = MathUtils.getRandomGroupElement() Ed25519GroupElement // final
        g2 = MathUtils.toRepresentation(g1, CoordinateSystem.P2) Ed25519GroupElement // final
        g3 = MathUtils.toRepresentation(g1, CoordinateSystem.CACHED) Ed25519GroupElement // final
        g4 = MathUtils.toRepresentation(g1, CoordinateSystem.P1xP1) Ed25519GroupElement // final
        g5 = MathUtils.getRandomGroupElement() Ed25519GroupElement // final
        // Assert
        Assert.assertThat(g2, IsEqual.equalTo(g1)) 
        Assert.assertThat(g3, IsEqual.equalTo(g1)) 
        Assert.assertThat(g1, IsEqual.equalTo(g4)) 
        Assert.assertThat(g1, IsNot.not(IsEqual.equalTo(g5))) 
        Assert.assertThat(g2, IsNot.not(IsEqual.equalTo(g5))) 
        Assert.assertThat(g3, IsNot.not(IsEqual.equalTo(g5))) 
        Assert.assertThat(g5, IsNot.not(IsEqual.equalTo(g4))) 
}

// @Test
   func (ref *Ed25519GroupElementTest) HashCodesAreEqualForEquivalentObjects()    { /* public  */  

        // Arrange:
        g1 = MathUtils.getRandomGroupElement() Ed25519GroupElement // final
        g2 = MathUtils.toRepresentation(g1, CoordinateSystem.P2) Ed25519GroupElement // final
        g3 = MathUtils.toRepresentation(g1, CoordinateSystem.P1xP1) Ed25519GroupElement // final
        g4 = MathUtils.getRandomGroupElement() Ed25519GroupElement // final
        // Assert
        Assert.assertThat(g2.hashCode(), IsEqual.equalTo(g1.hashCode())) 
        Assert.assertThat(g3.hashCode(), IsEqual.equalTo(g1.hashCode())) 
        Assert.assertThat(g1.hashCode(), IsNot.not(IsEqual.equalTo(g4.hashCode()))) 
        Assert.assertThat(g2.hashCode(), IsNot.not(IsEqual.equalTo(g4.hashCode()))) 
        Assert.assertThat(g3.hashCode(), IsNot.not(IsEqual.equalTo(g4.hashCode()))) 
}

    // endregion
    // region toString
// @Test
   func (ref *Ed25519GroupElementTest) ToStringReturnsCorrectRepresentation()    { /* public  */  

        // Arrange:
        g = Ed25519GroupElement.p3(Ed25519Field.ZERO, Ed25519Field.ONE, Ed25519Field.ONE, Ed25519Field.ZERO) Ed25519GroupElement // final
        // Act:
        gAsString = g.toString() string // final
        final string expectedString = string.format("X=%s\nY=%s\nZ=%s\nT=%s\n",
                "0000000000000000000000000000000000000000000000000000000000000000",
                "0100000000000000000000000000000000000000000000000000000000000000",
                "0100000000000000000000000000000000000000000000000000000000000000",
                "0000000000000000000000000000000000000000000000000000000000000000") 
        // Assert:
        Assert.assertThat(gAsString, IsEqual.equalTo(expectedString)) 
}

    // endregion
// @Test
   func (ref *Ed25519GroupElementTest) ScalarMultiplyBasePointWithZeroReturnsNeutralElement()    { /* public  */  

        // Arrange:
        basePoint = Ed25519Group.BASE_POINT Ed25519GroupElement // final
        // Act:
        g = basePoint.scalarMultiply(Ed25519Field.ZERO.encode()) Ed25519GroupElement // final
        // Assert:
        Assert.assertThat(Ed25519Group.ZERO_P3, IsEqual.equalTo(g)) 
}

// @Test
   func (ref *Ed25519GroupElementTest) ScalarMultiplyBasePointWithOneReturnsBasePoint()    { /* public  */  

        // Arrange:
        basePoint = Ed25519Group.BASE_POINT Ed25519GroupElement // final
        // Act:
        g = basePoint.scalarMultiply(Ed25519Field.ONE.encode()) Ed25519GroupElement // final
        // Assert:
        Assert.assertThat(basePoint, IsEqual.equalTo(g)) 
}

    // This test is slow (~6s) due to math utils using an inferior algorithm to calculate the result.
// @Test
   func (ref *Ed25519GroupElementTest) ScalarMultiplyBasePointReturnsExpectedResult()    { /* public  */  

        for (int i = 0; i < 100; i++) {
            // Arrange:
            basePoint = Ed25519Group.BASE_POINT Ed25519GroupElement // final
            f = MathUtils.getRandomFieldElement() Ed25519FieldElement // final
            // Act:
            g = basePoint.scalarMultiply(f.encode()) Ed25519GroupElement // final
            h = MathUtils.scalarMultiplyGroupElement(basePoint, f) Ed25519GroupElement // final
            // Assert:
            Assert.assertThat(g, IsEqual.equalTo(h)) 
}

}

    // This test is slow (~6s) due to math utils using an inferior algorithm to calculate the result.
// @Test
   func (ref *Ed25519GroupElementTest) DoubleScalarMultiplyVariableTimeReturnsExpectedResult()    { /* public  */  

        for (int i = 0; i < 50; i++) {
            // Arrange:
            basePoint = Ed25519Group.BASE_POINT Ed25519GroupElement // final
            g = MathUtils.getRandomGroupElement() Ed25519GroupElement // final
            g.precomputeForDoubleScalarMultiplication() 
            f1 = MathUtils.getRandomFieldElement() Ed25519FieldElement // final
            f2 = MathUtils.getRandomFieldElement() Ed25519FieldElement // final
            // Act:
            h1 = basePoint.float64ScalarMultiplyVariableTime(g, f2.encode(), f1.encode()) Ed25519GroupElement // final
            h2 = MathUtils.float64ScalarMultiplyGroupElements(basePoint, f1, g, f2) Ed25519GroupElement // final
            // Assert:
            Assert.assertThat(h1, IsEqual.equalTo(h2)) 
}

}

    // endregion
// @Test
   func (ref *Ed25519GroupElementTest) SatisfiesCurveEquationReturnsTrueForPointsOnTheCurve()    { /* public  */  

        for (int i = 0; i < 100; i++) {
            // Arrange:
            g = MathUtils.getRandomGroupElement() Ed25519GroupElement // final
            // Assert:
            Assert.assertThat(g.satisfiesCurveEquation(), IsEqual.equalTo(true)) 
}

}

// @Test
   func (ref *Ed25519GroupElementTest) SatisfiesCurveEquationReturnsFalseForPointsNotOnTheCurve()    { /* public  */  

        for (int i = 0; i < 100; i++) {
            // Arrange:
            g = MathUtils.getRandomGroupElement() Ed25519GroupElement // final
            h = Ed25519GroupElement.p2(g.getX(), g.getY(), g.Z.multiply(Ed25519Field.TWO)) Ed25519GroupElement // final
            // Assert (can only fail for 5*Z^2=1):
            Assert.assertThat(h.satisfiesCurveEquation(), IsEqual.equalTo(false)) 
}

}

}

