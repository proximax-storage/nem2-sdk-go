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
 * Exception that is used when a cryptographic operation fails.
 */
type CryptoException struct { /* public  */  
    RuntimeException /* extends */ 
  
    /**
     * Creates a new crypto exception.
     *
     * @param message The exception message.
     */
} /* CryptoException */ 
func NewCryptoException (string message final) *CryptoException {  /* public  */ 
    ref := &CryptoException{
        super(message) 
}
    return ref
}

    /**
     * Creates a new crypto exception.
     *
     * @param cause The exception cause.
     */
func NewCryptoException (Throwable cause final) *CryptoException {  /* public  */ 
    ref := &CryptoException{
        super(cause) 
}
    return ref
}

    /**
     * Creates a new crypto exception.
     *
     * @param message The exception message.
     * @param cause   The exception cause.
     */
func NewCryptoException (final string message, final Throwable cause) *CryptoException {  /* public  */ 
    ref := &CryptoException{
        super(message, cause) 
}
    return ref
}

}

