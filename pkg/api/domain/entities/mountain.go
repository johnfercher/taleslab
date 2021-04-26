package entities

type Mountain struct {
	MinX           int `json:"min_x"`
	RandX          int `json:"rand_x"`
	RandY          int `json:"rand_y"`
	MinY           int `json:"min_y"`
	MinComplexity  int `json:"min_complexity"`
	RandComplexity int `json:"rand_complexity"`
	MinHeight      int `json:"min_height"`
	RandHeight     int `json:"rand_height"`
}
