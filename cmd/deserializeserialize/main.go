package main

import (
	"fmt"
	"github.com/johnfercher/taleslab/internal/bytecompressor"
	"github.com/johnfercher/taleslab/pkg/mappers"
	"github.com/johnfercher/taleslab/pkg/talespire/talespirecoder"
	"log"
)

func main() {
	original := "H4sIAAAAAAAACzv369xFJgYmBgaG/wc689svvfDcPrlnfsmSe72MQLHdyQnn82LYHTe/nMMTdTw8AyRmwvCC0YGNAQkAAGMl9FdEAAAA"

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
