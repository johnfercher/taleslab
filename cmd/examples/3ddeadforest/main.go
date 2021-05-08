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
		Biome: taleslabconsts.DeadForestBiomeType,
		Ground: &taleslabcontracts.Ground{
			Width:             80,
			Length:            80,
			TerrainComplexity: 5,
			ForceBaseLand:     true,
		},
		Props: &taleslabcontracts.Props{
			StoneDensity: 300,
			TreeDensity:  15,
			MiscDensity:  130,
		},
		River: &taleslabcontracts.River{
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
