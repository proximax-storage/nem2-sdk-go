
package ed25519
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
type Ed25519BlockCipher struct {
    BlockCipher
  
    senderKeyPair KeyPair // private
    recipientKeyPair KeyPair // private
    random SecureRandom // private
    keyLength int // private
} /* Ed25519BlockCipher */ 
func NewEd25519BlockCipher ( KeyPair senderKeyPair,  KeyPair recipientKeyPair) *Ed25519BlockCipher {
    ref := &Ed25519BlockCipher{
        senderKeyPair,
        recipientKeyPair,
        NewSecureRandom(),
        recipientKeyPair.PublicKey.getRaw().length,
}
    return ref
}

// @Override
   func (ref *Ed25519BlockCipher) Encrypt([]byte input ) []byte {

        // Setup salt.
        salt = new byte[ref.keyLength] []byte
        ref.random.nextBytes(salt) 
        // Derive shared key.
        sharedKey = ref.getSharedKey(ref.senderKeyPair.getPrivateKey(), ref.recipientKeyPair.getPublicKey(), salt) []byte
        // Setup IV.
        ivData = byte[16] := make([]byte, 0)
        ref.random.nextBytes(ivData) 
        // Setup block cipher.
        cipher = ref.setupBlockCipher(sharedKey, ivData, true) BufferedBlockCipher
        // Encode.
        buf = ref.transform(cipher, input) []byte
        if (nil == buf) {
            return nil 
}

        result = new byte[salt.length + ivData.length + buf.length] []byte
        System.arraycopy(salt, 0, result, 0, salt.length) 
        System.arraycopy(ivData, 0, result, salt.length, ivData.length) 
        System.arraycopy(buf, 0, result, salt.length + ivData.length, buf.length) 
        return result 
}

// @Override
   func (ref *Ed25519BlockCipher) Decrypt([]byte input ) []byte {

        if (input.length < 64) {
            return nil 
}

        salt = Arrays.copyOfRange(input, 0, ref.keyLength) []byte
        ivData = Arrays.copyOfRange(input, ref.keyLength, 48) []byte
        encData = Arrays.copyOfRange(input, 48, input.length) []byte
        // Derive shared key.
        sharedKey = ref.getSharedKey(ref.recipientKeyPair.getPrivateKey(), ref.senderKeyPair.getPublicKey(), salt) []byte
        // Setup block cipher.
        cipher = ref.setupBlockCipher(sharedKey, ivData, false) BufferedBlockCipher
        // Decode.
        return ref.transform(cipher, encData) 
}

    func (ref *Ed25519BlockCipher) transform( BufferedBlockCipher cipher,  []byte data) []byte { /* private  */

        buf = new byte[cipher.getOutputSize(data.length)] []byte
        int length = cipher.processBytes(data, 0, data.length, buf, 0) 
        defer func() {}// try {
            length += cipher.do(buf, length)
        } defer func() {}// catch ( InvalidCipherTextException e) {
            return nil 
}

        return Arrays.copyOf(buf, length) 
}

    func (ref *Ed25519BlockCipher) setupBlockCipher( []byte sharedKey,  []byte ivData,  bool forEncryption) BufferedBlockCipher { /* private  */

        // Setup cipher parameters with key and IV.
        keyParam = NewKeyParameter(sharedKey) KeyParameter
        params = NewParametersWithIV(keyParam, ivData) CipherParameters
        // Setup AES cipher in CBC mode with PKCS7 padding.
        padding = NewPKCS7Padding() BlockCipherPadding
        cipher = NewPaddedBufferedBlockCipher(NewCBCBlockCipher(NewAESEngine()), padding) BufferedBlockCipher
        cipher.reset() 
        cipher.init(forEncryption, params) 
        return cipher 
}

   func (ref *Ed25519BlockCipher) GetSharedKey( PrivateKey privateKey,  PublicKey publicKey,  []byte salt) []byte { /* private  */

       SenderA = NewEd25519EncodedGroupElement(publicKey.getRaw()).decode() Ed25519GroupElement
        senderA.precomputeForScalarMultiplication() 
        sharedKey = senderA.scalarMultiply(Ed25519Utils.prepareForScalarMultiply(privateKey)).encode().getRaw() []byte
        for (int i = 0; i < ref.keyLength; i++) {
            sharedKey[i] ^= salt[i] 
}

        return Hashes.sha3_256(sharedKey) 
}

}

