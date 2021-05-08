package main

import (
	"context"
	"fmt"
	"github.com/johnfercher/taleslab/internal/bytecompressor"
	"github.com/johnfercher/taleslab/internal/talespireadapter/talespirecoder"
	"github.com/johnfercher/taleslab/pkg/proploader"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabcontracts"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabrepositories"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabservices"
	"log"
)

func main() {
	ctx := context.TODO()

	compressor := bytecompressor.New()
	encoder := talespirecoder.NewEncoder(compressor)

	assetLoader, err := proploader.NewPropLoader()
	if err != nil {
		log.Fatal(err.Error())
	}

	biomeRepository := taleslabrepositories.NewBiomeRepository(assetLoader)
	secondaryBiomeRepository := taleslabrepositories.NewBiomeRepository(assetLoader)
	mapService := taleslabservices.NewMapService(biomeRepository, secondaryBiomeRepository, encoder)

	inputMap := &taleslabcontracts.Map{
		Biome: taleslabconsts.TemperateForestBiomeType,
		Ground: &taleslabcontracts.Ground{
			Width:             70,
			Length:            70,
			TerrainComplexity: 5,
		},
		Props: &taleslabcontracts.Props{
			StoneDensity: 150,
			TreeDensity:  15,
			MiscDensity:  25,
		},
		Mountains: &taleslabcontracts.Mountains{
			MinX:           30,
			RandX:          5,
			MinY:           30,
			RandY:          5,
			MinComplexity:  5,
			RandComplexity: 2,
			MinHeight:      10,
			RandHeight:     10,
		},
		River: &taleslabcontracts.River{
			HasRiver: true,
		},
	}

	slab, err := mapService.Generate(ctx, inputMap)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(slab.Code)
	fmt.Println(slab.Size)
}
