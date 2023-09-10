package georeferencingservices

import (
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
	"github.com/johnfercher/tessadem-sdk/pkg/tessadem"
)

type GeoReferencingGridGenerator interface {
	Generate(tessademData *tessadem.AreaResponse) [][]taleslabentities.Element
}
