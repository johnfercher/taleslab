package taleslabhttp

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabcontracts"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabservices"
)

func MakeGenerateMap(service taleslabservices.MapService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		inputMap := request.(*taleslabcontracts.Map)
		return service.Generate(ctx, inputMap)
	}
}
