package proceduralservices

import (
	"context"
	"github.com/johnfercher/taleslab/pkg/procedural/proceduraldomain/proceduralentities"
)

type MapService interface {
	Generate(ctx context.Context, inputMap *proceduralentities.MapGeneration) (*proceduralentities.MapGenerated, error)
}
