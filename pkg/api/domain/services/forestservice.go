package services

import (
	"context"
	"github.com/johnfercher/taleslab/internal/api/apierror"
	"github.com/johnfercher/taleslab/pkg/api/contracts"
	"github.com/johnfercher/taleslab/pkg/api/domain/entities"
)

type ForestService interface {
	GenerateForest(ctx context.Context, contact *entities.Forest) (contracts.Slab, apierror.ApiError)
}
