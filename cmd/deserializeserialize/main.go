package main

import (
	"fmt"
	"github.com/johnfercher/taleslab/internal/slabdecoder"
	"github.com/johnfercher/taleslab/internal/slabencoder"
	"log"
)

func main() {
	original := "H4sIAAAAAAAAAzv369xFRgYOBoFFhxl/T5nv0uM+0+6WXaGaGAMDA+9FQUttEWP/+a9irQSOmjcyA8We7nn0L/jPar89u3iril94e/ICxbgmbr/Ktu6y18Rvp/e2N+1+BVL3Zl/41qnPU71mHuHezaHOXssBFNNOdXR3VV7iue+qVf2ul9J1TECx4BOz317dZ+OyJX65sG0ooxQrUOxw9rT/U++u9FkkLrMw/1x2ECMDDDTYI2gYRpFzQJZjQZFrwCN3wAGPmY5Y9DlgcwuaHDYzHbDZh6YPxT4OiB147IP7C4t9DVjtQzMTi30HsNqHJofFvgN47DuA1T6o3x3xhKcjbvvgZmELT0fcboGbicXvcHvwhCe2dHYAjz4GhuuLN+DyA1QOqx8coPqwxANMDmvahenDmpawuYUFRQ5b+mzAY+YBrGYyoMhhTxO4zWRwxBYuDChyqPrQ/I4nPLH5D58f4PHHgMVMIuIWa35nIMIPWOzDF0fYwxqWH/CENVZ3oroFNR5Q9eELF6z5D4//iPEDNrc04HFng8ObwBlQuQVQ9gIHkAwALIMy/wAHAAA="

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
