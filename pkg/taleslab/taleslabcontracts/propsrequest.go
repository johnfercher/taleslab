package taleslabcontracts

import validation "github.com/go-ozzo/ozzo-validation"

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

func (self Props) Validate() error {
	validate := validation.ValidateStruct(&self,
		validation.Field(&self.TreeDensity, validation.Min(0)),
		validation.Field(&self.StoneDensity, validation.Min(0)),
		validation.Field(&self.MiscDensity, validation.Min(0)),
	)

	return validate
}
