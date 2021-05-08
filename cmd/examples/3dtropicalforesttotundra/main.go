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
		Biome:          taleslabconsts.TemperateForestBiomeType,
		SecondaryBiome: taleslabconsts.TundraBiomeType,
		Ground: &taleslabentities.Ground{
			Width:             70,
			Length:            70,
			TerrainComplexity: 5,
			MinHeight:         5,
			ForceBaseLand:     true,
		},
		Props: &taleslabentities.Props{
			StoneDensity: 150,
			TreeDensity:  11,
			MiscDensity:  11,
		},
		Mountains: &taleslabentities.Mountains{
			MinX:           30,
			RandX:          5,
			MinY:           30,
			RandY:          5,
			MinComplexity:  5,
			RandComplexity: 2,
			MinHeight:      10,
			RandHeight:     10,
		},
		River: &taleslabentities.River{
			HasRiver: false,
		},
		Canyon: &taleslabentities.Canyon{
			HasCanyon:    false,
			CanyonOffset: 10,
		},
	}

	slab, err := mapService.Generate(ctx, inputMap)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(slab.Code)
	fmt.Println(slab.Size)
}
