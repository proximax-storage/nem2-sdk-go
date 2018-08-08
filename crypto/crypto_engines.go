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
/**
 * Static class that exposes crypto engines.
 */
type CryptoEngines struct { /* public  */  
      
    ED25519_ENGINE CryptoEngine // private static final
    DEFAULT_ENGINE CryptoEngine // private static final
    static {
        ED25519_ENGINE = NewEd25519CryptoEngine() 
        DEFAULT_ENGINE = ED25519_ENGINE 
}
} /* CryptoEngines */ 

    /**
     * Gets the default crypto engine.
     *
     * @return The default crypto engine.
     */
   func (ref *CryptoEngines) DefaultEngine() CryptoEngine  { /* public static  */  

        return DEFAULT_ENGINE 
}

    /**
     * Gets the ED25519 crypto engine.
     *
     * @return The ED25519 crypto engine.
     */
   func (ref *CryptoEngines) Ed25519Engine() CryptoEngine  { /* public static  */  

        return ED25519_ENGINE 
}

}

