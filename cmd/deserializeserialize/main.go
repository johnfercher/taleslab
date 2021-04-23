package main

import (
	"fmt"
	"github.com/johnfercher/taleslab/internal/slabdecoder"
	"github.com/johnfercher/taleslab/internal/slabencoder"
	"log"
)

func main() {
	original := "H4sIAAAAAAAACzv369xFJgZmBgaGrnxXhdcrr7vMqb6brRTL2s0IFHOV621ImxbtuMUvsu1qaONikNiRXLHTb4Iz3Pcctq9553rdHyQmBiQK+BgYmEAcKBDgPsDEwAAAJUKx/GAAAAA="

	slab, err := slabdecoder.Decode(original)
	if err != nil {
		log.Fatal(err)
	}

	slab.Assets[0].Layouts[0].RotationNew = 704

	slabBase64, err := slabencoder.Encode(slab)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(slabBase64)
}
