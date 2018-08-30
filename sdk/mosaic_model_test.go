package sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMosaicId(t *testing.T) {
	mosaicId, err := NewMosaicId(nil, "nem:xem")
	assert.Nil(t, err)

	assert.Equal(t, testIdGenerateNEM_XEMBigInt.Bytes(), mosaicId.Id.Bytes())
	assert.Equal(t, mosaicId.FullName, "nem:xem")
}

func TestNewMosaicIdFromIdViaConstructor(t *testing.T) {
	mosaicId, err := NewMosaicId(testIdGenerateNEMBigInt, "")
	assert.Nil(t, err)

	assert.Equal(t, testIdGenerateNEMBigInt, mosaicId.Id)
	assert.False(t, mosaicId.FullName != "")
}

//
func TestNewMosaic_ShouldCompareMosaicIdsForEquality(t *testing.T) {
	mosaicId, err := NewMosaicId(testIdGenerateNEMBigInt, "")
	assert.Nil(t, err)

	mosaicId2, err := NewMosaicId(testIdGenerateNEMBigInt, "")
	assert.Nil(t, err)
	assert.Equal(t, mosaicId, mosaicId2)
}
