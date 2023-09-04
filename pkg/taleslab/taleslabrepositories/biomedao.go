package taleslabrepositories

import (
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts/biometype"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
)

type BiomeDao struct {
	Type      biometype.BiomeType        `json:"biome_type"`
	Reliefs   []*taleslabentities.Relief `json:"reliefs"`
	StoneWall string                     `json:"stone_wall"`
}
