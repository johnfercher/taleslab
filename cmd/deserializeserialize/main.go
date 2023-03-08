package main

import (
	"fmt"
	"github.com/johnfercher/taleslab/internal/bytecompressor"
	"github.com/johnfercher/taleslab/internal/talespireadapter/talespirecoder"
	"log"
)

func main() {
	original := "H4sIAAAAAAAACjv369xFJgZGBgYGleWfZa6uaHSdselNJM93fUuQGAIAAKjgjvgoAAAA"

	slabCompressor := bytecompressor.New()
	decoder := talespirecoder.NewDecoder(slabCompressor)
	encoder := talespirecoder.NewEncoder(slabCompressor)

	slab, err := decoder.Decode(original)
	if err != nil {
		log.Fatal(err)
	}

	slabBase64, err := encoder.Encode(slab)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(slabBase64)
}
