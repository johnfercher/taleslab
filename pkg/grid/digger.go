package grid

import (
	"fmt"
	"github.com/johnfercher/go-rrt/pkg/rrt"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
	"math"
)

func DigRiver(grid [][]taleslabentities.Element) [][]taleslabentities.Element {
	xFrequency := 2.0

	x := len(grid)

	gain := 5.0
	offset := x / 2.0
	averageMin := math.MaxInt

	for i := 0; i < x; i++ {
		yNormalizedValue := float64(float64(i)/(float64(x)/(xFrequency)) + (math.Pi))

		randomY := uint16(gain*math.Sin(yNormalizedValue*math.Pi)) + uint16(offset)

		currentAverageMin := (grid[i][randomY].Height + grid[i][randomY+1].Height + grid[i][randomY+2].Height) / 3.0
		if currentAverageMin < averageMin {
			averageMin = currentAverageMin
		}

		grid[i][randomY] = taleslabentities.Element{
			Height:      getRiverHeight(averageMin),
			ElementType: taleslabconsts.Water,
		}
		grid[i][randomY+1] = taleslabentities.Element{
			Height:      getRiverHeight(averageMin),
			ElementType: taleslabconsts.Water,
		}
		grid[i][randomY+2] = taleslabentities.Element{
			Height:      getRiverHeight(averageMin),
			ElementType: taleslabconsts.Water,
		}
	}

	return grid
}

func getRiverHeight(height int) int {
	if height <= 0 {
		return height
	}

	return height - 1
}

func DigRiver3(grid [][]taleslabentities.Element) [][]taleslabentities.Element {
	min, max := getMinMax(grid)
	riverRRT := rrt.New[taleslabentities.Element](0.1)

	riverRRT.AddCollisionCondition(func(point taleslabentities.Element) bool {
		return point.Height > 10
	})
	riverRRT.AddStopCondition(func(testPoint *rrt.Point[taleslabentities.Element], finish *rrt.Point[taleslabentities.Element]) bool {
		return testPoint.DistanceTo(finish) <= 5
	})

	points := riverRRT.FindPath(max, min, grid)
	for _, point := range points {
		point.Println()
		x := int(point.X)
		y := int(point.Y)
		grid = digSquare(x, y, 3, grid[x][y].Height-1, grid)
	}

	return grid
}

