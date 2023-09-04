package taleslabservices

import (
	"fmt"
	"github.com/johnfercher/taleslab/internal/api/apierror"
	"github.com/johnfercher/taleslab/internal/math"
	"github.com/johnfercher/taleslab/pkg/grid"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts/biometype"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabrepositories"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabservices"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdto"
)

type assetsGenerator struct {
	biomeRepository    taleslabrepositories.BiomeRepository
	propsRepository    taleslabrepositories.PropRepository
	props              *taleslabdto.PropsDtoRequest
	biomeType          biometype.BiomeType
	secondaryBiomeType biometype.BiomeType
}

func NewAssetsGenerator(biomeRepository taleslabrepositories.BiomeRepository, propsRepository taleslabrepositories.PropRepository) taleslabservices.AssetsGenerator {
	return &assetsGenerator{
		biomeRepository: biomeRepository,
		propsRepository: propsRepository,
	}
}

func (self *assetsGenerator) SetBiome(biomeType biometype.BiomeType) taleslabservices.AssetsGenerator {
	if biomeType == "" {
		return self
	}

	self.biomeType = biomeType
	return self
}

func (self *assetsGenerator) SetSecondaryBiome(biomeType biometype.BiomeType) taleslabservices.AssetsGenerator {
	if biomeType == "" {
		return self
	}

	self.secondaryBiomeType = biomeType
	return self
}

func (self *assetsGenerator) SetProps(props *taleslabdto.PropsDtoRequest) taleslabservices.AssetsGenerator {
	self.props = props
	return self
}

func (self *assetsGenerator) Generate(world [][]taleslabentities.Element) (taleslabentities.Assets, apierror.ApiError) {
	biome := self.biomeRepository.GetBiome(self.biomeType)
	rotation := 768

	assets := self.generateWorldAssets(world, biome, rotation)
	propsGrid := self.generateDetailAssets(world)

	assets = self.appendPropsToSlab(assets, world, propsGrid)

	return assets, nil
}

func (self *assetsGenerator) generateWorldAssets(world [][]taleslabentities.Element, biome *taleslabentities.Biome, rotation int) []*taleslabentities.Asset {
	assets := []*taleslabentities.Asset{}

	// X axis
	for i, array := range world {
		// Y axis
		for j, element := range array {
			buildBlock, _ := biome.GetBuildingBlockFromElement(element.ElementType)
			prop := self.propsRepository.GetProp(buildBlock)

			minValue := element

			if i > 0 && world[i-1][j].Height < minValue.Height {
				minValue = world[i-1][j]
			}

			if i < len(world)-1 && world[i+1][j].Height < minValue.Height {
				minValue = world[i+1][j]
			}

			if j > 0 && world[i][j-1].Height < minValue.Height {
				minValue = world[i][j-1]
			}

			if j < len(world[i])-1 && world[i][j+1].Height < minValue.Height {
				minValue = world[i][j+1]
			}

			// Asset Layers
			for _, assetPart := range prop.Parts {
				// Fill Gaps
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

	return assets
}

func (self *assetsGenerator) generateDetailAssets(world [][]taleslabentities.Element) [][]taleslabentities.Element {
	worldWidth := len(world)
	worldLength := len(world[0])

	propsGrid := grid.GenerateElementGrid(worldWidth, worldLength, taleslabentities.Element{ElementType: taleslabconsts.None})

	biome := self.biomeRepository.GetBiome(self.biomeType)

	propsKey := []taleslabconsts.ElementType{
		taleslabconsts.Tree,
		taleslabconsts.Stone,
		taleslabconsts.Misc,
	}

	for i := 0; i < worldWidth; i++ {
		for j := 0; j < worldLength; j++ {
			// Avoid to add in limits
			if i == 0 || i == worldWidth-1 || j == 0 || j == worldLength-1 {
				continue
			}

			// Avoid to add to close
			if i > 1 && (propsGrid[i-1][j].ElementType != taleslabconsts.None || propsGrid[i-2][j].ElementType != taleslabconsts.None) {
				continue
			}

			// Avoid to add to close
			if j > 1 && (propsGrid[i][j-1].ElementType != taleslabconsts.None || propsGrid[i][j-2].ElementType != taleslabconsts.None) {
				continue
			}

			for key, _ := range biome.Reliefs {
				if world[i][j].ElementType != key {
					continue
				}

				for _, prop := range propsKey {
					if propsGrid[i][j].ElementType != taleslabconsts.None {
						continue
					}

					maxRand := 100
					weight, _ := biome.GetPropBlockWeight(key, prop)
					random := math.GetRandomValue(maxRand, fmt.Sprintf("%s-%s-add", key, prop))
					if float64(maxRand)*weight > float64(random) {
						propsGrid[i][j] = taleslabentities.Element{ElementType: prop}
					}
				}
			}

		}
	}

	/*propsGrid = grid.RandomlyFillEmptyGridSlots(world, propsGrid, self.props.StoneDensity, taleslabconsts.Stone, func(element taleslabentities.Element) bool {
		// Just to not add stone in an empty grid slot
		return element.ElementType != taleslabconsts.None
	})

	propsGrid = grid.RandomlyFillEmptyGridSlots(world, propsGrid, self.props.TreeDensity, taleslabconsts.Tree, func(element taleslabentities.Element) bool {
		return element.ElementType == taleslabconsts.Ground ||
			element.ElementType == taleslabconsts.Mountain ||
			element.ElementType == taleslabconsts.BaseGround ||
			element.ElementType == taleslabconsts.Water
	})

	if self.props.MiscDensity != 0 {
		propsGrid = grid.RandomlyFillEmptyGridSlots(world, propsGrid, self.props.MiscDensity, taleslabconsts.Misc, func(element taleslabentities.Element) bool {
			return element.ElementType == taleslabconsts.Ground ||
				element.ElementType == taleslabconsts.Mountain ||
				element.ElementType == taleslabconsts.BaseGround ||
				element.ElementType == taleslabconsts.Water
		})
	}*/

	return propsGrid
}

func (self *assetsGenerator) appendPropsToSlab(assets taleslabentities.Assets,
	world [][]taleslabentities.Element, gridProps [][]taleslabentities.Element) taleslabentities.Assets {
	for i, array := range world {
		for j, element := range array {
			reliefType := world[i][j].ElementType
			propType := gridProps[i][j].ElementType

			if propType != taleslabconsts.None {
				prop := self.getBiomeProp(i, len(world), reliefType, propType)
				if prop == nil {
					continue
				}

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

func (self *assetsGenerator) getBiomeProp(i, iMax int, reliefType taleslabconsts.ElementType, propType taleslabconsts.ElementType) *taleslabentities.Prop {
	biome := self.biomeRepository.GetBiome(self.biomeType)

	if self.secondaryBiomeType == "" {
		key, _ := biome.GetPropBlockFromElement(reliefType, propType)
		return self.propsRepository.GetProp(key)
	}

	option := math.GetRandomOption(i, iMax, 13.0)

	if option {
		key, _ := biome.GetPropBlockFromElement(reliefType, propType)
		return self.propsRepository.GetProp(key)
	}

	biome = self.biomeRepository.GetBiome(self.secondaryBiomeType)
	key, _ := biome.GetPropBlockFromElement(reliefType, propType)
	return self.propsRepository.GetProp(key)
}
