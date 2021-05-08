package grid

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateElementGrid(x, y int, defaultElement Element) [][]Element {
	unitGrid := [][]Element{}

	for i := 0; i < x; i++ {
		array := []Element{}
		for j := 0; j < y; j++ {
			array = append(array, defaultElement)
		}
		unitGrid = append(unitGrid, array)
	}

	return unitGrid
}

func RandomlyFillEmptyGridSlots(worldGrid [][]Element, propsGrid [][]Element,
	density int, elementType ElementType, mustAdd func(element Element) bool) [][]Element {
	width := len(worldGrid)
	length := len(worldGrid[0])

	for i := 0; i < width; i++ {
		for j := 0; j < length; j++ {
			// Custom validation
			if !mustAdd(worldGrid[i][j]) {
				continue
			}

			// Avoid to add in limits
			if i == 0 || i == width-1 || j == 0 || j == length-1 {
				continue
			}

			// Avoid to add to close
			if i > 1 && (propsGrid[i-1][j].ElementType != NoneType || propsGrid[i-2][j].ElementType != NoneType) {
				continue
			}

			// Avoid to add to close
			if j > 1 && (propsGrid[i][j-1].ElementType != NoneType || propsGrid[i][j-2].ElementType != NoneType) {
				continue
			}

			if rand.Int()%density == 0 {
				propsGrid[i][j] = Element{ElementType: elementType}
			}
		}
	}

	return propsGrid
}

func BuildTerrain(world [][]Element, asset [][]Element) [][]Element {
	xMax := len(world)
	yMax := len(world[0])

	assetXMax := len(asset)
	assetYMax := len(asset[0])

	newWorld := Copy(world)

	rand.Seed(time.Now().UnixNano())

	randomXPosition := rand.Intn(xMax - assetXMax)
	randomYPosition := rand.Intn(yMax - assetYMax)

	for i := 0; i < assetXMax; i++ {
		for j := 0; j < assetYMax; j++ {
			assetValue := asset[i][j]
			worldValue := world[i+randomXPosition][j+randomYPosition]

			if assetValue.Height > worldValue.Height {
				newWorld[i+randomXPosition][j+randomYPosition] = assetValue
			} else {
				newWorld[i+randomXPosition][j+randomYPosition] = worldValue
			}
		}
	}

	return newWorld
}

func Copy(gridOriginal [][]Element) [][]Element {
	x := len(gridOriginal)
	y := len(gridOriginal[0])

	gridNew := [][]Element{}

	for i := 0; i < x; i++ {
		array := []Element{}
		for j := 0; j < y; j++ {
			array = append(array, gridOriginal[i][j])
		}
		gridNew = append(gridNew, array)
	}

	return gridNew
}

func Print(grid [][]Element) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			fmt.Printf("(%s, %d)\t", grid[i][j].ElementType, grid[i][j].Height)
		}
		fmt.Println()
	}
	fmt.Println()
}

func PrintTypes(grid [][]Element) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			fmt.Printf("%s\t", grid[i][j].ElementType)
		}
		fmt.Println()
	}
	fmt.Println()
}

func PrintHeights(grid [][]Element) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			fmt.Printf("%d\t", grid[i][j].Height)
		}
		fmt.Println()
	}
	fmt.Println()
}
