package entities

// swagger:model
type Canyon struct {
	// Turn on a Canyon
	// required: false
	// example: true
	HasCanyon bool `json:"has_canyon"`
	// Move the Canyon
	// required: false
	// example: 10
	CanyonOffset uint `json:"canyon_offset"`
}