func GenRiver(xMin, yMin, xMax, yMax, currentMinHeight int, grid [][]taleslabentities.Element) [][]taleslabentities.Element {
	lengthX := len(grid)
	lengthY := len(grid[0])

	x := xMax
	y := yMax

	grid = digSquare(x, y, 3, currentMinHeight, grid)

	minDistance := math.MaxFloat64
	minXDistance := xMax
	minYDistance := yMax

	currentHeight := grid[xMax][yMax].Height
	minHeight := math.MaxInt
	minXHeight := xMax
	minYHeight := yMax

	// Up
	if y > 0 {
		distance := getDistance(xMin, yMin, xMax, y-1)
		if distance < minDistance {
			minDistance = distance
			minXDistance = xMax
			minYDistance = y - 1
		}
		height := grid[xMax][y-1].Height
		if height < minHeight && grid[xMax][y-1].ElementType != taleslabconsts.Water {
			minHeight = height
			minXHeight = xMax
			minYHeight = y - 1
		}
	}

	// Up-Left
	if x > 0 && y > 0 {
		distance := getDistance(xMin, yMin, x-1, y-1)
		if distance < minDistance {
			minDistance = distance
			minXDistance = x - 1
			minYDistance = y - 1
		}
		height := grid[x-1][y-1].Height
		if height < minHeight && grid[x-1][y-1].ElementType != taleslabconsts.Water {
			minHeight = height
			minXHeight = x - 1
			minYHeight = y - 1
		}
	}

	// Left
	if x > 0 {
		distance := getDistance(xMin, yMin, x-1, yMax)
		if distance < minDistance {
			minDistance = distance
			minXDistance = x - 1
			minYDistance = yMax
		}
		height := grid[x-1][yMax].Height
		if height < minHeight && grid[x-1][yMax].ElementType != taleslabconsts.Water {
			minHeight = height
			minXHeight = x - 1
			minYHeight = yMax
		}
	}

	// Down-Left
	if x > 0 && y < lengthY-1 {
		distance := getDistance(xMin, yMin, x-1, y+1)
		if distance < minDistance {
			minDistance = distance
			minXDistance = x - 1
			minYDistance = y + 1
		}
		height := grid[x-1][y+1].Height
		if height < minHeight && grid[x-1][y+1].ElementType != taleslabconsts.Water {
			minHeight = height
			minXHeight = x - 1
			minYHeight = y + 1
		}
	}

	// Down
	if y < lengthY-1 {
		distance := getDistance(xMin, yMin, xMax, y+1)
		if distance < minDistance {
			minDistance = distance
			minXDistance = xMax
			minYDistance = y + 1
		}
		height := grid[xMax][y+1].Height
		if height < minHeight && grid[xMax][y+1].ElementType != taleslabconsts.Water {
			minHeight = height
			minXHeight = xMax
			minYHeight = y + 1
		}
	}

	// Down-Right
	if x < lengthX-1 && y < lengthY-1 {
		distance := getDistance(xMin, yMin, x+1, y+1)
		if distance < minDistance {
			minDistance = distance
			minXDistance = x + 1
			minYDistance = y + 1
		}
		height := grid[x+1][y+1].Height
		if height < minHeight && grid[x+1][y+1].ElementType != taleslabconsts.Water {
			minHeight = height
			minXHeight = x + 1
			minYHeight = y + 1
		}
	}

	// Right
	if x < lengthX-1 {
		distance := getDistance(xMin, yMin, x+1, yMax)
		if distance < minDistance {
			minDistance = distance
			minXDistance = x + 1
			minYDistance = yMax
		}
		height := grid[x+1][yMax].Height
		if height < minHeight && grid[x+1][yMax].ElementType != taleslabconsts.Water {
			minHeight = height
			minXHeight = x + 1
			minYHeight = yMax
		}
	}

	// Up-Right
	if x < lengthX-1 && y > 0 {
		distance := getDistance(xMin, yMin, x+1, y-1)
		if distance < minDistance {
			minDistance = distance
			minXDistance = x + 1
			minYDistance = y - 1
		}
		height := grid[x+1][y-1].Height
		if height < minHeight && grid[x+1][y-1].ElementType != taleslabconsts.Water {
			minHeight = height
			minXHeight = x + 1
			minYHeight = y - 1
		}
	}

	if minHeight < currentMinHeight {
		currentMinHeight = minHeight
	}

	if minDistance <= 1 {
		grid = digSquare(minXDistance, minYDistance, 3, currentMinHeight, grid)
		return grid
	}

	newX := 0
	newY := 0

	if minHeight < currentHeight {
		newX = minXHeight
		newY = minYHeight
	} else {
		newX = minXDistance
		newY = minYDistance
	}

	fmt.Printf("%d, %d\n", newX, newY)

	return GenRiver(xMin, yMin, newX, newY, currentMinHeight, grid)
}

func digSquare(x, y, size, currentMinHeight int, grid [][]taleslabentities.Element) [][]taleslabentities.Element {
	squareSize := (size - 1) / 2

	for i := x - squareSize; i < x+squareSize; i++ {
		for j := y - squareSize; j < y+squareSize; j++ {
			if i > 0 && i < len(grid) && j > 0 && j < len(grid[0]) {
				grid[i][j].ElementType = taleslabconsts.Water
				grid[i][j].Height = currentMinHeight
			}
		}
	}

	return grid
}

func getMinMax(grid [][]taleslabentities.Element) (*rrt.Coordinate, *rrt.Coordinate) {
	minHeight := math.MaxInt
	min := &rrt.Coordinate{}
	maxHeight := 0
	max := &rrt.Coordinate{}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			elevation := grid[i][j].Height
			if elevation < minHeight {
				minHeight = elevation
				min = &rrt.Coordinate{X: float64(i), Y: float64(j)}
			}
			if elevation > maxHeight {
				maxHeight = elevation
				max = &rrt.Coordinate{X: float64(i), Y: float64(j)}
			}
		}
	}

	return min, max
}

func getDistance(t1X, t1Y, t2X, t2Y int) float64 {
	return math.Sqrt(float64(t1X-t2X)*float64(t1X-t2X) + float64(t1Y-t2Y)*float64(t1Y-t2Y))
}

func DigTerrainInOffset(baseTerrain [][]taleslabentities.Element, offset int) [][]taleslabentities.Element {
	yFrequency := 2.0

	y := len(baseTerrain[0])

	gain := 3.0

	for j := 0; j < y; j++ {
		yNormalizedValue := float64(float64(j)/(float64(y)/(yFrequency)) + (math.Pi))

		randomX := uint16(gain*math.Sin(yNormalizedValue*math.Pi)) + uint16(offset)

		for x := 0; x < int(randomX); x++ {
			baseTerrain[x][j] = taleslabentities.Element{
				Height:      1,
				ElementType: taleslabconsts.BaseGround,
			}
		}
	}

	return baseTerrain
}
