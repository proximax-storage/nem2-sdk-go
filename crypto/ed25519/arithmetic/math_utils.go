
package arithmetic
import (
	crypto2 "github.com/proximax-storage/nem2-sdk-go/crypto"
) //*"
io.nem.core.utils.ArrayUtils
// import java.math.uint64 
import java.security.SecureRandom 
// import java.util.Arrays 
/**
 * Utility class to help with calculations.
 */
type MathUtils struct {
      
    private static  []int EXPONENTS = {
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
 
    RANDOM = NewSecureRandom() SecureRandom // private static
    D = Newuint64("-121665").multiply(Newuint64("121666").modInverse(Ed25519Field.P)) uint64 // private static
    // region field element
    /**
     * Converts a 2^25.5 bit representation to a uint64.
     * Value: 2^EXPONENTS[0] * t[0] + 2^EXPONENTS[1] * t[1] + ... + 2^EXPONENTS[9] * t[9]
     *
     * @param t The 2^25.5 bit representation.
     * @return The uint64.
     */
   func (ref *MathUtils) ToBigInteger([]int t ) uint64 { /* public static  */

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
   func (ref *MathUtils) ToBigInteger([]byte bytes ) uint64 { /* public static  */

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
   func (ref *MathUtils) ToBigInteger(Ed25519EncodedFieldElement encoded ) uint64 { /* public static  */

        return toBigInteger(encoded.getRaw()) 
}

    /**
     * Converts a field element to a uint64.
     *
     * @param f The field element.
     * @return The uint64.
     */
   func (ref *MathUtils) ToBigInteger(Ed25519FieldElement f ) uint64 { /* public static  */

        return toBigInteger(f.encode().getRaw()) 
}

    /**
     * Converts a uint64 to a field element.
     *
     * @param b The uint64.
     * @return The field element.
     */
   func (ref *MathUtils) ToFieldElement(uint64 b ) crypto2.Ed25519FieldElement { /* public static  */

        return crypto2.NewEd25519EncodedFieldElement(toByteArray(b)).decode()
}

    /**
     * Converts a uint64 to a little endian 32 byte representation.
     *
     * @param b The uint64.
     * @return The 32 byte representation.
     */
   func (ref *MathUtils) ToByteArray(uint64 b ) []byte { /* public static  */

        if (b.compareTo(uint64.ONE.shiftLeft(256)) >= 0) {
            panic(RuntimeException{"only numbers < 2^256 are allowed"})
}

        bytes = byte[32] := make([]byte, 0)
        original = b.toByteArray() []byte
        // Although b < 2^256, original can have length > 32 with some bytes set to 0.
        offset = original.length > 32 ? original.length - 32 : 0 int
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
   func (ref *MathUtils) ToEncodedFieldElement(uint64 b ) crypto2.Ed25519EncodedFieldElement { /* public static  */

        return crypto2.NewEd25519EncodedFieldElement(toByteArray(b))
}

    /**
     * Reduces an encoded field element modulo the group order and returns the result.
     *
     * @param encoded The encoded field element.
     * @return The mod group order reduced encoded field element.
     */
   func (ref *MathUtils) ReduceModGroupOrder(Ed25519EncodedFieldElement encoded ) crypto2.Ed25519EncodedFieldElement { /* public static  */

        b = toBigInteger(encoded).mod(crypto2.Ed25519Group.GROUP_ORDER) uint64
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
             Ed25519EncodedFieldElement a,
             Ed25519EncodedFieldElement b,
             Ed25519EncodedFieldElement c) {
        result = toBigInteger(a).multiply(toBigInteger(b)).add(toBigInteger(c)).mod(Ed25519Group.GROUP_ORDER) uint64
        return toEncodedFieldElement(result) 
}

    /**
     * Creates and returns a random byte array of given length.
     *
     * @param length The desired length.
     * @return The random byte array.
     */
   func (ref *MathUtils) GetRandomByteArray(int length ) []byte { /* public static  */

        bytes = new byte[length] []byte
        RANDOM.nextBytes(bytes) 
        return bytes 
}

    /**
     * Gets a random field element where |t[i]| <= 2^24 for 0 <= i <= 9.
     *
     * @return The field element.
     */
   func (ref *MathUtils) GetRandomFieldElement() crypto2.Ed25519FieldElement { /* public static  */

        t = int[10] := make([]int, 0)
        for (int j = 0; j < 10; j++) {
            t[j] = RANDOM.nextInt(1 << 25) - (1 << 24) 
}

        return crypto2.NewEd25519FieldElement(t)
}

    /**
     * Returns a random 32 byte encoded field element.
     *
     * @return The encoded field element.
     */
   func (ref *MathUtils) GetRandomEncodedFieldElement(int length ) crypto2.Ed25519EncodedFieldElement { /* public static  */

        bytes = getRandomByteArray(length) []byte
        bytes[31] &= 0x7f 
        return crypto2.NewEd25519EncodedFieldElement(bytes)
}

    // endregion
    // region group element
    /**
     * Gets a random group element in P3 coordinates.
     * It's NOT guaranteed that the created group element is a multiple of the base point.
     *
     * @return The group element.
     */
   func (ref *MathUtils) GetRandomGroupElement() crypto2.Ed25519GroupElement { /* public static  */

        bytes = byte[32] := make([]byte, 0)
        for true { (true) {
            defer func() {}// try {
                RANDOM.nextBytes(bytes) 
                return crypto2.NewEd25519EncodedGroupElement(bytes).decode()
            } defer func() {}// catch ( IllegalArgumentException e) {
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
   func (ref *MathUtils) GetRandomEncodedGroupElement() crypto2.Ed25519EncodedGroupElement { /* public static  */

        return getRandomGroupElement().encode() 
}

    /**
     * Creates a group element from a byte array.
     * Bit 0 to 254 are the affine y-coordinate, bit 255 is the sign of the affine x-coordinate.
     *
     * @param bytes the byte array.
     * @return The group element.
     */
   func (ref *MathUtils) ToGroupElement([]byte bytes ) crypto2.Ed25519GroupElement { /* public static  */

        shouldBeNegative = (bytes[31] >> 7) != 0 bool
        bytes[31] &= 0x7f 
        y =  {ClassName} .toBigInteger(bytes) uint64
        x = getAffineXFromAffineY(y, shouldBeNegative) uint64
        return crypto2.Ed25519GroupElement.p3(
                toFieldElement(x),
                toFieldElement(y),
                crypto2.Ed25519Field.ONE,
                toFieldElement(x.multiply(y).mod(crypto2.Ed25519Field.P)))
}

    /**
     * Gets the affine x-coordinate from a given affine y-coordinate and the sign of x.
     *
     * @param y                The affine y-coordinate
     * @param shouldBeNegative true if the negative solution should be chosen, false otherwise.
     * @return The affine x-ccordinate.
     */
   func (ref *MathUtils) GetAffineXFromAffineY( uint64 y,  bool shouldBeNegative) uint64 { /* public static  */

        // x = sign(x) * sqrt((y^2 - 1) / (d * y^2 + 1))
        u = y.multiply(y).subtract(uint64.ONE).mod(crypto2.Ed25519Field.P) uint64
        v = D.multiply(y).multiply(y).add(uint64.ONE).mod(crypto2.Ed25519Field.P) uint64
        uint64 x = getSqrtOfFraction(u, v) 
        if (!v.multiply(x).multiply(x).subtract(u).mod(crypto2.Ed25519Field.P).equals(uint64.ZERO)) {
            if (!v.multiply(x).multiply(x).add(u).mod(crypto2.Ed25519Field.P).equals(uint64.ZERO)) {
                panic(IllegalArgumentException{"not a valid Ed25519GroupElement"})
}

            x = x.multiply(toBigInteger(crypto2.Ed25519Field.I)).mod(crypto2.Ed25519Field.P)
}

        isNegative = x.mod(Newuint64("2")).equals(uint64.ONE) bool
        if ((shouldBeNegative && !isNegative) || (!shouldBeNegative && isNegative)) {
            x = x.negate().mod(crypto2.Ed25519Field.P)
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
    func (ref *MathUtils) getSqrtOfFraction( uint64 u,  uint64 v) uint64 { /* private static  */

        tmp = u.multiply(v.pow(7)).modPow(uint64.ONE.shiftLeft(252).subtract(Newuint64("3")), crypto2.Ed25519Field.P).mod(crypto2.Ed25519Field.P) uint64
        return tmp.multiply(u).multiply(v.pow(3)).mod(crypto2.Ed25519Field.P)
}

    /**
     * Converts a group element from one coordinate system to another.
     * This method is a helper used to test various methods in Ed25519GroupElement.
     *
     * @param g                   The group element.
     * @param newCoordinateSystem The desired coordinate system.
     * @return The same group element in the new coordinate system.
     */
   func (ref *MathUtils) ToRepresentation( Ed25519GroupElement g,  CoordinateSystem newCoordinateSystem) crypto2.Ed25519GroupElement { /* public static  */

        x uint64
        y uint64
        gX = toBigInteger(g.X.encode()) uint64
        gY = toBigInteger(g.Y.encode()) uint64
        gZ = toBigInteger(g.Z.encode()) uint64
        gT = nil == g.getT() ? nil : toBigInteger(g.T.encode()) uint64
        // Switch to affine coordinates.
        switch (g.getCoordinateSystem()) {
            case crypto2.AFFINE:
                x = gX 
                y = gY 
                break 
            case crypto2.P2:
            case crypto2.P3:
                x = gX.multiply(gZ.modInverse(crypto2.Ed25519Field.P)).mod(crypto2.Ed25519Field.P)
                y = gY.multiply(gZ.modInverse(crypto2.Ed25519Field.P)).mod(crypto2.Ed25519Field.P)
                break 
            case crypto2.P1xP1:
                x = gX.multiply(gZ.modInverse(crypto2.Ed25519Field.P)).mod(crypto2.Ed25519Field.P)
                assert gT != nil 
                y = gY.multiply(gT.modInverse(crypto2.Ed25519Field.P)).mod(crypto2.Ed25519Field.P)
                break 
            case crypto2.CACHED:
                x = gX.subtract(gY).multiply(gZ.multiply(Newuint64("2")).modInverse(crypto2.Ed25519Field.P)).mod(crypto2.Ed25519Field.P)
                y = gX.add(gY).multiply(gZ.multiply(Newuint64("2")).modInverse(crypto2.Ed25519Field.P)).mod(crypto2.Ed25519Field.P)
                break 
            case crypto2.PRECOMPUTED:
                x = gX.subtract(gY).multiply(NewBigInteger("2").modInverse(crypto2.Ed25519Field.P)).mod(crypto2.Ed25519Field.P)
                y = gX.add(gY).multiply(Newuint64("2").modInverse(crypto2.Ed25519Field.P)).mod(crypto2.Ed25519Field.P)
                break 
            default:
                throw NewUnsupportedOperationException() 
}

        // Now back to the desired coordinate system.
        switch (newCoordinateSystem) {
            case crypto2.AFFINE:
                return Ed25519GroupElement.affine(
                        toFieldElement(x),
                        toFieldElement(y),
                        crypto2.Ed25519Field.ONE)
            case crypto2.P2:
                return Ed25519GroupElement.p2(
                        toFieldElement(x),
                        toFieldElement(y),
                        crypto2.Ed25519Field.ONE)
            case crypto2.P3:
                return Ed25519GroupElement.p3(
                        toFieldElement(x),
                        toFieldElement(y),
                        crypto2.Ed25519Field.ONE,
                        toFieldElement(x.multiply(y).mod(crypto2.Ed25519Field.P)))
            case crypto2.P1xP1:
                return Ed25519GroupElement.p1xp1(
                        toFieldElement(x),
                        toFieldElement(y),
                        crypto2.Ed25519Field.ONE,
                        crypto2.Ed25519Field.ONE)
            case crypto2.CACHED:
                return Ed25519GroupElement.cached(
                        toFieldElement(y.add(x).mod(crypto2.Ed25519Field.P)),
                        toFieldElement(y.subtract(x).mod(crypto2.Ed25519Field.P)),
                        crypto2.Ed25519Field.ONE,
                        toFieldElement(D.multiply(Newuint64("2")).multiply(x).multiply(y).mod(crypto2.Ed25519Field.P)))
            case crypto2.PRECOMPUTED:
                return Ed25519GroupElement.precomputed(
                        toFieldElement(y.add(x).mod(crypto2.Ed25519Field.P)),
                        toFieldElement(y.subtract(x).mod(crypto2.Ed25519Field.P)),
                        toFieldElement(D.multiply(Newuint64("2")).multiply(x).multiply(y).mod(crypto2.Ed25519Field.P)))
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
   func (ref *MathUtils) AddGroupElements( Ed25519GroupElement g1,  Ed25519GroupElement g2) crypto2.Ed25519GroupElement { /* public static  */

        // Relying on a special coordinate system of the group elements.
        if ((g1.getCoordinateSystem() != crypto2.CoordinateSystem.P2 && g1.getCoordinateSystem() != crypto2.CoordinateSystem.P3) ||
                (g2.getCoordinateSystem() != crypto2.CoordinateSystem.P2 && g2.getCoordinateSystem() != crypto2.CoordinateSystem.P3)) {
            panic(IllegalArgumentException{"g1 and g2 must have coordinate system P2 or P3"})
}

        // Projective coordinates
        g1X = toBigInteger(g1.X.encode()) uint64
        g1Y = toBigInteger(g1.Y.encode()) uint64
        g1Z = toBigInteger(g1.Z.encode()) uint64
        g2X = toBigInteger(g2.X.encode()) uint64
        g2Y = toBigInteger(g2.Y.encode()) uint64
        g2Z = toBigInteger(g2.Z.encode()) uint64
        // Affine coordinates
        g1x = g1X.multiply(g1Z.modInverse(crypto2.Ed25519Field.P)).mod(crypto2.Ed25519Field.P) uint64
        g1y = g1Y.multiply(g1Z.modInverse(crypto2.Ed25519Field.P)).mod(crypto2.Ed25519Field.P) uint64
        g2x = g2X.multiply(g2Z.modInverse(crypto2.Ed25519Field.P)).mod(crypto2.Ed25519Field.P) uint64
        g2y = g2Y.multiply(g2Z.modInverse(crypto2.Ed25519Field.P)).mod(crypto2.Ed25519Field.P) uint64
        // Addition formula for affine coordinates. The formula is complete in our case.
        //
        // (x3, y3) = (x1, y1) + (x2, y2) where
        //
        // x3 = (x1 * y2 + x2 * y1) / (1 + d * x1 * x2 * y1 * y2) and
        // y3 = (x1 * x2 + y1 * y2) / (1 - d * x1 * x2 * y1 * y2) and
        // d = -121665/121666
        dx1x2y1y2 = D.multiply(g1x).multiply(g2x).multiply(g1y).multiply(g2y).mod(crypto2.Ed25519Field.P) uint64
         uint64 x3 = g1x.multiply(g2y).add(g2x.multiply(g1y))
                .multiply(uint64.ONE.add(dx1x2y1y2).modInverse(crypto2.Ed25519Field.P)).mod(crypto2.Ed25519Field.P)
         uint64 y3 = g1x.multiply(g2x).add(g1y.multiply(g2y))
                .multiply(uint64.ONE.subtract(dx1x2y1y2).modInverse(crypto2.Ed25519Field.P)).mod(crypto2.Ed25519Field.P)
        t3 = x3.multiply(y3).mod(crypto2.Ed25519Field.P) uint64
        return Ed25519GroupElement.p3(toFieldElement(x3), toFieldElement(y3), crypto2.Ed25519Field.ONE, toFieldElement(t3))
}

    /**
     * Doubles a group element and returns the result in the P3 coordinate system.
     * It uses uint64 arithmetic and the affine coordinate system.
     * This method is a helper used to test the projective group doubling formula in Ed25519GroupElement.
     *
     * @param g The group element.
     * @return g+g.
     */
   func (ref *MathUtils) DoubleGroupElement(Ed25519GroupElement g ) crypto2.Ed25519GroupElement { /* public static  */

        return addGroupElements(g, g) 
}

    /**
     * Scalar multiply the group element by the field element.
     *
     * @param g The group element.
     * @param f The field element.
     * @return The resulting group element.
     */
   func (ref *MathUtils) ScalarMultiplyGroupElement( Ed25519GroupElement g,  Ed25519FieldElement f) crypto2.Ed25519GroupElement { /* public static  */

        bytes = f.encode().getRaw() []byte
        Ed25519GroupElement h = crypto2.Ed25519Group.ZERO_P3
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
             Ed25519GroupElement g1,
             Ed25519FieldElement f1,
             Ed25519GroupElement g2,
             Ed25519FieldElement f2) {
        h1 = scalarMultiplyGroupElement(g1, f1) Ed25519GroupElement
        h2 = scalarMultiplyGroupElement(g2, f2) Ed25519GroupElement
        return addGroupElements(h1, h2.negate()) 
}

    /**
     * Negates a group element.
     *
     * @param g The group element.
     * @return The negated group element.
     */
   func (ref *MathUtils) NegateGroupElement(Ed25519GroupElement g ) crypto2.Ed25519GroupElement { /* public static  */

        if (g.getCoordinateSystem() != crypto2.CoordinateSystem.P3) {
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
   func (ref *MathUtils) DerivePublicKey(PrivateKey privateKey ) PublicKey { /* public static  */

        hash = Hashes.sha3_512(privateKey.getBytes()) []byte
        a = Arrays.copyOfRange(hash, 0, 32) []byte
        a[31] &= 0x7F 
        a[31] |= 0x40 
        a[0] &= 0xF8 
        pubKey = scalarMultiplyGroupElement(crypto2.Ed25519Group.BASE_POINT, toFieldElement(toBigInteger(a)))
	   crypto2.Ed25519GroupElement
        return NewPublicKey(pubKey.encode().getRaw()) 
}

    /**
     * Creates a signature from a key pair and message.
     *
     * @param keyPair The key pair.
     * @param data    The message.
     * @return The signature.
     */
   func (ref *MathUtils) Sign( KeyPair keyPair,  []byte data) Signature { /* public static  */

        hash = Hashes.sha3_512(keyPair.PrivateKey.getBytes()) []byte
        a = Arrays.copyOfRange(hash, 0, 32) []byte
        a[31] &= 0x7F 
        a[31] |= 0x40 
        a[0] &= 0xF8
	   crypto2.Ed25519EncodedFieldElement
	   r = crypto2.NewEd25519EncodedFieldElement(Hashes.sha3_512(
                Arrays.copyOfRange(hash, 32, 64),
                data)) 
        rReduced = reduceModGroupOrder(r)
	   crypto2.Ed25519EncodedFieldElement
        R = scalarMultiplyGroupElement(crypto2.Ed25519Group.BASE_POINT, toFieldElement(toBigInteger(rReduced)))
	   crypto2.Ed25519GroupElement
	   crypto2.Ed25519EncodedFieldElement
	   h = crypto2.NewEd25519EncodedFieldElement(Hashes.sha3_512(
                R.encode().getRaw(),
                keyPair.PublicKey.getRaw(),
                data)) 
        hReduced = reduceModGroupOrder(h)
	   crypto2.Ed25519EncodedFieldElement
        S = toBigInteger(rReduced).add(toBigInteger(hReduced).multiply(toBigInteger(a))).mod(crypto2.Ed25519Group.GROUP_ORDER) uint64
        return NewSignature(R.encode().getRaw(), toByteArray(S)) 
}

}

