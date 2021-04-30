package main

import (
	"encoding/json"
	"fmt"
	"github.com/johnfercher/taleslab/internal/bytecompressor"
	"github.com/johnfercher/taleslab/pkg/talespire/talespirecoder"
	"log"
)

func main() {
	original := "H4sIAAAAAAAACzv369xFJgZGBgaG+W5vH53LXuDbZFV2h2Vuz36QGAIAAHcfOuEoAAAA"

	byteCompressor := bytecompressor.New()
	decoder := talespirecoder.NewDecoder(byteCompressor)
	encoder := talespirecoder.NewEncoder(byteCompressor)

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
