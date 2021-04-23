package main

import (
	"fmt"
	"github.com/johnfercher/taleslab/internal/slabdecoder"
	"github.com/johnfercher/taleslab/internal/slabencoder"
	"log"
)

func main() {
	original := "H4sIAAAAAAAACzv369xFJgZGBgaGI7lip98EZ7jvOWxf8871uj9IDAIOsDIwAABjbcV7KAAAAA=="

	slab, err := slabdecoder.Decode(original)
	if err != nil {
		log.Fatal(err)
	}

	slabBase64, err := slabencoder.Encode(slab)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(slabBase64)
}
