package gridhelper

import (
	"math"
	"math/rand"
	"time"
)

func TerrainGenerator(x, y int, xFrequency, yFrequency, gain float64) [][]uint16 {
	groundHeight := [][]uint16{}

	for i := 0; i < x; i++ {
		array := []uint16{}
		for j := 0; j < y; j++ {
			array = append(array, uint16(0))
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

			heightAvg := uint16((xHeight + yHeight) / 2.0)

			groundHeight[i][j] = heightAvg
		}
	}

	return groundHeight
}

func MountainGenerator(x, y int, gain float64) [][]uint16 {
	xFrequency := 2.0
	yFrequency := 2.0

	groundHeight := [][]uint16{}

	for i := 0; i < x; i++ {
		array := []uint16{}
		for j := 0; j < y; j++ {
			array = append(array, uint16(0))
		}
		groundHeight = append(groundHeight, array)
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

			heightAvg := uint16((xHeight + yHeight) / 2.0)

			if heightAvg > uint16(gain) {
				groundHeight[i][j] = heightAvg - uint16(gain)
			} else {
				groundHeight[i][j] = 0
			}
		}
	}

	return groundHeight
}
