package taleslabmappers

import (
	"testing"

	"github.com/johnfercher/taleslab/mocks"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts"
	"github.com/stretchr/testify/assert"
)

func TestEntitySlabFromTaleSpire(t *testing.T) {
	// Arrange
	taleSpireSlab := mocks.CreateTaleSpireSlab()

	// Act
	assets := AssetsFromTaleSpireSlab(taleSpireSlab)

	// Assert
	assert.NotNil(t, assets)
	asset := assets[0]

	assert.Equal(t, []byte{
		0x0, 0x0, 0x2c, 0x3e, 0x77, 0x5d, 0x2f, 0xca, 0x44, 0x4c,
		0x88, 0xd0, 0xf9, 0xb4, 0xa7, 0xaf, 0xaf, 0x5a,
	}, asset.ID)

	assert.Equal(t, 0, asset.Rotation)
	assert.Equal(t, 0, asset.Coordinates.X)
	assert.Equal(t, 0, asset.Coordinates.Y)
	assert.Equal(t, 0, asset.Coordinates.Z)
}

func TestTaleSpireSlabFromEntity(t *testing.T) {
	// Arrange
	assets := mocks.CreateAssets()

	// Act
	taleSpireSlab := TaleSpireSlabFromAssets(assets)

	// Assert
	assert.NotNil(t, taleSpireSlab)
	assert.Equal(t, taleslabconsts.MagicBytes, taleSpireSlab.MagicBytes)
	assert.Equal(t, int16(2), taleSpireSlab.Version)
	assert.Equal(t, int16(1), taleSpireSlab.AssetsCount)
	assert.Equal(t, taleSpireSlab.AssetsCount, int16(len(taleSpireSlab.Assets)))
	assert.Equal(t, []byte{
		0x0, 0x0, 0x2c, 0x3e, 0x77, 0x5d, 0x2f, 0xca, 0x44, 0x4c,
		0x88, 0xd0, 0xf9, 0xb4, 0xa7, 0xaf, 0xaf, 0x5a,
	}, taleSpireSlab.Assets[0].Id)
	assert.Equal(t, int16(1), taleSpireSlab.Assets[0].LayoutsCount)
	assert.Equal(t, taleSpireSlab.Assets[0].LayoutsCount, int16(len(taleSpireSlab.Assets[0].Layouts)))
	assert.Equal(t, uint16(0), taleSpireSlab.Assets[0].Layouts[0].Rotation)
	assert.Equal(t, uint16(0), taleSpireSlab.Assets[0].Layouts[0].Coordinates.X)
	assert.Equal(t, uint16(0), taleSpireSlab.Assets[0].Layouts[0].Coordinates.Y)
	assert.Equal(t, uint16(0), taleSpireSlab.Assets[0].Layouts[0].Coordinates.Z)
}
