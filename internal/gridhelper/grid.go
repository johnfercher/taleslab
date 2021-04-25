package gridhelper

import (
	"math/rand"
)

func GenerateRandomGridPositions(x, y, randomBias int) [][]bool {
	defaultValue := false
	groundSpots := [][]bool{}

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

			if i > 1 && (groundSpots[i-1][j] || groundSpots[i-2][j]) {
				continue
			}

			if j > 1 && (groundSpots[i][j-1] || groundSpots[i][j-2]) {
				continue
			}

			groundSpots[i][j] = rand.Int()%randomBias == 0
		}
	}

	return groundSpots
}

func GenerateExclusiveRandomGrid(x, y, randomBias int, unavailableSpots [][]bool) [][]bool {
	defaultValue := false
	groundSpots := [][]bool{}

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

			groundSpots[i][j] = rand.Int()%randomBias == 0
		}
	}
	return groundSpots
}
