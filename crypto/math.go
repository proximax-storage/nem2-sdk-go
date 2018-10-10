// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.


package crypto

import (
	rand2 "crypto/rand"
	"encoding/binary"
	"errors"
	"github.com/proximax-storage/nem2-sdk-go/utils"
	"io"
	"math/big"
)

//src/test/java/io/nem/core/crypto/ed25519/arithmetic/MathUtils.java
func BigInteger_ZERO() *big.Int {
	return big.NewInt(0)
}
func BigInteger_ONE() *big.Int {
	return big.NewInt(1)
}

//MathUtils Utility class to help with calculations.
type mathUtils struct {
	EXPONENTS []uint
	D         *big.Int
}

func (ref *mathUtils) random() (random [32]byte) {
	rand := rand2.Reader
	_, err := io.ReadFull(rand, random[:])
	if err != nil {
		panic(err)
	}

	return
}

//* Converts a 2^8 bit representation to a BigInteger.
//* Value: bytes[0] + 2^8 * bytes[1] + ...
//*
//* @param bytes The 2^8 bit representation.
//* @return The BigInteger.
func (ref *mathUtils) BytesToBigInteger(bytes []byte) *big.Int {

	//return utils.BytesToBigInteger(bytes)
	b := BigInteger_ZERO()
	for i, val := range bytes {
		el := (&big.Int{}).SetUint64(uint64(uint8(val)) & 0xff)
		//one := BigInteger_ONE()
		//one = one.Mul(one, el)
		b = b.Add(b, el.Lsh(el, uint(i*8)))
	}

	return b
}

/**
 * Converts an encoded field element to a *big.Int.
 *
 * @param encoded The encoded field element.
 * @return The *big.Int.
 */
func (ref *mathUtils) EncodedFieldToBigInteger(encoded *Ed25519EncodedFieldElement) *big.Int {

	return ref.BytesToBigInteger(encoded.Raw)
}

/**
 * Converts a field element to a *big.Int.
 *
 * @param f The field element.
 * @return The *big.Int.
 */
func (ref *mathUtils) FieldToBigInteger(f *Ed25519FieldElement) *big.Int {

	return ref.BytesToBigInteger(f.Encode().Raw)
}

//     * Converts a 2^25.5 bit representation to a BigInteger.
//     * Value: 2^EXPONENTS[0] * t[0] + 2^EXPONENTS[1] * t[1] + ... + 2^EXPONENTS[9] * t[9]
//     *
//     * @param t The 2^25.5 bit representation.
//     * @return The BigInteger.
func (ref *mathUtils) IntsToBigInteger(t []intRaw) *big.Int {
	b := BigInteger_ZERO()
	for i, val := range t {
		el := (&big.Int{}).SetInt64(int64(val))
		one := BigInteger_ONE()
		one = one.Mul(one, el)
		b = b.Add(b, one.Lsh(one, ref.EXPONENTS[i]))
	}

	return b
}

func (ref *mathUtils) GetRandomFieldElement() Ed25519FieldElement {

	t := make([]intRaw, 10)
	rand := rand2.Reader

	for j := range t {
		var v [2]byte
		_, err := rand.Read(v[:])
		if err != nil {
			panic(err)
		}

		t[j] = (&big.Int{}).SetBytes(v[:]).Lsh(big.NewInt(1), 25).Int64() - (1 << 24)
	}

	return Ed25519FieldElement{t}
}

