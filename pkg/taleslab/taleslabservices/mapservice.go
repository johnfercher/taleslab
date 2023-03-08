package taleslabservices

import (
	"context"
	"fmt"
	"github.com/johnfercher/taleslab/internal/api/apierror"
	"github.com/johnfercher/taleslab/internal/talespireadapter/talespirecoder"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabrepositories"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabservices"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdto"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabmappers"
	"net/http"
)

type mapService struct {
	biomeLoader          taleslabrepositories.BiomeRepository
	secondaryBiomeLoader taleslabrepositories.BiomeRepository
	encoder              talespirecoder.Encoder
}

func NewMapService(biomeLoader taleslabrepositories.BiomeRepository, secondaryBiomeLoader taleslabrepositories.BiomeRepository, encoder talespirecoder.Encoder) taleslabservices.SlabGenerator {
	return &mapService{
		biomeLoader:          biomeLoader,
		encoder:              encoder,
		secondaryBiomeLoader: secondaryBiomeLoader,
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

	assetsGenerator := NewAssetsGenerator(self.biomeLoader, self.secondaryBiomeLoader).
		SetBiome(inputMap.Biome).
		SetProps(inputMap.Props).
		SetSecondaryBiome(inputMap.SecondaryBiome)

	worldAssets, err := assetsGenerator.Generate(worldMatrix)
	if err != nil {
		return nil, err
	}

	slab := taleslabmappers.TaleSpireSlabFromAssets(worldAssets)

	base64, encodeError := self.encoder.Encode(slab)
	if err != nil {
		return nil, apierror.New(http.StatusInternalServerError, encodeError.Error())
	}

	size := float64(len(base64) / 1024)
	sizeStr := fmt.Sprintf("%f Kb", size)

	return &taleslabdto.MapDtoResponse{
		SlabVersion: "v2",
		Size:        sizeStr,
		Code:        base64,
	}, nil
}