package grid

import (
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
	"math"
)

func DigRiver(grid [][]taleslabentities.Element) [][]taleslabentities.Element {
	xFrequency := 2.0

	x := len(grid)

	gain := 5.0
	offset := x / 2.0

	for i := 0; i < x; i++ {
		yNormalizedValue := float64(float64(i)/(float64(x)/(xFrequency)) + (math.Pi))

		randomY := uint16(gain*math.Sin(yNormalizedValue*math.Pi)) + uint16(offset)

		grid[i][randomY] = taleslabentities.Element{
			Height:      1,
			ElementType: taleslabconsts.Water,
		}
		grid[i][randomY+1] = taleslabentities.Element{
			Height:      1,
			ElementType: taleslabconsts.Water,
		}
		grid[i][randomY+2] = taleslabentities.Element{
			Height:      1,
			ElementType: taleslabconsts.Water,
		}
	}

	return grid
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
