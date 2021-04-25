package gridhelper

import "math"

func MountainGenerator(x, y int, xFrequency, yFrequency, gain float64) [][]uint16 {
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
				groundHeight[i][j] = heightAvg
			} else {
				groundHeight[i][j] = uint16(gain)
			}
		}
	}

	return groundHeight

}
