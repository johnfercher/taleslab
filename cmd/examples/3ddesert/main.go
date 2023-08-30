package main

import (
	"context"
	"github.com/johnfercher/taleslab/internal/helper/bytecompressor"
	"github.com/johnfercher/taleslab/internal/helper/file"
	"github.com/johnfercher/taleslab/internal/helper/talespireadapter/talespirecoder"
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
		Biome: taleslabconsts.DesertBiomeType,
		Ground: &taleslabdto.GroundDtoRequest{
			Width:             100,
			Length:            100,
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

	err = file.SaveCodes(slab.Codes, "docs/codes/3ddesert.txt")
	if err != nil {
		log.Fatal(err)
	}
}
