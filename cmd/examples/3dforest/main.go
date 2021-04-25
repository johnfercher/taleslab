package main

import (
	"fmt"
	"github.com/johnfercher/taleslab/internal/gridhelper"
	"github.com/johnfercher/taleslab/pkg/assetloader"
	"github.com/johnfercher/taleslab/pkg/slab"
	"github.com/johnfercher/taleslab/pkg/slabcompressor"
	"github.com/johnfercher/taleslab/pkg/slabdecoder"
	"log"
)

func main() {
	loader := assetloader.NewAssetLoader()

	constructors, err := loader.GetConstructors()
	if err != nil {
		log.Fatalln(err)
	}

	ornaments, err := loader.GetOrnaments()
	if err != nil {
		log.Fatalln(err)
	}

	compressor := slabcompressor.New()
	encoder := slabdecoder.NewEncoder(compressor)

	slabGenerated := &slab.Slab{
		MagicBytes: slab.MagicBytes,
		Version:    2,
	}

	x := 50
	y := 50

	gridHeights := gridhelper.MountainGenerator(x, y, 2.0, 2.0, 20.0)
	gridStones := gridhelper.GenerateRandomGridPositions(x, y, 83)
	gridTrees := gridhelper.GenerateExclusiveRandomGrid(x, y, 11, gridStones)

	appendGroundToSlab(constructors, slabGenerated, gridHeights)
	appendStonesToSlab(ornaments, slabGenerated, gridHeights, gridStones)
	appendTreesToSlab(ornaments, slabGenerated, gridHeights, gridTrees)

	base64, err := encoder.Encode(slabGenerated)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(base64)
}

func appendStonesToSlab(ornaments map[string]assetloader.AssetInfo, generatedSlab *slab.Slab, gridHeights [][]uint16, gridStones [][]bool) {
	generatedSlab.AssetsCount++
	generatedSlab.Assets = append(generatedSlab.Assets,
		&slab.Asset{
			Id: ornaments["stone_big"].Id,
		})

	for i, array := range gridHeights {
		for j, element := range array {
			if gridStones[i][j] {
				addLayout(generatedSlab.Assets[1], uint16(i), uint16(j), element)
			}
		}
	}
}

func appendTreesToSlab(ornaments map[string]assetloader.AssetInfo, generatedSlab *slab.Slab, gridHeights [][]uint16, gridTrees [][]bool) {
	generatedSlab.AssetsCount++
	generatedSlab.Assets = append(generatedSlab.Assets,
		&slab.Asset{
			Id: ornaments["pine_tree_big"].Id,
		})

	for i, array := range gridHeights {
		for j, element := range array {
			if gridTrees[i][j] {
				addLayout(generatedSlab.Assets[2], uint16(i), uint16(j), element+1)
			}
		}
	}
}

func appendGroundToSlab(constructors map[string]assetloader.AssetInfo, generatedSlab *slab.Slab, gridHeights [][]uint16) {
	generatedSlab.AssetsCount++
	generatedSlab.Assets = append(generatedSlab.Assets,
		&slab.Asset{
			Id: constructors["ground_nature_small"].Id,
		})

	for i, array := range gridHeights {
		for j, element := range array {
			addLayout(generatedSlab.Assets[0], uint16(i), uint16(j), element)
			addLayout(generatedSlab.Assets[0], uint16(i), uint16(j), element-1)
			addLayout(generatedSlab.Assets[0], uint16(i), uint16(j), element-2)
		}
	}
}

func addLayout(asset *slab.Asset, x, y, z uint16) {
	layout := &slab.Bounds{
		Coordinates: &slab.Vector3d{
			X: x,
			Y: y,
			Z: z,
		},
		Rotation: y / 41,
	}

	asset.Layouts = append(asset.Layouts, layout)
	asset.LayoutsCount++
}
