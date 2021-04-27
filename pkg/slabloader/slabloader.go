package slabloader

import (
	"github.com/johnfercher/taleslab/pkg/slab"
	"github.com/johnfercher/taleslab/pkg/slabdecoder"
	"io/ioutil"
	"strings"
)


type SlabLoader interface {
	GetSlabs() (map[string]*slab.Slab, error)
}

type slabLoader struct {
	decoder slabdecoder.Decoder
}

func NewSlabLoader(decoder slabdecoder.Decoder) *slabLoader {
	return &slabLoader{
		decoder: decoder,
	}
}

func (self *slabLoader) GetSlabs() (map[string]*slab.Slab, error) {
	bytes, err := ioutil.ReadFile("./config/slabs/slabs.csv")
	if err != nil {
		return nil, err
	}

	//fmt.Println(string(bytes))

	stringFile := string(bytes)
	lines := strings.Split(stringFile,"\n")

	//slabs := []*slab.Slab{}
	slabs := make(map[string]*slab.Slab)
	for _, line := range lines{
		elements := strings.SplitN(line,",",4)
		slab,err := self.decoder.Decode(elements[3])
		if err != nil {
			return nil, err
		}
		slabs[elements[1]] = slab
	}


	return slabs, nil
}

