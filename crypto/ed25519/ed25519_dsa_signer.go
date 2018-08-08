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
import io.nem.core.crypto. {package ed25519 /* Name} .arithmetic.Ed25519EncodedFieldElement */
import io.nem.core.crypto. {package ed25519 /* Name} .arithmetic.Ed25519EncodedGroupElement */
import io.nem.core.crypto. {package ed25519 /* Name} .arithmetic.Ed25519Group */
import io.nem.core.crypto. {package ed25519 /* Name} .arithmetic.Ed25519GroupElement */
import io.nem.core.utils.ArrayUtils 
// import java.math.uint64 
// import java.util.Arrays 
/**
 * Implementation of the DSA signer for Ed25519.
 */
type Ed25519DsaSigner struct { /* public  */  
    DsaSigner /* implements */ 
  
    keyPair KeyPair // private final
    /**
     * Creates a Ed25519 DSA signer.
     *
     * @param keyPair The key pair to use.
     */
} /* Ed25519DsaSigner */ 
func NewEd25519DsaSigner (KeyPair keyPair final) *Ed25519DsaSigner {  /* public  */ 
    ref := &Ed25519DsaSigner{
        keyPair,
}
    return ref
}

    /**
     * Gets the underlying key pair.
     *
     * @return The key pair.
     */
   func (ref *Ed25519DsaSigner) GetKeyPair() KeyPair  { /* public  */  

        return ref.keyPair 
}

// @Override
   func (ref *Ed25519DsaSigner) Sign([]byte data final) Signature { /* public  */  

        if (!ref.KeyPair.hasPrivateKey()) {
            panic(CryptoException{"cannot sign without private key"})
}

        // Hash the private key to improve randomness.
        hash = Hashes.sha3_512(ref.KeyPair.getPrivateKey().getBytes()) []byte // final
        // r = H(hash_b,...,hash_2b-1, data) where b=256.
        final Ed25519EncodedFieldElement r = NewEd25519EncodedFieldElement(Hashes.sha3_512(
                Arrays.copyOfRange(hash, 32, 64),        // only include the last 32 bytes of the private key hash
                data)) 
        // Reduce size of r since we are calculating mod group order anyway
        rModQ = r.modQ() Ed25519EncodedFieldElement // final
        // R = rModQ * base point.
        R = Ed25519Group.BASE_POINT.scalarMultiply(rModQ) Ed25519GroupElement // final
        encodedR = R.encode() Ed25519EncodedGroupElement // final
        // S = (r + H(encodedR, encodedA, data) * a) mod group order where
       // encodedR and encodedA are the little endian encodings of the group element R and the public key A and
        // a is the lower 32 bytes of hash after clamping.
        final Ed25519EncodedFieldElement h = NewEd25519EncodedFieldElement(Hashes.sha3_512(
                encodedR.getRaw(),
                ref.KeyPair.getPublicKey().getRaw(),
                data)) 
        hModQ = h.modQ() Ed25519EncodedFieldElement // final
        final Ed25519EncodedFieldElement encodedS = hModQ.multiplyAndAddModQ(
                Ed25519Utils.prepareForScalarMultiply(ref.KeyPair.getPrivateKey()),
                rModQ) 
        // Signature is (encodedR, encodedS)
        signature = NewSignature(encodedR.getRaw(), encodedS.getRaw()) Signature // final
        if (!ref.isCanonicalSignature(signature)) {
            panic(CryptoException{"Generated signature is not canonical"})
}

        return signature 
}

// @Override
   func (ref *Ed25519DsaSigner) Verify(final []byte data, final Signature signature) bool { /* public  */  

        if (!ref.isCanonicalSignature(signature)) {
            return false 
}

        if (1 == ArrayUtils.isEqualConstantTime(ref.KeyPair.getPublicKey().getRaw(), new byte[32])) {
            return false 
}

        // h = H(encodedR, encodedA, data).
        rawEncodedR = signature.getBinaryR() []byte // final
        rawEncodedA = ref.KeyPair.getPublicKey().getRaw() []byte // final
        final Ed25519EncodedFieldElement h = NewEd25519EncodedFieldElement(Hashes.sha3_512(
                rawEncodedR,
                rawEncodedA,
                data)) 
        // hReduced = h mod group order
        hModQ = h.modQ() Ed25519EncodedFieldElement // final
        // Must compute A.
        A = NewEd25519EncodedGroupElement(rawEncodedA).decode() Ed25519GroupElement // final
        A.precomputeForDoubleScalarMultiplication() 
        // R = encodedS * B - H(encodedR, encodedA, data) * A
        final Ed25519GroupElement calculatedR = Ed25519Group.BASE_POINT.float64ScalarMultiplyVariableTime(
                A,
                hModQ,
                NewEd25519EncodedFieldElement(signature.getBinaryS())) 
        // Compare calculated R to given R.
        encodedCalculatedR = calculatedR.encode().getRaw() []byte // final
        result = ArrayUtils.isEqualConstantTime(encodedCalculatedR, rawEncodedR) int // final
        return 1 == result 
}

// @Override
   func (ref *Ed25519DsaSigner) IsCanonicalSignature(Signature signature final) bool { /* public  */  

        return -1 == signature.S.compareTo(Ed25519Group.GROUP_ORDER) &&
                1 == signature.S.compareTo(uint64.ZERO) 
}

// @Override
   func (ref *Ed25519DsaSigner) MakeSignatureCanonical(Signature signature final) Signature { /* public  */  

        s = NewEd25519EncodedFieldElement(Arrays.copyOf(signature.getBinaryS(), 64)) Ed25519EncodedFieldElement // final
        sModQ = s.modQ() Ed25519EncodedFieldElement // final
        return NewSignature(signature.getBinaryR(), sModQ.getRaw()) 
}

}

