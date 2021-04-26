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

	world := self.generateGround(forest)

	gridStones := gridhelper.GenerateRandomGridPositions(forest)
	gridTrees := gridhelper.GenerateExclusiveRandomGrid(forest, gridStones)

	self.appendGroundToSlab(constructors, slabGenerated, world, forest)
	self.appendStonesToSlab(ornaments, slabGenerated, world, gridStones)
	self.appendTreesToSlab(ornaments, slabGenerated, world, gridTrees)

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

func (self *forestService) generateGround(forest *entities.Forest) [][]uint16 {
	world := gridhelper.TerrainGenerator(forest.Ground.Width, forest.Ground.Length, 2.0, 2.0, forest.Ground.TerrainComplexity)

	rand.Seed(time.Now().UnixNano())

	iCount := rand.Intn(forest.Mountains.RandComplexity) + forest.Mountains.MinComplexity

	rand.Seed(time.Now().UnixNano())
	jCount := rand.Intn(forest.Mountains.RandComplexity) + forest.Mountains.MinComplexity

	for i := 0; i < iCount; i++ {
		for j := 0; j < jCount; j++ {
			rand.Seed(time.Now().UnixNano())
			mountainX := rand.Intn(forest.Mountains.RandX) + forest.Mountains.MinX

			rand.Seed(time.Now().UnixNano())
			mountainY := rand.Intn(forest.Mountains.RandY) + forest.Mountains.MinY

			rand.Seed(time.Now().UnixNano())
			gain := float64(rand.Intn(forest.Mountains.RandHeight) + forest.Mountains.MinHeight)

			generatedMountain := gridhelper.MountainGenerator(mountainX, mountainY, gain)
			world = gridhelper.BuildTerrain(world, generatedMountain)
		}
	}

	if forest.River.HasRiver {
		world = gridhelper.DigRiver(world)
	}

	return world
}
func (self *forestService) appendStonesToSlab(ornaments map[string]assetloader.AssetInfo, generatedSlab *slab.Slab, gridHeights [][]uint16, gridStones [][]bool) {
	generatedSlab.AssetsCount++
	generatedSlab.Assets = append(generatedSlab.Assets,
		&slab.Asset{
			Id: ornaments["stone_big"].Id,
		})

	for i, array := range gridHeights {
		for j, element := range array {
			if gridStones[i][j] {
				self.addLayout(generatedSlab.Assets[1], uint16(i), uint16(j), element)
			}
		}
	}
}

func (self *forestService) appendTreesToSlab(ornaments map[string]assetloader.AssetInfo, generatedSlab *slab.Slab, gridHeights [][]uint16, gridTrees [][]bool) {
	generatedSlab.AssetsCount++
	generatedSlab.Assets = append(generatedSlab.Assets,
		&slab.Asset{
			Id: ornaments["pine_tree_big"].Id,
		})

	for i, array := range gridHeights {
		for j, element := range array {
			if gridTrees[i][j] && gridHeights[i][j] > 0 {
				self.addLayout(generatedSlab.Assets[2], uint16(i), uint16(j), element+1)
			}
		}
	}
}

func (self *forestService) appendGroundToSlab(constructors map[string]assetloader.AssetInfo, generatedSlab *slab.Slab, gridHeights [][]uint16, forest *entities.Forest) {
	generatedSlab.AssetsCount++
	generatedSlab.Assets = append(generatedSlab.Assets,
		&slab.Asset{
			Id: constructors["ground_nature_small"].Id,
		})

	for i, array := range gridHeights {
		for j, element := range array {
			if !forest.Ground.ForceBaseLand && element == 0 {
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
				self.addLayout(generatedSlab.Assets[0], uint16(i), uint16(j), k)
			}
		}
	}
}

func (self *forestService) addLayout(asset *slab.Asset, x, y, z uint16) {
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
