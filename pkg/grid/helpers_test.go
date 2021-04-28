package grid

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestGenerateUintGrid(t *testing.T) {
	// Act
	grid := GenerateUintGrid(5, 5, Element{0, GroundType})

	// Assert
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			assert.Equal(t, Element{0, GroundType}, grid[i][j])
		}
	}
}

func TestBuildTerrain(t *testing.T) {
	// Arrange
	rand.Seed(time.Now().Unix())

	world2x2 := TerrainGenerator(5, 5, 2, 2, 15)

	//PrintHeights(world2x2)

	world1x1 := Duplicate(world2x2)

	PrintHeights(world1x1)

	mountain := MountainGenerator(5, 5, 45)

	PrintHeights(mountain)

	// Act
	world1x1 = BuildTerrain(world1x1, mountain)

	PrintHeights(world1x1)

	// Assert
	/*matched := false
	for i := 0 ; i < 5 && !matched ; i++ {
		for j := 0 ; j < 5 && !matched ; j++ {
			for i2 := 0 ; i2 < 3 && !matched ; i2++ {
				for j2 := 0 ; j2 < 3 && !matched ; j2++ {
					if world[i][j].Height != mountain[i2][j2].Height {
						break
					}

					if i2 == 2 && j2 == 2 {
						matched = true
						break
					}
				}
			}
		}
	}

	assert.True(t, matched)*/
}

func TestDuplicate(t *testing.T) {
	mountain := MountainGenerator(6, 6, 30)
	Print(mountain)

	duplicated := Duplicate(mountain)
	Print(duplicated)
}
