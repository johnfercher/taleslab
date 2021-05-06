package talespirecoder

func DecodeY(y int16) int16 {
	if y%1600 == 0 {
		result1 := y / 1600
		remain1 := y % 1600

		result2 := remain1 / 64

		return (result1 + (41 * result2)) * 160
	}

	return (y / 10) + 1
}

func EncodeY(y int16) int16 {
	if y%160 != 0 {
		return (y - 1) * 10
	}

	return int16(float64(y) / float64(160.0) * float64(1600.0))
}
