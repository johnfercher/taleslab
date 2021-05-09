package taleslabhttp

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabservices"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdto"
)

func MakeGenerateMap(service taleslabservices.MapService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		inputMap := request.(*taleslabdto.MapDtoRequest)
		return service.Generate(ctx, inputMap)
	}
}

func MakeGetGenerationsCount(service taleslabservices.MapService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return service.GetGenerationsCount(ctx)
	}
}
