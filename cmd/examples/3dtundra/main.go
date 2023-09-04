package main

import (
	"context"
	"github.com/johnfercher/talescoder/pkg/encoder"
	"github.com/johnfercher/taleslab/internal/file"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts/biometype"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdto"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabrepositories"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabservices"
	"log"
)

func main() {
	ctx := context.TODO()

	encoder := encoder.NewEncoder()

	propRepository := taleslabrepositories.NewPropRepository()
	biomeRepository := taleslabrepositories.NewBiomeRepository()
	mapService := taleslabservices.NewMapService(biomeRepository, propRepository, encoder)

	inputMap := &taleslabdto.MapDtoRequest{
		Biome: biometype.Tundra,
		Ground: &taleslabdto.GroundDtoRequest{
			Width:             50,
			Length:            50,
			TerrainComplexity: 5,
			ForceBaseLand:     true,
		},
		Props: &taleslabdto.PropsDtoRequest{
			StoneDensity: 100,
			TreeDensity:  15,
			MiscDensity:  15,
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
			HasRiver: true,
		},
	}

	slab, err := mapService.Generate(ctx, inputMap)
	if err != nil {
		log.Fatal(err)
	}

	err = file.SaveCodes(slab.Codes, "docs/codes/3dtundra.txt")
	if err != nil {
		log.Fatal(err)
	}
}
