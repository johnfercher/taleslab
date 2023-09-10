package proceduralservices

import (
	"github.com/johnfercher/taleslab/pkg/procedural/proceduraldomain/proceduralentities"
	"github.com/johnfercher/taleslab/pkg/shared/grid"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
)

type ProceduralGridGenerator interface {
	SetGround(ground *proceduralentities.Ground) ProceduralGridGenerator
	SetMountains(mountains *proceduralentities.Mountains) ProceduralGridGenerator
	SetRiver(river *grid.River) ProceduralGridGenerator
	SetCanyon(canyon *proceduralentities.Canyon) ProceduralGridGenerator
	Generate() ([][]taleslabentities.Element, error)
}
