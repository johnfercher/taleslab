package entities

// swagger:model
type Props struct {
	// Density of trees on the map
	// required: true
	// example: 11
	TreeDensity int `json:"tree_density"`
	// Density of stones on the map
	// required: false
	// example: 83
	StoneDensity int `json:"stone_density"`
	// Density of misc on the map
	// required: false
	// example: 83
	MiscDensity int `json:"misc_density"`
}
