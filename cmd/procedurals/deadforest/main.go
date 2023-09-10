package main

import (
	"context"
	"log"

	"github.com/johnfercher/talescoder/pkg/encoder"
	"github.com/johnfercher/taleslab/pkg/file"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts/biometype"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdto"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabrepositories"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabservices"
)

func main() {
	ctx := context.TODO()

	encoder := encoder.NewEncoder()
	propRepository, _ := taleslabrepositories.NewPropRepository()
	biomeRepository, _ := taleslabrepositories.NewBiomeRepository()
	sliceGenerator := taleslabservices.NewSlabSliceGenerator(biomeRepository, propRepository)
	slabGenerator := taleslabservices.NewSlabGenerator(sliceGenerator)

	mapService := taleslabservices.NewMapService(slabGenerator, encoder)

	inputMap := &taleslabdto.MapDtoRequest{
		Biome: biometype.DeadForest,
		Ground: &taleslabdto.GroundDtoRequest{
			Width:             50,
			Length:            50,
			TerrainComplexity: 5,
			ForceBaseLand:     true,
		},
		Props: &taleslabdto.PropsDtoRequest{
			StoneDensity: 300,
			TreeDensity:  15,
			MiscDensity:  130,
		},
	}

	slab, err := mapService.Generate(ctx, inputMap)
	if err != nil {
		log.Fatal(err)
	}

	err = file.SaveCodes(slab.Codes, "cmd/procedurals/deadforest/data.txt")
	if err != nil {
		log.Fatal(err)
	}
}
