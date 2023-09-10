package main

import (
	"context"
	"github.com/johnfercher/talescoder/pkg/encoder"
	"github.com/johnfercher/taleslab/pkg/file"
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
		Biome: biometype.Desert,
		Ground: &taleslabdto.GroundDtoRequest{
			Width:             50,
			Length:            50,
			TerrainComplexity: 5,
			ForceBaseLand:     true,
		},
		Props: &taleslabdto.PropsDtoRequest{
			StoneDensity: 300,
			TreeDensity:  180,
			MiscDensity:  350,
		},
		River: &taleslabdto.RiverDtoRequest{
			HasRiver: false,
		},
	}

	slab, err := mapService.Generate(ctx, inputMap)
	if err != nil {
		log.Fatal(err)
	}

	err = file.SaveCodes(slab.Codes, "docs/codes/3ddesert/data.txt")
	if err != nil {
		log.Fatal(err)
	}
}
