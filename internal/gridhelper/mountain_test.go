package gridhelper

import "testing"

func TestMountainGenerator(t *testing.T) {
	x := 50
	y := 50

	_ = MountainGenerator(x, y, 5.0)

}
