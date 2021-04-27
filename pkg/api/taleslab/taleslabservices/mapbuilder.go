package taleslabservices

import (
	"github.com/johnfercher/taleslab/internal/api/apierror"
	"github.com/johnfercher/taleslab/internal/grid"
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

type mapBuilder struct {
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
	groundBlock   string
	mountainBlock string
	treeBlock     string
	stoneBlock    string
}

func New(loader assetloader.AssetLoader, encoder slabdecoder.Encoder) *mapBuilder {
	return &mapBuilder{
		loader:  loader,
		encoder: encoder,
		biome:   entities.ForestBiome,
	}
}

func (self *mapBuilder) SetBiome(biome entities.Biome) services.MapBuilder {
	self.biome = biome

	switch self.biome {
	case entities.DesertBiome:
		self.groundBlock = "ground_sand_small"
		self.mountainBlock = "ground_sand_small"
		self.treeBlock = "cactus_small"
		self.stoneBlock = "stone_big"
		break
	default:
		self.groundBlock = "ground_nature_small"
		self.mountainBlock = "ground_nature_small"
		self.treeBlock = "pine_tree_big"
		self.stoneBlock = "stone_big"
		break
	}

	return self
}

func (self *mapBuilder) SetGround(ground *entities.Ground) services.MapBuilder {
	self.ground = ground
	return self
}

func (self *mapBuilder) SetMountains(mountains *entities.Mountains) services.MapBuilder {
	if mountains == nil {
		return self
	}
	self.mountains = mountains
	return self
}

func (self *mapBuilder) SetRiver(river *entities.River) services.MapBuilder {
	if river != nil {
		self.hasRiver = river.HasRiver
	}

	return self
}

func (self *mapBuilder) SetProps(props *entities.Props) services.MapBuilder {
	self.props = props
	return self
}

func (self *mapBuilder) Build() (string, apierror.ApiError) {
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
		return "", apierror.New(400, "GroundType must be provided")
	}

	world := grid.TerrainGenerator(self.ground.Width, self.ground.Length, 2.0, 2.0,
		self.ground.TerrainComplexity, self.ground.ForceBaseLand)

	if self.mountains != nil {
		mountains := self.generateMountainsGrid()
		for _, mountain := range mountains {
			world = grid.BuildTerrain(world, mountain)
		}
	}

	if self.hasRiver {
		world = grid.DigRiver(world)
	}

	var propsGrid [][]grid.Element

	if self.props != nil {
		propsGrid = grid.GenerateElementGrid(self.ground.Width, self.ground.Length, grid.Element{ElementType: grid.NoneType})
		propsGrid = grid.GenerateRandomGridPositions2(world, propsGrid, self.props.PropsDensity, grid.StoneType, func(element grid.Element) bool {
			return element.ElementType != grid.NoneType
		})
		propsGrid = grid.GenerateRandomGridPositions2(world, propsGrid, self.props.TreeDensity, grid.TreeType, func(element grid.Element) bool {
			return element.ElementType == grid.GroundType || element.ElementType == grid.MountainType
		})
	}

	self.appendGroundToSlab(constructors, slabGenerated, world)
	self.appendTerrainToSlab(constructors, slabGenerated, world)
	self.appendPropsToSlab(slabGenerated, world, propsGrid)

	base64, err := self.encoder.Encode(slabGenerated)
	if err != nil {
		return "", apierror.New(http.StatusInternalServerError, err.Error())
	}

	return base64, nil
}

func (self *mapBuilder) appendTerrainToSlab(constructors map[string]assetloader.AssetInfo, generatedSlab *slab.Slab, gridElements [][]grid.Element) {
	generatedSlab.AssetsCount++
	asset := &slab.Asset{
		Id:         constructors[self.mountainBlock].Id,
		Dimensions: constructors[self.mountainBlock].Dimensions,
	}

	for i, array := range gridElements {
		for j, element := range array {
			if element.Height == 0 || element.ElementType != grid.MountainType {
				continue
			}

			minValue := element

			if i > 0 && gridElements[i-1][j].Height < minValue.Height {
				minValue = gridElements[i-1][j]
			}

			if i < len(gridElements)-1 && gridElements[i+1][j].Height < minValue.Height {
				minValue = gridElements[i+1][j]
			}

			if j > 0 && gridElements[i][j-1].Height < minValue.Height {
				minValue = gridElements[i][j-1]
			}

			if j < len(gridElements[i])-1 && gridElements[i][j+1].Height < minValue.Height {
				minValue = gridElements[i][j+1]
			}

			// Use the minimum neighborhood height to fill empty spaces
			for k := minValue.Height; k <= element.Height; k++ {
				self.addLayout(asset, uint16(i), uint16(j), k)
			}
		}
	}

	generatedSlab.Assets = append(generatedSlab.Assets, asset)
}

