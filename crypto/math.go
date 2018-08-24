package crypto

import (
	rand2 "crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
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

	b := utils.BytesToBigInteger(bytes)
	//for i, val := range bytes {
	//	el := (&big.Int{}).SetUint64(uint64(uint8(val))) // & 0xff)
	//	//one := BigInteger_ONE()
	//	//one = one.Mul(one, el)
	//	b = b.Add(b, el.Lsh(el, uint(i*8)))
	//}

	return b
}

/**
 * Converts an encoded field element to a *big.Int.
 *
 * @param encoded The encoded field element.
 * @return The *big.Int.
 */
func (ref *mathUtils) EncodedFieldToBigInteger(encoded *Ed25519EncodedFieldElement) *big.Int { /* public static  */

	return ref.BytesToBigInteger(encoded.Raw)
}

/**
 * Converts a field element to a *big.Int.
 *
 * @param f The field element.
 * @return The *big.Int.
 */
func (ref *mathUtils) FieldToBigInteger(f *Ed25519FieldElement) *big.Int { /* public static  */

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
	x := big.NewInt(-121665)
	x = x.Mul(x, big.NewInt(121666))
	MathUtils.D = x.Mod(x, Ed25519Field.P)
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

/**
 * Scalar multiply the group element by the field element.
 *
 * @param g The group element.
 * @param f The field element.
 * @return The resulting group element.
 */
func (ref *mathUtils) ScalarMultiplyGroupElement(g *Ed25519GroupElement, f Ed25519FieldElement) *Ed25519GroupElement { /* public static  */

	bytes := f.Encode().Raw
	h := Ed25519Group.ZERO_P3
	for i := uint(254); i >= 0; i-- {
		h = ref.DoubleGroupElement(h)
		if utils.GetBitToBool(bytes, i) {
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

/**
 * Converts a *big.Int to a field element.
 *
 * @param b The *big.Int.
 * @return The field element.
 */
func (ref *mathUtils) ToFieldElement(b *big.Int) *Ed25519FieldElement { /* public static  */

	el, err := NewEd25519EncodedFieldElement(utils.BigIntToByteArray(b, 32))
	if err != nil {
		panic(err)
	}
	return el.Decode()
}

/**
 * Creates a group element from a byte array.
 * Bit 0 to 254 are the affine y-coordinate, bit 255 is the sign of the affine x-coordinate.
 *
 * @param bytes the byte array.
 * @return The group element.
 */
func (ref *mathUtils) ToGroupElement(bytes []byte) *Ed25519GroupElement { /* public static  */

	shouldBeNegative := (bytes[31] >> 7) != 0
	bytes[31] &= 0x7f
	y := MathUtils.BytesToBigInteger(bytes)
	x := ref.GetAffineXFromAffineY(y, shouldBeNegative)
	return NewEd25519GroupElementP3(
		ref.ToFieldElement(x),
		ref.ToFieldElement(y),
		Ed25519Field_ONE(),
		ref.ToFieldElement(x.Mul(x, y).Mod(x, Ed25519Field.P)))
}

/**
 * Gets the affine x-coordinate from a given affine y-coordinate and the sign of x.
 *
 * @param y                The affine y-coordinate
 * @param shouldBeNegative true if the negative solution should be chosen, false otherwise.
 * @return The affine x-ccordinate.
 */
func (ref *mathUtils) GetAffineXFromAffineY(y *big.Int, shouldBeNegative bool) *big.Int { /* public static  */

	// x = sign(x) * sqrt((y^2 - 1) / (d * y^2 + 1))
	u := y.Mul(y, y).Sub(y, BigInteger_ONE()).Mod(y, Ed25519Field.P)
	v := ref.D
	v.Mul(v, y).Mul(v, y).Add(v, BigInteger_ONE()).Mod(v, Ed25519Field.P)
	x := ref.getSqrtOfFraction(u, v)
	if v.Mul(v, x).Mul(v, x).Sub(v, u).Mod(v, Ed25519Field.P).Cmp(BigInteger_ZERO()) != 0 {
		if v.Mul(v, x).Mul(v, x).Add(v, u).Mod(v, Ed25519Field.P).Cmp(BigInteger_ZERO()) != 0 {
			panic(errors.New("not a valid Ed25519GroupElement"))
		}

		x = x.Mul(x, ref.IntsToBigInteger(Ed25519Field.I.Raw)).Mod(x, Ed25519Field.P)
	}

	isNegative := x.Mod(x, big.NewInt(2)).Cmp(BigInteger_ONE()) == 0 // final
	if (shouldBeNegative && !isNegative) || (!shouldBeNegative && isNegative) {
		x = x.Neg(x).Mod(x, Ed25519Field.P)
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
func (ref *mathUtils) getSqrtOfFraction(u *big.Int, v *big.Int) *big.Int { /* private static  */

	x := BigInteger_ONE()
	m := BigInteger_ONE()
	tree := big.NewInt(3)
	s := big.NewInt(7)
	tmp := u.Mul(u, v.Exp(v, s, m)).Exp(u, x.Lsh(x, 252).Sub(x, tree), Ed25519Field.P).Mod(u, Ed25519Field.P)
	return tmp.Mul(tmp, u).Mul(tmp, v.Exp(v, tree, m)).Mod(tmp, Ed25519Field.P)
}

/**
 * Converts a group element from one coordinate system to another.
 * This method is a helper used to test various methods in Ed25519GroupElement.
 *
 * @param g                   The group element.
 * @param newCoordinateSystem The desired coordinate system.
 * @return The same group element in the new coordinate system.
 */func (ref *mathUtils) ToRepresentation(g *Ed25519GroupElement, newCoorSys CoordinateSystem) (*Ed25519GroupElement, error) {
	gX := utils.BytesToBigInteger(g.X.Encode().Raw)
	gY := utils.BytesToBigInteger(g.Y.Encode().Raw)
	gZ := utils.BytesToBigInteger(g.Z.Encode().Raw)
	var gT *big.Int
	{
	}
	if g.T != nil {
		gT = utils.BytesToBigInteger(g.T.Encode().Raw)
	}

	// Switch to affine coordinates.
	switch g.coordinateSystem {
	case AFFINE:
		return ref.getNeeCoor(gX, gY, newCoorSys)
	case P2:
	case P3:
		x := gX.Mul(gX, gZ.ModInverse(gZ, Ed25519Field.P))
		x = x.Mod(x, Ed25519Field.P)
		y := gY.Mul(gY, gZ.ModInverse(gZ, Ed25519Field.P))
		y = y.Mod(y, Ed25519Field.P)
		return ref.getNeeCoor(x, y, newCoorSys)
	case P1xP1:
		x := gX.Mul(gX, gZ.ModInverse(gZ, Ed25519Field.P))
		x = x.Mod(x, Ed25519Field.P)
		if gT == nil {
			return nil, errors.New("coordinate T must not nil for P!XP1 ")
		}
		y := gY.Mul(gY, gT.ModInverse(gT, Ed25519Field.P))
		y = y.Mod(y, Ed25519Field.P)
		return ref.getNeeCoor(x, y, newCoorSys)
	case CACHED:
		x := gX.Sub(gX, gY)
		x = x.Mul(x, gZ.Mul(gZ, big.NewInt(2)))
		x = x.ModInverse(x, Ed25519Field.P)
		x = x.Mod(x, Ed25519Field.P)

		y := gX.Add(gX, gY)
		y = y.Mul(y, gZ.Mul(gZ, big.NewInt(2)))
		y = y.ModInverse(y, Ed25519Field.P)
		y = y.Mod(y, Ed25519Field.P)
		return ref.getNeeCoor(x, y, newCoorSys)
	case PRECOMPUTED:
		x := gX.Sub(gX, gY)
		x = x.Mul(x, gZ.Mul(gZ, big.NewInt(2)))
		x = x.ModInverse(x, Ed25519Field.P)
		x = x.Mod(x, Ed25519Field.P)

		y := gX.Add(gX, gY)
		y = y.Mul(y, big.NewInt(2))
		y = y.ModInverse(y, Ed25519Field.P)
		y = y.Mod(y, Ed25519Field.P)
		return ref.getNeeCoor(x, y, newCoorSys)
	}
	return nil, errors.New("NewUnsupportedOperationException")
}
func (ref *mathUtils) getNeeCoor(x, y *big.Int, newCoorSys CoordinateSystem) (*Ed25519GroupElement, error) {

	x1, err := toFieldElement(x)
	if err != nil {
		return nil, err
	}
	y1, err := toFieldElement(y)
	if err != nil {
		return nil, err
	}

	// Now back to the desired coordinate system.
	switch newCoorSys {
	case AFFINE:
		return NewEd25519GroupElementAffine(x1, y1, Ed25519Field_ONE()), nil
	case P2:
		return NewEd25519GroupElementP2(x1, y1, Ed25519Field_ONE()), nil
	case P3:
		m := x.Mul(x, y)
		z, err := toFieldElement(m.Mod(m, Ed25519Field.P))
		if err != nil {
			return nil, err
		}
		return NewEd25519GroupElementP3(x1, y1, Ed25519Field_ONE(), z), nil
	case P1xP1:
		return NewEd25519GroupElementP1XP1(x1, y1, Ed25519Field_ONE(), Ed25519Field_ONE()), nil
	case CACHED:
		m := y.Add(y, x)
		x1, err := toFieldElement(m.Mod(m, Ed25519Field.P))
		if err != nil {
			return nil, err
		}
		m = y.Sub(y, x)
		y1, err := toFieldElement(m.Mod(m, Ed25519Field.P))
		if err != nil {
			return nil, err
		}
		m = ref.D.Mul(ref.D, big.NewInt(2))
		m = m.Mul(m, x)
		m = m.Mul(m, y)
		z, err := toFieldElement(m.Mod(m, Ed25519Field.P))
		if err != nil {
			return nil, err
		}
		return NewEd25519GroupElementCached(x1, y1, Ed25519Field_ONE(), z), nil
	case PRECOMPUTED:
		m := y.Add(y, x)
		x1, err := toFieldElement(m.Mod(m, Ed25519Field.P))
		if err != nil {
			return nil, err
		}
		m = y.Sub(y, x)
		y1, err := toFieldElement(m.Mod(m, Ed25519Field.P))
		if err != nil {
			return nil, err
		}
		m = ref.D.Mul(ref.D, big.NewInt(2))
		m = m.Mul(m, x)
		m = m.Mul(m, y)
		z, err := toFieldElement(m.Mod(m, Ed25519Field.P))
		if err != nil {
			return nil, err
		}
		return NewEd25519GroupElementPrecomputed(x1, y1, z), nil
	}
	return nil, errors.New("NewUnsupportedOperationException")
}
func (ref *mathUtils) GetRandomEncodedFieldElement(length int) *Ed25519EncodedFieldElement {

	bytes := ref.GetRandomByteArray(length)
	bytes[31] &= 0x7f
	return &Ed25519EncodedFieldElement{Ed25519Field_ZERO_SHORT(), bytes}
}

/**
 * Gets a random group element in P3 coordinates.
 * It's NOT guaranteed that the created group element is a multiple of the base point.
 *
 * @return The group element.
 */
func (ref *mathUtils) GetRandomGroupElement() *Ed25519GroupElement { /* public static  */

	bytes := ref.random()
	gr, err := NewEd25519EncodedGroupElement(bytes[:])
	if err != nil {
		panic(err)
	}
	el, err := gr.Decode()
	if err != nil {
		panic(err)
	}

	return el
}

/**
 * Gets a random encoded group element.
 * It's NOT guaranteed that the created encoded group element is a multiple of the base point.
 *
 * @return The encoded group element.
 */
func (ref *mathUtils) GetRandomEncodedGroupElement() *Ed25519EncodedGroupElement { /* public static  */

	gr := ref.GetRandomGroupElement()
	el, err := gr.Encode()
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
		fmt.Print(err)
	}
	return int64(binary.LittleEndian.Uint64(bytes))
}

// Reduces an encoded field element modulo the group order and returns the result.
// *
// * @param encoded The encoded field element.
// * @return The mod group order reduced encoded field element.
func (ref *mathUtils) ReduceModGroupOrder(encoded *Ed25519EncodedFieldElement) *Ed25519EncodedFieldElement {

	b := ref.BytesToBigInteger(encoded.Raw)
	b.Mod(b, Ed25519Group.GROUP_ORDER)
	return ref.ToEncodedFieldElement(b)
}

/**
 * Converts a biginteger to an encoded field element.
 *
 * @param b The biginteger.
 * @return The encoded field element.
 */
func (ref *mathUtils) ToEncodedFieldElement(b *big.Int) *Ed25519EncodedFieldElement {

	return &Ed25519EncodedFieldElement{Ed25519Field_ZERO_SHORT(), utils.BigIntToByteArray(b, 32)}
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
	original := b.Bytes()
	// Although b < 2^256, original can have length > 32 with some bytes set to 0.
	offset := 0
	if len(original) > 32 {
		offset = len(original) - 32
	}
	for i := range original[:len(original)-offset] {
		bytes[len(original)-i-offset-1] = original[i+offset]
	}

	return bytes
}
func coorModify(g, b *big.Int) *big.Int {
	return g.Mul(g, (&big.Int{}).ModInverse(b, Ed25519Field.P)).Mod(g, Ed25519Field.P)
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
func (ref *mathUtils) AddGroupElements(g1 *Ed25519GroupElement, g2 *Ed25519GroupElement) *Ed25519GroupElement { /* public static  */

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
	x3 := ref.funcName(*g1x, g2y, g2x, g1y, (&big.Int{}).Add(one, dx1x2y1y2))
	y3 := ref.funcName(*g1x, g2x, g1y, g2y, (&big.Int{}).Sub(one, dx1x2y1y2))
	t3 := x3.Mul(x3, y3).Mod(x3, Ed25519Field.P)
	return NewEd25519GroupElementP3(
		ref.ToFieldElement(x3),
		ref.ToFieldElement(y3),
		Ed25519Field_ONE(),
		ref.ToFieldElement(t3))
}

func (ref *mathUtils) funcName(sp big.Int, g2y, g2x, g1y, one *big.Int) *big.Int {
	b := &sp
	return b.Mul(b, g2y).Add(b, (&big.Int{}).Mul(g2x, g1y)).Mul(b, one.ModInverse(one, Ed25519Field.P)).Mod(b, Ed25519Field.P)
}

/**
 * Doubles a group element and returns the result in the P3 coordinate system.
 * It uses *big.Int arithmetic and the affine coordinate system.
 * This method is a helper used to test the projective group doubling formula in Ed25519GroupElement.
 *
 * @param g The group element.
 * @return g+g.
 */
func (ref *mathUtils) DoubleGroupElement(g *Ed25519GroupElement) *Ed25519GroupElement { /* public static  */

	return ref.AddGroupElements(g, g)
}

/**
* Negates a group element.
*
* @param g The group element.
* @return The negated group element.
 */
func (ref *mathUtils) NegateGroupElement(g *Ed25519GroupElement) *Ed25519GroupElement { /* public static  */

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
func toFieldElement(b *big.Int) (*Ed25519FieldElement, error) {

	elem, err := NewEd25519EncodedFieldElement(utils.BigIntToByteArray(b, 32))
	if err != nil {
		return nil, err
	}

	return elem.Decode(), nil
}
