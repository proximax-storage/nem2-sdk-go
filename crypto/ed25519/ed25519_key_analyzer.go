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
import "github.com/proximax/nem2-go-sdk/sdk/core/crypto" //keyanalyzer"
import "github.com/proximax/nem2-go-sdk/sdk/core/crypto" //publickey"
/**
 * Implementation of the key analyzer for Ed25519.
 */
type Ed25519KeyAnalyzer struct { /* public  */  
    KeyAnalyzer /* implements */ 
  
    COMPRESSED_KEY_SIZE = 32 int // private static final
// @Override
   func (ref *Ed25519KeyAnalyzer) IsKeyCompressed(PublicKey publicKey final) bool { /* public  */  

       Return COMPRESSED_KEY_SIZE == publicKey.Raw.length 
}
} /* Ed25519KeyAnalyzer */ 

}

