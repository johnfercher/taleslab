package main

import (
	"encoding/json"
	"fmt"
	"github.com/johnfercher/taleslab/internal/helper/bytecompressor"
	talespirecoder2 "github.com/johnfercher/taleslab/internal/helper/talespireadapter/talespirecoder"
	"log"
)

func main() {
	original := "H4sIAAAAAAAACjv369xFJgZGBgaGrr5pd7ZJO3h2lzuYFp9PmAISQwAAcYRDSSgAAAA="

	byteCompressor := bytecompressor.New()
	decoder := talespirecoder2.NewDecoder(byteCompressor)
	encoder := talespirecoder2.NewEncoder(byteCompressor)

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
