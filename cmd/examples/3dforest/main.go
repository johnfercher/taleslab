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
	forestService := forestservices.NewForestService(encoder)

	forest := &entities.Forest{
		X:                 70,
		Y:                 70,
		TerrainComplexity: 5,
		OrnamentDensity:   83,
		TreeDensity:       11,
		Mountains: &entities.Mountain{
			MinX:           15,
			RandX:          30,
			MinY:           15,
			RandY:          30,
			MinComplexity:  3,
			RandComplexity: 6,
			MinHeight:      10,
			RandHeight:     10,
		},
		HasRiver: true,
	}

	slab, err := forestService.GenerateForest(ctx, forest)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(slab.Code)
}
