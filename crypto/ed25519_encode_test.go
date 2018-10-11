// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.


package crypto

import (
	"fmt"
	"github.com/proximax-storage/nem2-sdk-go/utils"
	"github.com/stretchr/testify/assert"
	"math/big"
	"runtime"
	"testing"
)

const numIter = 1000

// region Ed25519FieldElement

// TestNewEd25519FieldElement test correct lenght raw checking
func TestNewEd25519FieldElement(t *testing.T) {
	for i := 0; i < 100; i++ {
		_, err := NewEd25519FieldElement(make([]intRaw, i))
		if i == lenEd25519FieldElementRaw {
			assert.Nil(t, err)
		} else {
			assert.NotNil(t, err)
		}
	}

}
func TestIsNonZeroReturnsFalseIfFieldElementIsZero(t *testing.T) {

	test := [10]intRaw{}
	f, _ := NewEd25519FieldElement(test[:])

	assert.Equal(t, false, f.IsNonZero(), `f.isNonZero() must be false!`)
}

func TestIsNonZeroReturnsTrueIfFieldElementIsNonZero(t *testing.T) {

	test := [10]intRaw{}
	test[0] = 5
	f, _ := NewEd25519FieldElement(test[:])

	assert.Equal(t, true, f.IsNonZero(), `f.isNonZero() must be true!`)
}
func TestEd25519FieldElement_GetRawReturnsUnderlyingArray(t *testing.T) {

	values := [10]intRaw{}
	values[0] = 5
	values[6] = 15
	values[8] = -67
	f, _ := NewEd25519FieldElement(values[:])

	assert.Equal(t, values[:], f.Raw, `values and f.GetRaw() must by equal !`)
}

// TestAddReturnsCorrectResult test Ed25519FieldElement summaring corection
func TestAddReturnsCorrectResult(t *testing.T) {

	for i := 0; i < numIter; i++ {
		f1 := MathUtils.GetRandomFieldElement()
		f2 := MathUtils.GetRandomFieldElement()
		b1 := MathUtils.BytesToBigInteger(f1.Encode().Raw)
		b2 := MathUtils.BytesToBigInteger(f2.Encode().Raw)

		f3 := f1.add(f2)

		assertEquals(t, f3, b1.Add(b1, b2))
	}

}
func TestSubtractReturnsCorrectResult(t *testing.T) {

	for i := 0; i < numIter; i++ {

		f1 := MathUtils.GetRandomFieldElement()
		f2 := MathUtils.GetRandomFieldElement()
		b1 := MathUtils.BytesToBigInteger(f1.Encode().Raw)
		b2 := MathUtils.BytesToBigInteger(f2.Encode().Raw)

		f3 := f1.subtract(f2)

		assertEquals(t, f3, b1.Sub(b1, b2))
	}

}
func TestEd25519FieldAdd_Sub(t *testing.T) {

	for i := 0; i < numIter; i++ {
		f1 := MathUtils.GetRandomFieldElement()
		f2 := MathUtils.GetRandomFieldElement()
		b1 := MathUtils.BytesToBigInteger(f1.Encode().Raw)
		b2 := MathUtils.BytesToBigInteger(f2.Encode().Raw)

		f3 := f1.add(f2)
		assertEquals(t, f3, b1.Add(b1, b2))
		b3 := MathUtils.BytesToBigInteger(f3.Encode().Raw)

		f4 := f3.subtract(f2)
		assertEquals(t, f4, b3.Sub(b3, b2))

		assert.Equal(t, f1, f4)
		assertEquals(t, f4, b1.Sub(b1, b2))
	}
}

func TestNegateReturnsCorrectResult(t *testing.T) {

	for i := 0; i < numIter; i++ {
		f1 := MathUtils.GetRandomFieldElement()
		b1 := MathUtils.BytesToBigInteger(f1.Encode().Raw)

		f2 := f1.negate()

		assertEquals(t, f2, b1.Neg(b1))
	}

}
func TestEd25519FieldMultiply_Mirror(t *testing.T) {

	for i := 0; i < numIter; i++ {
		f1 := MathUtils.GetRandomFieldElement()
		f2 := MathUtils.GetRandomFieldElement()

		f3 := f1.multiply(f2)
		f4 := f2.multiply(f1)
		assert.Equal(t, f3, f4)
	}
}
func TestEd25519FieldElement_Muiltiply(t *testing.T) {
	raw := make([]intRaw, lenEd25519FieldElementRaw)
	for i := 0; i < numIter; i++ {
		for j := 0; j < lenEd25519FieldElementRaw; j++ {
			raw[j] = intRaw(i+1) - (1 << 24)
			f1 := Ed25519FieldElement{raw}
			f2 := MathUtils.GetRandomFieldElement()
			f3 := f1.multiply(f2)

			b1 := MathUtils.BytesToBigInteger(f1.Encode().Raw)
			b2 := MathUtils.BytesToBigInteger(f2.Encode().Raw)
			b3 := MathUtils.BytesToBigInteger(f3.Encode().Raw)
			assertEquals(t, f3, b3.Mul(b1, b2))
		}
	}
}
func TestMultiplyReturnsCorrectResult(t *testing.T) {

	for i := 0; i < numIter; i++ {
		f1 := MathUtils.GetRandomFieldElement()
		f2 := MathUtils.GetRandomFieldElement()
		b1 := MathUtils.BytesToBigInteger(f1.Encode().Raw)
		b2 := MathUtils.BytesToBigInteger(f2.Encode().Raw)

		f3 := f1.multiply(f2)
		if !assertEquals(t, f3, (&big.Int{}).Mul(b1, b2)) {
			assert.Equal(t, b1, b2, "iter=%d", i)
		}
		f4 := f2.multiply(f1)
		assert.Equal(t, f3, f4)
	}

}

