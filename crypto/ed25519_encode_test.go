package crypto

import (
	"github.com/stretchr/testify/assert"
	"math/big"
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

	assert.Equal(t, values[:], f.Raw, `values and f.getRaw() must by equal !`)
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

//
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
		if assertEquals(t, f3, (&big.Int{}).Mul(b1, b2)) {
			t.Log(b1, b2, i)
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

	assert.Equal(t, values[:], f.Raw, `values and f.getRaw() must by equal !`)
}
func TestDecodePlusEncodeDoesNotAlterTheEncodedFieldElement(t *testing.T) {

	for i := 0; i < numIter; i++ {
		original := MathUtils.GetRandomEncodedFieldElement(32)
		encoded := original.Decode().Encode()
		assert.Equal(t, original, encoded, `encoded and original must by equal !`)
	}

}

// endregion
// region modulo group order arithmetic
func TestEd25519EncodedFieldElement_ModQReturnsExpectedResult(t *testing.T) {

	for i := 0; i < numIter; i++ {
		encoded := &Ed25519EncodedFieldElement{Ed25519Field.ZERO_LONG, MathUtils.getRandomByteArray(64)}
		reduced1 := encoded.modQ()
		b1 := MathUtils.BytesToBigInteger(reduced1.Raw)

		reduced2 := MathUtils.ReduceModGroupOrder(encoded)

		assert.Equal(t, -1, b1.Cmp(Ed25519Field.P), `MathUtils.toBigInteger(reduced1).compareTo(Ed25519Field.P) and -1 must by equal !`)
		assert.Equal(t, 1, b1.Cmp(BigInteger_ONE()), `MathUtils.toBigInteger(reduced1).compareTo(Newuint64("-1")) and 1 must by equal !`)
		assert.Equal(t, reduced2.Raw, reduced1.Raw, `reduced1 and reduced2 must by equal ! Iter = %d`, i)
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
		assert.Equal(t, result1, result2, `result1 and result2 must by equal ! Iter %d`, i)
	}

}

func TestEncodeReturnsCorrectByteArrayForSimpleFieldElements(t *testing.T) {

	t1 := make([]intRaw, 10)
	t2 := make([]intRaw, 10)
	t2[0] = 1
	fieldElement1, err := NewEd25519FieldElement(t1)

	assert.Nil(t, err)
	fieldElement2, err := NewEd25519FieldElement(t2)
	assert.Nil(t, err)
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
			test[j] = MathUtils.getRandomIntRaw() - (1 << 27)
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

		values := MathUtils.getRandomByteArray(32)
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

func TestEqualsOnlyReturnsTrueForEquivalentObjects(t *testing.T) {

	encoded1 := MathUtils.GetRandomEncodedFieldElement(32)
	encoded2 := encoded1.Decode().Encode()
	encoded3 := MathUtils.GetRandomEncodedFieldElement(32)
	encoded4 := MathUtils.GetRandomEncodedFieldElement(32)
	// Assert:
	assert.Equal(t, encoded1, encoded2, `encoded1 and encoded2 must by equal !`)
	assert.NotEqual(t, encoded1, encoded3, `encoded1 and encoded3 must by not equal !`)
	assert.NotEqual(t, encoded1, encoded4, `encoded1 and encoded4 must by not equal !`)
	assert.NotEqual(t, encoded3, encoded4, `encoded3 and encoded4 must by not equal !`)
}
func TestEd25519EncodedGroupElement_GetAffineX(t *testing.T) {

	for i := 0; i < numIter; i++ {
		enc, err := MathUtils.GetRandomGroupElement()
		if err != nil {
			t.Fatal(err)
		}
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

	for i := 0; i < numIter; i++ {
		f1 := MathUtils.GetRandomFieldElement()
		b1 := MathUtils.BytesToBigInteger(f1.Encode().Raw)
		f2 := f1.square()

		assertEquals(t, f2, b1.Mul(b1, b1))
	}
}
func TestSqrtReturnsCorrectResult(t *testing.T) {

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
			for _ = range msgAndArgs {
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
