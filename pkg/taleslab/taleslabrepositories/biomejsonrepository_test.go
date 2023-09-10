package taleslabrepositories_test

import (
	"fmt"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts/biometype"

	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabrepositories"

	"github.com/stretchr/testify/assert"
)

func TestNewBiomeRepository_WhenCantCreate_ShouldReturnError(t *testing.T) {
	// Act
	sut, err := taleslabrepositories.NewBiomeRepository()

	// Assert
	assert.NotNil(t, err)
	assert.Nil(t, sut)
}

func TestNewBiomeRepository_WhenCanCreate_ShouldReturnRepository(t *testing.T) {
	// Act
	sut, err := taleslabrepositories.NewBiomeRepository(buildBiomePath())

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, sut)
	assert.Equal(t, "*taleslabrepositories.biomeJSONRepository", fmt.Sprintf("%T", sut))
}

func TestBiomeJSONRepository_GetBiomes(t *testing.T) {
	// Arrange
	sut, _ := taleslabrepositories.NewBiomeRepository(buildBiomePath())

	// Act
	biomes := sut.GetBiomes()

	// Assert
	assert.NotNil(t, biomes)
	assert.Equal(t, len(getMappedBiomes()), len(biomes), "different quantity mappeds x map returned")
}

func TestAssetLoader_GetProp(t *testing.T) {
	// Arrange
	mappedBiomes := getMappedBiomes()
	sut, _ := taleslabrepositories.NewBiomeRepository(buildBiomePath())

	// Act & Assert
	for key := range mappedBiomes {
		assert.NotNil(t, sut.GetBiome(biometype.BiomeType(key)), fmt.Sprintf("biome not loaded %s", key))
	}

	loadedProps := sut.GetBiomes()
	for key := range loadedProps {
		assert.True(t, mappedBiomes[string(key)], fmt.Sprintf("biome not mapped %s", key))
	}
}

func buildBiomePath() string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}

	dir = strings.Replace(dir, "/pkg/taleslab/taleslabrepositories", "", 1)
	return path.Join(dir, "/configs/biomes.json")
}

func getMappedBiomes() map[string]bool {
	return map[string]bool{
		"beach":              true,
		"swamp":              true,
		"temperate_forest":   true,
		"desert":             true,
		"tundra":             true,
		"subtropical_forest": true,
		"dead_forest":        true,
		"lava":               true,
	}
}
