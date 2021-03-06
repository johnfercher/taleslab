package taleslabservices

import (
	"github.com/johnfercher/taleslab/internal/api/apierror"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdto"
)

type MapBuilder interface {
	SetBiome(biome taleslabconsts.BiomeType) MapBuilder
	SetSecondaryBiome(biomeType taleslabconsts.BiomeType) MapBuilder
	SetGround(ground *taleslabdto.GroundDtoRequest) MapBuilder
	SetMountains(mountains *taleslabdto.MountainsDtoRequest) MapBuilder
	SetRiver(river *taleslabdto.RiverDtoRequest) MapBuilder
	SetCanyon(canyon *taleslabdto.CanyonDtoRequest) MapBuilder
	SetProps(props *taleslabdto.PropsDtoRequest) MapBuilder
	Build() (string, apierror.ApiError)
}
