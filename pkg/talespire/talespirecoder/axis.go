package talespirecoder

func DecodeX(y uint16) uint16 {
	return y / 100
}

func EncodeX(y uint16) uint16 {
	return y * 100
}

func DecodeY(y uint16) uint16 {
	result1 := y / 1600
	remain1 := y % 1600

	result2 := remain1 / 64

	return result1 + (41 * result2)
}

func EncodeY(y uint16) uint16 {
	return y * 1600
}

func DecodeZ(y uint16) uint16 {
	return y / 200
}

func EncodeZ(y uint16) uint16 {
	return y * 200
}
