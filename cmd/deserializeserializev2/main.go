package main

import (
	"encoding/json"
	"fmt"
	"github.com/johnfercher/taleslab/pkg/slabcompressor"
	"github.com/johnfercher/taleslab/pkg/slabdecoder/slabdecoderv2"
	"log"
)

func main() {
	original := "H4sIAAAAAAAAC1XRIQjCQBTG8TfD0pKmpaUlDbJkkgtrJosWbSYxmGzCwCQybGIwmQQRk8kym5YDQYtBMJlWtC2J3r/4wn1wv3fvhdOZPufEEpHWvLeZxK9wfU1P90ZcDsRUZFu/TByTUjCpXJORh/t4CQ/wpblOVgzcmlA75u/xA37ENX7Bb/gDf+Ip/sYzXNhfxUO8htd538TbeAfv4n18gA/xET7Gp/gMX8hfKZt+h/484OIe7uNFnI9Qle/5ARwqTauwAQAA"

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
