package main

import (
	"fmt"
	"github.com/johnfercher/taleslab/internal/bytecompressor"
	"github.com/johnfercher/taleslab/pkg/assetloader"
	"github.com/johnfercher/taleslab/pkg/mappers"
	"github.com/johnfercher/taleslab/pkg/taleslab/domain/entities"
	"github.com/johnfercher/taleslab/pkg/talespire/talespirecoder"
	"log"
)

func main() {
	loader := assetloader.NewAssetLoader()

	constructors, err := loader.GetConstructors()
	if err != nil {
		log.Fatalln(err)
	}

	compressor := bytecompressor.New()
	encoder := talespirecoder.NewEncoder(compressor)

	slab := entities.NewSlab()

	asset := constructors["ground_nature_small"]

	slab.AddAsset(&entities.Asset{
		Id: asset.Id,
	})

	xSize := 20
	ySize := 20
	zSize := 20

	for k := zSize; k > 0; k-- {
		for i := xSize - k; i > k; i-- {
			for j := ySize - k; j > k; j-- {
				layout := &entities.Bounds{
					Coordinates: &entities.Vector3d{
						X: uint16(i),
						Y: uint16(j),
						Z: uint16(k),
					},
					Rotation: 0,
				}

				slab.AddLayoutToAsset(asset.Id, layout)
			}
		}
	}

	taleSpireSlab := mappers.TaleSpireSlabFromEntity(slab)

	base64, err := encoder.Encode(taleSpireSlab)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(base64)
}
