package entities

type Biome string

const (
	ForestBiome Biome = "forest"
	DesertBiome Biome = "desert"
)

type Map struct {
	Biome     Biome      `json:"biome"`
	Ground    *Ground    `json:"ground,omitempty"`
	Mountains *Mountains `json:"mountains,omitempty"`
	River     *River     `json:"river,omitempty"`
	Props     *Props     `json:"props,omitempty"`
}
