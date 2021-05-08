package taleslabservices

import (
	"github.com/johnfercher/taleslab/internal/api/apierror"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabcontracts"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts"
)

type MapBuilder interface {
	SetBiome(biome taleslabconsts.BiomeType) MapBuilder
	SetSecondaryBiome(biomeType taleslabconsts.BiomeType) MapBuilder
	SetGround(ground *taleslabcontracts.Ground) MapBuilder
	SetMountains(mountains *taleslabcontracts.Mountains) MapBuilder
	SetRiver(river *taleslabcontracts.River) MapBuilder
	SetCanyon(canyon *taleslabcontracts.Canyon) MapBuilder
	SetProps(props *taleslabcontracts.Props) MapBuilder
	Build() (string, apierror.ApiError)
}
