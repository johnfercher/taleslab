package main

import (
	"encoding/json"
	"fmt"
	"github.com/johnfercher/taleslab/pkg/slabcompressor"
	"github.com/johnfercher/taleslab/pkg/slabdecoder/slabdecoderv2"
	"log"
)

func main() {
	original := "H4sIAAAAAAAAC03SIUwCcRiG8Q/CWc4C6dIlLJIuXcJjk3Tp0jmnzI0L7oIjSWK7jUTSRLpE0kRwpCt/LZLOuZE0kRiBJOmSm98TfMsTfvUtq/KzLjURKQY3H3dXR92XXW5fvse3iehMWv9rlmiDvlYutCbCQ7yHn+E+7uFt/AR3cQdv4se4hQte1dQPWrPXZlttsMG/8TVe4iv8DS/wJb7An/E5nuMz/BGf4hN8jI/wIZ7iCd7HYzzCQ7yHd3Af9/A23sJd3MGbuI1buOCVxvxos7022Gplg3/ha7zEV/grXuBLfIE/4XM8x2f4Az7FJ/gYv8eHeIpzUHONx3iEh/g53sF93MNP8Rbu4g7ewG3ckn/7BcaMYxBAAwAA"

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
