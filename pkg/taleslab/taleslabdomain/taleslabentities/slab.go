package taleslabentities

import (
	"github.com/johnfercher/taleslab/pkg/assetloader"
)

type Asset struct {
	Id          []byte
	Name        string
	Coordinates *Vector3d
	Rotation    int
	Dimensions  *assetloader.Dimensions
	OffsetZ     int
}

type Assets []*Asset

type Vector3d struct {
	X int
	Y int
	Z int
}
