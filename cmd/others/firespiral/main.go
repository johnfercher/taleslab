package main

import (
	"github.com/johnfercher/talescoder/pkg/encoder"
	"github.com/johnfercher/taleslab/pkg/file"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabmappers"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabrepositories"
	"log"
	"math"
)

func main() {
	propRepository := taleslabrepositories.NewPropRepository()

	encoder := encoder.NewEncoder()

	assets := taleslabentities.Assets{}

	asset := propRepository.GetProp("fire")

	radius := 8

	for i := 0.0; i < 4.0*3.14; i += 0.02 {
		cos := math.Cos(i)
		sin := math.Sin(i)

		xRounded := fix(float64(radius)*cos, 1)
		yRounded := fix(float64(radius)*sin, 1)

		xPositiveTranslated := radius + xRounded
		yPositiveTranslated := radius + yRounded

		asset := &taleslabentities.Asset{
			Id: asset.Parts[0].Id,
			Coordinates: &taleslabentities.Vector3d{
				X: xPositiveTranslated,
				Y: yPositiveTranslated,
				Z: int(i),
			},
			Rotation: 0,
		}

		assets = append(assets, asset)
	}

	taleSpireSlab := taleslabmappers.TaleSpireSlabFromAssets(assets)

	base64, err := encoder.Encode(taleSpireSlab)
	if err != nil {
		log.Fatal(err)
	}

	err = file.SaveCodes([][]string{{base64}}, "cmd/others/firespiral/data.txt")
	if err != nil {
		log.Fatal(err)
	}
}

func fix(value float64, fixValue int) int {
	division := value / float64(fixValue)

	divisionRounded := math.Round(division)
	top := int(divisionRounded * float64(fixValue))

	return top
}
