package crypto

import (
	"github.com/agl/ed25519/edwards25519"
	"github.com/proximax-storage/nem2-sdk-go/utils"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

const numIter = 1000

func TestIsNonZeroReturnsFalseIfFieldElementIsZero(t *testing.T) {

	test := [10]intRaw{}
	f, _ := NewEd25519FieldElement(test[:])

	if f.IsNonZero() {
		t.Error("f.isNonZero() must be true!")
	}
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
		b1 := utils.BytesToBigInteger(f1.Encode().Raw)
		b2 := utils.BytesToBigInteger(f2.Encode().Raw)

		f3 := f1.subtract(f2)

		assertEquals(t, f3, b1.Sub(b1, b2))
	}

}

//
func TestNegateReturnsCorrectResult(t *testing.T) {

	for i := 0; i < numIter; i++ {
		f1 := MathUtils.GetRandomFieldElement()
		b1 := utils.BytesToBigInteger(f1.Encode().Raw)

		f2 := f1.negate()

		assertEquals(t, f2, b1.Neg(b1))
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
		var h, f, g edwards25519.FieldElement
		for i, val := range f1.Raw {
			f[i] = int32(val)
		}
		for i, val := range f2.Raw {
			g[i] = int32(val)
		}
		// agl method use
		edwards25519.FeMul(&h, &f, &g)

		f4 := &Ed25519FieldElement{make([]intRaw, len(h))}
		for i, val := range h {
			f4.Raw[i] = intRaw(val)
		}

		assertEquals(t, f4, b1.Mul(b1, b2), f, g)
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
func assertEquals(t *testing.T, f *Ed25519FieldElement, b *big.Int, msgAndArgs ...interface{}) {

	msg := `b2.mod(Ed25519Field.P) and b.mod(Ed25519Field.P) must by equal ! %d = %d`
	b2 := MathUtils.BytesToBigInteger(f.Encode().Raw)
	args := []interface{}{msg, b.Uint64(), b2.Uint64()}
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
	assert.Equal(t, b.Mod(b, Ed25519Field.P), b2.Mod(b2, Ed25519Field.P), args...)

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
		assert.Equal(t, sqrt, fraction, `fraction.square().square() and sqrt.square().square() must by not equal !`)
		assert.Equal(t, sqrt.square(), fraction.square(), `fraction.square().square() and sqrt.square().square() must by not equal !`)
		assert.Equal(t, sqrt.square().square(), fraction.square().square(), `fraction.square().square() and sqrt.square().square() must by not equal !`)
		// (u / v) == +-1 * sqrt(u^2 / v^2) or (u / v) == +-i * sqrt(u^2 / v^2)
		assert.Equal(t, true, differsOnlyByAFactorOfAFourthRootOfOne(fraction, sqrt))
		break
	}
}

func differsOnlyByAFactorOfAFourthRootOfOne(x *Ed25519FieldElement, root *Ed25519FieldElement) bool {

	rootTimesI := root.multiply(Ed25519Field.I)

	return isEqualConstantTime(x.Encode().Raw, root.Encode().Raw) ||
		isEqualConstantTime(x.Encode().Raw, root.negate().Encode().Raw) ||
		isEqualConstantTime(x.Encode().Raw, rootTimesI.Encode().Raw) ||
		isEqualConstantTime(x.Encode().Raw, rootTimesI.negate().Encode().Raw)
}
