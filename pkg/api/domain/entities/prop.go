package entities

// swagger:model
type Props struct {
	// Density of trees on the map
	// required: true
	// example: 11
	TreeDensity int `json:"tree_density"`
	// Density of props on the map ( i.e rocks )
	// required: true
	// example: 83
	PropsDensity int `json:"props_density"`
}
