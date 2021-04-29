package vector

import "math"

func Distance(x1, y1, x2, y2 int) uint16 {
	a := math.Pow(float64(x2-x1), 2.0)
	b := math.Pow(float64(y2-y1), 2.0)
	c := uint16(math.Sqrt(a + b))
	return c
}
