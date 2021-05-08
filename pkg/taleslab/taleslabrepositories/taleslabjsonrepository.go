package taleslabrepositories

import (
	"encoding/json"
	"github.com/johnfercher/taleslab/pkg/assetloader"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
	"io/ioutil"
	"log"
)

type biomeJsonRepository struct {
	biomeType   taleslabconsts.BiomeType
	assetLoader assetloader.AssetLoader
	biomes      map[taleslabconsts.BiomeType]taleslabentities.Biome
}

func NewBiomeRepository(assetLoader assetloader.AssetLoader) *biomeJsonRepository {
	reposiktory := &biomeJsonRepository{
		assetLoader: assetLoader,
	}

	reposiktory.loadBiomes()

	return reposiktory
}

func (self *biomeJsonRepository) GetConstructorKeys() map[taleslabconsts.ElementType][]string {
	biome := self.biomes[self.biomeType]
	return biome.GroundBlocks
}

func (self *biomeJsonRepository) GetConstructorAssets(elementType taleslabconsts.ElementType) []string {
	return self.biomes[self.biomeType].GroundBlocks[elementType]
}

func (self *biomeJsonRepository) GetConstructor(id string) *assetloader.AssetInfo {
	return self.assetLoader.GetConstructor(id)
}

func (self *biomeJsonRepository) GetPropKeys() map[taleslabconsts.ElementType][]string {
	return self.biomes[self.biomeType].PropBlocks
}

func (self *biomeJsonRepository) GetPropAssets(elementType taleslabconsts.ElementType) []string {
	return self.biomes[self.biomeType].PropBlocks[elementType]
}

func (self *biomeJsonRepository) GetProp(id string) *assetloader.AssetInfo {
	return self.assetLoader.GetProp(id)
}

func (self *biomeJsonRepository) SetBiome(biomeType taleslabconsts.BiomeType) {
	self.biomeType = biomeType
}

func (self *biomeJsonRepository) GetBiome() taleslabconsts.BiomeType {
	return self.biomeType
}

func (self *biomeJsonRepository) GetStoneWall() string {
	return self.biomes[self.biomeType].StoneWall
}
func (self *biomeJsonRepository) loadBiomes() {
	bytes, err := ioutil.ReadFile("./config/assets/biomes.json")
	if err != nil {
		log.Fatal(err.Error())
	}

	biomes := []taleslabentities.Biome{}

	err = json.Unmarshal(bytes, &biomes)
	if err != nil {
		log.Fatal(err.Error())
	}

	biomeMap := make(map[taleslabconsts.BiomeType]taleslabentities.Biome)

	for _, biome := range biomes {
		biomeMap[biome.BiomeType] = biome
	}

	self.biomes = biomeMap
}
