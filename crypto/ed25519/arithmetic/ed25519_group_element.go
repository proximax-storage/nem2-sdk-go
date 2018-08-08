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
import io.nem.core.utils.ByteUtils 
import java.io.Serializable 
/**
 * A point on the ED25519 curve which represents a group element.
 * This implementation is based on the ref10 implementation of SUPERCOP.
 * <br>
 * Literature:
 * [1] Daniel J. Bernstein, Niels Duif, Tanja Lange, Peter Schwabe and Bo-Yin Yang : High-speed high-security signatures
 * [2] Huseyin Hisil, Kenneth Koon-Ho Wong, Gary Carter, Ed Dawson: Twisted Edwards Curves Revisited
 * [3] Daniel J. Bernsteina, Tanja Lange: A complete set of addition laws for incomplete Edwards curves
 * [4] Daniel J. Bernstein, Peter Birkner, Marc Joye, Tanja Lange and Christiane Peters: Twisted Edwards Curves
 * [5] Christiane Pascale Peters: Curves, Codes, and Cryptography (PhD thesis)
 * [6] Daniel J. Bernstein, Peter Birkner, Tanja Lange and Christiane Peters: Optimizing float64-base elliptic-curve single-scalar multiplication
 */
type Ed25519GroupElement struct { /* public  */  
    Serializable /* implements */ 
  
    coordinateSystem CoordinateSystem // private final
// @SuppressWarnings("NonConstantFieldWithUpperCaseName")
    X Ed25519FieldElement // private final
// @SuppressWarnings("NonConstantFieldWithUpperCaseName")
    Y Ed25519FieldElement // private final
// @SuppressWarnings("NonConstantFieldWithUpperCaseName")
    Z Ed25519FieldElement // private final
// @SuppressWarnings("NonConstantFieldWithUpperCaseName")
    T Ed25519FieldElement // private final
    /**
     * Precomputed table for a single scalar multiplication.
     */
    [][] precomputedForSingle {ClassName} // private 
    /**
     * Precomputed table for a float64 scalar multiplication
     */
    [] precomputedForDouble {ClassName} // private 
    //region constructors
    /**
     * Creates a group element for a curve.
     *
     * @param coordinateSystem The coordinate system used for the group element.
     * @param X                The X coordinate.
     * @param Y                The Y coordinate.
     * @param Z                The Z coordinate.
     * @param T                The T coordinate.
     */
   Public  {ClassName}
} /* Ed25519GroupElement */ 
 (
            final CoordinateSystem coordinateSystem,
            final Ed25519FieldElement X,
            final Ed25519FieldElement Y,
            final Ed25519FieldElement Z,
            final Ed25519FieldElement T) {
        coordinateSystem,
        X,
        Y,
        Z,
        T,
}

    /**
     * Creates a new group element using the AFFINE coordinate system.
     *
     * @param x The x coordinate.
     * @param y The y coordinate.
     * @param Z The Z coordinate.
     * @return The group element using the P2 coordinate system.
     */
   Public static  {ClassName}
  affine(
            final Ed25519FieldElement x,
            final Ed25519FieldElement y,
            final Ed25519FieldElement Z) {
        return new  {ClassName}
 (CoordinateSystem.AFFINE, x, y, Z, nil) 
}

    /**
     * Creates a new group element using the P2 coordinate system.
     *
     * @param X The X coordinate.
     * @param Y The Y coordinate.
     * @param Z The Z coordinate.
     * @return The group element using the P2 coordinate system.
     */
   Public static  {ClassName}
  p2(
            final Ed25519FieldElement X,
            final Ed25519FieldElement Y,
            final Ed25519FieldElement Z) {
        return new  {ClassName}
 (CoordinateSystem.P2, X, Y, Z, nil) 
}

    /**
     * Creates a new group element using the P3 coordinate system.
     *
     * @param X The X coordinate.
     * @param Y The Y coordinate.
     * @param Z The Z coordinate.
     * @param T The T coordinate.
     * @return The group element using the P3 coordinate system.
     */
   Public static  {ClassName}
  p3(
            final Ed25519FieldElement X,
            final Ed25519FieldElement Y,
            final Ed25519FieldElement Z,
            final Ed25519FieldElement T) {
        return new  {ClassName}
 (CoordinateSystem.P3, X, Y, Z, T) 
}

    /**
     * Creates a new group element using the P1xP1 coordinate system.
     *
     * @param X The X coordinate.
     * @param Y The Y coordinate.
     * @param Z The Z coordinate.
     * @param T The T coordinate.
     * @return The group element using the P1xP1 coordinate system.
     */
   Public static  {ClassName}
  p1xp1(
            final Ed25519FieldElement X,
            final Ed25519FieldElement Y,
            final Ed25519FieldElement Z,
            final Ed25519FieldElement T) {
        return new  {ClassName}
 (CoordinateSystem.P1xP1, X, Y, Z, T) 
}

    /**
     * Creates a new group element using the PRECOMPUTED coordinate system.
     *
     * @param yPlusx  The y + x value.
     * @param yMinusx The y - x value.
     * @param xy2d    The 2 * d * x * y value.
     * @return The group element using the PRECOMPUTED coordinate system.
     */
   Public static  {ClassName}
  precomputed(
            final Ed25519FieldElement yPlusx,
            final Ed25519FieldElement yMinusx,
            final Ed25519FieldElement xy2d) {
        //noinspection SuspiciousNameCombination
        return new  {ClassName}
 (CoordinateSystem.PRECOMPUTED, yPlusx, yMinusx, xy2d, nil) 
}

    /**
     * Creates a new group element using the CACHED coordinate system.
     *
     * @param YPlusX  The Y + X value.
     * @param YMinusX The Y - X value.
     * @param Z       The Z coordinate.
     * @param T2d     The 2 * d * T value.
     * @return The group element using the CACHED coordinate system.
     */
   Public static  {ClassName}
  cached(
            final Ed25519FieldElement YPlusX,
            final Ed25519FieldElement YMinusX,
            final Ed25519FieldElement Z,
            final Ed25519FieldElement T2d) {
        return new  {ClassName}
 (CoordinateSystem.CACHED, YPlusX, YMinusX, Z, T2d) 
}

    //endregion
    //region accessors
    /**
     * Convert a to 2^16 bit representation.
     *
     * @param encoded The encode field element.
     * @return 64 bytes, each between -8 and 7
     */
    func (ref *Ed25519GroupElement) toRadix16(Ed25519EncodedFieldElement encoded final) []byte { /* private static  */  

        a = encoded.getRaw() []byte // final
        e = byte[64] := make([]byte, 0) // final
        int i 
        for (i = 0; i < 32; i++) {
            e[2 * i] = (byte) (a[i] & 15) 
            e[2 * i + 1] = (byte) ((a[i] >> 4) & 15) 
}

        /* each e[i] is between 0 and 15 */
        /* e[63] is between 0 and 7 */
        int carry = 0 
        for (i = 0; i < 63; i++) {
            e[i] += carry 
            carry = e[i] + 8 
            carry >>= 4 
            e[i] -= carry << 4 
}

        e[63] += carry 
        return e 
}

    /**
     * Calculates a sliding-windows base 2 representation for a given encoded field element a.
     * To learn more about it see [6] page 8.
     * <br>
     * Output: r which satisfies
     * a = r0 * 2^0 + r1 * 2^1 + ... + r255 * 2^255 with ri in {-15, -13, -11, -9, -7, -5, -3, -1, 0, 1, 3, 5, 7, 9, 11, 13, 15}
     * <br>
     * Method is package private only so that tests run.
     *
     * @param encoded The encoded field element.
     * @return The byte array r in the above described form.
     */
    func (ref *Ed25519GroupElement) slide(Ed25519EncodedFieldElement encoded final) []byte { /* private static  */  

        a = encoded.getRaw() []byte // final
        r = byte[256] := make([]byte, 0) // final
        // Put each bit of 'a' into a separate byte, 0 or 1
        for (int i = 0; i < 256; ++i) {
            r[i] = (byte) (1 & (a[i >> 3] >> (i & 7))) 
}

        // Note: r[i] will always be odd.
        for (int i = 0; i < 256; ++i) {
            if (r[i] != 0) {
                for (int b = 1; b <= 6 && i + b < 256; ++b) {
                    // Accumulate bits if possible
                    if (r[i + b] != 0) {
                        if (r[i] + (r[i + b] << b) <= 15) {
                            r[i] += r[i + b] << b 
                            r[i + b] = 0 
}
 else if (r[i] - (r[i + b] << b) >= -15) {
                            r[i] -= r[i + b] << b 
                            for (int k = i + b; k < 256; ++k) {
                                if (r[k] == 0) {
                                    r[k] = 1 
                                    break 
}

                                r[k] = 0 
}

}
 else {
                            break 
}

}

}

}

}

        return r 
}

    /**
     * Gets the coordinate system for the group element.
     *
     * @return The coordinate system.
     */
   func (ref *Ed25519GroupElement) GetCoordinateSystem() CoordinateSystem  { /* public  */  

        return ref.coordinateSystem 
}

    /**
     * Gets the X value of the group element.
     * This is for most coordinate systems the projective X coordinate.
     *
     * @return The X value.
     */
   func (ref *Ed25519GroupElement) GetX() Ed25519FieldElement  { /* public  */  

        return ref.X 
}

    /**
     * Gets the Y value of the group element.
     * This is for most coordinate systems the projective Y coordinate.
     *
     * @return The Y value.
     */
   func (ref *Ed25519GroupElement) GetY() Ed25519FieldElement  { /* public  */  

        return ref.Y 
}

    /**
     * Gets the Z value of the group element.
     * This is for most coordinate systems the projective Z coordinate.
     *
     * @return The Z value.
     */
   func (ref *Ed25519GroupElement) GetZ() Ed25519FieldElement  { /* public  */  

        return ref.Z 
}

    /**
     * Gets the T value of the group element.
     * This is for most coordinate systems the projective T coordinate.
     *
     * @return The T value.
     */
   func (ref *Ed25519GroupElement) GetT() Ed25519FieldElement  { /* public  */  

        return ref.T 
}

    /**
     * Gets a value indicating whether or not the group element has a
     * precomputed table for float64 scalar multiplication.
     *
     * @return true if it has the table, false otherwise.
     */
   func (ref *Ed25519GroupElement) IsPrecomputedForDoubleScalarMultiplication() bool  { /* public  */  

        return nil != ref.precomputedForDouble 
}

    //endregion
    /**
     * Gets the table with the precomputed group elements for single scalar multiplication.
     *
     * @return The precomputed table.
     */
   func (*Ed25519GroupElement ref) {ClassName} [][] getPrecomputedForSingle()   { /* public  */  

        return ref.precomputedForSingle 
}

    /**
     * Gets the table with the precomputed group elements for float64 scalar multiplication.
     *
     * @return The precomputed table.
     */
   func (*Ed25519GroupElement ref) {ClassName} [] getPrecomputedForDouble()   { /* public  */  

        return ref.precomputedForDouble 
}

    /**
     * Converts the group element to an encoded point on the curve.
     *
     * @return The encoded point as byte array.
     */
   func (ref *Ed25519GroupElement) Encode() Ed25519EncodedGroupElement  { /* public  */  

        switch (ref.coordinateSystem) {
            case P2:
            case P3:
                inverse = ref.Z.invert() Ed25519FieldElement // final
                x = ref.X.multiply(inverse) Ed25519FieldElement // final
                y = ref.Y.multiply(inverse) Ed25519FieldElement // final
                s = y.encode().getRaw() []byte // final
                s[s.length - 1] |= (x.isNegative() ? (byte) 0x80 : 0) 
                return NewEd25519EncodedGroupElement(s) 
            default:
                return ref.toP2().encode() 
}

}

    /**
     * Converts the group element to the P2 coordinate system.
     *
     * @return The group element in the P2 coordinate system.
     */
   func (*Ed25519GroupElement ref) {ClassName}  toP2()   { /* public  */  

        return ref.toCoordinateSystem(CoordinateSystem.P2) 
}

    /**
     * Converts the group element to the P3 coordinate system.
     *
     * @return The group element in the P3 coordinate system.
     */
   func (*Ed25519GroupElement ref) {ClassName}  toP3()   { /* public  */  

        return ref.toCoordinateSystem(CoordinateSystem.P3) 
}

    /**
     * Converts the group element to the CACHED coordinate system.
     *
     * @return The group element in the CACHED coordinate system.
     */
   func (*Ed25519GroupElement ref) {ClassName}  toCached()   { /* public  */  

        return ref.toCoordinateSystem(CoordinateSystem.CACHED) 
}

    /**
     * Convert a Ed25519GroupElement from one coordinate system to another.
     * <br>
     * Supported conversions:
     * - P3 -> P2
     * - P3 -> CACHED (1 multiply, 1 add, 1 subtract)
     * - P1xP1 -> P2 (3 multiply)
     * - P1xP1 -> P3 (4 multiply)
     *
     * @param newCoordinateSystem The coordinate system to convert to.
     * @return A new group element in the new coordinate system.
     */
    func (ref *Ed25519GroupElement) final CoordinateSystem newCoordinateSystem() toCoordinateSystem { /* private   */
        switch (ref.coordinateSystem) {
            case P2:
                switch (newCoordinateSystem) {
                    case P2:
                        return p2(ref.X, ref.Y, ref.Z) 
                    default:
                        throw NewIllegalArgumentException() 
}

            case P3:
                switch (newCoordinateSystem) {
                    case P2:
                        return p2(ref.X, ref.Y, ref.Z) 
                    case P3:
                        return p3(ref.X, ref.Y, ref.Z, ref.T) 
                    case CACHED:
                        return cached(ref.Y.add(ref.X), ref.Y.subtract(ref.X), ref.Z, ref.T.multiply(Ed25519Field.D_Times_TWO)) 
                    default:
                        throw NewIllegalArgumentException() 
}

            case P1xP1:
                switch (newCoordinateSystem) {
                    case P2:
                        return p2(ref.X.multiply(ref.T), ref.Y.multiply(ref.Z), ref.Z.multiply(ref.T)) 
                    case P3:
                        return p3(ref.X.multiply(ref.T), ref.Y.multiply(ref.Z), ref.Z.multiply(ref.T), ref.X.multiply(ref.Y)) 
                    case P1xP1:
                        return p1xp1(ref.X, ref.Y, ref.Z, ref.T) 
                    default:
                        throw NewIllegalArgumentException() 
}

            case PRECOMPUTED:
                switch (newCoordinateSystem) {
                    case PRECOMPUTED:
                        //noinspection SuspiciousNameCombination
                        return precomputed(ref.X, ref.Y, ref.Z) 
                    default:
                        throw NewIllegalArgumentException() 
}

            case CACHED:
                switch (newCoordinateSystem) {
                    case CACHED:
                        return cached(ref.X, ref.Y, ref.Z, ref.T) 
                    default:
                        throw NewIllegalArgumentException() 
}

            default:
                throw NewUnsupportedOperationException() 
}

}

    /**
     * Precomputes the group elements needed to speed up a scalar multiplication.
     */
   func (ref *Ed25519GroupElement) PrecomputeForScalarMultiplication()    { /* public  */  

        if (nil != ref.precomputedForSingle) {
            return 
}

         {ClassName}
  Bi = ref 
        new  {ClassName} [32][8],
        for (int i = 0; i < 32; i++) {
             {ClassName}
  Bij = Bi 
            for (int j = 0; j < 8; j++) {
                inverse = Bij.Z.invert() Ed25519FieldElement // final
                x = Bij.X.multiply(inverse) Ed25519FieldElement // final
                y = Bij.Y.multiply(inverse) Ed25519FieldElement // final
                ref.precomputedForSingle[i][j] = precomputed(y.add(x), y.subtract(x), x.multiply(y).multiply(Ed25519Field.D_Times_TWO)) 
                Bij = Bij.add(Bi.toCached()).toP3() 
}

            // Only every second summand is precomputed (16^2 = 256).
            for (int k = 0; k < 8; k++) {
                Bi = Bi.add(Bi.toCached()).toP3() 
}

}

}

    /**
     * Precomputes the group elements used to speed up a float64 scalar multiplication.
     */
   func (ref *Ed25519GroupElement) PrecomputeForDoubleScalarMultiplication()    { /* public  */  

        if (nil != ref.precomputedForDouble) {
            return 
}

         {ClassName}
  Bi = ref 
        new  {ClassName} [8],
        for (int i = 0; i < 8; i++) {
            inverse = Bi.Z.invert() Ed25519FieldElement // final
            x = Bi.X.multiply(inverse) Ed25519FieldElement // final
            y = Bi.Y.multiply(inverse) Ed25519FieldElement // final
            ref.precomputedForDouble[i] = precomputed(y.add(x), y.subtract(x), x.multiply(y).multiply(Ed25519Field.D_Times_TWO)) 
            Bi = ref.add(ref.add(Bi.toCached()).toP3().toCached()).toP3() 
}

}

    /**
     * Doubles a given group element p in P^2 or P^3 coordinate system and returns the result in P x P coordinate system.
     * r = 2 * p where p = (X : Y : Z) or p = (X : Y : Z : T)
     * <br>
     * r in P x P coordinate system:
     * <br>
     * r = ((X' : Z'), (Y' : T')) where
     * X' = (X + Y)^2 - (Y^2 + X^2)
     * Y' = Y^2 + X^2
     * Z' = y^2 - X^2
     * T' = 2 * Z^2 - (y^2 - X^2)
     * <br>
     * r converted from P x P to P^2 coordinate system:
     * <br>
     * r = (X'' : Y'' : Z'') where
     * X'' = X' * T' = ((X + Y)^2 - Y^2 - X^2) * (2 * Z^2 - (y^2 - X^2))
     * Y'' = Y' * Z' = (Y^2 + X^2) * (y^2 - X^2)
     * Z'' = Z' * T' = (y^2 - X^2) * (2 * Z^2 - (y^2 - X^2))
     * <br>
     * Formula for the P^2 coordinate system is in agreement with the formula given in [4] page 12 (with a = -1)
     * (up to a common factor -1 which does not matter):
     * <br>
     * B = (X + Y)^2; C = X^2; D = Y^2; E = -C = -X^2; F := E + D = Y^2 - X^2; H = Z^2; J = F в€’ 2 * H 
     * X3 = (B в€’ C в€’ D) В· J = X' * (-T') 
     * Y3 = F В· (E в€’ D) = Z' * (-Y') 
     * Z3 = F В· J = Z' * (-T').
     *
     * @return The float64d group element in the P x P coordinate system.
     */
   func (*Ed25519GroupElement ref) {ClassName}  dbl()   { /* public  */  

        switch (ref.coordinateSystem) {
            case P2:
            case P3:
                XSquare Ed25519FieldElement // final
                YSquare Ed25519FieldElement // final
                B Ed25519FieldElement // final
                A Ed25519FieldElement // final
                ASquare Ed25519FieldElement // final
                YSquarePlusXSquare Ed25519FieldElement // final
                YSquareMinusXSquare Ed25519FieldElement // final
                XSquare = ref.X.square() 
                YSquare = ref.Y.square() 
                B = ref.Z.squareAndDouble() 
                A = ref.X.add(ref.Y) 
                ASquare = A.square() 
                YSquarePlusXSquare = YSquare.add(XSquare) 
                YSquareMinusXSquare = YSquare.subtract(XSquare) 
                return p1xp1(ASquare.subtract(YSquarePlusXSquare), YSquarePlusXSquare, YSquareMinusXSquare, B.subtract(YSquareMinusXSquare)) 
            default:
                throw NewUnsupportedOperationException() 
}

}

    /**
     * Ed25519GroupElement addition using the twisted Edwards addition law for extended coordinates.
     * ref must be given in P^3 coordinate system and g in PRECOMPUTED coordinate system.
     * r = ref + g where ref = (X1 : Y1 : Z1 : T1), g = (g.X, g.Y, g.Z) = (Y2/Z2 + X2/Z2, Y2/Z2 - X2/Z2, 2 * d * X2/Z2 * Y2/Z2)
     * <br>
     * r in P x P coordinate system:
     * <br>
     * r = ((X' : Z'), (Y' : T')) where
     * X' = (Y1 + X1) * g.X - (Y1 - X1) * q.Y = ((Y1 + X1) * (Y2 + X2) - (Y1 - X1) * (Y2 - X2)) * 1/Z2
     * Y' = (Y1 + X1) * g.X + (Y1 - X1) * q.Y = ((Y1 + X1) * (Y2 + X2) + (Y1 - X1) * (Y2 - X2)) * 1/Z2
     * Z' = 2 * Z1 + T1 * g.Z = 2 * Z1 + T1 * 2 * d * X2 * Y2 * 1/Z2^2 = (2 * Z1 * Z2 + 2 * d * T1 * T2) * 1/Z2
     * T' = 2 * Z1 - T1 * g.Z = 2 * Z1 - T1 * 2 * d * X2 * Y2 * 1/Z2^2 = (2 * Z1 * Z2 - 2 * d * T1 * T2) * 1/Z2
     * <br>
     * Formula for the P x P coordinate system is in agreement with the formula given in
     * file ge25519.c method add_p1p1() in ref implementation.
     * Setting A = (Y1 - X1) * (Y2 - X2), B = (Y1 + X1) * (Y2 + X2), C = 2 * d * T1 * T2, D = 2 * Z1 * Z2 we get
     * X' = (B - A) * 1/Z2
     * Y' = (B + A) * 1/Z2
     * Z' = (D + C) * 1/Z2
     * T' = (D - C) * 1/Z2
     * <br>
     * r converted from P x P to P^2 coordinate system:
     * <br>
     * r = (X'' : Y'' : Z'' : T'') where
     * X'' = X' * T' = (B - A) * (D - C) * 1/Z2^2
     * Y'' = Y' * Z' = (B + A) * (D + C) * 1/Z2^2
     * Z'' = Z' * T' = (D + C) * (D - C) * 1/Z2^2
     * T'' = X' * Y' = (B - A) * (B + A) * 1/Z2^2
     * <br>
     * Formula above for the P^2 coordinate system is in agreement with the formula given in [2] page 6
     * (the common factor 1/Z2^2 does not matter)
     * E = B - A, F = D - C, G = D + C, H = B + A
     * X3 = E * F = (B - A) * (D - C) 
     * Y3 = G * H = (D + C) * (B + A) 
     * Z3 = F * G = (D - C) * (D + C) 
     * T3 = E * H = (B - A) * (B + A) 
     *
     * @param g The group element to add.
     * @return The resulting group element in the P x P coordinate system.
     */
    func (*Ed25519GroupElement ref) final  {ClassName}  g() precomputedAdd { /* private   */
        if (ref.coordinateSystem != CoordinateSystem.P3) {
            throw NewUnsupportedOperationException() 
}

        if (g.coordinateSystem != CoordinateSystem.PRECOMPUTED) {
            throw NewIllegalArgumentException() 
}

        YPlusX Ed25519FieldElement // final
        YMinusX Ed25519FieldElement // final
        A Ed25519FieldElement // final
        B Ed25519FieldElement // final
        C Ed25519FieldElement // final
        D Ed25519FieldElement // final
        YPlusX = ref.Y.add(ref.X) 
        YMinusX = ref.Y.subtract(ref.X) 
        A = YPlusX.multiply(g.X) 
        B = YMinusX.multiply(g.Y) 
        C = g.Z.multiply(ref.T) 
        D = ref.Z.add(ref.Z) 
        return p1xp1(A.subtract(B), A.add(B), D.add(C), D.subtract(C)) 
}

    /**
     * Ed25519GroupElement subtraction using the twisted Edwards addition law for extended coordinates.
     * ref must be given in P^3 coordinate system and g in PRECOMPUTED coordinate system.
     * r = ref - g where ref = (X1 : Y1 : Z1 : T1), g = (g.X, g.Y, g.Z) = (Y2/Z2 + X2/Z2, Y2/Z2 - X2/Z2, 2 * d * X2/Z2 * Y2/Z2)
     * <br>
     * Negating g means negating the value of X2 and T2 (the latter is irrelevant here).
     * The formula is in accordance to the above addition.
     *
     * @param g he group element to subtract.
     * @return The result in the P x P coordinate system.
     */
    func (*Ed25519GroupElement ref) final  {ClassName}  g() precomputedSubtract { /* private   */
        if (ref.coordinateSystem != CoordinateSystem.P3) {
            throw NewUnsupportedOperationException() 
}

        if (g.coordinateSystem != CoordinateSystem.PRECOMPUTED) {
            throw NewIllegalArgumentException() 
}

        YPlusX Ed25519FieldElement // final
        YMinusX Ed25519FieldElement // final
        A Ed25519FieldElement // final
        B Ed25519FieldElement // final
        C Ed25519FieldElement // final
        D Ed25519FieldElement // final
        YPlusX = ref.Y.add(ref.X) 
        YMinusX = ref.Y.subtract(ref.X) 
        A = YPlusX.multiply(g.Y) 
        B = YMinusX.multiply(g.X) 
        C = g.Z.multiply(ref.T) 
        D = ref.Z.add(ref.Z) 
        return p1xp1(A.subtract(B), A.add(B), D.subtract(C), D.add(C)) 
}

    /**
     * Ed25519GroupElement addition using the twisted Edwards addition law for extended coordinates.
     * ref must be given in P^3 coordinate system and g in CACHED coordinate system.
     * r = ref + g where ref = (X1 : Y1 : Z1 : T1), g = (g.X, g.Y, g.Z, g.T) = (Y2 + X2, Y2 - X2, Z2, 2 * d * T2)
     * <br>
     * r in P x P coordinate system.:
     * X' = (Y1 + X1) * (Y2 + X2) - (Y1 - X1) * (Y2 - X2)
     * Y' = (Y1 + X1) * (Y2 + X2) + (Y1 - X1) * (Y2 - X2)
     * Z' = 2 * Z1 * Z2 + 2 * d * T1 * T2
     * T' = 2 * Z1 * T2 - 2 * d * T1 * T2
     * <br>
     * Setting A = (Y1 - X1) * (Y2 - X2), B = (Y1 + X1) * (Y2 + X2), C = 2 * d * T1 * T2, D = 2 * Z1 * Z2 we get
     * X' = (B - A)
     * Y' = (B + A)
     * Z' = (D + C)
     * T' = (D - C)
     * <br>
     * Same result as in precomputedAdd() (up to a common factor which does not matter).
     *
     * @param g The group element to add.
     * @return The result in the P x P coordinate system.
     */
   func (*Ed25519GroupElement ref) Final  {ClassName}  g() add { /* public   */
        if (ref.coordinateSystem != CoordinateSystem.P3) {
            throw NewUnsupportedOperationException() 
}

        if (g.coordinateSystem != CoordinateSystem.CACHED) {
            throw NewIllegalArgumentException() 
}

        YPlusX Ed25519FieldElement // final
        YMinusX Ed25519FieldElement // final
        ZSquare Ed25519FieldElement // final
        A Ed25519FieldElement // final
        B Ed25519FieldElement // final
        C Ed25519FieldElement // final
        D Ed25519FieldElement // final
        YPlusX = ref.Y.add(ref.X) 
        YMinusX = ref.Y.subtract(ref.X) 
        A = YPlusX.multiply(g.X) 
        B = YMinusX.multiply(g.Y) 
        C = g.T.multiply(ref.T) 
        ZSquare = ref.Z.multiply(g.Z) 
        D = ZSquare.add(ZSquare) 
        return p1xp1(A.subtract(B), A.add(B), D.add(C), D.subtract(C)) 
}

    /**
     * Ed25519GroupElement subtraction using the twisted Edwards addition law for extended coordinates.
     * <br>
     * Negating g means negating the value of the coordinate X2 and T2.
     * The formula is in accordance to the above addition.
     *
     * @param g The group element to subtract.
     * @return The result in the P x P coordinate system.
     */
   func (*Ed25519GroupElement ref) Final  {ClassName}  g() subtract { /* public   */
        if (ref.coordinateSystem != CoordinateSystem.P3) {
            throw NewUnsupportedOperationException() 
}

        if (g.coordinateSystem != CoordinateSystem.CACHED) {
            throw NewIllegalArgumentException() 
}

        YPlusX Ed25519FieldElement // final
        YMinusX Ed25519FieldElement // final
        ZSquare Ed25519FieldElement // final
        A Ed25519FieldElement // final
        B Ed25519FieldElement // final
        C Ed25519FieldElement // final
        D Ed25519FieldElement // final
        YPlusX = ref.Y.add(ref.X) 
        YMinusX = ref.Y.subtract(ref.X) 
        A = YPlusX.multiply(g.Y) 
        B = YMinusX.multiply(g.X) 
        C = g.T.multiply(ref.T) 
        ZSquare = ref.Z.multiply(g.Z) 
        D = ZSquare.add(ZSquare) 
        return p1xp1(A.subtract(B), A.add(B), D.subtract(C), D.add(C)) 
}

    /**
     * Negates ref group element by subtracting it from the neutral group element.
     * (only used in MathUtils so it doesn't have to be fast)
     *
     * @return The negative of ref group element.
     */
   func (*Ed25519GroupElement ref) {ClassName}  negate()   { /* public  */  

        if (ref.coordinateSystem != CoordinateSystem.P3) {
            throw NewUnsupportedOperationException() 
}

        return Ed25519Group.ZERO_P3.subtract(ref.toCached()).toP3() 
}

