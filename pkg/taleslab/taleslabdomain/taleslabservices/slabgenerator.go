package taleslabservices

import (
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdto"
)

type SlabGenerator interface {
	Generate(slabDto *taleslabdto.SlabDto) ([][]*taleslabentities.Slab, error)
}
