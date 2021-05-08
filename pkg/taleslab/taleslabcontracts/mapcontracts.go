package taleslabcontracts

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts"
)

// MapResponse response model
type MapResponse struct {
	// Version of the TaleSpire Slab
	SlabVersion string `json:"slab_version"`
	// Size of the base64 string
	Size string `json:"size"`
	// Code to insert in the game
	Code string `json:"code"`
}

// Response from API
// swagger:response mapRes
// nolint:deadcode,unused
type swaggMapRes struct {
	// in: body
	Map MapResponse
}

// Map request model
// swagger:model
type Map struct {
	// Biome type (subtropical_forest, temperate_forest, dead_forest, desert, tundra)
	// required: true
	// example: temperate_forest
	Biome taleslabconsts.BiomeType `json:"biome_type,omitempty"`
	// SecondaryBiome type (subtropical_forest, temperate_forest, dead_forest, desert, tundra)
	// required: false
	// example: tundra
	SecondaryBiome taleslabconsts.BiomeType `json:"secondary_biome_type,omitempty"`
	// required: true
	Ground *Ground `json:"ground,omitempty"`
	// required: false
	Mountains *Mountains `json:"mountains,omitempty"`
	// required: false
	River *River `json:"river,omitempty"`
	// required: false
	Canyon *Canyon `json:"canyon,omitempty"`
	// required: true
	Props *Props `json:"props,omitempty"`
}

func (self Map) Validate() error {
	err := validation.Errors{
		"map": validation.ValidateStruct(&self,
			validation.Field(&self.Biome, validation.Required, validation.By(taleslabconsts.ValidateBiomeType)),
			validation.Field(&self.SecondaryBiome, validation.By(taleslabconsts.ValidateBiomeType)),
			validation.Field(&self.Ground),
			validation.Field(&self.Mountains),
			validation.Field(&self.Canyon),
			validation.Field(&self.Props),
		),
	}.Filter()

	return err
}

// Request from API
// swagger:parameters Map
// nolint:deadcode,unused
type swaggMapReq struct {
	//in: body
	Map Map
}
