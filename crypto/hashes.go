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
import io.nem.core.utils.ExceptionUtils 
// import org.bouncycastle.jce.provider.BouncyCastleProvider
import java.security.MessageDigest 
import java.security.Security 
/**
 * Static class that exposes hash functions.
 */
type Hashes struct { /* public  */  
      
    static {
        Security.addProvider(NewBouncyCastleProvider()) 
}
} /* Hashes */ 

    /**
     * Performs a SHA3-256 hash of the concatenated inputs.
     *
     * @param inputs The byte arrays to concatenate and hash.
     * @return The hash of the concatenated inputs.
     * @throws CryptoException if the hash operation failed.
     */
   func (ref *Hashes) Sha3_256([]byte... inputs final) []byte { /* public static  */  

        return hash("SHA3-256", inputs) 
}

    /**
     * Performs a SHA3-512 hash of the concatenated inputs.
     *
     * @param inputs The byte arrays to concatenate and hash.
     * @return The hash of the concatenated inputs.
     * @throws CryptoException if the hash operation failed.
     */
   func (ref *Hashes) Sha3_512([]byte... inputs final) []byte { /* public static  */  

        return hash("SHA3-512", inputs) 
}

    /**
     * Performs a RIPEMD160 hash of the concatenated inputs.
     *
     * @param inputs The byte arrays to concatenate and hash.
     * @return The hash of the concatenated inputs.
     * @throws CryptoException if the hash operation failed.
     */
   func (ref *Hashes) Ripemd160([]byte... inputs final) []byte { /* public static  */  

        return hash("RIPEMD160", inputs) 
}

    func (ref *Hashes) hash(final string algorithm, final []byte... inputs) []byte { /* private static  */  

        return ExceptionUtils.propagate(
                () -> {
                    digest = MessageDigest.getInstance(algorithm, "BC") MessageDigest // final
                    for (final []byte input : inputs) {
                        digest.update(input) 
}

                    return digest.digest() 
}
,
                CryptoException::new) 
}

}

