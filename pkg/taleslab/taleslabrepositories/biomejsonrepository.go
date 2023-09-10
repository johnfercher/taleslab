package taleslabrepositories

import (
	"encoding/json"
	"log"
	"os"

	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts/biometype"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts/elementtype"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabentities"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabrepositories"
)

type biomeJSONRepository struct {
	path   string
	biomes map[biometype.BiomeType]*taleslabentities.Biome
}

func NewBiomeRepository(path ...string) taleslabrepositories.BiomeRepository {
	repository := &biomeJSONRepository{}
	if len(path) != 0 {
		repository.path = path[0]
	} else {
		repository.path = "./configs/biomes.json"
	}

	repository.loadBiomes()

	return repository
}

func (b *biomeJSONRepository) GetBiome(biomeType biometype.BiomeType) *taleslabentities.Biome {
	return b.biomes[biomeType]
}

func (b *biomeJSONRepository) loadBiomes() {
	bytes, err := os.ReadFile(b.path)
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
			Reliefs:   b.reliefsArrayToMap(biome.Reliefs),
			StoneWall: biome.StoneWall,
		}
	}

	b.biomes = biomeMap
}

func (b *biomeJSONRepository) reliefsArrayToMap(reliefs []*taleslabentities.Relief) map[elementtype.ElementType]*taleslabentities.Relief {
	reliefMap := make(map[elementtype.ElementType]*taleslabentities.Relief)

	for _, relief := range reliefs {
		reliefMap[elementtype.ElementType(relief.Key)] = relief
	}

	return reliefMap
}
