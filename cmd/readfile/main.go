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

	areaResponse, err := fileReader.ReadArea(ctx, "data/petropolis.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	worldMatrix := BuildNormalizedElevationMap(areaResponse)

	fmt.Println(len(worldMatrix), len(worldMatrix[0]))

	biome := biometype.TemperateForest

	worldMatrixSlices := grid.SliceTerrain(worldMatrix, 50)

	encoder := encoder.NewEncoder()
	propRepository := taleslabrepositories.NewPropRepository()
	biomeRepository := taleslabrepositories.NewBiomeRepository()

	response := &taleslabdto.MapDtoResponse{
		SlabVersion: "v2",
	}

	for _, worldMatrix := range worldMatrixSlices {
		sliceCode := []string{}
		for _, slice := range worldMatrix {
			assetsGenerator := taleslabservices.NewAssetsGenerator(biomeRepository, propRepository).
				SetBiome(biome)

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

	err = file.SaveCodes(response.Codes, "docs/codes/pet2.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func BuildNormalizedElevationMap(response *tessadem.AreaResponse) [][]taleslabentities.Element {
	hasOcean := false

	min, _ := getMinMax(response)

	if min <= 0 {
		hasOcean = true
		for i := 0; i < len(response.Results); i++ {
			for j := 0; j < len(response.Results[i]); j++ {
				response.Results[i][j].Elevation += math.Abs(min)
			}
		}
	}

	min, _ = getMinMax(response)

	elevation := [][]taleslabentities.Element{}

	for i := 0; i < len(response.Results); i++ {
		array := []taleslabentities.Element{}
		for j := 0; j < len(response.Results[i]); j++ {
			elevation := int(response.Results[i][j].Elevation - min)
			element := taleslabentities.Element{
				elevation,
				getBaseGroundType(hasOcean, elevation),
			}

			array = append(array, element)
		}
		elevation = append(elevation, array)
	}

	return elevation
}

func getBaseGroundType(hasOcean bool, elevation int) taleslabconsts.ElementType {
	if hasOcean && elevation <= 1 {
		return taleslabconsts.Water
	}

	if elevation <= 3 {
		return taleslabconsts.BaseGround
	}

	return taleslabconsts.Ground
}

func getMinMax(response *tessadem.AreaResponse) (float64, float64) {
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

	return min, max
}
