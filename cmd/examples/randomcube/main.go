package main

import (
	"fmt"
	"github.com/johnfercher/taleslab/pkg/assetloaderv2"
	slab2 "github.com/johnfercher/taleslab/pkg/slab"
	"github.com/johnfercher/taleslab/pkg/slab/slabv2"
	"github.com/johnfercher/taleslab/pkg/slabdecoder"
	"log"
	"math/rand"
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
				Id: constructors[0].Id,
			},
		},
	}

	xSize := 10
	ySize := 10
	zSize := 10

	for i := xSize; i > 0; i-- {
		for j := ySize; j > 0; j-- {
			for k := zSize; k > 0; k-- {
				if rand.Int()%2 == 0 {
					layout := &slabv2.Bounds{
						Coordinates: &slabv2.Vector3d{
							X: int16(slabv2.GainX * i),
							Y: int16(slabv2.GainY * j),
							Z: int16(slabv2.GainZ * k),
						},
						Rotation: 0,
					}

					slab.Assets[0].Layouts = append(slab.Assets[0].Layouts, layout)
					slab.Assets[0].LayoutsCount++
				}
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
