package taleslabservices

import (
	"github.com/johnfercher/taleslab/internal/api/apierror"
	"github.com/johnfercher/taleslab/internal/math"
	"github.com/johnfercher/taleslab/pkg/grid"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabservices"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdto"
	"math/rand"
	"time"
)

type matrixGenerator struct {
	ground    *taleslabdto.GroundDtoRequest
	mountains *taleslabdto.MountainsDtoRequest
	river     *taleslabdto.RiverDtoRequest
	canyon    *taleslabdto.CanyonDtoRequest
}

func NewMatrixGenerator() taleslabservices.MatrixGenerator {
	return &matrixGenerator{}
}

func (m *matrixGenerator) SetGround(ground *taleslabdto.GroundDtoRequest) taleslabservices.MatrixGenerator {
	m.ground = ground
	return m
}

func (m *matrixGenerator) SetMountains(mountains *taleslabdto.MountainsDtoRequest) taleslabservices.MatrixGenerator {
	if mountains == nil {
		return m
	}
	m.mountains = mountains
	return m
}

func (m *matrixGenerator) SetRiver(river *taleslabdto.RiverDtoRequest) taleslabservices.MatrixGenerator {
	if river != nil {
		m.river = river
	}

	return m
}

func (m *matrixGenerator) SetCanyon(canyon *taleslabdto.CanyonDtoRequest) taleslabservices.MatrixGenerator {
	if canyon != nil {
		m.canyon = canyon
	}

	return m
}

func (m *matrixGenerator) Generate() ([][]taleslabentities.Element, apierror.ApiError) {
	if m.ground == nil {
		return nil, apierror.New(400, "Ground must be provided")
	}

	world := grid.TerrainGenerator(m.ground.Width, m.ground.Length, 2.0, 2.0,
		m.ground.TerrainComplexity, m.ground.MinHeight, m.ground.ForceBaseLand)

	if m.mountains != nil {
		mountains := m.generateMountainsGrid(m.ground.MinHeight)
		for _, mountain := range mountains {
			world = grid.AppendTerrainRandomly(world, mountain)
		}
	}

	if m.canyon != nil && m.canyon.HasCanyon {
		world = grid.DigTerrainInOffset(world, m.canyon.CanyonOffset)
	}

	if m.river != nil && m.river.HasRiver {
		world = grid.DigRiver(world)
	}

	//grid.PrintHeights(world)

	return world, nil
}

func (m *matrixGenerator) generateMountainsGrid(minHeight int) []taleslabentities.ElementMatrix {
	mountains := []taleslabentities.ElementMatrix{}

	rand.Seed(time.Now().UnixNano())
	iCount := rand.Intn(m.mountains.RandComplexity) + m.mountains.MinComplexity

	rand.Seed(time.Now().UnixNano())
	jCount := rand.Intn(m.mountains.RandComplexity) + m.mountains.MinComplexity

	for i := 0; i < iCount; i++ {
		for j := 0; j < jCount; j++ {
			bothAxis := math.GetRandomValue(10, "bothAxis")

			balancedWidth := m.mountains.MinX
			balancedRandWidth := math.GetRandomValue(m.mountains.RandX, "balancedRandWidth")
			mountainX := balancedWidth + balancedRandWidth + bothAxis

			balancedLength := m.mountains.MinY
			balancedRandLength := math.GetRandomValue(m.mountains.RandY, "balancedRandLength")
			mountainY := balancedLength + balancedRandLength + bothAxis

			gain := float64(math.GetRandomValue(m.mountains.RandHeight, "m.mountains.RandHeight") + m.mountains.MinHeight)

			mountain := grid.MountainGenerator(mountainX, mountainY, gain, minHeight)
			mountains = append(mountains, mountain)
		}
	}

	return mountains
}
