package services

import (
	"context"
	"github.com/johnfercher/taleslab/internal/api/apierror"
	"github.com/johnfercher/taleslab/pkg/api/contracts"
	"github.com/johnfercher/taleslab/pkg/api/domain/entities"
)

type MapService interface {
	Generate(ctx context.Context, inputMap *entities.Map) (*contracts.MapResponse, apierror.ApiError)
}
