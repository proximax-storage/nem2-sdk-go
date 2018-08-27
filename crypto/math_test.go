package crypto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMathUtils_AddGroupElements_Neutral(t *testing.T) {
	neutral := NewEd25519GroupElementP3(
		Ed25519Field_ZERO(),
		Ed25519Field_ONE(),
		Ed25519Field_ONE(),
		Ed25519Field_ZERO())

	for i := 0; i < 0; i++ {
		g := MathUtils.GetRandomGroupElement()
		h1 := MathUtils.AddGroupElements(g, neutral)
		h2 := MathUtils.AddGroupElements(neutral, g)
		assert.True(t, g.Equals(h1), `g and h1 must by equal !`)
		assert.True(t, g.Equals(h2), `g and h2 must by equal !`)
	}
}

//Simple test for verifying that the MathUtils code works as expected.
func TestMathUtilsWorkAsExpected(t *testing.T) {

	for i := 0; i < numIter; i++ {
		g := MathUtils.GetRandomGroupElement()
		// P3 -> P2.
		h, err := MathUtils.ToRepresentation(g, P2)
		assert.Nil(t, err)

		assert.True(t, h.Equals(g), `h and g must by equal !`)
		// P3 -> P1xP1.
		h, err = MathUtils.ToRepresentation(g, P1xP1)
		assert.Nil(t, err)
		assert.True(t, g.Equals(h), `g and h must by equal !`)
		// P3 -> CACHED.
		h, err = MathUtils.ToRepresentation(g, CACHED)
		assert.Nil(t, err)
		assert.True(t, h.Equals(g), `h and g must by equal !`)
		// P3 -> P2 -> P3.
		g, err = MathUtils.ToRepresentation(g, P2)
		assert.Nil(t, err)
		h, err = MathUtils.ToRepresentation(g, P3)
		assert.Nil(t, err)
		assert.True(t, g.Equals(h), `g and h must by equal !`)
		// P3 -> P2 -> P1xP1.
		g, err = MathUtils.ToRepresentation(g, P2)
		assert.Nil(t, err)
		h, err = MathUtils.ToRepresentation(g, P1xP1)
		assert.Nil(t, err)

		assert.True(t, g.Equals(h), `g and h must by equal !`)
	}
}
func TestMathUtils_ScalarMultiplyGroupElement(t *testing.T) {

	for i := 0; i < 10; i++ {
		g := MathUtils.GetRandomGroupElement()
		h := MathUtils.ScalarMultiplyGroupElement(g, Ed25519Field.ZERO)
		assert.True(t, Ed25519Group.ZERO_P3().Equals(h), `Ed25519Group.ZERO_P3 and h must by equal !`)
	}

}

func TestMath_PrecomputedTableContainsExpectedGroupElements(t *testing.T) {

	defer testRecover(t)
	grEl := Ed25519Group.BASE_POINT()
	for i := 0; i < numIter; i++ {
		g, err := MathUtils.ToRepresentation(grEl, PRECOMPUTED)
		assert.Nil(t, err)
		//grEl.precomputedForSingle[0][0] = &Ed25519GroupElement{
		//	PRECOMPUTED,
		//	&Ed25519FieldElement{[]intRaw{25967493, -14356035, 29566456, 3660896, -12694345, 4014787, 27544626, -11754271, -6079156, 2047605}},
		//	&Ed25519FieldElement{[]intRaw{-12545711, 934262, -2722910, 3049990, -727428, 9406986, 12720692, 5043384, 19500929, -15469378}},
		//	&Ed25519FieldElement{[]intRaw{-8738181, 4489570, 9688441, -14785194, 10184609, -12363380, 29287919, 11864899, -24514362, -4438546}},
		//nil,
		//nil, nil,
		//}
		assert.True(t, grEl.precomputedForSingle[0][0].Equals(g), "iter = %d", i)
	}

}

func TestMathUtils_ReduceModGroupOrder(t *testing.T) {
	defer testRecover(t)

	for i := 0; i < numIter; i++ {
		encoded := MathUtils.GetRandomEncodedFieldElement(64)
		reduced1 := encoded.modQ()
		reduced2 := MathUtils.ReduceModGroupOrder(encoded)
		assert.True(t, reduced2.Equals(reduced1), "iter = %d", i)
	}

}
