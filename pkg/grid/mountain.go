package grid

import (
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
	"math"
	"math/rand"
	"time"
)

func TerrainGenerator(x, y int, xFrequency, yFrequency, gain float64, minHeight int, forceBaseLand bool) taleslabentities.ElementMatrix {
	groundHeight := [][]taleslabentities.Element{}

	for i := 0; i < x; i++ {
		array := []taleslabentities.Element{}
		for j := 0; j < y; j++ {
			array = append(array, taleslabentities.Element{Height: 0, ElementType: taleslabconsts.GroundType})
		}
		groundHeight = append(groundHeight, array)
	}

	rand.Seed(time.Now().UnixNano())
	randomShiftX := float64(rand.Intn(20.0)/10.0) + 1

	rand.Seed(time.Now().UnixNano())
	randomShiftY := float64(rand.Intn(20.0)/10.0) + 1

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			// i / x => normalized value between 0 and 1, represents each tile X axis
			// j / y => normalized value between 0 and 1, represents each tile Y axis
			// xFrequency | yFrequency => frequency of the senoidal wave that generates the mountain
			// + (math.Pi / 2) => shifts the senoidal wave starting point

			xNormalizedValue := float64(i)/(float64(x)/(xFrequency)) + (math.Pi / randomShiftX)
			yNormalizedValue := float64(j)/(float64(y)/(yFrequency)) + (math.Pi / randomShiftY)

			// xHeight / yHeight => multiplied by math.Pi because math.Sin only accepts RADs
			xHeight := (gain * math.Sin(xNormalizedValue*math.Pi)) + gain
			yHeight := (gain * math.Sin(yNormalizedValue*math.Pi)) + gain

			heightAvg := int((xHeight + yHeight) / 2.0)

			groundHeight[i][j].Height = heightAvg + minHeight

			// Remove Ground
			if !forceBaseLand && heightAvg == 0 {
				groundHeight[i][j].ElementType = taleslabconsts.WaterType
			}
		}
	}

	return groundHeight
}

func MountainGenerator(x, y int, gain float64, minHeight int) taleslabentities.ElementMatrix {
	xFrequency := 2.0
	yFrequency := 2.0

	mountainElements := [][]taleslabentities.Element{}

	for i := 0; i < x; i++ {
		array := []taleslabentities.Element{}
		for j := 0; j < y; j++ {
			array = append(array, taleslabentities.Element{Height: 0, ElementType: taleslabconsts.MountainType})
		}
		mountainElements = append(mountainElements, array)
	}

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			// i / x => normalized value between 0 and 1, represents each tile X axis
			// j / y => normalized value between 0 and 1, represents each tile Y axis
			// xFrequency | yFrequency => frequency of the senoidal wave that generates the mountain
			// + (math.Pi / 2) => shifts the senoidal wave starting point

			xNormalizedValue := float64(i)/(float64(x)/xFrequency) + (math.Pi / 2.0)
			yNormalizedValue := float64(j)/(float64(y)/yFrequency) + (math.Pi / 2.0)

			// xHeight / yHeight => multiplied by math.Pi because math.Sin only accepts RADs
			xHeight := (gain * math.Sin(xNormalizedValue*math.Pi)) + gain
			yHeight := (gain * math.Sin(yNormalizedValue*math.Pi)) + gain

			heightAvg := int((xHeight + yHeight) / 2.0)

			if heightAvg > int(gain) {
				mountainElements[i][j].Height = heightAvg - int(gain) + minHeight
			} else {
				mountainElements[i][j].Height = 0
			}
		}
	}

	return mountainElements
}

func SliceTerrain(base [][]taleslabentities.Element, sliceSize int) []taleslabentities.ElementMatrix {
	var slices []taleslabentities.ElementMatrix

	for i := 0; i < len(base); i += sliceSize {
		for j := 0; j < len(base[i]); j += sliceSize {
			//fmt.Printf("%d, %d\n", i, j)
			slices = append(slices, GetSliceInOffset(base, sliceSize, i, j))
		}
	}

	return slices
}

func GetSliceInOffset(base [][]taleslabentities.Element, sliceSize, offsetX, offsetY int) taleslabentities.ElementMatrix {
	xSliceSize := sliceSize
	ySliceSize := sliceSize

	if offsetX+sliceSize > len(base) {
		xSliceSize = sliceSize + len(base) - (offsetX + sliceSize)
	}

	if offsetY+sliceSize > len(base[0]) {
		ySliceSize = sliceSize + len(base) - (offsetY + sliceSize)
	}

	slice := GenerateElementGrid(xSliceSize, ySliceSize, taleslabentities.Element{0, taleslabconsts.GroundType})

	for i := 0; i+offsetX < len(base) && i < sliceSize; i++ {
		for j := 0; j+offsetY < len(base[i]) && j < sliceSize; j++ {
			//fmt.Printf("[%d] = %d, [%d] = %d\n", i, i+offsetX, j, j+offsetY)
			slice[i][j] = base[i+offsetX][j+offsetY]
		}
	}

	return slice
}
