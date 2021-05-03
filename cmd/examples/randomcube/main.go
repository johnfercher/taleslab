package main

/*
import (
	"fmt"
	"github.com/johnfercher/taleslab/internal/bytecompressor"
	"github.com/johnfercher/taleslab/pkg/assetloader"
	"github.com/johnfercher/taleslab/pkg/mappers"
	"github.com/johnfercher/taleslab/pkg/taleslab/domain/entities"
	"github.com/johnfercher/taleslab/pkg/talespire/talespirecoder"
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

	slab := entities.NewSlab()

	constructor := loader.GetConstructor("ground_nature_small")

	slab.AddAsset(&entities.Asset{
		Id: constructor.Id,
	})

	xSize := 50
	ySize := 50
	zSize := 3

	for i := xSize; i > 0; i-- {
		for j := ySize; j > 0; j-- {
			for k := zSize; k > 0; k-- {
				if rand.Int()%2 == 0 {
					layout := &entities.Bounds{
						Coordinates: &entities.Vector3d{
							X: i - 1,
							Y: j - 1,
							Z: k - 1,
						},
						Rotation: (j - 1) / 41,
					}

					slab.AddLayoutToAsset(constructor.Id, layout)
				}
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
*/
