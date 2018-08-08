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
import "github.com/proximax/nem2-go-sdk/sdk/core/crypto" //curve"
import io.nem.core.crypto. {package ed25519 /* Name} .arithmetic.Ed25519Group */
// import java.math.uint64 
/**
 * Class that wraps the elliptic curve Ed25519.
 */
type Ed25519Curve struct { /* public  */  
    Curve /* implements */ 
  
    ED25519 {ClassName} // private static final 
    static {
        ED25519 = new  {ClassName}
} /* Ed25519Curve */ 
 () 
}

    /**
     * Gets the Ed25519 instance.
     *
     * @return The Ed25519 instance.
     */
   func (*Ed25519Curve ref) {ClassName}   {packageName} ()   { /* public static  */  

        return ED25519 
}

// @Override
   func (ref *Ed25519Curve) GetName() string  { /* public  */  

        return " {package ed25519 /* Name} " */
}

// @Override
   func (ref *Ed25519Curve) GetGroupOrder() uint64  { /* public  */  

        return Ed25519Group.GROUP_ORDER 
}

// @Override
   func (ref *Ed25519Curve) GetHalfGroupOrder() uint64  { /* public  */  

        return Ed25519Group.GROUP_ORDER.shiftRight(1) 
}

}

