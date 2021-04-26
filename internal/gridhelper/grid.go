package gridhelper

import (
	"fmt"
	"github.com/johnfercher/taleslab/pkg/api/domain/entities"
	"math/rand"
	"time"
)

func GenerateUintGrid(x, y int, defaultValue uint16) [][]uint16 {
	unitGrid := [][]uint16{}

	for i := 0; i < x; i++ {
		array := []uint16{}
		for j := 0; j < y; j++ {
			array = append(array, defaultValue)
		}
		unitGrid = append(unitGrid, array)
	}

	return unitGrid
}

func GenerateBoolGrid(x, y int, defaultValue bool) [][]bool {
	boolGrid := [][]bool{}

	for i := 0; i < x; i++ {
		array := []bool{}
		for j := 0; j < y; j++ {
			array = append(array, defaultValue)
		}
		boolGrid = append(boolGrid, array)
	}

	return boolGrid
}

func GenerateRandomGridPositions(forest *entities.Forest) [][]bool {
	defaultValue := false
	groundSpots := GenerateBoolGrid(forest.Ground.Width, forest.Ground.Length, defaultValue)

	for i := 0; i < forest.Ground.Width; i++ {
		for j := 0; j < forest.Ground.Length; j++ {
			if i == 0 || i == forest.Ground.Width-1 || j == 0 || j == forest.Ground.Length-1 {
				continue
			}

			if i > 1 && (groundSpots[i-1][j] || groundSpots[i-2][j]) {
				continue
			}

			if j > 1 && (groundSpots[i][j-1] || groundSpots[i][j-2]) {
				continue
			}

			groundSpots[i][j] = rand.Int()%forest.Props.PropsDensity == 0
		}
	}

	return groundSpots
}

func GenerateExclusiveRandomGrid(forest *entities.Forest, unavailableSpots [][]bool) [][]bool {
	defaultValue := false
	x := forest.Ground.Width
	y := forest.Ground.Length

	groundSpots := GenerateBoolGrid(x, y, defaultValue)

	for i := 0; i < x; i++ {
		array := []bool{}
		for j := 0; j < y; j++ {
			array = append(array, defaultValue)
		}
		groundSpots = append(groundSpots, array)
	}

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if i == 0 || i == x-1 || j == 0 || j == y-1 {
				continue
			}

			if unavailableSpots[i][j] {
				continue
			}

			if i > 1 && (groundSpots[i-1][j] || groundSpots[i-2][j]) {
				continue
			}

			if j > 1 && (groundSpots[i][j-1] || groundSpots[i][j-2]) {
				continue
			}

			groundSpots[i][j] = rand.Int()%forest.Props.TreeDensity == 0
		}
	}
	return groundSpots
}

func BuildTerrain(world [][]uint16, asset [][]uint16) [][]uint16 {
	xMax := len(world)
	yMax := len(world[0])

	assetXMax := len(asset)
	assetYMax := len(asset[0])

	rand.Seed(time.Now().UnixNano())

	randomXPosition := rand.Intn(xMax - assetXMax)
	randomYPosition := rand.Intn(yMax - assetYMax)

	for i := 0; i < assetXMax; i++ {
		for j := 0; j < assetYMax; j++ {
			assetValue := asset[i][j]
			worldValue := world[i+randomXPosition][j+randomYPosition]

			if assetValue > worldValue {
				world[i+randomXPosition][j+randomYPosition] = assetValue
			}
		}
	}

	return world
}

func Print(grid [][]uint16) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			fmt.Printf("%d\t", grid[i][j])
		}
		fmt.Println()
	}
}
