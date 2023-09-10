package grid

import (
	"math"

	"github.com/johnfercher/taleslab/pkg/rand"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts/elementtype"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
)

func GenerateTerrain(maxX, maxY int, xFrequency, yFrequency, gain float64, minHeight int,
	forceBaseLand bool,
) taleslabentities.ElementMatrix {
	groundHeight := [][]taleslabentities.Element{}

	for i := 0; i < maxX; i++ {
		array := []taleslabentities.Element{}
		for j := 0; j < maxY; j++ {
			array = append(array, taleslabentities.Element{Height: 0, ElementType: elementtype.Ground})
		}
		groundHeight = append(groundHeight, array)
	}

	randomShiftX := float64(rand.Intn(20.0)/10.0) + 1
	randomShiftY := float64(rand.Intn(20.0)/10.0) + 1

	for i := 0; i < maxX; i++ {
		for j := 0; j < maxY; j++ {
			// i / maxX => normalized value between 0 and 1, represents each tile X axis
			// j / maxY => normalized value between 0 and 1, represents each tile Y axis
			// xFrequency | yFrequency => frequency of the senoidal wave that generates the mountain
			// + (rand.Pi / 2) => shifts the senoidal wave starting point

			xNormalizedValue := float64(i)/(float64(maxX)/(xFrequency)) + (math.Pi / randomShiftX)
			yNormalizedValue := float64(j)/(float64(maxY)/(yFrequency)) + (math.Pi / randomShiftY)

			// xHeight / yHeight => multiplied by rand.Pi because rand.Sin only accepts RADs
			xHeight := (gain * math.Sin(xNormalizedValue*math.Pi)) + gain
			yHeight := (gain * math.Sin(yNormalizedValue*math.Pi)) + gain

			heightAvg := int((xHeight + yHeight) / 2.0)

			groundHeight[i][j].Height = heightAvg + minHeight

			// Remove Ground
			if !forceBaseLand && heightAvg == 0 {
				groundHeight[i][j].ElementType = elementtype.Water
			}
		}
	}

	return groundHeight
}

func GenerateMountain(x, y int, gain float64, minHeight int) taleslabentities.ElementMatrix {
	xFrequency := 2.0
	yFrequency := 2.0

	mountainElements := [][]taleslabentities.Element{}

	for i := 0; i < x; i++ {
		array := []taleslabentities.Element{}
		for j := 0; j < y; j++ {
			array = append(array, taleslabentities.Element{Height: 0, ElementType: elementtype.Mountain})
		}
		mountainElements = append(mountainElements, array)
	}

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			// i / x => normalized value between 0 and 1, represents each tile X axis
			// j / y => normalized value between 0 and 1, represents each tile Y axis
			// xFrequency | yFrequency => frequency of the senoidal wave that generates the mountain
			// + (rand.Pi / 2) => shifts the senoidal wave starting point

			xNormalizedValue := float64(i)/(float64(x)/xFrequency) + (math.Pi / 2.0)
			yNormalizedValue := float64(j)/(float64(y)/yFrequency) + (math.Pi / 2.0)

			// xHeight / yHeight => multiplied by rand.Pi because rand.Sin only accepts RADs
			xHeight := (gain * math.Sin(xNormalizedValue*math.Pi)) + gain
			yHeight := (gain * math.Sin(yNormalizedValue*math.Pi)) + gain

			heightAvg := int((xHeight + yHeight) / 2.0)

			if heightAvg > int(gain) {
				mountainElements[i][j].Height = heightAvg - int(gain) + minHeight
			} else {
				mountainElements[i][j].Height = 0
				mountainElements[i][j].ElementType = elementtype.None
			}
		}
	}

	return mountainElements
}
