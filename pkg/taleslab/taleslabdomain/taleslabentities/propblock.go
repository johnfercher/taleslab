package taleslabentities

import "fmt"

type PropBlocks struct {
	Vegetation PropDistribution `json:"vegetation"`
	Stones     PropDistribution `json:"stones"`
	Misc       PropDistribution `json:"misc"`
}

func (p *PropBlocks) Print() {
	if p == nil {
		return
	}

	fmt.Printf("Vegetation: %v\n", p.Vegetation)
	fmt.Printf("Stones: %v\n", p.Stones)
	fmt.Printf("Misc: %v\n", p.Misc)
}
