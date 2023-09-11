package proceduralservices

import (
	"context"
	"errors"
	"github.com/johnfercher/talescoder/pkg/encoder"
	"github.com/johnfercher/taleslab/pkg/procedural/proceduraldomain/proceduralentities"
	proceduralservices2 "github.com/johnfercher/taleslab/pkg/procedural/proceduraldomain/proceduralservices"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts/biometype"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabservices"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabmappers"
)

type mapService struct {
	slabGenerator taleslabservices.SlabGenerator
	encoder       encoder.Encoder
}

func NewMapService(slabGenerator taleslabservices.SlabGenerator, encoder encoder.Encoder) proceduralservices2.MapService {
	return &mapService{
		slabGenerator: slabGenerator,
		encoder:       encoder,
	}
}

func (m *mapService) Generate(ctx context.Context, inputMap *proceduralentities.MapGeneration) (*proceduralentities.MapGenerated, error) {
	if inputMap.Biome == "" {
		return nil, errors.New("must provide at least one biome")
	}

	matrixGenerator := NewMatrixGenerator().
		SetMountains(inputMap.Mountains).
		SetGround(inputMap.Ground).
		SetRiver(inputMap.River).
		SetCanyon(inputMap.Canyon)

	world, err := matrixGenerator.Generate()
	if err != nil {
		return nil, err
	}

	slabDto := &taleslabentities.SlabGeneration{
		World:     world,
		SliceSize: 50,
		Biomes:    []biometype.BiomeType{inputMap.Biome},
	}

	if inputMap.SecondaryBiome != "" {
		slabDto.Biomes = append(slabDto.Biomes, inputMap.SecondaryBiome)
	}

	slabs, err := m.slabGenerator.Generate(slabDto)
	if err != nil {
		return nil, err
	}

	response := &proceduralentities.MapGenerated{
		SlabVersion: "v2",
	}

	for _, slab := range slabs {
		var codes []string
		for _, line := range slab {
			talespireSlab := taleslabmappers.TaleSpireSlabFromSlab(line)
			base64, err := m.encoder.Encode(talespireSlab)
			if err != nil {
				return nil, err
			}

			codes = append(codes, base64)
			response.Size += len(base64) / 1024
		}
		response.Codes = append(response.Codes, codes)
	}

	return response, nil
}