// ScalarMultiplyGroupElement scalar multiply the group element by the field element.
// * @param g The group element.
// * @param f The field element.
// * @return The resulting group element.
func (ref *mathUtils) ScalarMultiplyGroupElement(g *Ed25519GroupElement, f Ed25519FieldElement) *Ed25519GroupElement {

	bytes := f.Encode().Raw
	h := Ed25519Group.ZERO_P3()
	for i := uint(255); i > 0; i-- {
		h = ref.DoubleGroupElement(h)
		if utils.GetBitToBool(bytes, i-1) {
			h = ref.AddGroupElements(h, g)
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
func (ref *mathUtils) doubleScalarMultiplyGroupElements(
	g1 *Ed25519GroupElement,
	f1 Ed25519FieldElement,
	g2 *Ed25519GroupElement,
	f2 Ed25519FieldElement) *Ed25519GroupElement {
	h1 := ref.ScalarMultiplyGroupElement(g1, f1) // Ed25519GroupElement
	h2 := ref.ScalarMultiplyGroupElement(g2, f2) // Ed25519GroupElement

	h2Neg, err := h2.negate()
	if err != nil {
		panic(err)
	}
	return ref.AddGroupElements(h1, h2Neg)
}

// ToFieldElement Converts a *big.Int to a field element.
func (ref *mathUtils) ToFieldElement(b *big.Int) *Ed25519FieldElement {

	return (&Ed25519EncodedFieldElement{Ed25519Field_ZERO_SHORT(), ref.ToByteArray(b)}).Decode()
}

/**
 * Creates a group element from a byte array.
 * Bit 0 to 254 are the affine y-coordinate, bit 255 is the sign of the affine x-coordinate.
 *
 * @param bytes the byte array.
 * @return The group element.
 */
func (ref *mathUtils) ToGroupElement(bytes []byte) (*Ed25519GroupElement, error) {

	shouldBeNegative := (bytes[31] >> 7) != 0
	bytes[31] &= 0x7f
	y := ref.BytesToBigInteger(bytes)
	x, err := ref.GetAffineXFromAffineY(y, shouldBeNegative)
	if err != nil {
		return nil, err
	}
	return NewEd25519GroupElementP3(
		ref.ToFieldElement(x),
		ref.ToFieldElement(y),
		Ed25519Field_ONE(),
		ref.ToFieldElement(x.Mul(x, y).Mod(x, Ed25519Field.P))), nil
}

var errNoValidEd25519Group = errors.New("not a valid Ed25519GroupElement")

/**
 * Gets the affine x-coordinate from a given affine y-coordinate and the sign of x.
 *
 * @param y                The affine y-coordinate
 * @param shouldBeNegative true if the negative solution should be chosen, false otherwise.
 * @return The affine x-ccordinate.
 */
func (ref *mathUtils) GetAffineXFromAffineY(y *big.Int, shouldBeNegative bool) (*big.Int, error) {

	// x = sign(x) * sqrt((y^2 - 1) / (d * y^2 + 1))
	u := (&big.Int{}).Mul(y, y)
	u.Sub(u, BigInteger_ONE()).Mod(u, Ed25519Field.P)

	v := (&big.Int{}).Mul(ref.D, y)
	v.Mul(v, y).Add(v, BigInteger_ONE()).Mod(v, Ed25519Field.P)

	x := ref.getSqrtOfFraction(u, v)

	vx2 := (&big.Int{}).Mul(v, x)
	vx2.Mul(v, x)

	if v.Sub(vx2, u).Mod(v, Ed25519Field.P).Cmp(BigInteger_ZERO()) != 0 {
		if vx2.Add(vx2, u).Mod(vx2, Ed25519Field.P).Cmp(BigInteger_ZERO()) != 0 {
			return nil, errNoValidEd25519Group
		}

		x = x.Mul(x, ref.IntsToBigInteger(Ed25519Field.I.Raw)).Mod(x, Ed25519Field.P)
	}

	isNegative := x.Mod(x, big.NewInt(2)).Cmp(BigInteger_ONE()) == 0 // final
	if (shouldBeNegative && !isNegative) || (!shouldBeNegative && isNegative) {
		x = x.Neg(x).Mod(x, Ed25519Field.P)
	}

	return x, nil
}

/**
 * Calculates and returns the square root of a fraction of u and v.
 * The sign is unpredictable.
 *
 * @param u The nominator.
 * @param v The denominator.
 * @return Plus or minus the square root
 */
func (ref *mathUtils) getSqrtOfFraction(u *big.Int, v *big.Int) *big.Int { /* private static  */

	one := BigInteger_ONE()
	three := big.NewInt(3)
	powV3 := (&big.Int{}).Exp(v, three, nil)
	s := big.NewInt(7)
	powV7 := (&big.Int{}).Exp(v, s, nil)

	x := (&big.Int{}).Mul(u, powV7)
	x.Exp(x, one.Lsh(one, 252).Sub(one, three), Ed25519Field.P).Mod(x, Ed25519Field.P)
	x.Mul(x, u).Mul(x, powV3).Mod(x, Ed25519Field.P)

	return x
}

//ToRepresentation Converts a group element from one coordinate system to another.
// * This method is a helper used to test various methods in Ed25519GroupElement.
// *
// * @param g                   The group element.
// * @param newCoordinateSystem The desired coordinate system.
// * @return The same group element in the new coordinate system.
func (ref *mathUtils) ToRepresentation(g *Ed25519GroupElement, newCoorSys CoordinateSystem) (*Ed25519GroupElement, error) {
	gX := ref.BytesToBigInteger(g.X.Encode().Raw)
	gY := ref.BytesToBigInteger(g.Y.Encode().Raw)
	gZ := ref.BytesToBigInteger(g.Z.Encode().Raw)
	var gT *big.Int
	if g.T != nil {
		gT = ref.BytesToBigInteger(g.T.Encode().Raw)
	}

	// Switch to affine coordinates.
	switch g.coordinateSystem {
	case AFFINE:
		return ref.getNeeCoor(gX, gY, newCoorSys)
	case P2, P3:
		x := XMul_YModInverseAndMod_P(gX, gZ)
		y := XMul_YModInverseAndMod_P(gY, gZ)
		return ref.getNeeCoor(
			x,
			y,
			newCoorSys)
	case P1xP1:
		if gT == nil {
			return nil, errors.New("coordinate T must not nil for P!XP1 ")
		}
		x := XMul_YModInverseAndMod_P(gX, gZ)
		y := XMul_YModInverseAndMod_P(gY, gT)
		return ref.getNeeCoor(x, y, newCoorSys)
	case CACHED:
		z := (&big.Int{}).Mul(gZ, big.NewInt(2))
		x := XMul_YModInverseAndMod_P((&big.Int{}).Sub(gX, gY), z)
		y := XMul_YModInverseAndMod_P((&big.Int{}).Add(gX, gY), z)

		return ref.getNeeCoor(x, y, newCoorSys)
	case PRECOMPUTED:
		//safaty gX for next calculation
		x := (&big.Int{}).Sub(gX, gY)
		x = x.Mul(x, gZ.Mul(gZ, big.NewInt(2))).ModInverse(x, Ed25519Field.P).Mod(x, Ed25519Field.P)

		y := gX.Add(gX, gY)
		y = y.Mul(y, big.NewInt(2)).ModInverse(y, Ed25519Field.P).Mod(y, Ed25519Field.P)
		return ref.getNeeCoor(x, y, newCoorSys)
	}
	return nil, errors.New("NewUnsupportedOperationException")
}
func XMul_YModInverseAndMod_P(x *big.Int, y *big.Int) *big.Int {
	res := &big.Int{}
	return res.Mul(x, (&big.Int{}).ModInverse(y, Ed25519Field.P)).Mod(res, Ed25519Field.P)
}
func (ref *mathUtils) getNeeCoor(x, y *big.Int, newCoorSys CoordinateSystem) (*Ed25519GroupElement, error) {

	x1 := ref.ToFieldElement(x)
	y1 := ref.ToFieldElement(y)

	// Now back to the desired coordinate system.
	switch newCoorSys {
	case AFFINE:
		return NewEd25519GroupElementAffine(x1, y1, Ed25519Field_ONE()), nil
	case P2:
		return NewEd25519GroupElementP2(x1, y1, Ed25519Field_ONE()), nil
	case P3:
		m := x.Mul(x, y)
		z := ref.ToFieldElement(m.Mod(m, Ed25519Field.P))

		return NewEd25519GroupElementP3(x1, y1, Ed25519Field_ONE(), z), nil
	case P1xP1:
		return NewEd25519GroupElementP1XP1(x1, y1, Ed25519Field_ONE(), Ed25519Field_ONE()), nil
	case CACHED:
		m := (&big.Int{}).Add(y, x)
		x1 := ref.ToFieldElement(m.Mod(m, Ed25519Field.P))

		m = m.Sub(y, x)
		y1 := ref.ToFieldElement(m.Mod(m, Ed25519Field.P))

		//safaty D for next calculation
		m = (&big.Int{}).Mul(ref.D, big.NewInt(2))
		t := ref.ToFieldElement(m.Mul(m, x).Mul(m, y).Mod(m, Ed25519Field.P))

		return NewEd25519GroupElementCached(x1, y1, Ed25519Field_ONE(), t), nil
	case PRECOMPUTED:
		m := (&big.Int{}).Add(y, x)
		x1 := ref.ToFieldElement(m.Mod(m, Ed25519Field.P))
		m = m.Sub(y, x)
		y1 := ref.ToFieldElement(m.Mod(m, Ed25519Field.P))

		m = m.Mul(ref.D, big.NewInt(2)).Mul(m, x).Mul(m, y)
		z := ref.ToFieldElement(m.Mod(m, Ed25519Field.P))

		return NewEd25519GroupElementPrecomputed(x1, y1, z), nil
	}
	return nil, errors.New("NewUnsupportedOperationException")
}
func (ref *mathUtils) GetRandomEncodedFieldElement(length int) *Ed25519EncodedFieldElement {

	bytes := ref.GetRandomByteArray(length)
	bytes[31] &= 0x7f
	zero := Ed25519Field_ZERO_SHORT()
	if length == 64 {
		zero = Ed25519Field_ZERO_LONG()
	} else if length != 32 {
		panic(errors.New("wrong lenght bytes!"))
	}
	return &Ed25519EncodedFieldElement{zero, bytes}
}

//GetRandomGroupElement Gets a random group element in P3 coordinates.
// * It's NOT guaranteed that the created group element is a multiple of the base point.
func (ref *mathUtils) GetRandomGroupElement() (el *Ed25519GroupElement) {
	err := ref.tryToCreateObject(func() (err error) {

		bytes := ref.random()
		// we have garant 32 bytes
		gr := &Ed25519EncodedGroupElement{bytes[:]}

		el, err = gr.Decode()

		return err
	})

	if err != nil {
		panic(err)
	}

	return el
}

func (ref *mathUtils) tryToCreateObject(createObject func() (err error)) error {
	const numberTryCreateObject = 100000000
	err := errors.New("init value")
	for i := 0; (err != nil) && (i < numberTryCreateObject); i++ {
		err = createObject()
	}
	return err
}

/**
 * Gets a random encoded group element.
 * It's NOT guaranteed that the created encoded group element is a multiple of the base point.
 *
 * @return The encoded group element.
 */
func (ref *mathUtils) GetRandomEncodedGroupElement() (el *Ed25519EncodedGroupElement) {

	err := ref.tryToCreateObject(func() (err error) {
		gr := ref.GetRandomGroupElement()
		el, err = gr.Encode()
		return err
	})
	if err != nil {
		panic(err)
	}

	return el
}

/**
 * Creates and returns a random byte array of given length.
 *
 * @param length The desired length.
 * @return The random byte array.
 */
func (ref *mathUtils) GetRandomByteArray(length int) []byte {

	bytes := make([]byte, length)
	rand := rand2.Reader
	_, err := io.ReadFull(rand, bytes)
	if err != nil {
		panic(err)
	}
	return bytes
}

func (ref *mathUtils) GetRandomIntRaw() intRaw {

	bytes := make([]byte, 8)
	rand := rand2.Reader
	_, err := io.ReadFull(rand, bytes)
	if err != nil {
		panic(err)
	}
	return int64(binary.LittleEndian.Uint64(bytes))
}

// Reduces an encoded field element modulo the group order and returns the result.
// *
// * @param encoded The encoded field element.
// * @return The mod group order reduced encoded field element.
func (ref *mathUtils) ReduceModGroupOrder(encoded *Ed25519EncodedFieldElement) *Ed25519EncodedFieldElement {

	b := ref.BytesToBigInteger(encoded.Raw)

	return ref.ToEncodedFieldElement(b.Mod(b, Ed25519Group.GROUP_ORDER))
}

/**
 * Converts a biginteger to an encoded field element.
 *
 * @param b The biginteger.
 * @return The encoded field element.
 */
func (ref *mathUtils) ToEncodedFieldElement(b *big.Int) *Ed25519EncodedFieldElement {

	bytes := ref.ToByteArray(b)
	zero := Ed25519Field_ZERO_SHORT()
	if lenght := len(bytes); lenght == 64 {
		zero = Ed25519Field_ZERO_LONG()
	} else if lenght != 32 {
		panic(errors.New("wrong lenght bytes!"))
	}
	return &Ed25519EncodedFieldElement{zero, bytes}
}

/**
 * Converts a biginteger to a little endian 32 byte representation.
 *
 * @param b The biginteger.
 * @return The 32 byte representation.
 */
func (ref *mathUtils) ToByteArray(b *big.Int) []byte {

	if b.Cmp(BigInteger_ONE().Lsh(BigInteger_ONE(), 256)) >= 0 {
		panic(errors.New("only numbers < 2^256 are allowed"))
	}

	bytes := make([]byte, 32)
	original := utils.BigIntToByteArray(b, 32)
	return original
	// Although b < 2^256, original can have length > 32 with some bytes set to 0.
	offset := 0
	if len(original) > 32 {
		offset = len(original) - 32
	}
	for i, b := range original[offset:] {
		bytes[len(original)-i-offset-1] = b
	}

	return bytes
}
func coorModify(g, b *big.Int) *big.Int {
	x := (&big.Int{}).Mul(g, (&big.Int{}).ModInverse(b, Ed25519Field.P))
	return x.Mod(x, Ed25519Field.P)
}

/**
 * Adds two group elements and returns the result in P3 coordinate system.
 * It uses *big.Int arithmetic and the affine coordinate system.
 * This method is a helper used to test the projective group addition formulas in Ed25519GroupElement.
 *
 * @param g1 The first group element.
 * @param g2 The second group element.
 * @return The result of the addition.
 */
func (ref *mathUtils) AddGroupElements(g1 *Ed25519GroupElement, g2 *Ed25519GroupElement) *Ed25519GroupElement {

	// Relying on a special coordinate system of the group elements.
	if g1Coor, g2Coor := g1.GetCoordinateSystem(), g2.GetCoordinateSystem(); (g1Coor != P2 && g1Coor != P3) ||
		(g2Coor != P2 && g2Coor != P3) {
		panic(errors.New("g1 and g2 must have coordinate system P2 or P3"))
	}

	// Projective coordinates
	g1X := ref.EncodedFieldToBigInteger(g1.X.Encode())
	g1Y := ref.EncodedFieldToBigInteger(g1.Y.Encode())
	g1Z := ref.EncodedFieldToBigInteger(g1.Z.Encode())
	g2X := ref.EncodedFieldToBigInteger(g2.X.Encode())
	g2Y := ref.EncodedFieldToBigInteger(g2.Y.Encode())
	g2Z := ref.EncodedFieldToBigInteger(g2.Z.Encode())
	// Affine coordinates
	g1x := coorModify(g1X, g1Z)
	g1y := coorModify(g1Y, g1Z)
	g2x := coorModify(g2X, g2Z)
	g2y := coorModify(g2Y, g2Z)
	// Addition formula for affine coordinates. The formula is complete in our case.
	//
	// (x3, y3) = (x1, y1) + (x2, y2) where
	//
	// x3 = (x1 * y2 + x2 * y1) / (1 + d * x1 * x2 * y1 * y2) and
	// y3 = (x1 * x2 + y1 * y2) / (1 - d * x1 * x2 * y1 * y2) and
	// d = -121665/121666
	d := &big.Int{}
	dx1x2y1y2 := d.Mul(ref.D, g1x).Mul(d, g2x).Mul(d, g1y).Mul(d, g2y).Mod(d, Ed25519Field.P)

	one := BigInteger_ONE()
	x3 := ref.XMulY_Plus_ZMulT_DelD(*g1x, g2y, g2x, g1y, (&big.Int{}).Add(one, dx1x2y1y2))
	y3 := ref.XMulY_Plus_ZMulT_DelD(*g1x, g2x, g1y, g2y, (&big.Int{}).Sub(one, dx1x2y1y2))
	t3 := (&big.Int{}).Mul(x3, y3)
	t3.Mod(t3, Ed25519Field.P)

	return NewEd25519GroupElementP3(
		ref.ToFieldElement(x3),
		ref.ToFieldElement(y3),
		Ed25519Field_ONE(),
		ref.ToFieldElement(t3))
}

func (ref *mathUtils) XMulY_Plus_ZMulT_DelD(x big.Int, y, z, t, d *big.Int) *big.Int {
	b := &x
	return b.Mul(b, y).Add(b, (&big.Int{}).Mul(z, t)).Mul(b, d.ModInverse(d, Ed25519Field.P)).Mod(b, Ed25519Field.P)
}

/**
 * Doubles a group element and returns the result in the P3 coordinate system.
 * It uses *big.Int arithmetic and the affine coordinate system.
 * This method is a helper used to test the projective group doubling formula in Ed25519GroupElement.
 *
 * @param g The group element.
 * @return g+g.
 */
func (ref *mathUtils) DoubleGroupElement(g *Ed25519GroupElement) *Ed25519GroupElement {

	return ref.AddGroupElements(g, g)
}

/**
* Negates a group element.
*
* @param g The group element.
* @return The negated group element.
 */
func (ref *mathUtils) NegateGroupElement(g *Ed25519GroupElement) *Ed25519GroupElement {

	if g.GetCoordinateSystem() != P3 {
		panic(errors.New("g must have coordinate system P3"))
	}

	xNegate := g.X.negate()
	tNegate := g.T.negate()
	return NewEd25519GroupElementP3(&xNegate, g.GetY(), g.GetZ(), &tNegate)
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
func (ref *mathUtils) multiplyAndAddModGroupOrder(
	a *Ed25519EncodedFieldElement,
	b *Ed25519EncodedFieldElement,
	c *Ed25519EncodedFieldElement) *Ed25519EncodedFieldElement {
	res := ref.BytesToBigInteger(a.Raw)
	res.Mul(res, ref.BytesToBigInteger(b.Raw)).Add(res, ref.BytesToBigInteger(c.Raw)).Mod(res, Ed25519Group.GROUP_ORDER)

	return ref.ToEncodedFieldElement(res)
}

/**
 * Converts a *big.Int to an encoded field element.
 *
 * @param b The *big.Int.
 * @return The encoded field element.
 */
func (ref *mathUtils) toEncodedFieldElement(b *big.Int) (*Ed25519EncodedFieldElement, error) {

	return NewEd25519EncodedFieldElement(ref.ToByteArray(b))
}

var MathUtils mathUtils

func init() {
	MathUtils.EXPONENTS = []uint{
		0,
		26,
		26 + 25,
		2*26 + 25,
		2*26 + 2*25,
		3*26 + 2*25,
		3*26 + 3*25,
		4*26 + 3*25,
		4*26 + 4*25,
		5*26 + 4*25,
	}
	x, y := big.NewInt(-121665), big.NewInt(121666)
	MathUtils.D = x.Mul(x, y.ModInverse(y, Ed25519Field.P))
}
