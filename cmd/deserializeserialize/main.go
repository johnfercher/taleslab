package main

import (
	"fmt"
	"github.com/johnfercher/taleslab/internal/helper/bytecompressor"
	talespirecoder2 "github.com/johnfercher/taleslab/internal/helper/talespireadapter/talespirecoder"
	"log"
)

func main() {
	original := "H4sIAAAAAAAACjv369xFJgZGBgYGleWfZa6uaHSdselNJM93fUuQGAIAAKjgjvgoAAAA"

	slabCompressor := bytecompressor.New()
	decoder := talespirecoder2.NewDecoder(slabCompressor)
	encoder := talespirecoder2.NewEncoder(slabCompressor)

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
