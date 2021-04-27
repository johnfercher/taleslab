package slab

import "github.com/johnfercher/taleslab/pkg/assetloader"

type Asset struct {
	Id           []byte    `json:"id"`
	LayoutsCount int16     `json:"layouts_count"`
	Layouts      []*Bounds `json:"layouts"`
	Dimensions   *assetloader.Dimensions
}
