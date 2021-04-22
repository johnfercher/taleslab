package model

type Asset struct {
	Uuid         string    `json:"uuid"`
	LayoutsCount int16     `json:"layouts_count"`
	Layouts      []*Bounds `json:"layouts"`
}
