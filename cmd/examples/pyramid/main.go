package main

import (
	"fmt"
	"github.com/johnfercher/taleslab/pkg/assetloader"
	slab2 "github.com/johnfercher/taleslab/pkg/slab"
	"github.com/johnfercher/taleslab/pkg/slabcompressor"
	"github.com/johnfercher/taleslab/pkg/slabdecoder"
	"log"
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
				Id: constructors["nature_1"].Id,
			},
		},
	}

	xSize := 20
	ySize := 20
	zSize := 20

	for k := zSize; k > 0; k-- {
		for i := xSize - k; i > k; i-- {
			for j := ySize - k; j > k; j-- {
				layout := &slab2.Bounds{
					Coordinates: &slab2.Vector3d{
						X: uint16(i),
						Y: uint16(j),
						Z: uint16(k),
					},
					Rotation: 0,
				}

				slab.Assets[0].Layouts = append(slab.Assets[0].Layouts, layout)
				slab.Assets[0].LayoutsCount++
			}
		}
	}

	base64, err := encoder.Encode(slab)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(base64)
}
