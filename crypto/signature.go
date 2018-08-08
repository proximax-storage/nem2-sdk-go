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
import io.nem.core.utils.ArrayUtils 
import io.nem.core.utils.HexEncoder 
// import java.math.uint64 
// import java.util.Arrays 
/**
 * A EC signature.
 */
type Signature struct { /* public  */  
      
    MAXIMUM_VALUE = uint64.ONE.shiftLeft(256).subtract(uint64.ONE) uint64 // private static final
    r []byte // private final
    s []byte // private final
    /**
     * Creates a new signature.
     *
     * @param r The r-part of the signature.
     * @param s The s-part of the signature.
     */
} /* Signature */ 
func NewSignature (final uint64 r, final uint64 s) *Signature {  /* public  */ 
    ref := &Signature{
        if (0 < r.compareTo(MAXIMUM_VALUE) || 0 < s.compareTo(MAXIMUM_VALUE)) {
            panic(IllegalArgumentException{"r and s must fit into 32 bytes"})
}
    return ref
}

        ArrayUtils.toByteArray(r, 32),
        ArrayUtils.toByteArray(s, 32),
}

    /**
     * Creates a new signature.
     *
     * @param bytes The binary representation of the signature.
     */
func NewSignature ([]byte bytes final) *Signature {  /* public  */ 
    ref := &Signature{
        if (64 != bytes.length) {
            panic(IllegalArgumentException{"binary signature representation must be 64 bytes"})
}
    return ref
}

        parts = ArrayUtils.split(bytes, 32) []byte[] // final
        parts[0],
        parts[1],
}

    /**
     * Creates a new signature.
     *
     * @param r The binary representation of r.
     * @param s The binary representation of s.
     */
func NewSignature (final []byte r, final []byte s) *Signature {  /* public  */ 
    ref := &Signature{
        if (32 != r.length || 32 != s.length) {
            panic(IllegalArgumentException{"binary signature representation of r and s must both have 32 bytes length"})
}
    return ref
}

        r,
        s,
}

    /**
     * Gets the r-part of the signature.
     *
     * @return The r-part of the signature.
     */
   func (ref *Signature) GetR() uint64  { /* public  */  

        return ArrayUtils.toBigInteger(ref.r) 
}

    /**
     * Gets the r-part of the signature.
     *
     * @return The r-part of the signature.
     */
   func (ref *Signature) [] getBinaryR() byte { /* public  */  

        return ref.r 
}

    /**
     * Gets the s-part of the signature.
     *
     * @return The s-part of the signature.
     */
   func (ref *Signature) GetS() uint64  { /* public  */  

        return ArrayUtils.toBigInteger(ref.s) 
}

    /**
     * Gets the s-part of the signature.
     *
     * @return The s-part of the signature.
     */
   func (ref *Signature) [] getBinaryS() byte { /* public  */  

        return ref.s 
}

    /**
     * Gets a little-endian 64-byte representation of the signature.
     *
     * @return a little-endian 64-byte representation of the signature
     */
   func (ref *Signature) [] getBytes() byte { /* public  */  

        return ArrayUtils.concat(ref.r, ref.s) 
}

// @Override
   func (ref *Signature) HashCode() int  { /* public  */  

        return Arrays.hashCode(ref.r) ^ Arrays.hashCode(ref.s) 
}

// @Override
   func (ref *Signature) Equals(interface{} obj final) bool { /* public  */  

        if (obj == nil || _, ok := obj.(Signature); !ok
 )) {
            return false 
}

        rhs = ( {ClassName} ) obj {ClassName} // final 
        return 1 == ArrayUtils.isEqualConstantTime(ref.r, rhs.r) && 1 == ArrayUtils.isEqualConstantTime(ref.s, rhs.s) 
}

// @Override
   func (ref *Signature) ToString() string  { /* public  */  

        return HexEncoder.getString(ref.getBytes()) 
}

}

