package mocks

import (
	"github.com/johnfercher/taleslab/pkg/taleslab/domain/consts"
	"github.com/johnfercher/taleslab/pkg/taleslab/domain/entities"
	"github.com/johnfercher/taleslab/pkg/talespire/talespirecontracts"
)

func CreateTaleSpireSlab() *talespirecontracts.Slab {
	return &talespirecontracts.Slab{
		MagicBytes:  consts.MagicBytes,
		Version:     2,
		AssetsCount: 1,
		Assets: []*talespirecontracts.Asset{
			{
				Id:           []byte{0x0, 0x0, 0x2c, 0x3e, 0x77, 0x5d, 0x2f, 0xca, 0x44, 0x4c, 0x88, 0xd0, 0xf9, 0xb4, 0xa7, 0xaf, 0xaf, 0x5a},
				LayoutsCount: 1,
				Layouts: []*talespirecontracts.Bounds{
					{
						Coordinates: &talespirecontracts.Vector3d{
							X: 0,
							Y: 0,
							Z: 0,
						},
						Rotation: 0,
					},
				},
			},
		},
	}
}

func CreateBase64Slab() string {
	return "H4sIAAAAAAAE/wAoANf/zvrO0QIAAQAAACw+d10vykRMiND5tKevr1oBAAAAAAAAAAAAAAAAAAEAAP//7YI+iCgAAAA="
}

func CreateSlab() *entities.Slab {
	slab := entities.NewSlab()

	asset := &entities.Asset{
		Id: []byte{0x0, 0x0, 0x2c, 0x3e, 0x77, 0x5d, 0x2f, 0xca, 0x44, 0x4c, 0x88, 0xd0, 0xf9, 0xb4, 0xa7, 0xaf, 0xaf, 0x5a},
		Layouts: []*entities.Bounds{
			{
				Coordinates: &entities.Vector3d{
					X: 0,
					Y: 0,
					Z: 0,
				},
				Rotation: 0,
			},
		},
	}

	slab.AddAsset(asset)

	return slab
}
