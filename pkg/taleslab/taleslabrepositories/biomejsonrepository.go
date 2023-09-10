package taleslabrepositories

import (
	"encoding/json"
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

func NewBiomeRepository(path ...string) (taleslabrepositories.BiomeRepository, error) {
	repository := &biomeJSONRepository{}
	if len(path) != 0 {
		repository.path = path[0]
	} else {
		repository.path = "./configs/biomes.json"
	}

	err := repository.loadBiomes()
	if err != nil {
		return nil, err
	}

	return repository, nil
}

func (b *biomeJSONRepository) GetBiome(biomeType biometype.BiomeType) *taleslabentities.Biome {
	return b.biomes[biomeType]
}

func (b *biomeJSONRepository) GetBiomes() map[biometype.BiomeType]*taleslabentities.Biome {
	return b.biomes
}

func (b *biomeJSONRepository) loadBiomes() error {
	bytes, err := os.ReadFile(b.path)
	if err != nil {
		return err
	}

	biomes := []*BiomeDao{}

	err = json.Unmarshal(bytes, &biomes)
	if err != nil {
		return err
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
	return nil
}

func (b *biomeJSONRepository) reliefsArrayToMap(reliefs []*taleslabentities.Relief) map[elementtype.ElementType]*taleslabentities.Relief {
	reliefMap := make(map[elementtype.ElementType]*taleslabentities.Relief)

	for _, relief := range reliefs {
		reliefMap[elementtype.ElementType(relief.Key)] = relief
	}

	return reliefMap
}
