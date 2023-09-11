package grid_test

import (
	grid2 "github.com/johnfercher/taleslab/pkg/shared/grid"
	"testing"

	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
	"github.com/stretchr/testify/assert"
)

func TestSliceTerrain_PerfectCase(t *testing.T) {
	// Arrange
	sliceSize := 50
	basicElement := taleslabentities.Element{}
	world := grid2.GenerateElementGrid(100, 100, basicElement)

	// Act
	matrix := grid2.SliceTerrain(world, sliceSize)

	// Assert
	assert.Equal(t, 2, len(matrix))
	assert.Equal(t, 2, len(matrix[0]))
	assert.Equal(t, 2, len(matrix[1]))
	for _, line := range matrix {
		for _, worldMatrix := range line {
			assert.Equal(t, sliceSize, len(worldMatrix))
			for _, lineWorld := range worldMatrix {
				assert.Equal(t, sliceSize, len(lineWorld))
				for _, element := range lineWorld {
					assert.Equal(t, basicElement, element)
				}
			}
		}
	}
}

func TestSliceTerrain_ImperfectCaseInX(t *testing.T) {
	// Arrange
	sliceSize := 50
	basicElement := taleslabentities.Element{}
	world := grid2.GenerateElementGrid(120, 100, basicElement)

	// Act
	matrix := grid2.SliceTerrain(world, sliceSize)

	// Assert
	assert.Equal(t, 3, len(matrix))
	assert.Equal(t, 2, len(matrix[0]))
	assert.Equal(t, 2, len(matrix[1]))
	assert.Equal(t, sliceSize, len(matrix[0][0]))
	assert.Equal(t, sliceSize, len(matrix[1][0]))
	assert.Equal(t, 20, len(matrix[2][0]))
	assert.Equal(t, sliceSize, len(matrix[0][1]))
	assert.Equal(t, sliceSize, len(matrix[1][1]))
	assert.Equal(t, 20, len(matrix[2][1]))
	for _, line := range matrix {
		for _, worldMatrix := range line {
			for _, lineWorld := range worldMatrix {
				assert.Equal(t, sliceSize, len(lineWorld))
				for _, element := range lineWorld {
					assert.Equal(t, basicElement, element)
				}
			}
		}
	}
}

func TestSliceTerrain_ImperfectCaseInY(t *testing.T) {
	// Arrange
	sliceSize := 50
	basicElement := taleslabentities.Element{}
	world := grid2.GenerateElementGrid(100, 120, basicElement)

	// Act
	matrix := grid2.SliceTerrain(world, sliceSize)

	// Assert
	assert.Equal(t, 2, len(matrix))
	assert.Equal(t, 3, len(matrix[0]))
	assert.Equal(t, 3, len(matrix[1]))
	for _, line := range matrix {
		for _, worldMatrix := range line {
			assert.Equal(t, sliceSize, len(worldMatrix))
			for _, lineWorld := range worldMatrix {
				for _, element := range lineWorld {
					assert.Equal(t, basicElement, element)
				}
			}
		}
	}
}
