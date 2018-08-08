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
package ed25519 /*  {packageName}  */
import "github.com/proximax/nem2-go-sdk/sdk/core/crypto" //privatekey"
import io.nem.core.test.Utils 
// import org.hamcrest.core.IsEqual
// import org.junit.Assert
// import org.junit.Test
// import java.math.uint64 
type Ed25519UtilsTest struct { /* public  */  
      
    //region prepareForScalarMultiply
// @Test
   func (ref *Ed25519UtilsTest) PrepareForScalarMultiplyReturnsClampedValue()    { /* public  */  

        // Arrange:
        privateKey = NewPrivateKey(Newuint64(Utils.generateRandomBytes(32))) PrivateKey // final
        // Act:
        a = Ed25519Utils.prepareForScalarMultiply(privateKey).getRaw() []byte // final
        // Assert:
        Assert.assertThat(a[31] & 0x40, IsEqual.equalTo(0x40)) 
        Assert.assertThat(a[31] & 0x80, IsEqual.equalTo(0x0)) 
        Assert.assertThat(a[0] & 0x7, IsEqual.equalTo(0x0)) 
}
} /* Ed25519UtilsTest */ 

    //endregion
}

