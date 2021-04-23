package main

import (
	"fmt"
	"github.com/johnfercher/taleslab/internal/slabdecoder2"
	"github.com/johnfercher/taleslab/internal/slabencoder2"
	"log"
)

func main() {
	original := "H4sIAAAAAAAACzv369xFJgYWBgaGOUbn1Jf/MPXtWx++7fPm7qOMQLHukv7gXZJ5nu33BRK7JkuUgMRctQ/3/1+6yG3ftSk5PevN3oLEdii+f+SUaOo4I3JS44pNYX9AYqxAfEGDgaUfyAmQc2DxAgkyODCByAAhEA0A/IkpR3wAAAA="

	slab, err := slabdecoder2.Decode(original)
	if err != nil {
		log.Fatal(err)
	}

	slab.Assets[0].Layouts[0].RotationNew = 704

	slabBase64, err := slabencoder2.Encode(slab)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(slabBase64)
}
