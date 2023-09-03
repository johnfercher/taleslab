package taleslabentities

import (
	"fmt"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts/biometype"
)

type Biome struct {
	Type      biometype.BiomeType `json:"biome_type"`
	Reliefs   Reliefs             `json:"reliefs"`
	StoneWall string              `json:"stone_wall"`
}

func (b *Biome) Print() {
	if b == nil {
		return
	}

	fmt.Printf("Biome: %s\n", b.Type)
	fmt.Printf("Stone: %s\n", b.StoneWall)
	b.Reliefs.Print()
}

type Reliefs struct {
	Water      *Relief `json:"water"`
	BaseGround *Relief `json:"base_ground"`
	Ground     *Relief `json:"ground"`
	Mountain   *Relief `json:"mountain"`
}

func (rs *Reliefs) Print() {
	if rs == nil {
		return
	}

	rs.Water.Print("water")
	rs.BaseGround.Print("base ground")
	rs.Ground.Print("ground")
	rs.Ground.Print("mountain")
}

type Relief struct {
	BuildingBlocks []string    `json:"building_blocks"`
	PropBlocks     *PropBlocks `json:"prop_blocks"`
}

func (r *Relief) Print(label string) {
	if r == nil {
		return
	}

	fmt.Printf("Relief: %s\n", label)
	fmt.Printf("BuildingBlocks: %v\n", r.BuildingBlocks)
	r.PropBlocks.Print()
	fmt.Println()
}

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
