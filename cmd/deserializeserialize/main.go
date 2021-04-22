package main

import (
	"fmt"
	"github.com/johnfercher/taleslab/internal/slabdecoder"
	"github.com/johnfercher/taleslab/internal/slabencoder"
	"log"
)

func main() {
	original := "H4sIAAAAAAAAAzv369xFRgYOBoFFhxl/T5nv0uM+0+6WXaGaGAMDA+9FQUttEWP/+a9irQSOmjcyA8We7nn0L/jPar89u3iril94e3IAxbgmbr/Ktu6y18Rvp/e2N+1+xQQUe7MvfOvU56leM49w7+ZQZ69lBIpppzq6uyov8dx31ap+10vpOpC64BOz317dZ+OyJX65sG0ooxQLUOxw9rT/U++u9FkkLrMw/1x2EEgvBDTYI2gI5kCVc0CWY0GRa0CRY0CRO+CAx0xHLGY6YHMLmhw2tzhgcwsHihxWtzhgcwtErgGrfQwoctj83oDVPhZUM7H4HaYe1e8MKHLYzDyA1UxoQDviNhMuhy1cYGZhi1tH3G6Bm4nFf3C/4/EfNrfA7cHmTobrizdguIUBRQ5ruDhgk4PFOzYzYWGNTY4DRQ6bmQyOUDkGPHLY/IDVTBYUd2KNW6xyHATNhLkFe14hNVxQzcSabx1xxxG+sG5weBM4A5omFkDZCxxAMgAxCMPdeAUAAA=="

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
