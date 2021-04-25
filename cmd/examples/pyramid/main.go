package main

import (
	"fmt"
	"github.com/johnfercher/taleslab/pkg/assetloaderv2"
	slab2 "github.com/johnfercher/taleslab/pkg/slab"
	"github.com/johnfercher/taleslab/pkg/slab/slabv2"
	"github.com/johnfercher/taleslab/pkg/slabdecoder"
	"log"
)

func main() {
	loader := assetloaderv2.NewAssetLoaderV2()

	constructors, err := loader.GetConstructors()
	if err != nil {
		log.Fatalln(err)
	}

	builder := slabdecoder.NewSlabEncoderBuilder()
	encoder := builder.Build()

	slab := &slabv2.Slab{
		MagicBytes:  slab2.MagicBytes,
		Version:     2,
		AssetsCount: 1,
		Assets: []*slabv2.Asset{
			{
				Id: constructors["nature"].Id,
			},
		},
	}

	xSize := 20
	ySize := 20
	zSize := 20

	for k := zSize; k > 0; k-- {
		for i := xSize - k; i > k; i-- {
			for j := ySize - k; j > k; j-- {
				layout := &slabv2.Bounds{
					Coordinates: &slabv2.Vector3d{
						X: uint16(slabv2.GainX * i),
						Y: uint16(slabv2.GainY * j),
						Z: uint16(slabv2.GainZ * k),
					},
					Rotation: 0,
				}

				slab.Assets[0].Layouts = append(slab.Assets[0].Layouts, layout)
				slab.Assets[0].LayoutsCount++
			}
		}
	}

	base64, err := encoder.Encode(&slab2.Aggregator{
		SlabV2: slab,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(base64)
}
