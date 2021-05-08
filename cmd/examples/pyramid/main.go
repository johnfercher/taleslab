package main

import (
	"fmt"
	"github.com/johnfercher/taleslab/internal/bytecompressor"
	"github.com/johnfercher/taleslab/internal/talespireadapter/talespirecoder"
	"github.com/johnfercher/taleslab/pkg/assetloader"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabmappers"
	"log"
)

func main() {
	loader, err := assetloader.NewAssetLoader()
	if err != nil {
		log.Fatal(err.Error())
	}

	compressor := bytecompressor.New()
	encoder := talespirecoder.NewEncoder(compressor)

	slab := taleslabentities.NewSlab()

	asset := loader.GetConstructor("ground_nature_small")

	slab.AddAsset(&taleslabentities.Asset{
		Id: asset.AssertParts[0].Id,
	})

	xSize := 20
	ySize := 20
	zSize := 20

	for k := zSize; k > 0; k-- {
		for i := xSize - k; i > k; i-- {
			for j := ySize - k; j > k; j-- {
				layout := &taleslabentities.Bounds{
					Coordinates: &taleslabentities.Vector3d{
						X: i,
						Y: j,
						Z: k,
					},
					Rotation: 0,
				}

				slab.AddLayoutToAsset(asset.AssertParts[0].Id, layout)
			}
		}
	}

	taleSpireSlab := taleslabmappers.TaleSpireSlabFromEntity(slab)

	base64, err := encoder.Encode(taleSpireSlab)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(base64)
}
