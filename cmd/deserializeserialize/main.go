package main

import (
	"fmt"
	"github.com/johnfercher/taleslab/internal/bytecompressor"
	"github.com/johnfercher/taleslab/pkg/talespire/talespirecoder"
	"log"
)

func main() {
	original := "H4sIAAAAAAAACz2RoWtCURTGz32IMyg4gygIhgWLYazZDGIxrMkLxhsFs+2lFXn4kgzExcEQs0EFm8ciGBYn/gPCcGEgIt7j+bDcD873O+c7h7s5bbYexYio9d6e9MJjbfx9WO+a4fOLq6V/vpbTwaXe/xvVc6/pjnG1x1n/96G0b8y754/U/rMgNc5W/qdPq8Yi/zZcmZ2RWtk91WJgmFQtlKDiU151mVENkqrVOHzSfuFEhRMV7jY3Dp90vnAWnAVnwVlwBI7AETgCR3eO76p1BsfoY+Qx8hk5jFzGHoz9Gfcw9mbcwbiLsVdkNDcymhsZzY2MzlPfwrfwLXwi39N+39N+39N+pwn9lyAm27v8rNx6BRp5flQEAgAA"

	slabCompressor := bytecompressor.New()
	decoder := talespirecoder.NewDecoder(slabCompressor)
	encoder := talespirecoder.NewEncoder(slabCompressor)

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
