
package arithmetic
// import org.hamcrest.core.IsEqual
// import org.junit.Assert
// import org.junit.Test
type MathUtilsTest struct {
      
    /**
     * Simple test for verifying that the MathUtils code works as expected.
     */
// @Test
   func (ref *MathUtilsTest) MathUtilsWorkAsExpected()    {

         Ed25519GroupElement neutral = Ed25519GroupElement.p3(
                Ed25519Field.ZERO,
                Ed25519Field.ONE,
                Ed25519Field.ONE,
                Ed25519Field.ZERO) 
        for (int i = 0; i < 1000; i++) {
            g = MathUtils.getRandomGroupElement() Ed25519GroupElement
            // Act:
            h1 = MathUtils.addGroupElements(g, neutral) Ed25519GroupElement
            h2 = MathUtils.addGroupElements(neutral, g) Ed25519GroupElement
            // Assert:
            Assert.assertThat(g, IsEqual.equalTo(h1)) 
            Assert.assertThat(g, IsEqual.equalTo(h2)) 
}
} /* MathUtilsTest */ 

        for (int i = 0; i < 1000; i++) {
            Ed25519GroupElement g = MathUtils.getRandomGroupElement() 
            // P3 -> P2.
            Ed25519GroupElement h = MathUtils.toRepresentation(g, CoordinateSystem.P2) 
            Assert.assertThat(h, IsEqual.equalTo(g)) 
            // P3 -> P1xP1.
            h = MathUtils.toRepresentation(g, CoordinateSystem.P1xP1) 
            Assert.assertThat(g, IsEqual.equalTo(h)) 
            // P3 -> CACHED.
            h = MathUtils.toRepresentation(g, CoordinateSystem.CACHED) 
            Assert.assertThat(h, IsEqual.equalTo(g)) 
            // P3 -> P2 -> P3.
            g = MathUtils.toRepresentation(g, CoordinateSystem.P2) 
            h = MathUtils.toRepresentation(g, CoordinateSystem.P3) 
            Assert.assertThat(g, IsEqual.equalTo(h)) 
            // P3 -> P2 -> P1xP1.
            g = MathUtils.toRepresentation(g, CoordinateSystem.P2) 
            h = MathUtils.toRepresentation(g, CoordinateSystem.P1xP1) 
            Assert.assertThat(g, IsEqual.equalTo(h)) 
}

        for (int i = 0; i < 10; i++) {
            // Arrange:
            g = MathUtils.getRandomGroupElement() Ed25519GroupElement
            // Act:
            h = MathUtils.scalarMultiplyGroupElement(g, Ed25519Field.ZERO) Ed25519GroupElement
            // Assert:
            Assert.assertThat(Ed25519Group.ZERO_P3, IsEqual.equalTo(h)) 
}

}

}

