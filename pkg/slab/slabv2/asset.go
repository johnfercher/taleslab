package slabv2

type Asset struct {
	Id           []byte    `json:"id"`
	LayoutsCount int16     `json:"layouts_count"`
	Layouts      []*Bounds `json:"layouts"`
}
