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
	GetSlabs() map[string]*entities.Slab
	GetSlabById(id string) *entities.Slab
}

type slabLoader struct {
	decoder talespirecoder.Decoder
	slabs   map[string]*entities.Slab
}

func NewSlabLoader(decoder talespirecoder.Decoder) (*slabLoader, error) {
	slabLoader := &slabLoader{
		decoder: decoder,
	}

	err := slabLoader.loadSlabs()
	if err != nil {
		return nil, err
	}

	return slabLoader, nil
}

func (self *slabLoader) GetSlabs() map[string]*entities.Slab {
	return self.slabs
}

func (self *slabLoader) GetSlabById(id string) *entities.Slab {
	return self.slabs[id]
}

func (self *slabLoader) loadSlabs() error {
	bytes, err := ioutil.ReadFile("./config/assets/slabs.csv")
	if err != nil {
		return err
	}

	stringFile := string(bytes)
	lines := strings.Split(stringFile, "\n")

	taleSpireSlabs := make(map[string]*talespirecontracts.Slab)
	for _, line := range lines {
		elements := strings.SplitN(line, ",", 4)
		slab, err := self.decoder.Decode(elements[3])
		if err != nil {
			return err
		}
		taleSpireSlabs[elements[1]] = slab
	}

	slabs := make(map[string]*entities.Slab)
	for key, taleSpireSlab := range taleSpireSlabs {
		slabs[key] = mappers.EntitySlabFromTaleSpire(taleSpireSlab)
	}

	return nil
}