// region EncodedFieldElement

func TestNewEd25519EncodedFieldElement(t *testing.T) {
	for i := 0; i < 100; i++ {
		_, err := NewEd25519EncodedFieldElement(make([]byte, i))
		if (i == 32) || (i == 64) {
			assert.Nil(t, err)
		} else {
			assert.NotNil(t, err)
		}
	}

}
func TestEncodedFieldElement_ReturnsFalse(t *testing.T) {

	test := make([]byte, 32)
	f, _ := NewEd25519EncodedFieldElement(test)

	assert.Equal(t, false, f.IsNonZero(), `f.isNonZero() must be false!`)
}

func TestEncodedFieldElement_ReturnsTrue(t *testing.T) {

	test := make([]byte, 32)
	test[0] = 5
	f, _ := NewEd25519EncodedFieldElement(test)

	assert.Equal(t, true, f.IsNonZero(), `f.isNonZero() must be true!`)
}

func TestEd25519EncodedFieldElement_GetRawReturnsUnderlyingArray(t *testing.T) {

	values := make([]byte, 32)
	values[0] = 5
	values[6] = 15
	values[8] = 67
	f, _ := NewEd25519EncodedFieldElement(values)

	assert.Equal(t, values[:], f.Raw, `values and f.GetRaw() must by equal !`)
}
func TestDecodePlusEncodeDoesNotAlterTheEncodedFieldElement(t *testing.T) {

	for i := 0; i < numIter; i++ {
		original := MathUtils.GetRandomEncodedFieldElement(32)
		encoded := original.Decode().Encode()
		assert.True(t, original.Equals(encoded), `encoded and original must by equal !`)
	}

}

// endregion
// region modulo group order arithmetic
func TestEd25519EncodedFieldElement_ModQReturnsExpectedResult(t *testing.T) {

	for i := 0; i < numIter; i++ {
		encoded := &Ed25519EncodedFieldElement{Ed25519Field_ZERO_LONG(), MathUtils.GetRandomByteArray(64)}
		reduced1 := encoded.modQ()
		b1 := MathUtils.BytesToBigInteger(reduced1.Raw)

		reduced2 := MathUtils.ReduceModGroupOrder(encoded)

		assert.Equal(t, -1, b1.Cmp(Ed25519Field.P), `MathUtils.toBigInteger(reduced1).compareTo(Ed25519Field.P) and -1 must by equal !`)
		assert.Equal(t, 1, b1.Cmp(BigInteger_ONE()), `MathUtils.toBigInteger(reduced1).compareTo(Newuint64("-1")) and 1 must by equal !`)
		assert.True(t, reduced2.Equals(reduced1), `reduced1 and reduced2 must by equal ! Iter = %d`, i)
	}

}
func TestEd25519EncodedFieldElement_MultiplyAndAddModQReturnsExpectedResult(t *testing.T) {

	for i := 0; i < numIter; i++ {
		encoded1 := MathUtils.GetRandomEncodedFieldElement(32)
		encoded2 := MathUtils.GetRandomEncodedFieldElement(32)
		encoded3 := MathUtils.GetRandomEncodedFieldElement(32)
		result1 := encoded1.multiplyAndAddModQ(encoded2, encoded3)
		b1 := MathUtils.BytesToBigInteger(result1.Raw)
		result2 := MathUtils.multiplyAndAddModGroupOrder(encoded1, encoded2, encoded3)
		assert.Equal(t, -1, b1.Cmp(Ed25519Field.P), `MathUtils.toBigInteger(result1).compareTo(Ed25519Field.P) and -1 must by equal !`)
		assert.Equal(t, 1, b1.Cmp(big.NewInt(-1)), `MathUtils.toBigInteger(result1).compareTo(Newuint64("-1")) and 1 must by equal !`)
		assert.True(t, result1.Equals(result2), `result1 and result2 must by equal ! Iter %d`, i)
	}

}

func TestEd25519FieldElementEncode_ReturnsCorrectByteArrayForSimpleFieldElements(t *testing.T) {

	t1 := make([]intRaw, 10)
	t2 := make([]intRaw, 10)
	t2[0] = 1
	fieldElement1 := &Ed25519FieldElement{t1}

	fieldElement2 := &Ed25519FieldElement{t2}
	// Act:
	encoded1 := fieldElement1.Encode()
	encoded2 := fieldElement2.Encode()
	// Assert:
	want1, err := MathUtils.toEncodedFieldElement(BigInteger_ZERO())
	assert.Nil(t, err)
	want2, err := MathUtils.toEncodedFieldElement(BigInteger_ONE())
	assert.Nil(t, err)
	assert.Equal(t, want1, encoded1, `encoded1 and MathUtils.toEncodedFieldElement(*big.Int.ZERO) must by equal !`)
	assert.Equal(t, want2, encoded2, `encoded2 and MathUtils.toEncodedFieldElement(*big.Int.ONE) must by equal !`)
}

func TestEncodeReturnsCorrectByteArrayIfJthBitOfTiIsSetToOne(t *testing.T) {

	for i := 0; i < 10; i++ {

		test := make([]intRaw, 10)
		for j := uint(0); j < 24; j++ {
			test[i] = 1 << j
			fieldElement, err := NewEd25519FieldElement(test)
			assert.Nil(t, err)
			b := MathUtils.IntsToBigInteger(test)
			b.Mod(b, Ed25519Field.P)
			// Act:
			encoded := fieldElement.Encode()
			// Assert:
			want, err := MathUtils.toEncodedFieldElement(b)
			assert.Nil(t, err)
			assert.Equal(t, want, encoded, `encoded and MathUtils.toEncodedFieldElement(b) must by equal !`)
		}
	}
}

