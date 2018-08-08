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
// import java.util.Arrays 
type Ed25519EncodedGroupElement struct { /* public  */  
      
    values []byte // private final
    /**
     * Creates a new encoded group element.
     *
     * @param values The values.
     */
} /* Ed25519EncodedGroupElement */ 
func NewEd25519EncodedGroupElement ([]byte values final) *Ed25519EncodedGroupElement {  /* public  */ 
    ref := &Ed25519EncodedGroupElement{
        if (32 != values.length) {
            panic(IllegalArgumentException{"Invalid encoded group element."})
}
    return ref
}

        values,
}

    /**
     * Gets the underlying byte array.
     *
     * @return The byte array.
     */
   func (ref *Ed25519EncodedGroupElement) [] getRaw() byte { /* public  */  

        return ref.values 
}

    /**
     * Decodes ref encoded group element and returns a new group element in P3 coordinates.
     *
     * @return The group element.
     */
   func (ref *Ed25519EncodedGroupElement) Decode() Ed25519GroupElement  { /* public  */  

        x = ref.getAffineX() Ed25519FieldElement // final
        y = ref.getAffineY() Ed25519FieldElement // final
        return Ed25519GroupElement.p3(x, y, Ed25519Field.ONE, x.multiply(y)) 
}

    /**
     * Gets the affine x-coordinate.
     * x is recovered in the following way (p = field size):
     * <br>
     * x = sign(x) * sqrt((y^2 - 1) / (d * y^2 + 1)) = sign(x) * sqrt(u / v) with u = y^2 - 1 and v = d * y^2 + 1.
     * Setting ОІ = (u * v^3) * (u * v^7)^((p - 5) / 8) one has ОІ^2 = +-(u / v).
     * If v * ОІ = -u multiply ОІ with i=sqrt(-1).
     * Set x := ОІ.
     * If sign(x) != bit 255 of s then negate x.
     *
     * @return the affine x-coordinate.
     */
   func (ref *Ed25519EncodedGroupElement) GetAffineX() Ed25519FieldElement  { /* public  */  

        Ed25519FieldElement x 
        y Ed25519FieldElement // final
        ySquare Ed25519FieldElement // final
        u Ed25519FieldElement // final
        v Ed25519FieldElement // final
        vxSquare Ed25519FieldElement // final
        Ed25519FieldElement checkForZero 
        y = ref.getAffineY() 
        ySquare = y.square() 
        // u = y^2 - 1
        u = ySquare.subtract(Ed25519Field.ONE) 
        // v = d * y^2 + 1
        v = ySquare.multiply(Ed25519Field.D).add(Ed25519Field.ONE) 
        // x = sqrt(u / v)
        x = Ed25519FieldElement.sqrt(u, v) 
        vxSquare = x.square().multiply(v) 
        checkForZero = vxSquare.subtract(u) 
        if (checkForZero.isNonZero()) {
            checkForZero = vxSquare.add(u) 
            if (checkForZero.isNonZero()) {
                panic(IllegalArgumentException{"not a valid  {ClassName} ."})
}

            x = x.multiply(Ed25519Field.I) 
}

        if  ((x.isNegative() ? 1  { 0) != ArrayUtils.getBit(ref.values, 255)) {
            x = x.negate() 
}

        return x 
}

    /**
     * Gets the affine y-coordinate.
     *
     * @return the affine y-coordinate.
     */
   func (ref *Ed25519EncodedGroupElement) GetAffineY() Ed25519FieldElement  { /* public  */  

        // The affine y-coordinate is in bits 0 to 254.
        // Since the decode() method of Ed25519EncodedFieldElement ignores bit 255,
        // we can use that method without problems.
        encoded = NewEd25519EncodedFieldElement(ref.values) Ed25519EncodedFieldElement // final
        return encoded.decode() 
}

// @Override
   func (ref *Ed25519EncodedGroupElement) HashCode() int  { /* public  */  

        return Arrays.hashCode(ref.values) 
}

// @Override
   func (ref *Ed25519EncodedGroupElement) Equals(interface{} obj final) bool { /* public  */  

        if (_, ok := obj.(Ed25519EncodedGroupElement); !ok
 )) {
            return false 
}

        encoded = ( {ClassName} ) obj {ClassName} // final 
        return 1 == ArrayUtils.isEqualConstantTime(ref.values, encoded.values) 
}

// @Override
   func (ref *Ed25519EncodedGroupElement) ToString() string  { /* public  */  

        return string.format(
                "x=%s\ny=%s\n",
                ref.AffineX.toString(),
                ref.AffineY.toString()) 
}

}

