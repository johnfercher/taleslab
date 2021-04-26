package gridhelper

import (
	"math"
)

func DigRiver(grid [][]uint16) [][]uint16 {
	xFrequency := 2.0

	x := len(grid)

	gain := 5.0
	offset := x / 2.0

	for i := 0; i < x; i++ {
		xNormalizedValue := float64(float64(i)/(float64(x)/(xFrequency)) + (math.Pi))

		xHeight := uint16(gain*math.Sin(xNormalizedValue*math.Pi)) + uint16(offset)

		grid[i][xHeight] = 0
		grid[i][xHeight+1] = 0
		grid[i][xHeight-1] = 0
	}

	return grid
}
