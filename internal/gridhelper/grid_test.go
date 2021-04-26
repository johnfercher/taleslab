package gridhelper

import (
	"math/rand"
	"testing"
	"time"
)

func TestBuildTerrain(t *testing.T) {
	rand.Seed(time.Now().Unix())

	x := 20
	y := 20

	world := [][]uint16{}

	for i := 0; i < x; i++ {
		array := []uint16{}
		for j := 0; j < y; j++ {
			array = append(array, uint16(0))
		}
		world = append(world, array)
	}

	Print(world)

	mountain := MountainGenerator(6, 6, 30)

	Print(mountain)

	world = BuildTerrain(world, mountain)

	Print(world)
}
