package taleslabservices

import (
	"context"

	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdto"
)

type MapService interface {
	Generate(ctx context.Context, inputMap *taleslabdto.MapDtoRequest) (*taleslabdto.MapDtoResponse, error)
}
