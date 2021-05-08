package main

import (
	"context"
	"fmt"
	"github.com/johnfercher/taleslab/internal/bytecompressor"
	"github.com/johnfercher/taleslab/internal/talespireadapter/talespirecoder"
	"github.com/johnfercher/taleslab/pkg/assetloader"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabrepositories"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabservices"
	"log"
)

func main() {
	ctx := context.TODO()

	compressor := bytecompressor.New()
	encoder := talespirecoder.NewEncoder(compressor)

	assetLoader, err := assetloader.NewAssetLoader()
	if err != nil {
		log.Fatal(err.Error())
	}

	biomeLoader := taleslabrepositories.NewBiomeRepository(assetLoader)
	secondaryBiomeLoader := taleslabrepositories.NewBiomeRepository(assetLoader)
	mapService := taleslabservices.NewMapService(biomeLoader, secondaryBiomeLoader, encoder)

	inputMap := &taleslabentities.Map{
		Biome: taleslabconsts.DesertBiomeType,
		Ground: &taleslabentities.Ground{
			Width:             100,
			Length:            100,
			TerrainComplexity: 5,
			ForceBaseLand:     true,
		},
		Props: &taleslabentities.Props{
			StoneDensity: 300,
			TreeDensity:  180,
			MiscDensity:  350,
		},
		River: &taleslabentities.River{
			HasRiver: false,
		},
	}

	slab, err := mapService.Generate(ctx, inputMap)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(slab.Code)
	fmt.Println(slab.Size)
}
