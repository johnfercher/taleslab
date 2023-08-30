package mocks

import (
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
)

func CreateAssets() taleslabentities.Assets {
	assets := taleslabentities.Assets{}

	asset := &taleslabentities.Asset{
		Id: []byte{0x0, 0x0, 0x2c, 0x3e, 0x77, 0x5d, 0x2f, 0xca, 0x44, 0x4c, 0x88, 0xd0, 0xf9, 0xb4, 0xa7, 0xaf, 0xaf, 0x5a},
		Coordinates: &taleslabentities.Vector3d{
			X: 0,
			Y: 0,
			Z: 0,
		},
		Rotation: 0,
	}

	assets = append(assets, asset)

	return assets
}
