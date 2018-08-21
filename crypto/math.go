package crypto

import (
	rand2 "crypto/rand"
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
	random    [32]byte
	EXPONENTS []uint
	D         *big.Int
}

//* Converts a 2^8 bit representation to a BigInteger.
//* Value: bytes[0] + 2^8 * bytes[1] + ...
//*
//* @param bytes The 2^8 bit representation.
//* @return The BigInteger.
func (ref *mathUtils) BytesToBigInteger(bytes []byte) *big.Int {

	b := BigInteger_ZERO()
	for i, val := range bytes {
		el := (&big.Int{}).SetUint64(uint64(uint8(val)) & 0xff)
		one := BigInteger_ONE()
		one = one.Mul(one, el)
		b = b.Add(b, one.Lsh(one, uint(i*8)))
	}

	return b
}

//     * Converts a 2^25.5 bit representation to a BigInteger.
//     * Value: 2^EXPONENTS[0] * t[0] + 2^EXPONENTS[1] * t[1] + ... + 2^EXPONENTS[9] * t[9]
//     *
//     * @param t The 2^25.5 bit representation.
//     * @return The BigInteger.
func (ref *mathUtils) IntsToBigInteger(t []int) *big.Int {
	b := BigInteger_ZERO()
	for i, val := range t {
		el := (&big.Int{}).SetInt64(int64(val) & 0xff)
		one := BigInteger_ONE()
		one = one.Mul(one, el)
		b = b.Add(b, one.Lsh(one, ref.EXPONENTS[i]))
	}

	return b
}

var MathUtils mathUtils

func init() {
	rand := rand2.Reader
	_, err := io.ReadFull(rand, MathUtils.random[:])
	if err != nil {
		fmt.Print(err)
	}
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
func (ref *mathUtils) GetRandomGroupElement() (*Ed25519GroupElement, error) {

	grElem, err := NewEd25519EncodedGroupElement(ref.random[:])
	if err != nil {
		return nil, err
	}
	return grElem.Decode()
}
func (ref *mathUtils) GetRandomFieldElement() *Ed25519FieldElement {

	t := make([]intRaw, 10)
	rand := rand2.Reader

	for j := range t {
		var v [4]byte
		_, err := rand.Read(v[:])
		if err != nil {
			fmt.Println(err)
			return nil
		}
		//(1 << 25)
		t[j] = (&big.Int{}).SetBytes(v[:]).Int64() - (1 << 24)
	}

	return &Ed25519FieldElement{t}
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
		return Ed25519GroupElementAffine(x1, y1, Ed25519Field.ONE), nil
	case P2:
		return Ed25519GroupElementP2(x1, y1, Ed25519Field.ONE), nil
	case P3:
		m := x.Mul(x, y)
		z, err := toFieldElement(m.Mod(m, Ed25519Field.P))
		if err != nil {
			return nil, err
		}
		return Ed25519GroupElementP3(x1, y1, Ed25519Field.ONE, z), nil
	case P1xP1:
		return Ed25519GroupElementP1XP1(x1, y1, Ed25519Field.ONE, Ed25519Field.ONE), nil
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
		return Ed25519GroupElementCached(x1, y1, Ed25519Field.ONE, z), nil
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
		return Ed25519GroupElementPrecomputed(x1, y1, z), nil
	}
	return nil, errors.New("NewUnsupportedOperationException")
}
func toFieldElement(b *big.Int) (*Ed25519FieldElement, error) {

	elem, err := NewEd25519EncodedFieldElement(utils.BigIntToByteArray(b, 32))
	if err != nil {
		return nil, err
	}

	return elem.Decode(), nil
}
