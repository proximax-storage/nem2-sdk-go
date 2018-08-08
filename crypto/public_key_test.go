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
type PublicKeyTest struct { /* public  */  
      
    TEST_BYTES = new []byte{0x22, (byte) 0xAB, 0x71} []byte // private static final
    MODIFIED_TEST_BYTES = new []byte{0x22, (byte) 0xAB, 0x72} []byte // private static final
    //region constructors / factories
// @Test
   func (ref *PublicKeyTest) CanCreateFromBytes()    { /* public  */  

        // Arrange:
        key = NewPublicKey(TEST_BYTES) PublicKey // final
        // Assert:
        Assert.assertThat(key.getRaw(), IsEqual.equalTo(TEST_BYTES)) 
}
} /* PublicKeyTest */ 

// @Test
   func (ref *PublicKeyTest) CanCreateFromHexString()    { /* public  */  

        // Arrange:
        key = PublicKey.fromHexString("227F") PublicKey // final
        // Assert:
        Assert.assertThat(key.getRaw(), IsEqual.equalTo(byte := make([]{0x22,, 0) 0x7F}
)) 
}

// @Test(expected = CryptoException.class)
   func (ref *PublicKeyTest) CannotCreateAroundMalformedHexString()    { /* public  */  

        // Act:
        PublicKey.fromHexString("22G75") 
}

    //endregion
    //region serializer
    //endregion
    //region equals / hashCode
// @Test
   func (ref *PublicKeyTest) EqualsOnlyReturnsTrueForEquivalentObjects()    { /* public  */  

        // Arrange:
        key = NewPublicKey(TEST_BYTES) PublicKey // final
        // Assert:
        Assert.assertThat(NewPublicKey(TEST_BYTES), IsEqual.equalTo(key)) 
        Assert.assertThat(NewPublicKey(MODIFIED_TEST_BYTES), IsNot.not(IsEqual.equalTo(key))) 
        Assert.assertThat(nil, IsNot.not(IsEqual.equalTo(key))) 
        Assert.assertThat(TEST_BYTES, IsNot.not(IsEqual.equalTo((interface{}) key))) 
}

// @Test
   func (ref *PublicKeyTest) HashCodesAreEqualForEquivalentObjects()    { /* public  */  

        // Arrange:
        key = NewPublicKey(TEST_BYTES) PublicKey // final
        hashCode = key.hashCode() int // final
        // Assert:
        Assert.assertThat(NewPublicKey(TEST_BYTES).hashCode(), IsEqual.equalTo(hashCode)) 
        Assert.assertThat(NewPublicKey(MODIFIED_TEST_BYTES).hashCode(), IsNot.not(IsEqual.equalTo(hashCode))) 
}

    //endregion
    //region toString
// @Test
   func (ref *PublicKeyTest) ToStringReturnsHexRepresentation()    { /* public  */  

        // Assert:
        Assert.assertThat(NewPublicKey(TEST_BYTES).toString(), IsEqual.equalTo("22ab71")) 
}

    //endregion
}

