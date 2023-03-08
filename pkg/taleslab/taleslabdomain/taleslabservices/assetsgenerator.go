package taleslabservices

import (
	"github.com/johnfercher/taleslab/internal/api/apierror"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdto"
)

type AssetsGenerator interface {
	SetBiome(biome taleslabconsts.BiomeType) AssetsGenerator
	SetSecondaryBiome(biomeType taleslabconsts.BiomeType) AssetsGenerator
	SetProps(props *taleslabdto.PropsDtoRequest) AssetsGenerator
	Generate(world [][]taleslabentities.Element) (taleslabentities.Assets, apierror.ApiError)
}
