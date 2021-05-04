package taleslabservices

import (
	"context"
	"fmt"
	"github.com/johnfercher/taleslab/internal/api/apierror"
	"github.com/johnfercher/taleslab/pkg/biomeloader"
	"github.com/johnfercher/taleslab/pkg/taleslab/contracts"
	"github.com/johnfercher/taleslab/pkg/taleslab/domain/entities"
	"github.com/johnfercher/taleslab/pkg/talespire/talespirecoder"
)

type mapService struct {
	biomeLoader          biomeloader.BiomeLoader
	secondaryBiomeLoader biomeloader.BiomeLoader
	encoder              talespirecoder.Encoder
}

func NewMapService(biomeLoader biomeloader.BiomeLoader, secondaryBiomeLoader biomeloader.BiomeLoader, encoder talespirecoder.Encoder) *mapService {
	return &mapService{
		biomeLoader:          biomeLoader,
		encoder:              encoder,
		secondaryBiomeLoader: secondaryBiomeLoader,
	}
}

func (self *mapService) Generate(ctx context.Context, inputMap *entities.Map) (*contracts.MapResponse, apierror.ApiError) {
	builder := New(self.biomeLoader, self.secondaryBiomeLoader, self.encoder).
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

	return &contracts.MapResponse{
		SlabVersion: "v2",
		Size:        sizeStr,
		Code:        base64,
	}, nil
}
