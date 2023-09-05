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
	biomeType biometype.BiomeType
	biomes    map[biometype.BiomeType]*taleslabentities.Biome
}

func NewBiomeRepository() taleslabrepositories.BiomeRepository {
	repository := &biomeJsonRepository{}

	repository.loadBiomes()

	return repository
}

func (self *biomeJsonRepository) GetBiome(biomeType biometype.BiomeType) *taleslabentities.Biome {
	return self.biomes[biomeType]
}

func (self *biomeJsonRepository) loadBiomes() {
	bytes, err := ioutil.ReadFile("./docs/configs/biomes.json")
	if err != nil {
		log.Fatal(err.Error())
	}

	biomes := []*BiomeDao{}

	err = json.Unmarshal(bytes, &biomes)
	if err != nil {
		log.Fatal(err.Error())
	}

	biomeMap := make(map[biometype.BiomeType]*taleslabentities.Biome)

	for _, biome := range biomes {
		biomeMap[biome.Type] = &taleslabentities.Biome{
			Type:      biome.Type,
			Reliefs:   self.reliefsArrayToMap(biome.Reliefs),
			StoneWall: biome.StoneWall,
		}
	}

	self.biomes = biomeMap
}

func (self *biomeJsonRepository) reliefsArrayToMap(reliefs []*taleslabentities.Relief) map[taleslabconsts.ElementType]*taleslabentities.Relief {
	reliefMap := make(map[taleslabconsts.ElementType]*taleslabentities.Relief)

	for _, relief := range reliefs {
		reliefMap[taleslabconsts.ElementType(relief.Key)] = relief
	}

	return reliefMap
}
