package main

import (
	"context"
	"fmt"
	"github.com/johnfercher/taleslab/internal/bytecompressor"
	"github.com/johnfercher/taleslab/pkg/assetloader"
	"github.com/johnfercher/taleslab/pkg/biomeloader"
	"github.com/johnfercher/taleslab/pkg/taleslab/domain/entities"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslab/taleslabservices"
	"github.com/johnfercher/taleslab/pkg/talespire/talespirecoder"
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

	biomeLoader := biomeloader.NewBiomeLoader(assetLoader)
	secondaryBiomeLoader := biomeloader.NewBiomeLoader(assetLoader)
	mapService := taleslabservices.NewMapService(biomeLoader, secondaryBiomeLoader, encoder)

	inputMap := &entities.Map{
		Biome: entities.DesertTestType,
		Ground: &entities.Ground{
			Width:             50,
			Length:            50,
			TerrainComplexity: 5,
			ForceBaseLand:     true,
		},
		Props: &entities.Props{
			StoneDensity: 350,
			TreeDensity:  350,
			MiscDensity:  1000,
		},
		River: &entities.River{
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
