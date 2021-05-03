package slabloader

import (
	"github.com/johnfercher/taleslab/pkg/mappers"
	"github.com/johnfercher/taleslab/pkg/taleslab/domain/entities"
	"github.com/johnfercher/taleslab/pkg/talespire/talespirecoder"
	"github.com/johnfercher/taleslab/pkg/talespire/talespirecontracts"
	"io/ioutil"
	"strings"
)

type SlabLoader interface {
	GetSlabs() (map[string]*entities.Slab, error)
}

type slabLoader struct {
	decoder talespirecoder.Decoder
}

func NewSlabLoader(decoder talespirecoder.Decoder) *slabLoader {
	return &slabLoader{
		decoder: decoder,
	}
}

func (self *slabLoader) GetSlabs() (map[string]*entities.Slab, error) {
	bytes, err := ioutil.ReadFile("./config/assets/slabs.csv")
	if err != nil {
		return nil, err
	}

	stringFile := string(bytes)
	lines := strings.Split(stringFile, "\n")

	taleSpireSlabs := make(map[string]*talespirecontracts.Slab)
	for _, line := range lines {
		elements := strings.SplitN(line, ",", 4)
		slab, err := self.decoder.Decode(elements[3])
		if err != nil {
			return nil, err
		}
		taleSpireSlabs[elements[1]] = slab
	}

	slabs := make(map[string]*entities.Slab)
	for key, taleSpireSlab := range taleSpireSlabs {
		slabs[key] = mappers.EntitySlabFromTaleSpire(taleSpireSlab)
	}

	return slabs, nil
}
