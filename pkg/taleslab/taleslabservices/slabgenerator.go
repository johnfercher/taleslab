package taleslabservices

import (
	"errors"
	"github.com/johnfercher/taleslab/pkg/grid"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabservices"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdto"
)

type slabGenerator struct {
	slabSliceGenerator taleslabservices.SlabSliceGenerator
}

func NewSlabGenerator(slabSliceGenerator taleslabservices.SlabSliceGenerator) taleslabservices.SlabGenerator {
	return &slabGenerator{
		slabSliceGenerator: slabSliceGenerator,
	}
}

func (s *slabGenerator) Generate(slabDto *taleslabdto.SlabDto) ([][]*taleslabentities.Slab, error) {
	if len(slabDto.Biomes) == 0 {
		return nil, errors.New("must provide at least one biome")
	}

	slabs := [][]*taleslabentities.Slab{}
	worldMatrixSlices := grid.SliceTerrain(slabDto.World, slabDto.SliceSize)

	currentX := 0
	currentY := 0
	for _, worldMatrix := range worldMatrixSlices {
		line := []*taleslabentities.Slab{}
		for _, slice := range worldMatrix {
			sliceDto := &taleslabdto.SliceDto{
				World: slice,
				FullDimension: &taleslabentities.Dimensions{
					Width:  len(slabDto.World),
					Length: len(slabDto.World[0]),
				},
				SliceDimension: &taleslabentities.Dimensions{
					Width:  len(slice),
					Length: len(slice[0]),
				},
				OffsetX: currentX,
				OffsetY: currentY,
				Biomes:  slabDto.Biomes,
			}

			slab, err := s.slabSliceGenerator.Generate(sliceDto)
			if err != nil {
				return nil, err
			}
			line = append(line, slab)
			currentY += slabDto.SliceSize
		}
		slabs = append(slabs, line)
		currentY = 0
		currentX += slabDto.SliceSize
	}

	return slabs, nil
}
