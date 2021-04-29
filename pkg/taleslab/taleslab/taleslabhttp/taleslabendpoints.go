package taleslabhttp

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/johnfercher/taleslab/pkg/taleslab/domain/entities"
	"github.com/johnfercher/taleslab/pkg/taleslab/domain/services"
)

func MakeGenerateMap(service services.MapService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		inputMap := request.(*entities.Map)
		return service.Generate(ctx, inputMap)
	}
}
