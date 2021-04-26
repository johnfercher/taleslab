package entities

// swagger:model
type River struct {
	// Defines if there will be a river on the map
	// required: true
	// example: true
	HasRiver bool `json:"has_river"`
}
