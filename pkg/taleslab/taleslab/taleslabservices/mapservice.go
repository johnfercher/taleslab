package taleslabservices

import (
	"context"
	"fmt"
	"github.com/johnfercher/taleslab/internal/api/apierror"
	"github.com/johnfercher/taleslab/pkg/assetloader"
	"github.com/johnfercher/taleslab/pkg/taleslab/contracts"
	"github.com/johnfercher/taleslab/pkg/taleslab/domain/entities"
	"github.com/johnfercher/taleslab/pkg/talespire/talespirecoder"
)

type mapService struct {
	loader  assetloader.AssetLoader
	encoder talespirecoder.Encoder
}

func NewMapService(encoder talespirecoder.Encoder) *mapService {
	return &mapService{
		encoder: encoder,
		loader:  assetloader.NewAssetLoader(),
	}
}

func (self *mapService) Generate(ctx context.Context, inputMap *entities.Map) (*contracts.MapResponse, apierror.ApiError) {
	builder := New(self.loader, self.encoder).
		SetBiome(inputMap.Biome).
		SetMountains(inputMap.Mountains).
		SetGround(inputMap.Ground).
		SetProps(inputMap.Props).
		SetRiver(inputMap.River)

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
