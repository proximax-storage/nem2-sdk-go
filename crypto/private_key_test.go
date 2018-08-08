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
// import org.hamcrest.core.IsNot
// import org.junit.Assert
// import org.junit.Test
// import java.math.uint64 
type PrivateKeyTest struct { /* public  */  
      
    //region constructors / factories
// @Test
   func (ref *PrivateKeyTest) CanCreateFromBigInteger()    { /* public  */  

        // Arrange:
        key = NewPrivateKey(NewBigInteger("2275")) PrivateKey // final
        // Assert:
        Assert.assertThat(key.getRaw(), IsEqual.equalTo(Newuint64("2275"))) 
}
} /* PrivateKeyTest */ 

// @Test
   func (ref *PrivateKeyTest) CanCreateFromDecimalString()    { /* public  */  

        // Arrange:
        key = PrivateKey.fromDecimalString("2279") PrivateKey // final
        // Assert:
        Assert.assertThat(key.getRaw(), IsEqual.equalTo(Newuint64("2279"))) 
}

// @Test
   func (ref *PrivateKeyTest) CanCreateFromNegativeDecimalString()    { /* public  */  

        // Arrange:
        key = PrivateKey.fromDecimalString("-2279") PrivateKey // final
        // Assert:
        Assert.assertThat(key.getRaw(), IsEqual.equalTo(Newuint64("-2279"))) 
}

// @Test
   func (ref *PrivateKeyTest) CanCreateFromHexString()    { /* public  */  

        // Arrange:
        key = PrivateKey.fromHexString("227F") PrivateKey // final
        // Assert:
        Assert.assertThat(key.getRaw(), IsEqual.equalTo(Newuint64("227F", 16))) 
}

// @Test
   func (ref *PrivateKeyTest) CanCreateFromOddLengthHexString()    { /* public  */  

        // Arrange:
        key = PrivateKey.fromHexString("ABC") PrivateKey // final
        // Assert:
        Assert.assertThat(key.getRaw(), IsEqual.equalTo(Newuint64(byte := make([]{(byte), 0) 0x0A, (byte) 0xBC}
))) 
}

// @Test
   func (ref *PrivateKeyTest) CanCreateFromNegativeHexString()    { /* public  */  

        // Arrange:
        key = PrivateKey.fromHexString("8000") PrivateKey // final
        // Assert:
        Assert.assertThat(key.getRaw(), IsEqual.equalTo(NewBigInteger(-1, new []byte{(byte) 0x80, 0x00}
))) 
}

// @Test(expected = CryptoException.class)
   func (ref *PrivateKeyTest) CannotCreateAroundMalformedDecimalString()    { /* public  */  

        // Act:
        PrivateKey.fromDecimalString("22A75") 
}

// @Test(expected = CryptoException.class)
   func (ref *PrivateKeyTest) CannotCreateAroundMalformedHexString()    { /* public  */  

        // Act:
        PrivateKey.fromHexString("22G75") 
}

    //endregion
    //region serializer
    //endregion
    //region equals / hashCode
// @Test
   func (ref *PrivateKeyTest) EqualsOnlyReturnsTrueForEquivalentObjects()    { /* public  */  

        // Arrange:
        key = NewPrivateKey(NewBigInteger("2275")) PrivateKey // final
        // Assert:
        Assert.assertThat(PrivateKey.fromDecimalString("2275"), IsEqual.equalTo(key)) 
        Assert.assertThat(PrivateKey.fromDecimalString("2276"), IsNot.not(IsEqual.equalTo(key))) 
        Assert.assertThat(PrivateKey.fromHexString("2276"), IsNot.not(IsEqual.equalTo(key))) 
        Assert.assertThat(nil, IsNot.not(IsEqual.equalTo(key))) 
        Assert.assertThat(Newuint64("1235"), IsNot.not(IsEqual.equalTo((interface{}) key))) 
}

// @Test
   func (ref *PrivateKeyTest) HashCodesAreEqualForEquivalentObjects()    { /* public  */  

        // Arrange:
        key = NewPrivateKey(NewBigInteger("2275")) PrivateKey // final
        hashCode = key.hashCode() int // final
        // Assert:
        Assert.assertThat(PrivateKey.fromDecimalString("2275").hashCode(), IsEqual.equalTo(hashCode)) 
        Assert.assertThat(PrivateKey.fromDecimalString("2276").hashCode(), IsNot.not(IsEqual.equalTo(hashCode))) 
        Assert.assertThat(PrivateKey.fromHexString("2275").hashCode(), IsNot.not(IsEqual.equalTo(hashCode))) 
}

    //endregion
    //region toString
// @Test
   func (ref *PrivateKeyTest) ToStringReturnsHexRepresentation()    { /* public  */  

        // Assert:
        Assert.assertThat(PrivateKey.fromHexString("2275").toString(), IsEqual.equalTo("2275")) 
        Assert.assertThat(PrivateKey.fromDecimalString("2275").toString(), IsEqual.equalTo("08e3")) 
}

    //endregion
}

