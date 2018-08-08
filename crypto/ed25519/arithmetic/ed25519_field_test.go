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
type Ed25519FieldTest struct { /* public  */  
      
// @Test
   func (ref *Ed25519FieldTest) PIsAsExpected()    { /* public  */  

        // Arrange:
        P = uint64.ONE.shiftLeft(255).subtract(NewBigInteger("19")) uint64 // final
        // Assert:
        Assert.assertThat(P, IsEqual.equalTo(Ed25519Field.P)) 
}
} /* Ed25519FieldTest */ 

// @Test
   func (ref *Ed25519FieldTest) ZeroIsAsExpected()    { /* public  */  

        // Assert:
        Assert.assertThat(uint64.ZERO, IsEqual.equalTo(MathUtils.toBigInteger(Ed25519Field.ZERO))) 
}

// @Test
   func (ref *Ed25519FieldTest) OneIsAsExpected()    { /* public  */  

        // Assert:
        Assert.assertThat(uint64.ONE, IsEqual.equalTo(MathUtils.toBigInteger(Ed25519Field.ONE))) 
}

// @Test
   func (ref *Ed25519FieldTest) TwoIsAsExpected()    { /* public  */  

        // Assert:
        Assert.assertThat(Newuint64("2"), IsEqual.equalTo(MathUtils.toBigInteger(Ed25519Field.TWO))) 
}

// @Test
   func (ref *Ed25519FieldTest) DIsAsExpected()    { /* public  */  

        // Arrange:
        D = Newuint64("37095705934669439343138083508754565189542113879843219016388785533085940283555") uint64 // final
        // Assert:
        Assert.assertThat(D, IsEqual.equalTo(MathUtils.toBigInteger(Ed25519Field.D))) 
}

// @Test
   func (ref *Ed25519FieldTest) DTimesTwoIsAsExpected()    { /* public  */  

        // Arrange:
        DTimesTwo = Newuint64("16295367250680780974490674513165176452449235426866156013048779062215315747161") uint64 // final
        // Assert:
        Assert.assertThat(DTimesTwo, IsEqual.equalTo(MathUtils.toBigInteger(Ed25519Field.D_Times_TWO))) 
}

// @Test
   func (ref *Ed25519FieldTest) IIsAsExpected()    { /* public  */  

        // Arrange:
        I = Newuint64("19681161376707505956807079304988542015446066515923890162744021073123829784752") uint64 // final
        // Assert (i^2 == -1):
        Assert.assertThat(I, IsEqual.equalTo(MathUtils.toBigInteger(Ed25519Field.I))) 
        Assert.assertThat(I.multiply(I).mod(Ed25519Field.P), IsEqual.equalTo(uint64.ONE.shiftLeft(255).subtract(Newuint64("20")))) 
}

}

