package slabv1

type Asset struct {
	Id           []byte    `json:"id,omitempty"`
	LayoutsCount int16     `json:"layouts_count,omitempty"`
	Layouts      []*Bounds `json:"layouts,omitempty"`
}
