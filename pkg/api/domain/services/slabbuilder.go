package services

import (
	"github.com/johnfercher/taleslab/internal/api/apierror"
	"github.com/johnfercher/taleslab/pkg/api/domain/entities"
)

type SlabBuilder interface {
	SetBiome(biome entities.Biome) SlabBuilder
	SetGround(ground *entities.Ground) SlabBuilder
	SetMountains(mountains *entities.Mountains) SlabBuilder
	SetRiver(river *entities.River) SlabBuilder
	SetProps(props *entities.Props) SlabBuilder
	Build() (string, apierror.ApiError)
}
