// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package crypto

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/proximax-storage/nem2-sdk-go/utils"
	"math/big"
)

func PrepareForScalarMultiply(key *PrivateKey) *Ed25519EncodedFieldElement {

	hash, err := HashesSha3_512(key.Raw)
	if err != nil {
		fmt.Println(err)
	}
	a := hash[:32]
	a[31] &= 0x7F
	a[31] |= 0x40
	a[0] &= 0xF8
	return &Ed25519EncodedFieldElement{Ed25519Field_ZERO_SHORT(), a}
}

type Ed25519KeyAnalyzer struct {
}

func NewEd25519KeyAnalyzer() *Ed25519KeyAnalyzer {
	return &Ed25519KeyAnalyzer{}
}

const COMPRESSED_KEY_SIZE = 32

func (ref *Ed25519KeyAnalyzer) IsKeyCompressed(publicKey *PublicKey) bool {

	return COMPRESSED_KEY_SIZE == len(publicKey.Raw)
}

// Represents the underlying finite field for Ed25519.
//  The field has p = 2^255 - 19 elements.
type ed25519Field struct {
	P                              *big.Int
	ZERO, ONE, TWO, D, D_Times_TWO Ed25519FieldElement
	I                              Ed25519FieldElement
}

const Ed25519FieldI = "b0a00e4a271beec478e42fad0618432fa7d7fb3d99004d2b0bdfc14f8024832b"

func Ed25519Field_D() *Ed25519FieldElement {

	return &Ed25519FieldElement{[]intRaw{
		-10913610, 13857413, -15372611, 6949391, 114729, -8787816, -6275908, -3247719, -18696448, -12055116,
	}}
	// this calculatios return result as code before
	//s := big.NewInt(-121665)
	//d := utils.BigIntToByteArray(
	//	s.Mul(s, (&big.Int{}).ModInverse(big.NewInt(121666), Ed25519Field_P)).Mod(s, Ed25519Field_P),
	//	32)
	//el := *(&Ed25519EncodedFieldElement{Ed25519Field_ZERO_SHORT(), d}).Decode()
	//fmt.Printf("%+v", el.Raw)
	//
	//return el
}

func Ed25519Field_D_Times_TWO() *Ed25519FieldElement {
	// return getD().multiply(*(Ed25519Field_TWO())),
	return &Ed25519FieldElement{[]intRaw{-21827239, -5839606, -30745221, 13898782, 229458, 15978800, -12551817, -6495438, 29715968, 9444199}}
}

// this method replace on three methods for one base constant
//func getFieldElement(value intRaw) Ed25519FieldElement {
//
//	f := make([]intRaw, 10)
//	f[0] = value
//	return Ed25519FieldElement{f}
//}

