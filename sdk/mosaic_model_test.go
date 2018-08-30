package sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMosaicId(t *testing.T) {
	id, err := NewMosaicId(nil, "nem:xem")
	assert.Nil(t, err)

	t.Log(id)
}
