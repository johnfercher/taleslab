package grid_test

import (
	"github.com/johnfercher/taleslab/pkg/grid"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
	"testing"
)

func TestDigRiver(t *testing.T) {
	world := grid.GenerateElementGrid(30, 30, taleslabentities.Element{
		1,
		taleslabconsts.Ground,
	})

	world = grid.DigRiver(world)

	world.Print()
}

func TestDigRiver2(t *testing.T) {
	world := grid.GenerateElementGrid(5, 5, taleslabentities.Element{
		1,
		taleslabconsts.Ground,
	})

	world[0][0].Height = 2

	x := len(world) - 1
	y := len(world[0]) - 1
	world[x][y].Height = 0

	//world.Print()

	world = grid.DigRiver2(world)
	world.Print()
}
