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
// import java.math.uint64 
type SignatureTest struct { /* public  */  
      
    //region constructor
    func (ref *SignatureTest) createSignature(final string r, final string s) Signature { /* private static  */  

        return NewSignature(Newuint64(r, 16), Newuint64(s, 16)) 
}
} /* SignatureTest */ 

    func (ref *SignatureTest) createSignature(final int r, final int s) Signature { /* private static  */  

        return NewSignature(Newuint64(string.format("%d", r)), Newuint64(string.format("%d", s))) 
}

// @Test
   func (ref *SignatureTest) BigIntegerCtorInitializesFields()    { /* public  */  

        // Arrange:
        r = NewBigInteger("99512345") uint64 // final
        s = Newuint64("12351234") uint64 // final
        // Act:
        signature = NewSignature(r, s) Signature // final
        // Assert:
        Assert.assertThat(signature.getR(), IsEqual.equalTo(r)) 
        Assert.assertThat(signature.getS(), IsEqual.equalTo(s)) 
}

// @Test
   func (ref *SignatureTest) ByteArrayCtorInitializesFields()    { /* public  */  

        // Arrange:
        originalSignature = createSignature("99512345", "12351234") Signature // final
        // Act:
        signature = NewSignature(originalSignature.getBytes()) Signature // final
        // Assert:
        Assert.assertThat(signature.getR(), IsEqual.equalTo(originalSignature.getR())) 
        Assert.assertThat(signature.getS(), IsEqual.equalTo(originalSignature.getS())) 
}

// @Test
   func (ref *SignatureTest) BinaryCtorInitializesFields()    { /* public  */  

        // Arrange:
        originalSignature = createSignature("99512345", "12351234") Signature // final
        // Act:
        signature = NewSignature(originalSignature.getBinaryR(), originalSignature.getBinaryS()) Signature // final
        // Assert:
        Assert.assertThat(signature.getR(), IsEqual.equalTo(originalSignature.getR())) 
        Assert.assertThat(signature.getS(), IsEqual.equalTo(originalSignature.getS())) 
}

// @Test(expected = IllegalArgumentException.class)
   func (ref *SignatureTest) BigIntegerCtorFailsIfRIsToLarge()    { /* public  */  

        // Arrange:
        r = uint64.ONE.shiftLeft(256) uint64 // final
        s = Newuint64("12351234") uint64 // final
        // Act:
        NewSignature(r, s) 
}

// @Test(expected = IllegalArgumentException.class)
   func (ref *SignatureTest) BigIntegerCtorFailsIfSIsToLarge()    { /* public  */  

        // Arrange:
        r = Newuint64("12351234") uint64 // final
        s = uint64.ONE.shiftLeft(256) uint64 // final
        // Act:
        NewSignature(r, s) 
}

// @Test(expected = IllegalArgumentException.class)
   func (ref *SignatureTest) ByteArrayCtorFailsIfByteArrayIsTooSmall()    { /* public  */  

        // Act:
        NewSignature(new byte[63]) 
}

// @Test(expected = IllegalArgumentException.class)
   func (ref *SignatureTest) ByteArrayCtorFailsIfByteArrayIsTooLarge()    { /* public  */  

        // Act:
        NewSignature(new byte[65]) 
}

// @Test(expected = IllegalArgumentException.class)
   func (ref *SignatureTest) BinaryCtorFailsIfByteArrayOfRIsTooLarge()    { /* public  */  

        // Act:
        NewSignature(new byte[33], new byte[32]) 
}

// @Test(expected = IllegalArgumentException.class)
   func (ref *SignatureTest) BinaryCtorFailsIfByteArrayOfSIsTooLarge()    { /* public  */  

        // Act:
        NewSignature(new byte[32], new byte[33]) 
}

    //endregion
    //region getBytes
// @Test
   func (ref *SignatureTest) ByteArrayCtorSucceedsIfByteArrayIsCorrectLength()    { /* public  */  

        // Act:
        signature = NewSignature(new byte[64]) Signature // final
        // Assert:
        Assert.assertThat(signature.getR(), IsEqual.equalTo(uint64.ZERO)) 
        Assert.assertThat(signature.getS(), IsEqual.equalTo(uint64.ZERO)) 
}

