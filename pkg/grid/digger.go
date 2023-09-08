package grid

import (
	"fmt"
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

func DigRiver2(grid [][]taleslabentities.Element) [][]taleslabentities.Element {
	min, max := getMinMax(grid)
	return GenRiver(min.X, min.Y, max.X, max.Y, max.Z, grid)
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
			minDistance = distanceq
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

func getMinMax(grid [][]taleslabentities.Element) (taleslabentities.Vector3d, taleslabentities.Vector3d) {
	min := math.MaxInt
	minX := 0
	minY := 0
	max := 0
	maxX := 0
	maxY := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			elevation := grid[i][j].Height
			if elevation < min {
				min = elevation
				minX = i
				minY = j
			}
			if elevation > max {
				max = elevation
				maxX = i
				maxY = j
			}
		}
	}

	return taleslabentities.Vector3d{X: minX, Y: minY, Z: min}, taleslabentities.Vector3d{X: maxX, Y: maxY, Z: max}
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
