package taleslabentities

type Prop struct {
	ID    string  `json:"id"`
	Parts []*Part `json:"asset_parts"`
}

type Part struct {
	ID         []byte      `json:"id"`
	Dimensions *Dimensions `json:"dimensions"`
	OffsetZ    int         `json:"offset_z"`
	Name       string      `json:"name"`
}

type Dimensions struct {
	Width  int `json:"width"`
	Length int `json:"length"`
	Height int `json:"height"`
}
