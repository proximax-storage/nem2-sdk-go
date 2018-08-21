package crypto

import (
	"github.com/agl/ed25519/edwards25519"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

const numIter = 1000

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
func TestGetRawReturnsUnderlyingArray(t *testing.T) {

	values := [10]intRaw{}
	values[0] = 5
	values[6] = 15
	values[8] = -67
	f, _ := NewEd25519FieldElement(values[:])

	assert.Equal(t, values[:], f.Raw, `values and f.getRaw() must by equal !`)
}
func TestDecodePlusEncodeDoesNotAlterTheEncodedFieldElement(t *testing.T) { /* public  */

	// Act:
	for i := 0; i < numIter; i++ {
		// Arrange:
		original := MathUtils.GetRandomEncodedFieldElement(32)
		encoded := original.Decode().Encode()
		// Assert:
		assert.Equal(t, original, encoded, `encoded and original must by equal !`)
	}

}

// endregion
// region modulo group order arithmetic
// @Test
func TestModQReturnsExpectedResult(t *testing.T) { /* public  */

	for i := 0; i < numIter; i++ {
		// Arrange:
		encoded := &Ed25519EncodedFieldElement{Ed25519Field.ZERO_LONG, MathUtils.getRandomByteArray(64)}
		// Act:
		reduced1 := encoded.modQ()
		b1 := MathUtils.BytesToBigInteger(reduced1.Raw)

		reduced2 := MathUtils.ReduceModGroupOrder(encoded)
		// Assert:
		assert.Equal(t, -1, b1.Cmp(Ed25519Field.P), `MathUtils.toBigInteger(reduced1).compareTo(Ed25519Field.P) and -1 must by equal !`)
		assert.Equal(t, 1, b1.Cmp(BigInteger_ONE()), `MathUtils.toBigInteger(reduced1).compareTo(Newuint64("-1")) and 1 must by equal !`)
		assert.Equal(t, reduced2, reduced1, `reduced1 and reduced2 must by equal !`)
	}

}
func TestMultiplyAndAddModQReturnsExpectedResult(t *testing.T) { /* public  */

	for i := 0; i < numIter; i++ {
		// Arrange:
		encoded1 := MathUtils.GetRandomEncodedFieldElement(32)
		encoded2 := MathUtils.GetRandomEncodedFieldElement(32)
		encoded3 := MathUtils.GetRandomEncodedFieldElement(32)
		// Act:
		result1 := encoded1.multiplyAndAddModQ(encoded2, encoded3)
		b1 := MathUtils.BytesToBigInteger(result1.Raw)
		result2 := MathUtils.multiplyAndAddModGroupOrder(encoded1, encoded2, encoded3)
		// Assert:
		assert.Equal(t, -1, b1.Cmp(Ed25519Field.P), `MathUtils.toBigInteger(result1).compareTo(Ed25519Field.P) and -1 must by equal !`)
		assert.Equal(t, 1, b1.Cmp(big.NewInt(-1)), `MathUtils.toBigInteger(result1).compareTo(Newuint64("-1")) and 1 must by equal !`)
		assert.Equal(t, result1, result2, `result1 and result2 must by equal !`)
	}

}

// TestAddReturnsCorrectResult test Ed25519FieldElement summaring corection
func TestAddReturnsCorrectResult(t *testing.T) {

	for i := 0; i < numIter; i++ {
		// Arrange:
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

func TestMultiplyReturnsCorrectResult(t *testing.T) {

	for i := 0; i < numIter; i++ {
		f1 := MathUtils.GetRandomFieldElement()
		f2 := MathUtils.GetRandomFieldElement()
		b1 := MathUtils.BytesToBigInteger(f1.Encode().Raw)
		b2 := MathUtils.BytesToBigInteger(f2.Encode().Raw)

		f3 := f1.multiply(f2)
		assertEquals(t, f3, b1.Mul(b1, b2), b1, b2)
		f4 := f2.multiply(f1)
		assert.Equal(t, f3, f4)
		var h, f, g edwards25519.FieldElement
		for i, val := range f1.Raw {
			f[i] = int32(val)
		}
		for i, val := range f2.Raw {
			g[i] = int32(val)
		}
		// agl method use
		edwards25519.FeMul(&h, &f, &g)

		f4 = &Ed25519FieldElement{make([]intRaw, len(h))}
		for i, val := range h {
			f4.Raw[i] = intRaw(val)
		}

		assertEquals(t, f4, b1, f, g)
		assert.Equal(t, f3, f4, "equal two methods")
	}

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
		assert.Equal(t, sqrt, fraction, `fraction and sqrt must by not equal !`)
		assert.Equal(t, sqrt.square(), fraction.square(), `fraction.square().square() and sqrt.square().square() must by not equal !`)
		assert.Equal(t, sqrt.square().square(), fraction.square().square(), `fraction.square().square() and sqrt.square().square() must by not equal !`)
		// (u / v) == +-1 * sqrt(u^2 / v^2) or (u / v) == +-i * sqrt(u^2 / v^2)
		assert.Equal(t, true, differsOnlyByAFactorOfAFourthRootOfOne(fraction, sqrt))
		break
	}
}

func assertEquals(t *testing.T, f *Ed25519FieldElement, b1 *big.Int, msgAndArgs ...interface{}) {

	msg := `b2.mod(Ed25519Field.P) and b.mod(Ed25519Field.P) must by equal ! %d = %d`
	b2 := MathUtils.BytesToBigInteger(f.Encode().Raw)

	args := []interface{}{msg, b1.Uint64(), b2.Uint64()}
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
	assert.Equal(t, (&big.Int{}).Mod(b1, Ed25519Field.P), b2.Mod(b2, Ed25519Field.P), args...)

}
func differsOnlyByAFactorOfAFourthRootOfOne(x *Ed25519FieldElement, root *Ed25519FieldElement) bool {

	rootTimesI := root.multiply(Ed25519Field.I)

	return isEqualConstantTime(x.Encode().Raw, root.Encode().Raw) ||
		isEqualConstantTime(x.Encode().Raw, root.negate().Encode().Raw) ||
		isEqualConstantTime(x.Encode().Raw, rootTimesI.Encode().Raw) ||
		isEqualConstantTime(x.Encode().Raw, rootTimesI.negate().Encode().Raw)
}
