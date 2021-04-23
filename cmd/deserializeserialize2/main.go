package main

import (
	"encoding/json"
	"fmt"
	"github.com/johnfercher/taleslab/internal/slabdecoder2"
	"github.com/johnfercher/taleslab/internal/slabencoder2"
	"log"
)

func main() {
	original := "H4sIAAAAAAAACzv369xFJgYmBgaG7pL+4F2SeZ7t9wUSuyZLlDACxXYovn/klGjqOCNyUuOKTWF/QOr6gRIBgg4sQCZDAKsDkzQziOUAlAIAoQYiAEwAAAA="

	slab, err := slabdecoder2.Decode(original)
	if err != nil {
		log.Fatal(err)
	}

	slabJsonBytes, err := json.Marshal(slab)
	if err != nil {
		log.Fatal(slabJsonBytes)
	}

	fmt.Println(string(slabJsonBytes))

	//slab.Assets[0].Layouts[0].Rotation = 704

	slabBase64, err := slabencoder2.Encode(slab)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(slabBase64)
}
