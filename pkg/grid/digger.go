package grid

import (
	"fmt"
	"github.com/johnfercher/go-rrt/pkg/rrt"
	mathRRT "github.com/johnfercher/go-rrt/pkg/rrt/math"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
	"math"
)

func DigRiver(grid [][]taleslabentities.Element, river *River) [][]taleslabentities.Element {
	points := findRiverRandomPath(grid, river)

	points = GetFilledPoints(points, grid)

	for i := 0; i < len(points)-1; i++ {
		grid = digRiverBetweenPoints(points[i], points[i+1], grid)
	}

	return grid
}

func GetFilledPoints(points []*mathRRT.Point[taleslabentities.Element], grid [][]taleslabentities.Element) []*mathRRT.Point[taleslabentities.Element] {
	var newPoints []*mathRRT.Point[taleslabentities.Element]

	for i := len(points) - 1; i > 0; i-- {
		fmt.Printf("%s -> %s\n", points[i-1].GetString(), points[i].GetString())
		newPoints = append(newPoints, getPointsBetweenPoints(points[i-1], points[i], grid)...)
	}

	return newPoints
}

func getPointsBetweenPoints(a *mathRRT.Point[taleslabentities.Element], b *mathRRT.Point[taleslabentities.Element], grid [][]taleslabentities.Element) []*mathRRT.Point[taleslabentities.Element] {
	radian := mathRRT.Radian(a, b)

	points := make(map[string]*mathRRT.Point[taleslabentities.Element])

	i := 0.0
	for {
		deltaX := math.Sin(radian) * float64(i)
		x := int(a.X) + int(deltaX)

		deltaY := math.Cos(radian) * float64(i)
		y := int(a.Y) + int(deltaY)

		point := &mathRRT.Point[taleslabentities.Element]{
			X:    float64(x),
			Y:    float64(y),
			Data: grid[x][y],
		}

		points[fmt.Sprintf("%d-%d", x, y)] = point
		i += 0.5
		if mathRRT.Distance(point, b) < 1 {
			break
		}
	}

	var arr []*mathRRT.Point[taleslabentities.Element]
	for _, point := range points {
		arr = append(arr, point)
	}

	var reverse []*mathRRT.Point[taleslabentities.Element]
	for i := len(arr) - 1; i >= 0; i-- {
		reverse = append(reverse, arr[i])
	}

	return reverse
}

func digRiverBetweenPoints(a *mathRRT.Point[taleslabentities.Element], b *mathRRT.Point[taleslabentities.Element], grid [][]taleslabentities.Element) [][]taleslabentities.Element {
	minX, maxX := getMinMaxX(a, b)
	minY, maxY := getMinMaxY(a, b)

	minHeight := math.MaxInt
	for i := minX; i < maxX+1; i++ {
		for j := minY; j < maxY+1; j++ {
			height := grid[i][j].Height - 1
			if height < 0 {
				height = 0
			}
			if height < minHeight {
				minHeight = height
			}
			digSquare(i, j, 3, minHeight, grid)
		}
	}

	return grid
}

func getMinMaxX(a *mathRRT.Point[taleslabentities.Element], b *mathRRT.Point[taleslabentities.Element]) (int, int) {
	if a.X < b.X {
		return int(a.X), int(b.X)
	}

	return int(b.X), int(a.X)
}

func getMinMaxY(a *mathRRT.Point[taleslabentities.Element], b *mathRRT.Point[taleslabentities.Element]) (int, int) {
	if a.Y < b.Y {
		return int(a.Y), int(b.Y)
	}

	return int(b.Y), int(a.Y)
}

func findRiverRandomPath(grid [][]taleslabentities.Element, river *River) []*mathRRT.Point[taleslabentities.Element] {
	var min *mathRRT.Coordinate
	var max *mathRRT.Coordinate
	start := river.Start
	end := river.End

	if start != nil && end != nil {
		max = &mathRRT.Coordinate{X: float64(start.X), Y: float64(start.Y)}
		min = &mathRRT.Coordinate{X: float64(end.X), Y: float64(end.Y)}
	} else {
		min, max = getMinMaxHeights(grid)
	}

	riverRRT := rrt.New[taleslabentities.Element](5, 10000, 15)
	currentMax := grid[int(max.X)][int(max.Y)].Height
	riverRRT.AddCollisionCondition(func(point taleslabentities.Element) bool {
		if point.Height <= currentMax+river.HeightCutThreshold {
			currentMax = point.Height
			return false
		}

		return true
	})

	riverRRT.AddStopCondition(func(testPoint *mathRRT.Point[taleslabentities.Element], finish *mathRRT.Point[taleslabentities.Element]) bool {
		return testPoint.DistanceTo(finish) <= 5
	})

	return riverRRT.FindPath(max, min, grid)
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

func getMinMaxHeights(grid [][]taleslabentities.Element) (*mathRRT.Coordinate, *mathRRT.Coordinate) {
	minHeight := math.MaxInt
	min := &mathRRT.Coordinate{}
	maxHeight := 0
	max := &mathRRT.Coordinate{}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			elevation := grid[i][j].Height
			if elevation < minHeight {
				minHeight = elevation
				min = &mathRRT.Coordinate{X: float64(i), Y: float64(j)}
			}
			if elevation > maxHeight {
				maxHeight = elevation
				max = &mathRRT.Coordinate{X: float64(i), Y: float64(j)}
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
