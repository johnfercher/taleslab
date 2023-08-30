package main

import (
	"context"
	"github.com/johnfercher/taleslab/internal/bytecompressor"
	"github.com/johnfercher/taleslab/internal/file"
	"github.com/johnfercher/taleslab/internal/talespireadapter/talespirecoder"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdto"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabrepositories"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabservices"
	"log"
)

func main() {
	ctx := context.TODO()

	compressor := bytecompressor.New()
	encoder := talespirecoder.NewEncoder(compressor)

	propRepository := taleslabrepositories.NewPropRepository()

	biomeRepository := taleslabrepositories.NewBiomeRepository(propRepository)
	secondaryBiomeRepository := taleslabrepositories.NewBiomeRepository(propRepository)
	mapService := taleslabservices.NewMapService(biomeRepository, secondaryBiomeRepository, encoder)

	inputMap := &taleslabdto.MapDtoRequest{
		Biome:          taleslabconsts.TemperateForestBiomeType,
		SecondaryBiome: taleslabconsts.TundraBiomeType,
		Ground: &taleslabdto.GroundDtoRequest{
			Width:             70,
			Length:            70,
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
		River: &taleslabdto.RiverDtoRequest{
			HasRiver: false,
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

	err = file.SaveCodes(slab.Codes, "docs/codes/3dtemperateforest2tundra.txt")
	if err != nil {
		log.Fatal(err)
	}
}
