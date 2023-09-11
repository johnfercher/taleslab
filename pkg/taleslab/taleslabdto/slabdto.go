package taleslabdto

import (
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts/biometype"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
)

type SlabDto struct {
	World     [][]taleslabentities.Element
	SliceSize int
	Biomes    []biometype.BiomeType
}
