package main

import (
	"context"
	"fmt"
	"github.com/johnfercher/talescoder/pkg/encoder"
	"github.com/johnfercher/taleslab/pkg/georeferencing/georeferencingservices"
	"github.com/johnfercher/taleslab/pkg/shared/file"
	"github.com/johnfercher/taleslab/pkg/shared/grid"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts/biometype"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabmappers"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabrepositories"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabservices"
	"github.com/johnfercher/tessadem-sdk/pkg/tessadem"
	"log"
)

func main() {
	worldMatrix, err := generateWorldFromTessadem()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	river := &grid.River{
		HeightCutThreshold: 2,
	}
	worldMatrix = grid.DigRiver(worldMatrix, river)

	encoder := encoder.NewEncoder()
	propRepository, _ := taleslabrepositories.NewPropRepository()
	biomeRepository, _ := taleslabrepositories.NewBiomeRepository()
	sliceGenerator := taleslabservices.NewSlabSliceGenerator(biomeRepository, propRepository)
	slabGenerator := taleslabservices.NewSlabGenerator(sliceGenerator)

	slabGeneration := &taleslabentities.SlabGeneration{
		SliceSize: 50,
		Biomes:    []biometype.BiomeType{biometype.TemperateForest},
		World:     worldMatrix,
	}

	slabs, err := slabGenerator.Generate(slabGeneration)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	var sliceCodes [][]string
	for _, slab := range slabs {
		var codes []string
		for _, line := range slab {
			talespireSlab := taleslabmappers.TaleSpireSlabFromSlab(line)
			base64, err := encoder.Encode(talespireSlab)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			codes = append(codes, base64)
		}
		sliceCodes = append(sliceCodes, codes)
	}

	err = file.SaveCodes(sliceCodes, "cmd/elevations/petropolis/data.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func generateWorldFromTessadem() ([][]taleslabentities.Element, error) {
	ctx := context.TODO()
	fileReader := tessadem.NewFileReader()

	areaResponse, err := fileReader.ReadArea(ctx, "data/elevation/petropolis.json")
	if err != nil {
		return nil, err
	}

	geoGenerator := georeferencingservices.NewGeoReferencingGridGenerator()
	return geoGenerator.Generate(areaResponse), nil
}
