package taleslabentities

import "fmt"

type PropBlocks struct {
	Vegetation []string `json:"vegetation"`
	Stones     []string `json:"stones"`
	Misc       []string `json:"misc"`
}

func (p *PropBlocks) Print() {
	if p == nil {
		return
	}

	fmt.Printf("Vegetation: %v\n", p.Vegetation)
	fmt.Printf("Stones: %v\n", p.Stones)
	fmt.Printf("Misc: %v\n", p.Misc)
}
