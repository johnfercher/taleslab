package taleslabservices

import (
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdto"
)

type SlabSliceGenerator interface {
	Generate(sliceDto *taleslabdto.SliceDto) (*taleslabentities.Slab, error)
}
