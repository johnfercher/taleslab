package taleslabservices

import (
	"github.com/johnfercher/taleslab/internal/api/apierror"
	"github.com/johnfercher/taleslab/internal/math"
	"github.com/johnfercher/taleslab/internal/talespireadapter/talespirecoder"
	"github.com/johnfercher/taleslab/pkg/assetloader"
	"github.com/johnfercher/taleslab/pkg/grid"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabrepositories"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabservices"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabmappers"
	"math/rand"
	"net/http"
	"time"
)

type mapBuilder struct {
	biomeRepository          taleslabrepositories.BiomeRepository
	secondaryBiomeRepository taleslabrepositories.BiomeRepository
	encoder                  talespirecoder.Encoder
	props                    *taleslabentities.Props
	ground                   *taleslabentities.Ground
	mountains                *taleslabentities.Mountains
	river                    *taleslabentities.River
	canyon                   *taleslabentities.Canyon
}

func NewMapBuilder(biomeLoader taleslabrepositories.BiomeRepository, secondaryBiomeLoader taleslabrepositories.BiomeRepository, encoder talespirecoder.Encoder) *mapBuilder {
	return &mapBuilder{
		biomeRepository:          biomeLoader,
		encoder:                  encoder,
		secondaryBiomeRepository: secondaryBiomeLoader,
	}
}

func (self *mapBuilder) SetBiome(biomeType taleslabconsts.BiomeType) taleslabservices.MapBuilder {
	self.biomeRepository.SetBiome(biomeType)
	return self
}

func (self *mapBuilder) SetSecondaryBiome(biomeType taleslabconsts.BiomeType) taleslabservices.MapBuilder {
	if biomeType == "" {
		return self
	}

	self.secondaryBiomeRepository.SetBiome(biomeType)
	return self
}

func (self *mapBuilder) SetGround(ground *taleslabentities.Ground) taleslabservices.MapBuilder {
	self.ground = ground
	return self
}

func (self *mapBuilder) SetMountains(mountains *taleslabentities.Mountains) taleslabservices.MapBuilder {
	if mountains == nil {
		return self
	}
	self.mountains = mountains
	return self
}

func (self *mapBuilder) SetRiver(river *taleslabentities.River) taleslabservices.MapBuilder {
	if river != nil {
		self.river = river
	}

	return self
}

func (self *mapBuilder) SetCanyon(canyon *taleslabentities.Canyon) taleslabservices.MapBuilder {
	if canyon != nil {
		self.canyon = canyon
	}

	return self
}

func (self *mapBuilder) SetProps(props *taleslabentities.Props) taleslabservices.MapBuilder {
	self.props = props
	return self
}

func (self *mapBuilder) Build() (string, apierror.ApiError) {
	assets := taleslabentities.Assets{}

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

	var propsGrid [][]taleslabentities.Element

	if self.props != nil {
		propsGrid = grid.GenerateElementGrid(self.ground.Width, self.ground.Length, taleslabentities.Element{ElementType: taleslabconsts.NoneType})

		propsGrid = grid.RandomlyFillEmptyGridSlots(world, propsGrid, self.props.StoneDensity, taleslabconsts.StoneType, func(element taleslabentities.Element) bool {
			// Just to not add stone in an empty grid slot
			return element.ElementType != taleslabconsts.NoneType
		})

		propsGrid = grid.RandomlyFillEmptyGridSlots(world, propsGrid, self.props.TreeDensity, taleslabconsts.TreeType, func(element taleslabentities.Element) bool {
			return element.ElementType == taleslabconsts.GroundType ||
				element.ElementType == taleslabconsts.MountainType ||
				element.ElementType == taleslabconsts.BaseGroundType
		})

		if self.props.MiscDensity != 0 {
			propsGrid = grid.RandomlyFillEmptyGridSlots(world, propsGrid, self.props.MiscDensity, taleslabconsts.MiscType, func(element taleslabentities.Element) bool {
				return element.ElementType == taleslabconsts.GroundType ||
					element.ElementType == taleslabconsts.MountainType ||
					element.ElementType == taleslabconsts.BaseGroundType
			})
		}
	}

	groundBlocks := self.biomeRepository.GetConstructorKeys()
	propKeys := self.biomeRepository.GetPropKeys()

	for key := range groundBlocks {
		assets = self.appendConstructionSlab(key, assets, world)
	}

	for key := range propKeys {
		assets = self.appendPropsToSlab(key, assets, world, propsGrid)
	}

	taleSpireSlab := taleslabmappers.TaleSpireSlabFromAssets(assets)

	base64, err := self.encoder.Encode(taleSpireSlab)
	if err != nil {
		return "", apierror.New(http.StatusInternalServerError, err.Error())
	}

	return base64, nil
}

