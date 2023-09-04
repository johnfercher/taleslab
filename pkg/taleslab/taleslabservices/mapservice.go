package taleslabservices

import (
	"context"
	"github.com/johnfercher/talescoder/pkg/encoder"
	"github.com/johnfercher/taleslab/internal/api/apierror"
	"github.com/johnfercher/taleslab/pkg/grid"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabrepositories"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabservices"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdto"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabmappers"
	"net/http"
)

type mapService struct {
	biomeRepository taleslabrepositories.BiomeRepository
	propsRepository taleslabrepositories.PropRepository
	encoder         encoder.Encoder
}

func NewMapService(biomeRepository taleslabrepositories.BiomeRepository, propsRepository taleslabrepositories.PropRepository, encoder encoder.Encoder) taleslabservices.SlabGenerator {
	return &mapService{
		biomeRepository: biomeRepository,
		propsRepository: propsRepository,
		encoder:         encoder,
	}
}

func (self *mapService) Generate(ctx context.Context, inputMap *taleslabdto.MapDtoRequest) (*taleslabdto.MapDtoResponse, apierror.ApiError) {
	matrixGenerator := NewMatrixGenerator().
		SetMountains(inputMap.Mountains).
		SetGround(inputMap.Ground).
		SetRiver(inputMap.River).
		SetCanyon(inputMap.Canyon)

	worldMatrix, err := matrixGenerator.Generate()
	if err != nil {
		return nil, err
	}

	worldMatrixSlices := grid.SliceTerrain(worldMatrix, 50)

	response := &taleslabdto.MapDtoResponse{
		SlabVersion: "v2",
	}

	for _, worldMatrix := range worldMatrixSlices {
		sliceCode := []string{}
		for _, slice := range worldMatrix {
			assetsGenerator := NewAssetsGenerator(self.biomeRepository, self.propsRepository).
				SetBiome(inputMap.Biome).
				SetProps(inputMap.Props).
				SetSecondaryBiome(inputMap.SecondaryBiome)

			worldAssets, err := assetsGenerator.Generate(slice)
			if err != nil {
				return nil, err
			}

			slab := taleslabmappers.TaleSpireSlabFromAssets(worldAssets)

			base64, encodeError := self.encoder.Encode(slab)
			if err != nil {
				return nil, apierror.New(http.StatusInternalServerError, encodeError.Error())
			}

			sliceCode = append(sliceCode, base64)
			response.Size += len(base64) / 1024
		}

		response.Codes = append(response.Codes, sliceCode)
	}

	return response, nil
}
