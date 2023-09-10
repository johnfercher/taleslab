package taleslabentities

import (
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts/biometype"
)

type SliceGeneration struct {
	World          ElementMatrix
	FullDimension  *Dimensions
	SliceDimension *Dimensions
	OffsetX        int
	OffsetY        int
	Biomes         []biometype.BiomeType
}
