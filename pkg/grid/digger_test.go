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

/*func TestGetFilledPoints(t *testing.T) {
	// Arrange
	var points = []*math.Point[taleslabentities.Element]{}
	points = append(points, &math.Point[taleslabentities.Element]{
		X: 0,
		Y: 0,
	})
	points = append(points, &math.Point[taleslabentities.Element]{
		X: 0,
		Y: 5,
	})
	points = append(points, &math.Point[taleslabentities.Element]{
		X: 5,
		Y: 5,
	})
	points = append(points, &math.Point[taleslabentities.Element]{
		X: 10,
		Y: 10,
	})

	// Act
	newPoints := grid.GetFilledPoints(points)

	// Assert
	assert.Equal(t, true, newPoints)
}
*/
