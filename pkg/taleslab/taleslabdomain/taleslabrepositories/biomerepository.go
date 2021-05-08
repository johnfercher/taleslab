package taleslabrepositories

import (
	"github.com/johnfercher/taleslab/pkg/assetloader"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts"
)

type BiomeRepository interface {
	SetBiome(taleslabconsts.BiomeType)
	GetBiome() taleslabconsts.BiomeType
	GetConstructorKeys() map[taleslabconsts.ElementType][]string
	GetConstructorAssets(elementType taleslabconsts.ElementType) []string
	GetConstructor(id string) *assetloader.AssetInfo
	GetPropKeys() map[taleslabconsts.ElementType][]string
	GetPropAssets(elementType taleslabconsts.ElementType) []string
	GetProp(id string) *assetloader.AssetInfo
	GetStoneWall() string
}
