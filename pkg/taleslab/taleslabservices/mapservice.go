package taleslabservices

import (
	"context"
	"fmt"
	"github.com/johnfercher/taleslab/internal/api/apierror"
	"github.com/johnfercher/taleslab/internal/talespireadapter/talespirecoder"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabrepositories"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdto"
)

type mapService struct {
	biomeLoader          taleslabrepositories.BiomeRepository
	secondaryBiomeLoader taleslabrepositories.BiomeRepository
	encoder              talespirecoder.Encoder
}

func NewMapService(biomeLoader taleslabrepositories.BiomeRepository, secondaryBiomeLoader taleslabrepositories.BiomeRepository, encoder talespirecoder.Encoder) *mapService {
	return &mapService{
		biomeLoader:          biomeLoader,
		encoder:              encoder,
		secondaryBiomeLoader: secondaryBiomeLoader,
	}
}

func (self *mapService) Generate(ctx context.Context, inputMap *taleslabdto.MapDtoRequest) (*taleslabdto.MapDtoResponse, apierror.ApiError) {
	builder := NewMapBuilder(self.biomeLoader, self.secondaryBiomeLoader, self.encoder).
		SetBiome(inputMap.Biome).
		SetSecondaryBiome(inputMap.SecondaryBiome).
		SetMountains(inputMap.Mountains).
		SetGround(inputMap.Ground).
		SetProps(inputMap.Props).
		SetRiver(inputMap.River).
		SetCanyon(inputMap.Canyon)

	base64, err := builder.Build()
	if err != nil {
		return nil, err
	}

	size := float64(len(base64) / 1024)
	sizeStr := fmt.Sprintf("%f Kb", size)

	return &taleslabdto.MapDtoResponse{
		SlabVersion: "v2",
		Size:        sizeStr,
		Code:        base64,
	}, nil
}
