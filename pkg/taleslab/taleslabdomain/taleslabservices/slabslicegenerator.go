package taleslabservices

import (
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
)

type SlabSliceGenerator interface {
	Generate(sliceGeneration *taleslabentities.SliceGeneration) (*taleslabentities.Slab, error)
}
