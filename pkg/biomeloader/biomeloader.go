package biomeloader

import (
	"encoding/json"
	"github.com/johnfercher/taleslab/pkg/assetloader"
	"github.com/johnfercher/taleslab/pkg/grid"
	"github.com/johnfercher/taleslab/pkg/slabloader"
	"github.com/johnfercher/taleslab/pkg/taleslab/domain/entities"
	"io/ioutil"
	"log"
)

type Biome struct {
	BiomeType    entities.BiomeType            `json:"biome_type"`
	GroundBlocks map[grid.ElementType][]string `json:"ground_blocks"`
	PropBlocks   map[grid.ElementType][]string `json:"prop_blocks"`
	StoneWall    string                        `json:"stone_wall"`
}

type BiomeLoader interface {
	SetBiome(entities.BiomeType)
	GetBiome() entities.BiomeType
	GetConstructorKeys() map[grid.ElementType][]string
	GetConstructorAssets(elementType grid.ElementType) []string
	GetConstructor(id string) *assetloader.AssetInfo
	GetPropKeys() map[grid.ElementType][]string
	GetPropAssets(elementType grid.ElementType) []string
	GetProp(id string) *assetloader.AssetInfo
	GetStoneWall() string
}

type biomeLoader struct {
	biomeType   entities.BiomeType
	assetLoader assetloader.AssetLoader
	slabLoader  slabloader.SlabLoader
	biomes      map[entities.BiomeType]Biome
}

func NewBiomeLoader(assetLoader assetloader.AssetLoader, slabLoader slabloader.SlabLoader) *biomeLoader {
	biomeLoader := &biomeLoader{
		assetLoader: assetLoader,
		slabLoader:  slabLoader,
	}

	biomeLoader.loadBiomes()

	return biomeLoader
}

func (self *biomeLoader) GetConstructorKeys() map[grid.ElementType][]string {
	biome := self.biomes[self.biomeType]
	return biome.GroundBlocks
}

func (self *biomeLoader) GetConstructorAssets(elementType grid.ElementType) []string {
	return self.biomes[self.biomeType].GroundBlocks[elementType]
}

func (self *biomeLoader) GetConstructor(id string) *assetloader.AssetInfo {
	return self.assetLoader.GetConstructor(id)
}

func (self *biomeLoader) GetPropKeys() map[grid.ElementType][]string {
	return self.biomes[self.biomeType].PropBlocks
}

func (self *biomeLoader) GetPropAssets(elementType grid.ElementType) []string {
	return self.biomes[self.biomeType].PropBlocks[elementType]
}

func (self *biomeLoader) GetProp(id string) *assetloader.AssetInfo {
	return self.assetLoader.GetProp(id)
}

func (self *biomeLoader) SetBiome(biomeType entities.BiomeType) {
	self.biomeType = biomeType
}

func (self *biomeLoader) GetBiome() entities.BiomeType {
	return self.biomeType
}

func (self *biomeLoader) GetStoneWall() string {
	return self.biomes[self.biomeType].StoneWall
}
func (self *biomeLoader) loadBiomes() {
	bytes, err := ioutil.ReadFile("./config/assets/biomes.json")
	if err != nil {
		log.Fatal(err.Error())
	}

	biomes := []Biome{}

	err = json.Unmarshal(bytes, &biomes)
	if err != nil {
		log.Fatal(err.Error())
	}

	biomeMap := make(map[entities.BiomeType]Biome)

	for _, biome := range biomes {
		biomeMap[biome.BiomeType] = biome
	}

	self.biomes = biomeMap
}
