package forestservices

import (
	"github.com/johnfercher/taleslab/internal/api/apierror"
	"github.com/johnfercher/taleslab/internal/gridhelper"
	"github.com/johnfercher/taleslab/pkg/api/domain/entities"
	"github.com/johnfercher/taleslab/pkg/api/domain/services"
	"github.com/johnfercher/taleslab/pkg/assetloader"
	"github.com/johnfercher/taleslab/pkg/slab"
	"github.com/johnfercher/taleslab/pkg/slabdecoder"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type slabBuilder struct {
	loader       assetloader.AssetLoader
	encoder      slabdecoder.Encoder
	biome        entities.Biome
	props        *entities.Props
	ground       *entities.Ground
	mountains    [][][]uint16
	hasRiver     bool
	propsInfo    map[string]assetloader.AssetInfo
	constructors map[string]assetloader.AssetInfo
	// TODO: Improve
	groundBlock string
	treeBlock   string
	stoneBlock  string
}

func New(loader assetloader.AssetLoader, encoder slabdecoder.Encoder) *slabBuilder {
	return &slabBuilder{
		loader:  loader,
		encoder: encoder,
		biome:   entities.ForestBiome,
	}
}

func (self *slabBuilder) SetBiome(biome entities.Biome) services.SlabBuilder {
	self.biome = biome

	switch self.biome {
	case entities.DesertBiome:
		self.groundBlock = "ground_sand_small"
		self.treeBlock = "cactus_small"
		self.stoneBlock = "stone_big"
		break
	default:
		self.groundBlock = "ground_nature_small"
		self.treeBlock = "pine_tree_big"
		self.stoneBlock = "stone_big"
		break
	}

	return self
}

func (self *slabBuilder) SetGround(ground *entities.Ground) services.SlabBuilder {
	self.ground = ground
	return self
}

func (self *slabBuilder) SetMountains(mountains *entities.Mountains) services.SlabBuilder {
	if mountains == nil {
		return self
	}

	rand.Seed(time.Now().UnixNano())

	iCount := rand.Intn(mountains.RandComplexity) + mountains.MinComplexity

	rand.Seed(time.Now().UnixNano())
	jCount := rand.Intn(mountains.RandComplexity) + mountains.MinComplexity

	for i := 0; i < iCount; i++ {
		for j := 0; j < jCount; j++ {
			rand.Seed(time.Now().UnixNano())
			mountainX := rand.Intn(mountains.RandX) + mountains.MinX

			rand.Seed(time.Now().UnixNano())
			mountainY := rand.Intn(mountains.RandY) + mountains.MinY

			rand.Seed(time.Now().UnixNano())
			gain := float64(rand.Intn(mountains.RandHeight) + mountains.MinHeight)

			generatedMountain := gridhelper.MountainGenerator(mountainX, mountainY, gain)
			self.mountains = append(self.mountains, generatedMountain)
		}
	}

	return self
}

func (self *slabBuilder) SetRiver(river *entities.River) services.SlabBuilder {
	if river != nil {
		self.hasRiver = river.HasRiver
	}

	return self
}

func (self *slabBuilder) SetProps(props *entities.Props) services.SlabBuilder {
	self.props = props
	return self
}

func (self *slabBuilder) Build() (string, apierror.ApiError) {
	constructors, err := self.loader.GetConstructors()
	if err != nil {
		log.Fatalln(err)
	}

	self.constructors = constructors

	propsInfo, err := self.loader.GetProps()
	if err != nil {
		log.Fatalln(err)
	}

	self.propsInfo = propsInfo

	slabGenerated := &slab.Slab{
		MagicBytes: slab.MagicBytes,
		Version:    2,
	}

	if self.ground == nil {
		return "", apierror.New(400, "Ground must be provided")
	}

	world := gridhelper.TerrainGenerator(self.ground.Width, self.ground.Length, 2.0, 2.0, self.ground.TerrainComplexity)

	if self.mountains != nil {
		for _, mountain := range self.mountains {
			world = gridhelper.BuildTerrain(world, mountain)
		}
	}

	if self.hasRiver {
		world = gridhelper.DigRiver(world)
	}

	var gridStones [][]bool
	var gridTrees [][]bool
	if self.props != nil {
		gridStones = gridhelper.GenerateRandomGridPositions(self.ground.Width, self.ground.Length, self.props.PropsDensity)
		gridTrees = gridhelper.GenerateExclusiveRandomGrid(self.ground.Width, self.ground.Length, self.props.TreeDensity, gridStones)
	}

	self.appendGroundToSlab(constructors, slabGenerated, world)
	self.appendStonesToSlab(slabGenerated, world, gridStones)
	self.appendTreesToSlab(slabGenerated, world, gridTrees)

	base64, err := self.encoder.Encode(slabGenerated)
	if err != nil {
		return "", apierror.New(http.StatusInternalServerError, err.Error())
	}

	return base64, nil
}

func (self *slabBuilder) appendGroundToSlab(constructors map[string]assetloader.AssetInfo, generatedSlab *slab.Slab, gridHeights [][]uint16) {
	generatedSlab.AssetsCount++
	generatedSlab.Assets = append(generatedSlab.Assets,
		&slab.Asset{
			Id: constructors[self.groundBlock].Id,
		})

	for i, array := range gridHeights {
		for j, element := range array {
			if !self.ground.ForceBaseLand && element == 0 {
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

func (self *slabBuilder) appendStonesToSlab(generatedSlab *slab.Slab, gridHeights [][]uint16, gridStones [][]bool) {
	generatedSlab.AssetsCount++
	generatedSlab.Assets = append(generatedSlab.Assets,
		&slab.Asset{
			Id: self.propsInfo[self.stoneBlock].Id,
		})

	for i, array := range gridHeights {
		for j, element := range array {
			if gridStones[i][j] {
				self.addLayout(generatedSlab.Assets[1], uint16(i), uint16(j), element)
			}
		}
	}
}

func (self *slabBuilder) appendTreesToSlab(generatedSlab *slab.Slab, gridHeights [][]uint16, gridTrees [][]bool) {
	generatedSlab.AssetsCount++
	generatedSlab.Assets = append(generatedSlab.Assets,
		&slab.Asset{
			Id: self.propsInfo[self.treeBlock].Id,
		})

	for i, array := range gridHeights {
		for j, element := range array {
			if gridTrees[i][j] && gridHeights[i][j] > 0 {
				self.addLayout(generatedSlab.Assets[2], uint16(i), uint16(j), element+1)
			}
		}
	}
}

func (self *slabBuilder) addLayout(asset *slab.Asset, x, y, z uint16) {
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
