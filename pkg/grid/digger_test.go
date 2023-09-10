package grid_test

import (
	"github.com/johnfercher/taleslab/pkg/grid"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts/elementtype"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
	"testing"
)

func TestDigRiver(t *testing.T) {
	world := grid.GenerateElementGrid(30, 30, taleslabentities.Element{
		1,
		elementtype.Ground,
	})

	world = grid.DigRiver(world)

	world.Print()
}
