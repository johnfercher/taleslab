package taleslabservices

import (
	"context"
	"github.com/johnfercher/taleslab/internal/api/apierror"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabcontracts"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
)

type MapService interface {
	Generate(ctx context.Context, inputMap *taleslabentities.Map) (*taleslabcontracts.MapResponse, apierror.ApiError)
}
