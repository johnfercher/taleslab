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
	biomes map[biometype.BiomeType]*taleslabentities.Biome
}

func NewBiomeRepository() taleslabrepositories.BiomeRepository {
	repository := &biomeJSONRepository{}

	repository.loadBiomes()

	return repository
}

func (b *biomeJSONRepository) GetBiome(biomeType biometype.BiomeType) *taleslabentities.Biome {
	return b.biomes[biomeType]
}

func (b *biomeJSONRepository) loadBiomes() {
	bytes, err := os.ReadFile("./configs/biomes.json")
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
