package talespirecoder

import (
	"fmt"
	"github.com/johnfercher/taleslab/internal/bytecompressor"
	"github.com/johnfercher/taleslab/pkg/taleslab/domain/consts"
	"github.com/johnfercher/taleslab/pkg/talespire/talespirecontracts"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEncoder(t *testing.T) {
	// Act
	sut := NewEncoder(nil)

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*talespirecoder.encoder", fmt.Sprintf("%T", sut))
}

func TestEncoder_EncodeIntegrationHappy(t *testing.T) {
	// Arrange
	byteCompressor := bytecompressor.New()
	sut := NewEncoder(byteCompressor)
	slab := mockSlab()

	// Act
	b64, err := sut.Encode(slab)

	// Assert
	assert.Nil(t, err)
	assert.NotEmpty(t, b64)
	assert.Equal(t, "H4sIAAAAAAAE/wAoANf/zvrO0QIAAQAAACw+d10vykRMiND5tKevr1oBAAAAAAAAAAAAAAAAAAEAAP//7YI+iCgAAAA=", b64)
}

func mockSlab() *talespirecontracts.Slab {
	return &talespirecontracts.Slab{
		MagicBytes:  consts.MagicBytes,
		Version:     2,
		AssetsCount: 1,
		Assets: []*talespirecontracts.Asset{
			{
				Id:           []byte{0x0, 0x0, 0x2c, 0x3e, 0x77, 0x5d, 0x2f, 0xca, 0x44, 0x4c, 0x88, 0xd0, 0xf9, 0xb4, 0xa7, 0xaf, 0xaf, 0x5a},
				LayoutsCount: 1,
				Layouts: []*talespirecontracts.Bounds{
					{
						Coordinates: &talespirecontracts.Vector3d{
							X: 0,
							Y: 0,
							Z: 0,
						},
						Rotation: 0,
					},
				},
			},
		},
	}
}
