package services

import (
	"context"
	"github.com/johnfercher/taleslab/internal/api/apierror"
	"github.com/johnfercher/taleslab/pkg/api/contracts"
	"github.com/johnfercher/taleslab/pkg/api/domain/entities"
)

type ForestService interface {
	Generate(ctx context.Context, forest *entities.Map) (*contracts.Map, apierror.ApiError)
}