func Ed25519Field_ZERO() *Ed25519FieldElement {
	return &Ed25519FieldElement{[]intRaw{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}
}
func Ed25519Field_ONE() *Ed25519FieldElement {
	return &Ed25519FieldElement{[]intRaw{1, 0, 0, 0, 0, 0, 0, 0, 0, 0}}
}
func Ed25519Field_ZERO_SHORT() []byte {
	return make([]byte, 32)
}
func Ed25519Field_ZERO_LONG() []byte {
	return make([]byte, 64)
}
func Ed25519Field_TWO() *Ed25519FieldElement {
	return &Ed25519FieldElement{[]intRaw{2, 0, 0, 0, 0, 0, 0, 0, 0, 0}}
}

// P: 2^255 - 19
func Ed25519Field_P() *big.Int {
	const Ed25519FieldP = "7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed"
	p := (&big.Int{}).Lsh(big.NewInt(1), 255)
	//SetBytes(utils.MustHexDecodeString(Ed25519FieldP))
	return p.Sub(p, big.NewInt(19))
}

var Ed25519Field = ed25519Field{
	Ed25519Field_P(),
	*(Ed25519Field_ZERO()),
	*(Ed25519Field_ONE()),
	*(Ed25519Field_TWO()),
	*(Ed25519Field_D()),
	*(Ed25519Field_D_Times_TWO()),
	// I ^ 2 = -1
	*((&Ed25519EncodedFieldElement{Ed25519Field_ZERO_SHORT(), utils.MustHexDecodeString(Ed25519FieldI)}).Decode()),
}

type intRaw = int64

const lenEd25519FieldElementRaw = 10

// Ed25519FieldElement Represents a element of the finite field with p=2^255-19 elements.
// Raw[0] ... Raw[9], represent the integer
// Raw[0] + 2^26 * Raw[1] + 2^51 * Raw[2] + 2^77 * Raw[3] + 2^102 * Raw[4] + ... + 2^230 * Raw[9].
// Bounds on each Raw[i] vary depending on context.
// This implementation is based on the ref10 implementation of SUPERCOP.
type Ed25519FieldElement struct {
	Raw []intRaw
}

// NewEd25519FieldElement Creates a field element.
// Raw The 2^25.5 bit representation of the field element.
func NewEd25519FieldElement(Raw []intRaw) (*Ed25519FieldElement, error) {
	if len(Raw) != lenEd25519FieldElementRaw {
		return nil, errors.New("Invalid 2^25.5 bit representation.")
	}
	return &Ed25519FieldElement{Raw}, nil
}

// Ed25519FieldElementSqrt Calculates and returns one of the square roots of u / v.
// * x = (u * v^3) * (u * v^7)^((p - 5) / 8) ==> x^2 = +-(u / v).
// * Note that ref means x can be sqrt(u / v), -sqrt(u / v), +i * sqrt(u / v), -i * sqrt(u / v).
// *
// * @param u The nominator of the fraction.
// * @param v The denominator of the fraction.
// * @return The square root of u / v.
func Ed25519FieldElementSqrt(u Ed25519FieldElement, v Ed25519FieldElement) Ed25519FieldElement {

	// v3 = v^3
	v3 := v.square().multiply(v)
	// x = (v3^2) * v * u = u * v^7
	x := v3.square().multiply(v).multiply(u)
	//  x = (u * v^7)^((q - 5) / 8)
	x = x.pow2to252sub4().multiply(x) // 2^252 - 3
	// x = u * v^3 * (u * v^7)^((q - 5) / 8)
	x = v3.multiply(u).multiply(x)
	return x
}

//IsNonZero gets a value indicating whether or not the field element is non-zero.
func (ref Ed25519FieldElement) IsNonZero() bool {

	return ref.Encode().IsNonZero()
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
func (ref Ed25519FieldElement) add(g Ed25519FieldElement) Ed25519FieldElement {
	h := make([]intRaw, 10)
	for i := range h {
		h[i] = ref.Raw[i] + g.Raw[i]
	}
	return Ed25519FieldElement{h}
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
 */func (ref Ed25519FieldElement) subtract(g Ed25519FieldElement) Ed25519FieldElement {
	h := make([]intRaw, 10)
	for i := range h {
		h[i] = ref.Raw[i] - g.Raw[i]
	}

	return Ed25519FieldElement{h}
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
 */func (ref Ed25519FieldElement) negate() Ed25519FieldElement {

	h := make([]intRaw, 10)
	for i := range h {
		h[i] = -ref.Raw[i]
	}

	return Ed25519FieldElement{h}
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
 */func (ref Ed25519FieldElement) multiply(g Ed25519FieldElement) Ed25519FieldElement {
	g1_19 := 19 * g.Raw[1] /* 1.959375*2^29 */
	g2_19 := 19 * g.Raw[2] /* 1.959375*2^30; still ok */
	g3_19 := 19 * g.Raw[3]
	g4_19 := 19 * g.Raw[4]
	g5_19 := 19 * g.Raw[5]
	g6_19 := 19 * g.Raw[6]
	g7_19 := 19 * g.Raw[7]
	g8_19 := 19 * g.Raw[8]
	g9_19 := 19 * g.Raw[9]
	f1_2 := 2 * ref.Raw[1]
	f3_2 := 2 * ref.Raw[3]
	f5_2 := 2 * ref.Raw[5]
	f7_2 := 2 * ref.Raw[7]
	f9_2 := 2 * ref.Raw[9]
	f0g0 := ref.Raw[0] * g.Raw[0]
	f0g1 := ref.Raw[0] * g.Raw[1]
	f0g2 := ref.Raw[0] * g.Raw[2]
	f0g3 := ref.Raw[0] * g.Raw[3]
	f0g4 := ref.Raw[0] * g.Raw[4]
	f0g5 := ref.Raw[0] * g.Raw[5]
	f0g6 := ref.Raw[0] * g.Raw[6]
	f0g7 := ref.Raw[0] * g.Raw[7]
	f0g8 := ref.Raw[0] * g.Raw[8]
	f0g9 := ref.Raw[0] * g.Raw[9]
	f1g0 := ref.Raw[1] * g.Raw[0]
	f1g1_2 := f1_2 * g.Raw[1]
	f1g2 := ref.Raw[1] * g.Raw[2]
	f1g3_2 := f1_2 * g.Raw[3]
	f1g4 := ref.Raw[1] * g.Raw[4]
	f1g5_2 := f1_2 * g.Raw[5]
	f1g6 := ref.Raw[1] * g.Raw[6]
	f1g7_2 := f1_2 * g.Raw[7]
	f1g8 := ref.Raw[1] * g.Raw[8]
	f1g9_38 := f1_2 * g9_19
	f2g0 := ref.Raw[2] * g.Raw[0]
	f2g1 := ref.Raw[2] * g.Raw[1]
	f2g2 := ref.Raw[2] * g.Raw[2]
	f2g3 := ref.Raw[2] * g.Raw[3]
	f2g4 := ref.Raw[2] * g.Raw[4]
	f2g5 := ref.Raw[2] * g.Raw[5]
	f2g6 := ref.Raw[2] * g.Raw[6]
	f2g7 := ref.Raw[2] * g.Raw[7]
	f2g8_19 := ref.Raw[2] * g8_19
	f2g9_19 := ref.Raw[2] * g9_19
	f3g0 := ref.Raw[3] * g.Raw[0]
	f3g1_2 := f3_2 * g.Raw[1]
	f3g2 := ref.Raw[3] * g.Raw[2]
	f3g3_2 := f3_2 * g.Raw[3]
	f3g4 := ref.Raw[3] * g.Raw[4]
	f3g5_2 := f3_2 * g.Raw[5]
	f3g6 := ref.Raw[3] * g.Raw[6]
	f3g7_38 := f3_2 * g7_19
	f3g8_19 := ref.Raw[3] * g8_19
	f3g9_38 := f3_2 * g9_19
	f4g0 := ref.Raw[4] * g.Raw[0]
	f4g1 := ref.Raw[4] * g.Raw[1]
	f4g2 := ref.Raw[4] * g.Raw[2]
	f4g3 := ref.Raw[4] * g.Raw[3]
	f4g4 := ref.Raw[4] * g.Raw[4]
	f4g5 := ref.Raw[4] * g.Raw[5]
	f4g6_19 := ref.Raw[4] * g6_19
	f4g7_19 := ref.Raw[4] * g7_19
	f4g8_19 := ref.Raw[4] * g8_19
	f4g9_19 := ref.Raw[4] * g9_19
	f5g0 := ref.Raw[5] * g.Raw[0]
	f5g1_2 := f5_2 * g.Raw[1]
	f5g2 := ref.Raw[5] * g.Raw[2]
	f5g3_2 := f5_2 * g.Raw[3]
	f5g4 := ref.Raw[5] * g.Raw[4]
	f5g5_38 := f5_2 * g5_19
	f5g6_19 := ref.Raw[5] * g6_19
	f5g7_38 := f5_2 * g7_19
	f5g8_19 := ref.Raw[5] * g8_19
	f5g9_38 := f5_2 * g9_19
	f6g0 := ref.Raw[6] * g.Raw[0]
	f6g1 := ref.Raw[6] * g.Raw[1]
	f6g2 := ref.Raw[6] * g.Raw[2]
	f6g3 := ref.Raw[6] * g.Raw[3]
	f6g4_19 := ref.Raw[6] * g4_19
	f6g5_19 := ref.Raw[6] * g5_19
	f6g6_19 := ref.Raw[6] * g6_19
	f6g7_19 := ref.Raw[6] * g7_19
	f6g8_19 := ref.Raw[6] * g8_19
	f6g9_19 := ref.Raw[6] * g9_19
	f7g0 := ref.Raw[7] * g.Raw[0]
	f7g1_2 := f7_2 * g.Raw[1]
	f7g2 := ref.Raw[7] * g.Raw[2]
	f7g3_38 := f7_2 * g3_19
	f7g4_19 := ref.Raw[7] * g4_19
	f7g5_38 := f7_2 * g5_19
	f7g6_19 := ref.Raw[7] * g6_19
	f7g7_38 := f7_2 * g7_19
	f7g8_19 := ref.Raw[7] * g8_19
	f7g9_38 := f7_2 * g9_19
	f8g0 := ref.Raw[8] * g.Raw[0]
	f8g1 := ref.Raw[8] * g.Raw[1]
	f8g2_19 := ref.Raw[8] * g2_19
	f8g3_19 := ref.Raw[8] * g3_19
	f8g4_19 := ref.Raw[8] * g4_19
	f8g5_19 := ref.Raw[8] * g5_19
	f8g6_19 := ref.Raw[8] * g6_19
	f8g7_19 := ref.Raw[8] * g7_19
	f8g8_19 := ref.Raw[8] * g8_19
	f8g9_19 := ref.Raw[8] * g9_19
	f9g0 := ref.Raw[9] * g.Raw[0]
	f9g1_38 := f9_2 * g1_19
	f9g2_19 := ref.Raw[9] * g2_19
	f9g3_38 := f9_2 * g3_19
	f9g4_19 := ref.Raw[9] * g4_19
	f9g5_38 := f9_2 * g5_19
	f9g6_19 := ref.Raw[9] * g6_19
	f9g7_38 := f9_2 * g7_19
	f9g8_19 := ref.Raw[9] * g8_19
	f9g9_38 := f9_2 * g9_19
	/**
	 * Remember: 2^255 congruent 19 modulo p.
	 * h = h[0] * 2^0 + h[1] * 2^26 + h[2] * 2^(26+25) + h[3] * 2^(26+25+26) + ... + h[9] * 2^(5*26+5*25).
	 * So to get the real number we would have to multiply the coefficients with the corresponding powers of 2.
	 * To get an idea what is going on below, look at the calculation of h[0]:
	 * h[0] is the coefficient to the power 2^0 so it collects (sums) all products that have the power 2^0.
	 * f0 * g0 really is f0 * 2^0 * g0 * 2^0 = (f0 * g0) * 2^0.
	 * f1 * g9 really is f1 * 2^26 * g9 * 2^230 = f1 * g9 * 2^256 = 2 * f1 * g9 * 2^255 congruent 2 * 19 * f1 * g9 * 2^0 modulo p.
	 * f2 * g8 really is f2 * 2^51 * g8 * 2^204 = f2 * g8 * 2^255 congruent 19 * f2 * g8 * 2^0 modulo p.
	 * and so on...
	 */
	h := []intRaw{
		f0g0 + f1g9_38 + f2g8_19 + f3g7_38 + f4g6_19 + f5g5_38 + f6g4_19 + f7g3_38 + f8g2_19 + f9g1_38,
		f0g1 + f1g0 + f2g9_19 + f3g8_19 + f4g7_19 + f5g6_19 + f6g5_19 + f7g4_19 + f8g3_19 + f9g2_19,
		f0g2 + f1g1_2 + f2g0 + f3g9_38 + f4g8_19 + f5g7_38 + f6g6_19 + f7g5_38 + f8g4_19 + f9g3_38,
		f0g3 + f1g2 + f2g1 + f3g0 + f4g9_19 + f5g8_19 + f6g7_19 + f7g6_19 + f8g5_19 + f9g4_19,
		f0g4 + f1g3_2 + f2g2 + f3g1_2 + f4g0 + f5g9_38 + f6g8_19 + f7g7_38 + f8g6_19 + f9g5_38,
		f0g5 + f1g4 + f2g3 + f3g2 + f4g1 + f5g0 + f6g9_19 + f7g8_19 + f8g7_19 + f9g6_19,
		f0g6 + f1g5_2 + f2g4 + f3g3_2 + f4g2 + f5g1_2 + f6g0 + f7g9_38 + f8g8_19 + f9g7_38,
		f0g7 + f1g6 + f2g5 + f3g4 + f4g3 + f5g2 + f6g1 + f7g0 + f8g9_19 + f9g8_19,
		f0g8 + f1g7_2 + f2g6 + f3g5_2 + f4g4 + f5g3_2 + f6g2 + f7g1_2 + f8g0 + f9g9_38,
		f0g9 + f1g8 + f2g7 + f3g6 + f4g5 + f5g4 + f6g3 + f7g2 + f8g1 + f9g0,
	}
	/**
	 * |h[0]| <= (1.65*1.65*2^52*(1+19+19+19+19)+1.65*1.65*2^50*(38+38+38+38+38))
	 * i.e. |h[0]| <= 1.4*2^60; narrower ranges for h[2], h[4], h[6], h[8]
	 * |h[1]| <= (1.65*1.65*2^51*(1+1+19+19+19+19+19+19+19+19))
	 * i.e. |h[1]| <= 1.7*2^59; narrower ranges for h[3], h[5], h[7], h[9]
	 */
	carry0 := (h[0] + (1 << 25)) >> 26
	h[1] += carry0
	h[0] -= carry0 << 26
	carry4 := (h[4] + (1 << 25)) >> 26
	h[5] += carry4
	h[4] -= carry4 << 26
	/* |h[0]| <= 2^25 */
	/* |h[4]| <= 2^25 */
	/* |h[1]| <= 1.71*2^59 */
	/* |h[5]| <= 1.71*2^59 */
	carry1 := (h[1] + (1 << 24)) >> 25
	h[2] += carry1
	h[1] -= carry1 << 25
	carry5 := (h[5] + (1 << 24)) >> 25
	h[6] += carry5
	h[5] -= carry5 << 25
	/* |h[1]| <= 2^24; from now on fits into int32 */
	/* |h[5]| <= 2^24; from now on fits into int32 */
	/* |h[2]| <= 1.41*2^60 */
	/* |h[6]| <= 1.41*2^60 */
	carry2 := (h[2] + (1 << 25)) >> 26
	h[3] += carry2
	h[2] -= carry2 << 26
	carry6 := (h[6] + (1 << 25)) >> 26
	h[7] += carry6
	h[6] -= carry6 << 26
	/* |h[2]| <= 2^25; from now on fits into int32 unchanged */
	/* |h[6]| <= 2^25; from now on fits into int32 unchanged */
	/* |h[3]| <= 1.71*2^59 */
	/* |h[7]| <= 1.71*2^59 */
	carry3 := (h[3] + (1 << 24)) >> 25
	h[4] += carry3
	h[3] -= carry3 << 25
	carry7 := (h[7] + (1 << 24)) >> 25
	h[8] += carry7
	h[7] -= carry7 << 25
	/* |h[3]| <= 2^24; from now on fits into int32 unchanged */
	/* |h[7]| <= 2^24; from now on fits into int32 unchanged */
	/* |h[4]| <= 1.72*2^34 */
	/* |h[8]| <= 1.41*2^60 */
	carry4 = (h[4] + (1 << 25)) >> 26
	h[5] += carry4
	h[4] -= carry4 << 26
	carry8 := (h[8] + (1 << 25)) >> 26
	h[9] += carry8
	h[8] -= carry8 << 26
	/* |h[4]| <= 2^25; from now on fits into int32 unchanged */
	/* |h[8]| <= 2^25; from now on fits into int32 unchanged */
	/* |h[5]| <= 1.01*2^24 */
	/* |h[9]| <= 1.71*2^59 */
	carry9 := (h[9] + (1 << 24)) >> 25
	h[0] += carry9 * 19
	h[9] -= carry9 << 25
	/* |h[9]| <= 2^24; from now on fits into int32 unchanged */
	/* |h[0]| <= 1.1*2^39 */
	carry0 = (h[0] + (1 << 25)) >> 26
	h[1] += carry0
	h[0] -= carry0 << 26
	/* |h[0]| <= 2^25; from now on fits into int32 unchanged */
	/* |h[1]| <= 1.01*2^24 */

	return Ed25519FieldElement{h}
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
func (ref Ed25519FieldElement) square() Ed25519FieldElement {

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
 */func (ref Ed25519FieldElement) squareAndDouble() Ed25519FieldElement {

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
 */func (ref Ed25519FieldElement) squareAndOptionalDouble(dbl bool) Ed25519FieldElement {

	f0_2 := 2 * ref.Raw[0]
	f1_2 := 2 * ref.Raw[1]
	f2_2 := 2 * ref.Raw[2]
	f3_2 := 2 * ref.Raw[3]
	f4_2 := 2 * ref.Raw[4]
	f5_2 := 2 * ref.Raw[5]
	f6_2 := 2 * ref.Raw[6]
	f7_2 := 2 * ref.Raw[7]
	f5_38 := 38 * ref.Raw[5] /* 1.959375*2^30 */
	f6_19 := 19 * ref.Raw[6] /* 1.959375*2^30 */
	f7_38 := 38 * ref.Raw[7] /* 1.959375*2^30 */
	f8_19 := 19 * ref.Raw[8] /* 1.959375*2^30 */
	f9_38 := 38 * ref.Raw[9] /* 1.959375*2^30 */

	f0f0 := ref.Raw[0] * ref.Raw[0]
	f0f1_2 := f0_2 * ref.Raw[1]
	f0f2_2 := f0_2 * ref.Raw[2]
	f0f3_2 := f0_2 * ref.Raw[3]
	f0f4_2 := f0_2 * ref.Raw[4]
	f0f5_2 := f0_2 * ref.Raw[5]
	f0f6_2 := f0_2 * ref.Raw[6]
	f0f7_2 := f0_2 * ref.Raw[7]
	f0f8_2 := f0_2 * ref.Raw[8]
	f0f9_2 := f0_2 * ref.Raw[9]

	f1f1_2 := f1_2 * ref.Raw[1]
	f1f2_2 := f1_2 * ref.Raw[2]
	f1f3_4 := f1_2 * f3_2
	f1f4_2 := f1_2 * ref.Raw[4]
	f1f5_4 := f1_2 * f5_2
	f1f6_2 := f1_2 * ref.Raw[6]
	f1f7_4 := f1_2 * f7_2
	f1f8_2 := f1_2 * ref.Raw[8]
	f1f9_76 := f1_2 * f9_38

	f2f2 := ref.Raw[2] * ref.Raw[2]
	f2f3_2 := f2_2 * ref.Raw[3]
	f2f4_2 := f2_2 * ref.Raw[4]
	f2f5_2 := f2_2 * ref.Raw[5]
	f2f6_2 := f2_2 * ref.Raw[6]
	f2f7_2 := f2_2 * ref.Raw[7]
	f2f8_38 := f2_2 * f8_19
	f2f9_38 := ref.Raw[2] * f9_38

	f3f3_2 := f3_2 * ref.Raw[3]
	f3f4_2 := f3_2 * ref.Raw[4]
	f3f5_4 := f3_2 * f5_2
	f3f6_2 := f3_2 * ref.Raw[6]
	f3f7_76 := f3_2 * f7_38
	f3f8_38 := f3_2 * f8_19
	f3f9_76 := f3_2 * f9_38

	f4f4 := ref.Raw[4] * ref.Raw[4]
	f4f5_2 := f4_2 * ref.Raw[5]
	f4f6_38 := f4_2 * f6_19
	f4f7_38 := ref.Raw[4] * f7_38
	f4f8_38 := f4_2 * f8_19
	f4f9_38 := ref.Raw[4] * f9_38

	f5f5_38 := ref.Raw[5] * f5_38
	f5f6_38 := f5_2 * f6_19
	f5f7_76 := f5_2 * f7_38
	f5f8_38 := f5_2 * f8_19
	f5f9_76 := f5_2 * f9_38

	f6f6_19 := ref.Raw[6] * f6_19
	f6f7_38 := ref.Raw[6] * f7_38
	f6f8_38 := f6_2 * f8_19
	f6f9_38 := ref.Raw[6] * f9_38

	f7f7_38 := ref.Raw[7] * f7_38
	f7f8_38 := f7_2 * f8_19
	f7f9_76 := f7_2 * f9_38
	f8f8_19 := ref.Raw[8] * f8_19
	f8f9_38 := ref.Raw[8] * f9_38
	f9f9_38 := ref.Raw[9] * f9_38

	var h [10]intRaw
	h[0] = f0f0 + f1f9_76 + f2f8_38 + f3f7_76 + f4f6_38 + f5f5_38
	h[1] = f0f1_2 + f2f9_38 + f3f8_38 + f4f7_38 + f5f6_38
	h[2] = f0f2_2 + f1f1_2 + f3f9_76 + f4f8_38 + f5f7_76 + f6f6_19
	h[3] = f0f3_2 + f1f2_2 + f4f9_38 + f5f8_38 + f6f7_38
	h[4] = f0f4_2 + f1f3_4 + f2f2 + f5f9_76 + f6f8_38 + f7f7_38
	h[5] = f0f5_2 + f1f4_2 + f2f3_2 + f6f9_38 + f7f8_38
	h[6] = f0f6_2 + f1f5_4 + f2f4_2 + f3f3_2 + f7f9_76 + f8f8_19
	h[7] = f0f7_2 + f1f6_2 + f2f5_2 + f3f4_2 + f8f9_38
	h[8] = f0f8_2 + f1f7_4 + f2f6_2 + f3f5_4 + f4f4 + f9f9_38
	h[9] = f0f9_2 + f1f8_2 + f2f7_2 + f3f6_2 + f4f5_2
	if dbl {
		for i, val := range h {
			h[i] += val
		}
	}

	carry0 := (h[0] + (1 << 25)) >> 26
	h[1] += carry0
	h[0] -= carry0 << 26
	carry4 := (h[4] + (1 << 25)) >> 26
	h[5] += carry4
	h[4] -= carry4 << 26
	carry1 := (h[1] + (1 << 24)) >> 25
	h[2] += carry1
	h[1] -= carry1 << 25
	carry5 := (h[5] + (1 << 24)) >> 25
	h[6] += carry5
	h[5] -= carry5 << 25
	carry2 := (h[2] + (1 << 25)) >> 26
	h[3] += carry2
	h[2] -= carry2 << 26
	carry6 := (h[6] + (1 << 25)) >> 26
	h[7] += carry6
	h[6] -= carry6 << 26
	carry3 := (h[3] + (1 << 24)) >> 25
	h[4] += carry3
	h[3] -= carry3 << 25
	carry7 := (h[7] + (1 << 24)) >> 25
	h[8] += carry7
	h[7] -= carry7 << 25
	carry4 = (h[4] + (1 << 25)) >> 26
	h[5] += carry4
	h[4] -= carry4 << 26
	carry8 := (h[8] + (1 << 25)) >> 26
	h[9] += carry8
	h[8] -= carry8 << 26
	carry9 := (h[9] + (1 << 24)) >> 25
	h[0] += carry9 * 19
	h[9] -= carry9 << 25
	carry0 = (h[0] + (1 << 25)) >> 26
	h[1] += carry0
	h[0] -= carry0 << 26
	return Ed25519FieldElement{h[:]}
}

/**
 * Invert ref field element and return the result.
 * The inverse is found via Fermat's little theorem:
 * a^p congruent a mod p and therefore a^(p-2) congruent a^-1 mod p
 *
 * @return The inverse of ref field element.
 */func (ref Ed25519FieldElement) invert() Ed25519FieldElement {

	// comments describe how exponent is created
	// 2 == 2 * 1
	f0 := ref.square()
	// 9 == 9
	f1 := ref.pow2to9()
	// 11 == 9 + 2
	f0 = f0.multiply(f1)
	// 2^252 - 2^2
	f1 = ref.pow2to252sub4()
	// 2^255 - 2^5
	for i := 1; i < 4; i++ {
		f1 = f1.square()
	}

	// 2^255 - 21
	return f1.multiply(f0)
}

//pow2to9 Ccmputes ref field element to the power of (2^9) and returns the result.
func (ref Ed25519FieldElement) pow2to9() Ed25519FieldElement {

	// 2 == 2 * 1
	f := ref.square()
	// 4 == 2 * 2
	f = f.square()
	// 8 == 2 * 4
	f = f.square()
	// 9 == 1 + 8
	return ref.multiply(f)
}

//pow2to252sub4 computes ref field element to the power of (2^252 - 4) and returns the result.
func (ref Ed25519FieldElement) pow2to252sub4() Ed25519FieldElement {

	// 2 == 2 * 1
	f0 := ref.square()
	// 9
	f1 := ref.pow2to9()
	// 11 == 9 + 2
	f0 = f0.multiply(f1)
	// 22 == 2 * 11
	f0 = f0.square()
	// 31 == 22 + 9
	f0 = f1.multiply(f0)
	// 2^6 - 2^1
	f1 = f0.square()
	// 2^10 - 2^5
	for i := 1; i < 5; i++ {
		f1 = f1.square()
	}

	// 2^10 - 2^0
	f0 = f1.multiply(f0)
	// 2^11 - 2^1
	f1 = f0.square()
	// 2^20 - 2^10
	for i := 1; i < 10; i++ {
		f1 = f1.square()
	}

	// 2^20 - 2^0
	f1 = f1.multiply(f0)
	// 2^21 - 2^1
	f2 := f1.square()
	// 2^40 - 2^20
	for i := 1; i < 20; i++ {
		f2 = f2.square()
	}

	// 2^40 - 2^0
	f1 = f2.multiply(f1)
	// 2^41 - 2^1
	f1 = f1.square()
	// 2^50 - 2^10
	for i := 1; i < 10; i++ {
		f1 = f1.square()
	}

	// 2^50 - 2^0
	f0 = f1.multiply(f0)
	// 2^51 - 2^1
	f1 = f0.square()
	// 2^100 - 2^50
	for i := 1; i < 50; i++ {
		f1 = f1.square()
	}

	// 2^100 - 2^0
	f1 = f1.multiply(f0)
	// 2^101 - 2^1
	f2 = f1.square()
	// 2^200 - 2^100
	for i := 1; i < 100; i++ {
		f2 = f2.square()
	}

	// 2^200 - 2^0
	f1 = f2.multiply(f1)
	// 2^201 - 2^1
	f1 = f1.square()
	// 2^250 - 2^50
	for i := 1; i < 50; i++ {
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
 * h = h[0] + 2^25 * h[1] + 2^(26+25) * h[2] + ... + 2^230 * h[9] where 0 <= |hi| < 2^27 for all i=0,...,9.
 * h congruent r modulo p, i.e. h = r + q * p for some suitable 0 <= r < p and an integer q.
 * <br>
 * Then q = [2^-255 * (h + 19 * 2^-25 * h[9] + 1/2)] where [x] = floor(x).
 * <br>
 * Proof:
 * We begin with some very raw estimation for the bounds of some expressions:
 *     |h| < 2^230 * 2^30 = 2^260 ==> |r + q * p| < 2^260 ==> |q| < 2^10.
 *         ==> -1/4 <= a := 19^2 * 2^-255 * q < 1/4.
 *     |h - 2^230 * h[9]| = |h[0] + ... + 2^204 * h[8]| < 2^204 * 2^30 = 2^234.
 *         ==> -1/4 <= b := 19 * 2^-255 * (h - 2^230 * h[9]) < 1/4
 * Therefore 0 < 1/2 - a - b < 1.
 * Set x := r + 19 * 2^-255 * r + 1/2 - a - b then
 *     0 <= x < 255 - 20 + 19 + 1 = 2^255 ==> 0 <= 2^-255 * x < 1. Since q is an integer we have
 *     [q + 2^-255 * x] = q        (1)
 * Have a closer look at x:
 *     x = h - q * (2^255 - 19) + 19 * 2^-255 * (h - q * (2^255 - 19)) + 1/2 - 19^2 * 2^-255 * q - 19 * 2^-255 * (h - 2^230 * h[9])
 *       = h - q * 2^255 + 19 * q + 19 * 2^-255 * h - 19 * q + 19^2 * 2^-255 * q + 1/2 - 19^2 * 2^-255 * q - 19 * 2^-255 * h + 19 * 2^-25 * h[9]
 *       = h + 19 * 2^-25 * h[9] + 1/2 - q^255.
 * Inserting the expression for x into (1) we get the desired expression for q.
 * }
 * </pre>
 *
 * @return The mod p reduced field element
 */func (ref Ed25519FieldElement) modP() Ed25519FieldElement {

	h := []intRaw{
		ref.Raw[0],
		ref.Raw[1],
		ref.Raw[2],
		ref.Raw[3],
		ref.Raw[4],
		ref.Raw[5],
		ref.Raw[6],
		ref.Raw[7],
		ref.Raw[8],
		ref.Raw[9],
	}
	// Calculate q
	q := (19*h[9] + (1 << 24)) >> 25
	q = (h[0] + q) >> 26
	q = (h[1] + q) >> 25
	q = (h[2] + q) >> 26
	q = (h[3] + q) >> 25
	q = (h[4] + q) >> 26
	q = (h[5] + q) >> 25
	q = (h[6] + q) >> 26
	q = (h[7] + q) >> 25
	q = (h[8] + q) >> 26
	q = (h[9] + q) >> 25
	// r = h - q * p = h - 2^255 * q + 19 * q
	// First add 19 * q then discard the bit 255
	h[0] += 19 * q
	carry0 := h[0] >> 26
	h[1] += carry0
	h[0] -= carry0 << 26
	carry1 := h[1] >> 25
	h[2] += carry1
	h[1] -= carry1 << 25
	carry2 := h[2] >> 26
	h[3] += carry2
	h[2] -= carry2 << 26
	carry3 := h[3] >> 25
	h[4] += carry3
	h[3] -= carry3 << 25
	carry4 := h[4] >> 26
	h[5] += carry4
	h[4] -= carry4 << 26
	carry5 := h[5] >> 25
	h[6] += carry5
	h[5] -= carry5 << 25
	carry6 := h[6] >> 26
	h[7] += carry6
	h[6] -= carry6 << 26
	carry7 := h[7] >> 25
	h[8] += carry7
	h[7] -= carry7 << 25
	carry8 := h[8] >> 26
	h[9] += carry8
	h[8] -= carry8 << 26
	carry9 := h[9] >> 25
	h[9] -= carry9 << 25

	return Ed25519FieldElement{h}
}

/**
 * Encodes a given field element in its 32 byte 2^8 bit representation. This is done in two steps.
 * Step 1: Reduce the value of the field element modulo p.
 * Step 2: Convert the field element to the 32 byte representation.
 *
 * @return Encoded field element (32 bytes).
 */func (ref Ed25519FieldElement) Encode() *Ed25519EncodedFieldElement {

	// Step 1:
	g := ref.modP()
	h := g.Raw
	// Step 2:
	s := make([]byte, 32)
	s[0] = (byte)(h[0])
	s[1] = (byte)(h[0] >> 8)
	s[2] = (byte)(h[0] >> 16)
	s[3] = (byte)((h[0] >> 24) | (h[1] << 2))
	s[4] = (byte)(h[1] >> 6)
	s[5] = (byte)(h[1] >> 14)
	s[6] = (byte)((h[1] >> 22) | (h[2] << 3))
	s[7] = (byte)(h[2] >> 5)
	s[8] = (byte)(h[2] >> 13)
	s[9] = (byte)((h[2] >> 21) | (h[3] << 5))
	s[10] = (byte)(h[3] >> 3)
	s[11] = (byte)(h[3] >> 11)
	s[12] = (byte)((h[3] >> 19) | (h[4] << 6))
	s[13] = (byte)(h[4] >> 2)
	s[14] = (byte)(h[4] >> 10)
	s[15] = (byte)(h[4] >> 18)
	s[16] = (byte)(h[5])
	s[17] = (byte)(h[5] >> 8)
	s[18] = (byte)(h[5] >> 16)
	s[19] = (byte)((h[5] >> 24) | (h[6] << 1))
	s[20] = (byte)(h[6] >> 7)
	s[21] = (byte)(h[6] >> 15)
	s[22] = (byte)((h[6] >> 23) | (h[7] << 3))
	s[23] = (byte)(h[7] >> 5)
	s[24] = (byte)(h[7] >> 13)
	s[25] = (byte)((h[7] >> 21) | (h[8] << 4))
	s[26] = (byte)(h[8] >> 4)
	s[27] = (byte)(h[8] >> 12)
	s[28] = (byte)((h[8] >> 20) | (h[9] << 6))
	s[29] = (byte)(h[9] >> 2)
	s[30] = (byte)(h[9] >> 10)
	s[31] = (byte)(h[9] >> 18)
	return &Ed25519EncodedFieldElement{Ed25519Field_ZERO_SHORT(), s}
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
 */func (ref Ed25519FieldElement) IsNegative() bool {

	return ref.Encode().IsNegative()
}
func (ref *Ed25519FieldElement) Equals(ge *Ed25519FieldElement) bool {

	return ref.Encode().Equals(ge.Encode())
}
func (ref *Ed25519FieldElement) String() string {

	return ref.Encode().String()
}

// Ed25519EncodedFieldElement Represents a field element of the finite field with p=2^255-19 elements.
// * The value of the field element is held in 2^8 bit representation, i.e. in a byte array.
// * The length of the array must be 32 or 64.
type Ed25519EncodedFieldElement struct {
	zero []byte
	Raw  []byte
}

//NewEd25519EncodedFieldElement creates a new encoded field element.
// Raw must to have leght 32 or 64 bytes
func NewEd25519EncodedFieldElement(Raw []byte) (*Ed25519EncodedFieldElement, error) {

	switch len(Raw) {
	case 32:
		return &Ed25519EncodedFieldElement{Ed25519Field_ZERO_SHORT(), Raw}, nil
	case 64:
		return &Ed25519EncodedFieldElement{Ed25519Field_ZERO_LONG(), Raw}, nil
	}
	return nil, errors.New("Invalid 2^8 bit representation.")
}
func (ref *Ed25519EncodedFieldElement) Equals(ge *Ed25519EncodedFieldElement) bool {
	return isEqualConstantTime(ref.Raw, ge.Raw)
}
func (ref *Ed25519EncodedFieldElement) threeBytesToLong(b []byte, offset int) intRaw {

	return intRaw(int64(b[offset]) | int64(b[offset+1])<<8 | int64(b[offset+2])<<16)
}

func (ref *Ed25519EncodedFieldElement) fourBytesToLong(b []byte, offset int) intRaw {

	return intRaw(int64(b[offset]) | int64(b[offset+1])<<8 | int64(b[offset+2])<<16 | int64(b[offset+3])<<24)
}

//IsNegative Return true if ref is in {1,3,5,...,q-2}
//  Return false if ref is in {0,2,4,...,q-1}
//* Preconditions:
//* |x| bounded by 1.1*2^26,1.1*2^25,1.1*2^26,1.1*2^25,etc.
//*
//* @return true if ref is in {1,3,5,...,q-2}, false otherwise.
func (ref *Ed25519EncodedFieldElement) IsNegative() bool {

	return (ref.Raw[0] & 1) != 0
}

/**
 * Gets a value indicating whether or not the field element is non-zero.
 *
 * @return 1 if it is non-zero, 0 otherwise.
 */func (ref *Ed25519EncodedFieldElement) IsNonZero() bool {

	return !isEqualConstantTime(ref.Raw, ref.zero)
}

/**
 * Decodes ref encoded (32 byte) representation to a field element in its 10 byte 2^25.5 representation.
 * The most significant bit is discarded.
 *
 * @return The field element in its 2^25.5 bit representation.
 */
func (ref *Ed25519EncodedFieldElement) Decode() *Ed25519FieldElement {

	h := []intRaw{
		ref.fourBytesToLong(ref.Raw, 0),
		ref.threeBytesToLong(ref.Raw, 4) << 6,
		ref.threeBytesToLong(ref.Raw, 7) << 5,
		ref.threeBytesToLong(ref.Raw, 10) << 3,
		ref.threeBytesToLong(ref.Raw, 13) << 2,
		ref.fourBytesToLong(ref.Raw, 16),
		ref.threeBytesToLong(ref.Raw, 20) << 7,
		ref.threeBytesToLong(ref.Raw, 23) << 5,
		ref.threeBytesToLong(ref.Raw, 26) << 4,
		(ref.threeBytesToLong(ref.Raw, 29) & 0x7FFFFF) << 2,
	}

	// Remember: 2^255 congruent 19 modulo p
	carry9 := (h[9] + (1 << 24)) >> 25
	h[0] += carry9 * 19
	h[9] -= carry9 << 25
	carry1 := (h[1] + (1 << 24)) >> 25
	h[2] += carry1
	h[1] -= carry1 << 25
	carry3 := (h[3] + (1 << 24)) >> 25
	h[4] += carry3
	h[3] -= carry3 << 25
	carry5 := (h[5] + (1 << 24)) >> 25
	h[6] += carry5
	h[5] -= carry5 << 25
	carry7 := (h[7] + (1 << 24)) >> 25
	h[8] += carry7
	h[7] -= carry7 << 25
	carry0 := (h[0] + (1 << 25)) >> 26
	h[1] += carry0
	h[0] -= carry0 << 26
	carry2 := (h[2] + (1 << 25)) >> 26
	h[3] += carry2
	h[2] -= carry2 << 26
	carry4 := (h[4] + (1 << 25)) >> 26
	h[5] += carry4
	h[4] -= carry4 << 26
	carry6 := (h[6] + (1 << 25)) >> 26
	h[7] += carry6
	h[6] -= carry6 << 26
	carry8 := (h[8] + (1 << 25)) >> 26
	h[9] += carry8
	h[8] -= carry8 << 26

	return &Ed25519FieldElement{h}
}

// modQ Reduces ref encoded field element (64 bytes) modulo the group order q.
// return Encoded field element (32 bytes).
func (ref *Ed25519EncodedFieldElement) modQ() *Ed25519EncodedFieldElement {

	// s0, ..., s22 have 21 bits, s23 has 29 bits
	s0 := 0x1FFFFF & ref.threeBytesToLong(ref.Raw, 0)
	s1 := 0x1FFFFF & (ref.fourBytesToLong(ref.Raw, 2) >> 5)
	s2 := 0x1FFFFF & (ref.threeBytesToLong(ref.Raw, 5) >> 2)
	s3 := 0x1FFFFF & (ref.fourBytesToLong(ref.Raw, 7) >> 7)
	s4 := 0x1FFFFF & (ref.fourBytesToLong(ref.Raw, 10) >> 4)
	s5 := 0x1FFFFF & (ref.threeBytesToLong(ref.Raw, 13) >> 1)
	s6 := 0x1FFFFF & (ref.fourBytesToLong(ref.Raw, 15) >> 6)
	s7 := 0x1FFFFF & (ref.threeBytesToLong(ref.Raw, 18) >> 3)
	s8 := 0x1FFFFF & ref.threeBytesToLong(ref.Raw, 21)
	s9 := 0x1FFFFF & (ref.fourBytesToLong(ref.Raw, 23) >> 5)
	s10 := 0x1FFFFF & (ref.threeBytesToLong(ref.Raw, 26) >> 2)
	s11 := 0x1FFFFF & (ref.fourBytesToLong(ref.Raw, 28) >> 7)
	s12 := 0x1FFFFF & (ref.fourBytesToLong(ref.Raw, 31) >> 4)
	s13 := 0x1FFFFF & (ref.threeBytesToLong(ref.Raw, 34) >> 1)
	s14 := 0x1FFFFF & (ref.fourBytesToLong(ref.Raw, 36) >> 6)
	s15 := 0x1FFFFF & (ref.threeBytesToLong(ref.Raw, 39) >> 3)
	s16 := 0x1FFFFF & ref.threeBytesToLong(ref.Raw, 42)
	s17 := 0x1FFFFF & (ref.fourBytesToLong(ref.Raw, 44) >> 5)
	s18 := 0x1FFFFF & (ref.threeBytesToLong(ref.Raw, 47) >> 2)

	s19 := 0x1FFFFF & (ref.fourBytesToLong(ref.Raw, 49) >> 7)

	s20 := 0x1FFFFF & (ref.fourBytesToLong(ref.Raw, 52) >> 4)

	s21 := 0x1FFFFF & (ref.threeBytesToLong(ref.Raw, 55) >> 1)

	s22 := 0x1FFFFF & (ref.fourBytesToLong(ref.Raw, 57) >> 6)

	s23 := ref.fourBytesToLong(ref.Raw, 60) >> 3

	/**
	 * Lots of magic numbers :)
	 * To understand what's going on below, note that
	 *
	 * (1) q = 2^252 + q0 where q0 = 27742317777372353535851937790883648493.
	 * (2) s11 is the coefficient of 2^(11*21), s23 is the coefficient of 2^(^23*21) and 2^252 = 2^((23-11) * 21)).
	 * (3) 2^252 congruent -q0 modulo q.
	 * (4) -q0 = 666643 * 2^0 + 470296 * 2^21 + 654183 * 2^(2*21) - 997805 * 2^(3*21) + 136657 * 2^(4*21) - 683901 * 2^(5*21)
	 *
	 * Thus
	 * s23 * 2^(23*11) = s23 * 2^(12*21) * 2^(11*21) = s3 * 2^252 * 2^(11*21) congruent
	 * s23 * (666643 * 2^0 + 470296 * 2^21 + 654183 * 2^(2*21) - 997805 * 2^(3*21) + 136657 * 2^(4*21) - 683901 * 2^(5*21)) * 2^(11*21) modulo q =
	 * s23 * (666643 * 2^(11*21) + 470296 * 2^(12*21) + 654183 * 2^(13*21) - 997805 * 2^(14*21) + 136657 * 2^(15*21) - 683901 * 2^(16*21)).
	 *
	 * The same procedure is then applied for s22,...,s18.
	 */
	s11 += s23 * 666643
	s12 += s23 * 470296
	s13 += s23 * 654183
	s14 -= s23 * 997805
	s15 += s23 * 136657
	s16 -= s23 * 683901
	s10 += s22 * 666643
	s11 += s22 * 470296
	s12 += s22 * 654183
	s13 -= s22 * 997805
	s14 += s22 * 136657
	s15 -= s22 * 683901
	s9 += s21 * 666643
	s10 += s21 * 470296
	s11 += s21 * 654183
	s12 -= s21 * 997805
	s13 += s21 * 136657
	s14 -= s21 * 683901
	s8 += s20 * 666643
	s9 += s20 * 470296
	s10 += s20 * 654183
	s11 -= s20 * 997805
	s12 += s20 * 136657
	s13 -= s20 * 683901
	s7 += s19 * 666643
	s8 += s19 * 470296
	s9 += s19 * 654183
	s10 -= s19 * 997805
	s11 += s19 * 136657
	s12 -= s19 * 683901
	s6 += s18 * 666643
	s7 += s18 * 470296
	s8 += s18 * 654183
	s9 -= s18 * 997805
	s10 += s18 * 136657
	s11 -= s18 * 683901
	/**
	 * Time to reduce the coefficient in order not to get an overflow.
	 */
	carry6 := (s6 + (1 << 20)) >> 21
	s7 += carry6
	s6 -= carry6 << 21
	carry8 := (s8 + (1 << 20)) >> 21
	s9 += carry8
	s8 -= carry8 << 21
	carry10 := (s10 + (1 << 20)) >> 21
	s11 += carry10
	s10 -= carry10 << 21
	carry12 := (s12 + (1 << 20)) >> 21
	s13 += carry12
	s12 -= carry12 << 21
	carry14 := (s14 + (1 << 20)) >> 21
	s15 += carry14
	s14 -= carry14 << 21
	carry16 := (s16 + (1 << 20)) >> 21
	s17 += carry16
	s16 -= carry16 << 21
	carry7 := (s7 + (1 << 20)) >> 21
	s8 += carry7
	s7 -= carry7 << 21
	carry9 := (s9 + (1 << 20)) >> 21
	s10 += carry9
	s9 -= carry9 << 21
	carry11 := (s11 + (1 << 20)) >> 21
	s12 += carry11
	s11 -= carry11 << 21
	carry13 := (s13 + (1 << 20)) >> 21
	s14 += carry13
	s13 -= carry13 << 21
	carry15 := (s15 + (1 << 20)) >> 21
	s16 += carry15
	s15 -= carry15 << 21
	/**
	 * Continue with above procedure.
	 */
	s5 += s17 * 666643
	s6 += s17 * 470296
	s7 += s17 * 654183
	s8 -= s17 * 997805
	s9 += s17 * 136657
	s10 -= s17 * 683901
	s4 += s16 * 666643
	s5 += s16 * 470296
	s6 += s16 * 654183
	s7 -= s16 * 997805
	s8 += s16 * 136657
	s9 -= s16 * 683901
	s3 += s15 * 666643
	s4 += s15 * 470296
	s5 += s15 * 654183
	s6 -= s15 * 997805
	s7 += s15 * 136657
	s8 -= s15 * 683901
	s2 += s14 * 666643
	s3 += s14 * 470296
	s4 += s14 * 654183
	s5 -= s14 * 997805
	s6 += s14 * 136657
	s7 -= s14 * 683901
	s1 += s13 * 666643
	s2 += s13 * 470296
	s3 += s13 * 654183
	s4 -= s13 * 997805
	s5 += s13 * 136657
	s6 -= s13 * 683901
	s0 += s12 * 666643
	s1 += s12 * 470296
	s2 += s12 * 654183
	s3 -= s12 * 997805
	s4 += s12 * 136657
	s5 -= s12 * 683901
	s12 = 0
	/**
	 * Reduce coefficients again.
	 */
	carry0 := (s0 + (1 << 20)) >> 21
	s1 += carry0
	s0 -= carry0 << 21
	carry2 := (s2 + (1 << 20)) >> 21
	s3 += carry2
	s2 -= carry2 << 21
	carry4 := (s4 + (1 << 20)) >> 21
	s5 += carry4
	s4 -= carry4 << 21
	carry6 = (s6 + (1 << 20)) >> 21
	s7 += carry6
	s6 -= carry6 << 21
	carry8 = (s8 + (1 << 20)) >> 21
	s9 += carry8
	s8 -= carry8 << 21
	carry10 = (s10 + (1 << 20)) >> 21
	s11 += carry10
	s10 -= carry10 << 21
	carry1 := (s1 + (1 << 20)) >> 21
	s2 += carry1
	s1 -= carry1 << 21
	carry3 := (s3 + (1 << 20)) >> 21
	s4 += carry3
	s3 -= carry3 << 21
	carry5 := (s5 + (1 << 20)) >> 21
	s6 += carry5
	s5 -= carry5 << 21
	carry7 = (s7 + (1 << 20)) >> 21
	s8 += carry7
	s7 -= carry7 << 21
	carry9 = (s9 + (1 << 20)) >> 21
	s10 += carry9
	s9 -= carry9 << 21
	carry11 = (s11 + (1 << 20)) >> 21
	s12 += carry11
	s11 -= carry11 << 21
	s0 += s12 * 666643
	s1 += s12 * 470296
	s2 += s12 * 654183
	s3 -= s12 * 997805
	s4 += s12 * 136657
	s5 -= s12 * 683901
	s12 = 0
	carry0 = s0 >> 21
	s1 += carry0
	s0 -= carry0 << 21
	carry1 = s1 >> 21
	s2 += carry1
	s1 -= carry1 << 21
	carry2 = s2 >> 21
	s3 += carry2
	s2 -= carry2 << 21
	carry3 = s3 >> 21
	s4 += carry3
	s3 -= carry3 << 21
	carry4 = s4 >> 21
	s5 += carry4
	s4 -= carry4 << 21
	carry5 = s5 >> 21
	s6 += carry5
	s5 -= carry5 << 21
	carry6 = s6 >> 21
	s7 += carry6
	s6 -= carry6 << 21
	carry7 = s7 >> 21
	s8 += carry7
	s7 -= carry7 << 21
	carry8 = s8 >> 21
	s9 += carry8
	s8 -= carry8 << 21
	carry9 = s9 >> 21
	s10 += carry9
	s9 -= carry9 << 21
	carry10 = s10 >> 21
	s11 += carry10
	s10 -= carry10 << 21
	carry11 = s11 >> 21
	s12 += carry11
	s11 -= carry11 << 21
	s0 += s12 * 666643
	s1 += s12 * 470296
	s2 += s12 * 654183
	s3 -= s12 * 997805
	s4 += s12 * 136657
	s5 -= s12 * 683901
	carry0 = s0 >> 21
	s1 += carry0
	s0 -= carry0 << 21
	carry1 = s1 >> 21
	s2 += carry1
	s1 -= carry1 << 21
	carry2 = s2 >> 21
	s3 += carry2
	s2 -= carry2 << 21
	carry3 = s3 >> 21
	s4 += carry3
	s3 -= carry3 << 21
	carry4 = s4 >> 21
	s5 += carry4
	s4 -= carry4 << 21
	carry5 = s5 >> 21
	s6 += carry5
	s5 -= carry5 << 21
	carry6 = s6 >> 21
	s7 += carry6
	s6 -= carry6 << 21
	carry7 = s7 >> 21
	s8 += carry7
	s7 -= carry7 << 21
	carry8 = s8 >> 21
	s9 += carry8
	s8 -= carry8 << 21
	carry9 = s9 >> 21
	s10 += carry9
	s9 -= carry9 << 21
	carry10 = s10 >> 21
	s11 += carry10
	s10 -= carry10 << 21
	// s0, ..., s11 got 21 bits each.
	result := make([]byte, 32)
	result[0] = byte(s0)
	result[1] = (byte)(s0 >> 8)
	result[2] = (byte)((s0 >> 16) | (s1 << 5))
	result[3] = (byte)(s1 >> 3)
	result[4] = (byte)(s1 >> 11)
	result[5] = (byte)((s1 >> 19) | (s2 << 2))
	result[6] = (byte)(s2 >> 6)
	result[7] = (byte)((s2 >> 14) | (s3 << 7))
	result[8] = (byte)(s3 >> 1)
	result[9] = (byte)(s3 >> 9)
	result[10] = (byte)((s3 >> 17) | (s4 << 4))
	result[11] = (byte)(s4 >> 4)
	result[12] = (byte)(s4 >> 12)
	result[13] = (byte)((s4 >> 20) | (s5 << 1))
	result[14] = (byte)(s5 >> 7)
	result[15] = (byte)((s5 >> 15) | (s6 << 6))
	result[16] = (byte)(s6 >> 2)
	result[17] = (byte)(s6 >> 10)
	result[18] = (byte)((s6 >> 18) | (s7 << 3))
	result[19] = (byte)(s7 >> 5)
	result[20] = (byte)(s7 >> 13)
	result[21] = (byte)(s8)
	result[22] = (byte)(s8 >> 8)
	result[23] = (byte)((s8 >> 16) | (s9 << 5))
	result[24] = (byte)(s9 >> 3)
	result[25] = (byte)(s9 >> 11)
	result[26] = (byte)((s9 >> 19) | (s10 << 2))
	result[27] = (byte)(s10 >> 6)
	result[28] = (byte)((s10 >> 14) | (s11 << 7))
	result[29] = (byte)(s11 >> 1)
	result[30] = (byte)(s11 >> 9)
	result[31] = (byte)(s11 >> 17)
	return &Ed25519EncodedFieldElement{Ed25519Field_ZERO_SHORT(), result}
}

/**
 * Multiplies ref encoded field element with another and adds a third.
 * The result is reduced modulo the group order.
 * <br>
 * See the comments in the method modQ() for an explanation of the algorithm.
 *
 * @param b The encoded field element which is multiplied with ref.
 * @param c The third encoded field element which is added.
 * @return The encoded field element (32 bytes).
 */
func (ref *Ed25519EncodedFieldElement) multiplyAndAddModQ(
	b *Ed25519EncodedFieldElement,
	c *Ed25519EncodedFieldElement) *Ed25519EncodedFieldElement {
	a0 := 0x1FFFFF & ref.threeBytesToLong(ref.Raw, 0)

	a1 := 0x1FFFFF & (ref.fourBytesToLong(ref.Raw, 2) >> 5)

	a2 := 0x1FFFFF & (ref.threeBytesToLong(ref.Raw, 5) >> 2)

	a3 := 0x1FFFFF & (ref.fourBytesToLong(ref.Raw, 7) >> 7)

	a4 := 0x1FFFFF & (ref.fourBytesToLong(ref.Raw, 10) >> 4)

	a5 := 0x1FFFFF & (ref.threeBytesToLong(ref.Raw, 13) >> 1)

	a6 := 0x1FFFFF & (ref.fourBytesToLong(ref.Raw, 15) >> 6)

	a7 := 0x1FFFFF & (ref.threeBytesToLong(ref.Raw, 18) >> 3)

	a8 := 0x1FFFFF & ref.threeBytesToLong(ref.Raw, 21)

	a9 := 0x1FFFFF & (ref.fourBytesToLong(ref.Raw, 23) >> 5)

	a10 := 0x1FFFFF & (ref.threeBytesToLong(ref.Raw, 26) >> 2)

	a11 := ref.fourBytesToLong(ref.Raw, 28) >> 7

	b0 := 0x1FFFFF & ref.threeBytesToLong(b.Raw, 0)

	b1 := 0x1FFFFF & (ref.fourBytesToLong(b.Raw, 2) >> 5)

	b2 := 0x1FFFFF & (ref.threeBytesToLong(b.Raw, 5) >> 2)

	b3 := 0x1FFFFF & (ref.fourBytesToLong(b.Raw, 7) >> 7)

	b4 := 0x1FFFFF & (ref.fourBytesToLong(b.Raw, 10) >> 4)

	b5 := 0x1FFFFF & (ref.threeBytesToLong(b.Raw, 13) >> 1)

	b6 := 0x1FFFFF & (ref.fourBytesToLong(b.Raw, 15) >> 6)

	b7 := 0x1FFFFF & (ref.threeBytesToLong(b.Raw, 18) >> 3)

	b8 := 0x1FFFFF & ref.threeBytesToLong(b.Raw, 21)

	b9 := 0x1FFFFF & (ref.fourBytesToLong(b.Raw, 23) >> 5)

	b10 := 0x1FFFFF & (ref.threeBytesToLong(b.Raw, 26) >> 2)

	b11 := ref.fourBytesToLong(b.Raw, 28) >> 7

	c0 := 0x1FFFFF & ref.threeBytesToLong(c.Raw, 0)

	c1 := 0x1FFFFF & (ref.fourBytesToLong(c.Raw, 2) >> 5)

	c2 := 0x1FFFFF & (ref.threeBytesToLong(c.Raw, 5) >> 2)

	c3 := 0x1FFFFF & (ref.fourBytesToLong(c.Raw, 7) >> 7)

	c4 := 0x1FFFFF & (ref.fourBytesToLong(c.Raw, 10) >> 4)

	c5 := 0x1FFFFF & (ref.threeBytesToLong(c.Raw, 13) >> 1)

	c6 := 0x1FFFFF & (ref.fourBytesToLong(c.Raw, 15) >> 6)

	c7 := 0x1FFFFF & (ref.threeBytesToLong(c.Raw, 18) >> 3)

	c8 := 0x1FFFFF & ref.threeBytesToLong(c.Raw, 21)

	c9 := 0x1FFFFF & (ref.fourBytesToLong(c.Raw, 23) >> 5)

	c10 := 0x1FFFFF & (ref.threeBytesToLong(c.Raw, 26) >> 2)

	c11 := ref.fourBytesToLong(c.Raw, 28) >> 7

	s0 := c0 + a0*b0
	s1 := c1 + a0*b1 + a1*b0
	s2 := c2 + a0*b2 + a1*b1 + a2*b0
	s3 := c3 + a0*b3 + a1*b2 + a2*b1 + a3*b0
	s4 := c4 + a0*b4 + a1*b3 + a2*b2 + a3*b1 + a4*b0
	s5 := c5 + a0*b5 + a1*b4 + a2*b3 + a3*b2 + a4*b1 + a5*b0
	s6 := c6 + a0*b6 + a1*b5 + a2*b4 + a3*b3 + a4*b2 + a5*b1 + a6*b0
	s7 := c7 + a0*b7 + a1*b6 + a2*b5 + a3*b4 + a4*b3 + a5*b2 + a6*b1 + a7*b0
	s8 := c8 + a0*b8 + a1*b7 + a2*b6 + a3*b5 + a4*b4 + a5*b3 + a6*b2 + a7*b1 + a8*b0
	s9 := c9 + a0*b9 + a1*b8 + a2*b7 + a3*b6 + a4*b5 + a5*b4 + a6*b3 + a7*b2 + a8*b1 + a9*b0
	s10 := c10 + a0*b10 + a1*b9 + a2*b8 + a3*b7 + a4*b6 + a5*b5 + a6*b4 + a7*b3 + a8*b2 + a9*b1 + a10*b0
	s11 := c11 + a0*b11 + a1*b10 + a2*b9 + a3*b8 + a4*b7 + a5*b6 + a6*b5 + a7*b4 + a8*b3 + a9*b2 + a10*b1 + a11*b0
	s12 := a1*b11 + a2*b10 + a3*b9 + a4*b8 + a5*b7 + a6*b6 + a7*b5 + a8*b4 + a9*b3 + a10*b2 + a11*b1
	s13 := a2*b11 + a3*b10 + a4*b9 + a5*b8 + a6*b7 + a7*b6 + a8*b5 + a9*b4 + a10*b3 + a11*b2
	s14 := a3*b11 + a4*b10 + a5*b9 + a6*b8 + a7*b7 + a8*b6 + a9*b5 + a10*b4 + a11*b3
	s15 := a4*b11 + a5*b10 + a6*b9 + a7*b8 + a8*b7 + a9*b6 + a10*b5 + a11*b4
	s16 := a5*b11 + a6*b10 + a7*b9 + a8*b8 + a9*b7 + a10*b6 + a11*b5
	s17 := a6*b11 + a7*b10 + a8*b9 + a9*b8 + a10*b7 + a11*b6
	s18 := a7*b11 + a8*b10 + a9*b9 + a10*b8 + a11*b7
	s19 := a8*b11 + a9*b10 + a10*b9 + a11*b8
	s20 := a9*b11 + a10*b10 + a11*b9
	s21 := a10*b11 + a11*b10
	s22 := a11 * b11
	s23 := intRaw(0)
	carry0 := (s0 + (1 << 20)) >> 21
	s1 += carry0
	s0 -= carry0 << 21
	carry2 := (s2 + (1 << 20)) >> 21
	s3 += carry2
	s2 -= carry2 << 21
	carry4 := (s4 + (1 << 20)) >> 21
	s5 += carry4
	s4 -= carry4 << 21
	carry6 := (s6 + (1 << 20)) >> 21
	s7 += carry6
	s6 -= carry6 << 21
	carry8 := (s8 + (1 << 20)) >> 21
	s9 += carry8
	s8 -= carry8 << 21
	carry10 := (s10 + (1 << 20)) >> 21
	s11 += carry10
	s10 -= carry10 << 21
	carry12 := (s12 + (1 << 20)) >> 21
	s13 += carry12
	s12 -= carry12 << 21
	carry14 := (s14 + (1 << 20)) >> 21
	s15 += carry14
	s14 -= carry14 << 21
	carry16 := (s16 + (1 << 20)) >> 21
	s17 += carry16
	s16 -= carry16 << 21
	carry18 := (s18 + (1 << 20)) >> 21
	s19 += carry18
	s18 -= carry18 << 21
	carry20 := (s20 + (1 << 20)) >> 21
	s21 += carry20
	s20 -= carry20 << 21
	carry22 := (s22 + (1 << 20)) >> 21
	s23 += carry22
	s22 -= carry22 << 21
	carry1 := (s1 + (1 << 20)) >> 21
	s2 += carry1
	s1 -= carry1 << 21
	carry3 := (s3 + (1 << 20)) >> 21
	s4 += carry3
	s3 -= carry3 << 21
	carry5 := (s5 + (1 << 20)) >> 21
	s6 += carry5
	s5 -= carry5 << 21
	carry7 := (s7 + (1 << 20)) >> 21
	s8 += carry7
	s7 -= carry7 << 21
	carry9 := (s9 + (1 << 20)) >> 21
	s10 += carry9
	s9 -= carry9 << 21
	carry11 := (s11 + (1 << 20)) >> 21
	s12 += carry11
	s11 -= carry11 << 21
	carry13 := (s13 + (1 << 20)) >> 21
	s14 += carry13
	s13 -= carry13 << 21
	carry15 := (s15 + (1 << 20)) >> 21
	s16 += carry15
	s15 -= carry15 << 21
	carry17 := (s17 + (1 << 20)) >> 21
	s18 += carry17
	s17 -= carry17 << 21
	carry19 := (s19 + (1 << 20)) >> 21
	s20 += carry19
	s19 -= carry19 << 21
	carry21 := (s21 + (1 << 20)) >> 21
	s22 += carry21
	s21 -= carry21 << 21
	s11 += s23 * 666643
	s12 += s23 * 470296
	s13 += s23 * 654183
	s14 -= s23 * 997805
	s15 += s23 * 136657
	s16 -= s23 * 683901
	s10 += s22 * 666643
	s11 += s22 * 470296
	s12 += s22 * 654183
	s13 -= s22 * 997805
	s14 += s22 * 136657
	s15 -= s22 * 683901
	s9 += s21 * 666643
	s10 += s21 * 470296
	s11 += s21 * 654183
	s12 -= s21 * 997805
	s13 += s21 * 136657
	s14 -= s21 * 683901
	s8 += s20 * 666643
	s9 += s20 * 470296
	s10 += s20 * 654183
	s11 -= s20 * 997805
	s12 += s20 * 136657
	s13 -= s20 * 683901
	s7 += s19 * 666643
	s8 += s19 * 470296
	s9 += s19 * 654183
	s10 -= s19 * 997805
	s11 += s19 * 136657
	s12 -= s19 * 683901
	s6 += s18 * 666643
	s7 += s18 * 470296
	s8 += s18 * 654183
	s9 -= s18 * 997805
	s10 += s18 * 136657
	s11 -= s18 * 683901
	carry6 = (s6 + (1 << 20)) >> 21
	s7 += carry6
	s6 -= carry6 << 21
	carry8 = (s8 + (1 << 20)) >> 21
	s9 += carry8
	s8 -= carry8 << 21
	carry10 = (s10 + (1 << 20)) >> 21
	s11 += carry10
	s10 -= carry10 << 21
	carry12 = (s12 + (1 << 20)) >> 21
	s13 += carry12
	s12 -= carry12 << 21
	carry14 = (s14 + (1 << 20)) >> 21
	s15 += carry14
	s14 -= carry14 << 21
	carry16 = (s16 + (1 << 20)) >> 21
	s17 += carry16
	s16 -= carry16 << 21
	carry7 = (s7 + (1 << 20)) >> 21
	s8 += carry7
	s7 -= carry7 << 21
	carry9 = (s9 + (1 << 20)) >> 21
	s10 += carry9
	s9 -= carry9 << 21
	carry11 = (s11 + (1 << 20)) >> 21
	s12 += carry11
	s11 -= carry11 << 21
	carry13 = (s13 + (1 << 20)) >> 21
	s14 += carry13
	s13 -= carry13 << 21
	carry15 = (s15 + (1 << 20)) >> 21
	s16 += carry15
	s15 -= carry15 << 21
	s5 += s17 * 666643
	s6 += s17 * 470296
	s7 += s17 * 654183
	s8 -= s17 * 997805
	s9 += s17 * 136657
	s10 -= s17 * 683901
	s4 += s16 * 666643
	s5 += s16 * 470296
	s6 += s16 * 654183
	s7 -= s16 * 997805
	s8 += s16 * 136657
	s9 -= s16 * 683901
	s3 += s15 * 666643
	s4 += s15 * 470296
	s5 += s15 * 654183
	s6 -= s15 * 997805
	s7 += s15 * 136657
	s8 -= s15 * 683901
	s2 += s14 * 666643
	s3 += s14 * 470296
	s4 += s14 * 654183
	s5 -= s14 * 997805
	s6 += s14 * 136657
	s7 -= s14 * 683901
	s1 += s13 * 666643
	s2 += s13 * 470296
	s3 += s13 * 654183
	s4 -= s13 * 997805
	s5 += s13 * 136657
	s6 -= s13 * 683901
	s0 += s12 * 666643
	s1 += s12 * 470296
	s2 += s12 * 654183
	s3 -= s12 * 997805
	s4 += s12 * 136657
	s5 -= s12 * 683901
	s12 = 0
	carry0 = (s0 + (1 << 20)) >> 21
	s1 += carry0
	s0 -= carry0 << 21
	carry2 = (s2 + (1 << 20)) >> 21
	s3 += carry2
	s2 -= carry2 << 21
	carry4 = (s4 + (1 << 20)) >> 21
	s5 += carry4
	s4 -= carry4 << 21
	carry6 = (s6 + (1 << 20)) >> 21
	s7 += carry6
	s6 -= carry6 << 21
	carry8 = (s8 + (1 << 20)) >> 21
	s9 += carry8
	s8 -= carry8 << 21
	carry10 = (s10 + (1 << 20)) >> 21
	s11 += carry10
	s10 -= carry10 << 21
	carry1 = (s1 + (1 << 20)) >> 21
	s2 += carry1
	s1 -= carry1 << 21
	carry3 = (s3 + (1 << 20)) >> 21
	s4 += carry3
	s3 -= carry3 << 21
	carry5 = (s5 + (1 << 20)) >> 21
	s6 += carry5
	s5 -= carry5 << 21
	carry7 = (s7 + (1 << 20)) >> 21
	s8 += carry7
	s7 -= carry7 << 21
	carry9 = (s9 + (1 << 20)) >> 21
	s10 += carry9
	s9 -= carry9 << 21
	carry11 = (s11 + (1 << 20)) >> 21
	s12 += carry11
	s11 -= carry11 << 21
	s0 += s12 * 666643
	s1 += s12 * 470296
	s2 += s12 * 654183
	s3 -= s12 * 997805
	s4 += s12 * 136657
	s5 -= s12 * 683901
	s12 = 0
	carry0 = s0 >> 21
	s1 += carry0
	s0 -= carry0 << 21
	carry1 = s1 >> 21
	s2 += carry1
	s1 -= carry1 << 21
	carry2 = s2 >> 21
	s3 += carry2
	s2 -= carry2 << 21
	carry3 = s3 >> 21
	s4 += carry3
	s3 -= carry3 << 21
	carry4 = s4 >> 21
	s5 += carry4
	s4 -= carry4 << 21
	carry5 = s5 >> 21
	s6 += carry5
	s5 -= carry5 << 21
	carry6 = s6 >> 21
	s7 += carry6
	s6 -= carry6 << 21
	carry7 = s7 >> 21
	s8 += carry7
	s7 -= carry7 << 21
	carry8 = s8 >> 21
	s9 += carry8
	s8 -= carry8 << 21
	carry9 = s9 >> 21
	s10 += carry9
	s9 -= carry9 << 21
	carry10 = s10 >> 21
	s11 += carry10
	s10 -= carry10 << 21
	carry11 = s11 >> 21
	s12 += carry11
	s11 -= carry11 << 21
	s0 += s12 * 666643
	s1 += s12 * 470296
	s2 += s12 * 654183
	s3 -= s12 * 997805
	s4 += s12 * 136657
	s5 -= s12 * 683901
	carry0 = s0 >> 21
	s1 += carry0
	s0 -= carry0 << 21
	carry1 = s1 >> 21
	s2 += carry1
	s1 -= carry1 << 21
	carry2 = s2 >> 21
	s3 += carry2
	s2 -= carry2 << 21
	carry3 = s3 >> 21
	s4 += carry3
	s3 -= carry3 << 21
	carry4 = s4 >> 21
	s5 += carry4
	s4 -= carry4 << 21
	carry5 = s5 >> 21
	s6 += carry5
	s5 -= carry5 << 21
	carry6 = s6 >> 21
	s7 += carry6
	s6 -= carry6 << 21
	carry7 = s7 >> 21
	s8 += carry7
	s7 -= carry7 << 21
	carry8 = s8 >> 21
	s9 += carry8
	s8 -= carry8 << 21
	carry9 = s9 >> 21
	s10 += carry9
	s9 -= carry9 << 21
	carry10 = s10 >> 21
	s11 += carry10
	s10 -= carry10 << 21
	result := make([]byte, 32)
	result[0] = (byte)(s0)
	result[1] = (byte)(s0 >> 8)
	result[2] = (byte)((s0 >> 16) | (s1 << 5))
	result[3] = (byte)(s1 >> 3)
	result[4] = (byte)(s1 >> 11)
	result[5] = (byte)((s1 >> 19) | (s2 << 2))
	result[6] = (byte)(s2 >> 6)
	result[7] = (byte)((s2 >> 14) | (s3 << 7))
	result[8] = (byte)(s3 >> 1)
	result[9] = (byte)(s3 >> 9)
	result[10] = (byte)((s3 >> 17) | (s4 << 4))
	result[11] = (byte)(s4 >> 4)
	result[12] = (byte)(s4 >> 12)
	result[13] = (byte)((s4 >> 20) | (s5 << 1))
	result[14] = (byte)(s5 >> 7)
	result[15] = (byte)((s5 >> 15) | (s6 << 6))
	result[16] = (byte)(s6 >> 2)
	result[17] = (byte)(s6 >> 10)
	result[18] = (byte)((s6 >> 18) | (s7 << 3))
	result[19] = (byte)(s7 >> 5)
	result[20] = (byte)(s7 >> 13)
	result[21] = (byte)(s8)
	result[22] = (byte)(s8 >> 8)
	result[23] = (byte)((s8 >> 16) | (s9 << 5))
	result[24] = (byte)(s9 >> 3)
	result[25] = (byte)(s9 >> 11)
	result[26] = (byte)((s9 >> 19) | (s10 << 2))
	result[27] = (byte)(s10 >> 6)
	result[28] = (byte)((s10 >> 14) | (s11 << 7))
	result[29] = (byte)(s11 >> 1)
	result[30] = (byte)(s11 >> 9)
	result[31] = (byte)(s11 >> 17)
	return &Ed25519EncodedFieldElement{Ed25519Field_ZERO_SHORT(), result}
}

func (ref *Ed25519EncodedFieldElement) String() string {

	return hex.EncodeToString(ref.Raw)
}

// Ed25519EncodedGroupElement
type Ed25519EncodedGroupElement struct {
	Raw []byte
}

//NewEd25519EncodedGroupElement creates a new encoded group element.
func NewEd25519EncodedGroupElement(Raw []byte) (*Ed25519EncodedGroupElement, error) {
	if 32 != len(Raw) {
		return nil, errors.New("Invalid encoded group element.")
	}

	return &Ed25519EncodedGroupElement{Raw}, nil
}

//Decode Decodes ref encoded group element and returns a new group element in P3 coordinates.
func (ref *Ed25519EncodedGroupElement) Decode() (*Ed25519GroupElement, error) {

	x, err := ref.GetAffineX()
	if err != nil {
		return nil, err
	}
	y, err := ref.GetAffineY()
	if err != nil {
		return nil, err
	}

	t := x.multiply(*y)
	return NewEd25519GroupElementP3(x, y, Ed25519Field_ONE(), &t), nil
}

//GetAffineX gets the affine x-coordinate.
// * x is recovered in the following way (p = field size):
// * <br>
// * x = sign(x) * sqrt((y^2 - 1) / (d * y^2 + 1)) = sign(x) * sqrt(u / v) with u = y^2 - 1 and v = d * y^2 + 1.
// * Setting ОІ = (u * v^3) * (u * v^7)^((p - 5) / 8) one has ОІ^2 = +-(u / v).
// * If v * ОІ = -u multiply ОІ with i=sqrt(-1).
// * Set x := ОІ.
// * If sign(x) != bit 255 of s then negate x.
//* @return the affine x-coordinate.
func (ref *Ed25519EncodedGroupElement) GetAffineX() (*Ed25519FieldElement, error) {

	y, err := ref.GetAffineY()
	if err != nil {
		return nil, err
	}
	ySquare := y.square()
	// u = y^2 - 1
	u := ySquare.subtract(Ed25519Field.ONE)
	// v = d * y^2 + 1
	v := ySquare.multiply(Ed25519Field.D).add(Ed25519Field.ONE)
	// x = sqrt(u / v)
	x := Ed25519FieldElementSqrt(u, v)
	vxSquare := x.square().multiply(v)
	checkForZero := vxSquare.subtract(u)
	if checkForZero.IsNonZero() {
		checkForZero = vxSquare.add(u)
		if checkForZero.IsNonZero() {
			return nil, errors.New("not a valid Ed25519EncodedGroupElement.")
		}

		x = x.multiply(Ed25519Field.I)
	}

	if x.IsNegative() != utils.GetBitToBool(ref.Raw, 255) {
		x = x.negate()
	}

	return &x, nil
}

/**
 * Gets the affine y-coordinate.
 *
 * @return the affine y-coordinate.
 */func (ref *Ed25519EncodedGroupElement) GetAffineY() (*Ed25519FieldElement, error) {

	// The affine y-coordinate is in bits 0 to 254.
	// Since the decode() method of Ed25519EncodedFieldElement ignores bit 255,
	// we can use that method without problems.
	encoded, err := NewEd25519EncodedFieldElement(ref.Raw)
	if err != nil {
		return nil, err
	}
	return encoded.Decode(), nil
}
func (ref *Ed25519EncodedGroupElement) Equals(ge *Ed25519EncodedGroupElement) bool {

	return isEqualConstantTime(ref.Raw, ge.Raw)
}
func (ref *Ed25519EncodedGroupElement) String() string {

	x, err := ref.GetAffineX()
	if err != nil {
		return err.Error()
	}
	y, err := ref.GetAffineY()
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf(
		"x=%s\ny=%s\n", x.String(), y.String())
}

//CoordinateSystem Available coordinate systems for a group element.
type CoordinateSystem int

const (
	// Affine coordinate system (x, y).
	AFFINE CoordinateSystem = iota
	// Projective coordinate system (X:Y:Z) satisfying x=X/Z, y=Y/Z.
	P2
	// Extended projective coordinate system (X:Y:Z:T) satisfying x=X/Z, y=Y/Z, XY=ZT.
	P3
	// Completed coordinate system ((X:Z), (Y:T)) satisfying x=X/Z, y=Y/T.
	P1xP1
	// Precomputed coordinate system (y+x, y-x, 2dxy).
	PRECOMPUTED
	// Cached coordinate system (Y+X, Y-X, Z, 2dT).
	CACHED
)

//ed25519Group Represents the underlying group for Ed25519.
type ed25519Group struct {
	GROUP_ORDER *big.Int
	basePoint   *Ed25519GroupElement
}

// different representations of zero

func (ref *ed25519Group) BASE_POINT() *Ed25519GroupElement {
	return ref.basePoint.copy()
}
func (ref *ed25519Group) ZERO_P2() *Ed25519GroupElement {
	return NewEd25519GroupElementP2(Ed25519Field_ZERO(), Ed25519Field_ONE(), Ed25519Field_ONE())
}
func (ref *ed25519Group) ZERO_P3() *Ed25519GroupElement {
	return NewEd25519GroupElementP3(Ed25519Field_ZERO(), Ed25519Field_ONE(), Ed25519Field_ONE(), Ed25519Field_ZERO())
}
func (ref *ed25519Group) ZERO_PRECOMPUTED() *Ed25519GroupElement {
	return NewEd25519GroupElementPrecomputed(Ed25519Field_ONE(), Ed25519Field_ONE(), Ed25519Field_ZERO())
}

//Ed25519Groupp
var Ed25519Group = &ed25519Group{}

const ed25519GroupOrder = "27742317777372353535851937790883648493"
const ed25519GroupRawElement = "b0a00e4a271beec478e42fad0618432fa7d7fb3d99004d2b0bdfc14f8024832b"

func init() {

	//Ed25519Group.GROUP_ORDER
	// 2^252 + 27742317777372353535851937790883648493

	rInt, ok := (&big.Int{}).SetString(ed25519GroupOrder, 10)
	if !ok {
		panic(errors.New(ed25519GroupOrder + " is wrang value for big.Int!"))
	}

	z := (&big.Int{}).Lsh(BigInteger_ONE(), 252)
	Ed25519Group.GROUP_ORDER = z.Add(z, rInt)

	/**
	 * <pre>{@code
	 * (x, 4/5); x > 0
	 * }</pre>
	 */
	var err error
	Ed25519Group.basePoint, err = getBasePoint()
	if err != nil {
		panic(err)
	}
}

const ed25519GroupBasePoint = "5866666666666666666666666666666666666666666666666666666666666666"

func getBasePoint() (*Ed25519GroupElement, error) {

	rawEncodedGroupElement, err := hex.DecodeString(ed25519GroupBasePoint)
	if err != nil {
		return nil, err
	}
	grElem, err := NewEd25519EncodedGroupElement(rawEncodedGroupElement)
	if err != nil {
		return nil, err
	}
	basePoint, err := grElem.Decode()
	if err != nil {
		return nil, err
	}
	basePoint.precomputedForSingle = basePrecSingle
	basePoint.precomputedForDouble = basePrecDouble

	return basePoint, nil
}

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
//Ed25519GroupElement
type Ed25519GroupElement struct {
	coordinateSystem CoordinateSystem

	X *Ed25519FieldElement

	Y *Ed25519FieldElement

	Z *Ed25519FieldElement

	T *Ed25519FieldElement
	// Precomputed table for a single scalar multiplication.
	precomputedForSingle [][]*Ed25519GroupElement
	// Precomputed table for a float64 scalar multiplication
	precomputedForDouble []*Ed25519GroupElement
	//region constructors
}

//NewEd25519GroupElement creates a group element for a curve.
func NewEd25519GroupElement(coordinateSystem CoordinateSystem,
	x *Ed25519FieldElement,
	y *Ed25519FieldElement,
	z *Ed25519FieldElement,
	t *Ed25519FieldElement) *Ed25519GroupElement {
	return &Ed25519GroupElement{
		coordinateSystem,
		x,
		y,
		z,
		t,
		nil,
		nil,
	}
}

//NewEd25519GroupElementAffine creates a new group element using the AFFINE coordinate system.
func NewEd25519GroupElementAffine(
	x *Ed25519FieldElement,
	y *Ed25519FieldElement,
	z *Ed25519FieldElement) *Ed25519GroupElement {
	return NewEd25519GroupElement(AFFINE, x, y, z, nil)
}

//NewEd25519GroupElementP2 creates a new group element using the P2 coordinate system.
func NewEd25519GroupElementP2(
	x *Ed25519FieldElement,
	y *Ed25519FieldElement,
	z *Ed25519FieldElement) *Ed25519GroupElement {
	return NewEd25519GroupElement(P2, x, y, z, nil)
}

//NewEd25519GroupElementP3 creates a new group element using the P3 coordinate system.
func NewEd25519GroupElementP3(
	x *Ed25519FieldElement,
	y *Ed25519FieldElement,
	z *Ed25519FieldElement,
	t *Ed25519FieldElement) *Ed25519GroupElement {
	return NewEd25519GroupElement(P3, x, y, z, t)
}

//NewEd25519GroupElementP1XP1 Creates a new group element using the P1xP1 coordinate system.
func NewEd25519GroupElementP1XP1(
	x *Ed25519FieldElement,
	y *Ed25519FieldElement,
	z *Ed25519FieldElement,
	t *Ed25519FieldElement) *Ed25519GroupElement {
	return NewEd25519GroupElement(P1xP1, x, y, z, t)
}

//NewEd25519GroupElementPrecomputed сreates a new group element using the PRECOMPUTED coordinate system.
//(CoordinateSystem.PRECOMPUTED, yPlusx, yMinusx, xy2d, nil)
func NewEd25519GroupElementPrecomputed(
	x *Ed25519FieldElement,
	y *Ed25519FieldElement,
	z *Ed25519FieldElement) *Ed25519GroupElement {
	return NewEd25519GroupElement(PRECOMPUTED, x, y, z, nil)
}

// NewEd25519GroupElementCached сreates a new group element using the CACHED coordinate system.
//(CoordinateSystem.CACHED, YPlusX, YMinusX, Z, T2d)
func NewEd25519GroupElementCached(
	x *Ed25519FieldElement,
	y *Ed25519FieldElement,
	z *Ed25519FieldElement,
	t *Ed25519FieldElement) *Ed25519GroupElement {
	return NewEd25519GroupElement(CACHED, x, y, z, t)
}
func (ref *Ed25519GroupElement) Equals(ge *Ed25519GroupElement) (res bool) {

	if ref.coordinateSystem != ge.coordinateSystem {
		defer func() {
			err := recover()
			if err != nil {
				fmt.Printf("%v, %d, %d", err, ref.coordinateSystem, ge.coordinateSystem)
				res = false
			}
		}()
		ge = ge.toCoordinateSystem(ref.coordinateSystem)
	}

	switch ref.coordinateSystem {
	case P2, P3:
		if ref.Z.Equals(ge.Z) {
			return ref.X.Equals(ge.X) && ref.Y.Equals(ge.Y)
		}

		x1 := ref.X.multiply(*(ge.Z))
		y1 := ref.Y.multiply(*(ge.Z))
		x2 := ge.X.multiply(*(ref.Z))
		y2 := ge.Y.multiply(*(ref.Z))
		return x1.Equals(&x2) && y1.Equals(&y2)
	case P1xP1:
		return ref.toP2().Equals(ge)
	case PRECOMPUTED:
		return ref.X.Equals(ge.X) && ref.Y.Equals(ge.Y) && ref.Z.Equals(ge.Z)
	case CACHED:
		if ref.Z.Equals(ge.Z) {
			return ref.X.Equals(ge.X) && ref.Y.Equals(ge.Y) && ref.T.Equals(ge.T)
		}

		x3 := ref.X.multiply(*(ge.Z))
		y3 := ref.Y.multiply(*(ge.Z))
		t3 := ref.T.multiply(*(ge.Z))
		x4 := ge.X.multiply(*(ref.Z))
		y4 := ge.Y.multiply(*(ref.Z))
		t4 := ge.T.multiply(*(ref.Z))
		return x3.Equals(&x4) && y3.Equals(&y4) && t3.Equals(&t4)
	default:
		panic(errors.New("not supportet coor type"))
	}
	return false

}

// Convert a to 2^16 bit representation.
// *
// * @param encoded The encode field element.
// * @return 64 bytes, each between -8 and 7
func (ref *Ed25519GroupElement) toRadix16(encoded *Ed25519EncodedFieldElement) []int8 {

	a := encoded.Raw
	e := make([]int8, 64)
	for i := 0; i < 32; i++ {
		e[2*i] = int8(a[i]) & 15
		e[2*i+1] = (int8(a[i]) >> 4) & 15
	}

	/* each e[i] is between 0 and 15 */
	/* e[63] is between 0 and 7 */
	carry := 0
	for i := 0; i < 63; i++ {
		e[i] += int8(carry)
		carry = int(e[i]) + 8
		carry >>= 4
		e[i] -= int8(carry) << 4
	}

	e[63] += int8(carry)
	return e
}

/* each e[i] is between 0 and 15 */
/* e[63] is between 0 and 7 */
//int carry = 0;
//for (i = 0; i < 63; i++) {
//e[i] += carry;
//carry = e[i] + 8;
//carry >>= 4;
//e[i] -= carry << 4;
//}
//e[63] += carry;

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
 */func (ref *Ed25519GroupElement) slide(encoded *Ed25519EncodedFieldElement) []int8 {

	a := encoded.Raw
	r := make([]int8, 256)
	// Put each bit of 'a' into a separate byte, 0 or 1
	for i := 0; i < 256; i++ {
		r[i] = 1 & (int8(a[i>>3]) >> uint(i&7))
	}
	//todo: algorimt must be simple!
	// Note: r[i] will always be odd.
	for i := uint(0); i < 256; i++ {
		if r[i] != 0 {
			for b := uint(1); b <= 6 && i+b < 256; b++ {
				// Accumulate bits if possible
				ib := uint(i + b)
				if r[ib] != 0 {
					if r[i]+(r[ib]<<b) <= 15 {
						r[i] += r[ib] << b
						r[ib] = 0
					} else if r[i]-(r[ib]<<b) >= -15 {
						r[i] -= r[ib] << b
						for k := ib; k < 256; k++ {
							if r[k] == 0 {
								r[k] = 1
								break
							}

							r[k] = 0
						}

					} else {
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
 */func (ref *Ed25519GroupElement) GetCoordinateSystem() CoordinateSystem {

	return ref.coordinateSystem
}

/**
 * Gets the X value of the group element.
 * This is for most coordinate systems the projective X coordinate.
 *
 * @return The X value.
 */func (ref *Ed25519GroupElement) GetX() *Ed25519FieldElement {

	return ref.X
}

/**
 * Gets the Y value of the group element.
 * This is for most coordinate systems the projective Y coordinate.
 *
 * @return The Y value.
 */func (ref *Ed25519GroupElement) GetY() *Ed25519FieldElement {

	return ref.Y
}

/**
 * Gets the Z value of the group element.
 * This is for most coordinate systems the projective Z coordinate.
 *
 * @return The Z value.
 */func (ref *Ed25519GroupElement) GetZ() *Ed25519FieldElement {

	return ref.Z
}

/**
 * Gets the T value of the group element.
 * This is for most coordinate systems the projective T coordinate.
 *
 * @return The T value.
 */func (ref *Ed25519GroupElement) GetT() *Ed25519FieldElement {

	return ref.T
}

/**
 * Gets a value indicating whether or not the group element has a
 * precomputed table for float64 scalar multiplication.
 *
 * @return true if it has the table, false otherwise.
 */
func (ref *Ed25519GroupElement) IsPrecomputedForDoubleScalarMultiplication() bool {

	return nil != ref.precomputedForDouble
}

// Converts the group element to an encoded point on the curve.
// *
// * @return The encoded point as byte array.
func (ref *Ed25519GroupElement) Encode() (*Ed25519EncodedGroupElement, error) {

	switch ref.coordinateSystem {
	case P2, P3:
		inverse := ref.Z.invert()
		x := ref.X.multiply(inverse)
		y := ref.Y.multiply(inverse)
		s := y.Encode().Raw
		if x.IsNegative() {
			s[len(s)-1] |= 0x80
		} else {
			s[len(s)-1] |= 0
		}
		return NewEd25519EncodedGroupElement(s)
	}
	return ref.toP2().Encode()

}

// Converts the group element to the P2 coordinate system.
func (ref *Ed25519GroupElement) toP2() *Ed25519GroupElement {

	return ref.toCoordinateSystem(P2)
}

// Converts the group element to the P3 coordinate system.
func (ref *Ed25519GroupElement) toP3() *Ed25519GroupElement {

	return ref.toCoordinateSystem(P3)
}

// Converts the group element to the CACHED coordinate system.
func (ref *Ed25519GroupElement) toCached() *Ed25519GroupElement {

	return ref.toCoordinateSystem(CACHED)
}

// toCoordinateSystem convert a Ed25519GroupElement from one coordinate system to another.
//* Supported conversions:
//* - P3 -> P2
//* - P3 -> CACHED (1 multiply, 1 add, 1 subtract)
//* - P1xP1 -> P2 (3 multiply)
//* - P1xP1 -> P3 (4 multiply)
//*
//* @param newCoordinateSystem The coordinate system to convert to.
//* @return A new group element in the new coordinate system.
func (ref *Ed25519GroupElement) toCoordinateSystem(newCoordinateSystem CoordinateSystem) *Ed25519GroupElement {
	switch ref.coordinateSystem {
	case P2:
		switch newCoordinateSystem {
		case P2:
			return NewEd25519GroupElementP2(ref.X, ref.Y, ref.Z)
		default:
			panic("NewIllegalArgumentException P2")
		}

	case P3:
		switch newCoordinateSystem {
		case P2:
			return NewEd25519GroupElementP2(ref.X, ref.Y, ref.Z)
		case P3:
			return NewEd25519GroupElementP3(ref.X, ref.Y, ref.Z, ref.T)
		case CACHED:
			X := ref.Y.add(*(ref.X))
			Y := ref.Y.subtract(*(ref.X))
			t := ref.T.multiply(Ed25519Field.D_Times_TWO)
			return NewEd25519GroupElementCached(&X, &Y, ref.Z, &t)
		default:
			panic("NewIllegalArgumentException P3")
		}

	case P1xP1:
		switch newCoordinateSystem {
		case P2:
			x := ref.X.multiply(*(ref.T))
			y := ref.Y.multiply(*(ref.Z))
			z := ref.Z.multiply(*(ref.T))
			return NewEd25519GroupElementP2(&x, &y, &z)
		case P3:
			x := ref.X.multiply(*(ref.T))
			y := ref.Y.multiply(*(ref.Z))
			z := ref.Z.multiply(*(ref.T))
			t := ref.X.multiply(*(ref.Y))
			return NewEd25519GroupElementP3(&x, &y, &z, &t)
		case P1xP1:
			return NewEd25519GroupElementP1XP1(ref.X, ref.Y, ref.Z, ref.T)
		default:
			panic("NewIllegalArgumentException P1xP1")
		}

	case PRECOMPUTED:
		switch newCoordinateSystem {
		case PRECOMPUTED:
			//noinspection SuspiciousNameCombination
			return NewEd25519GroupElementPrecomputed(ref.X, ref.Y, ref.Z)
		default:
			panic("NewIllegalArgumentException PRECOMPUTED")
		}

	case CACHED:
		switch newCoordinateSystem {
		case CACHED:
			return NewEd25519GroupElementCached(ref.X, ref.Y, ref.Z, ref.T)
		default:
			panic("NewIllegalArgumentException CACHED")
		}

	default:
		panic(fmt.Sprintf("NewIllegalArgumentException(%d)", ref.coordinateSystem))
	}
	return nil
}

// make copy ref object
func (ref *Ed25519GroupElement) copy() *Ed25519GroupElement {
	el := &Ed25519GroupElement{
		ref.coordinateSystem,
		ref.X,
		ref.Y,
		ref.Z,
		ref.T,
		make([][]*Ed25519GroupElement, len(ref.precomputedForSingle)),
		make([]*Ed25519GroupElement, len(ref.precomputedForDouble)),
	}
	for i, val := range ref.precomputedForSingle {
		el.precomputedForSingle[i] = make([]*Ed25519GroupElement, 8)
		for j, valElem := range val {
			el.precomputedForSingle[i][j] = valElem
		}
	}

	for i, val := range ref.precomputedForDouble {
		el.precomputedForDouble[i] = val
	}

	return el
}

// PrecomputeForScalarMultiplication precomputes the group elements needed to speed up a scalar multiplication.
func (ref *Ed25519GroupElement) PrecomputeForScalarMultiplication() {

	if (ref.precomputedForSingle != nil) || len(ref.precomputedForSingle) > 0 {
		return
	}

	Bi := ref.copy()
	ref.precomputedForSingle = make([][]*Ed25519GroupElement, 32)
	for i := range ref.precomputedForSingle {
		ref.precomputedForSingle[i] = make([]*Ed25519GroupElement, 8)
		Bij := Bi.copy()
		for j := range ref.precomputedForSingle[i] {
			inverse := Bij.Z.invert()
			x := Bij.X.multiply(inverse)
			y := Bij.Y.multiply(inverse)

			X := y.add(x)
			Y := y.subtract(x)
			Z := x.multiply(y).multiply(Ed25519Field.D_Times_TWO)
			ref.precomputedForSingle[i][j] = NewEd25519GroupElementPrecomputed(&X, &Y, &Z)
			Bij = Bij.add(Bi.toCached()).toP3()
		}
		// Only every second summand is precomputed (16^2 = 256).
		for k := 0; k < 8; k++ {
			Bi = Bi.add(Bi.toCached()).toP3()
		}

	}

}

//PrecomputeForDoubleScalarMultiplication Precomputes the group elements used to speed up a float64 scalar multiplication.
func (ref *Ed25519GroupElement) PrecomputeForDoubleScalarMultiplication() {

	if len(ref.precomputedForDouble) > 0 {
		return
	}

	ref.precomputedForDouble = make([]*Ed25519GroupElement, 8)
	Bi := ref.copy()
	for i := range ref.precomputedForDouble {
		inverse := Bi.Z.invert()
		x := Bi.X.multiply(inverse)
		y := Bi.Y.multiply(inverse)

		X := y.add(x)
		Y := y.subtract(x)
		Z := x.multiply(y).multiply(Ed25519Field.D_Times_TWO)
		ref.precomputedForDouble[i] = NewEd25519GroupElementPrecomputed(&X, &Y, &Z)
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
 */func (ref *Ed25519GroupElement) dbl() *Ed25519GroupElement {

	switch ref.coordinateSystem {
	case P2, P3:
		XSquare := ref.X.square()
		YSquare := ref.Y.square()
		B := ref.Z.squareAndDouble()
		A := ref.X.add(*(ref.Y))
		ASquare := A.square()
		YSquarePlusXSquare := YSquare.add(XSquare)
		YSquareMinusXSquare := YSquare.subtract(XSquare)
		X := ASquare.subtract(YSquarePlusXSquare)
		T := B.subtract(YSquareMinusXSquare)
		return NewEd25519GroupElementP1XP1(&X, &YSquarePlusXSquare, &YSquareMinusXSquare, &T)
	default:
		panic("NewUnsupportedOperationException()")
	}

}

/**
 * Ed25519GroupElement addition using the twisted Edwards addition law for extended coordinates.
 * ref must be given in P^3 coordinate system and g in PRECOMPUTED coordinate system.
 * r = ref + g where ref = (X1 : Y1 : Z1 : T1), g = (g.X, g.Y, g.Z) = (Y2/Z2 + X2/Z2, Y2/Z2 - X2/Z2, 2 * d * X2/Z2 * Y2/Z2)
 * r in P x P coordinate system:
 * r = ((X' : Z'), (Y' : T')) where
 * X' = (Y1 + X1) * g.X - (Y1 - X1) * q.Y = ((Y1 + X1) * (Y2 + X2) - (Y1 - X1) * (Y2 - X2)) * 1/Z2
 * Y' = (Y1 + X1) * g.X + (Y1 - X1) * q.Y = ((Y1 + X1) * (Y2 + X2) + (Y1 - X1) * (Y2 - X2)) * 1/Z2
 * Z' = 2 * Z1 + T1 * g.Z = 2 * Z1 + T1 * 2 * d * X2 * Y2 * 1/Z2^2 = (2 * Z1 * Z2 + 2 * d * T1 * T2) * 1/Z2
 * T' = 2 * Z1 - T1 * g.Z = 2 * Z1 - T1 * 2 * d * X2 * Y2 * 1/Z2^2 = (2 * Z1 * Z2 - 2 * d * T1 * T2) * 1/Z2
 * <br>
 * Formula for the P x P coordinate system is in agreement with the formula given in
 * file ge25519.c method add_p1p1() in this implementation.
 * Setting A = (Y1 - X1) * (Y2 - X2), B = (Y1 + X1) * (Y2 + X2), C = 2 * d * T1 * T2, D = 2 * Z1 * Z2 we get
 * X' = (B - A) * 1/Z2
 * Y' = (B + A) * 1/Z2
 * Z' = (D + C) * 1/Z2
 * T' = (D - C) * 1/Z2
 * r converted from P x P to P^2 coordinate system:
 * r = (X'' : Y'' : Z'' : T'') where
 * X'' = X' * T' = (B - A) * (D - C) * 1/Z2^2
 * Y'' = Y' * Z' = (B + A) * (D + C) * 1/Z2^2
 * Z'' = Z' * T' = (D + C) * (D - C) * 1/Z2^2
 * T'' = X' * Y' = (B - A) * (B + A) * 1/Z2^2
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
 */func (ref *Ed25519GroupElement) precomputedAdd(g *Ed25519GroupElement) (*Ed25519GroupElement, error) {
	if ref.coordinateSystem != P3 {
		return nil, errors.New("NewUnsupportedOperationException(")
	}

	if g.coordinateSystem != PRECOMPUTED {
		return nil, errors.New("NewIllegalArgumentException(")
	}

	YPlusX := ref.Y.add(*(ref.X))
	YMinusX := ref.Y.subtract(*(ref.X))
	//A = (Y1 - X1) * (Y2 - X2)
	A := YMinusX.multiply(*(g.Y))

	//B = (Y1 + X1) * (Y2 + X2)
	B := YPlusX.multiply(*(g.X))

	//C = 2 * d * T1 * T2
	C := ref.T.multiply(*(g.Z))

	//D = 2 * Z1 * Z2
	D := ref.Z.add(*(ref.Z))

	//X' = (Y1 + X1) * g.X - (Y1 - X1) * q.Y
	x := B.subtract(A)
	//(Y1 + X1) * g.X + (Y1 - X1) * q.Y
	y := A.add(B)
	//2 * Z1 + T1 * g.Z
	z := D.add(C)
	t := D.subtract(C)

	return NewEd25519GroupElementP1XP1(&x, &y, &z, &t), nil
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
 */func (ref Ed25519GroupElement) precomputedSubtract(g *Ed25519GroupElement) *Ed25519GroupElement {
	if ref.coordinateSystem != P3 {
		panic("NewUnsupportedOperationException")
	}

	if g.coordinateSystem != PRECOMPUTED {
		panic("NewIllegalArgumentException")
	}

	YPlusX := ref.Y.add(*(ref.X))
	YMinusX := ref.Y.subtract(*(ref.X))
	A := YPlusX.multiply(*(g.Y))
	B := YMinusX.multiply(*(g.X))
	C := g.Z.multiply(*(ref.T))
	D := ref.Z.add(*(ref.Z))

	x := A.subtract(B)
	y := A.add(B)
	z := D.subtract(C)
	t := D.add(C)

	return NewEd25519GroupElementP1XP1(&x, &y, &z, &t)
}

// add Ed25519GroupElement addition using the twisted Edwards addition law for extended coordinates.
// * ref must be given in P^3 coordinate system and g in CACHED coordinate system.
// * r = ref + g where ref = (X1 : Y1 : Z1 : T1), g = (g.X, g.Y, g.Z, g.T) = (Y2 + X2, Y2 - X2, Z2, 2 * d * T2)
// * <br>
// * r in P x P coordinate system.:
// * X' = (Y1 + X1) * (Y2 + X2) - (Y1 - X1) * (Y2 - X2)
// * Y' = (Y1 + X1) * (Y2 + X2) + (Y1 - X1) * (Y2 - X2)
// * Z' = 2 * Z1 * Z2 + 2 * d * T1 * T2
// * T' = 2 * Z1 * T2 - 2 * d * T1 * T2
// * <br>
// * Setting A = (Y1 - X1) * (Y2 - X2), B = (Y1 + X1) * (Y2 + X2), C = 2 * d * T1 * T2, D = 2 * Z1 * Z2 we get
// * X' = (B - A)
// * Y' = (B + A)
// * Z' = (D + C)
// * T' = (D - C)
// * <br>
// * Same result as in precomputedAdd() (up to a common factor which does not matter).
// *
// * @return The result in the P x P coordinate system.
func (ref *Ed25519GroupElement) add(g *Ed25519GroupElement) *Ed25519GroupElement {
	if ref.coordinateSystem != P3 {
		panic("NewUnsupportedOperationException")
	}

	if g.coordinateSystem != CACHED {
		panic("NewIllegalArgumentException")
	}

	//!!!B = (Y1 + X1) * (Y2 + X2)
	B := ref.Y.add(*(ref.X)).multiply(*(g.X))

	//!!!A = (Y1 - X1) * (Y2 - X2)
	A := ref.Y.subtract(*(ref.X)).multiply(*(g.Y))

	//C = (2 * d* T2) * T1
	C := g.T.multiply(*(ref.T))

	//D = 2 * Z1 * Z2
	ZSquare := ref.Z.multiply(*(g.Z))
	D := ZSquare.add(ZSquare)

	// !!! X' = (B - A)
	x := B.subtract(A)
	y := A.add(B)
	z := D.add(C)
	t := D.subtract(C)

	return NewEd25519GroupElementP1XP1(&x, &y, &z, &t)
}

/**
 * Ed25519GroupElement subtraction using the twisted Edwards addition law for extended coordinates.
 * <br>
 * Negating g means negating the value of the coordinate X2 and T2.
 * The formula is in accordance to the above addition.
 *
 * @param g The group element to subtract.
 * @return The result in the P x P coordinate system.
 */func (ref *Ed25519GroupElement) subtract(g *Ed25519GroupElement) *Ed25519GroupElement {
	if ref.coordinateSystem != P3 {
		panic("NewUnsupportedOperationException")
	}

	if g.coordinateSystem != CACHED {
		panic("NewIllegalArgumentException")
	}

	YPlusX := ref.Y.add(*(ref.X))
	YMinusX := ref.Y.subtract(*(ref.X))
	A := YPlusX.multiply(*(g.Y))
	B := YMinusX.multiply(*(g.X))
	C := g.T.multiply(*(ref.T))
	ZSquare := ref.Z.multiply(*(g.Z))
	D := ZSquare.add(ZSquare)

	x := A.subtract(B)
	y := A.add(B)
	z := D.subtract(C)
	t := D.add(C)

	return NewEd25519GroupElementP1XP1(&x, &y, &z, &t)
}

// negate Negates ref group element by subtracting it from the neutral group element.
// * (only used in MathUtils so it doesn't have to be fast)
func (ref *Ed25519GroupElement) negate() (*Ed25519GroupElement, error) {

	if ref.coordinateSystem != P3 {
		return nil, errors.New("NewUnsupportedOperationException")
	}

	return Ed25519Group.ZERO_P3().subtract(ref.toCached()).toP3(), nil
}

// Constant-time conditional move.
// * @param u The group element to return if b == 1.
// * @param b in {0, 1}
// * @return u if b == 1; ref if b == 0; nil otherwise.
func (ref *Ed25519GroupElement) cmov(u *Ed25519GroupElement, b int) (*Ed25519GroupElement, error) {

	if b == 1 {
		return u, nil
	} else if b == 0 {
		return ref, nil
	}
	return nil, errors.New("parameter 'b' must by in range {0,1}")
}

/**
 * Look up 16^i r_i B in the precomputed table.
 * No secret array indices, no secret branching.
 * Constant time.
 * Must have previously precomputed.
 *
 * @param pos = i/2 for i in {0, 2, 4,..., 62}
 * @param b   = r_i
 * @return The Ed25519GroupElement
 */func (ref *Ed25519GroupElement) Select(pos int, b int) (*Ed25519GroupElement, error) {
	// Is r_i negative?
	bNegative := isNegativeConstantTime(b)

	// |r_i|
	bAbs := b - (((-bNegative) & b) << 1)
	// 16^i |r_i| B
	t := Ed25519Group.ZERO_PRECOMPUTED()
	for i, el := range ref.precomputedForSingle[pos] {
		tt, err := t.cmov(el, IsConstantTimeByteEq(bAbs, i+1))
		if err != nil {
			return nil, err
		} else if tt == nil {
			return nil, errors.New("get nil object - don't work ")
		}
		t = tt
	}
	// -16^i |r_i| B
	//noinspection SuspiciousNameCombination
	z := t.Z.negate()
	tMinus := NewEd25519GroupElementPrecomputed(t.Y, t.X, &z)
	// 16^i r_i B
	return t.cmov(tMinus, bNegative)
}

/**
 * h = a * B where a = a[0]+256*a[1]+...+256^31 a[31] and
 * B is ref point. If its lookup table has not been precomputed, it
 * will be at the start of the method (and cached for later calls).
 * Constant time.
 * @param a The encoded field element.
 * @return The resulting group element.
 */
func (ref *Ed25519GroupElement) scalarMultiply(a *Ed25519EncodedFieldElement) (*Ed25519GroupElement, error) {

	e := ref.toRadix16(a)
	h := Ed25519Group.ZERO_P3()
	for i := 1; i < 64; i += 2 {
		g, err := ref.Select(i/2, int(e[i]))
		if err != nil {
			return nil, err
		}
		h, err = h.precomputedAdd(g)
		if err != nil {
			return nil, err
		}
		h = h.toP3()
	}

	h = h.dbl().toP2().dbl().toP2().dbl().toP2().dbl().toP3()
	for i := 0; i < 64; i += 2 {
		g, err := ref.Select(i/2, int(e[i]))
		if err != nil {
			return nil, err
		}
		h, err = h.precomputedAdd(g)
		if err != nil {
			return nil, err
		}
		h = h.toP3()
	}

	return h, nil
}

//doubleScalarMultiplyVariableTime r = b * B - a * A  where
// * a and b are encoded field elements and
// * B is ref point.
// * A must have been previously precomputed for float64 scalar multiplication.
// *
// * @param A in P3 coordinate system.
// * @param a = The first encoded field element.
// * @param b = The second encoded field element.
// * @return The resulting group element.
func (ref *Ed25519GroupElement) doubleScalarMultiplyVariableTime(
	A *Ed25519GroupElement,
	a *Ed25519EncodedFieldElement,
	b *Ed25519EncodedFieldElement) (r *Ed25519GroupElement, err error) {
	aSlide := ref.slide(a)
	bSlide := ref.slide(b)

	r = Ed25519Group.ZERO_P2()
	flag := false
	for i := 255; i >= 0; i-- {
		flag = flag || (aSlide[i] != 0) || (bSlide[i] != 0)
		if flag {

			t := r.dbl()
			if aSlide[i] > 0 {
				t = t.toP3().precomputedSubtract(A.precomputedForDouble[aSlide[i]/2])
			} else if aSlide[i] < 0 {
				t, err = t.toP3().precomputedAdd(A.precomputedForDouble[(-aSlide[i])/2])
				if err != nil {
					return nil, err
				}
			}

			if bSlide[i] > 0 {
				t, err = t.toP3().precomputedAdd(ref.precomputedForDouble[bSlide[i]/2])
				if err != nil {
					return nil, err
				}
			} else if bSlide[i] < 0 {
				t = t.toP3().precomputedSubtract(ref.precomputedForDouble[(-bSlide[i])/2])
			}

			r = t.toP2()
		}

	}

	return r, nil
}

// SatisfiesCurveEquation Verify that the group element satisfies the curve equation.
//* @return true if the group element satisfies the curve equation, false otherwise.
func (ref *Ed25519GroupElement) SatisfiesCurveEquation() bool {

	switch ref.coordinateSystem {
	case P2, P3:
		inverse := ref.Z.invert()
		x := ref.X.multiply(inverse)
		y := ref.Y.multiply(inverse)
		xSquare := x.square()
		ySquare := y.square()
		dXSquareYSquare := Ed25519Field.D.multiply(xSquare).multiply(ySquare)
		return bytes.Equal(Ed25519Field_ONE().add(dXSquareYSquare).add(xSquare).Encode().Raw, ySquare.Encode().Raw)
	}
	return ref.toP2().SatisfiesCurveEquation()

}

func (ref *Ed25519GroupElement) String() string {

	return fmt.Sprintf(
		"X=%s\nY=%s\nZ=%s\nT=%s\n",
		ref.X.String(),
		ref.Y.String(),
		ref.Z.String(),
		ref.T.String())
}