// @Test
func TestEncodeReturnsCorrectByteArray(t *testing.T) {

	for i := 0; i < numIter; i++ {

		test := make([]intRaw, 10)
		for j := uint(0); j < 10; j++ {
			//random.nextInt(1 << 28) - (1 << 27);
			test[j] = MathUtils.GetRandomIntRaw() - (1 << 27)
		}

		fieldElement, err := NewEd25519FieldElement(test)
		assert.Nil(t, err)
		b := MathUtils.IntsToBigInteger(test)
		encoded := fieldElement.Encode()
		// Assert:
		want, err := MathUtils.toEncodedFieldElement(b.Mod(b, Ed25519Field.P))
		assert.Nil(t, err)
		assert.Equal(t, want, encoded, `encoded and MathUtils.toEncodedFieldElement(b.mod(Ed25519Field.P)) must by equal !`)
	}

}

// region isNegative
// @Test
func TestIsNegativeReturnsCorrectResult(t *testing.T) {

	for i := 0; i < numIter; i++ {

		values := MathUtils.GetRandomByteArray(32)
		values[31] &= 0x7F
		encoded, err := NewEd25519EncodedFieldElement(values)
		assert.Nil(t, err)

		b := MathUtils.BytesToBigInteger(values)
		b = b.Mod(b, Ed25519Field.P)
		b = b.Mod(b, big.NewInt(2))
		isNegative := b.Int64() == BigInteger_ONE().Int64()
		// Assert:
		assert.Equal(t, encoded.IsNegative(), isNegative, `encoded.isNegative() and isNegative must by equal !`)
	}

}

func TestEncodedFieldElement_EqualsOnlyReturnsTrueForEquivalentObjects(t *testing.T) {

	encoded1 := MathUtils.GetRandomEncodedFieldElement(32)
	encoded2 := encoded1.Decode().Encode()
	encoded3 := MathUtils.GetRandomEncodedFieldElement(32)
	encoded4 := MathUtils.GetRandomEncodedFieldElement(32)
	// Assert:
	assert.True(t, encoded1.Equals(encoded2), `encoded1 and encoded2 must by equal !`)
	assert.NotEqual(t, encoded1, encoded3, `encoded1 and encoded3 must by not equal !`)
	assert.NotEqual(t, encoded1, encoded4, `encoded1 and encoded4 must by not equal !`)
	assert.NotEqual(t, encoded3, encoded4, `encoded3 and encoded4 must by not equal !`)
}
func TestEd25519EncodedGroupElement_GetAffineX(t *testing.T) {

	defer testRecover(t)
	for i := 0; i < numIter; i++ {
		enc := MathUtils.GetRandomGroupElement()
		encoded, err := enc.Encode()
		if err != nil {
			t.Fatal(err)
		}

		affineX1, err := encoded.GetAffineX()
		if err != nil {
			t.Fatal(err)
		}
		g, err := encoded.Decode()
		if err != nil {
			t.Fatal(err)
		}
		X2, err := MathUtils.ToRepresentation(g, AFFINE)
		if err != nil {
			t.Fatal(err)
		}
		affineX2 := X2.GetX()

		assert.Equal(t, affineX2, affineX1, `affineX1 and affineX2 must by  equal !`)
	}
}
func TestSquareReturnsCorrectResult(t *testing.T) {

	defer func() {
		err := recover()
		t.Log(err)
	}()
	for i := 0; i < numIter; i++ {
		f1 := MathUtils.GetRandomFieldElement()
		b1 := MathUtils.BytesToBigInteger(f1.Encode().Raw)
		f2 := f1.square()

		assertEquals(t, f2, b1.Mul(b1, b1))
	}
}
func TestSqrtReturnsCorrectResult(t *testing.T) {

	defer func() {
		err := recover()
		t.Log(err)
	}()
	for i := 0; i < numIter; i++ {
		u := MathUtils.GetRandomFieldElement()
		uSquare := u.square()
		v := MathUtils.GetRandomFieldElement()
		vSquare := v.square()
		fraction := u.multiply(v.invert())

		sqrt := Ed25519FieldElementSqrt(uSquare, vSquare)

		// (u / v)^4 == (sqrt(u^2 / v^2))^4.
		assert.Equal(t, sqrt, fraction, `fraction and sqrt must by  equal !`)
		assert.Equal(t, sqrt.square(), fraction.square(), `fraction.square().square() and sqrt.square().square() must by  equal !`)
		assert.Equal(t, sqrt.square().square(), fraction.square().square(), `fraction.square().square() and sqrt.square().square() must by  equal !`)
		// (u / v) == +-1 * sqrt(u^2 / v^2) or (u / v) == +-i * sqrt(u^2 / v^2)
		assert.Equal(t, true, differsOnlyByAFactorOfAFourthRootOfOne(fraction, sqrt))
		break
	}
}

//end region
// region Ed25519GroupElementTest
func TestCanBeCreatedWithP2Coordinates(t *testing.T) {

	g := NewEd25519GroupElementP2(Ed25519Field_ZERO(), Ed25519Field_ONE(), Ed25519Field_ONE())

	assert.Equal(t, g.GetCoordinateSystem(), P2, `g.GetCoordinateSystem() and P2 must by equal !`)
	assert.Equal(t, g.GetX(), Ed25519Field_ZERO(), `g.GetX() and Ed25519Field.ZERO must by equal !`)
	assert.Equal(t, g.GetY(), Ed25519Field_ONE(), `g.GetY() and Ed25519Field.ONE must by equal !`)
	assert.Equal(t, g.GetZ(), Ed25519Field_ONE(), `g.GetZ() and Ed25519Field.ONE must by equal !`)
	assert.Nil(t, g.GetT())
}

