package main

import (
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabdomain/taleslabconsts/biometype"
	"github.com/johnfercher/taleslab/pkg/taleslab/taleslabrepositories"
)

func main() {
	props := taleslabrepositories.NewPropRepository()
	biomes := taleslabrepositories.NewBiomeRepository(props)

	biome := biomes.GetBiome(biometype.Beach)
	biome.Print()
}
