package main

import (
	"fmt"
	"github.com/johnfercher/taleslab/internal/bytecompressor"
	"github.com/johnfercher/taleslab/internal/talespireadapter/talespirecoder"
	"github.com/johnfercher/taleslab/pkg/assetloader"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabmappers"
	"log"
	"math/rand"
)

func main() {
	loader, err := assetloader.NewAssetLoader()
	if err != nil {
		log.Fatal(err.Error())
	}

	compressor := bytecompressor.New()
	encoder := talespirecoder.NewEncoder(compressor)

	assets := taleslabentities.Assets{}

	constructor := loader.GetConstructor("ground_nature_small")

	xSize := 50
	ySize := 50
	zSize := 3

	for i := xSize; i > 0; i-- {
		for j := ySize; j > 0; j-- {
			for k := zSize; k > 0; k-- {
				if rand.Int()%2 == 0 {
					asset := &taleslabentities.Asset{
						Id: constructor.AssertParts[0].Id,
						Coordinates: &taleslabentities.Vector3d{
							X: i - 1,
							Y: j - 1,
							Z: k - 1,
						},
						Rotation: (j - 1) / 41,
					}

					assets = append(assets, asset)
				}
			}
		}
	}

	taleSpireSlab := taleslabmappers.TaleSpireSlabFromAssets(assets)

	base64, err := encoder.Encode(taleSpireSlab)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(base64)
}
