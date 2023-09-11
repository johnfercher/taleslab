package main

import (
	"context"
	"github.com/johnfercher/taleslab/pkg/procedural/proceduraldomain/proceduralentities"
	"github.com/johnfercher/taleslab/pkg/procedural/proceduralservices"
	"github.com/johnfercher/taleslab/pkg/shared/file"
	"github.com/johnfercher/taleslab/pkg/shared/grid"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
	"log"

	"github.com/johnfercher/talescoder/pkg/encoder"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts/biometype"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabrepositories"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabservices"
)

func main() {
	ctx := context.TODO()

	encoder := encoder.NewEncoder()
	propRepository, _ := taleslabrepositories.NewPropRepository()
	biomeRepository, _ := taleslabrepositories.NewBiomeRepository()
	sliceGenerator := taleslabservices.NewSlabSliceGenerator(biomeRepository, propRepository)
	slabGenerator := taleslabservices.NewSlabGenerator(sliceGenerator)

	mapService := proceduralservices.NewMapService(slabGenerator, encoder)

	inputMap := &proceduralentities.MapGeneration{
		Biome: biometype.Beach,
		Ground: &proceduralentities.Ground{
			Width:             50,
			Length:            50,
			TerrainComplexity: 5,
			MinHeight:         5,
			ForceBaseLand:     true,
		},
		Props: &proceduralentities.Props{
			StoneDensity: 83,
			TreeDensity:  15,
		},
		Mountains: &proceduralentities.Mountains{
			MinX:           30,
			RandX:          5,
			MinY:           30,
			RandY:          5,
			MinComplexity:  5,
			RandComplexity: 2,
			MinHeight:      10,
			RandHeight:     10,
		},
		River: &grid.River{
			Start:              &taleslabentities.Vector3d{49, 49, 0},
			End:                &taleslabentities.Vector3d{0, 0, 0},
			HeightCutThreshold: 5,
		},
		Canyon: &proceduralentities.Canyon{
			HasCanyon:    true,
			CanyonOffset: 10,
		},
	}

	slab, err := mapService.Generate(ctx, inputMap)
	if err != nil {
		log.Fatal(err)
	}

	err = file.SaveCodes(slab.Codes, "cmd/procedurals/beach/data.txt")
	if err != nil {
		log.Fatal(err)
	}
}
