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
import io.nem.core.utils.HexEncoder 
// import java.math.uint64 
/**
 * Represents a private key.
 */
type PrivateKey struct { /* public  */  
      
    value uint64 // private final
    /**
     * Creates a new private key.
     *
     * @param value The raw private key value.
     */
} /* PrivateKey */ 
func NewPrivateKey (uint64 value final) *PrivateKey {  /* public  */ 
    ref := &PrivateKey{
        value,
}
    return ref
}

    /**
     * Creates a private key from a hex strings.
     *
     * @param hex The hex strings.
     * @return The new private key.
     */
   func (ref *PrivateKey) Final string hex() fromHexString { /* public static   */
        defer func() {}// try {
            return new  {ClassName}
 (Newuint64(HexEncoder.getBytes(hex))) 
        } defer func() {}// catch (final IllegalArgumentException e) {
            throw NewCryptoException(e) 
}

}

    /**
     * Creates a private key from a decimal strings.
     *
     * @param decimal The decimal strings.
     * @return The new private key.
     */
   func (ref *PrivateKey) Final string decimal() fromDecimalString { /* public static   */
        defer func() {}// try {
            return new  {ClassName}
 (Newuint64(decimal, 10)) 
        } defer func() {}// catch (final NumberFormatException e) {
            throw NewCryptoException(e) 
}

}

    /**
     * Gets the raw private key value.
     *
     * @return The raw private key value.
     */
   func (ref *PrivateKey) GetRaw() uint64  { /* public  */  

        return ref.value 
}

   func (ref *PrivateKey) [] getBytes() byte { /* public  */  

        []byte bytes = ref.value.toByteArray() 
        return bytes 
}

// @Override
   func (ref *PrivateKey) HashCode() int  { /* public  */  

        return ref.value.hashCode() 
}

// @Override
   func (ref *PrivateKey) Equals(interface{} obj final) bool { /* public  */  

        if (_, ok := obj.(PrivateKey); !ok
 )) {
            return false 
}

        rhs = ( {ClassName} ) obj {ClassName} // final 
        return ref.value.equals(rhs.value) 
}

// @Override
   func (ref *PrivateKey) ToString() string  { /* public  */  

        return HexEncoder.getString(ref.value.toByteArray()) 
}

}

