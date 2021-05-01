package taleslabservices

import (
	"github.com/johnfercher/taleslab/internal/api/apierror"
	"github.com/johnfercher/taleslab/internal/math"
	"github.com/johnfercher/taleslab/pkg/biomeloader"
	"github.com/johnfercher/taleslab/pkg/grid"
	"github.com/johnfercher/taleslab/pkg/mappers"
	"github.com/johnfercher/taleslab/pkg/taleslab/domain/entities"
	"github.com/johnfercher/taleslab/pkg/taleslab/domain/services"
	"github.com/johnfercher/taleslab/pkg/talespire/talespirecoder"
	"math/rand"
	"net/http"
	"time"
)

type mapBuilder struct {
	biomeLoader biomeloader.BiomeLoader
	encoder     talespirecoder.Encoder
	props       *entities.Props
	ground      *entities.Ground
	mountains   *entities.Mountains
	river       *entities.River
	canyon      *entities.Canyon
}

func New(biomeLoader biomeloader.BiomeLoader, encoder talespirecoder.Encoder) *mapBuilder {
	return &mapBuilder{
		biomeLoader: biomeLoader,
		encoder:     encoder,
	}
}

func (self *mapBuilder) SetBiome(biomeType entities.BiomeType) services.MapBuilder {
	self.biomeLoader.SetBiome(biomeType)
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
		self.river = river
	}

	return self
}

func (self *mapBuilder) SetCanyon(canyon *entities.Canyon) services.MapBuilder {
	if canyon != nil {
		self.canyon = canyon
	}

	return self
}

func (self *mapBuilder) SetProps(props *entities.Props) services.MapBuilder {
	self.props = props
	return self
}

func (self *mapBuilder) Build() (string, apierror.ApiError) {
	slabGenerated := entities.NewSlab()

	if self.ground == nil {
		return "", apierror.New(400, "GroundType must be provided")
	}

	world := grid.TerrainGenerator(self.ground.Width, self.ground.Length, 2.0, 2.0,
		self.ground.TerrainComplexity, self.ground.MinHeight, self.ground.ForceBaseLand)

	if self.mountains != nil {
		mountains := self.generateMountainsGrid(self.ground.MinHeight)
		for _, mountain := range mountains {
			world = grid.BuildTerrain(world, mountain)
		}
	}

	if self.canyon != nil && self.canyon.HasCanyon {
		world = grid.DigCanyon(world, self.canyon.CanyonOffset)
	}

	//grid.PrintHeights(world)

	if self.river != nil && self.river.HasRiver {
		world = grid.DigRiver(world)
	}

	var propsGrid [][]grid.Element

	if self.props != nil {
		propsGrid = grid.GenerateElementGrid(self.ground.Width, self.ground.Length, grid.Element{ElementType: grid.NoneType})

		propsGrid = grid.RandomlyFillEmptyGridSlots(world, propsGrid, self.props.StoneDensity, grid.StoneType, func(element grid.Element) bool {
			return element.ElementType != grid.NoneType
		})

		propsGrid = grid.RandomlyFillEmptyGridSlots(world, propsGrid, self.props.TreeDensity, grid.TreeType, func(element grid.Element) bool {
			return element.ElementType == grid.GroundType || element.ElementType == grid.MountainType
		})

		if self.props.MiscDensity != 0 {
			propsGrid = grid.RandomlyFillEmptyGridSlots(world, propsGrid, self.props.MiscDensity, grid.MiscType, func(element grid.Element) bool {
				return element.ElementType != grid.NoneType
			})
		}
	}

	groundBlocks := self.biomeLoader.GetConstructorKeys()
	propKeys := self.biomeLoader.GetPropKeys()

	for key := range groundBlocks {
		self.appendConstructionSlab(key, slabGenerated, world)
	}

	for key := range propKeys {
		self.appendPropsToSlab(key, slabGenerated, world, propsGrid)
	}

	taleSpireSlab := mappers.TaleSpireSlabFromEntity(slabGenerated)

	base64, err := self.encoder.Encode(taleSpireSlab)
	if err != nil {
		return "", apierror.New(http.StatusInternalServerError, err.Error())
	}

	return base64, nil
}

