package main

import (
	"fmt"
	"github.com/johnfercher/taleslab/pkg/assetloader"
	slab2 "github.com/johnfercher/taleslab/pkg/slab"
	"github.com/johnfercher/taleslab/pkg/slabcompressor"
	"github.com/johnfercher/taleslab/pkg/slabdecoder"
	"log"
	"math"
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

	radius := 5.0

	for i := 0.0; i < 2.0*3.14; i += 0.2 {
		cos := math.Cos(i)
		sin := math.Sin(i)

		xRounded := fix(radius*cos, uint16(1))
		yRounded := fix(radius*sin, uint16(1))

		xPositiveTranslated := uint16(radius) + xRounded
		yPositiveTranslated := uint16(radius) + yRounded

		layout := &slab2.Bounds{
			Coordinates: &slab2.Vector3d{
				X: xPositiveTranslated,
				Y: yPositiveTranslated,
				Z: 0,
			},
			Rotation: 0,
		}

		slab.Assets[0].Layouts = append(slab.Assets[0].Layouts, layout)
		slab.Assets[0].LayoutsCount++
	}

	base64, err := encoder.Encode(slab)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(base64)
}

func fix(value float64, fixValue uint16) uint16 {
	division := value / float64(fixValue)

	divisionRounded := math.Round(division)
	top := uint16(divisionRounded * float64(fixValue))

	return top
}
