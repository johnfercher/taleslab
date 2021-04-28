package main

import (
	"context"
	"fmt"
	"github.com/johnfercher/taleslab/pkg/api/domain/entities"
	"github.com/johnfercher/taleslab/pkg/api/taleslab/taleslabservices"
	"github.com/johnfercher/taleslab/pkg/slabcompressor"
	"github.com/johnfercher/taleslab/pkg/slabdecoder"
	"log"
)

func main() {
	ctx := context.TODO()

	compressor := slabcompressor.New()
	encoder := slabdecoder.NewEncoder(compressor)
	mapGenerator := taleslabservices.NewMapService(encoder)

	inputMap := &entities.Map{
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
		River: &entities.River{
			HasRiver: false,
		},
	}

	slab, err := mapGenerator.Generate(ctx, inputMap)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(slab.Code)
}
