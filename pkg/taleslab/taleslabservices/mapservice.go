package taleslabservices

import (
	"context"

	"github.com/johnfercher/talescoder/pkg/encoder"
	"github.com/johnfercher/taleslab/pkg/grid"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabrepositories"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabservices"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdto"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabmappers"
)

type mapService struct {
	biomeRepository taleslabrepositories.BiomeRepository
	propsRepository taleslabrepositories.PropRepository
	encoder         encoder.Encoder
}

func NewMapService(biomeRepository taleslabrepositories.BiomeRepository, propsRepository taleslabrepositories.PropRepository,
	encoder encoder.Encoder,
) taleslabservices.SlabGenerator {
	return &mapService{
		biomeRepository: biomeRepository,
		propsRepository: propsRepository,
		encoder:         encoder,
	}
}

func (m *mapService) Generate(ctx context.Context, inputMap *taleslabdto.MapDtoRequest) (*taleslabdto.MapDtoResponse, error) {
	matrixGenerator := NewMatrixGenerator().
		SetMountains(inputMap.Mountains).
		SetGround(inputMap.Ground).
		SetRiver(inputMap.River).
		SetCanyon(inputMap.Canyon)

	worldMatrix, err := matrixGenerator.Generate()
	if err != nil {
		return nil, err
	}

	maxWidth := len(worldMatrix)
	maxLength := len(worldMatrix[0])
	squareSize := 50

	worldMatrixSlices := grid.SliceTerrain(worldMatrix, squareSize)

	response := &taleslabdto.MapDtoResponse{
		SlabVersion: "v2",
	}

	currentX := 0
	currentY := 0
	for _, worldMatrix := range worldMatrixSlices {
		sliceCode := []string{}
		for _, slice := range worldMatrix {
			assetsGenerator := NewAssetsGenerator(m.biomeRepository, m.propsRepository, maxWidth, maxLength).
				SetBiome(inputMap.Biome).
				SetSecondaryBiome(inputMap.SecondaryBiome)

			worldAssets, err := assetsGenerator.Generate(slice, currentX, currentY)
			if err != nil {
				return nil, err
			}

			slab := taleslabmappers.TaleSpireSlabFromAssets(worldAssets)

			base64, err := m.encoder.Encode(slab)
			if err != nil {
				return nil, err
			}

			sliceCode = append(sliceCode, base64)
			response.Size += len(base64) / 1024
			currentX += squareSize
		}

		response.Codes = append(response.Codes, sliceCode)
		currentX = 0
		currentY += squareSize
	}

	return response, nil
}
