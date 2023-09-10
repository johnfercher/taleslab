package taleslabdto

import (
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts/biometype"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
)

type SliceDto struct {
	World          taleslabentities.ElementMatrix
	FullDimension  *taleslabentities.Dimensions
	SliceDimension *taleslabentities.Dimensions
	OffsetX        int
	OffsetY        int
	Biomes         []biometype.BiomeType
}
