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
		Biome:          biometype.TemperateForest,
		SecondaryBiome: biometype.Tundra,
		Ground: &taleslabdto.GroundDtoRequest{
			Width:             50,
			Length:            50,
			TerrainComplexity: 5,
			MinHeight:         5,
			ForceBaseLand:     true,
		},
		Props: &taleslabdto.PropsDtoRequest{
			StoneDensity: 150,
			TreeDensity:  11,
			MiscDensity:  11,
		},
		Mountains: &taleslabdto.MountainsDtoRequest{
			MinX:           30,
			RandX:          5,
			MinY:           30,
			RandY:          5,
			MinComplexity:  5,
			RandComplexity: 2,
			MinHeight:      10,
			RandHeight:     10,
		},
		Canyon: &taleslabdto.CanyonDtoRequest{
			HasCanyon:    false,
			CanyonOffset: 10,
		},
	}

	slab, err := mapService.Generate(ctx, inputMap)
	if err != nil {
		log.Fatal(err)
	}

	err = file.SaveCodes(slab.Codes, "cmd/transitions/temperateforesttotundra/data.txt")
	if err != nil {
		log.Fatal(err)
	}
}