// @Test
func TestCanBeCreatedWithP3Coordinates(t *testing.T) {

	g := NewEd25519GroupElementP3(Ed25519Field_ZERO(), Ed25519Field_ONE(), Ed25519Field_ONE(), Ed25519Field_ZERO())

	assert.Equal(t, g.GetCoordinateSystem(), P3, `g.GetCoordinateSystem() and P3 must by equal !`)
	assert.Equal(t, g.GetX(), Ed25519Field_ZERO(), `g.GetX() and Ed25519Field.ZERO must by equal !`)
	assert.Equal(t, g.GetY(), Ed25519Field_ONE(), `g.GetY() and Ed25519Field.ONE must by equal !`)
	assert.Equal(t, g.GetZ(), Ed25519Field_ONE(), `g.GetZ() and Ed25519Field.ONE must by equal !`)
}

func TestCanBeCreatedWithP1P1Coordinates(t *testing.T) {

	g := NewEd25519GroupElementP1XP1(Ed25519Field_ZERO(), Ed25519Field_ONE(), Ed25519Field_ONE(), Ed25519Field_ONE())

	assert.Equal(t, g.GetCoordinateSystem(), P1xP1, `g.GetCoordinateSystem() and P1xP1 must by equal !`)
	assert.Equal(t, g.GetX(), Ed25519Field_ZERO(), `g.GetX() and Ed25519Field.ZERO must by equal !`)
	assert.Equal(t, g.GetY(), Ed25519Field_ONE(), `g.GetY() and Ed25519Field.ONE must by equal !`)
	assert.Equal(t, g.GetZ(), Ed25519Field_ONE(), `g.GetZ() and Ed25519Field.ONE must by equal !`)
	assert.Equal(t, g.GetT(), Ed25519Field_ONE(), `g.GetT() and Ed25519Field.ONE must by equal !`)
}

func TestCanBeCreatedWithPrecompCoordinates(t *testing.T) {

	g := NewEd25519GroupElementPrecomputed(Ed25519Field_ONE(), Ed25519Field_ONE(), Ed25519Field_ZERO())

	assert.Equal(t, g.GetCoordinateSystem(), PRECOMPUTED, `g.GetCoordinateSystem() and PRECOMPUTED must by equal !`)
	assert.Equal(t, g.GetX(), Ed25519Field_ONE(), `g.GetX() and Ed25519Field_ONE() must by equal !`)
	assert.Equal(t, g.GetY(), Ed25519Field_ONE(), `g.GetY() and Ed25519Field_ONE() must by equal !`)
	assert.Equal(t, g.GetZ(), Ed25519Field_ZERO(), `g.GetZ() and Ed25519Field_ZERO() must by equal !`)
	assert.Nil(t, g.GetT())
}

// @Test
func TestCanBeCreatedWithCachedCoordinates(t *testing.T) {

	g := NewEd25519GroupElementCached(Ed25519Field_ONE(), Ed25519Field_ONE(), Ed25519Field_ONE(), Ed25519Field_ZERO())

	assert.Equal(t, g.GetCoordinateSystem(), CACHED, `g.GetCoordinateSystem() and CACHED must by equal !`)
	assert.Equal(t, g.GetX(), Ed25519Field_ONE(), `g.GetX() and Ed25519Field_ONE() must by equal !`)
	assert.Equal(t, g.GetY(), Ed25519Field_ONE(), `g.GetY() and Ed25519Field_ONE() must by equal !`)
	assert.Equal(t, g.GetZ(), Ed25519Field_ONE(), `g.GetZ() and Ed25519Field_ONE() must by equal !`)
	assert.Equal(t, g.GetT(), Ed25519Field_ZERO(), `g.GetT() and Ed25519Field_ZERO() must by equal !`)
}

func TestCanBeCreatedWithSpecifiedCoordinates(t *testing.T) {

	g := NewEd25519GroupElement(
		P3,
		Ed25519Field_ZERO(),
		Ed25519Field_ONE(),
		Ed25519Field_ONE(),
		Ed25519Field_ZERO())

	assert.Equal(t, g.GetCoordinateSystem(), P3, `g.GetCoordinateSystem() and P3 must by equal !`)
	assert.Equal(t, g.GetX(), Ed25519Field_ZERO(), `g.GetX() and Ed25519Field_ZERO() must by equal !`)
	assert.Equal(t, g.GetY(), Ed25519Field_ONE(), `g.GetY() and Ed25519Field_ONE() must by equal !`)
	assert.Equal(t, g.GetZ(), Ed25519Field_ONE(), `g.GetZ() and Ed25519Field_ONE() must by equal !`)
	assert.Equal(t, g.GetT(), Ed25519Field_ZERO(), `g.GetT() and Ed25519Field_ZERO() must by equal !`)
}

// @Test
func TestConstructorUsingEncodedGroupElementReturnsExpectedResult(t *testing.T) { /* public  */

	defer testRecover(t)
	for i := 0; i < 100; i++ {

		g := MathUtils.GetRandomGroupElement()

		encoded, err := g.Encode()
		assert.Nil(t, err)

		h1, err := encoded.Decode()
		assert.Nil(t, err)
		assert.Equal(t, g, h1, `h1 and h2 must by equal !`)

		h2, err := MathUtils.ToGroupElement(encoded.Raw)
		if (err == errNoValidEd25519Group) || !assert.Nil(t, err) {
			continue
		}

		if !assert.Equal(t, h1, h2, `h1 and h2 must by equal !`) {
			assert.Equal(t, h1.X, h2.X)
			assert.Equal(t, h1.Y, h2.Y)
			assert.Equal(t, h1.Z, h2.Z)
			assert.Equal(t, h1.T, h2.T)
		}
	}
}
func TestEncodeReturnsExpectedResult(t *testing.T) {

	defer func() {
		err := recover()
		t.Log(err)
	}()
	for i := 0; i < 100; i++ {

		g := MathUtils.GetRandomGroupElement()
		encoded, err := g.Encode()
		assert.Nil(t, err)
		bytes := utils.BigIntToByteArray(MathUtils.FieldToBigInteger(g.GetY()), 32)

		b := MathUtils.FieldToBigInteger(g.GetX())
		if b.Mod(b, big.NewInt(2)).Cmp(MathUtils.FieldToBigInteger(Ed25519Field_ONE())) == 0 {
			bytes[31] |= 0x80
		}

		assert.Equal(t, encoded.Raw, bytes)
	}

}

