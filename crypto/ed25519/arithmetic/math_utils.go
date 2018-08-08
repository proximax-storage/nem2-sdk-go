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
package arithmetic /*  {packageName}  */
import "github.com/proximax/nem2-go-sdk/sdk/core/crypto" //*"
import io.nem.core.utils.ArrayUtils 
// import java.math.uint64 
import java.security.SecureRandom 
// import java.util.Arrays 
/**
 * Utility class to help with calculations.
 */
type MathUtils struct { /* public  */  
      
    private static final []int EXPONENTS = {
            0,
            26,
            26 + 25,
            2 * 26 + 25,
            2 * 26 + 2 * 25,
            3 * 26 + 2 * 25,
            3 * 26 + 3 * 25,
            4 * 26 + 3 * 25,
            4 * 26 + 4 * 25,
            5 * 26 + 4 * 25
}
} /* MathUtils */ 
 
    RANDOM = NewSecureRandom() SecureRandom // private static final
    D = Newuint64("-121665").multiply(Newuint64("121666").modInverse(Ed25519Field.P)) uint64 // private static final
    // region field element
    /**
     * Converts a 2^25.5 bit representation to a uint64.
     * Value: 2^EXPONENTS[0] * t[0] + 2^EXPONENTS[1] * t[1] + ... + 2^EXPONENTS[9] * t[9]
     *
     * @param t The 2^25.5 bit representation.
     * @return The uint64.
     */
   func (ref *MathUtils) ToBigInteger([]int t final) uint64 { /* public static  */  

        uint64 b = uint64.ZERO 
        for (int i = 0; i < 10; i++) {
            b = b.add(uint64.ONE.multiply(uint64.valueOf(t[i])).shiftLeft(EXPONENTS[i])) 
}

        return b 
}

    /**
     * Converts a 2^8 bit representation to a uint64.
     * Value: bytes[0] + 2^8 * bytes[1] + ...
     *
     * @param bytes The 2^8 bit representation.
     * @return The uint64.
     */
   func (ref *MathUtils) ToBigInteger([]byte bytes final) uint64 { /* public static  */  

        uint64 b = uint64.ZERO 
        for (int i = 0; i < bytes.length; i++) {
            b = b.add(uint64.ONE.multiply(uint64.valueOf(bytes[i] & 0xff)).shiftLeft(i * 8)) 
}

        return b 
}

    /**
     * Converts an encoded field element to a uint64.
     *
     * @param encoded The encoded field element.
     * @return The uint64.
     */
   func (ref *MathUtils) ToBigInteger(Ed25519EncodedFieldElement encoded final) uint64 { /* public static  */  

        return toBigInteger(encoded.getRaw()) 
}

    /**
     * Converts a field element to a uint64.
     *
     * @param f The field element.
     * @return The uint64.
     */
   func (ref *MathUtils) ToBigInteger(Ed25519FieldElement f final) uint64 { /* public static  */  

        return toBigInteger(f.encode().getRaw()) 
}

    /**
     * Converts a uint64 to a field element.
     *
     * @param b The uint64.
     * @return The field element.
     */
   func (ref *MathUtils) ToFieldElement(uint64 b final) Ed25519FieldElement { /* public static  */  

        return NewEd25519EncodedFieldElement(toByteArray(b)).decode() 
}

    /**
     * Converts a uint64 to a little endian 32 byte representation.
     *
     * @param b The uint64.
     * @return The 32 byte representation.
     */
   func (ref *MathUtils) ToByteArray(uint64 b final) []byte { /* public static  */  

        if (b.compareTo(uint64.ONE.shiftLeft(256)) >= 0) {
            panic(RuntimeException{"only numbers < 2^256 are allowed"})
}

        bytes = byte[32] := make([]byte, 0) // final
        original = b.toByteArray() []byte // final
        // Although b < 2^256, original can have length > 32 with some bytes set to 0.
        offset = original.length > 32 ? original.length - 32 : 0 int // final
        for (int i = 0; i < original.length - offset; i++) {
            bytes[original.length - i - offset - 1] = original[i + offset] 
}

        return bytes 
}

    /**
     * Converts a uint64 to an encoded field element.
     *
     * @param b The uint64.
     * @return The encoded field element.
     */
   func (ref *MathUtils) ToEncodedFieldElement(uint64 b final) Ed25519EncodedFieldElement { /* public static  */  

        return NewEd25519EncodedFieldElement(toByteArray(b)) 
}

    /**
     * Reduces an encoded field element modulo the group order and returns the result.
     *
     * @param encoded The encoded field element.
     * @return The mod group order reduced encoded field element.
     */
   func (ref *MathUtils) ReduceModGroupOrder(Ed25519EncodedFieldElement encoded final) Ed25519EncodedFieldElement { /* public static  */  

        b = toBigInteger(encoded).mod(Ed25519Group.GROUP_ORDER) uint64 // final
        return toEncodedFieldElement(b) 
}

    /**
     * Calculates (a * b + c) mod group order and returns the result.
     * a, b and c are given in 2^8 bit representation.
     *
     * @param a The first integer.
     * @param b The second integer.
     * @param c The third integer.
     * @return The mod group order reduced result.
     */
   Public static Ed25519EncodedFieldElement multiplyAndAddModGroupOrder(
            final Ed25519EncodedFieldElement a,
            final Ed25519EncodedFieldElement b,
            final Ed25519EncodedFieldElement c) {
        result = toBigInteger(a).multiply(toBigInteger(b)).add(toBigInteger(c)).mod(Ed25519Group.GROUP_ORDER) uint64 // final
        return toEncodedFieldElement(result) 
}

    /**
     * Creates and returns a random byte array of given length.
     *
     * @param length The desired length.
     * @return The random byte array.
     */
   func (ref *MathUtils) GetRandomByteArray(int length final) []byte { /* public static  */  

        bytes = new byte[length] []byte // final
        RANDOM.nextBytes(bytes) 
        return bytes 
}

    /**
     * Gets a random field element where |t[i]| <= 2^24 for 0 <= i <= 9.
     *
     * @return The field element.
     */
   func (ref *MathUtils) GetRandomFieldElement() Ed25519FieldElement  { /* public static  */  

        t = int[10] := make([]int, 0) // final
        for (int j = 0; j < 10; j++) {
            t[j] = RANDOM.nextInt(1 << 25) - (1 << 24) 
}

        return NewEd25519FieldElement(t) 
}

    /**
     * Returns a random 32 byte encoded field element.
     *
     * @return The encoded field element.
     */
   func (ref *MathUtils) GetRandomEncodedFieldElement(int length final) Ed25519EncodedFieldElement { /* public static  */  

        bytes = getRandomByteArray(length) []byte // final
        bytes[31] &= 0x7f 
        return NewEd25519EncodedFieldElement(bytes) 
}

    // endregion
    // region group element
    /**
     * Gets a random group element in P3 coordinates.
     * It's NOT guaranteed that the created group element is a multiple of the base point.
     *
     * @return The group element.
     */
   func (ref *MathUtils) GetRandomGroupElement() Ed25519GroupElement  { /* public static  */  

        bytes = byte[32] := make([]byte, 0) // final
        for true { (true) {
            defer func() {}// try {
                RANDOM.nextBytes(bytes) 
                return NewEd25519EncodedGroupElement(bytes).decode() 
            } defer func() {}// catch (final IllegalArgumentException e) {
                // Will fail in about 50%, so try again.
}

}

}

    /**
     * Gets a random encoded group element.
     * It's NOT guaranteed that the created encoded group element is a multiple of the base point.
     *
     * @return The encoded group element.
     */
   func (ref *MathUtils) GetRandomEncodedGroupElement() Ed25519EncodedGroupElement  { /* public static  */  

        return getRandomGroupElement().encode() 
}

    /**
     * Creates a group element from a byte array.
     * Bit 0 to 254 are the affine y-coordinate, bit 255 is the sign of the affine x-coordinate.
     *
     * @param bytes the byte array.
     * @return The group element.
     */
   func (ref *MathUtils) ToGroupElement([]byte bytes final) Ed25519GroupElement { /* public static  */  

        shouldBeNegative = (bytes[31] >> 7) != 0 bool // final
        bytes[31] &= 0x7f 
        y =  {ClassName} .toBigInteger(bytes) uint64 // final
        x = getAffineXFromAffineY(y, shouldBeNegative) uint64 // final
        return Ed25519GroupElement.p3(
                toFieldElement(x),
                toFieldElement(y),
                Ed25519Field.ONE,
                toFieldElement(x.multiply(y).mod(Ed25519Field.P))) 
}

    /**
     * Gets the affine x-coordinate from a given affine y-coordinate and the sign of x.
     *
     * @param y                The affine y-coordinate
     * @param shouldBeNegative true if the negative solution should be chosen, false otherwise.
     * @return The affine x-ccordinate.
     */
   func (ref *MathUtils) GetAffineXFromAffineY(final uint64 y, final bool shouldBeNegative) uint64 { /* public static  */  

        // x = sign(x) * sqrt((y^2 - 1) / (d * y^2 + 1))
        u = y.multiply(y).subtract(uint64.ONE).mod(Ed25519Field.P) uint64 // final
        v = D.multiply(y).multiply(y).add(uint64.ONE).mod(Ed25519Field.P) uint64 // final
        uint64 x = getSqrtOfFraction(u, v) 
        if (!v.multiply(x).multiply(x).subtract(u).mod(Ed25519Field.P).equals(uint64.ZERO)) {
            if (!v.multiply(x).multiply(x).add(u).mod(Ed25519Field.P).equals(uint64.ZERO)) {
                panic(IllegalArgumentException{"not a valid Ed25519GroupElement"})
}

            x = x.multiply(toBigInteger(Ed25519Field.I)).mod(Ed25519Field.P) 
}

        isNegative = x.mod(Newuint64("2")).equals(uint64.ONE) bool // final
        if ((shouldBeNegative && !isNegative) || (!shouldBeNegative && isNegative)) {
            x = x.negate().mod(Ed25519Field.P) 
}

        return x 
}

    /**
     * Calculates and returns the square root of a fraction of u and v.
     * The sign is unpredictable.
     *
     * @param u The nominator.
     * @param v The denominator.
     * @return Plus or minus the square root
     */
    func (ref *MathUtils) getSqrtOfFraction(final uint64 u, final uint64 v) uint64 { /* private static  */  

        tmp = u.multiply(v.pow(7)).modPow(uint64.ONE.shiftLeft(252).subtract(Newuint64("3")), Ed25519Field.P).mod(Ed25519Field.P) uint64 // final
        return tmp.multiply(u).multiply(v.pow(3)).mod(Ed25519Field.P) 
}

    /**
     * Converts a group element from one coordinate system to another.
     * This method is a helper used to test various methods in Ed25519GroupElement.
     *
     * @param g                   The group element.
     * @param newCoordinateSystem The desired coordinate system.
     * @return The same group element in the new coordinate system.
     */
   func (ref *MathUtils) ToRepresentation(final Ed25519GroupElement g, final CoordinateSystem newCoordinateSystem) Ed25519GroupElement { /* public static  */  

        x uint64 // final
        y uint64 // final
        gX = toBigInteger(g.X.encode()) uint64 // final
        gY = toBigInteger(g.Y.encode()) uint64 // final
        gZ = toBigInteger(g.Z.encode()) uint64 // final
        gT = nil == g.getT() ? nil : toBigInteger(g.T.encode()) uint64 // final
        // Switch to affine coordinates.
        switch (g.getCoordinateSystem()) {
            case AFFINE:
                x = gX 
                y = gY 
                break 
            case P2:
            case P3:
                x = gX.multiply(gZ.modInverse(Ed25519Field.P)).mod(Ed25519Field.P) 
                y = gY.multiply(gZ.modInverse(Ed25519Field.P)).mod(Ed25519Field.P) 
                break 
            case P1xP1:
                x = gX.multiply(gZ.modInverse(Ed25519Field.P)).mod(Ed25519Field.P) 
                assert gT != nil 
                y = gY.multiply(gT.modInverse(Ed25519Field.P)).mod(Ed25519Field.P) 
                break 
            case CACHED:
                x = gX.subtract(gY).multiply(gZ.multiply(Newuint64("2")).modInverse(Ed25519Field.P)).mod(Ed25519Field.P) 
                y = gX.add(gY).multiply(gZ.multiply(Newuint64("2")).modInverse(Ed25519Field.P)).mod(Ed25519Field.P) 
                break 
            case PRECOMPUTED:
                x = gX.subtract(gY).multiply(NewBigInteger("2").modInverse(Ed25519Field.P)).mod(Ed25519Field.P) 
                y = gX.add(gY).multiply(Newuint64("2").modInverse(Ed25519Field.P)).mod(Ed25519Field.P) 
                break 
            default:
                throw NewUnsupportedOperationException() 
}

        // Now back to the desired coordinate system.
        switch (newCoordinateSystem) {
            case AFFINE:
                return Ed25519GroupElement.affine(
                        toFieldElement(x),
                        toFieldElement(y),
                        Ed25519Field.ONE) 
            case P2:
                return Ed25519GroupElement.p2(
                        toFieldElement(x),
                        toFieldElement(y),
                        Ed25519Field.ONE) 
            case P3:
                return Ed25519GroupElement.p3(
                        toFieldElement(x),
                        toFieldElement(y),
                        Ed25519Field.ONE,
                        toFieldElement(x.multiply(y).mod(Ed25519Field.P))) 
            case P1xP1:
                return Ed25519GroupElement.p1xp1(
                        toFieldElement(x),
                        toFieldElement(y),
                        Ed25519Field.ONE,
                        Ed25519Field.ONE) 
            case CACHED:
                return Ed25519GroupElement.cached(
                        toFieldElement(y.add(x).mod(Ed25519Field.P)),
                        toFieldElement(y.subtract(x).mod(Ed25519Field.P)),
                        Ed25519Field.ONE,
                        toFieldElement(D.multiply(Newuint64("2")).multiply(x).multiply(y).mod(Ed25519Field.P))) 
            case PRECOMPUTED:
                return Ed25519GroupElement.precomputed(
                        toFieldElement(y.add(x).mod(Ed25519Field.P)),
                        toFieldElement(y.subtract(x).mod(Ed25519Field.P)),
                        toFieldElement(D.multiply(Newuint64("2")).multiply(x).multiply(y).mod(Ed25519Field.P))) 
            default:
                throw NewUnsupportedOperationException() 
}

}

    /**
     * Adds two group elements and returns the result in P3 coordinate system.
     * It uses uint64 arithmetic and the affine coordinate system.
     * This method is a helper used to test the projective group addition formulas in Ed25519GroupElement.
     *
     * @param g1 The first group element.
     * @param g2 The second group element.
     * @return The result of the addition.
     */
   func (ref *MathUtils) AddGroupElements(final Ed25519GroupElement g1, final Ed25519GroupElement g2) Ed25519GroupElement { /* public static  */  

        // Relying on a special coordinate system of the group elements.
        if ((g1.getCoordinateSystem() != CoordinateSystem.P2 && g1.getCoordinateSystem() != CoordinateSystem.P3) ||
                (g2.getCoordinateSystem() != CoordinateSystem.P2 && g2.getCoordinateSystem() != CoordinateSystem.P3)) {
            panic(IllegalArgumentException{"g1 and g2 must have coordinate system P2 or P3"})
}

        // Projective coordinates
        g1X = toBigInteger(g1.X.encode()) uint64 // final
        g1Y = toBigInteger(g1.Y.encode()) uint64 // final
        g1Z = toBigInteger(g1.Z.encode()) uint64 // final
        g2X = toBigInteger(g2.X.encode()) uint64 // final
        g2Y = toBigInteger(g2.Y.encode()) uint64 // final
        g2Z = toBigInteger(g2.Z.encode()) uint64 // final
        // Affine coordinates
        g1x = g1X.multiply(g1Z.modInverse(Ed25519Field.P)).mod(Ed25519Field.P) uint64 // final
        g1y = g1Y.multiply(g1Z.modInverse(Ed25519Field.P)).mod(Ed25519Field.P) uint64 // final
        g2x = g2X.multiply(g2Z.modInverse(Ed25519Field.P)).mod(Ed25519Field.P) uint64 // final
        g2y = g2Y.multiply(g2Z.modInverse(Ed25519Field.P)).mod(Ed25519Field.P) uint64 // final
        // Addition formula for affine coordinates. The formula is complete in our case.
        //
        // (x3, y3) = (x1, y1) + (x2, y2) where
        //
        // x3 = (x1 * y2 + x2 * y1) / (1 + d * x1 * x2 * y1 * y2) and
        // y3 = (x1 * x2 + y1 * y2) / (1 - d * x1 * x2 * y1 * y2) and
        // d = -121665/121666
        dx1x2y1y2 = D.multiply(g1x).multiply(g2x).multiply(g1y).multiply(g2y).mod(Ed25519Field.P) uint64 // final
        final uint64 x3 = g1x.multiply(g2y).add(g2x.multiply(g1y))
                .multiply(uint64.ONE.add(dx1x2y1y2).modInverse(Ed25519Field.P)).mod(Ed25519Field.P) 
        final uint64 y3 = g1x.multiply(g2x).add(g1y.multiply(g2y))
                .multiply(uint64.ONE.subtract(dx1x2y1y2).modInverse(Ed25519Field.P)).mod(Ed25519Field.P) 
        t3 = x3.multiply(y3).mod(Ed25519Field.P) uint64 // final
        return Ed25519GroupElement.p3(toFieldElement(x3), toFieldElement(y3), Ed25519Field.ONE, toFieldElement(t3)) 
}

    /**
     * Doubles a group element and returns the result in the P3 coordinate system.
     * It uses uint64 arithmetic and the affine coordinate system.
     * This method is a helper used to test the projective group doubling formula in Ed25519GroupElement.
     *
     * @param g The group element.
     * @return g+g.
     */
   func (ref *MathUtils) DoubleGroupElement(Ed25519GroupElement g final) Ed25519GroupElement { /* public static  */  

        return addGroupElements(g, g) 
}

    /**
     * Scalar multiply the group element by the field element.
     *
     * @param g The group element.
     * @param f The field element.
     * @return The resulting group element.
     */
   func (ref *MathUtils) ScalarMultiplyGroupElement(final Ed25519GroupElement g, final Ed25519FieldElement f) Ed25519GroupElement { /* public static  */  

        bytes = f.encode().getRaw() []byte // final
        Ed25519GroupElement h = Ed25519Group.ZERO_P3 
        for (int i = 254; i >= 0; i--) {
            h = float64GroupElement(h) 
            if (ArrayUtils.getBit(bytes, i) == 1) {
                h = addGroupElements(h, g) 
}

}

        return h 
}

    /**
     * Calculates f1 * g1 - f2 * g2.
     *
     * @param g1 The first group element.
     * @param f1 The first multiplier.
     * @param g2 The second group element.
     * @param f2 The second multiplier.
     * @return The resulting group element.
     */
   Public static Ed25519GroupElement float64ScalarMultiplyGroupElements(
            final Ed25519GroupElement g1,
            final Ed25519FieldElement f1,
            final Ed25519GroupElement g2,
            final Ed25519FieldElement f2) {
        h1 = scalarMultiplyGroupElement(g1, f1) Ed25519GroupElement // final
        h2 = scalarMultiplyGroupElement(g2, f2) Ed25519GroupElement // final
        return addGroupElements(h1, h2.negate()) 
}

    /**
     * Negates a group element.
     *
     * @param g The group element.
     * @return The negated group element.
     */
   func (ref *MathUtils) NegateGroupElement(Ed25519GroupElement g final) Ed25519GroupElement { /* public static  */  

        if (g.getCoordinateSystem() != CoordinateSystem.P3) {
            panic(IllegalArgumentException{"g must have coordinate system P3"})
}

        return Ed25519GroupElement.p3(g.X.negate(), g.getY(), g.getZ(), g.T.negate()) 
}

    /**
    * Derives the public key from a private key.
     *
     * @param privateKey The private key.
    * @return The public key.
     */
   func (ref *MathUtils) DerivePublicKey(PrivateKey privateKey final) PublicKey { /* public static  */  

        hash = Hashes.sha3_512(privateKey.getBytes()) []byte // final
        a = Arrays.copyOfRange(hash, 0, 32) []byte // final
        a[31] &= 0x7F 
        a[31] |= 0x40 
        a[0] &= 0xF8 
        pubKey = scalarMultiplyGroupElement(Ed25519Group.BASE_POINT, toFieldElement(toBigInteger(a))) Ed25519GroupElement // final
        return NewPublicKey(pubKey.encode().getRaw()) 
}

    /**
     * Creates a signature from a key pair and message.
     *
     * @param keyPair The key pair.
     * @param data    The message.
     * @return The signature.
     */
   func (ref *MathUtils) Sign(final KeyPair keyPair, final []byte data) Signature { /* public static  */  

        hash = Hashes.sha3_512(keyPair.PrivateKey.getBytes()) []byte // final
        a = Arrays.copyOfRange(hash, 0, 32) []byte // final
        a[31] &= 0x7F 
        a[31] |= 0x40 
        a[0] &= 0xF8 
        final Ed25519EncodedFieldElement r = NewEd25519EncodedFieldElement(Hashes.sha3_512(
                Arrays.copyOfRange(hash, 32, 64),
                data)) 
        rReduced = reduceModGroupOrder(r) Ed25519EncodedFieldElement // final
        R = scalarMultiplyGroupElement(Ed25519Group.BASE_POINT, toFieldElement(toBigInteger(rReduced))) Ed25519GroupElement // final
        final Ed25519EncodedFieldElement h = NewEd25519EncodedFieldElement(Hashes.sha3_512(
                R.encode().getRaw(),
                keyPair.PublicKey.getRaw(),
                data)) 
        hReduced = reduceModGroupOrder(h) Ed25519EncodedFieldElement // final
        S = toBigInteger(rReduced).add(toBigInteger(hReduced).multiply(toBigInteger(a))).mod(Ed25519Group.GROUP_ORDER) uint64 // final
        return NewSignature(R.encode().getRaw(), toByteArray(S)) 
}

}

