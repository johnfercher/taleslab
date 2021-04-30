package entities

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
	MinHeight uint16 `json:"min_height"`
	// Forces all 0 height tiles to have ground tiles
	// required: true
	// example: false
	ForceBaseLand bool `json:"force_base_land"`
}
