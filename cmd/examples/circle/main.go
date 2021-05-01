package main

import (
	"fmt"
	"github.com/johnfercher/taleslab/internal/bytecompressor"
	"github.com/johnfercher/taleslab/pkg/assetloader"
	"github.com/johnfercher/taleslab/pkg/mappers"
	"github.com/johnfercher/taleslab/pkg/taleslab/domain/entities"
	"github.com/johnfercher/taleslab/pkg/talespire/talespirecoder"
	"log"
	"math"
)

func main() {
	loader := assetloader.NewAssetLoader()

	compressor := bytecompressor.New()
	encoder := talespirecoder.NewEncoder(compressor)

	slab := entities.NewSlab()

	asset := loader.GetConstructor("ground_nature_small")
	slab.AddAsset(&entities.Asset{
		Id: asset.Id,
	})

	radius := 5.0

	for i := 0.0; i < 2.0*3.14; i += 0.2 {
		cos := math.Cos(i)
		sin := math.Sin(i)

		xRounded := fix(radius*cos, uint16(1))
		yRounded := fix(radius*sin, uint16(1))

		xPositiveTranslated := uint16(radius) + xRounded
		yPositiveTranslated := uint16(radius) + yRounded

		layout := &entities.Bounds{
			Coordinates: &entities.Vector3d{
				X: xPositiveTranslated,
				Y: yPositiveTranslated,
				Z: 0,
			},
			Rotation: 0,
		}

		slab.AddLayoutToAsset(asset.Id, layout)
	}

	taleSpireSlab := mappers.TaleSpireSlabFromEntity(slab)

	base64, err := encoder.Encode(taleSpireSlab)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(base64)
}

func fix(value float64, fixValue uint16) uint16 {
	division := value / float64(fixValue)

	divisionRounded := math.Round(division)
	top := uint16(divisionRounded * float64(fixValue))

	return top
}
