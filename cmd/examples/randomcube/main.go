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
				Id: constructors["nature_1"].Id,
			},
		},
	}

	xSize := 50
	ySize := 50
	zSize := 3

	for i := xSize; i > 0; i-- {
		for j := ySize; j > 0; j-- {
			for k := zSize; k > 0; k-- {
				if rand.Int()%2 == 0 {
					layout := &slabv2.Bounds{
						Coordinates: &slabv2.Vector3d{
							X: uint16((i - 1) * slabv2.GainX),
							Y: uint16(j - 1),
							Z: uint16((k - 1) * slabv2.GainZ),
						},
						Rotation: uint16((j - 1) / 41),
					}

					slab.Assets[0].Layouts = append(slab.Assets[0].Layouts, layout)
					slab.Assets[0].LayoutsCount++
				}
			}
		}
	}

	aggs := &slab2.Aggregator{
		SlabV2: slab,
	}

	//bytesJson, err := json.Marshal(aggs)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//jsonString := string(bytesJson)
	//fmt.Println(jsonString)

	base64, err := encoder.Encode(aggs)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(base64)
}
