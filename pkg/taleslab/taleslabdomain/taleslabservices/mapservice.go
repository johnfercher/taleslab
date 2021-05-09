package taleslabservices

import (
	"context"
	"github.com/johnfercher/taleslab/internal/api/apierror"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdto"
)

type MapService interface {
	Generate(ctx context.Context, inputMap *taleslabdto.MapDtoRequest) (*taleslabdto.MapDtoResponse, apierror.ApiError)
}
