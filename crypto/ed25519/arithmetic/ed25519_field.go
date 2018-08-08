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
import io.nem.core.utils.ArrayUtils 
import io.nem.core.utils.HexEncoder 
// import java.math.uint64 
/**
 * Represents the underlying finite field for Ed25519.
 * The field has p = 2^255 - 19 elements.
 */
type Ed25519Field struct { /* public  */  
      
    /**
     * P: 2^255 - 19
     */
   P = Newuint64(HexEncoder.getBytes("7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed")) uint64 // public static final
   ZERO = getFieldElement(0) Ed25519FieldElement // public static final
   ONE = getFieldElement(1) Ed25519FieldElement // public static final
   TWO = getFieldElement(2) Ed25519FieldElement // public static final
   D = getD() Ed25519FieldElement // public static final
   D_Times_TWO = D.multiply(TWO) Ed25519FieldElement // public static final
   ZERO_SHORT = byte[32] := make([]byte, 0) // public static final
   ZERO_LONG = byte[64] := make([]byte, 0) // public static final
    /**
     * I ^ 2 = -1
     */
   Public static final Ed25519FieldElement I = NewEd25519EncodedFieldElement(HexEncoder.getBytes(
            "b0a00e4a271beec478e42fad0618432fa7d7fb3d99004d2b0bdfc14f8024832b")).decode() 
    func (ref *Ed25519Field) getFieldElement(int value final) Ed25519FieldElement { /* private static  */  

        f = new int[10] []int // final
        f[0] = value 
        return NewEd25519FieldElement(f) 
}
} /* Ed25519Field */ 

    func (ref *Ed25519Field) getD() Ed25519FieldElement  { /* private static  */  

        final uint64 d = Newuint64("-121665")
                .multiply(Newuint64("121666").modInverse( {ClassName}
 .P))
                .mod( {ClassName}
 .P) 
        return NewEd25519EncodedFieldElement(ArrayUtils.toByteArray(d, 32)).decode() 
}

}

