package entities

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
)

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

func ValidateBiomeType(value interface{}) error {
	valueBiome, ok := value.(BiomeType)
	if !ok {
		return validation.Errors{
			"invalid_biome": errors.New("biome should be: tropical_forest, temperate_forest, desert or tundra"),
		}.Filter()
	}

	switch valueBiome {
	case TropicalForestBiomeType:
		return nil
	case TemperateForestBiomeType:
		return nil
	case DesertBiomeType:
		return nil
	case TundraBiomeType:
		return nil
	case BeachBiomeType:
		return nil
	default:
		return validation.Errors{
			"invalid_biome": errors.New("biome should be: tropical_forest, temperate_forest, desert or tundra"),
		}.Filter()
	}
}

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

func (self Map) Validate() error {
	err := validation.Errors{
		"map": validation.ValidateStruct(&self,
			validation.Field(&self.Biome, validation.Required, validation.By(ValidateBiomeType)),
			validation.Field(&self.Ground),
			validation.Field(&self.Mountains),
			validation.Field(&self.Canyon),
			validation.Field(&self.Props),
		),
	}.Filter()

	return err
}

// swagger:parameters Map
type swaggMapReq struct {
	//in: body
	Map Map
}
