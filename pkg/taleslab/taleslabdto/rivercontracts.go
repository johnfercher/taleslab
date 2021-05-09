package taleslabdto

// RiverDtoRequest response model
// swagger:model
type RiverDtoRequest struct {
	// Defines if there will be a river on the map
	// required: true
	// example: true
	HasRiver bool `json:"has_river"`
}
