package grid_test

import (
	"github.com/johnfercher/taleslab/pkg/shared/grid"
	"testing"

	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts/elementtype"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
	"github.com/stretchr/testify/assert"
)

func TestGenerateElementGrid(t *testing.T) {
	// Act
	elementGrind := grid.GenerateElementGrid(5, 5, taleslabentities.Element{Height: 0, ElementType: elementtype.Ground})

	// Assert
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			assert.Equal(t, taleslabentities.Element{Height: 0, ElementType: elementtype.Ground}, elementGrind[i][j])
		}
	}
}