//(expected = IllegalArgumentException.class)
func TestToP2ThrowsIfGroupElementHasPrecompRepresentation(t *testing.T) {

	defer func() {
		err := recover()
		t.Log(err)
	}()
	grEl := MathUtils.GetRandomGroupElement()

	g, err := MathUtils.ToRepresentation(grEl, PRECOMPUTED)
	assert.Nil(t, err)

	g.toP2()
}

//(expected = IllegalArgumentException.class)
func TestToP2ThrowsIfGroupElementHasCachedRepresentation(t *testing.T) { /* public  */
	defer func() {
		err := recover()
		t.Log(err)
	}()
	grEl := MathUtils.GetRandomGroupElement()

	g, err := MathUtils.ToRepresentation(grEl, CACHED)
	assert.Nil(t, err)

	g.toP2()
}

func TestToP2ReturnsExpectedResultIfGroupElementHasP2Representation(t *testing.T) {

	defer func() {
		err := recover()
		t.Log(err)
	}()
	for i := 0; i < 10; i++ {

		grEl := MathUtils.GetRandomGroupElement()

		g, err := MathUtils.ToRepresentation(grEl, P2)
		assert.Nil(t, err)

		h := g.toP2()

		assert.Equal(t, h, g, `h and g must by equal !`)
		assert.Equal(t, h.GetCoordinateSystem(), P2, `h.GetCoordinateSystem() and P2 must by equal !`)
		assert.Equal(t, h.GetX(), g.GetX(), `h.GetX() and g.GetX() must by equal !`)
		assert.Equal(t, h.GetY(), g.GetY(), `h.GetY() and g.GetY() must by equal !`)
		assert.Equal(t, h.GetZ(), g.GetZ(), `h.GetZ() and g.GetZ() must by equal !`)
		assert.Nil(t, h.GetT())
	}

}

// @Test
func TestToP2ReturnsExpectedResultIfGroupElementHasP3Representation(t *testing.T) {

	defer testRecover(t)
	for i := 0; i < 10; i++ {

		g := MathUtils.GetRandomGroupElement()

		h1 := g.toP2()
		h2, err := MathUtils.ToRepresentation(g, P2)
		assert.Nil(t, err)

		assert.Equal(t, h1, h2, `h1 and h2 must by equal !`)
		assert.Equal(t, h1.GetCoordinateSystem(), P2, `h1.GetCoordinateSystem() and P2 must by equal !`)
		assert.Equal(t, h1.GetX(), g.GetX(), `h1.GetX() and g.GetX() must by equal !`)
		assert.Equal(t, h1.GetY(), g.GetY(), `h1.GetY() and g.GetY() must by equal !`)
		assert.Equal(t, h1.GetZ(), g.GetZ(), `h1.GetZ() and g.GetZ() must by equal !`)
		assert.Nil(t, h1.GetT())
	}
}

func testRecover(t *testing.T) {
	err := recover()
	switch err.(type) {
	case error:
		buf := make([]byte, 1024)
		runtime.Stack(buf, true)
		t.Fatal(err, string(buf))
	case nil:
	default:
		t.Log(err)

	}
}

//(expected = IllegalArgumentException.class)
func TestToCachedThrowsIfGroupElementHasP2Representation(t *testing.T) {

	defer testRecover(t)
	grEl := MathUtils.GetRandomGroupElement()

	g, err := MathUtils.ToRepresentation(grEl, P2)
	assert.Nil(t, err)
	g.toCached()
}

// @Test(expected = IllegalArgumentException.class)
func TestToCachedThrowsIfGroupElementHasPrecompRepresentation(t *testing.T) {

	defer testRecover(t)
	grEl := MathUtils.GetRandomGroupElement()

	g, err := MathUtils.ToRepresentation(grEl, PRECOMPUTED)
	assert.Nil(t, err)

	g.toCached()
}

//(expected = IllegalArgumentException.class)
func TestToCachedThrowsIfGroupElementHasP1P1Representation(t *testing.T) {

	defer func() {
		err := recover()
		t.Log(err)
	}()
	grEl := MathUtils.GetRandomGroupElement()
	g, err := MathUtils.ToRepresentation(grEl, P1xP1)
	assert.Nil(t, err)
	g.toCached()
}

func TestToCachedReturnsExpectedResultIfGroupElementHasCachedRepresentation(t *testing.T) {

	defer testRecover(t)

	for i := 0; i < 10; i++ {
		grEl := MathUtils.GetRandomGroupElement()

		g, err := MathUtils.ToRepresentation(grEl, CACHED)
		assert.Nil(t, err)
		h := g.toCached()
		assert.Equal(t, h, g, `h and g must by equal !`)
		assert.Equal(t, h.GetCoordinateSystem(), CACHED, `h.GetCoordinateSystem() and CACHED must by equal !`)
		assert.Equal(t, h, g, `h and g must by equal !`)
		assert.Equal(t, h.GetX(), g.GetX(), `h.GetX() and g.GetX() must by equal !`)
		assert.Equal(t, h.GetY(), g.GetY(), `h.GetY() and g.GetY() must by equal !`)
		assert.Equal(t, h.GetZ(), g.GetZ(), `h.GetZ() and g.GetZ() must by equal !`)
		assert.Equal(t, h.GetT(), g.GetT(), `h.GetT() and g.GetT() must by equal !`)
	}

}

