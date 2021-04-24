package main

import (
	"encoding/json"
	"fmt"
	"github.com/johnfercher/taleslab/pkg/slabcompressor"
	"github.com/johnfercher/taleslab/pkg/slabdecoder/slabdecoderv2"
	"log"
)

func main() {
	original := "H4sIAAAAAAAACzv369xFJgZGBgYGr8DUtPLzq33W2rhWTfLKuAkSQwAAbgFGQigAAAA="

	slabCompressor := slabcompressor.New()
	decoder := slabdecoderv2.NewDecoderV2(slabCompressor)
	encoder := slabdecoderv2.NewEncoderV2(slabCompressor)

	slab, err := decoder.Decode(original)
	if err != nil {
		log.Fatal(err)
	}

	/*slab.Assets[0].Layouts = nil
	slab.Assets[0].LayoutsCount = 0

	layoutNew := &slabv2.Bounds{
		Coordinates: &slabv2.Vector3d{K
			X: 0,
			Y: 0,
			Z: 0,
		},
		Rotation: 0,
	}

	slab.Assets[0].Layouts = append(slab.Assets[0].Layouts, layoutNew)
	slab.Assets[0].LayoutsCount++

	layoutNew = &slabv2.Bounds{
		Coordinates: &slabv2.Vector3d{
			X: 0,
			Y: 0,
			Z: 200,
		},
		Rotation: 0,
	}

	slab.Assets[0].Layouts = append(slab.Assets[0].Layouts, layoutNew)
	slab.Assets[0].LayoutsCount++

	layoutNew = &slabv2.Bounds{
		Coordinates: &slabv2.Vector3d{
			X: 200,
			Y: 0,
			Z: 0,
		},
		Rotation: 0,
	}

	slab.Assets[0].Layouts = append(slab.Assets[0].Layouts, layoutNew)
	slab.Assets[0].LayoutsCount++

	layoutNew = &slabv2.Bounds{
		Coordinates: &slabv2.Vector3d{
			X: 0,
			Y: 9400,
			Z: 0,
		},
		Rotation: 0,
	}

	slab.Assets[0].Layouts = append(slab.Assets[0].Layouts, layoutNew)
	slab.Assets[0].LayoutsCount++*/

	slabJsonBytes, err := json.Marshal(slab)
	if err != nil {
		log.Fatal(slabJsonBytes)
	}

	slabJsonString := string(slabJsonBytes)
	fmt.Println(slabJsonString)

	slabBase64, err := encoder.Encode(slab)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(slabBase64)
}
