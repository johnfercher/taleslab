package taleslabentities

import "fmt"

type Vector3d struct {
	X int
	Y int
	Z int
}

func (v *Vector3d) Print(label string) {
	fmt.Printf("%s(x: %d, y: %d, z: %d)\n", label, v.X, v.Y, v.Z)
}
