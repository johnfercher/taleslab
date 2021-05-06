package main

import (
	"fmt"
	"github.com/johnfercher/taleslab/internal/bytecompressor"
	"github.com/johnfercher/taleslab/pkg/mappers"
	"github.com/johnfercher/taleslab/pkg/talespire/talespirecoder"
	"log"
)

func main() {
	original := "H4sIAAAAAAAACzv369xFJgZGBgaG3ckJ5/Ni2B03v5zDE3U8PIOdAQQmAKUagPgEA4QGkjwQ2oGtgfEEWE0DYwqUZoDTAFUoNPRYAAAA"

	slabCompressor := bytecompressor.New()
	decoder := talespirecoder.NewDecoder(slabCompressor)
	encoder := talespirecoder.NewEncoder(slabCompressor)

	taleSlab, err := decoder.Decode(original)
	if err != nil {
		log.Fatal(err)
	}

	slab := mappers.EntitySlabFromTaleSpire(taleSlab)
	fmt.Println(slab)

	slabBase64, err := encoder.Encode(taleSlab)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(slabBase64)
}
