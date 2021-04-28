package services

import (
	"context"
	"github.com/johnfercher/taleslab/internal/api/apierror"
	"github.com/johnfercher/taleslab/pkg/taleslab/contracts"
	"github.com/johnfercher/taleslab/pkg/taleslab/domain/entities"
)

type MapService interface {
	Generate(ctx context.Context, inputMap *entities.Map) (*contracts.MapResponse, apierror.ApiError)
}