// @Test
func TestToCachedReturnsExpectedResultIfGroupElementHasP3Representation(t *testing.T) {

	defer testRecover(t)
	for i := 0; i < 10; i++ {
		grEl := MathUtils.GetRandomGroupElement()
		h1 := grEl.toCached()

		h2, err := MathUtils.ToRepresentation(grEl, CACHED)
		assert.Nil(t, err)
		assert.Equal(t, h1.GetCoordinateSystem(), CACHED, `h1.GetCoordinateSystem() and CACHED must by equal !`)
		assert.True(t, h1.Equals(grEl), `h and grEl must by equal !`)

		x := grEl.GetX()
		gYX := grEl.Y.add(*x)
		gY_X := grEl.Y.subtract(*x)
		gTM := grEl.T.multiply(Ed25519Field.D_Times_TWO)

		assert.True(t, h1.GetX().Equals(&gYX), `h1.GetX() and grEl.Y.add(grEl.GetX()) must by equal !`)
		assert.True(t, h1.GetY().Equals(&gY_X), `h1.GetY() and grEl.Y.subtract(grEl.GetX()) must by equal !`)
		assert.True(t, h1.GetZ().Equals(grEl.GetZ()), `h1.GetZ() and grEl.GetZ() must by equal !`)
		assert.True(t, h1.GetT().Equals(&gTM), `h1.GetT() and grEl.T.multiply(Ed25519Field.D_Times_TWO) must by equal !`)
		if !assert.True(t, h1.Equals(h2), `h1 and h2 must by equal ! i=%d`, i) {
			assert.True(t, h1.GetX().Equals(h2.GetX()), `h1.GetX() and grEl.Y.add(grEl.GetX()) must by equal !`)
			assert.True(t, h1.GetY().Equals(h2.GetY()), `h1.GetY() and grEl.Y.subtract(grEl.GetX()) must by equal !`)
			assert.True(t, h1.GetZ().Equals(h2.GetZ()), `h1.GetZ() and grEl.GetZ() must by equal !`)

		}
	}

}

func TestPrecomputedTableContainsExpectedGroupElements(t *testing.T) {

	defer testRecover(t)
	grEl := Ed25519Group.BASE_POINT()
	// Act + Assert:
	for i := 0; i < 32; i++ {
		h := grEl.copy()
		for j := 0; j < 8; j++ {
			g, err := MathUtils.ToRepresentation(h, PRECOMPUTED)
			assert.Nil(t, err)
			assert.True(t, g.Equals(Ed25519Group.BASE_POINT().precomputedForSingle[i][j]), "iter = %d, %d", i, j)
			h = MathUtils.AddGroupElements(h, grEl)
		}

		for k := 0; k < 8; k++ {
			grEl = MathUtils.AddGroupElements(grEl, grEl)
		}
	}

}

func TestDblPrecomputedTableContainsExpectedGroupElements(t *testing.T) {

	defer testRecover(t)
	grEl := Ed25519Group.BASE_POINT()
	h := MathUtils.AddGroupElements(grEl, grEl)
	// Act + Assert:
	for i := 0; i < 8; i++ {
		g, err := MathUtils.ToRepresentation(grEl, PRECOMPUTED)
		assert.Nil(t, err)
		assert.True(t, Ed25519Group.BASE_POINT().precomputedForDouble[i].Equals(g), "iter=%d", i)
		grEl = MathUtils.AddGroupElements(grEl, h)
	}

}

func TestDblReturnsExpectedResult(t *testing.T) {

	defer testRecover(t)
	for i := 0; i < numIter; i++ {
		g := MathUtils.GetRandomGroupElement()

		h1 := g.dbl()
		h2 := MathUtils.DoubleGroupElement(g)
		assert.True(t, h2.Equals(h1), `h2 and h1 must by equal !`)
	}

}

func TestEd25519GroupElemenAdd_AddingNeutralGroupElementDoesNotChangeGroupElement(t *testing.T) {

	defer testRecover(t)
	neutral := NewEd25519GroupElementP3(
		Ed25519Field_ZERO(),
		Ed25519Field_ONE(),
		Ed25519Field_ONE(),
		Ed25519Field_ZERO())
	for i := 0; i < numIter; i++ {
		g := MathUtils.GetRandomGroupElement()

		h1 := g.add(neutral.toCached())
		h2 := neutral.add(g.toCached())
		assert.True(t, g.Equals(h1), `g and h1 must by equal ! i=%d`, i)
		assert.True(t, g.Equals(h2), `g and h2 must by equal ! i=%d`, i)
	}

}

func TestEd25519GroupElemenAddReturnsExpectedResult(t *testing.T) {

	defer testRecover(t)
	for i := 0; i < numIter; i++ {
		g1 := MathUtils.GetRandomGroupElement()
		g2 := MathUtils.GetRandomGroupElement()
		h1 := g1.add(g2.toCached())
		h2 := MathUtils.AddGroupElements(g1, g2)
		assert.True(t, h2.Equals(h1), `h2 and h1 must by equal ! i=%d`, i)
	}

}

func TestSubReturnsExpectedResult(t *testing.T) {

	defer testRecover(t)
	for i := 0; i < numIter; i++ {
		g1 := MathUtils.GetRandomGroupElement()
		g2 := MathUtils.GetRandomGroupElement()
		h1 := g1.subtract(g2.toCached())
		h2 := MathUtils.AddGroupElements(g1, MathUtils.NegateGroupElement(g2))
		assert.True(t, h2.Equals(h1), `h2 and h1 must by equal !`)
	}

}

