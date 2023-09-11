package grid

import (
	"fmt"

	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
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
