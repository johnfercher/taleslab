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

	fmt.Println(len(worldMatrix), len(worldMatrix[0]))

	biome := biometype.Beach
	secondaryBiome := biometype.Beach
	props := &taleslabdto.PropsDtoRequest{
		StoneDensity: 100,
		TreeDensity:  15,
		MiscDensity:  15,
	}

	worldMatrixSlices := grid.SliceTerrain(worldMatrix, 50)

	encoder := encoder.NewEncoder()
	propRepository := taleslabrepositories.NewPropRepository()
	biomeRepository := taleslabrepositories.NewBiomeRepository(propRepository)
	secondaryBiomeRepository := taleslabrepositories.NewBiomeRepository(propRepository)

	response := &taleslabdto.MapDtoResponse{
		SlabVersion: "v2",
	}

	for _, worldMatrix := range worldMatrixSlices {
		sliceCode := []string{}
		for _, slice := range worldMatrix {
			assetsGenerator := taleslabservices.NewAssetsGenerator(biomeRepository, secondaryBiomeRepository).
				SetBiome(biome).
				SetProps(props).
				SetSecondaryBiome(secondaryBiome)

			worldAssets, err := assetsGenerator.Generate(slice)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			slab := taleslabmappers.TaleSpireSlabFromAssets(worldAssets)

			base64, encodeError := encoder.Encode(slab)
			if err != nil {
				fmt.Println(encodeError.Error())
				return
			}

			sliceCode = append(sliceCode, base64)
			response.Size += len(base64) / 1024
		}

		response.Codes = append(response.Codes, sliceCode)
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
	hasOcean := false

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

	if min <= 0 {
		fmt.Println(min)
		hasOcean = true
	}

	fmt.Println(min, max)

	elevation := [][]taleslabentities.Element{}

	for i := 0; i < len(response.Results); i++ {
		array := []taleslabentities.Element{}
		for j := 0; j < len(response.Results[i]); j++ {
			elevation := int(response.Results[i][j].Elevation - min)
			element := taleslabentities.Element{
				elevation,
				getBaseGroundType(hasOcean, int(response.Results[i][j].Elevation)),
			}

			array = append(array, element)
		}
		elevation = append(elevation, array)
	}

	return elevation
}

func getBaseGroundType(hasOcean bool, elevation int) taleslabconsts.ElementType {
	if !hasOcean {
		return taleslabconsts.GroundType
	}

	if elevation <= 0 {
		return taleslabconsts.WaterType
	}

	if elevation <= 2 {
		return taleslabconsts.BaseGroundType
	}

	return taleslabconsts.GroundType
}
