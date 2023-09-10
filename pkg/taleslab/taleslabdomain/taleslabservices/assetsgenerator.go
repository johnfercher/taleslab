package taleslabservices

import (
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts/biometype"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
)

type AssetsGenerator interface {
	SetBiome(biome biometype.BiomeType) AssetsGenerator
	SetSecondaryBiome(biomeType biometype.BiomeType) AssetsGenerator
	Generate(world [][]taleslabentities.Element, currentX, currentY int) (taleslabentities.Assets, error)
}
