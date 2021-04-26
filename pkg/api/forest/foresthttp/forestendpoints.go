package foresthttp

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/johnfercher/taleslab/pkg/api/domain/entities"
	"github.com/johnfercher/taleslab/pkg/api/domain/services"
)

func MakeGenerateForest(service services.ForestService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		forest := request.(*entities.Forest)
		return service.GenerateForest(ctx, forest)
	}
}