func (self *mapBuilder) appendGroundToSlab(constructors map[string]assetloader.AssetInfo, generatedSlab *slab.Slab, gridHeights [][]grid.Element) {
	generatedSlab.AssetsCount++
	asset := &slab.Asset{
		Id:         constructors[self.groundBlock].Id,
		Dimensions: constructors[self.groundBlock].Dimensions,
	}

	for i, array := range gridHeights {
		for j, element := range array {
			if element.ElementType != grid.GroundType {
				continue
			}

			minValue := element

			if i > 0 && gridHeights[i-1][j].Height < minValue.Height {
				minValue = gridHeights[i-1][j]
			}

			if i < len(gridHeights)-1 && gridHeights[i+1][j].Height < minValue.Height {
				minValue = gridHeights[i+1][j]
			}

			if j > 0 && gridHeights[i][j-1].Height < minValue.Height {
				minValue = gridHeights[i][j-1]
			}

			if j < len(gridHeights[i])-1 && gridHeights[i][j+1].Height < minValue.Height {
				minValue = gridHeights[i][j+1]
			}

			// Use the minimum neighborhood height to fill empty spaces
			for k := minValue.Height; k <= element.Height; k++ {
				self.addLayout(asset, uint16(i), uint16(j), k)
			}
		}
	}

	generatedSlab.Assets = append(generatedSlab.Assets, asset)
}

func (self *mapBuilder) appendPropsToSlab(generatedSlab *slab.Slab, gridHeights [][]grid.Element, gridProps [][]grid.Element) {
	generatedSlab.AssetsCount++
	stoneProp := &slab.Asset{
		Id:         self.propsInfo[self.stoneBlock].Id,
		Dimensions: self.propsInfo[self.stoneBlock].Dimensions,
	}

	generatedSlab.AssetsCount++
	treeProp := &slab.Asset{
		Id:         self.propsInfo[self.treeBlock].Id,
		Dimensions: self.propsInfo[self.treeBlock].Dimensions,
	}

	for i, array := range gridHeights {
		for j, element := range array {
			if gridProps[i][j].ElementType == grid.StoneType {
				self.addLayout(stoneProp, uint16(i), uint16(j), element.Height)
				continue
			}

			if gridProps[i][j].ElementType == grid.TreeType {
				self.addLayout(treeProp, uint16(i), uint16(j), element.Height)
				continue
			}
		}
	}

	generatedSlab.Assets = append(generatedSlab.Assets, stoneProp)
	generatedSlab.Assets = append(generatedSlab.Assets, treeProp)
}

func (self *mapBuilder) generateMountainsGrid() [][][]grid.Element {
	mountainsGrid := [][][]grid.Element{}

	rand.Seed(time.Now().UnixNano())
	iCount := rand.Intn(self.mountains.RandComplexity) + self.mountains.MinComplexity

	rand.Seed(time.Now().UnixNano())
	jCount := rand.Intn(self.mountains.RandComplexity) + self.mountains.MinComplexity

	for i := 0; i < iCount; i++ {
		for j := 0; j < jCount; j++ {
			rand.Seed(time.Now().UnixNano())
			bothAxis := rand.Intn(10)

			rand.Seed(time.Now().UnixNano())
			balancedWidth := self.mountains.MinX
			balancedRandWidth := rand.Intn(self.mountains.RandX)
			mountainX := balancedWidth + balancedRandWidth + bothAxis

			rand.Seed(time.Now().UnixNano())
			balancedLength := self.mountains.MinY
			balancedRandLength := rand.Intn(self.mountains.RandY)
			mountainY := balancedLength + balancedRandLength + bothAxis

			rand.Seed(time.Now().UnixNano())
			gain := float64(rand.Intn(self.mountains.RandHeight) + self.mountains.MinHeight)

			generatedMountain := grid.MountainGenerator(mountainX, mountainY, gain)
			mountainsGrid = append(mountainsGrid, generatedMountain)
		}
	}
	return mountainsGrid
}

func (self *mapBuilder) addLayout(asset *slab.Asset, x, y, z uint16) {
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
