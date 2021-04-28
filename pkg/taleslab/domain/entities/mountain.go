package entities

// swagger:model
type Mountains struct {
	// Defines the minimum width of the mountains in the map
	// required: true
	// example: 15
	MinX int `json:"min_x"`
	// Defines a maximum value to be added to the minimum width
	// required: true
	// example: 30
	RandX int `json:"rand_x"`
	// Defines the minimum length of the mountains in the map
	// required: true
	// example: 15
	MinY int `json:"min_y"`
	// Defines a maximum value to be added to the minimum length
	// required: true
	// example: 30
	RandY int `json:"rand_y"`
	// Minimum amount of mountains on the map
	// required: true
	// example: 3
	MinComplexity int `json:"min_complexity"`
	// Defines a maximum value to be added to the amount of mountains on the map
	// required: true
	// example: 6
	RandComplexity int `json:"rand_complexity"`
	// Defines the minimum height of the mountains on the map
	// required: true
	// example: 10
	MinHeight int `json:"min_height"`
	// Defines a maximum value to be added to the minimum height of the mountains on the map
	// required: true
	// example: 10
	RandHeight int `json:"rand_height"`
}
