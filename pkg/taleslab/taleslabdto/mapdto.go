package taleslabdto

import (
	"github.com/johnfercher/taleslab/pkg/grid"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts/biometype"
)

type MapDtoResponse struct {
	SlabVersion string     `json:"slab_version"`
	Size        int        `json:"size"`
	Codes       [][]string `json:"codes"`
}

type MapDtoRequest struct {
	Biome          biometype.BiomeType  `json:"biome_type,omitempty"`
	SecondaryBiome biometype.BiomeType  `json:"secondary_biome_type,omitempty"`
	Ground         *GroundDtoRequest    `json:"ground,omitempty"`
	Mountains      *MountainsDtoRequest `json:"mountains,omitempty"`
	River          *grid.River          `json:"river,omitempty"`
	Canyon         *CanyonDtoRequest    `json:"canyon,omitempty"`
	Props          *PropsDtoRequest     `json:"props,omitempty"`
}
