package grid_test

import (
	"github.com/johnfercher/taleslab/pkg/grid"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateElementGrid(t *testing.T) {
	// Acts
	elementGrind := grid.GenerateElementGrid(5, 5, taleslabentities.Element{Height: 0, ElementType: taleslabconsts.GroundType})

	// Assert
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			assert.Equal(t, taleslabentities.Element{Height: 0, ElementType: taleslabconsts.GroundType}, elementGrind[i][j])
		}
	}
}
