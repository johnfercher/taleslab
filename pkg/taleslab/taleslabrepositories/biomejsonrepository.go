package taleslabrepositories

import (
	"encoding/json"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts/biometype"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabrepositories"
	"io/ioutil"
	"log"
)

type biomeJsonRepository struct {
	biomeType      biometype.BiomeType
	propRepository taleslabrepositories.PropRepository
	biomes         map[biometype.BiomeType]taleslabentities.Biome
}

func NewBiomeRepository(propRepository taleslabrepositories.PropRepository) taleslabrepositories.BiomeRepository {
	reposiktory := &biomeJsonRepository{
		propRepository: propRepository,
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

func (self *biomeJsonRepository) GetPropKeys() map[taleslabconsts.ElementType][]string {
	return self.biomes[self.biomeType].PropBlocks
}

func (self *biomeJsonRepository) GetPropAssets(elementType taleslabconsts.ElementType) []string {
	return self.biomes[self.biomeType].PropBlocks[elementType]
}

func (self *biomeJsonRepository) GetProp(id string) *taleslabentities.Prop {
	return self.propRepository.GetProp(id)
}

func (self *biomeJsonRepository) SetBiome(biomeType biometype.BiomeType) {
	self.biomeType = biomeType
}

func (self *biomeJsonRepository) GetBiome() biometype.BiomeType {
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

	biomeMap := make(map[biometype.BiomeType]taleslabentities.Biome)

	for _, biome := range biomes {
		biomeMap[biome.BiomeType] = biome
	}

	self.biomes = biomeMap
}
