package taleslabrepositories

import (
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts/biometype"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
)

type BiomeRepository interface {
	SetBiome(biometype.BiomeType)
	GetBiome() biometype.BiomeType
	GetConstructorKeys() map[taleslabconsts.ElementType][]string
	GetConstructorAssets(elementType taleslabconsts.ElementType) []string
	GetPropKeys() map[taleslabconsts.ElementType][]string
	GetPropAssets(elementType taleslabconsts.ElementType) []string
	GetProp(id string) *taleslabentities.Prop
	GetStoneWall() string
}
