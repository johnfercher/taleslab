package taleslabdto

type GroundDtoRequest struct {
	Width             int     `json:"width"`
	Length            int     `json:"length"`
	TerrainComplexity float64 `json:"terrain_complexity"`
	MinHeight         int     `json:"min_height"`
	ForceBaseLand     bool    `json:"force_base_land"`
}
