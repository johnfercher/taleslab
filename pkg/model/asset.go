package model

type Asset struct {
	Uuid         string    `json:"uuid"`
	LayoutsCount int       `json:"layouts_count"`
	Layouts      []*Bounds `json:"layouts"`
}
