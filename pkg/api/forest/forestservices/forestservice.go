package forestservices

import (
	"context"
	"fmt"
	"github.com/johnfercher/taleslab/internal/api/apierror"
	"github.com/johnfercher/taleslab/internal/gridhelper"
	"github.com/johnfercher/taleslab/pkg/api/contracts"
	"github.com/johnfercher/taleslab/pkg/api/domain/entities"
	"github.com/johnfercher/taleslab/pkg/assetloader"
	"github.com/johnfercher/taleslab/pkg/slab"
	"github.com/johnfercher/taleslab/pkg/slabdecoder"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type forestService struct {
	loader  assetloader.AssetLoader
	encoder slabdecoder.Encoder
}

func NewForestService(encoder slabdecoder.Encoder) *forestService {
	return &forestService{
		encoder: encoder,
		loader:  assetloader.NewAssetLoader(),
	}
}

func (self *forestService) GenerateForest(ctx context.Context, forest *entities.Forest) (contracts.Slab, apierror.ApiError) {

	constructors, err := self.loader.GetConstructors()
	if err != nil {
		log.Fatalln(err)
	}

	ornaments, err := self.loader.GetOrnaments()
	if err != nil {
		log.Fatalln(err)
	}

	slabGenerated := &slab.Slab{
		MagicBytes: slab.MagicBytes,
		Version:    2,
	}

	world := generateGround(forest.X, forest.Y, forest.TerrainComplexity, forest.Mountains)

	gridStones := gridhelper.GenerateRandomGridPositions(forest.X, forest.Y, forest.OrnamentDensity)
	gridTrees := gridhelper.GenerateExclusiveRandomGrid(forest.X, forest.Y, forest.TreeDensity, gridStones)

	appendGroundToSlab(constructors, slabGenerated, world)
	appendStonesToSlab(ornaments, slabGenerated, world, gridStones)
	appendTreesToSlab(ornaments, slabGenerated, world, gridTrees)

	base64, err := self.encoder.Encode(slabGenerated)

	if err != nil {
		return contracts.Slab{}, apierror.New(http.StatusInternalServerError, err.Error())
	}

	size := float64(len(base64) / 1024)
	sizeStr := fmt.Sprintf("%f Kb", size)

	slabContract := contracts.Slab{
		SlabVersion: "V2",
		Code:        base64,
		Size:        sizeStr,
	}

	return slabContract, nil
}

func generateGround(worldX, worldY int, terrainComplexity float64, mountain *entities.Mountain) [][]uint16 {
	world := gridhelper.TerrainGenerator(worldX, worldY, 2.0, 2.0, terrainComplexity)

	rand.Seed(time.Now().UnixNano())

	iCount := rand.Intn(mountain.RandComplexity) + mountain.MinComplexity

	rand.Seed(time.Now().UnixNano())
	jCount := rand.Intn(mountain.RandComplexity) + mountain.MinComplexity

	for i := 0; i < iCount; i++ {
		for j := 0; j < jCount; j++ {
			rand.Seed(time.Now().UnixNano())
			mountainX := rand.Intn(mountain.RandX) + mountain.MinX

			rand.Seed(time.Now().UnixNano())
			mountainY := rand.Intn(mountain.RandY) + mountain.MinY

			rand.Seed(time.Now().UnixNano())
			gain := float64(rand.Intn(mountain.RandHeight) + mountain.MinHeight)

			generatedMountain := gridhelper.MountainGenerator(mountainX, mountainY, gain)
			world = gridhelper.BuildTerrain(world, generatedMountain)
		}
	}
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