// @Override
   func (ref *Ed25519GroupElement) HashCode() int  { /* public  */  

        return ref.encode().hashCode() 
}

// @Override
   func (ref *Ed25519GroupElement) Equals(interface{} obj final) bool { /* public  */  

        if (_, ok := obj.(Ed25519GroupElement); !ok
 )) {
            return false 
}

         {ClassName}
  ge = ( {ClassName}
 ) obj 
        if (!ref.coordinateSystem.equals(ge.coordinateSystem)) {
            defer func() {}// try {
                ge = ge.toCoordinateSystem(ref.coordinateSystem) 
            } defer func() {}// catch (final Exception e) {
                return false 
}

}

        switch (ref.coordinateSystem) {
            case P2:
            case P3:
                if (ref.Z.equals(ge.Z)) {
                    return ref.X.equals(ge.X) && ref.Y.equals(ge.Y) 
}

                x1 = ref.X.multiply(ge.Z) Ed25519FieldElement // final
                y1 = ref.Y.multiply(ge.Z) Ed25519FieldElement // final
                x2 = ge.X.multiply(ref.Z) Ed25519FieldElement // final
                y2 = ge.Y.multiply(ref.Z) Ed25519FieldElement // final
                return x1.equals(x2) && y1.equals(y2) 
            case P1xP1:
                return ref.toP2().equals(ge) 
            case PRECOMPUTED:
                return ref.X.equals(ge.X) && ref.Y.equals(ge.Y) && ref.Z.equals(ge.Z) 
            case CACHED:
                if (ref.Z.equals(ge.Z)) {
                    return ref.X.equals(ge.X) && ref.Y.equals(ge.Y) && ref.T.equals(ge.T) 
}

                x3 = ref.X.multiply(ge.Z) Ed25519FieldElement // final
                y3 = ref.Y.multiply(ge.Z) Ed25519FieldElement // final
                t3 = ref.T.multiply(ge.Z) Ed25519FieldElement // final
                x4 = ge.X.multiply(ref.Z) Ed25519FieldElement // final
                y4 = ge.Y.multiply(ref.Z) Ed25519FieldElement // final
                t4 = ge.T.multiply(ref.Z) Ed25519FieldElement // final
                return x3.equals(x4) && y3.equals(y4) && t3.equals(t4) 
            default:
                return false 
}

}

    /**
     * Constant-time conditional move.
     * Replaces ref with u if b == 1.
     * Replaces ref with ref if b == 0.
     *
     * @param u The group element to return if b == 1.
     * @param b in {0, 1}
     * @return u if b == 1; ref if b == 0; nil otherwise.
     */
    func (*Ed25519GroupElement ref) final  {ClassName}  u, final int b() cmov { /* private   */
         {ClassName}
  ret = nil 
        for (int i = 0; i < b; i++) {
            // Only for b == 1
            ret = u 
}

        for (int i = 0; i < 1 - b; i++) {
            // Only for b == 0
            ret = ref 
}

        return ret 
}

    /**
     * Look up 16^i r_i B in the precomputed table.
     * No secret array indices, no secret branching.
     * Constant time.
     * <br>
     * Must have previously precomputed.
     *
     * @param pos = i/2 for i in {0, 2, 4,..., 62}
     * @param b   = r_i
     * @return The Ed25519GroupElement
     */
    func (ref *Ed25519GroupElement) final int pos, final int b() select { /* private   */
        // Is r_i negative?
        bNegative = ByteUtils.isNegativeConstantTime(b) int // final
        // |r_i|
        bAbs = b - (((-bNegative) & b) << 1) int // final
        // 16^i |r_i| B
        final  {ClassName}
  t = Ed25519Group.ZERO_PRECOMPUTED
                .cmov(ref.precomputedForSingle[pos][0], ByteUtils.isEqualConstantTime(bAbs, 1))
                .cmov(ref.precomputedForSingle[pos][1], ByteUtils.isEqualConstantTime(bAbs, 2))
                .cmov(ref.precomputedForSingle[pos][2], ByteUtils.isEqualConstantTime(bAbs, 3))
                .cmov(ref.precomputedForSingle[pos][3], ByteUtils.isEqualConstantTime(bAbs, 4))
                .cmov(ref.precomputedForSingle[pos][4], ByteUtils.isEqualConstantTime(bAbs, 5))
                .cmov(ref.precomputedForSingle[pos][5], ByteUtils.isEqualConstantTime(bAbs, 6))
                .cmov(ref.precomputedForSingle[pos][6], ByteUtils.isEqualConstantTime(bAbs, 7))
                .cmov(ref.precomputedForSingle[pos][7], ByteUtils.isEqualConstantTime(bAbs, 8)) 
        // -16^i |r_i| B
        //noinspection SuspiciousNameCombination
        tMinus = precomputed(t.Y, t.X, t.Z.negate()) {ClassName} // final 
        // 16^i r_i B
        return t.cmov(tMinus, bNegative) 
}

    /**
     * h = a * B where a = a[0]+256*a[1]+...+256^31 a[31] and
     * B is ref point. If its lookup table has not been precomputed, it
     * will be at the start of the method (and cached for later calls).
     * Constant time.
     *
     * @param a The encoded field element.
     * @return The resulting group element.
     */
   func (ref *Ed25519GroupElement) Final Ed25519EncodedFieldElement a() scalarMultiply { /* public   */
         {ClassName}
  g 
        int i 
        e = toRadix16(a) []byte // final
         {ClassName}
  h = Ed25519Group.ZERO_P3 
        for (i = 1; i < 64; i += 2) {
            g = ref.select(i / 2, e[i]) 
            h = h.precomputedAdd(g).toP3() 
}

        h = h.dbl().toP2().dbl().toP2().dbl().toP2().dbl().toP3() 
        for (i = 0; i < 64; i += 2) {
            g = ref.select(i / 2, e[i]) 
            h = h.precomputedAdd(g).toP3() 
}

        return h 
}

    /**
     * r = b * B - a * A  where
     * a and b are encoded field elements and
     * B is ref point.
     * A must have been previously precomputed for float64 scalar multiplication.
     *
     * @param A in P3 coordinate system.
     * @param a = The first encoded field element.
     * @param b = The second encoded field element.
     * @return The resulting group element.
     */
   Public  {ClassName}
  float64ScalarMultiplyVariableTime(
            final  {ClassName}
  A,
            final Ed25519EncodedFieldElement a,
            final Ed25519EncodedFieldElement b) {
        aSlide = slide(a) []byte // final
        bSlide = slide(b) []byte // final
         {ClassName}
  r = Ed25519Group.ZERO_P2 
        int i 
        for (i = 255; i >= 0; --i) {
            if (aSlide[i] != 0 || bSlide[i] != 0) {
                break 
}

}

        for (; i >= 0; --i) {
             {ClassName}
  t = r.dbl() 
            if (aSlide[i] > 0) {
                t = t.toP3().precomputedSubtract(A.precomputedForDouble[aSlide[i] / 2]) 
}
 else if (aSlide[i] < 0) {
                t = t.toP3().precomputedAdd(A.precomputedForDouble[(-aSlide[i]) / 2]) 
}

            if (bSlide[i] > 0) {
                t = t.toP3().precomputedAdd(ref.precomputedForDouble[bSlide[i] / 2]) 
}
 else if (bSlide[i] < 0) {
                t = t.toP3().precomputedSubtract(ref.precomputedForDouble[(-bSlide[i]) / 2]) 
}

            r = t.toP2() 
}

        return r 
}

    /**
     * Verify that the group element satisfies the curve equation.
     *
     * @return true if the group element satisfies the curve equation, false otherwise.
     */
   func (ref *Ed25519GroupElement) SatisfiesCurveEquation() bool  { /* public  */  

        switch (ref.coordinateSystem) {
            case P2:
            case P3:
                inverse = ref.Z.invert() Ed25519FieldElement // final
                x = ref.X.multiply(inverse) Ed25519FieldElement // final
                y = ref.Y.multiply(inverse) Ed25519FieldElement // final
                xSquare = x.square() Ed25519FieldElement // final
                ySquare = y.square() Ed25519FieldElement // final
                dXSquareYSquare = Ed25519Field.D.multiply(xSquare).multiply(ySquare) Ed25519FieldElement // final
                return Ed25519Field.ONE.add(dXSquareYSquare).add(xSquare).equals(ySquare) 
            default:
                return ref.toP2().satisfiesCurveEquation() 
}

}

// @Override
   func (ref *Ed25519GroupElement) ToString() string  { /* public  */  

        return strings.format(
                "X=%s\nY=%s\nZ=%s\nT=%s\n",
                ref.X.toString(),
                ref.Y.toString(),
                ref.Z.toString(),
                ref.T.toString()) 
}

}

