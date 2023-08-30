package taleslabservices

import (
	"github.com/johnfercher/taleslab/internal/api/apierror"
	"github.com/johnfercher/taleslab/internal/helper/math"
	"github.com/johnfercher/taleslab/pkg/grid"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabrepositories"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabservices"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdto"
)

type assetsGenerator struct {
	biomeRepository          taleslabrepositories.BiomeRepository
	secondaryBiomeRepository taleslabrepositories.BiomeRepository
	props                    *taleslabdto.PropsDtoRequest
}

func NewAssetsGenerator(biomeLoader taleslabrepositories.BiomeRepository, secondaryBiomeLoader taleslabrepositories.BiomeRepository) taleslabservices.AssetsGenerator {
	return &assetsGenerator{
		biomeRepository:          biomeLoader,
		secondaryBiomeRepository: secondaryBiomeLoader,
	}
}

func (self *assetsGenerator) SetBiome(biomeType taleslabconsts.BiomeType) taleslabservices.AssetsGenerator {
	self.biomeRepository.SetBiome(biomeType)
	return self
}

func (self *assetsGenerator) SetSecondaryBiome(biomeType taleslabconsts.BiomeType) taleslabservices.AssetsGenerator {
	if biomeType == "" {
		return self
	}

	self.secondaryBiomeRepository.SetBiome(biomeType)
	return self
}

func (self *assetsGenerator) SetProps(props *taleslabdto.PropsDtoRequest) taleslabservices.AssetsGenerator {
	self.props = props
	return self
}

func (self *assetsGenerator) Generate(world [][]taleslabentities.Element) (taleslabentities.Assets, apierror.ApiError) {
	worldWidth := len(world)
	worldLength := len(world[0])
	assets := taleslabentities.Assets{}

	var propsGrid [][]taleslabentities.Element

	if self.props != nil {
		propsGrid = grid.GenerateElementGrid(worldWidth, worldLength, taleslabentities.Element{ElementType: taleslabconsts.NoneType})

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

	propKeys := self.biomeRepository.GetPropKeys()
	constructorKeys := self.biomeRepository.GetConstructorKeys()

	for key := range constructorKeys {
		assets = self.appendConstructionSlab(key, assets, world)
	}

	for key := range propKeys {
		assets = self.appendPropsToSlab(key, assets, world, propsGrid)
	}

	return assets, nil
}

func (self *assetsGenerator) appendConstructionSlab(elementType taleslabconsts.ElementType, assets taleslabentities.Assets,
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

			for _, assetPart := range prop.Parts {
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
								Id:         stoneWallProp.Parts[0].Id,
								Name:       stoneWallProp.Parts[0].Name,
								Dimensions: stoneWallProp.Parts[0].Dimensions,
								OffsetZ:    stoneWallProp.Parts[0].OffsetZ,
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

func (self *assetsGenerator) appendPropsToSlab(elementType taleslabconsts.ElementType, assets taleslabentities.Assets,
	gridHeights [][]taleslabentities.Element, gridProps [][]taleslabentities.Element) taleslabentities.Assets {
	for i, array := range gridHeights {
		for j, element := range array {
			if gridProps[i][j].ElementType == elementType {
				prop := self.getBiomeProp(i, len(gridHeights), elementType)

				for id, assetPart := range prop.Parts {
					asset := &taleslabentities.Asset{
						Id:         prop.Parts[id].Id,
						Name:       prop.Parts[id].Name,
						Dimensions: prop.Parts[id].Dimensions,
						OffsetZ:    prop.Parts[id].OffsetZ,
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

func (self *assetsGenerator) addCoordinates(asset *taleslabentities.Asset, x, y, z, rotation int) {
	asset.Coordinates = &taleslabentities.Vector3d{
		X: x * asset.Dimensions.Width,
		Y: y * asset.Dimensions.Length,
		Z: z * asset.Dimensions.Height,
	}
	asset.Rotation = rotation + (y * asset.Dimensions.Length / 41)
}

func (self *assetsGenerator) getBiomeConstructor(i, iMax int, elementType taleslabconsts.ElementType) *taleslabentities.Prop {
	option := math.GetRandomOption(i, iMax, 6.0)

	if self.secondaryBiomeRepository.GetBiome() == "" {
		option = true
	}

	if option {
		elementsKeys := self.biomeRepository.GetConstructorAssets(elementType)
		elementKey := elementsKeys[math.GetRandomValue(len(elementsKeys), "constructors")]
		return self.biomeRepository.GetProp(elementKey)
	}

	elementsKeys := self.secondaryBiomeRepository.GetConstructorAssets(elementType)
	elementKey := elementsKeys[math.GetRandomValue(len(elementsKeys), "constructors")]
	return self.secondaryBiomeRepository.GetProp(elementKey)
}

func (self *assetsGenerator) getBiomeProp(i, iMax int, elementType taleslabconsts.ElementType) *taleslabentities.Prop {
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
