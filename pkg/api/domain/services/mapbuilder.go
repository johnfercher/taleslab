package services

import (
	"github.com/johnfercher/taleslab/internal/api/apierror"
	"github.com/johnfercher/taleslab/pkg/api/domain/entities"
)

type MapBuilder interface {
	SetBiome(biome entities.Biome) MapBuilder
	SetGround(ground *entities.Ground) MapBuilder
	SetMountains(mountains *entities.Mountains) MapBuilder
	SetRiver(river *entities.River) MapBuilder
	SetProps(props *entities.Props) MapBuilder
	Build() (string, apierror.ApiError)
}
