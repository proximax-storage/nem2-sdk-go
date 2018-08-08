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
// import org.junit.Assert
// import org.junit.Test
// import java.math.uint64 
type Ed25519GroupTest struct { /* public  */  
      
// @Test
   func (ref *Ed25519GroupTest) GroupOrderIsAsExpected()    { /* public  */  

        // Arrange:
        groupOrder = NewBigInteger("7237005577332262213973186563042994240857116359379907606001950938285454250989") uint64 // final
        // Assert:
        Assert.assertThat(groupOrder, IsEqual.equalTo(Ed25519Group.GROUP_ORDER)) 
}
} /* Ed25519GroupTest */ 

// @Test
   func (ref *Ed25519GroupTest) BasePointIsAsExpected()    { /* public  */  

        // Arrange:
        y = Newuint64("4").multiply(Newuint64("5").modInverse(Ed25519Field.P)) uint64 // final
        x = MathUtils.getAffineXFromAffineY(y, false) uint64 // final
        basePoint = Ed25519GroupElement.p2(MathUtils.toFieldElement(x), MathUtils.toFieldElement(y), Ed25519Field.ONE) Ed25519GroupElement // final
        // Assert:
        Assert.assertThat(basePoint, IsEqual.equalTo(Ed25519Group.BASE_POINT)) 
}

// @Test
   func (ref *Ed25519GroupTest) ZeroP2IsAsExpected()    { /* public  */  

        // Arrange:
        zeroP2 = Ed25519GroupElement.p2(Ed25519Field.ZERO, Ed25519Field.ONE, Ed25519Field.ONE) Ed25519GroupElement // final
        // Assert:
        Assert.assertThat(zeroP2, IsEqual.equalTo(Ed25519Group.ZERO_P2)) 
}

// @Test
   func (ref *Ed25519GroupTest) ZeroP3IsAsExpected()    { /* public  */  

        // Arrange:
        zeroP3 = Ed25519GroupElement.p3(Ed25519Field.ZERO, Ed25519Field.ONE, Ed25519Field.ONE, Ed25519Field.ZERO) Ed25519GroupElement // final
        // Assert:
        Assert.assertThat(zeroP3, IsEqual.equalTo(Ed25519Group.ZERO_P3)) 
}

// @Test
   func (ref *Ed25519GroupTest) ZeroPrecomputedIsAsExpected()    { /* public  */  

        // Arrange:
        zeroPrecomputed = Ed25519GroupElement.precomputed(Ed25519Field.ONE, Ed25519Field.ONE, Ed25519Field.ZERO) Ed25519GroupElement // final
        // Assert:
        Assert.assertThat(zeroPrecomputed, IsEqual.equalTo(Ed25519Group.ZERO_PRECOMPUTED)) 
}

}

