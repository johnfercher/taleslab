package grid

import "testing"

func TestMountainGenerator(t *testing.T) {
	x := 50
	y := 50

	mountain := MountainGenerator(x, y, 5.0)

	PrintTypes(mountain)
}
