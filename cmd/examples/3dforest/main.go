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
	mapService := taleslabservices.NewMapService(encoder)

	inputMap := &entities.Map{
		Biome: entities.ForestBiome,
		Ground: &entities.Ground{
			Width:             70,
			Length:            70,
			TerrainComplexity: 5,
			ForceBaseLand:     false,
		},
		Props: &entities.Props{
			PropsDensity: 83,
			TreeDensity:  11,
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
			HasRiver: true,
		},
	}

	slab, err := mapService.Generate(ctx, inputMap)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(slab.Code)
}
