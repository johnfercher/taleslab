package main

import (
	"log"

	"github.com/johnfercher/talescoder/pkg/encoder"
	"github.com/johnfercher/taleslab/pkg/file"
	"github.com/johnfercher/taleslab/pkg/rand"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabmappers"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabrepositories"
)

func main() {
	propRepository, _ := taleslabrepositories.NewPropRepository()

	encoder := encoder.NewEncoder()

	slab := &taleslabentities.Slab{}

	constructor := propRepository.GetProp("ground_nature_small")

	xSize := 50
	ySize := 50
	zSize := 3

	for i := xSize; i > 0; i-- {
		for j := ySize; j > 0; j-- {
			for k := zSize; k > 0; k-- {
				if rand.Int()%2 == 0 {
					asset := &taleslabentities.Asset{
						ID: constructor.Parts[0].ID,
						Coordinates: &taleslabentities.Vector3d{
							X: i - 1,
							Y: j - 1,
							Z: k - 1,
						},
						Rotation: (j - 1) / 41,
					}

					slab.Assets = append(slab.Assets, asset)
				}
			}
		}
	}

	taleSpireSlab := taleslabmappers.TaleSpireSlabFromSlab(slab)

	base64, err := encoder.Encode(taleSpireSlab)
	if err != nil {
		log.Fatal(err)
	}

	err = file.SaveCodes([][]string{{base64}}, "cmd/others/randomcube/data.txt")
	if err != nil {
		log.Fatal(err)
	}
}
