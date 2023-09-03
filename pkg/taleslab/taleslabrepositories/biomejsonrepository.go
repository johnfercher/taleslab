package taleslabrepositories

import (
	"encoding/json"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts/biometype"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabrepositories"
	"io/ioutil"
	"log"
)

type biomeJsonRepository struct {
	biomeType      biometype.BiomeType
	propRepository taleslabrepositories.PropRepository
	biomes         map[biometype.BiomeType]*taleslabentities.Biome
}

func NewBiomeRepository(propRepository taleslabrepositories.PropRepository) taleslabrepositories.BiomeRepository {
	repository := &biomeJsonRepository{
		propRepository: propRepository,
	}

	repository.loadBiomes()

	return repository
}

func (self *biomeJsonRepository) GetBiome(biomeType biometype.BiomeType) *taleslabentities.Biome {
	return self.biomes[biomeType]
}

func (self *biomeJsonRepository) loadBiomes() {
	bytes, err := ioutil.ReadFile("./config/assets/biomes.json")
	if err != nil {
		log.Fatal(err.Error())
	}

	biomes := []*taleslabentities.Biome{}

	err = json.Unmarshal(bytes, &biomes)
	if err != nil {
		log.Fatal(err.Error())
	}

	biomeMap := make(map[biometype.BiomeType]*taleslabentities.Biome)

	for _, biome := range biomes {
		biomeMap[biome.Type] = biome
	}

	self.biomes = biomeMap
}
