package biometype

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
)

// BiomeType request model
//
//swagger:type string
type BiomeType string

const (
	SubTropicalForest BiomeType = "subtropical_forest"
	TemperateForest   BiomeType = "temperate_forest"
	DeadForest        BiomeType = "dead_forest"
	Swamp             BiomeType = "swamp"
	Lava              BiomeType = "lava"
	Desert            BiomeType = "desert"
	Tundra            BiomeType = "tundra"
	Beach             BiomeType = "beach"
)

func ValidateBiomeType(value interface{}) error {
	valueBiome, ok := value.(BiomeType)
	if !ok {
		return validation.Errors{
			"invalid_biome": errors.New("biome should be: subtropical_forest, temperate_forest, dead_forest, desert or tundra"),
		}.Filter()
	}

	switch valueBiome {
	case SubTropicalForest:
		return nil
	case TemperateForest:
		return nil
	case DeadForest:
		return nil
	case Desert:
		return nil
	case Tundra:
		return nil
	default:
		return validation.Errors{
			"invalid_biome": errors.New("biome should be: subtropical_forest, temperate_forest, dead_forest, desert or tundra"),
		}.Filter()
	}
}
