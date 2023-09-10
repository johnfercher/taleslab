package taleslabservices

import (
	"errors"
	"fmt"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts/biometype"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdto"

	"github.com/johnfercher/taleslab/pkg/grid"
	"github.com/johnfercher/taleslab/pkg/rand"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts/elementtype"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabrepositories"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabservices"
)

type slabSliceGenerator struct {
	biomeRepository taleslabrepositories.BiomeRepository
	propsRepository taleslabrepositories.PropRepository
}

func NewSlabSliceGenerator(biomeRepository taleslabrepositories.BiomeRepository, propsRepository taleslabrepositories.PropRepository) taleslabservices.SlabSliceGenerator {
	return &slabSliceGenerator{
		biomeRepository: biomeRepository,
		propsRepository: propsRepository,
	}
}

func (a *slabSliceGenerator) Generate(sliceDto *taleslabdto.SliceDto) (*taleslabentities.Slab, error) {
	rotation := 768

	if len(sliceDto.Biomes) == 0 {
		return nil, errors.New("you must provide at least one biome")
	}

	slab := a.generateWorldAssets(sliceDto, rotation)
	propsGrid := a.generateDetailAssets(sliceDto)

	slab = a.appendPropsToSlab(slab, propsGrid, sliceDto)

	return slab, nil
}

func (a *slabSliceGenerator) generateWorldAssets(sliceDto *taleslabdto.SliceDto, rotation int) *taleslabentities.Slab {
	slab := &taleslabentities.Slab{}
	world := sliceDto.World
	offsetX := sliceDto.OffsetX
	maxWidth := sliceDto.FullDimension.Width

	// X axis
	for i, array := range world {
		// Y axis
		for j, element := range array {
			prop := a.getBiomeBuildingBlock(offsetX+i, maxWidth, element.ElementType, sliceDto.Biomes)

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
					slab.Assets = append(slab.Assets, asset)
				}
			}
		}
	}

	return slab
}

func (a *slabSliceGenerator) generateDetailAssets(sliceDto *taleslabdto.SliceDto) [][]taleslabentities.Element {
	worldWidth := sliceDto.SliceDimension.Width
	worldLength := sliceDto.SliceDimension.Length
	world := sliceDto.World

	propsGrid := grid.GenerateElementGrid(worldWidth, worldLength, taleslabentities.Element{ElementType: elementtype.None})

	biome := a.biomeRepository.GetBiome(sliceDto.Biomes[0])

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

func (a *slabSliceGenerator) appendPropsToSlab(slab *taleslabentities.Slab, gridProps [][]taleslabentities.Element, sliceDto *taleslabdto.SliceDto) *taleslabentities.Slab {
	world := sliceDto.World
	offsetX := sliceDto.OffsetX
	maxWidth := sliceDto.FullDimension.Width

	for i, array := range world {
		for j, element := range array {
			reliefType := world[i][j].ElementType
			propType := gridProps[i][j].ElementType

			if propType != elementtype.None {
				prop := a.getBiomeProp(offsetX+i, maxWidth, reliefType, propType, sliceDto.Biomes)
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

					slab.Assets = append(slab.Assets, asset)
				}
			}
		}
	}

	return slab
}

func (a *slabSliceGenerator) addCoordinates(asset *taleslabentities.Asset, x, y, z, rotation int) {
	asset.Coordinates = &taleslabentities.Vector3d{
		X: x * asset.Dimensions.Width,
		Y: y * asset.Dimensions.Length,
		Z: z * asset.Dimensions.Height,
	}
	asset.Rotation = rotation + (y * asset.Dimensions.Length / 41)
}

func (a *slabSliceGenerator) getBiomeProp(currentXCoordinate, maxXCoordinate int, reliefType elementtype.ElementType,
	propType elementtype.ElementType, biomes []biometype.BiomeType,
) *taleslabentities.Prop {
	if len(biomes) == 1 {
		biome := a.biomeRepository.GetBiome(biomes[0])
		key, _ := biome.GetPropBlockFromElement(reliefType, propType)
		return a.propsRepository.GetProp(key)
	}

	option := rand.Option(currentXCoordinate, maxXCoordinate, 13.0)

	if option {
		biome := a.biomeRepository.GetBiome(biomes[0])
		key, _ := biome.GetPropBlockFromElement(reliefType, propType)
		return a.propsRepository.GetProp(key)
	}

	biome := a.biomeRepository.GetBiome(biomes[1])
	key, _ := biome.GetPropBlockFromElement(reliefType, propType)
	return a.propsRepository.GetProp(key)
}

func (a *slabSliceGenerator) getBiomeBuildingBlock(i, iMax int, reliefType elementtype.ElementType, biomes []biometype.BiomeType) *taleslabentities.Prop {
	if len(biomes) == 1 {
		biome := a.biomeRepository.GetBiome(biomes[0])
		key, _ := biome.GetBuildingBlockFromElement(reliefType)
		return a.propsRepository.GetProp(key)
	}

	option := rand.Option(i, iMax, 13.0)

	if option {
		biome := a.biomeRepository.GetBiome(biomes[0])
		key, _ := biome.GetBuildingBlockFromElement(reliefType)
		return a.propsRepository.GetProp(key)
	}

	biome := a.biomeRepository.GetBiome(biomes[1])
	key, _ := biome.GetBuildingBlockFromElement(reliefType)
	return a.propsRepository.GetProp(key)
}
