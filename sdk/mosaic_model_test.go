package sdk

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestNewMosaicId(t *testing.T) {
	mosaicId, err := NewMosaicId(nil, "nem:xem")
	assert.Nil(t, err)

	assert.Equal(t, big.NewInt(-3087871471161192663).Int64(), mosaicId.Id.Int64())
	assert.Equal(t, mosaicId.FullName, "nem:xem")
}

func TestNewMosaicIdFromIdViaConstructor(t *testing.T) {
	mosaicId, err := NewMosaicId(big.NewInt(-8884663987180930485), "")
	assert.Nil(t, err)

	assert.Equal(t, big.NewInt(-8884663987180930485), mosaicId.Id)
	assert.False(t, mosaicId.FullName != "")
}

//
func TestNewMosaic_ShouldCompareMosaicIdsForEquality(t *testing.T) {
	mosaicId, err := NewMosaicId(big.NewInt(-8884663987180930485), "")
	assert.Nil(t, err)

	mosaicId2, err := NewMosaicId(big.NewInt(-8884663987180930485), "")
	assert.Nil(t, err)
	assert.Equal(t, mosaicId, mosaicId2)
}
