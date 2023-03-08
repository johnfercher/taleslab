package grid

import (
	"fmt"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
	"testing"
)

func TestMountainGenerator(t *testing.T) {
	x := 50
	y := 50

	mountain := MountainGenerator(x, y, 5.0, 5)

	PrintTypes(mountain)
}

func TestGetSliceInOffset(t *testing.T) {
	grid := GenerateElementGrid(9, 10, taleslabentities.Element{Height: 1, ElementType: taleslabconsts.GroundType})

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			grid[i][j].Height = i + j
		}
	}

	grid.Print()
	fmt.Println()

	slices := SliceTerrain(grid, 5)
	for _, slice := range slices {
		slice.Print()
		fmt.Println()
	}

}
