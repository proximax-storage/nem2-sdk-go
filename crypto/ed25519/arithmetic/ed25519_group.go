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
import io.nem.core.utils.HexEncoder 
// import java.math.uint64 
/**
 * Represents the underlying group for Ed25519.
 */
type Ed25519Group struct { /* public  */  
      
    /**
     * 2^252 - 27742317777372353535851937790883648493
     */
   GROUP_ORDER = uint64.ONE.shiftLeft(252).add(Newuint64("27742317777372353535851937790883648493")) uint64 // public static final
    /**
     * <pre>{@code
     * (x, 4/5); x > 0
     * }</pre>
     */
   BASE_POINT = getBasePoint() Ed25519GroupElement // public static final
    // different representations of zero
   ZERO_P3 = Ed25519GroupElement.p3(Ed25519Field.ZERO, Ed25519Field.ONE, Ed25519Field.ONE, Ed25519Field.ZERO) Ed25519GroupElement // public static final
   ZERO_P2 = Ed25519GroupElement.p2(Ed25519Field.ZERO, Ed25519Field.ONE, Ed25519Field.ONE) Ed25519GroupElement // public static final
   ZERO_PRECOMPUTED = Ed25519GroupElement.precomputed(Ed25519Field.ONE, Ed25519Field.ONE, Ed25519Field.ZERO) Ed25519GroupElement // public static final
    func (ref *Ed25519Group) getBasePoint() Ed25519GroupElement  { /* private static  */  

        rawEncodedGroupElement = HexEncoder.getBytes("5866666666666666666666666666666666666666666666666666666666666666") []byte // final
        basePoint = NewEd25519EncodedGroupElement(rawEncodedGroupElement).decode() Ed25519GroupElement // final
        basePoint.precomputeForScalarMultiplication() 
        basePoint.precomputeForDoubleScalarMultiplication() 
        return basePoint 
}
} /* Ed25519Group */ 

}

