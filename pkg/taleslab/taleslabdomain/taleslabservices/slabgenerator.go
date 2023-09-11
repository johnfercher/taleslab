package taleslabservices

import (
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
)

type SlabGenerator interface {
	Generate(slabDto *taleslabentities.SlabGeneration) ([][]*taleslabentities.Slab, error)
}
