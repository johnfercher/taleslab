package entities

type Forest struct {
	X                 int       `json:"x"`
	Y                 int       `json:"y"`
	TreeDensity       int       `json:"tree_density"`
	OrnamentDensity   int       `json:"ornament_density"`
	TerrainComplexity float64   `json:"terrain_complexity"`
	Mountains         *Mountain `json:"mountains"`
}