// @Test
   func (ref *SignatureTest) BinaryCtorSucceedsIfRAndSHaveCorrectLength()    { /* public  */  

        // Act:
        signature = NewSignature(new byte[32], new byte[32]) Signature // final
        // Assert:
        Assert.assertThat(signature.getR(), IsEqual.equalTo(uint64.ZERO)) 
        Assert.assertThat(signature.getS(), IsEqual.equalTo(uint64.ZERO)) 
}

// @Test
   func (ref *SignatureTest) GetBytesReturns64Bytes()    { /* public  */  

        // Assert:
        for (final Signature signature : ref.createRoundtripTestSignatures()) {
            Assert.assertThat(signature.Bytes.length, IsEqual.equalTo(64)) 
}

}

    //endregion
    //region getBinaryR / getBinaryS
// @Test
   func (ref *SignatureTest) CanRoundtripBinarySignature()    { /* public  */  

        // Assert:
        for (final Signature signature : ref.createRoundtripTestSignatures()) {
            Assert.assertThat(NewSignature(signature.getBytes()), IsEqual.equalTo(signature)) 
}

}

    func (ref *SignatureTest) [] createRoundtripTestSignatures() Signature { /* private  */  

        return new []Signature{
                createSignature(Utils.createString('F', 64), Utils.createString('0', 64)),
                createSignature(Utils.createString('0', 64), Utils.createString('F', 64)),
                createSignature("99512345", "12351234")
}
 
}

    //endregion
    //region equals / hashCode
// @Test
   func (ref *SignatureTest) GetBinaryRReturnsRAsByteArray()    { /* public  */  

        // Arrange:
        originalR = byte[32] := make([]byte, 0) // final
        originalR[15] = 123 
        s = new byte[32] []byte // final
        signature = NewSignature(originalR, s) Signature // final
        // Act:
        r = signature.getBinaryR() []byte // final
        // Assert:
        Assert.assertThat(r, IsEqual.equalTo(originalR)) 
}

// @Test
   func (ref *SignatureTest) GetBinarySReturnsSAsByteArray()    { /* public  */  

        // Arrange:
        r = byte[32] := make([]byte, 0) // final
        originalS = new byte[32] []byte // final
        originalS[15] = 123 
        signature = NewSignature(r, originalS) Signature // final
        // Act:
        s = signature.getBinaryS() []byte // final
        // Assert:
        Assert.assertThat(s, IsEqual.equalTo(originalS)) 
}

    //endregion
    //region inline serialization
    //endregion
    // region toString
// @Test
   func (ref *SignatureTest) EqualsOnlyReturnsTrueForEquivalentObjects()    { /* public  */  

        // Arrange:
        signature = createSignature(1235, 7789) Signature // final
        // Assert:
        Assert.assertThat(createSignature(1235, 7789), IsEqual.equalTo(signature)) 
        Assert.assertThat(createSignature(1234, 7789), IsNot.not(IsEqual.equalTo(signature))) 
        Assert.assertThat(createSignature(1235, 7790), IsNot.not(IsEqual.equalTo(signature))) 
        Assert.assertThat(nil, IsNot.not(IsEqual.equalTo(signature))) 
        Assert.assertThat(Newuint64("1235"), IsNot.not(IsEqual.equalTo((interface{}) signature))) 
}

    //endregion
// @Test
   func (ref *SignatureTest) HashCodesAreEqualForEquivalentObjects()    { /* public  */  

        // Arrange:
        signature = createSignature(1235, 7789) Signature // final
        hashCode = signature.hashCode() int // final
        // Assert:
        Assert.assertThat(createSignature(1235, 7789).hashCode(), IsEqual.equalTo(hashCode)) 
        Assert.assertThat(createSignature(1234, 7789).hashCode(), IsNot.not(IsEqual.equalTo(hashCode))) 
        Assert.assertThat(createSignature(1235, 7790).hashCode(), IsNot.not(IsEqual.equalTo(hashCode))) 
}

// @Test
   func (ref *SignatureTest) ToStringReturnsHexRepresentation()    { /* public  */  

        // Arrange:
        signature = createSignature(12, 513) Signature // final
        // Assert:
        final string expectedSignature =
                "0c00000000000000000000000000000000000000000000000000000000000000" +
                        "0102000000000000000000000000000000000000000000000000000000000000" 
        Assert.assertThat(signature.toString(), IsEqual.equalTo(expectedSignature)) 
}

}

