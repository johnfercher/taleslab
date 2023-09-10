package taleslabservices

import (
	"github.com/johnfercher/taleslab/pkg/grid"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdto"
)

type MatrixGenerator interface {
	SetGround(ground *taleslabdto.GroundDtoRequest) MatrixGenerator
	SetMountains(mountains *taleslabdto.MountainsDtoRequest) MatrixGenerator
	SetRiver(river *grid.River) MatrixGenerator
	SetCanyon(canyon *taleslabdto.CanyonDtoRequest) MatrixGenerator
	Generate() ([][]taleslabentities.Element, error)
}
