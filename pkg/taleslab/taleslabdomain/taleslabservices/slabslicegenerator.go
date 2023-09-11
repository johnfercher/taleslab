package taleslabservices

import (
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
)

type SlabSliceGenerator interface {
	Generate(sliceDto *taleslabentities.SliceGeneration) (*taleslabentities.Slab, error)
}
