package main

import (
	"github.com/johnfercher/taleslab/internal/helper/bytecompressor"
	"github.com/johnfercher/taleslab/internal/helper/file"
	"github.com/johnfercher/taleslab/internal/helper/talespireadapter/talespirecoder"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabmappers"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabrepositories"
	"log"
)

func main() {
	propRepository := taleslabrepositories.NewPropRepository()

	compressor := bytecompressor.New()
	encoder := talespirecoder.NewEncoder(compressor)

	assets := taleslabentities.Assets{}

	asset := propRepository.GetProp("ground_nature_small")

	xSize := 20
	ySize := 20
	zSize := 20

	for k := zSize; k > 0; k-- {
		for i := xSize - k; i > k; i-- {
			for j := ySize - k; j > k; j-- {
				asset := &taleslabentities.Asset{
					Id: asset.Parts[0].Id,
					Coordinates: &taleslabentities.Vector3d{
						X: i,
						Y: j,
						Z: k,
					},
					Rotation: 0,
				}

				assets = append(assets, asset)
			}
		}
	}

	taleSpireSlab := taleslabmappers.TaleSpireSlabFromAssets(assets)

	base64, err := encoder.Encode(taleSpireSlab)
	if err != nil {
		log.Fatal(err)
	}

	err = file.SaveCodes([]string{base64}, "docs/codes/pyramid.txt")
	if err != nil {
		log.Fatal(err)
	}
}
