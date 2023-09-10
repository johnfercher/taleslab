package main

import (
	"github.com/johnfercher/taleslab/pkg/shared/file"
	"log"

	"github.com/johnfercher/talescoder/pkg/encoder"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabmappers"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabrepositories"
)

func main() {
	propRepository, _ := taleslabrepositories.NewPropRepository()

	encoder := encoder.NewEncoder()

	slab := &taleslabentities.Slab{}

	asset := propRepository.GetProp("ground_nature_small")

	xSize := 20
	ySize := 20
	zSize := 20

	for k := zSize; k > 0; k-- {
		for i := xSize - k; i > k; i-- {
			for j := ySize - k; j > k; j-- {
				asset := &taleslabentities.Asset{
					ID: asset.Parts[0].ID,
					Coordinates: &taleslabentities.Vector3d{
						X: i,
						Y: j,
						Z: k,
					},
					Rotation: 0,
				}

				slab.Assets = append(slab.Assets, asset)
			}
		}
	}

	taleSpireSlab := taleslabmappers.TaleSpireSlabFromSlab(slab)

	base64, err := encoder.Encode(taleSpireSlab)
	if err != nil {
		log.Fatal(err)
	}

	err = file.SaveCodes([][]string{{base64}}, "cmd/others/pyramid/data.txt")
	if err != nil {
		log.Fatal(err)
	}
}
