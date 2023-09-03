package main

import (
	"context"
	"fmt"
	"github.com/johnfercher/talescoder/pkg/encoder"
	"github.com/johnfercher/taleslab/internal/file"
	"github.com/johnfercher/taleslab/pkg/grid"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts/biometype"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdto"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabmappers"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabrepositories"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabservices"
	"github.com/johnfercher/tessadem-sdk/pkg/tessadem"
	"math"
)

func main() {
	ctx := context.TODO()

	fileReader := tessadem.NewFileReader()

	areaResponse, err := fileReader.ReadArea(ctx, "file.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	worldMatrix := BuildNormalizedElevationMap(areaResponse)

	inputMap := &taleslabdto.MapDtoRequest{
		Biome: biometype.Tundra,
		Ground: &taleslabdto.GroundDtoRequest{
			Width:             128,
			Length:            128,
			TerrainComplexity: 5,
			ForceBaseLand:     true,
		},
		Props: &taleslabdto.PropsDtoRequest{
			StoneDensity: 100,
			TreeDensity:  15,
			MiscDensity:  15,
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
		River: &taleslabdto.RiverDtoRequest{
			HasRiver: true,
		},
	}

	worldSlices := grid.SliceTerrain(worldMatrix, 50)

	encoder := encoder.NewEncoder()
	propRepository := taleslabrepositories.NewPropRepository()
	biomeRepository := taleslabrepositories.NewBiomeRepository(propRepository)
	secondaryBiomeRepository := taleslabrepositories.NewBiomeRepository(propRepository)

	response := &taleslabdto.MapDtoResponse{
		SlabVersion: "v2",
	}

	for _, slice := range worldSlices {
		assetsGenerator := taleslabservices.NewAssetsGenerator(biomeRepository, secondaryBiomeRepository).
			SetBiome(inputMap.Biome).
			SetProps(inputMap.Props).
			SetSecondaryBiome(inputMap.SecondaryBiome)

		worldAssets, apiErr := assetsGenerator.Generate(slice)
		if apiErr != nil {
			fmt.Println(err.Error())
			return
		}

		slab := taleslabmappers.TaleSpireSlabFromAssets(worldAssets)

		base64, err := encoder.Encode(slab)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		response.Codes = append(response.Codes, base64)
		response.Size += len(base64) / 1024
	}

	err = file.SaveCodes(response.Codes, "docs/codes/pet.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func BuildNormalizedElevationMap(response *tessadem.AreaResponse) [][]taleslabentities.Element {
	min := math.MaxFloat64
	max := 0.0

	for i := 0; i < len(response.Results); i++ {
		for j := 0; j < len(response.Results[i]); j++ {
			elevation := response.Results[i][j].Elevation
			if elevation < min {
				min = elevation
			} else if elevation > max {
				max = elevation
			}
		}
	}

	fmt.Println(min, max)

	elevation := [][]taleslabentities.Element{}

	for i := 0; i < len(response.Results); i++ {
		array := []taleslabentities.Element{}
		for j := 0; j < len(response.Results[i]); j++ {
			element := taleslabentities.Element{
				int(response.Results[i][j].Elevation - min),
				taleslabconsts.GroundType,
			}
			array = append(array, element)
		}
		elevation = append(elevation, array)
	}

	return elevation
}
