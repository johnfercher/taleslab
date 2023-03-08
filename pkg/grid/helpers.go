package grid

import (
	"fmt"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
	"math/rand"
	"time"
)

func GenerateElementGrid(x, y int, defaultElement taleslabentities.Element) taleslabentities.ElementMatrix {
	unitGrid := [][]taleslabentities.Element{}

	for i := 0; i < x; i++ {
		array := []taleslabentities.Element{}
		for j := 0; j < y; j++ {
			array = append(array, defaultElement)
		}
		unitGrid = append(unitGrid, array)
	}

	return unitGrid
}

func RandomlyFillEmptyGridSlots(worldGrid [][]taleslabentities.Element, propsGrid [][]taleslabentities.Element,
	density int, elementType taleslabconsts.ElementType, mustAdd func(element taleslabentities.Element) bool) [][]taleslabentities.Element {
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
			if i > 1 && (propsGrid[i-1][j].ElementType != taleslabconsts.NoneType || propsGrid[i-2][j].ElementType != taleslabconsts.NoneType) {
				continue
			}

			// Avoid to add to close
			if j > 1 && (propsGrid[i][j-1].ElementType != taleslabconsts.NoneType || propsGrid[i][j-2].ElementType != taleslabconsts.NoneType) {
				continue
			}

			if rand.Int()%density == 0 {
				propsGrid[i][j] = taleslabentities.Element{ElementType: elementType}
			}
		}
	}

	return propsGrid
}

func AppendTerrainRandomly(baseTerrain [][]taleslabentities.Element, terrainToAppend [][]taleslabentities.Element) [][]taleslabentities.Element {
	xMax := len(baseTerrain)
	yMax := len(baseTerrain[0])

	assetXMax := len(terrainToAppend)
	assetYMax := len(terrainToAppend[0])

	newWorld := Copy(baseTerrain)

	rand.Seed(time.Now().UnixNano())

	randomXPosition := rand.Intn(xMax - assetXMax)
	randomYPosition := rand.Intn(yMax - assetYMax)

	for i := 0; i < assetXMax; i++ {
		for j := 0; j < assetYMax; j++ {
			assetValue := terrainToAppend[i][j]
			worldValue := baseTerrain[i+randomXPosition][j+randomYPosition]

			if assetValue.Height > worldValue.Height {
				newWorld[i+randomXPosition][j+randomYPosition] = assetValue
			} else {
				newWorld[i+randomXPosition][j+randomYPosition] = worldValue
			}
		}
	}

	return newWorld
}

func Copy(gridOriginal [][]taleslabentities.Element) [][]taleslabentities.Element {
	x := len(gridOriginal)
	y := len(gridOriginal[0])

	gridNew := [][]taleslabentities.Element{}

	for i := 0; i < x; i++ {
		array := []taleslabentities.Element{}
		for j := 0; j < y; j++ {
			array = append(array, gridOriginal[i][j])
		}
		gridNew = append(gridNew, array)
	}

	return gridNew
}

func Print(grid [][]taleslabentities.Element) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			fmt.Printf("(%s, %d)\t", grid[i][j].ElementType, grid[i][j].Height)
		}
		fmt.Println()
	}
	fmt.Println()
}

func PrintTypes(grid [][]taleslabentities.Element) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			fmt.Printf("%s\t", grid[i][j].ElementType)
		}
		fmt.Println()
	}
	fmt.Println()
}

func PrintHeights(grid [][]taleslabentities.Element) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			fmt.Printf("%d\t", grid[i][j].Height)
		}
		fmt.Println()
	}
	fmt.Println()
}
