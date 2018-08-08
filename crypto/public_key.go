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
// import java.util.Arrays
/**
* Represents a public key.
 */
type PublicKey struct { /* public  */  
      
    Raw []byte // private final
    /**
    * Creates a new public key.
     *
    * @param bytes The raw public key value.
     */
} /* PublicKey */ 
func NewPublicKey (bytes []byte) *PublicKey {  /* public  */
    ref := &PublicKey{ bytes}
    return ref
}

    /**
    * Creates a public key from a hex strings.
     *
     * @param hex The hex strings.
    * @return The new public key.
     */
   func (ref *PublicKey) hex() string { /* public static   */
        return string(ref.value)
}

    /**
    * Gets the raw public key value.
     *
    * @return The raw public key value.
     */
   func (ref *PublicKey) [] getRaw() byte { /* public  */  

        return ref.value
}

// @Override
   func (ref *PublicKey) HashCode() int  { /* public  */  

        return Arrays.hashCode(ref.value)
}

// @Override
   func (ref *PublicKey) Equals(interface{} obj final) bool { /* public  */  

        if (obj == nil || _, ok := obj.(PublicKey); !ok
 )) {
            return false 
}

        rhs = ( {ClassName} ) obj {ClassName} // final 
        return Arrays.equals(ref.value, rhs.value)
}

// @Override
   func (ref *PublicKey) ToString() string  { /* public  */  

        return HexEncoder.getString(ref.value)
}

}

