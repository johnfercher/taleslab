package proceduralservices

import (
	"errors"
	"github.com/johnfercher/taleslab/pkg/procedural/proceduraldomain/proceduralentities"
	"github.com/johnfercher/taleslab/pkg/procedural/proceduraldomain/proceduralservices"
	"github.com/johnfercher/taleslab/pkg/shared/grid"
	"github.com/johnfercher/taleslab/pkg/shared/rand"

	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
)

type proceduralGridGenerator struct {
	ground    *proceduralentities.Ground
	mountains *proceduralentities.Mountains
	river     *grid.River
	canyon    *proceduralentities.Canyon
}

func NewMatrixGenerator() proceduralservices.ProceduralGridGenerator {
	return &proceduralGridGenerator{}
}

func (m *proceduralGridGenerator) SetGround(ground *proceduralentities.Ground) proceduralservices.ProceduralGridGenerator {
	m.ground = ground
	return m
}

func (m *proceduralGridGenerator) SetMountains(mountains *proceduralentities.Mountains) proceduralservices.ProceduralGridGenerator {
	if mountains == nil {
		return m
	}
	m.mountains = mountains
	return m
}

func (m *proceduralGridGenerator) SetRiver(river *grid.River) proceduralservices.ProceduralGridGenerator {
	if river != nil {
		m.river = river
	}

	return m
}

func (m *proceduralGridGenerator) SetCanyon(canyon *proceduralentities.Canyon) proceduralservices.ProceduralGridGenerator {
	if canyon != nil {
		m.canyon = canyon
	}

	return m
}

func (m *proceduralGridGenerator) Generate() ([][]taleslabentities.Element, error) {
	if m.ground == nil {
		return nil, errors.New("ground must be provided")
	}

	world := grid.GenerateTerrain(m.ground.Width, m.ground.Length, 2.0, 2.0,
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

	if m.river != nil {
		world = grid.DigRiver(world, m.river)
	}

	// grid.PrintHeights(world)

	return world, nil
}

func (m *proceduralGridGenerator) generateMountainsGrid(minHeight int) []taleslabentities.ElementMatrix {
	mountains := []taleslabentities.ElementMatrix{}

	iCount := rand.Intn(m.mountains.RandComplexity) + m.mountains.MinComplexity
	jCount := rand.Intn(m.mountains.RandComplexity) + m.mountains.MinComplexity

	for i := 0; i < iCount; i++ {
		for j := 0; j < jCount; j++ {
			bothAxis := rand.DifferentIntn(10, "bothAxis")

			balancedWidth := m.mountains.MinX
			balancedRandWidth := rand.DifferentIntn(m.mountains.RandX, "balancedRandWidth")
			mountainX := balancedWidth + balancedRandWidth + bothAxis

			balancedLength := m.mountains.MinY
			balancedRandLength := rand.DifferentIntn(m.mountains.RandY, "balancedRandLength")
			mountainY := balancedLength + balancedRandLength + bothAxis

			gain := float64(rand.DifferentIntn(m.mountains.RandHeight, "m.mountains.RandHeight") + m.mountains.MinHeight)

			mountain := grid.GenerateMountain(mountainX, mountainY, gain, minHeight)
			mountains = append(mountains, mountain)
		}
	}

	return mountains
}
