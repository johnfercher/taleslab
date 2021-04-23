package main

import (
	"fmt"
	"github.com/johnfercher/taleslab/internal/slabencoder"
	"github.com/johnfercher/taleslab/pkg/model"
	"log"
)

func main() {
	slab := MockSlab()

	slabBase64, err := slabencoder.Encode(slab)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(slabBase64)
}

func MockSlab() *model.Slab {
	return &model.Slab{
		MagicHex: []string{
			"CE",
			"FA",
			"CE",
			"D1",
		},
		Version:     1,
		AssetsCount: 2,
		Assets: []*model.Asset{
			{
				Uuid:         "8daf6ef1-017e-6b48-be6d-1fa261569cd4",
				LayoutsCount: 4,
				Layouts: []*model.Bounds{
					{
						Center: &model.Vector3{
							X: 6.0,
							Y: 1.25,
							Z: 65.0,
						},
						Extents: &model.Vector3{
							X: 1,
							Y: 1.25,
							Z: 1,
						},
						Rotation: 8,
					},
					{
						Center: &model.Vector3{
							X: 4.0,
							Y: 1.25,
							Z: 65.0,
						},
						Extents: &model.Vector3{
							X: 1,
							Y: 1.25,
							Z: 1,
						},
						Rotation: 4,
					},
					{
						Center: &model.Vector3{
							X: 9.0,
							Y: 1.25,
							Z: 62.0,
						},
						Extents: &model.Vector3{
							X: 1,
							Y: 1.25,
							Z: 1,
						},
						Rotation: 8,
					},
					{
						Center: &model.Vector3{
							X: 7.0,
							Y: 1.25,
							Z: 63.0,
						},
						Extents: &model.Vector3{
							X: 1,
							Y: 1.25,
							Z: 1,
						},
						Rotation: 0,
					},
				},
			},
			{
				Uuid:         "84210e62-8964-1144-86b9-223ac358c64b",
				LayoutsCount: 4,
				Layouts: []*model.Bounds{
					{
						Center: &model.Vector3{
							X: 7.0,
							Y: 3.0,
							Z: 60.0,
						},
						Extents: &model.Vector3{
							X: 1,
							Y: 1,
							Z: 1,
						},
						Rotation: 8,
					},
					{
						Center: &model.Vector3{
							X: 6.0,
							Y: 1.0,
							Z: 59.0,
						},
						Extents: &model.Vector3{
							X: 1,
							Y: 1,
							Z: 1,
						},
						Rotation: 8,
					},
					{
						Center: &model.Vector3{
							X: 4.0,
							Y: 1.0,
							Z: 63.0,
						},
						Extents: &model.Vector3{
							X: 1,
							Y: 1,
							Z: 1,
						},
						Rotation: 8,
					},
					{
						Center: &model.Vector3{
							X: 4.0,
							Y: 1.0,
							Z: 61.0,
						},
						Extents: &model.Vector3{
							X: 1,
							Y: 1,
							Z: 1,
						},
						Rotation: 8,
					},
				},
			},
		},
		Bounds: &model.Bounds{
			Center: &model.Vector3{
				X: 6.5,
				Y: 2.0,
				Z: 62.0,
			},
			Extents: &model.Vector3{
				X: 3.5,
				Y: 2,
				Z: 4,
			},
			Rotation: 0,
		},
	}
}
