package services

import (
	"github.com/johnfercher/taleslab/pkg/api/domain/entities"
)

type SlabBuilder interface {
	Init() SlabBuilder
	DefineGround(forest *entities.Forest) SlabBuilder
}
