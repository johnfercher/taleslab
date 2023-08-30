package talespirecoder

import (
	"fmt"
	"github.com/johnfercher/taleslab/internal/helper/bytecompressor"
	"github.com/johnfercher/taleslab/mocks"
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
	slab := mocks.CreateTaleSpireSlab()

	// Act
	b64, err := sut.Encode(slab)

	// Assert
	assert.Nil(t, err)
	assert.NotEmpty(t, b64)
	assert.Equal(t, "H4sIAAAAAAAE/wAoANf/zvrO0QIAAQAAACw+d10vykRMiND5tKevr1oBAAAAAAAAAAAAAAAAAAEAAP//7YI+iCgAAAA=", b64)
}
