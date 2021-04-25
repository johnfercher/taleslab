package gridhelper

import (
	"fmt"
	"testing"
)

func TestBuildTerrain(t *testing.T) {
	x := 10
	y := 10

	world := [][]uint16{}

	for i := 0; i < x; i++ {
		array := []uint16{}
		for j := 0; j < y; j++ {
			array = append(array, uint16(0))
		}
		world = append(world, array)
	}

	mountain := MountainGenerator(3, 3, 2, 2, 20)

	world = BuildTerrain(world, mountain)

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			fmt.Printf("%d\t", world[i][j])
		}
		println()
	}

}
