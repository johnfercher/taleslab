package taleslabservices

import (
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts/elementtype"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabservices"
	"github.com/johnfercher/tessadem-sdk/pkg/tessadem"
	"math"
)

type geoReferencingGridGenerator struct {
	hasElevationAtOceanLevel bool
	waterHeightLimit         int
	baseGroundHeightLimit    int
	groundHeightLimit        int
}

func NewGeoReferencingGridGenerator() taleslabservices.GeoReferencingGridGenerator {
	return &geoReferencingGridGenerator{
		waterHeightLimit:      1,
		baseGroundHeightLimit: 3,
		groundHeightLimit:     10,
	}
}

func (g *geoReferencingGridGenerator) Generate(tessademData *tessadem.AreaResponse) [][]taleslabentities.Element {
	tessademData = g.normalizeInput(tessademData)

	min, _ := g.getMinMax(tessademData)

	elevation := [][]taleslabentities.Element{}

	for i := 0; i < len(tessademData.Results); i++ {
		array := []taleslabentities.Element{}
		for j := 0; j < len(tessademData.Results[i]); j++ {
			elevation := int(tessademData.Results[i][j].Elevation - min)
			element := taleslabentities.Element{
				elevation,
				g.getBaseGroundType(elevation),
			}

			array = append(array, element)
		}
		elevation = append(elevation, array)
	}

	return elevation
}

func (g *geoReferencingGridGenerator) getMinMax(response *tessadem.AreaResponse) (float64, float64) {
	min := math.MaxFloat64
	max := 0.0

	for i := 0; i < len(response.Results); i++ {
		for j := 0; j < len(response.Results[i]); j++ {
			elevation := response.Results[i][j].Elevation
			if elevation < min {
				min = elevation
			} else if elevation > max {
				max = elevation
			}
		}
	}

	return min, max
}

func (g *geoReferencingGridGenerator) getBaseGroundType(elevation int) elementtype.ElementType {
	if g.hasElevationAtOceanLevel && elevation <= g.waterHeightLimit {
		return elementtype.Water
	}

	if elevation <= g.baseGroundHeightLimit {
		return elementtype.BaseGround
	}

	if elevation <= g.baseGroundHeightLimit {
		return elementtype.Ground
	}

	return elementtype.Mountain
}

func (g *geoReferencingGridGenerator) normalizeInput(tessademData *tessadem.AreaResponse) *tessadem.AreaResponse {
	min, _ := g.getMinMax(tessademData)

	if min <= 0 {
		g.hasElevationAtOceanLevel = true
		for i := 0; i < len(tessademData.Results); i++ {
			for j := 0; j < len(tessademData.Results[i]); j++ {
				tessademData.Results[i][j].Elevation += math.Abs(min)
			}
		}
	}

	return tessademData
}
