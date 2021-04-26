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
	"net/http"
)

type desertService struct {
	loader  assetloader.AssetLoader
	encoder slabdecoder.Encoder
}

func NewDesertService(encoder slabdecoder.Encoder) *desertService {
	return &desertService{
		encoder: encoder,
		loader:  assetloader.NewAssetLoader(),
	}
}

func (self *desertService) GenerateForest(ctx context.Context, forest *entities.Forest) (contracts.Slab, apierror.ApiError) {
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

	gridStones := gridhelper.GenerateRandomGridPositions(forest.X, forest.Y, forest.OrnamentDensity)
	gridTrees := gridhelper.GenerateExclusiveRandomGrid(forest.X, forest.Y, forest.TreeDensity, gridStones)

	self.appendGroundToSlab(constructors, slabGenerated, world, forest)
	self.appendStonesToSlab(ornaments, slabGenerated, world, gridStones)
	self.appendCactusToSlab(ornaments, slabGenerated, world, gridTrees)

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

func (self *desertService) generateGround(forest *entities.Forest) [][]uint16 {
	world := gridhelper.TerrainGenerator(forest.X, forest.Y, 3.0, 3.0, forest.TerrainComplexity)

	if forest.HasRiver {
		world = gridhelper.DigRiver(world)
	}

	return world
}
func (self *desertService) appendStonesToSlab(ornaments map[string]assetloader.AssetInfo, generatedSlab *slab.Slab, gridHeights [][]uint16, gridStones [][]bool) {
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

func (self *desertService) appendCactusToSlab(ornaments map[string]assetloader.AssetInfo, generatedSlab *slab.Slab, gridHeights [][]uint16, gridTrees [][]bool) {
	generatedSlab.AssetsCount++
	generatedSlab.Assets = append(generatedSlab.Assets,
		&slab.Asset{
			Id: ornaments["cactus_small"].Id,
		})

	for i, array := range gridHeights {
		for j, element := range array {
			if gridTrees[i][j] && gridHeights[i][j] > 0 {
				self.addLayout(generatedSlab.Assets[2], uint16(i), uint16(j), element+1)
			}
		}
	}
}

func (self *desertService) appendGroundToSlab(constructors map[string]assetloader.AssetInfo, generatedSlab *slab.Slab, gridHeights [][]uint16, forest *entities.Forest) {
	generatedSlab.AssetsCount++
	generatedSlab.Assets = append(generatedSlab.Assets,
		&slab.Asset{
			Id: constructors["ground_sand_small"].Id,
		})

	for i, array := range gridHeights {
		for j, element := range array {
			if !forest.ForceBaseLand && element == 0 {
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

func (self *desertService) addLayout(asset *slab.Asset, x, y, z uint16) {
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
