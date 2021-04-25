package main

import (
	"fmt"
	"github.com/johnfercher/taleslab/pkg/assetloader"
	slab2 "github.com/johnfercher/taleslab/pkg/slab"
	"github.com/johnfercher/taleslab/pkg/slabcompressor"
	"github.com/johnfercher/taleslab/pkg/slabdecoder"
	"log"
	"math/rand"
)

func main() {
	loader := assetloader.NewAssetLoader()

	constructors, err := loader.GetConstructors()
	if err != nil {
		log.Fatalln(err)
	}

	compressor := slabcompressor.New()
	encoder := slabdecoder.NewEncoder(compressor)

	slab := &slab2.Slab{
		MagicBytes:  slab2.MagicBytes,
		Version:     2,
		AssetsCount: 1,
		Assets: []*slab2.Asset{
			{
				Id: constructors["ground_nature_small"].Id,
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
					layout := &slab2.Bounds{
						Coordinates: &slab2.Vector3d{
							X: uint16(i - 1),
							Y: uint16(j - 1),
							Z: uint16(k - 1),
						},
						Rotation: uint16((j - 1) / 41),
					}

					slab.Assets[0].Layouts = append(slab.Assets[0].Layouts, layout)
					slab.Assets[0].LayoutsCount++
				}
			}
		}
	}

	base64, err := encoder.Encode(slab)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(base64)
}
