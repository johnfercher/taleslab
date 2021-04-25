package gridhelper

import (
	"fmt"
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

	mountain := MountainGenerator(6, 6, 30)

	for i := 0; i < len(mountain); i++ {
		for j := 0; j < len(mountain[0]); j++ {
			fmt.Printf("%d\t", mountain[i][j])
		}
		println()
	}

	world = BuildTerrain(world, mountain)

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			fmt.Printf("%d\t", world[i][j])
		}
		println()
	}
}
