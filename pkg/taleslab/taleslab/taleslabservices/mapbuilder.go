package taleslabservices

import (
	"github.com/johnfercher/taleslab/internal/api/apierror"
	"github.com/johnfercher/taleslab/pkg/assetloader"
	"github.com/johnfercher/taleslab/pkg/grid"
	"github.com/johnfercher/taleslab/pkg/mappers"
	"github.com/johnfercher/taleslab/pkg/taleslab/domain/entities"
	"github.com/johnfercher/taleslab/pkg/taleslab/domain/services"
	"github.com/johnfercher/taleslab/pkg/talespire/talespirecoder"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type mapBuilder struct {
	loader       assetloader.AssetLoader
	encoder      talespirecoder.Encoder
	biome        entities.Biome
	props        *entities.Props
	ground       *entities.Ground
	mountains    *entities.Mountains
	hasRiver     bool
	propsInfo    map[string]assetloader.AssetInfo
	constructors map[string]assetloader.AssetInfo
	// TODO: Improve
	groundBlocks map[grid.ElementType][]string
	propBlocks   map[grid.ElementType][]string
}

func New(loader assetloader.AssetLoader, encoder talespirecoder.Encoder) *mapBuilder {
	return &mapBuilder{
		loader:       loader,
		encoder:      encoder,
		biome:        entities.ForestBiome,
		groundBlocks: make(map[grid.ElementType][]string),
		propBlocks:   make(map[grid.ElementType][]string),
	}
}

func (self *mapBuilder) SetBiome(biome entities.Biome) services.MapBuilder {
	self.biome = biome

	switch self.biome {
	case entities.DesertBiome:
		self.groundBlocks[grid.GroundType] = []string{"ground_sand_small"}
		self.groundBlocks[grid.MountainType] = []string{"ground_sand_small"}
		self.propBlocks[grid.TreeType] = []string{"cactus_small"}
		self.propBlocks[grid.StoneType] = []string{"stone_big"}
		break
	case entities.TundraBiome:
		self.groundBlocks[grid.GroundType] = []string{"ground_snow_small"}
		self.groundBlocks[grid.MountainType] = []string{"ground_snow_small"}
		self.propBlocks[grid.TreeType] = []string{"snow_pine_tree_big", "snow_pine_tree_big", "dead_tree_big"}
		self.propBlocks[grid.StoneType] = []string{"snow_stone_small"}
		break
	default:
		self.groundBlocks[grid.GroundType] = []string{"ground_nature_small"}
		self.groundBlocks[grid.MountainType] = []string{"ground_nature_small"}
		self.propBlocks[grid.TreeType] = []string{"pine_tree_big"}
		self.propBlocks[grid.StoneType] = []string{"stone_big"}
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

	slabGenerated := entities.NewSlab()

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

		propsGrid = grid.RandomlyFillEmptyGridSlots(world, propsGrid, self.props.PropsDensity, grid.StoneType, func(element grid.Element) bool {
			return element.ElementType != grid.NoneType
		})

		propsGrid = grid.RandomlyFillEmptyGridSlots(world, propsGrid, self.props.TreeDensity, grid.TreeType, func(element grid.Element) bool {
			return element.ElementType == grid.GroundType || element.ElementType == grid.MountainType
		})
	}

	for key, _ := range self.groundBlocks {
		self.appendConstructionSlab(constructors, key, slabGenerated, world)
	}

	for key, _ := range self.propBlocks {
		self.appendPropsToSlab(propsGrid, key, slabGenerated, world)
	}

	taleSpireSlab := mappers.TaleSpireSlabFromEntity(slabGenerated)

	base64, err := self.encoder.Encode(taleSpireSlab)
	if err != nil {
		return "", apierror.New(http.StatusInternalServerError, err.Error())
	}

	return base64, nil
}

func (self *mapBuilder) appendConstructionSlab(constructors map[string]assetloader.AssetInfo, elementType grid.ElementType, generatedSlab *entities.Slab, gridHeights [][]grid.Element) {
	elementMax := len(self.groundBlocks[elementType])
	elementRand := rand.Intn(elementMax)

	asset := &entities.Asset{
		Id:         constructors[self.groundBlocks[elementType][elementRand]].Id,
		Dimensions: constructors[self.groundBlocks[elementType][elementRand]].Dimensions,
	}

	for i, array := range gridHeights {
		for j, element := range array {
			if element.ElementType != elementType {
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
				self.addLayout(asset, uint16(i), uint16(j), k, 768)
			}
		}
	}

	generatedSlab.AddAsset(asset)
}

func (self *mapBuilder) appendPropsToSlab(gridProps [][]grid.Element, elementType grid.ElementType, generatedSlab *entities.Slab, gridHeights [][]grid.Element) {
	assets := []*entities.Asset{}
	elementMax := len(self.propBlocks[elementType])
	elementKeys := self.propBlocks[elementType]

	for _, elementKey := range self.propBlocks[elementType] {
		asset := &entities.Asset{
			Id:         self.propsInfo[elementKey].Id,
			Dimensions: self.propsInfo[elementKey].Dimensions,
		}

		assets = append(assets, asset)
	}

	for i, array := range gridHeights {
		for j, element := range array {
			rand.Seed(time.Now().UnixNano())

			elementRand := rand.Intn(elementMax)
			elementKey := elementKeys[elementRand]
			prop := self.propsInfo[elementKey]
			offsetZ := prop.OffsetZ

			if gridProps[i][j].ElementType == elementType {
				self.addLayout(assets[elementRand], uint16(i), uint16(j), element.Height+offsetZ, 768)
			}
		}
	}

	for _, asset := range assets {
		generatedSlab.AddAsset(asset)
	}
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

func (self *mapBuilder) addLayout(asset *entities.Asset, x, y, z, rotation uint16) {
	layout := &entities.Bounds{
		Coordinates: &entities.Vector3d{
			X: x * uint16(asset.Dimensions.Width),
			Y: y * uint16(asset.Dimensions.Length),
			Z: z * uint16(asset.Dimensions.Height),
		},
		Rotation: rotation + (y * uint16(asset.Dimensions.Length) / 41),
	}

	asset.Layouts = append(asset.Layouts, layout)
}
