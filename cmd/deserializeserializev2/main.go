package main

import (
	"encoding/json"
	"fmt"
	"github.com/johnfercher/taleslab/pkg/slabcompressor"
	"github.com/johnfercher/taleslab/pkg/slabdecoder/slabdecoderv2"
	"log"
)

func main() {
	original := "H4sIAAAAAAAAC03QIQjCUBSF4fsWTAqCacnkkmnJ9CyKYcmkybRk0mQSHpgEg0EEQRDWTCaTTYvpgbA0k8lkM81ieH/YKed+8Vyb24cnSkSqyU39NofOqrfTmZ4GobiY2HPH0LWJcBuHOMA+rmDBX+X8di1PbPEVn/ER7/EaL/AMj/EI93EXt3AT13ENl3DO3g/DXzjFd3zBJ5zgLV7iOZ7gGA9whDXm4aaBfVyWQv6/IWqusAEAAA=="

	slabCompressor := slabcompressor.New()
	decoder := slabdecoderv2.NewDecoderV2(slabCompressor)
	encoder := slabdecoderv2.NewEncoderV2(slabCompressor)

	slab, err := decoder.Decode(original)
	if err != nil {
		log.Fatal(err)
	}

	slabJsonBytes, err := json.Marshal(slab)
	if err != nil {
		log.Fatal(slabJsonBytes)
	}

	fmt.Println(string(slabJsonBytes))

	slabBase64, err := encoder.Encode(slab)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(slabBase64)
}
