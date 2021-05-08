package taleslabservices

import (
	"github.com/johnfercher/taleslab/internal/api/apierror"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
)

type MapBuilder interface {
	SetBiome(biome taleslabconsts.BiomeType) MapBuilder
	SetSecondaryBiome(biomeType taleslabconsts.BiomeType) MapBuilder
	SetGround(ground *taleslabentities.Ground) MapBuilder
	SetMountains(mountains *taleslabentities.Mountains) MapBuilder
	SetRiver(river *taleslabentities.River) MapBuilder
	SetCanyon(canyon *taleslabentities.Canyon) MapBuilder
	SetProps(props *taleslabentities.Props) MapBuilder
	Build() (string, apierror.ApiError)
}
