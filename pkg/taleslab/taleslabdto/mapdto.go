package taleslabdto

import validation "github.com/go-ozzo/ozzo-validation"
import "github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts"

// MapDtoResponse response model
// swagger:model
type MapDtoResponse struct {
	// Version of the TaleSpire Slab
	SlabVersion string `json:"slab_version"`
	// Size of the base64 string
	Size string `json:"size"`
	// Code to insert in the game
	Code string `json:"code"`
}

// MapDtoRequest request model
// swagger:model
type MapDtoRequest struct {
	// Biome type (subtropical_forest, temperate_forest, dead_forest, desert, tundra)
	// required: true
	// example: temperate_forest
	Biome taleslabconsts.BiomeType `json:"biome_type,omitempty"`
	// SecondaryBiome type (subtropical_forest, temperate_forest, dead_forest, desert, tundra)
	// required: false
	// example: tundra
	SecondaryBiome taleslabconsts.BiomeType `json:"secondary_biome_type,omitempty"`
	// required: true
	Ground *GroundDtoRequest `json:"ground,omitempty"`
	// required: false
	Mountains *MountainsDtoRequest `json:"mountains,omitempty"`
	// required: false
	River *RiverDtoRequest `json:"river,omitempty"`
	// required: false
	Canyon *CanyonDtoRequest `json:"canyon,omitempty"`
	// required: true
	Props *PropsDtoRequest `json:"props,omitempty"`
}

func (self MapDtoRequest) Validate() error {
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

// Response from API
// swagger:response swaggMapRes
// nolint:deadcode,unused
type swaggMapRes struct {
	// in: body
	Map MapDtoResponse
}

// Request from API
// swagger:parameters MapDtoRequest
// nolint:deadcode,unused
type swaggMapReq struct {
	//in: body
	Map MapDtoRequest
}
