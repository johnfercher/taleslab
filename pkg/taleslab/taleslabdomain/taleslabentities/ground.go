package taleslabentities

import validation "github.com/go-ozzo/ozzo-validation"

// swagger:model
type Ground struct {
	// World map width
	// required: true
	// example: 70
	Width int `json:"width"`
	// World map length
	// required: true
	// example: 70
	Length int `json:"length"`
	// Defines how "wavy" the base terrain will be
	// required: true
	// example: 5
	TerrainComplexity float64 `json:"terrain_complexity"`
	// Defines the minimum height
	// false: true
	// example: 5
	MinHeight int `json:"min_height"`
	// Forces all 0 height tiles to have ground tiles
	// required: true
	// example: false
	ForceBaseLand bool `json:"force_base_land"`
}

func (self Ground) Validate() error {
	validate := validation.ValidateStruct(&self,
		validation.Field(&self.Width, validation.Required, validation.Min(10), validation.Max(70)),
		validation.Field(&self.Length, validation.Required, validation.Min(10), validation.Max(70)),
		validation.Field(&self.TerrainComplexity, validation.Required, validation.Max(10.0)),
		validation.Field(&self.MinHeight, validation.Min(0)),
	)

	return validate
}
