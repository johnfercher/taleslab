package taleslabconsts

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
)

// BiomeType request model
//
//swagger:type string
type BiomeType string

const (
	SubTropicalForestBiomeType BiomeType = "subtropical_forest"
	TemperateForestBiomeType   BiomeType = "temperate_forest"
	DeadForestBiomeType        BiomeType = "dead_forest"
	DeadLandBiomeType          BiomeType = "dead_land"
	DesertBiomeType            BiomeType = "desert"
	TundraBiomeType            BiomeType = "tundra"
	BeachBiomeType             BiomeType = "beach"
)

func ValidateBiomeType(value interface{}) error {
	valueBiome, ok := value.(BiomeType)
	if !ok {
		return validation.Errors{
			"invalid_biome": errors.New("biome should be: subtropical_forest, temperate_forest, dead_forest, desert or tundra"),
		}.Filter()
	}

	switch valueBiome {
	case SubTropicalForestBiomeType:
		return nil
	case TemperateForestBiomeType:
		return nil
	case DeadForestBiomeType:
		return nil
	case DesertBiomeType:
		return nil
	case TundraBiomeType:
		return nil
	default:
		return validation.Errors{
			"invalid_biome": errors.New("biome should be: subtropical_forest, temperate_forest, dead_forest, desert or tundra"),
		}.Filter()
	}
}
