package main

import (
	"encoding/json"
	"fmt"
	"github.com/johnfercher/taleslab/pkg/slabcompressor"
	"github.com/johnfercher/taleslab/pkg/slabdecoder"
	"log"
)

func main() {
	original := "H4sIAAAAAAAACzv369xFJgZGBgaGVROORFw+w+HQ/ll4cnSXvwxIDAEA6LpdlygAAAA="

	slabCompressor := slabcompressor.New()
	decoder := slabdecoder.NewDecoder(slabCompressor)
	encoder := slabdecoder.NewEncoder(slabCompressor)

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
