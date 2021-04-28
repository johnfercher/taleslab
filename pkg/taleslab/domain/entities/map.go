package entities

// Biome request model
// example: forest
type Biome string

const (
	ForestBiome Biome = "forest"
	DesertBiome Biome = "desert"
	TundraBiome Biome = "tundra"
)

// Map request model
// swagger:model
type Map struct {

	// Base Biome type (Desert, Forest, Tundra)
	// required: false
	Biome Biome `json:"biome,omitempty"`
	// required: true
	Ground *Ground `json:"ground,omitempty"`
	// required: true
	Mountains *Mountains `json:"mountains,omitempty"`
	// required: true
	River *River `json:"river,omitempty"`
	// required: true
	Props *Props `json:"props,omitempty"`
}

// swagger:parameters Map
type swaggMapReq struct {
	//in: body
	Map Map
}
