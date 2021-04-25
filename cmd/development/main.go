package main

import (
	"encoding/json"
	"fmt"
	"github.com/johnfercher/taleslab/pkg/slabcompressor"
	"github.com/johnfercher/taleslab/pkg/slabdecoder/slabdecoderv2"
	"log"
)

func main() {
	original := "H4sIAAAAAAAACzv369xFJgZGBgaGmBk5a7t7P7msvvr25L2wXgOQGAIAAPn3NnkoAAAA"

	slabCompressor := slabcompressor.New()
	decoder := slabdecoderv2.NewDecoderV2(slabCompressor)
	encoder := slabdecoderv2.NewEncoderV2(slabCompressor)

	slab, err := decoder.Decode(original)
	if err != nil {
		log.Fatal(err)
	}

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
