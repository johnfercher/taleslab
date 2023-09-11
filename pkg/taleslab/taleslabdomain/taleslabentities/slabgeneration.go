package taleslabentities

import (
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts/biometype"
)

type SlabGeneration struct {
	World     [][]Element
	SliceSize int
	Biomes    []biometype.BiomeType
}
