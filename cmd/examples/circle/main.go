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
	"math"
)

func main() {
	loader, err := assetloader.NewAssetLoader()
	if err != nil {
		log.Fatal(err.Error())
	}

	compressor := bytecompressor.New()
	encoder := talespirecoder.NewEncoder(compressor)

	slab := entities.NewSlab()

	asset := loader.GetConstructor("ground_nature_small")
	slab.AddAsset(&entities.Asset{
		Id: asset.Id,
	})

	radius := 5

	for i := 0.0; i < 2.0*3.14; i += 0.2 {
		cos := math.Cos(i)
		sin := math.Sin(i)

		xRounded := fix(float64(radius)*cos, 1)
		yRounded := fix(float64(radius)*sin, 1)

		xPositiveTranslated := radius + xRounded
		yPositiveTranslated := radius + yRounded

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

func fix(value float64, fixValue int) int {
	division := value / float64(fixValue)

	divisionRounded := math.Round(division)
	top := int(divisionRounded * float64(fixValue))

	return top
}
*/