func TestGroupElement_EqualsOnlyReturnsTrueForEquivalentObjects(t *testing.T) {

	defer testRecover(t)
	g1 := MathUtils.GetRandomGroupElement()
	g2, err := MathUtils.ToRepresentation(g1, P2)
	assert.Nil(t, err)
	g3, err := MathUtils.ToRepresentation(g1, CACHED)
	assert.Nil(t, err)
	g4, err := MathUtils.ToRepresentation(g1, P1xP1)
	assert.Nil(t, err)
	g5 := MathUtils.GetRandomGroupElement()
	// Assert
	assert.True(t, g2.Equals(g1), `g2 and g1 must by equal !`)
	assert.True(t, g3.Equals(g1), `g3 and g1 must by equal !`)
	assert.True(t, g1.Equals(g4), `g1 and g4 must by equal !`)
	assert.NotEqual(t, g1, g5, `g1 and g5 must by not equal !`)
	assert.NotEqual(t, g2, g5, `g2 and g5 must by not equal !`)
	assert.NotEqual(t, g3, g5, `g3 and g5 must by not equal !`)
	assert.NotEqual(t, g5, g4, `g5 and g4 must by not equal !`)
}

func TestEd25519GroupElementP3String_ReturnsCorrectRepresentation(t *testing.T) {

	g := NewEd25519GroupElementP3(Ed25519Field_ZERO(), Ed25519Field_ONE(), Ed25519Field_ONE(), Ed25519Field_ZERO())
	gAsString := g.String()
	expectedString := fmt.Sprintf("X=%s\nY=%s\nZ=%s\nT=%s\n",
		"0000000000000000000000000000000000000000000000000000000000000000",
		"0100000000000000000000000000000000000000000000000000000000000000",
		"0100000000000000000000000000000000000000000000000000000000000000",
		"0000000000000000000000000000000000000000000000000000000000000000")
	assert.Equal(t, gAsString, expectedString, `gAsString and expectedString must by equal !`)
}

//
func TestScalarMultiplyBasePointWithZeroReturnsNeutralElement(t *testing.T) {

	basePoint := Ed25519Group.BASE_POINT()
	g, err := basePoint.scalarMultiply(Ed25519Field.ZERO.Encode())
	assert.Nil(t, err)

	assert.True(t, Ed25519Group.ZERO_P3().Equals(g), `Ed25519Group.ZERO_P3 and g must by equal !`)
}
func TestScalarMultiplyBasePointWithOneReturnsBasePoint(t *testing.T) {

	defer testRecover(t)
	basePoint := Ed25519Group.BASE_POINT()
	g, err := basePoint.scalarMultiply(Ed25519Field.ONE.Encode())
	assert.Nil(t, err)

	assert.True(t, basePoint.Equals(g), `basePoint and g must by equal !`)
}

func TestScalarMultiplyBasePointReturnsExpectedResult(t *testing.T) {

	defer testRecover(t)
	for i := 0; i < 100; i++ {
		basePoint := Ed25519Group.BASE_POINT()
		f := MathUtils.GetRandomFieldElement()
		g, err := basePoint.scalarMultiply(f.Encode())
		assert.Nil(t, err)
		h := MathUtils.ScalarMultiplyGroupElement(basePoint, f)
		if !assert.True(t, g.Equals(h), `g and h must by equal !`) {
			assert.Equal(t, h.GetX(), g.GetX())
			assert.Equal(t, h.GetY(), g.GetY())
			assert.Equal(t, h.GetZ(), g.GetZ())
			assert.Equal(t, h.GetT(), g.GetT())
		}
	}

}

// This test is slow (~6s) due to math utils using an inferior algorithm to calculate the result.
func TestDoubleScalarMultiplyVariableTimeReturnsExpectedResult(t *testing.T) {

	defer testRecover(t)

	for i := 0; i < 50; i++ {
		basePoint := Ed25519Group.BASE_POINT()
		g := MathUtils.GetRandomGroupElement()
		g.PrecomputeForDoubleScalarMultiplication()
		f1 := MathUtils.GetRandomFieldElement()
		f2 := MathUtils.GetRandomFieldElement()
		h1, err := basePoint.doubleScalarMultiplyVariableTime(g, f2.Encode(), f1.Encode())
		assert.Nil(t, err)
		h2 := MathUtils.doubleScalarMultiplyGroupElements(basePoint, f1, g, f2)
		assert.True(t, h1.Equals(h2), `h1 and h2 must by equal !`)
	}

}

// endregion
func TestSatisfiesCurveEquationReturnsTrueForPointsOnTheCurve(t *testing.T) {

	defer testRecover(t)

	for i := 0; i < 100; i++ {
		g := MathUtils.GetRandomGroupElement()
		assert.Truef(t, g.SatisfiesCurveEquation(), `g.satisfiesCurveEquation() must be true!`)
	}

}

func TestSatisfiesCurveEquationReturnsFalseForPointsNotOnTheCurve(t *testing.T) {

	defer testRecover(t)
	for i := 0; i < 100; i++ {
		g := MathUtils.GetRandomGroupElement()

		gZ := g.Z.multiply(Ed25519Field.TWO)
		h := NewEd25519GroupElementP2(g.GetX(), g.GetY(), &gZ)
		// Assert (can only fail for 5*Z^2=1):
		assert.False(t, h.SatisfiesCurveEquation(), `h.satisfiesCurveEquation() must be false!`)
	}

}

// endregion
// region EncodedGroupElement

func TestEd25519EncodedGroupElement_CanBeCreatedFromByteArray(t *testing.T) {

	_, err := NewEd25519EncodedGroupElement(make([]byte, 32))
	assert.Nil(t, err)
}

