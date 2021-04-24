# Circle

**Printscreen**
![version_size](../../docs/images/circle.png)

**Base64**
```bash
H4sIAAAAAAAE/0SPrQ0CQRCFvz1+FnECe+4MFHAFbBAkiCsACqAAcBiSLQCBwKBIqIAeMKjLYamAGsCQmdkLT+yXTWbee9N8mmeGA8bXu/ueLvPD4hxeYTstgdZDDEZqY1zCyANrmA0hbv7cDYA93PrGd89YJq4y+x+d7T0wn44i8VfV+iL5qgqdixMjBeaTo74dRZoHlk/qA9pP5rR3nu4ojOIr91JB6yEGgF8AAAD//ynsjnYgAQAA
```

**[Code](../../cmd/examples/circle/main.go)**
```golang
package main

import (
	"fmt"
	"github.com/johnfercher/taleslab/pkg/assetloaderv2"
	slab2 "github.com/johnfercher/taleslab/pkg/slab"
	"github.com/johnfercher/taleslab/pkg/slab/slabv2"
	"github.com/johnfercher/taleslab/pkg/slabdecoder"
	"log"
	"math"
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

	radius := 5.0

	for i := 0.0; i < 2.0*3.14; i += 0.2 {
		layout := &slabv2.Bounds{
			Coordinates: &slabv2.Vector3d{
				X: 1000 + fix(slabv2.GainX*radius*math.Cos(i), slabv2.GainX),
				Y: 16000 + fix(slabv2.GainY*radius*math.Sin(i), slabv2.GainY),
				Z: 0,
			},
			Rotation: 0,
		}

		slab.Assets[0].Layouts = append(slab.Assets[0].Layouts, layout)
		slab.Assets[0].LayoutsCount++
	}

	base64, err := encoder.Encode(&slab2.Aggregator{
		SlabV2: slab,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(base64)
}

func fix(value float64, fixValue float64) int16 {
	division := value / fixValue

	divisionRounded := math.Round(division)

	return int16(divisionRounded * fixValue)
}
```