func (self *mapBuilder) appendConstructionSlab(elementType taleslabconsts.ElementType, assets taleslabentities.Assets,
	gridHeights [][]taleslabentities.Element) taleslabentities.Assets {
	lastStoneWallX := -4
	lastStoneWallY := -4

	for i, array := range gridHeights {
		for j, element := range array {
			if element.ElementType != elementType {
				continue
			}

			prop := self.getBiomeConstructor(i, len(gridHeights), elementType)

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

			for _, assetPart := range prop.AssertParts {
				if element.Height-minValue.Height > 1 && (minValue.ElementType == taleslabconsts.BaseGroundType) {
					if math.Distance(lastStoneWallX, lastStoneWallY, i, j) > 2 {
						lastStoneWallX = i
						lastStoneWallY = j

						for k := int(element.Height); k >= int(minValue.Height); k-- {
							rotation := math.GetRandomRotation(minValue.ElementType == taleslabconsts.BaseGroundType, 2, "stone_wall_rotation")
							randomDistanceY := math.GetRandomValue(2, "y")
							randomDistanceX := math.GetRandomValue(2, "x")

							stoneWallProp := self.biomeRepository.GetProp(self.biomeRepository.GetStoneWall())

							stoneWall := &taleslabentities.Asset{
								Id:         stoneWallProp.AssertParts[0].Id,
								Name:       stoneWallProp.AssertParts[0].Name,
								Dimensions: stoneWallProp.AssertParts[0].Dimensions,
								OffsetZ:    stoneWallProp.AssertParts[0].OffsetZ,
							}

							self.addCoordinates(stoneWall, i+randomDistanceX, j+randomDistanceY, int(k+stoneWall.OffsetZ)/3.0, rotation)
							assets = append(assets, stoneWall)
						}
					}
				} else {
					for k := minValue.Height; k <= element.Height; k++ {
						asset := &taleslabentities.Asset{
							Id:         assetPart.Id,
							Name:       assetPart.Name,
							Dimensions: assetPart.Dimensions,
							OffsetZ:    assetPart.OffsetZ,
						}

						self.addCoordinates(asset, i, j, k+asset.OffsetZ, 768)
						assets = append(assets, asset)
					}
				}
			}
		}
	}

	return assets
}

func (self *mapBuilder) appendPropsToSlab(elementType taleslabconsts.ElementType, assets taleslabentities.Assets,
	gridHeights [][]taleslabentities.Element, gridProps [][]taleslabentities.Element) taleslabentities.Assets {
	for i, array := range gridHeights {
		for j, element := range array {
			if gridProps[i][j].ElementType == elementType {
				prop := self.getBiomeProp(i, len(gridHeights), elementType)

				for id, assetPart := range prop.AssertParts {
					asset := &taleslabentities.Asset{
						Id:         prop.AssertParts[id].Id,
						Name:       prop.AssertParts[id].Name,
						Dimensions: prop.AssertParts[id].Dimensions,
						OffsetZ:    prop.AssertParts[id].OffsetZ,
					}

					rotation := math.GetRandomRotation(true, 5, "props")
					self.addCoordinates(asset, i, j, element.Height+assetPart.OffsetZ, rotation)

					assets = append(assets, asset)
				}
			}
		}
	}

	return assets
}

func (self *mapBuilder) generateMountainsGrid(minHeight int) [][][]taleslabentities.Element {
	mountainsGrid := [][][]taleslabentities.Element{}

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

func (self *mapBuilder) addCoordinates(asset *taleslabentities.Asset, x, y, z, rotation int) {
	asset.Coordinates = &taleslabentities.Vector3d{
		X: x * asset.Dimensions.Width,
		Y: y * asset.Dimensions.Length,
		Z: z * asset.Dimensions.Height,
	}
	asset.Rotation = rotation + (y * asset.Dimensions.Length / 41)
}

func (self *mapBuilder) getBiomeConstructor(i, iMax int, elementType taleslabconsts.ElementType) *assetloader.AssetInfo {
	option := math.GetRandomOption(i, iMax, 6.0)

	if self.secondaryBiomeRepository.GetBiome() == "" {
		option = true
	}

	if option {
		elementsKeys := self.biomeRepository.GetConstructorAssets(elementType)
		elementKey := elementsKeys[math.GetRandomValue(len(elementsKeys), "constructors")]
		return self.biomeRepository.GetConstructor(elementKey)
	}

	elementsKeys := self.secondaryBiomeRepository.GetConstructorAssets(elementType)
	elementKey := elementsKeys[math.GetRandomValue(len(elementsKeys), "constructors")]
	return self.secondaryBiomeRepository.GetConstructor(elementKey)
}

func (self *mapBuilder) getBiomeProp(i, iMax int, elementType taleslabconsts.ElementType) *assetloader.AssetInfo {
	option := math.GetRandomOption(i, iMax, 13.0)

	if self.secondaryBiomeRepository.GetBiome() == "" {
		option = true
	}

	if option {
		elementsKeys := self.biomeRepository.GetPropAssets(elementType)
		elementKey := elementsKeys[math.GetRandomValue(len(elementsKeys), "props")]
		return self.biomeRepository.GetProp(elementKey)
	}

	elementsKeys := self.secondaryBiomeRepository.GetPropAssets(elementType)
	elementKey := elementsKeys[math.GetRandomValue(len(elementsKeys), "props")]
	return self.secondaryBiomeRepository.GetProp(elementKey)
}
