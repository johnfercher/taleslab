package entities

// BiomeType request model
// example: temperate_forest
type BiomeType string

const (
	TropicalForestBiomeType  BiomeType = "tropical_forest"
	TemperateForestBiomeType BiomeType = "temperate_forest"
	DesertBiomeType          BiomeType = "desert"
	TundraBiomeType          BiomeType = "tundra"
	BeachBiomeType           BiomeType = "beach"
)

// Map request model
// swagger:model
type Map struct {
	// Base Biome type (desert, tropical_forest, temperate_forest, tundra)
	// required: false
	Biome BiomeType `json:"biome_type,omitempty"`
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

// swagger:parameters Map
type swaggMapReq struct {
	//in: body
	Map Map
}
