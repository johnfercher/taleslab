package grid_test

import (
	"testing"

	"github.com/johnfercher/taleslab/pkg/grid"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts/elementtype"
	"github.com/stretchr/testify/assert"
)

func TestGenerateTerrain_WhenForceBaseLandIsTrue(t *testing.T) {
	// Arrange
	forceBaseLand := true
	minHeight := 1
	maxX := 100
	maxY := 100
	xFrequency := 2.0
	yFrequency := 2.0
	gain := 5.0

	// Act
	terrain := grid.GenerateTerrain(maxX, maxY, xFrequency, yFrequency, gain, minHeight, forceBaseLand)

	// Assert
	assert.Equal(t, maxX, len(terrain))
	assert.Equal(t, maxY, len(terrain[0]))
	for _, line := range terrain {
		for _, element := range line {
			assert.Equal(t, elementtype.Ground, element.ElementType)
			assert.LessOrEqual(t, minHeight, element.Height)
		}
	}
}

func TestGenerateTerrain_WhenForceBaseLandIsFalse(t *testing.T) {
	// Arrange
	forceBaseLand := false
	minHeight := 1
	maxX := 100
	maxY := 100
	xFrequency := 2.0
	yFrequency := 2.0
	gain := 5.0

	// Act
	terrain := grid.GenerateTerrain(maxX, maxY, xFrequency, yFrequency, gain, minHeight, forceBaseLand)

	// Assert
	assert.Equal(t, maxX, len(terrain))
	assert.Equal(t, maxY, len(terrain[0]))
	for _, line := range terrain {
		for _, element := range line {
			groundOrWater := element.ElementType == elementtype.Ground || element.ElementType == elementtype.Water
			assert.True(t, groundOrWater)
			assert.LessOrEqual(t, minHeight, element.Height)
		}
	}
}

func TestGenerateMountain(t *testing.T) {
	// Arrange
	maxX := 10
	maxY := 10
	gain := 5.0
	minHeight := 5

	// Act
	mountain := grid.GenerateMountain(maxX, maxY, gain, minHeight)

	// Assert
	assert.Equal(t, maxX, len(mountain))
	assert.Equal(t, maxY, len(mountain[0]))
	for _, line := range mountain {
		for _, element := range line {
			noneOrMountain := element.ElementType == elementtype.None || element.ElementType == elementtype.Mountain
			assert.True(t, noneOrMountain)
			if element.ElementType == elementtype.None {
				assert.Equal(t, 0, element.Height)
			} else {
				assert.LessOrEqual(t, minHeight, element.Height)
			}
		}
	}
}

func TestAppendTerrainRandomly(t *testing.T) {
	// Arrange
	forceBaseLand := true
	minHeight := 1
	maxX := 100
	maxY := 100
	xFrequency := 2.0
	yFrequency := 2.0
	gain := 5.0

	terrain := grid.GenerateTerrain(maxX, maxY, xFrequency, yFrequency, gain, minHeight, forceBaseLand)

	mountainMaxX := 10
	mountainMaxY := 10
	mountainGain := 5.0
	mountainMinHeight := 5

	mountain := grid.GenerateMountain(mountainMaxX, mountainMaxY, mountainGain, mountainMinHeight)

	// Act
	newTerrain := grid.AppendTerrainRandomly(terrain, mountain)

	// Assert
	assert.Equal(t, maxX, len(newTerrain))
	assert.Equal(t, maxY, len(newTerrain[0]))
	for _, line := range newTerrain {
		for _, element := range line {
			groundOrMountain := element.ElementType == elementtype.Ground || element.ElementType == elementtype.Mountain
			assert.True(t, groundOrMountain)
			assert.LessOrEqual(t, minHeight, element.Height)
		}
	}
}
