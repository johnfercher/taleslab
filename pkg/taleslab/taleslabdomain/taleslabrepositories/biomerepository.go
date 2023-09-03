package taleslabrepositories

import (
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts/biometype"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
)

type BiomeRepository interface {
	GetBiome(biometype.BiomeType) *taleslabentities.Biome
}
