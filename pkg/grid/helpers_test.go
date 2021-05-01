package grid_test

import (
	"github.com/johnfercher/taleslab/pkg/grid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateElementGrid(t *testing.T) {
	// Act
	elementGrind := grid.GenerateElementGrid(5, 5, grid.Element{0, grid.GroundType})

	// Assert
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			assert.Equal(t, grid.Element{0, grid.GroundType}, elementGrind[i][j])
		}
	}
}
