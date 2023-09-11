package proceduralentities

import (
	"github.com/johnfercher/taleslab/pkg/shared/grid"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts/biometype"
)

type MapGenerated struct {
	SlabVersion string
	Size        int
	Codes       [][]string
}

type MapGeneration struct {
	Biome          biometype.BiomeType
	SecondaryBiome biometype.BiomeType
	Ground         *Ground
	Mountains      *Mountains
	River          *grid.River
	Canyon         *Canyon
	Props          *Props
}
