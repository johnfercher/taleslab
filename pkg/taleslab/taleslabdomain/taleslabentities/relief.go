package taleslabentities

import (
	"fmt"

	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts/elementtype"
)

type Relief struct {
	Key            string      `json:"key"`
	BuildingBlocks []string    `json:"building_blocks"`
	PropBlocks     *PropBlocks `json:"prop_blocks"`
}

func (r *Relief) Print(label elementtype.ElementType) {
	if r == nil {
		return
	}

	fmt.Printf("Relief: %s\n", label)
	fmt.Printf("BuildingBlocks: %v\n", r.BuildingBlocks)
	r.PropBlocks.Print()
	fmt.Println()
}
