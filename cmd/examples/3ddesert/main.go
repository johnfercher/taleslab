package main

import (
	"context"
	"fmt"
	"github.com/johnfercher/taleslab/pkg/api/domain/entities"
	"github.com/johnfercher/taleslab/pkg/api/forest/forestservices"
	"github.com/johnfercher/taleslab/pkg/slabcompressor"
	"github.com/johnfercher/taleslab/pkg/slabdecoder"
	"log"
)

func main() {
	ctx := context.TODO()

	compressor := slabcompressor.New()
	encoder := slabdecoder.NewEncoder(compressor)
	mapGenerator := forestservices.NewMapGenerator(encoder)

	forest := &entities.Map{
		Biome: entities.DesertBiome,
		Ground: &entities.Ground{
			Width:             70,
			Length:            70,
			TerrainComplexity: 5,
			ForceBaseLand:     true,
		},
		Props: &entities.Props{
			PropsDensity: 223,
			TreeDensity:  113,
		},
		Mountains: &entities.Mountains{
			MinX:           15,
			RandX:          30,
			MinY:           15,
			RandY:          30,
			MinComplexity:  3,
			RandComplexity: 6,
			MinHeight:      10,
			RandHeight:     10,
		},
		River: &entities.River{
			HasRiver: false,
		},
	}

	slab, err := mapGenerator.Generate(ctx, forest)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(slab.Code)
}
