package taleslabentities

import (
	"fmt"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts"
)

type Element struct {
	Height      int
	ElementType taleslabconsts.ElementType
}

func (e *Element) Print() {
	fmt.Printf("%s:%d", e.ElementType, e.Height)
}

type ElementMatrix [][]Element

func (m ElementMatrix) Print() {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			m[i][j].Print()
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
