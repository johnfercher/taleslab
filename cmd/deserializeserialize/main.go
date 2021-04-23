package main

import (
	"encoding/json"
	"fmt"
	"github.com/johnfercher/taleslab/internal/slabdecoder"
	"github.com/johnfercher/taleslab/internal/slabencoder"
	"log"
)

func main() {
	original := "H4sIAAAAAAAAAzv369xFRgYmBt6LgpbaIsb+81/FWgkcNW9kYmBgeLrn0b/gP6v99uzirSp+4e3JyAADDfYIGoZR5ByQ5TigMtcXb8Ap9yZwBlwMwgapY2AAAFC/RiOgAAAA"

	slab, err := slabdecoder.Decode(original)
	if err != nil {
		log.Fatal(err)
	}

	slabBytes, err := json.Marshal(slab)
	if err != nil {
		log.Fatal(err)
	}

	slabString := string(slabBytes)

	fmt.Println(slabString)

	slabBase64, err := slabencoder.Encode(slab)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(slabBase64)
}
