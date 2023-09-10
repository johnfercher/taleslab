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

	propRepository := taleslabrepositories.NewPropRepository()
	biomeRepository := taleslabrepositories.NewBiomeRepository()
	mapService := taleslabservices.NewMapService(biomeRepository, propRepository, encoder)

	inputMap := &taleslabdto.MapDtoRequest{
		Biome: biometype.Lava,
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
		River: &taleslabdto.RiverDtoRequest{
			HasRiver: false,
		},
	}

	slab, err := mapService.Generate(ctx, inputMap)
	if err != nil {
		log.Fatal(err)
	}

	err = file.SaveCodes(slab.Codes, "cmd/procedurals/lava/data.txt")
	if err != nil {
		log.Fatal(err)
	}
}
