package taleslabservices

import (
	"context"
	"github.com/johnfercher/taleslab/internal/api/apierror"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabcontracts"
)

type MapService interface {
	Generate(ctx context.Context, inputMap *taleslabcontracts.Map) (*taleslabcontracts.MapResponse, apierror.ApiError)
}
