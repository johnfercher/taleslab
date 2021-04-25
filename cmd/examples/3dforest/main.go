package main

import (
	"fmt"
	"github.com/johnfercher/taleslab/pkg/assetloader"
	"github.com/johnfercher/taleslab/pkg/slab"
	"github.com/johnfercher/taleslab/pkg/slabcompressor"
	"github.com/johnfercher/taleslab/pkg/slabdecoder"
	"log"
	"math/rand"
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

	gridHeights := generateGridHeights(20, 20)
	gridStones := generateGridStones(20, 20)
	gridTrees := generateGridTrees(20, 20, gridStones)

	appendGroundToSlab(constructors, slabGenerated, gridHeights)
	appendStonesToSlab(ornaments, slabGenerated, gridHeights, gridStones)
	appendTreesToSlab(ornaments, slabGenerated, gridHeights, gridTrees)

	base64, err := encoder.Encode(slabGenerated)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(base64)
}

func appendStonesToSlab(ornaments map[string]assetloader.AssetInfo, generatedSlab *slab.Slab, gridHeights [][]int, gridStones [][]bool) {
	generatedSlab.AssetsCount++
	generatedSlab.Assets = append(generatedSlab.Assets,
		&slab.Asset{
			Id: ornaments["big_stone"].Id,
		})

	for i, array := range gridHeights {
		for j, element := range array {
			if gridStones[i][j] {
				addLayout(generatedSlab.Assets[2], i, j, element)
			}
		}
	}
}

func appendTreesToSlab(ornaments map[string]assetloader.AssetInfo, generatedSlab *slab.Slab, gridHeights [][]int, gridTrees [][]bool) {
	generatedSlab.AssetsCount++
	generatedSlab.Assets = append(generatedSlab.Assets,
		&slab.Asset{
			Id: ornaments["pine_tree"].Id,
		})

	for i, array := range gridHeights {
		for j, element := range array {
			if gridTrees[i][j] {
				addLayout(generatedSlab.Assets[3], i, j, element+1)
			}
		}
	}
}

func appendGroundToSlab(constructors map[string]assetloader.AssetInfo, generatedSlab *slab.Slab, gridHeights [][]int) {
	generatedSlab.AssetsCount++
	generatedSlab.Assets = append(generatedSlab.Assets,
		&slab.Asset{
			Id: constructors["nature"].Id,
		})

	generatedSlab.AssetsCount++
	generatedSlab.Assets = append(generatedSlab.Assets,
		&slab.Asset{
			Id: constructors["nature_with_stones"].Id,
		})

	for i, array := range gridHeights {
		for j, element := range array {
			useNormalGround := rand.Int()%2 == 0

			if useNormalGround {
				addLayout(generatedSlab.Assets[0], i, j, element)
			} else {
				addLayout(generatedSlab.Assets[1], i, j, element)
			}

			addLayout(generatedSlab.Assets[0], i, j, element-1)
			addLayout(generatedSlab.Assets[0], i, j, element-2)
		}
	}
}

func generateGridHeights(x, y int) [][]int {
	base := 7.0
	mainValue := 3
	groundHeight := [][]int{}

	for i := 0; i < x; i++ {
		array := []int{}
		for j := 0; j < y; j++ {
			array = append(array, mainValue)
		}
		groundHeight = append(groundHeight, array)
	}

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			lastXHeight := base
			lastYHeight := base

			if i > 1 {
				lastXHeight = float64(groundHeight[i-1][j]+groundHeight[i-2][j]) / 2.0
			}

			if j > 1 {
				lastYHeight = float64(groundHeight[i][j-1]+groundHeight[i][j-2]) / 2.0
			}

			avgHeight := (lastXHeight + lastYHeight) / 2.0

			keepAvgHeight := rand.Int()%2 == 0
			if keepAvgHeight {
				groundHeight[i][j] = int(avgHeight)
				continue
			}

			increaseHeight := rand.Int()%3 != 0
			if increaseHeight {
				groundHeight[i][j] = int(avgHeight) + 1
				continue
			}

			if int(avgHeight)-1 > 3 {
				groundHeight[i][j] = int(avgHeight) - 1
			}

			continue

		}
	}

	fmt.Printf("\n**** Grid Heights ****\n")
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			fmt.Printf("%d\t", groundHeight[i][j])
		}
		fmt.Println()
	}

	return groundHeight
}

func generateGridStones(x, y int) [][]bool {
	defaultValue := false
	groundStones := [][]bool{}

	for i := 0; i < x; i++ {
		array := []bool{}
		for j := 0; j < y; j++ {
			array = append(array, defaultValue)
		}
		groundStones = append(groundStones, array)
	}

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if i == 0 || i == x-1 || j == 0 || j == y-1 {
				continue
			}

			if i > 1 && (groundStones[i-1][j] || groundStones[i-2][j]) {
				continue
			}

			if j > 1 && (groundStones[i][j-1] || groundStones[i][j-2]) {
				continue
			}

			groundStones[i][j] = rand.Int()%41 == 0
		}
	}

	fmt.Printf("\n**** Grid Stones ****\n")
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			fmt.Printf("%v\t", groundStones[i][j])
		}
		fmt.Println()
	}

	return groundStones
}

func generateGridTrees(x, y int, gridStones [][]bool) [][]bool {
	defaultValue := false
	groundStones := [][]bool{}

	for i := 0; i < x; i++ {
		array := []bool{}
		for j := 0; j < y; j++ {
			array = append(array, defaultValue)
		}
		groundStones = append(groundStones, array)
	}

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if i == 0 || i == x-1 || j == 0 || j == y-1 {
				continue
			}

			if gridStones[i][j] {
				continue
			}

			if i > 1 && (groundStones[i-1][j] || groundStones[i-2][j]) {
				continue
			}

			if j > 1 && (groundStones[i][j-1] || groundStones[i][j-2]) {
				continue
			}

			groundStones[i][j] = rand.Int()%5 == 0
		}
	}

	fmt.Printf("\n**** Grid Trees ****\n")
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			fmt.Printf("%v\t", groundStones[i][j])
		}
		fmt.Println()
	}

	return groundStones
}

func addLayout(asset *slab.Asset, x, y, z int) {
	layout := &slab.Bounds{
		Coordinates: &slab.Vector3d{
			X: uint16(slab.GainX * x),
			Y: uint16(slab.GainY * y),
			Z: uint16(slab.GainZ * z),
		},
		Rotation: 0,
	}

	asset.Layouts = append(asset.Layouts, layout)
	asset.LayoutsCount++
}
