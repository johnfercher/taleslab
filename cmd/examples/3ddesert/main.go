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
	assetLoader := assetloader.NewAssetLoader()
	biomeLoader := biomeloader.NewBiomeLoader(assetLoader)
	mapService := taleslabservices.NewMapService(biomeLoader, encoder)

	inputMap := &entities.Map{
		Biome: entities.DesertBiomeType,
		Ground: &entities.Ground{
			Width:             100,
			Length:            100,
			TerrainComplexity: 5,
			ForceBaseLand:     true,
		},
		Props: &entities.Props{
			StoneDensity: 300,
			TreeDensity:  180,
			MiscDensity:  350,
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
