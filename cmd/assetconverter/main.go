package main

import (
	"encoding/json"
	"fmt"
	"github.com/johnfercher/taleslab/internal/bytecompressor"
	"github.com/johnfercher/taleslab/pkg/slabconverter"
	"github.com/johnfercher/taleslab/pkg/talespire/talespirecoder"
)

func main() {
	slabToDecode := "H4sIAAAAAAAACzv369xFJgYmBgaGXrt/95qMI13b9h/waHA2lGAEirml6zw7f26/W7fCD3GLj5cUWIFirxh6mAKEGBhOANkMklAaCCYANTTwQPggmgFOAwCSbtN7ZAAAAA=="

	compressor := bytecompressor.New()
	decoder := talespirecoder.NewDecoder(compressor)
	encoder := talespirecoder.NewEncoder(compressor)
	converter := slabconverter.New(decoder, encoder)

	convertedSlab, err := converter.ConvertToSlab(slabToDecode, "Name", "Type")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	slabBytes, err := json.Marshal(convertedSlab)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	jsonString := string(slabBytes)

	fmt.Println(jsonString)

	//fmt.Println(decodedSlab)
	return
}
