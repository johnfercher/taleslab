package main

import (
	"context"
	"github.com/johnfercher/taleslab/pkg/grid"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
	"log"

	"github.com/johnfercher/talescoder/pkg/encoder"
	"github.com/johnfercher/taleslab/pkg/file"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts/biometype"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdto"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabrepositories"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabservices"
)

func main() {
	ctx := context.TODO()

	encoder := encoder.NewEncoder()

	propRepository, _ := taleslabrepositories.NewPropRepository()
	biomeRepository, _ := taleslabrepositories.NewBiomeRepository()

	mapService := taleslabservices.NewMapService(biomeRepository, propRepository, encoder)

	inputMap := &taleslabdto.MapDtoRequest{
		Biome: biometype.Beach,
		Ground: &taleslabdto.GroundDtoRequest{
			Width:             50,
			Length:            50,
			TerrainComplexity: 5,
			MinHeight:         5,
			ForceBaseLand:     true,
		},
		Props: &taleslabdto.PropsDtoRequest{
			StoneDensity: 83,
			TreeDensity:  15,
		},
		Mountains: &taleslabdto.MountainsDtoRequest{
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
		Canyon: &taleslabdto.CanyonDtoRequest{
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
