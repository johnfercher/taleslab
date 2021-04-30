package grid

import (
	"math"
)

func DigRiver(grid [][]Element) [][]Element {
	xFrequency := 2.0

	x := len(grid)

	gain := 5.0
	offset := x / 2.0

	for i := 0; i < x; i++ {
		yNormalizedValue := float64(float64(i)/(float64(x)/(xFrequency)) + (math.Pi))

		randomY := uint16(gain*math.Sin(yNormalizedValue*math.Pi)) + uint16(offset)

		grid[i][randomY] = Element{
			Height:      0,
			ElementType: RiverType,
		}
		grid[i][randomY+1] = Element{
			Height:      0,
			ElementType: RiverType,
		}
		grid[i][randomY+2] = Element{
			Height:      0,
			ElementType: RiverType,
		}
	}

	return grid
}

func DigCanyon(grid [][]Element, offset uint) [][]Element {
	yFrequency := 2.0

	y := len(grid[0])

	gain := 3.0

	for j := 0; j < y; j++ {
		yNormalizedValue := float64(float64(j)/(float64(y)/(yFrequency)) + (math.Pi))

		randomX := uint16(gain*math.Sin(yNormalizedValue*math.Pi)) + uint16(offset)

		for x := 0; x < int(randomX); x++ {
			grid[x][j] = Element{
				Height:      1,
				ElementType: BaseGroundType,
			}
		}
	}

	return grid
}
