package taleslabservices

import (
	"github.com/johnfercher/taleslab/pkg/grid"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdto"
)

type ProceduralGridGenerator interface {
	SetGround(ground *taleslabdto.GroundDtoRequest) ProceduralGridGenerator
	SetMountains(mountains *taleslabdto.MountainsDtoRequest) ProceduralGridGenerator
	SetRiver(river *grid.River) ProceduralGridGenerator
	SetCanyon(canyon *taleslabdto.CanyonDtoRequest) ProceduralGridGenerator
	Generate() ([][]taleslabentities.Element, error)
}
