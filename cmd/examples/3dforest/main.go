package main

import (
	"fmt"
	"github.com/johnfercher/taleslab/internal/gridhelper"
	"github.com/johnfercher/taleslab/pkg/assetloader"
	"github.com/johnfercher/taleslab/pkg/slab"
	"github.com/johnfercher/taleslab/pkg/slabcompressor"
	"github.com/johnfercher/taleslab/pkg/slabdecoder"
	"log"
	"math/rand"
	"time"
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

	worldX := 70
	worldY := 70

	world := generateGround(worldX, worldY)

	gridStones := gridhelper.GenerateRandomGridPositions(worldX, worldY, 83)
	gridTrees := gridhelper.GenerateExclusiveRandomGrid(worldX, worldY, 11, gridStones)

	appendGroundToSlab(constructors, slabGenerated, world)
	appendStonesToSlab(ornaments, slabGenerated, world, gridStones)
	appendTreesToSlab(ornaments, slabGenerated, world, gridTrees)

	base64, err := encoder.Encode(slabGenerated)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(base64)
}

func generateGround(worldX, worldY int) [][]uint16 {
	world := gridhelper.TerrainGenerator(worldX, worldY, 2.0, 2.0, 5.0)

	rand.Seed(time.Now().UnixNano())

	iCount := rand.Intn(6) + 3

	rand.Seed(time.Now().UnixNano())
	jCount := rand.Intn(6) + 3

	for i := 0; i < iCount; i++ {
		for j := 0; j < jCount; j++ {
			rand.Seed(time.Now().UnixNano())
			mountainX := rand.Intn(30) + 15

			rand.Seed(time.Now().UnixNano())
			mountainY := rand.Intn(30) + 15

			rand.Seed(time.Now().UnixNano())
			gain := float64(rand.Intn(10.0) + 10.0)

			mountain := gridhelper.MountainGenerator(mountainX, mountainY, gain)
			world = gridhelper.BuildTerrain(world, mountain)
		}
	}

	world = gridhelper.DigRiver(world)

	return world
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
			if gridTrees[i][j] && gridHeights[i][j] > 0 {
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
			if element == 0 {
				continue
			}

			minValue := element

			if i > 0 && gridHeights[i-1][j] < minValue {
				minValue = gridHeights[i-1][j]
			}

			if i < len(gridHeights)-1 && gridHeights[i+1][j] < minValue {
				minValue = gridHeights[i+1][j]
			}

			if j > 0 && gridHeights[i][j-1] < minValue {
				minValue = gridHeights[i][j-1]
			}

			if j < len(gridHeights[i])-1 && gridHeights[i][j+1] < minValue {
				minValue = gridHeights[i][j+1]
			}

			// Use the minimum neighborhood height to fill empty spaces
			for k := minValue; k <= element; k++ {
				addLayout(generatedSlab.Assets[0], uint16(i), uint16(j), k)
			}
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
