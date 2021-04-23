package model

type Asset struct {
	Uuid         string    `json:"uuid"`
	Id           []int8    `json:"id"`
	LayoutsCount int16     `json:"layouts_count"`
	Layouts      []*Bounds `json:"layouts"`
}
