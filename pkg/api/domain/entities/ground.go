package entities

type Ground struct {
	Width             int     `json:"width"`
	Length            int     `json:"length"`
	TerrainComplexity float64 `json:"terrain_complexity"`
	ForceBaseLand     bool    `json:"force_base_land"`
}
