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
	biomeLoader := biomeloader.NewBiomeLoader()
	mapService := taleslabservices.NewMapService(assetLoader, biomeLoader, encoder)

	inputMap := &entities.Map{
		Biome: entities.DesertBiomeType,
		Ground: &entities.Ground{
			Width:             70,
			Length:            70,
			TerrainComplexity: 5,
			ForceBaseLand:     true,
		},
		Props: &entities.Props{
			PropsDensity: 223,
			TreeDensity:  113,
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
}
