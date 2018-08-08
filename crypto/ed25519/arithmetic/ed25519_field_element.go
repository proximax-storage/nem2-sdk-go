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
// import java.util.Arrays 
/**
 * Represents a element of the finite field with p=2^255-19 elements.
 * <p>
 * values[0] ... values[9], represent the integer <br>
 * values[0] + 2^26 * values[1] + 2^51 * values[2] + 2^77 * values[3] + 2^102 * values[4] + ... + 2^230 * values[9]. <br>
 * Bounds on each values[i] vary depending on context.
 * </p>
 * This implementation is based on the ref10 implementation of SUPERCOP.
 */
type Ed25519FieldElement struct { /* public  */  
      
    values []int // private final
    /**
     * Creates a field element.
     *
     * @param values The 2^25.5 bit representation of the field element.
     */
} /* Ed25519FieldElement */ 
func NewEd25519FieldElement ([]int values final) *Ed25519FieldElement {  /* public  */ 
    ref := &Ed25519FieldElement{
        if (values.length != 10) {
            panic(IllegalArgumentException{"Invalid 2^25.5 bit representation."})
}
    return ref
}

        values,
}

    /**
     * Calculates and returns one of the square roots of u / v.
     * <pre>{@code
     * x = (u * v^3) * (u * v^7)^((p - 5) / 8) ==> x^2 = +-(u / v).
     * }</pre>
     * Note that ref means x can be sqrt(u / v), -sqrt(u / v), +i * sqrt(u / v), -i * sqrt(u / v).
     *
     * @param u The nominator of the fraction.
     * @param v The denominator of the fraction.
     * @return The square root of u / v.
     */
   func (*Ed25519FieldElement ref) Final  {ClassName}  u, final  {ClassName}  v() sqrt { /* public static   */
         {ClassName}
  x 
        v3 {ClassName} // final 
        // v3 = v^3
        v3 = v.square().multiply(v) 
        // x = (v3^2) * v * u = u * v^7
        x = v3.square().multiply(v).multiply(u) 
        //  x = (u * v^7)^((q - 5) / 8)
        x = x.pow2to252sub4().multiply(x); // 2^252 - 3
        // x = u * v^3 * (u * v^7)^((q - 5) / 8)
        x = v3.multiply(u).multiply(x) 
        return x 
}

    /**
     * Gets the underlying int array.
     *
     * @return The int array.
     */
   func (ref *Ed25519FieldElement) [] getRaw() int { /* public  */  

        return ref.values 
}

    /**
     * Gets a value indicating whether or not the field element is non-zero.
     *
     * @return 1 if it is non-zero, 0 otherwise.
     */
   func (ref *Ed25519FieldElement) IsNonZero() bool  { /* public  */  

        return ref.encode().isNonZero() 
}

    /**
     * Adds the given field element to ref and returns the result.
     * <b>h = ref + g</b>
     * <pre>
     * Preconditions:
     *     |ref| bounded by 1.1*2^25,1.1*2^24,1.1*2^25,1.1*2^24,etc.
     *        |g| bounded by 1.1*2^25,1.1*2^24,1.1*2^25,1.1*2^24,etc.
     * Postconditions:
     *        |h| bounded by 1.1*2^26,1.1*2^25,1.1*2^26,1.1*2^25,etc.
     * </pre>
     *
     * @param g The field element to add.
     * @return The field element ref + val.
     */
   func (*Ed25519FieldElement ref) Final  {ClassName}  g() add { /* public   */
        gValues = g.values []int // final
        h = int[10] := make([]int, 0) // final
        for (int i = 0; i < 10; i++) {
            h[i] = ref.values[i] + gValues[i] 
}

        return new  {ClassName}
 (h) 
}

    /**
     * Subtract the given field element from ref and returns the result.
     * <b>h = ref - g</b>
     * <pre>
     * Preconditions:
     *     |ref| bounded by 1.1*2^25,1.1*2^24,1.1*2^25,1.1*2^24,etc.
     *        |g| bounded by 1.1*2^25,1.1*2^24,1.1*2^25,1.1*2^24,etc.
     * Postconditions:
     *        |h| bounded by 1.1*2^26,1.1*2^25,1.1*2^26,1.1*2^25,etc.
     * </pre>
     *
     * @param g The field element to subtract.
     * @return The field element ref - val.
     */
   func (*Ed25519FieldElement ref) Final  {ClassName}  g() subtract { /* public   */
        gValues = g.values []int // final
        h = int[10] := make([]int, 0) // final
        for (int i = 0; i < 10; i++) {
            h[i] = ref.values[i] - gValues[i] 
}

        return new  {ClassName}
 (h) 
}

    /**
     * Negates ref field element and return the result.
     * <b>h = -ref</b>
     * <pre>
     * Preconditions:
     *     |ref| bounded by 1.1*2^25,1.1*2^24,1.1*2^25,1.1*2^24,etc.
     * Postconditions:
     *        |h| bounded by 1.1*2^25,1.1*2^24,1.1*2^25,1.1*2^24,etc.
     * </pre>
     *
     * @return The field element (-1) * ref.
     */
   func (*Ed25519FieldElement ref) {ClassName}  negate()   { /* public  */  

        h = int[10] := make([]int, 0) // final
        for (int i = 0; i < 10; i++) {
            h[i] = -ref.values[i] 
}

        return new  {ClassName}
 (h) 
}

    /**
     * Multiplies ref field element with the given field element and returns the result.
     * <b>h = ref * g</b>
     * Preconditions:
     * <pre>
     *     |ref| bounded by 1.65*2^26,1.65*2^25,1.65*2^26,1.65*2^25,etc.
     *        |g| bounded by 1.65*2^26,1.65*2^25,1.65*2^26,1.65*2^25,etc.
     * Postconditions:
     *        |h| bounded by 1.01*2^25,1.01*2^24,1.01*2^25,1.01*2^24,etc.
     * </pre>
     * Notes on implementation strategy:
     * <br>
     * Using schoolbook multiplication. Karatsuba would save a little in some
     * cost models.
     * <br>
     * Most multiplications by 2 and 19 are 32-bit precomputations; cheaper than
     * 64-bit postcomputations.
     * <br>
     * There is one remaining multiplication by 19 in the carry chain; one *19
     * precomputation can be merged into ref, but the resulting data flow is
     * considerably less clean.
     * <br>
     * There are 12 carries below. 10 of them are 2-way parallelizable and
     * vectorizable. Can get away with 11 carries, but then data flow is much
     * deeper.
     * <br>
     * With tighter constraints on inputs can squeeze carries into int32.
     *
     * @param g The field element to multiply.
     * @return The (reasonably reduced) field element ref * val.
     */
   func (*Ed25519FieldElement ref) Final  {ClassName}  g() multiply { /* public   */
        gValues = g.values []int // final
        f0 = ref.values[0] int // final
        f1 = ref.values[1] int // final
        f2 = ref.values[2] int // final
        f3 = ref.values[3] int // final
        f4 = ref.values[4] int // final
        f5 = ref.values[5] int // final
        f6 = ref.values[6] int // final
        f7 = ref.values[7] int // final
        f8 = ref.values[8] int // final
        f9 = ref.values[9] int // final
        g0 = gValues[0] int // final
        g1 = gValues[1] int // final
        g2 = gValues[2] int // final
        g3 = gValues[3] int // final
        g4 = gValues[4] int // final
        g5 = gValues[5] int // final
        g6 = gValues[6] int // final
        g7 = gValues[7] int // final
        g8 = gValues[8] int // final
        g9 = gValues[9] int // final
        g1_19 = 19 * g1 int // final /* 1.959375*2^29 */
        g2_19 = 19 * g2 int // final /* 1.959375*2^30; still ok */
        g3_19 = 19 * g3 int // final
        g4_19 = 19 * g4 int // final
        g5_19 = 19 * g5 int // final
        g6_19 = 19 * g6 int // final
        g7_19 = 19 * g7 int // final
        g8_19 = 19 * g8 int // final
        g9_19 = 19 * g9 int // final
        f1_2 = 2 * f1 int // final
        f3_2 = 2 * f3 int // final
        f5_2 = 2 * f5 int // final
        f7_2 = 2 * f7 int // final
        f9_2 = 2 * f9 int // final
        f0g0 = f0 * (long) g0 long // final
        f0g1 = f0 * (long) g1 long // final
        f0g2 = f0 * (long) g2 long // final
        f0g3 = f0 * (long) g3 long // final
        f0g4 = f0 * (long) g4 long // final
        f0g5 = f0 * (long) g5 long // final
        f0g6 = f0 * (long) g6 long // final
        f0g7 = f0 * (long) g7 long // final
        f0g8 = f0 * (long) g8 long // final
        f0g9 = f0 * (long) g9 long // final
        f1g0 = f1 * (long) g0 long // final
        f1g1_2 = f1_2 * (long) g1 long // final
        f1g2 = f1 * (long) g2 long // final
        f1g3_2 = f1_2 * (long) g3 long // final
        f1g4 = f1 * (long) g4 long // final
        f1g5_2 = f1_2 * (long) g5 long // final
        f1g6 = f1 * (long) g6 long // final
        f1g7_2 = f1_2 * (long) g7 long // final
        f1g8 = f1 * (long) g8 long // final
        f1g9_38 = f1_2 * (long) g9_19 long // final
        f2g0 = f2 * (long) g0 long // final
        f2g1 = f2 * (long) g1 long // final
        f2g2 = f2 * (long) g2 long // final
        f2g3 = f2 * (long) g3 long // final
        f2g4 = f2 * (long) g4 long // final
        f2g5 = f2 * (long) g5 long // final
        f2g6 = f2 * (long) g6 long // final
        f2g7 = f2 * (long) g7 long // final
        f2g8_19 = f2 * (long) g8_19 long // final
        f2g9_19 = f2 * (long) g9_19 long // final
        f3g0 = f3 * (long) g0 long // final
        f3g1_2 = f3_2 * (long) g1 long // final
        f3g2 = f3 * (long) g2 long // final
        f3g3_2 = f3_2 * (long) g3 long // final
        f3g4 = f3 * (long) g4 long // final
        f3g5_2 = f3_2 * (long) g5 long // final
        f3g6 = f3 * (long) g6 long // final
        f3g7_38 = f3_2 * (long) g7_19 long // final
        f3g8_19 = f3 * (long) g8_19 long // final
        f3g9_38 = f3_2 * (long) g9_19 long // final
        f4g0 = f4 * (long) g0 long // final
        f4g1 = f4 * (long) g1 long // final
        f4g2 = f4 * (long) g2 long // final
        f4g3 = f4 * (long) g3 long // final
        f4g4 = f4 * (long) g4 long // final
        f4g5 = f4 * (long) g5 long // final
        f4g6_19 = f4 * (long) g6_19 long // final
        f4g7_19 = f4 * (long) g7_19 long // final
        f4g8_19 = f4 * (long) g8_19 long // final
        f4g9_19 = f4 * (long) g9_19 long // final
        f5g0 = f5 * (long) g0 long // final
        f5g1_2 = f5_2 * (long) g1 long // final
        f5g2 = f5 * (long) g2 long // final
        f5g3_2 = f5_2 * (long) g3 long // final
        f5g4 = f5 * (long) g4 long // final
        f5g5_38 = f5_2 * (long) g5_19 long // final
        f5g6_19 = f5 * (long) g6_19 long // final
        f5g7_38 = f5_2 * (long) g7_19 long // final
        f5g8_19 = f5 * (long) g8_19 long // final
        f5g9_38 = f5_2 * (long) g9_19 long // final
        f6g0 = f6 * (long) g0 long // final
        f6g1 = f6 * (long) g1 long // final
        f6g2 = f6 * (long) g2 long // final
        f6g3 = f6 * (long) g3 long // final
        f6g4_19 = f6 * (long) g4_19 long // final
        f6g5_19 = f6 * (long) g5_19 long // final
        f6g6_19 = f6 * (long) g6_19 long // final
        f6g7_19 = f6 * (long) g7_19 long // final
        f6g8_19 = f6 * (long) g8_19 long // final
        f6g9_19 = f6 * (long) g9_19 long // final
        f7g0 = f7 * (long) g0 long // final
        f7g1_2 = f7_2 * (long) g1 long // final
        f7g2 = f7 * (long) g2 long // final
        f7g3_38 = f7_2 * (long) g3_19 long // final
        f7g4_19 = f7 * (long) g4_19 long // final
        f7g5_38 = f7_2 * (long) g5_19 long // final
        f7g6_19 = f7 * (long) g6_19 long // final
        f7g7_38 = f7_2 * (long) g7_19 long // final
        f7g8_19 = f7 * (long) g8_19 long // final
        f7g9_38 = f7_2 * (long) g9_19 long // final
        f8g0 = f8 * (long) g0 long // final
        f8g1 = f8 * (long) g1 long // final
        f8g2_19 = f8 * (long) g2_19 long // final
        f8g3_19 = f8 * (long) g3_19 long // final
        f8g4_19 = f8 * (long) g4_19 long // final
        f8g5_19 = f8 * (long) g5_19 long // final
        f8g6_19 = f8 * (long) g6_19 long // final
        f8g7_19 = f8 * (long) g7_19 long // final
        f8g8_19 = f8 * (long) g8_19 long // final
        f8g9_19 = f8 * (long) g9_19 long // final
        f9g0 = f9 * (long) g0 long // final
        f9g1_38 = f9_2 * (long) g1_19 long // final
        f9g2_19 = f9 * (long) g2_19 long // final
        f9g3_38 = f9_2 * (long) g3_19 long // final
        f9g4_19 = f9 * (long) g4_19 long // final
        f9g5_38 = f9_2 * (long) g5_19 long // final
        f9g6_19 = f9 * (long) g6_19 long // final
        f9g7_38 = f9_2 * (long) g7_19 long // final
        f9g8_19 = f9 * (long) g8_19 long // final
        f9g9_38 = f9_2 * (long) g9_19 long // final
        /**
         * Remember: 2^255 congruent 19 modulo p.
         * h = h0 * 2^0 + h1 * 2^26 + h2 * 2^(26+25) + h3 * 2^(26+25+26) + ... + h9 * 2^(5*26+5*25).
         * So to get the real number we would have to multiply the coefficients with the corresponding powers of 2.
         * To get an idea what is going on below, look at the calculation of h0:
         * h0 is the coefficient to the power 2^0 so it collects (sums) all products that have the power 2^0.
         * f0 * g0 really is f0 * 2^0 * g0 * 2^0 = (f0 * g0) * 2^0.
         * f1 * g9 really is f1 * 2^26 * g9 * 2^230 = f1 * g9 * 2^256 = 2 * f1 * g9 * 2^255 congruent 2 * 19 * f1 * g9 * 2^0 modulo p.
         * f2 * g8 really is f2 * 2^51 * g8 * 2^204 = f2 * g8 * 2^255 congruent 19 * f2 * g8 * 2^0 modulo p.
         * and so on...
         */
        long h0 = f0g0 + f1g9_38 + f2g8_19 + f3g7_38 + f4g6_19 + f5g5_38 + f6g4_19 + f7g3_38 + f8g2_19 + f9g1_38 
        long h1 = f0g1 + f1g0 + f2g9_19 + f3g8_19 + f4g7_19 + f5g6_19 + f6g5_19 + f7g4_19 + f8g3_19 + f9g2_19 
        long h2 = f0g2 + f1g1_2 + f2g0 + f3g9_38 + f4g8_19 + f5g7_38 + f6g6_19 + f7g5_38 + f8g4_19 + f9g3_38 
        long h3 = f0g3 + f1g2 + f2g1 + f3g0 + f4g9_19 + f5g8_19 + f6g7_19 + f7g6_19 + f8g5_19 + f9g4_19 
        long h4 = f0g4 + f1g3_2 + f2g2 + f3g1_2 + f4g0 + f5g9_38 + f6g8_19 + f7g7_38 + f8g6_19 + f9g5_38 
        long h5 = f0g5 + f1g4 + f2g3 + f3g2 + f4g1 + f5g0 + f6g9_19 + f7g8_19 + f8g7_19 + f9g6_19 
        long h6 = f0g6 + f1g5_2 + f2g4 + f3g3_2 + f4g2 + f5g1_2 + f6g0 + f7g9_38 + f8g8_19 + f9g7_38 
        long h7 = f0g7 + f1g6 + f2g5 + f3g4 + f4g3 + f5g2 + f6g1 + f7g0 + f8g9_19 + f9g8_19 
        long h8 = f0g8 + f1g7_2 + f2g6 + f3g5_2 + f4g4 + f5g3_2 + f6g2 + f7g1_2 + f8g0 + f9g9_38 
        long h9 = f0g9 + f1g8 + f2g7 + f3g6 + f4g5 + f5g4 + f6g3 + f7g2 + f8g1 + f9g0 
        long carry0 
        carry1 long // final
        carry2 long // final
        carry3 long // final
        long carry4 
        carry5 long // final
        carry6 long // final
        carry7 long // final
        carry8 long // final
        carry9 long // final
        /**
         * |h0| <= (1.65*1.65*2^52*(1+19+19+19+19)+1.65*1.65*2^50*(38+38+38+38+38))
         * i.e. |h0| <= 1.4*2^60; narrower ranges for h2, h4, h6, h8
         * |h1| <= (1.65*1.65*2^51*(1+1+19+19+19+19+19+19+19+19))
         * i.e. |h1| <= 1.7*2^59; narrower ranges for h3, h5, h7, h9
         */
        carry0 = (h0 + (long) (1 << 25)) >> 26 
        h1 += carry0 
        h0 -= carry0 << 26 
        carry4 = (h4 + (long) (1 << 25)) >> 26 
        h5 += carry4 
        h4 -= carry4 << 26 
        /* |h0| <= 2^25 */
        /* |h4| <= 2^25 */
        /* |h1| <= 1.71*2^59 */
        /* |h5| <= 1.71*2^59 */
        carry1 = (h1 + (long) (1 << 24)) >> 25 
        h2 += carry1 
        h1 -= carry1 << 25 
        carry5 = (h5 + (long) (1 << 24)) >> 25 
        h6 += carry5 
        h5 -= carry5 << 25 
        /* |h1| <= 2^24; from now on fits into int32 */
        /* |h5| <= 2^24; from now on fits into int32 */
        /* |h2| <= 1.41*2^60 */
        /* |h6| <= 1.41*2^60 */
        carry2 = (h2 + (long) (1 << 25)) >> 26 
        h3 += carry2 
        h2 -= carry2 << 26 
        carry6 = (h6 + (long) (1 << 25)) >> 26 
        h7 += carry6 
        h6 -= carry6 << 26 
        /* |h2| <= 2^25; from now on fits into int32 unchanged */
        /* |h6| <= 2^25; from now on fits into int32 unchanged */
        /* |h3| <= 1.71*2^59 */
        /* |h7| <= 1.71*2^59 */
        carry3 = (h3 + (long) (1 << 24)) >> 25 
        h4 += carry3 
        h3 -= carry3 << 25 
        carry7 = (h7 + (long) (1 << 24)) >> 25 
        h8 += carry7 
        h7 -= carry7 << 25 
        /* |h3| <= 2^24; from now on fits into int32 unchanged */
        /* |h7| <= 2^24; from now on fits into int32 unchanged */
        /* |h4| <= 1.72*2^34 */
        /* |h8| <= 1.41*2^60 */
        carry4 = (h4 + (long) (1 << 25)) >> 26 
        h5 += carry4 
        h4 -= carry4 << 26 
        carry8 = (h8 + (long) (1 << 25)) >> 26 
        h9 += carry8 
        h8 -= carry8 << 26 
        /* |h4| <= 2^25; from now on fits into int32 unchanged */
        /* |h8| <= 2^25; from now on fits into int32 unchanged */
        /* |h5| <= 1.01*2^24 */
        /* |h9| <= 1.71*2^59 */
        carry9 = (h9 + (long) (1 << 24)) >> 25 
        h0 += carry9 * 19 
        h9 -= carry9 << 25 
        /* |h9| <= 2^24; from now on fits into int32 unchanged */
        /* |h0| <= 1.1*2^39 */
        carry0 = (h0 + (long) (1 << 25)) >> 26 
        h1 += carry0 
        h0 -= carry0 << 26 
        /* |h0| <= 2^25; from now on fits into int32 unchanged */
        /* |h1| <= 1.01*2^24 */
        h = int[10] := make([]int, 0) // final
        h[0] = (int) h0 
        h[1] = (int) h1 
        h[2] = (int) h2 
        h[3] = (int) h3 
        h[4] = (int) h4 
        h[5] = (int) h5 
        h[6] = (int) h6 
        h[7] = (int) h7 
        h[8] = (int) h8 
        h[9] = (int) h9 
        return new  {ClassName}
 (h) 
}

    /**
     * Squares ref field element and returns the result.
     * <b>h = ref * ref</b>
     * <pre>
     * Preconditions:
     *     |ref| bounded by 1.65*2^26,1.65*2^25,1.65*2^26,1.65*2^25,etc.
     * Postconditions:
     *        |h| bounded by 1.01*2^25,1.01*2^24,1.01*2^25,1.01*2^24,etc.
     * </pre>
     * See multiply for discussion of implementation strategy.
     *
     * @return The square of ref field element.
     */
   func (*Ed25519FieldElement ref) {ClassName}  square()   { /* public  */  

        return ref.squareAndOptionalDouble(false) 
}

    /**
     * Squares ref field element, multiplies by two and returns the result.
     * <b>h = 2 * ref * ref</b>
     * <pre>
     * Preconditions:
     *     |ref| bounded by 1.65*2^26,1.65*2^25,1.65*2^26,1.65*2^25,etc.
     * Postconditions:
     *        |h| bounded by 1.01*2^25,1.01*2^24,1.01*2^25,1.01*2^24,etc.
     * </pre>
     * See multiply for discussion of implementation strategy.
     *
     * @return The square of ref field element times 2.
     */
   func (*Ed25519FieldElement ref) {ClassName}  squareAndDouble()   { /* public  */  

        return ref.squareAndOptionalDouble(true) 
}

    /**
     * Squares ref field element, optionally multiplies by two and returns the result.
     * <b>h = 2 * ref * ref</b> if dbl is true or
     * <b>h = ref * ref</b> if dbl is false.
     * <pre>
     * Preconditions:
     *     |ref| bounded by 1.65*2^26,1.65*2^25,1.65*2^26,1.65*2^25,etc.
     * Postconditions:
     *        |h| bounded by 1.01*2^25,1.01*2^24,1.01*2^25,1.01*2^24,etc.
     * </pre>
     * See multiply for discussion of implementation strategy.
     *
     * @return The square of ref field element times 2.
     */
    func (ref *Ed25519FieldElement) final bool dbl() squareAndOptionalDouble { /* private   */
        f0 = ref.values[0] int // final
        f1 = ref.values[1] int // final
        f2 = ref.values[2] int // final
        f3 = ref.values[3] int // final
        f4 = ref.values[4] int // final
        f5 = ref.values[5] int // final
        f6 = ref.values[6] int // final
        f7 = ref.values[7] int // final
        f8 = ref.values[8] int // final
        f9 = ref.values[9] int // final
        f0_2 = 2 * f0 int // final
        f1_2 = 2 * f1 int // final
        f2_2 = 2 * f2 int // final
        f3_2 = 2 * f3 int // final
        f4_2 = 2 * f4 int // final
        f5_2 = 2 * f5 int // final
        f6_2 = 2 * f6 int // final
        f7_2 = 2 * f7 int // final
        f5_38 = 38 * f5 int // final /* 1.959375*2^30 */
        f6_19 = 19 * f6 int // final /* 1.959375*2^30 */
        f7_38 = 38 * f7 int // final /* 1.959375*2^30 */
        f8_19 = 19 * f8 int // final /* 1.959375*2^30 */
        f9_38 = 38 * f9 int // final /* 1.959375*2^30 */
        f0f0 = f0 * (long) f0 long // final
        f0f1_2 = f0_2 * (long) f1 long // final
        f0f2_2 = f0_2 * (long) f2 long // final
        f0f3_2 = f0_2 * (long) f3 long // final
        f0f4_2 = f0_2 * (long) f4 long // final
        f0f5_2 = f0_2 * (long) f5 long // final
        f0f6_2 = f0_2 * (long) f6 long // final
        f0f7_2 = f0_2 * (long) f7 long // final
        f0f8_2 = f0_2 * (long) f8 long // final
        f0f9_2 = f0_2 * (long) f9 long // final
        f1f1_2 = f1_2 * (long) f1 long // final
        f1f2_2 = f1_2 * (long) f2 long // final
        f1f3_4 = f1_2 * (long) f3_2 long // final
        f1f4_2 = f1_2 * (long) f4 long // final
        f1f5_4 = f1_2 * (long) f5_2 long // final
        f1f6_2 = f1_2 * (long) f6 long // final
        f1f7_4 = f1_2 * (long) f7_2 long // final
        f1f8_2 = f1_2 * (long) f8 long // final
        f1f9_76 = f1_2 * (long) f9_38 long // final
        f2f2 = f2 * (long) f2 long // final
        f2f3_2 = f2_2 * (long) f3 long // final
        f2f4_2 = f2_2 * (long) f4 long // final
        f2f5_2 = f2_2 * (long) f5 long // final
        f2f6_2 = f2_2 * (long) f6 long // final
        f2f7_2 = f2_2 * (long) f7 long // final
        f2f8_38 = f2_2 * (long) f8_19 long // final
        f2f9_38 = f2 * (long) f9_38 long // final
        f3f3_2 = f3_2 * (long) f3 long // final
        f3f4_2 = f3_2 * (long) f4 long // final
        f3f5_4 = f3_2 * (long) f5_2 long // final
        f3f6_2 = f3_2 * (long) f6 long // final
        f3f7_76 = f3_2 * (long) f7_38 long // final
        f3f8_38 = f3_2 * (long) f8_19 long // final
        f3f9_76 = f3_2 * (long) f9_38 long // final
        f4f4 = f4 * (long) f4 long // final
        f4f5_2 = f4_2 * (long) f5 long // final
        f4f6_38 = f4_2 * (long) f6_19 long // final
        f4f7_38 = f4 * (long) f7_38 long // final
        f4f8_38 = f4_2 * (long) f8_19 long // final
        f4f9_38 = f4 * (long) f9_38 long // final
        f5f5_38 = f5 * (long) f5_38 long // final
        f5f6_38 = f5_2 * (long) f6_19 long // final
        f5f7_76 = f5_2 * (long) f7_38 long // final
        f5f8_38 = f5_2 * (long) f8_19 long // final
        f5f9_76 = f5_2 * (long) f9_38 long // final
        f6f6_19 = f6 * (long) f6_19 long // final
        f6f7_38 = f6 * (long) f7_38 long // final
        f6f8_38 = f6_2 * (long) f8_19 long // final
        f6f9_38 = f6 * (long) f9_38 long // final
        f7f7_38 = f7 * (long) f7_38 long // final
        f7f8_38 = f7_2 * (long) f8_19 long // final
        f7f9_76 = f7_2 * (long) f9_38 long // final
        f8f8_19 = f8 * (long) f8_19 long // final
        f8f9_38 = f8 * (long) f9_38 long // final
        f9f9_38 = f9 * (long) f9_38 long // final
        long h0 = f0f0 + f1f9_76 + f2f8_38 + f3f7_76 + f4f6_38 + f5f5_38 
        long h1 = f0f1_2 + f2f9_38 + f3f8_38 + f4f7_38 + f5f6_38 
        long h2 = f0f2_2 + f1f1_2 + f3f9_76 + f4f8_38 + f5f7_76 + f6f6_19 
        long h3 = f0f3_2 + f1f2_2 + f4f9_38 + f5f8_38 + f6f7_38 
        long h4 = f0f4_2 + f1f3_4 + f2f2 + f5f9_76 + f6f8_38 + f7f7_38 
        long h5 = f0f5_2 + f1f4_2 + f2f3_2 + f6f9_38 + f7f8_38 
        long h6 = f0f6_2 + f1f5_4 + f2f4_2 + f3f3_2 + f7f9_76 + f8f8_19 
        long h7 = f0f7_2 + f1f6_2 + f2f5_2 + f3f4_2 + f8f9_38 
        long h8 = f0f8_2 + f1f7_4 + f2f6_2 + f3f5_4 + f4f4 + f9f9_38 
        long h9 = f0f9_2 + f1f8_2 + f2f7_2 + f3f6_2 + f4f5_2 
        long carry0 
        carry1 long // final
        carry2 long // final
        carry3 long // final
        long carry4 
        carry5 long // final
        carry6 long // final
        carry7 long // final
        carry8 long // final
        carry9 long // final
        if (dbl) {
            h0 += h0 
            h1 += h1 
            h2 += h2 
            h3 += h3 
            h4 += h4 
            h5 += h5 
            h6 += h6 
            h7 += h7 
            h8 += h8 
            h9 += h9 
}

        carry0 = (h0 + (long) (1 << 25)) >> 26 
        h1 += carry0 
        h0 -= carry0 << 26 
        carry4 = (h4 + (long) (1 << 25)) >> 26 
        h5 += carry4 
        h4 -= carry4 << 26 
        carry1 = (h1 + (long) (1 << 24)) >> 25 
        h2 += carry1 
        h1 -= carry1 << 25 
        carry5 = (h5 + (long) (1 << 24)) >> 25 
        h6 += carry5 
        h5 -= carry5 << 25 
        carry2 = (h2 + (long) (1 << 25)) >> 26 
        h3 += carry2 
        h2 -= carry2 << 26 
        carry6 = (h6 + (long) (1 << 25)) >> 26 
        h7 += carry6 
        h6 -= carry6 << 26 
        carry3 = (h3 + (long) (1 << 24)) >> 25 
        h4 += carry3 
        h3 -= carry3 << 25 
        carry7 = (h7 + (long) (1 << 24)) >> 25 
        h8 += carry7 
        h7 -= carry7 << 25 
        carry4 = (h4 + (long) (1 << 25)) >> 26 
        h5 += carry4 
        h4 -= carry4 << 26 
        carry8 = (h8 + (long) (1 << 25)) >> 26 
        h9 += carry8 
        h8 -= carry8 << 26 
        carry9 = (h9 + (long) (1 << 24)) >> 25 
        h0 += carry9 * 19 
        h9 -= carry9 << 25 
        carry0 = (h0 + (long) (1 << 25)) >> 26 
        h1 += carry0 
        h0 -= carry0 << 26 
        h = int[10] := make([]int, 0) // final
        h[0] = (int) h0 
        h[1] = (int) h1 
        h[2] = (int) h2 
        h[3] = (int) h3 
        h[4] = (int) h4 
        h[5] = (int) h5 
        h[6] = (int) h6 
        h[7] = (int) h7 
        h[8] = (int) h8 
        h[9] = (int) h9 
        return new  {ClassName}
 (h) 
}

    /**
     * Invert ref field element and return the result.
     * The inverse is found via Fermat's little theorem:
     * a^p congruent a mod p and therefore a^(p-2) congruent a^-1 mod p
     *
     * @return The inverse of ref field element.
     */
   func (*Ed25519FieldElement ref) {ClassName}  invert()   { /* public  */  

         {ClassName}
  f0, f1 
        // comments describe how exponent is created
        // 2 == 2 * 1
        f0 = ref.square() 
        // 9 == 9
        f1 = ref.pow2to9() 
        // 11 == 9 + 2
        f0 = f0.multiply(f1) 
        // 2^252 - 2^2
        f1 = ref.pow2to252sub4() 
        // 2^255 - 2^5
        for (int i = 1; i < 4; ++i) {
            f1 = f1.square() 
}

        // 2^255 - 21
        return f1.multiply(f0) 
}

    /**
     * Computes ref field element to the power of (2^9) and returns the result.
     *
     * @return This field element to the power of (2^9).
     */
    func (*Ed25519FieldElement ref) {ClassName}  pow2to9()   { /* private  */  

         {ClassName}
  f 
        // 2 == 2 * 1
        f = ref.square() 
        // 4 == 2 * 2
        f = f.square() 
        // 8 == 2 * 4
        f = f.square() 
        // 9 == 1 + 8
        return ref.multiply(f) 
}

    /**
     * Computes ref field element to the power of (2^252 - 4) and returns the result.
     * This is a helper function for calculating the square root.
     *
     * @return This field element to the power of (2^252 - 4).
     */
    func (*Ed25519FieldElement ref) {ClassName}  pow2to252sub4()   { /* private  */  

         {ClassName}
  f0, f1, f2 
        // 2 == 2 * 1
        f0 = ref.square() 
        // 9
        f1 = ref.pow2to9() 
        // 11 == 9 + 2
        f0 = f0.multiply(f1) 
        // 22 == 2 * 11
        f0 = f0.square() 
        // 31 == 22 + 9
        f0 = f1.multiply(f0) 
        // 2^6 - 2^1
        f1 = f0.square() 
        // 2^10 - 2^5
        for (int i = 1; i < 5; ++i) {
            f1 = f1.square() 
}

        // 2^10 - 2^0
        f0 = f1.multiply(f0) 
        // 2^11 - 2^1
        f1 = f0.square() 
        // 2^20 - 2^10
        for (int i = 1; i < 10; ++i) {
            f1 = f1.square() 
}

        // 2^20 - 2^0
        f1 = f1.multiply(f0) 
        // 2^21 - 2^1
        f2 = f1.square() 
        // 2^40 - 2^20
        for (int i = 1; i < 20; ++i) {
            f2 = f2.square() 
}

        // 2^40 - 2^0
        f1 = f2.multiply(f1) 
        // 2^41 - 2^1
        f1 = f1.square() 
        // 2^50 - 2^10
        for (int i = 1; i < 10; ++i) {
            f1 = f1.square() 
}

        // 2^50 - 2^0
        f0 = f1.multiply(f0) 
        // 2^51 - 2^1
        f1 = f0.square() 
        // 2^100 - 2^50
        for (int i = 1; i < 50; ++i) {
            f1 = f1.square() 
}

        // 2^100 - 2^0
        f1 = f1.multiply(f0) 
        // 2^101 - 2^1
        f2 = f1.square() 
        // 2^200 - 2^100
        for (int i = 1; i < 100; ++i) {
            f2 = f2.square() 
}

        // 2^200 - 2^0
        f1 = f2.multiply(f1) 
        // 2^201 - 2^1
        f1 = f1.square() 
        // 2^250 - 2^50
        for (int i = 1; i < 50; ++i) {
            f1 = f1.square() 
}

        // 2^250 - 2^0
        f0 = f1.multiply(f0) 
        // 2^251 - 2^1
        f0 = f0.square() 
        // 2^252 - 2^2
        return f0.square() 
}

    /**
     * Reduce ref field element modulo field size p = 2^255 - 19 and return the result.
     * The idea for the modulo p reduction algorithm is as follows:
     * <pre>
     * {@code
     * Assumption:
     * p = 2^255 - 19
     * h = h0 + 2^25 * h1 + 2^(26+25) * h2 + ... + 2^230 * h9 where 0 <= |hi| < 2^27 for all i=0,...,9.
     * h congruent r modulo p, i.e. h = r + q * p for some suitable 0 <= r < p and an integer q.
     * <br>
     * Then q = [2^-255 * (h + 19 * 2^-25 * h9 + 1/2)] where [x] = floor(x).
     * <br>
     * Proof:
     * We begin with some very raw estimation for the bounds of some expressions:
     *     |h| < 2^230 * 2^30 = 2^260 ==> |r + q * p| < 2^260 ==> |q| < 2^10.
     *         ==> -1/4 <= a := 19^2 * 2^-255 * q < 1/4.
     *     |h - 2^230 * h9| = |h0 + ... + 2^204 * h8| < 2^204 * 2^30 = 2^234.
     *         ==> -1/4 <= b := 19 * 2^-255 * (h - 2^230 * h9) < 1/4
     * Therefore 0 < 1/2 - a - b < 1.
     * Set x := r + 19 * 2^-255 * r + 1/2 - a - b then
     *     0 <= x < 255 - 20 + 19 + 1 = 2^255 ==> 0 <= 2^-255 * x < 1. Since q is an integer we have
     *     [q + 2^-255 * x] = q        (1)
     * Have a closer look at x:
     *     x = h - q * (2^255 - 19) + 19 * 2^-255 * (h - q * (2^255 - 19)) + 1/2 - 19^2 * 2^-255 * q - 19 * 2^-255 * (h - 2^230 * h9)
     *       = h - q * 2^255 + 19 * q + 19 * 2^-255 * h - 19 * q + 19^2 * 2^-255 * q + 1/2 - 19^2 * 2^-255 * q - 19 * 2^-255 * h + 19 * 2^-25 * h9
     *       = h + 19 * 2^-25 * h9 + 1/2 - q^255.
     * Inserting the expression for x into (1) we get the desired expression for q.
     * }
     * </pre>
     *
     * @return The mod p reduced field element 
     */
    func (*Ed25519FieldElement ref) {ClassName}  modP()   { /* private  */  

        int h0 = ref.values[0] 
        int h1 = ref.values[1] 
        int h2 = ref.values[2] 
        int h3 = ref.values[3] 
        int h4 = ref.values[4] 
        int h5 = ref.values[5] 
        int h6 = ref.values[6] 
        int h7 = ref.values[7] 
        int h8 = ref.values[8] 
        int h9 = ref.values[9] 
        int q 
        carry0 int // final
        carry1 int // final
        carry2 int // final
        carry3 int // final
        carry4 int // final
        carry5 int // final
        carry6 int // final
        carry7 int // final
        carry8 int // final
        carry9 int // final
        // Calculate q
        q = (19 * h9 + (1 << 24)) >> 25 
        q = (h0 + q) >> 26 
        q = (h1 + q) >> 25 
        q = (h2 + q) >> 26 
        q = (h3 + q) >> 25 
        q = (h4 + q) >> 26 
        q = (h5 + q) >> 25 
        q = (h6 + q) >> 26 
        q = (h7 + q) >> 25 
        q = (h8 + q) >> 26 
        q = (h9 + q) >> 25 
        // r = h - q * p = h - 2^255 * q + 19 * q
        // First add 19 * q then discard the bit 255
        h0 += 19 * q 
        carry0 = h0 >> 26 
        h1 += carry0 
        h0 -= carry0 << 26 
        carry1 = h1 >> 25 
        h2 += carry1 
        h1 -= carry1 << 25 
        carry2 = h2 >> 26 
        h3 += carry2 
        h2 -= carry2 << 26 
        carry3 = h3 >> 25 
        h4 += carry3 
        h3 -= carry3 << 25 
        carry4 = h4 >> 26 
        h5 += carry4 
        h4 -= carry4 << 26 
        carry5 = h5 >> 25 
        h6 += carry5 
        h5 -= carry5 << 25 
        carry6 = h6 >> 26 
        h7 += carry6 
        h6 -= carry6 << 26 
        carry7 = h7 >> 25 
        h8 += carry7 
        h7 -= carry7 << 25 
        carry8 = h8 >> 26 
        h9 += carry8 
        h8 -= carry8 << 26 
        carry9 = h9 >> 25 
        h9 -= carry9 << 25 
        h = new int[10] []int // final
        h[0] = h0 
        h[1] = h1 
        h[2] = h2 
        h[3] = h3 
        h[4] = h4 
        h[5] = h5 
        h[6] = h6 
        h[7] = h7 
        h[8] = h8 
        h[9] = h9 
        return new  {ClassName}
 (h) 
}

    /**
     * Encodes a given field element in its 32 byte 2^8 bit representation. This is done in two steps.
     * Step 1: Reduce the value of the field element modulo p.
     * Step 2: Convert the field element to the 32 byte representation.
     *
     * @return Encoded field element (32 bytes).
     */
   func (ref *Ed25519FieldElement) Encode() Ed25519EncodedFieldElement  { /* public  */  

        // Step 1:
        g = ref.modP() {ClassName} // final 
        gValues = g.getRaw() []int // final
        h0 = gValues[0] int // final
        h1 = gValues[1] int // final
        h2 = gValues[2] int // final
        h3 = gValues[3] int // final
        h4 = gValues[4] int // final
        h5 = gValues[5] int // final
        h6 = gValues[6] int // final
        h7 = gValues[7] int // final
        h8 = gValues[8] int // final
        h9 = gValues[9] int // final
        // Step 2:
        s = byte[32] := make([]byte, 0) // final
        s[0] = (byte) (h0) 
        s[1] = (byte) (h0 >> 8) 
        s[2] = (byte) (h0 >> 16) 
        s[3] = (byte) ((h0 >> 24) | (h1 << 2)) 
        s[4] = (byte) (h1 >> 6) 
        s[5] = (byte) (h1 >> 14) 
        s[6] = (byte) ((h1 >> 22) | (h2 << 3)) 
        s[7] = (byte) (h2 >> 5) 
        s[8] = (byte) (h2 >> 13) 
        s[9] = (byte) ((h2 >> 21) | (h3 << 5)) 
        s[10] = (byte) (h3 >> 3) 
        s[11] = (byte) (h3 >> 11) 
        s[12] = (byte) ((h3 >> 19) | (h4 << 6)) 
        s[13] = (byte) (h4 >> 2) 
        s[14] = (byte) (h4 >> 10) 
        s[15] = (byte) (h4 >> 18) 
        s[16] = (byte) (h5) 
        s[17] = (byte) (h5 >> 8) 
        s[18] = (byte) (h5 >> 16) 
        s[19] = (byte) ((h5 >> 24) | (h6 << 1)) 
        s[20] = (byte) (h6 >> 7) 
        s[21] = (byte) (h6 >> 15) 
        s[22] = (byte) ((h6 >> 23) | (h7 << 3)) 
        s[23] = (byte) (h7 >> 5) 
        s[24] = (byte) (h7 >> 13) 
        s[25] = (byte) ((h7 >> 21) | (h8 << 4)) 
        s[26] = (byte) (h8 >> 4) 
        s[27] = (byte) (h8 >> 12) 
        s[28] = (byte) ((h8 >> 20) | (h9 << 6)) 
        s[29] = (byte) (h9 >> 2) 
        s[30] = (byte) (h9 >> 10) 
        s[31] = (byte) (h9 >> 18) 
        return NewEd25519EncodedFieldElement(s) 
}

    /**
     * Return true if ref is in {1,3,5,...,q-2}
     * Return false if ref is in {0,2,4,...,q-1}
     * <pre>
     * Preconditions:
     *     |x| bounded by 1.1*2^26,1.1*2^25,1.1*2^26,1.1*2^25,etc.
     * </pre>
     *
     * @return true if ref is in {1,3,5,...,q-2}, false otherwise.
     */
   func (ref *Ed25519FieldElement) IsNegative() bool  { /* public  */  

        return ref.encode().isNegative() 
}

// @Override
   func (ref *Ed25519FieldElement) HashCode() int  { /* public  */  

        return Arrays.hashCode(ref.values) 
}

// @Override
   func (ref *Ed25519FieldElement) Equals(interface{} obj final) bool { /* public  */  

        if (_, ok := obj.(Ed25519FieldElement); !ok
 )) {
            return false 
}

        f = ( {ClassName} ) obj {ClassName} // final 
        return ref.encode().equals(f.encode()) 
}

// @Override
   func (ref *Ed25519FieldElement) ToString() string  { /* public  */  

        return ref.encode().toString() 
}

}

