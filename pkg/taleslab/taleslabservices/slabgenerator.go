package taleslabservices

import (
	"errors"
	"github.com/johnfercher/taleslab/pkg/shared/grid"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabservices"
)

type slabGenerator struct {
	slabSliceGenerator taleslabservices.SlabSliceGenerator
}

func NewSlabGenerator(slabSliceGenerator taleslabservices.SlabSliceGenerator) taleslabservices.SlabGenerator {
	return &slabGenerator{
		slabSliceGenerator: slabSliceGenerator,
	}
}

func (s *slabGenerator) Generate(slabGeneration *taleslabentities.SlabGeneration) ([][]*taleslabentities.Slab, error) {
	if len(slabGeneration.Biomes) == 0 {
		return nil, errors.New("must provide at least one biome")
	}

	slabs := [][]*taleslabentities.Slab{}
	worldMatrixSlices := grid.SliceTerrain(slabGeneration.World, slabGeneration.SliceSize)

	currentX := 0
	currentY := 0
	for _, worldMatrix := range worldMatrixSlices {
		line := []*taleslabentities.Slab{}
		for _, slice := range worldMatrix {
			sliceGeneration := &taleslabentities.SliceGeneration{
				World: slice,
				FullDimension: &taleslabentities.Dimensions{
					Width:  len(slabGeneration.World),
					Length: len(slabGeneration.World[0]),
				},
				SliceDimension: &taleslabentities.Dimensions{
					Width:  len(slice),
					Length: len(slice[0]),
				},
				OffsetX: currentX,
				OffsetY: currentY,
				Biomes:  slabGeneration.Biomes,
			}

			slab, err := s.slabSliceGenerator.Generate(sliceGeneration)
			if err != nil {
				return nil, err
			}
			line = append(line, slab)
			currentY += slabGeneration.SliceSize
		}
		slabs = append(slabs, line)
		currentY = 0
		currentX += slabGeneration.SliceSize
	}

	return slabs, nil
}
