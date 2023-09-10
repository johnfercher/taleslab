package grid

import (
	"fmt"

	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts/elementtype"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
)

func SliceTerrain(base [][]taleslabentities.Element, sliceSize int) [][]taleslabentities.ElementMatrix {
	var matrix [][]taleslabentities.ElementMatrix

	for i := 0; i < len(base); i += sliceSize {
		var slices []taleslabentities.ElementMatrix
		for j := 0; j < len(base[i]); j += sliceSize {
			fmt.Printf("%d, %d\n", i, j)
			slices = append(slices, getSliceInOffset(base, sliceSize, i, j))
		}
		matrix = append(matrix, slices)
	}

	return matrix
}

func getSliceInOffset(base [][]taleslabentities.Element, sliceSize, offsetX, offsetY int) taleslabentities.ElementMatrix {
	xSliceSize := sliceSize
	ySliceSize := sliceSize

	if offsetX+sliceSize > len(base) {
		xSliceSize = sliceSize + len(base) - (offsetX + sliceSize)
	}

	if offsetY+sliceSize > len(base[0]) {
		ySliceSize = sliceSize + len(base[0]) - (offsetY + sliceSize)
	}

	slice := GenerateElementGrid(xSliceSize, ySliceSize, taleslabentities.Element{Height: 0, ElementType: elementtype.Ground})

	for i := 0; i+offsetX < len(base) && i < sliceSize; i++ {
		for j := 0; j+offsetY < len(base[i]) && j < sliceSize; j++ {
			// fmt.Printf("[%d] = %d, [%d] = %d\n", i, i+offsetX, j, j+offsetY)
			slice[i][j] = base[i+offsetX][j+offsetY]
		}
	}

	return slice
}