func (self *mapBuilder) appendConstructionSlab(elementType grid.ElementType, generatedSlab *entities.Slab, gridHeights [][]grid.Element) {
	assets := []*entities.Asset{}

	elementMax := len(self.biomeLoader.GetConstructorAssets(elementType))

	for _, elementKey := range self.biomeLoader.GetConstructorAssets(elementType) {
		prop := self.biomeLoader.GetConstructor(elementKey)
		asset := &entities.Asset{
			Id:         prop.Id,
			Name:       prop.Name,
			Dimensions: prop.Dimensions,
			OffsetZ:    prop.OffsetZ,
		}

		assets = append(assets, asset)
	}

	stoneWallProp := self.biomeLoader.GetProp(self.biomeLoader.GetStoneWall())

	stoneWall := &entities.Asset{
		Id:         stoneWallProp.Id,
		Name:       stoneWallProp.Name,
		Dimensions: stoneWallProp.Dimensions,
		OffsetZ:    stoneWallProp.OffsetZ,
	}

	lastStoneWallX := -4
	lastStoneWallY := -4

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

			if element.Height-minValue.Height > 1 && (minValue.ElementType == grid.BaseGroundType) {
				if math.Distance(lastStoneWallX, lastStoneWallY, i, j) > 2 {
					lastStoneWallX = i
					lastStoneWallY = j

					for k := int(element.Height); k >= int(minValue.Height); k-- {
						rotation := math.GetRandomRotation(minValue.ElementType == grid.BaseGroundType, 2, "stone_wall_rotation")
						randomDistanceY := math.GetRandomValue(2, "y")
						randomDistanceX := math.GetRandomValue(2, "x")

						self.addLayout(stoneWall, i+randomDistanceX, j+randomDistanceY, int(k+stoneWall.OffsetZ)/3.0, rotation)
					}
				}
			} else {
				elementRand := math.GetRandomValue(elementMax, "ground")

				// Use the minimum neighborhood height to fill empty spaces
				for k := minValue.Height; k <= element.Height; k++ {
					self.addLayout(assets[elementRand], i, j, k+assets[elementRand].OffsetZ, 768)
				}
			}

		}
	}

	for _, asset := range assets {
		generatedSlab.AddAsset(asset)
	}

	generatedSlab.AddAsset(stoneWall)
}

func (self *mapBuilder) appendPropsToSlab(elementType grid.ElementType, generatedSlab *entities.Slab, gridHeights [][]grid.Element, gridProps [][]grid.Element) {
	assets := []*entities.Asset{}
	elementMax := len(self.biomeLoader.GetPropAssets(elementType))

	for _, elementKey := range self.biomeLoader.GetPropAssets(elementType) {
		prop := self.biomeLoader.GetProp(elementKey)
		asset := &entities.Asset{
			Id:         prop.Id,
			Name:       prop.Name,
			Dimensions: prop.Dimensions,
			OffsetZ:    prop.OffsetZ,
		}

		assets = append(assets, asset)
	}

	for i, array := range gridHeights {
		for j, element := range array {
			rand.Seed(time.Now().UnixNano())
			elementRand := math.GetRandomValue(elementMax, "props")

			if gridProps[i][j].ElementType == elementType {
				rotation := math.GetRandomRotation(true, 5, "props")
				self.addLayout(assets[elementRand], i, j, element.Height+assets[elementRand].OffsetZ, rotation)
			}
		}
	}

	for _, asset := range assets {
		generatedSlab.AddAsset(asset)
	}
}

func (self *mapBuilder) generateMountainsGrid(minHeight int) [][][]grid.Element {
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

			generatedMountain := grid.MountainGenerator(mountainX, mountainY, gain, minHeight)
			mountainsGrid = append(mountainsGrid, generatedMountain)
		}
	}
	return mountainsGrid
}

func (self *mapBuilder) addLayout(asset *entities.Asset, x, y, z, rotation int) {
	layout := &entities.Bounds{
		Coordinates: &entities.Vector3d{
			X: x * asset.Dimensions.Width,
			Y: y * asset.Dimensions.Length,
			Z: z * asset.Dimensions.Height,
		},
		Rotation: rotation + (y * asset.Dimensions.Length / 41),
	}

	asset.Layouts = append(asset.Layouts, layout)
}
