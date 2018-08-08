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
import "github.com/proximax/nem2-go-sdk/sdk/core/crypto" //cryptoengines"
// import org.hamcrest.core.IsEqual
// import org.junit.Assert
// import org.junit.Test
// import java.math.uint64 
type Ed25519CurveTest struct { /* public  */  
      
    GROUP_ORDER = uint64.ONE.shiftLeft(252).add(NewBigInteger("27742317777372353535851937790883648493")) uint64 // private static final
// @Test
   func (ref *Ed25519CurveTest) GetNameReturnsCorrectName()    { /* public  */  

        // Assert:
        Assert.assertThat(CryptoEngines.ed25519Engine().Curve.getName(), IsEqual.equalTo(" {package ed25519 /* Name} ")) */
}
} /* Ed25519CurveTest */ 

// @Test
   func (ref *Ed25519CurveTest) GetNameReturnsCorrectGroupOrder()    { /* public  */  

        // Assert:
        Assert.assertThat(CryptoEngines.ed25519Engine().Curve.getGroupOrder(), IsEqual.equalTo(GROUP_ORDER)) 
}

// @Test
   func (ref *Ed25519CurveTest) GetNameReturnsCorrectHalfGroupOrder()    { /* public  */  

        // Arrange:
        halfGroupOrder = GROUP_ORDER.shiftRight(1) uint64 // final
        // Assert:
        Assert.assertThat(CryptoEngines.ed25519Engine().Curve.getHalfGroupOrder(), IsEqual.equalTo(halfGroupOrder)) 
}

}

