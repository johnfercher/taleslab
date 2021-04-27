package taleslabservices

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

type mapbuilder struct {
	loader       assetloader.AssetLoader
	encoder      slabdecoder.Encoder
	biome        entities.Biome
	props        *entities.Props
	ground       *entities.Ground
	mountains    *entities.Mountains
	hasRiver     bool
	propsInfo    map[string]assetloader.AssetInfo
	constructors map[string]assetloader.AssetInfo
	// TODO: Improve
	groundBlock string
	treeBlock   string
	stoneBlock  string
}

func New(loader assetloader.AssetLoader, encoder slabdecoder.Encoder) *mapbuilder {
	return &mapbuilder{
		loader:  loader,
		encoder: encoder,
		biome:   entities.ForestBiome,
	}
}

func (self *mapbuilder) SetBiome(biome entities.Biome) services.MapBuilder {
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

func (self *mapbuilder) SetGround(ground *entities.Ground) services.MapBuilder {
	self.ground = ground
	return self
}

func (self *mapbuilder) SetMountains(mountains *entities.Mountains) services.MapBuilder {
	if mountains == nil {
		return self
	}
	self.mountains = mountains
	return self
}

func (self *mapbuilder) SetRiver(river *entities.River) services.MapBuilder {
	if river != nil {
		self.hasRiver = river.HasRiver
	}

	return self
}

func (self *mapbuilder) SetProps(props *entities.Props) services.MapBuilder {
	self.props = props
	return self
}

func (self *mapbuilder) Build() (string, apierror.ApiError) {
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

	gridWidth := self.ground.Width
	gridLength := self.ground.Length

	world := gridhelper.TerrainGenerator(gridWidth, gridLength, 2.0, 2.0, self.ground.TerrainComplexity)

	if self.mountains != nil {
		mountains := self.generateMountainsGrid()
		for _, mountain := range mountains {
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

func (self *mapbuilder) appendGroundToSlab(constructors map[string]assetloader.AssetInfo, generatedSlab *slab.Slab, gridHeights [][]uint16) {
	generatedSlab.AssetsCount++
	generatedSlab.Assets = append(generatedSlab.Assets,
		&slab.Asset{
			Id:         constructors[self.groundBlock].Id,
			Dimensions: constructors[self.groundBlock].Dimensions,
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

func (self *mapbuilder) appendStonesToSlab(generatedSlab *slab.Slab, gridHeights [][]uint16, gridStones [][]bool) {
	generatedSlab.AssetsCount++
	generatedSlab.Assets = append(generatedSlab.Assets,
		&slab.Asset{
			Id:         self.propsInfo[self.stoneBlock].Id,
			Dimensions: self.constructors[self.groundBlock].Dimensions,
		})

	for i, array := range gridHeights {
		for j, element := range array {
			if gridStones[i][j] {
				self.addLayout(generatedSlab.Assets[1], uint16(i), uint16(j), element)
			}
		}
	}
}

func (self *mapbuilder) appendTreesToSlab(generatedSlab *slab.Slab, gridHeights [][]uint16, gridTrees [][]bool) {
	generatedSlab.AssetsCount++
	generatedSlab.Assets = append(generatedSlab.Assets,
		&slab.Asset{
			Id:         self.propsInfo[self.treeBlock].Id,
			Dimensions: self.constructors[self.groundBlock].Dimensions,
		})

	for i, array := range gridHeights {
		for j, element := range array {
			if gridTrees[i][j] && gridHeights[i][j] > 0 {
				self.addLayout(generatedSlab.Assets[2], uint16(i), uint16(j), element+1)
			}
		}
	}
}

func (self *mapbuilder) generateMountainsGrid() [][][]uint16 {
	mountainsGrid := [][][]uint16{}

	rand.Seed(time.Now().UnixNano())

	iCount := rand.Intn(self.mountains.RandComplexity) + self.mountains.MinComplexity

	rand.Seed(time.Now().UnixNano())
	jCount := rand.Intn(self.mountains.RandComplexity) + self.mountains.MinComplexity

	for i := 0; i < iCount; i++ {
		for j := 0; j < jCount; j++ {
			rand.Seed(time.Now().UnixNano())

			balancedWidth := self.mountains.MinX
			balancedRandWidth := rand.Intn(self.mountains.RandX)
			mountainX := balancedWidth + balancedRandWidth

			balancedLength := self.mountains.MinY
			balancedRandLength := rand.Intn(self.mountains.RandY)
			mountainY := balancedLength + balancedRandLength

			rand.Seed(time.Now().UnixNano())
			gain := float64(rand.Intn(self.mountains.RandHeight) + self.mountains.MinHeight)

			generatedMountain := gridhelper.MountainGenerator(mountainX, mountainY, gain)
			mountainsGrid = append(mountainsGrid, generatedMountain)
		}
	}
	return mountainsGrid
}

func (self *mapbuilder) addLayout(asset *slab.Asset, x, y, z uint16) {
	layout := &slab.Bounds{
		Coordinates: &slab.Vector3d{
			X: x * uint16(asset.Dimensions.Width),
			Y: y * uint16(asset.Dimensions.Length),
			Z: z * uint16(asset.Dimensions.Height),
		},
		Rotation: y * uint16(asset.Dimensions.Length) / 41,
	}

	asset.Layouts = append(asset.Layouts, layout)
	asset.LayoutsCount++
}
