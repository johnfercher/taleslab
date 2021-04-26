package forestservices

import (
	"context"
	"fmt"
	"github.com/johnfercher/taleslab/internal/api/apierror"
	"github.com/johnfercher/taleslab/pkg/api/contracts"
	"github.com/johnfercher/taleslab/pkg/api/domain/entities"
	"github.com/johnfercher/taleslab/pkg/assetloader"
	"github.com/johnfercher/taleslab/pkg/slabdecoder"
)

type mapGenerator struct {
	loader  assetloader.AssetLoader
	encoder slabdecoder.Encoder
}

func NewMapGenerator(encoder slabdecoder.Encoder) *mapGenerator {
	return &mapGenerator{
		encoder: encoder,
		loader:  assetloader.NewAssetLoader(),
	}
}

func (self *mapGenerator) Generate(ctx context.Context, forest *entities.Map) (*contracts.Map, apierror.ApiError) {
	builder := New(self.loader, self.encoder).
		SetBiome(forest.Biome).
		SetMountains(forest.Mountains).
		SetGround(forest.Ground).
		SetProps(forest.Props).
		SetRiver(forest.River)

	base64, err := builder.Build()
	if err != nil {
		return nil, err
	}

	size := float64(len(base64) / 1024)
	sizeStr := fmt.Sprintf("%f Kb", size)

	return &contracts.Map{
		SlabVersion: "v2",
		Size:        sizeStr,
		Code:        base64,
	}, nil
}
