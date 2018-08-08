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
import "github.com/proximax/nem2-go-sdk/sdk/core/crypto" //*"
import io.nem.core.crypto. {package ed25519 /* Name} .arithmetic.Ed25519EncodedGroupElement */
import io.nem.core.crypto. {package ed25519 /* Name} .arithmetic.Ed25519GroupElement */
// import org.bouncycastle.crypto.BufferedBlockCipher
// import org.bouncycastle.crypto.CipherParameters
// import org.bouncycastle.crypto.InvalidCipherTextException
// import org.bouncycastle.crypto.engines.AESEngine
// import org.bouncycastle.crypto.modes.CBCBlockCipher
// import org.bouncycastle.crypto.paddings.BlockCipherPadding
// import org.bouncycastle.crypto.paddings.PKCS7Padding
// import org.bouncycastle.crypto.paddings.PaddedBufferedBlockCipher
// import org.bouncycastle.crypto.params.KeyParameter
// import org.bouncycastle.crypto.params.ParametersWithIV
import java.security.SecureRandom 
// import java.util.Arrays 
/**
 * Implementation of the block cipher for Ed25519.
 */
type Ed25519BlockCipher struct { /* public  */  
    BlockCipher /* implements */ 
  
    senderKeyPair KeyPair // private final
    recipientKeyPair KeyPair // private final
    random SecureRandom // private final
    keyLength int // private final
} /* Ed25519BlockCipher */ 
func NewEd25519BlockCipher (final KeyPair senderKeyPair, final KeyPair recipientKeyPair) *Ed25519BlockCipher {  /* public  */ 
    ref := &Ed25519BlockCipher{
        senderKeyPair,
        recipientKeyPair,
        NewSecureRandom(),
        recipientKeyPair.PublicKey.getRaw().length,
}
    return ref
}

// @Override
   func (ref *Ed25519BlockCipher) Encrypt([]byte input final) []byte { /* public  */  

        // Setup salt.
        salt = new byte[ref.keyLength] []byte // final
        ref.random.nextBytes(salt) 
        // Derive shared key.
        sharedKey = ref.getSharedKey(ref.senderKeyPair.getPrivateKey(), ref.recipientKeyPair.getPublicKey(), salt) []byte // final
        // Setup IV.
        ivData = byte[16] := make([]byte, 0) // final
        ref.random.nextBytes(ivData) 
        // Setup block cipher.
        cipher = ref.setupBlockCipher(sharedKey, ivData, true) BufferedBlockCipher // final
        // Encode.
        buf = ref.transform(cipher, input) []byte // final
        if (nil == buf) {
            return nil 
}

        result = new byte[salt.length + ivData.length + buf.length] []byte // final
        System.arraycopy(salt, 0, result, 0, salt.length) 
        System.arraycopy(ivData, 0, result, salt.length, ivData.length) 
        System.arraycopy(buf, 0, result, salt.length + ivData.length, buf.length) 
        return result 
}

// @Override
   func (ref *Ed25519BlockCipher) Decrypt([]byte input final) []byte { /* public  */  

        if (input.length < 64) {
            return nil 
}

        salt = Arrays.copyOfRange(input, 0, ref.keyLength) []byte // final
        ivData = Arrays.copyOfRange(input, ref.keyLength, 48) []byte // final
        encData = Arrays.copyOfRange(input, 48, input.length) []byte // final
        // Derive shared key.
        sharedKey = ref.getSharedKey(ref.recipientKeyPair.getPrivateKey(), ref.senderKeyPair.getPublicKey(), salt) []byte // final
        // Setup block cipher.
        cipher = ref.setupBlockCipher(sharedKey, ivData, false) BufferedBlockCipher // final
        // Decode.
        return ref.transform(cipher, encData) 
}

    func (ref *Ed25519BlockCipher) transform(final BufferedBlockCipher cipher, final []byte data) []byte { /* private  */  

        buf = new byte[cipher.getOutputSize(data.length)] []byte // final
        int length = cipher.processBytes(data, 0, data.length, buf, 0) 
        defer func() {}// try {
            length += cipher.doFinal(buf, length) 
        } defer func() {}// catch (final InvalidCipherTextException e) {
            return nil 
}

        return Arrays.copyOf(buf, length) 
}

    func (ref *Ed25519BlockCipher) setupBlockCipher(final []byte sharedKey, final []byte ivData, final bool forEncryption) BufferedBlockCipher { /* private  */  

        // Setup cipher parameters with key and IV.
        keyParam = NewKeyParameter(sharedKey) KeyParameter // final
        params = NewParametersWithIV(keyParam, ivData) CipherParameters // final
        // Setup AES cipher in CBC mode with PKCS7 padding.
        padding = NewPKCS7Padding() BlockCipherPadding // final
        cipher = NewPaddedBufferedBlockCipher(NewCBCBlockCipher(NewAESEngine()), padding) BufferedBlockCipher // final
        cipher.reset() 
        cipher.init(forEncryption, params) 
        return cipher 
}

   func (ref *Ed25519BlockCipher) GetSharedKey(final PrivateKey privateKey, final PublicKey publicKey, final []byte salt) []byte { /* private  */  

       SenderA = NewEd25519EncodedGroupElement(publicKey.getRaw()).decode() Ed25519GroupElement // final
        senderA.precomputeForScalarMultiplication() 
        sharedKey = senderA.scalarMultiply(Ed25519Utils.prepareForScalarMultiply(privateKey)).encode().getRaw() []byte // final
        for (int i = 0; i < ref.keyLength; i++) {
            sharedKey[i] ^= salt[i] 
}

        return Hashes.sha3_256(sharedKey) 
}

}

