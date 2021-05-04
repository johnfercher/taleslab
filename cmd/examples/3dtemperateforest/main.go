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
		Biome: entities.TemperateForestBiomeType,
		Ground: &entities.Ground{
			Width:             70,
			Length:            70,
			TerrainComplexity: 5,
		},
		Props: &entities.Props{
			StoneDensity: 150,
			TreeDensity:  15,
			MiscDensity:  25,
		},
		Mountains: &entities.Mountains{
			MinX:           30,
			RandX:          5,
			MinY:           30,
			RandY:          5,
			MinComplexity:  5,
			RandComplexity: 2,
			MinHeight:      10,
			RandHeight:     10,
		},
		River: &entities.River{
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