//(expected = IllegalArgumentException.class)
func TestEd25519EncodedGroupElement_CannotBeCreatedFromArrayWithIncorrectLength(t *testing.T) {

	_, err := NewEd25519EncodedGroupElement(make([]byte, 30))
	assert.NotNil(t, err)
}

// region getRaw
func TestGetRawReturnsUnderlyingArray(t *testing.T) {

	values := make([]byte, 32) // final
	values[0] = 5
	values[6] = 15
	values[23] = 256 - 67
	encoded, err := NewEd25519EncodedGroupElement(values)
	assert.Nil(t, err)

	assert.Equal(t, values, encoded.Raw, `values and encoded.Raw must by equal !`)
}

func TestDecodePlusEncodeDoesNotAlterTheEncodedGroupElement(t *testing.T) {

	defer func() {
		err := recover()
		t.Log(err)
	}()
	for i := 0; i < numIter; i++ {
		original := MathUtils.GetRandomEncodedGroupElement()
		grEl, err := original.Decode()
		assert.Nil(t, err)
		encoded, err := grEl.Encode()
		assert.Equal(t, encoded, original, `encoded and original must by equal !`)
	}

}

// endregion
func Test_GetAffineXReturnsExpectedResult(t *testing.T) {

	defer func() {
		err := recover()
		t.Log(err)
	}()
	for i := 0; i < 1000; i++ {
		encoded, err := MathUtils.GetRandomGroupElement().Encode()
		assert.Nil(t, err)
		affineX1, err := encoded.GetAffineX() // final
		assert.Nil(t, err)

		encEl, err := encoded.Decode()
		assert.Nil(t, err)
		el, err := MathUtils.ToRepresentation(encEl, AFFINE)
		affineX2 := el.GetX() // final
		assert.Equal(t, affineX1, affineX2, `affineX1 and affineX2 must by equal !`)
	}

}

//(expected = IllegalArgumentException.class)
func TestGetAffineXThrowsIfEncodedGroupElementIsInvalid(t *testing.T) {

	g := NewEd25519GroupElementP2(Ed25519Field_ONE(), &Ed25519Field.D, Ed25519Field_ONE()) // Ed25519GroupElement
	encoded, err := g.Encode()
	assert.Nil(t, err)
	encoded.GetAffineX()
}

// @Test
func TestGetAffineYReturnsExpectedResult(t *testing.T) {

	defer func() {
		err := recover()
		t.Log(err)
	}()
	for i := 0; i < numIter; i++ {
		encoded := MathUtils.GetRandomEncodedGroupElement()
		affineY1, err := encoded.GetAffineY()
		assert.Nil(t, err)

		encEl, err := encoded.Decode()
		assert.Nil(t, err)
		el, err := MathUtils.ToRepresentation(encEl, AFFINE)
		assert.Nil(t, err)
		affineY2 := el.GetY() // final
		assert.Equal(t, affineY1, affineY2, `affineY1 and affineY2 must by equal !`)
	}

}

// @Test
func TestEqualsOnlyReturnsTrueForEquivalentObjects(t *testing.T) {

	defer testRecover(t)
	g1 := MathUtils.GetRandomEncodedGroupElement()
	g2Enc, err := g1.Decode()
	assert.Nil(t, err)
	g2, err := g2Enc.Encode()
	assert.Nil(t, err)
	g3 := MathUtils.GetRandomEncodedGroupElement()
	g4 := MathUtils.GetRandomEncodedGroupElement()
	// Assert
	assert.Equal(t, g1, g2, `g2 and g1 must by equal !`)
	assert.NotEqual(t, g1, g3, `g1 and g3 must by not equal !`)
	assert.NotEqual(t, g2, g4, `g2 and g4 must by not equal !`)
	assert.NotEqual(t, g3, g4, `g3 and g4 must by not equal !`)
}

func TestEd25519GroupP2ElementString_ReturnsCorrectRepresentation(t *testing.T) {

	encoded, err := NewEd25519GroupElementP2(Ed25519Field_ZERO(), Ed25519Field_ONE(), Ed25519Field_ONE()).Encode()
	assert.Nil(t, err)
	encodedAsString := encoded.String()
	expectedString := fmt.Sprintf("x=%s\ny=%s\n",
		"0000000000000000000000000000000000000000000000000000000000000000",
		"0100000000000000000000000000000000000000000000000000000000000000")
	assert.Equal(t, encodedAsString, expectedString, `encodedAsString and expectedString must by equal !`)
}

// endregion

func assertEquals(t *testing.T, f Ed25519FieldElement, b1 *big.Int, msgAndArgs ...interface{}) bool {

	msg := `b2.mod(%+v) and b.mod(%+v) must by equal !`
	b2 := MathUtils.BytesToBigInteger(f.Encode().Raw)

	args := []interface{}{msg, b1, b2}
	if len(msgAndArgs) > 0 {
		i := 0
		if arg, ok := msgAndArgs[0].(string); ok {
			args[0] = arg
			i = 1
		} else {
			for range msgAndArgs {
				msg += " %+v"
			}
			args[0] = msg
		}
		args = append(args, msgAndArgs[i:])
	}
	return assert.Equal(t, (&big.Int{}).Mod(b1, Ed25519Field.P), (&big.Int{}).Mod(b2, Ed25519Field.P), args...)

}

func differsOnlyByAFactorOfAFourthRootOfOne(x Ed25519FieldElement, root Ed25519FieldElement) bool {

	rootTimesI := root.multiply(Ed25519Field.I)

	return isEqualConstantTime(x.Encode().Raw, root.Encode().Raw) ||
		isEqualConstantTime(x.Encode().Raw, root.negate().Encode().Raw) ||
		isEqualConstantTime(x.Encode().Raw, rootTimesI.Encode().Raw) ||
		isEqualConstantTime(x.Encode().Raw, rootTimesI.negate().Encode().Raw)
}
