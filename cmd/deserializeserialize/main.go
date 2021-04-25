package main

import (
	"encoding/json"
	"fmt"
	"github.com/johnfercher/taleslab/pkg/slabdecoder"
	"log"
)

func main() {
	slabs := []string{
		//"H4sIAAAAAAAAAzv369xFRgYmBt6LgpbaIsb+81/FWgkcNW9kYmBgeLrn0b/gP6v99uzirSp+4e3JyAADDfYIGoZR5ByQ5TigMtcXb8Ap9yZwBlwMwgapY2AAAFC/RiOgAAAA", // Version 1
		"H4sIAAAAAAAACzv369xFJgZGBgaGmBk5a7t7P7msvvr25L2wXgOQGBQwMzAAAKBJcHsoAAAA", // Version 2
	}

	slabDecoderBuilder := slabdecoder.NewSlabDecoderBuilder()
	decoder := slabDecoderBuilder.Build()

	slabEncoderBuilder := slabdecoder.NewSlabEncoderBuilder()
	encoder := slabEncoderBuilder.Build()

	for _, slab := range slabs {
		slab, err := decoder.Decode(slab)
		if err != nil {
			log.Fatal(err)
		}

		slabBytes, err := json.Marshal(slab)
		if err != nil {
			log.Fatal(err)
		}

		slabString := string(slabBytes)

		fmt.Println(slabString)

		slabBase64, err := encoder.Encode(slab)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(slabBase64)
	}
}
