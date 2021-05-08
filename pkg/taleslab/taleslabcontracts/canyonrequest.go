package taleslabcontracts

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// swagger:model
type Canyon struct {
	// Turn on a Canyon
	// required: false
	// example: true
	HasCanyon bool `json:"has_canyon"`
	// Move the Canyon
	// required: false
	// example: 10
	CanyonOffset int `json:"canyon_offset"`
}

func (self Canyon) Validate() error {
	if !self.HasCanyon {
		return nil
	}

	return validation.ValidateStruct(&self,
		validation.Field(&self.CanyonOffset, validation.Required, validation.Min(5)),
	)
}
