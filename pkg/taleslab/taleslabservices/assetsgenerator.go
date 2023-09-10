package taleslabservices

import (
	"fmt"

	"github.com/johnfercher/taleslab/pkg/grid"
	"github.com/johnfercher/taleslab/pkg/rand"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts/biometype"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts/elementtype"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabrepositories"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabservices"
)

type assetsGenerator struct {
	biomeRepository    taleslabrepositories.BiomeRepository
	propsRepository    taleslabrepositories.PropRepository
	biomeType          biometype.BiomeType
	secondaryBiomeType biometype.BiomeType
	MaxWidth           int
	MaxLength          int
	CurrentX           int
	CurrentY           int
	River              bool
}

func NewAssetsGenerator(biomeRepository taleslabrepositories.BiomeRepository, propsRepository taleslabrepositories.PropRepository,
	maxWidth, maxLength int,
) taleslabservices.AssetsGenerator {
	return &assetsGenerator{
		biomeRepository: biomeRepository,
		propsRepository: propsRepository,
		MaxWidth:        maxWidth,
		MaxLength:       maxLength,
	}
}

func (a *assetsGenerator) SetBiome(biomeType biometype.BiomeType) taleslabservices.AssetsGenerator {
	if biomeType == "" {
		return a
	}

	a.biomeType = biomeType
	return a
}

func (a *assetsGenerator) SetSecondaryBiome(biomeType biometype.BiomeType) taleslabservices.AssetsGenerator {
	if biomeType == "" {
		return a
	}

	a.secondaryBiomeType = biomeType
	return a
}

func (a *assetsGenerator) Generate(world [][]taleslabentities.Element, currentX, currentY int) (taleslabentities.Assets, error) {
	a.CurrentX = currentX
	a.CurrentY = currentY

	rotation := 768

	assets := a.generateWorldAssets(world, rotation)
	propsGrid := a.generateDetailAssets(world)

	assets = a.appendPropsToSlab(assets, world, propsGrid)

	return assets, nil
}

func (a *assetsGenerator) generateWorldAssets(world [][]taleslabentities.Element, rotation int) []*taleslabentities.Asset {
	assets := []*taleslabentities.Asset{}

	// X axis
	for i, array := range world {
		// Y axis
		for j, element := range array {
			prop := a.getBiomeBuildingBlock(a.CurrentX+i, a.MaxWidth, element.ElementType)

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
						ID:         assetPart.ID,
						Name:       assetPart.Name,
						Dimensions: assetPart.Dimensions,
						OffsetZ:    assetPart.OffsetZ,
					}

					a.addCoordinates(asset, i, j, k+asset.OffsetZ, rotation)
					assets = append(assets, asset)
				}
			}
		}
	}

	return assets
}

func (a *assetsGenerator) generateDetailAssets(world [][]taleslabentities.Element) [][]taleslabentities.Element {
	worldWidth := len(world)
	worldLength := len(world[0])

	propsGrid := grid.GenerateElementGrid(worldWidth, worldLength, taleslabentities.Element{ElementType: elementtype.None})

	biome := a.biomeRepository.GetBiome(a.biomeType)

	propsKey := []elementtype.ElementType{
		elementtype.Tree,
		elementtype.Stone,
		elementtype.Misc,
	}

	for i := 0; i < worldWidth; i++ {
		for j := 0; j < worldLength; j++ {
			// Avoid to add in limits
			if i == 0 || i == worldWidth-1 || j == 0 || j == worldLength-1 {
				continue
			}

			// Avoid to add to close
			if i > 1 && (propsGrid[i-1][j].ElementType != elementtype.None || propsGrid[i-2][j].ElementType != elementtype.None) {
				continue
			}

			// Avoid to add to close
			if j > 1 && (propsGrid[i][j-1].ElementType != elementtype.None || propsGrid[i][j-2].ElementType != elementtype.None) {
				continue
			}

			for key := range biome.Reliefs {
				if world[i][j].ElementType != key {
					continue
				}

				for _, prop := range propsKey {
					if propsGrid[i][j].ElementType != elementtype.None {
						continue
					}

					maxRand := 100
					weight, _ := biome.GetPropBlockWeight(key, prop)
					random := rand.DifferentIntn(maxRand, fmt.Sprintf("%s-%s-add", key, prop))
					if float64(maxRand)*weight > float64(random) {
						propsGrid[i][j] = taleslabentities.Element{ElementType: prop}
					}
				}
			}
		}
	}

	return propsGrid
}

func (a *assetsGenerator) appendPropsToSlab(assets taleslabentities.Assets,
	world [][]taleslabentities.Element, gridProps [][]taleslabentities.Element,
) taleslabentities.Assets {
	for i, array := range world {
		for j, element := range array {
			reliefType := world[i][j].ElementType
			propType := gridProps[i][j].ElementType

			if propType != elementtype.None {
				prop := a.getBiomeProp(a.CurrentX+i, a.MaxWidth, reliefType, propType)
				if prop == nil {
					continue
				}

				for id, assetPart := range prop.Parts {
					asset := &taleslabentities.Asset{
						ID:         prop.Parts[id].ID,
						Name:       prop.Parts[id].Name,
						Dimensions: prop.Parts[id].Dimensions,
						OffsetZ:    prop.Parts[id].OffsetZ,
					}

					rotation := rand.DifferentRotation(true, 5, "props")
					a.addCoordinates(asset, i, j, element.Height+assetPart.OffsetZ, rotation)

					assets = append(assets, asset)
				}
			}
		}
	}

	return assets
}

func (a *assetsGenerator) addCoordinates(asset *taleslabentities.Asset, x, y, z, rotation int) {
	asset.Coordinates = &taleslabentities.Vector3d{
		X: x * asset.Dimensions.Width,
		Y: y * asset.Dimensions.Length,
		Z: z * asset.Dimensions.Height,
	}
	asset.Rotation = rotation + (y * asset.Dimensions.Length / 41)
}

func (a *assetsGenerator) getBiomeProp(currentXCoordinate, maxXCoordinate int, reliefType elementtype.ElementType,
	propType elementtype.ElementType,
) *taleslabentities.Prop {
	biome := a.biomeRepository.GetBiome(a.biomeType)

	if a.secondaryBiomeType == "" {
		key, _ := biome.GetPropBlockFromElement(reliefType, propType)
		return a.propsRepository.GetProp(key)
	}

	option := rand.Option(currentXCoordinate, maxXCoordinate, 13.0)

	if option {
		key, _ := biome.GetPropBlockFromElement(reliefType, propType)
		return a.propsRepository.GetProp(key)
	}

	biome = a.biomeRepository.GetBiome(a.secondaryBiomeType)
	key, _ := biome.GetPropBlockFromElement(reliefType, propType)
	return a.propsRepository.GetProp(key)
}

func (a *assetsGenerator) getBiomeBuildingBlock(i, iMax int, reliefType elementtype.ElementType) *taleslabentities.Prop {
	biome := a.biomeRepository.GetBiome(a.biomeType)

	if a.secondaryBiomeType == "" {
		key, _ := biome.GetBuildingBlockFromElement(reliefType)
		return a.propsRepository.GetProp(key)
	}

	option := rand.Option(i, iMax, 13.0)

	if option {
		key, _ := biome.GetBuildingBlockFromElement(reliefType)
		return a.propsRepository.GetProp(key)
	}

	biome = a.biomeRepository.GetBiome(a.secondaryBiomeType)
	key, _ := biome.GetBuildingBlockFromElement(reliefType)
	return a.propsRepository.GetProp(key)
}
