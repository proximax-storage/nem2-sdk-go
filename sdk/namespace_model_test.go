// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package sdk

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestGenerateNamespacePath_GeneratesCorrectWellKnownRootPath(t *testing.T) {
	ids, err := GenerateNamespacePath("nem")
	assert.Nil(t, err)

	assert.Equal(t, len(ids), 1, `ids.size() and 1 must by equal !`)

	assert.Equal(t, big.NewInt(-8884663987180930485).Int64(), ids[0].Int64())
}

// @Test
func TestNamespacePath_GeneratesCorrectWellKnownChildPath(t *testing.T) {
	ids, err := GenerateNamespacePath("nem.xem")
	assert.Nil(t, err)
	assert.Equal(t, len(ids), 2, `ids.size() and 2 must by equal !`)

	assert.Equal(t, big.NewInt(-8884663987180930485).Int64(), ids[0].Int64())
	assert.Equal(t, big.NewInt(-3087871471161192663).Int64(), ids[1].Int64(), `NewBigInteger(testNewXemBigInt) and ids.get(1) must by equal !`)
}

// @Test
func TestNamespacePathSupportsMultiLevelNamespaces(t *testing.T) {
	ids := make([]*big.Int, 3)
	var err error
	ids[0], err = generateId("foo", big.NewInt(0))
	assert.Nil(t, err)
	ids[1], err = generateId("bar", ids[0])
	assert.Nil(t, err)
	ids[2], err = generateId("baz", ids[1])
	assert.Nil(t, err)
	ids1, err := GenerateNamespacePath("foo.bar.baz")
	assert.Nil(t, err)
	assert.Equal(t, ids1, ids, `GenerateNamespacePath("foo.bar.baz") and ids must by equal !`)
}

// @Test
func TestNamespacePathRejectsNamesWithTooManyParts(t *testing.T) {
	_, err := GenerateNamespacePath("a.b.c.d")
	assert.Equal(t, ErrNamespaceToManyPart, err, "Err 'too many parts' must return")
	_, err = GenerateNamespacePath("a.b.c.d.e")
	assert.Equal(t, ErrNamespaceToManyPart, err, "Err 'too many parts' must return")

}

// @Test
func TestMosaicIdGeneratesCorrectWellKnowId(t *testing.T) {
	id, err := generateMosaicId("nem", "xem")
	assert.Nil(t, err)
	assert.Equal(t, big.NewInt(-3087871471161192663).Int64(), id.Int64())
}

// @Test
func TestMosaicIdSupportMultiLevelMosaics(t *testing.T) {
	var err error
	ids := make([]*big.Int, 4)
	ids[0], err = generateId("foo", big.NewInt(0))
	assert.Nil(t, err)
	ids[1], err = generateId("bar", ids[0])
	assert.Nil(t, err)
	ids[2], err = generateId("baz", ids[1])
	assert.Nil(t, err)
	ids[3], err = generateId("tokens", ids[2])
	assert.Nil(t, err)
	ids1, err := generateMosaicId("foo.bar.baz", "tokens")
	assert.Equal(t, ids1.Int64(), ids[3].Int64(), `GenerateMosaicId("foo.bar.baz" and "tokens" must by equal !`)
}
