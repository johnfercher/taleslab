package talespirecoder

import (
	"errors"
	"fmt"
	"github.com/johnfercher/taleslab/internal/bytecompressor"
	"github.com/johnfercher/taleslab/mocks"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestNewDecoder(t *testing.T) {
	// Act
	sut := NewDecoder(nil)

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*talespirecoder.decoder", fmt.Sprintf("%T", sut))
}

func TestDecoder_Decode_WhenCannotDecompress(t *testing.T) {
	// Arrange
	b64 := mocks.CreateBase64Slab()
	errToReturn := errors.New("cannot_decompress")

	byteCompressor := &mocks.ByteCompressor{}
	byteCompressor.On("BufferFromBase64", mock.Anything).Return(nil, errToReturn)

	sut := NewDecoder(byteCompressor)

	// Act
	slab, err := sut.Decode(b64)

	// Assert
	assert.NotNil(t, err)
	assert.Equal(t, errToReturn, err)
	assert.Nil(t, slab)
	byteCompressor.AssertNumberOfCalls(t, "BufferFromBase64", 1)
	byteCompressor.AssertCalled(t, "BufferFromBase64", b64)
}

func TestDecoder_Decode_IntegrationHappy(t *testing.T) {
	// Arrange
	b64 := mocks.CreateBase64Slab()

	byteCompressor := bytecompressor.New()

	sut := NewDecoder(byteCompressor)

	// Act
	slab, err := sut.Decode(b64)

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, slab)
	assert.Equal(t, taleslabconsts.MagicBytes, slab.MagicBytes)
	assert.Equal(t, int16(2), slab.Version)
	assert.Equal(t, int16(1), slab.AssetsCount)
	assert.Equal(t, slab.AssetsCount, int16(len(slab.Assets)))
	assert.Equal(t, []byte{0x0, 0x0, 0x2c, 0x3e, 0x77, 0x5d, 0x2f, 0xca, 0x44, 0x4c,
		0x88, 0xd0, 0xf9, 0xb4, 0xa7, 0xaf, 0xaf, 0x5a}, slab.Assets[0].Id)
	assert.Equal(t, int16(1), slab.Assets[0].LayoutsCount)
	assert.Equal(t, slab.Assets[0].LayoutsCount, int16(len(slab.Assets[0].Layouts)))
	assert.Equal(t, uint16(0), slab.Assets[0].Layouts[0].Rotation)
	assert.Equal(t, uint16(0), slab.Assets[0].Layouts[0].Coordinates.X)
	assert.Equal(t, uint16(0), slab.Assets[0].Layouts[0].Coordinates.Y)
	assert.Equal(t, uint16(0), slab.Assets[0].Layouts[0].Coordinates.Z)
}
